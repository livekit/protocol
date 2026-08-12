// Copyright 2026 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"math"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

const lockTrackerMaxReadHolders = 8

// rwLockTracker extends lockTracker with read-holder tracking. The write
// holder uses the embedded single-holder bookkeeping (writes are exclusive);
// readers are concurrent, so their gids live in atomic slots.
//
// Read-holder accounting: trackRLock records the reader's gid in a free
// slot, or increments roverflow when all slots are taken. trackRUnlock
// removes exactly one entry: the caller's own gid if present, else an
// overflow credit, else an arbitrary slot — Go permits unlocking from a
// goroutine other than the locker, so like go-deadlock we keep the count
// consistent at the price of a possibly misattributed entry in that rare
// case. Invariant: occupied slots plus roverflow equals the reader count.
// Typed atomics guarantee 64-bit alignment regardless of field layout.
type rwLockTracker struct {
	lockTracker
	rheld     atomic.Int32
	roverflow atomic.Int32
	rgids     [lockTrackerMaxReadHolders]atomic.Int64
}

func (t *rwLockTracker) trackRLock() {
	atomic.AddInt32(&t.waiting, -1)

	gid := goid.Get()
	claimed := false
	for i := range t.rgids {
		if t.rgids[i].CompareAndSwap(0, gid) {
			claimed = true
			break
		}
	}
	if !claimed {
		t.roverflow.Add(1)
	}

	if t.rheld.Add(1) == 1 {
		atomic.StoreUint32(&t.ts, atomic.LoadUint32(&lowResTime))

		if atomic.LoadUint32(&enableLockTrackerStackTrace) == 1 {
			n := runtime.Callers(2, t.stack[:lockTrackerMaxStackDepth])
			t.stack = t.stack[:n]
		}
	}
}

func (t *rwLockTracker) trackRUnlock() {
	gid := goid.Get()
	removed := false
	for i := range t.rgids {
		if t.rgids[i].CompareAndSwap(gid, 0) {
			removed = true
			break
		}
	}
	for !removed {
		if o := t.roverflow.Load(); o > 0 {
			removed = t.roverflow.CompareAndSwap(o, o-1)
			continue
		}
		removed = true
		for i := range t.rgids {
			if g := t.rgids[i].Load(); g != 0 && t.rgids[i].CompareAndSwap(g, 0) {
				break
			}
		}
	}

	if t.rheld.Add(-1) == 0 {
		atomic.StoreUint32(&t.ts, math.MaxUint32)
	}
}

// newRWLockTracker allocates a tracker without side effects; lazyInitTracker
// registers whichever allocation wins publication.
func newRWLockTracker() *rwLockTracker {
	return &rwLockTracker{
		lockTracker: lockTracker{
			stack: make([]uintptr, lockTrackerMaxStackDepth),
			ts:    math.MaxUint32,
			rw:    true,
		},
	}
}

type RWMutex struct {
	sync.RWMutex
	t *rwLockTracker
}

func (m *RWMutex) Lock() {
	t := lazyInitTracker(&m.t, newRWLockTracker)
	if t != nil {
		t.trackWait()
		defer t.trackLock()
	}
	m.RWMutex.Lock()
}

func (m *RWMutex) Unlock() {
	if t := loadTracker(&m.t); t != nil {
		t.trackUnlock()
	}
	m.RWMutex.Unlock()
}

func (m *RWMutex) RLock() {
	t := lazyInitTracker(&m.t, newRWLockTracker)
	if t != nil {
		t.trackWait()
		defer t.trackRLock()
	}
	m.RWMutex.RLock()
}

func (m *RWMutex) RUnlock() {
	if t := loadTracker(&m.t); t != nil {
		t.trackRUnlock()
	}
	m.RWMutex.RUnlock()
}
