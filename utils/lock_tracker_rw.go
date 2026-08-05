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
	"unsafe"

	"github.com/petermattis/goid"
	"golang.org/x/exp/slices"
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
type rwLockTracker struct {
	lockTracker
	rheld     int32
	roverflow int32
	rgids     [lockTrackerMaxReadHolders]int64
}

func (t *rwLockTracker) trackRLock() {
	atomic.AddInt32(&t.waiting, -1)

	gid := goid.Get()
	claimed := false
	for i := range t.rgids {
		if atomic.CompareAndSwapInt64(&t.rgids[i], 0, gid) {
			claimed = true
			break
		}
	}
	if !claimed {
		atomic.AddInt32(&t.roverflow, 1)
	}

	if atomic.AddInt32(&t.rheld, 1) == 1 {
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
		if atomic.CompareAndSwapInt64(&t.rgids[i], gid, 0) {
			removed = true
			break
		}
	}
	for !removed {
		if o := atomic.LoadInt32(&t.roverflow); o > 0 {
			removed = atomic.CompareAndSwapInt32(&t.roverflow, o, o-1)
			continue
		}
		removed = true
		for i := range t.rgids {
			if g := atomic.LoadInt64(&t.rgids[i]); g != 0 && atomic.CompareAndSwapInt64(&t.rgids[i], g, 0) {
				break
			}
		}
	}

	if atomic.AddInt32(&t.rheld, -1) == 0 {
		atomic.StoreUint32(&t.ts, math.MaxUint32)
	}
}

func newRWLockTracker() *rwLockTracker {
	t := &rwLockTracker{
		lockTracker: lockTracker{
			stack: make([]uintptr, lockTrackerMaxStackDepth),
			ts:    math.MaxUint32,
		},
	}
	t.ref = rwRefs.add(unsafe.Pointer(t))
	runtime.SetFinalizer(t, func(t *rwLockTracker) {
		rwRefs.remove(t.ref)
	})
	return t
}

//go:norace
//go:nosplit
func scanRWTrackedLocks(refs []uintptr, minTS uint32) []*StuckLock {
	var stuck []*StuckLock
	for _, ref := range refs {
		if ref != 0 {
			t := (*rwLockTracker)(unsafe.Pointer(ref))
			ts := atomic.LoadUint32(&t.ts)
			waiting := atomic.LoadInt32(&t.waiting)
			if ts <= minTS && waiting > 0 {
				var gids []int64
				strength := HolderShared
				held := atomic.LoadInt32(&t.rheld)
				if gid := t.gid; gid != 0 {
					gids = append(gids, gid)
					held++
					strength = HolderExclusive
				}
				for i := range t.rgids {
					if gid := atomic.LoadInt64(&t.rgids[i]); gid != 0 {
						gids = append(gids, gid)
					}
				}
				stuck = append(stuck, &StuckLock{
					stack:          slices.Clone(t.stack),
					ts:             ts,
					waiting:        waiting,
					held:           held,
					holderStrength: strength,
					gids:           gids,
				})
			}
		}
	}
	return stuck
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
