// Copyright 2023 LiveKit, Inc.
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
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/petermattis/goid"
	"golang.org/x/exp/slices"
)

const lockTrackerMaxStackDepth = 16

var lockTrackerEnabled = false
var enableLockTrackerOnce sync.Once
var lowResTime uint32 = uint32(time.Now().Unix())
var enableLockTrackerStackTrace uint32

// EnableLockTracker enable lock tracking background worker. This should be
// called during init
func EnableLockTracker() {
	enableLockTrackerOnce.Do(func() {
		lockTrackerEnabled = true
		go updateLowResTime()
	})
}

func ToggleLockTrackerStackTraces(enable bool) {
	var v uint32
	if enable {
		v = 1
	}
	atomic.StoreUint32(&enableLockTrackerStackTrace, v)
}

func updateLowResTime() {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		atomic.StoreUint32(&lowResTime, uint32(t.Unix()))
	}
}

// weakRefList is a registry of tracker pointers held as uintptrs so they
// don't keep their owners alive; finalizers clear the slots. All lists share
// weakRefLock.
type weakRefList struct {
	refs []uintptr
	free []int
}

var weakRefLock sync.Mutex
var mutexRefs, rwRefs weakRefList

func (l *weakRefList) add(p unsafe.Pointer) int {
	weakRefLock.Lock()
	defer weakRefLock.Unlock()
	if fi := len(l.free) - 1; fi >= 0 {
		ref := l.free[fi]
		l.refs[ref] = uintptr(p)
		l.free = l.free[:fi]
		return ref
	}
	l.refs = append(l.refs, uintptr(p))
	return len(l.refs) - 1
}

func (l *weakRefList) remove(ref int) {
	weakRefLock.Lock()
	defer weakRefLock.Unlock()
	l.refs[ref] = 0
	l.free = append(l.free, ref)
}

func (l *weakRefList) count() int {
	return len(l.refs) - len(l.free)
}

func NumMutexes() int {
	weakRefLock.Lock()
	defer weakRefLock.Unlock()
	return mutexRefs.count() + rwRefs.count()
}

// ScanTrackedLocks check all lock trackers
func ScanTrackedLocks(threshold time.Duration) []*StuckLock {
	minTS := uint32(time.Now().Add(-threshold).Unix())

	weakRefLock.Lock()
	defer weakRefLock.Unlock()
	return append(scanTrackedLocks(mutexRefs.refs, minTS), scanRWTrackedLocks(rwRefs.refs, minTS)...)
}

var nextScanMin int
var nextScanMinRW int

// ScanTrackedLocksI check lock trackers incrementally n at a time
func ScanTrackedLocksI(threshold time.Duration, n int) []*StuckLock {
	minTS := uint32(time.Now().Add(-threshold).Unix())
	if n <= 0 {
		n = 10000
	}

	weakRefLock.Lock()
	defer weakRefLock.Unlock()

	window := func(size int, next *int) (int, int) {
		min := *next
		max := min + n
		if size <= max {
			max = size
			*next = 0
		} else {
			*next = max
		}
		return min, max
	}

	min, max := window(len(mutexRefs.refs), &nextScanMin)
	stuck := scanTrackedLocks(mutexRefs.refs[min:max], minTS)
	min, max = window(len(rwRefs.refs), &nextScanMinRW)
	return append(stuck, scanRWTrackedLocks(rwRefs.refs[min:max], minTS)...)
}

//go:norace
//go:nosplit
func scanTrackedLocks(refs []uintptr, minTS uint32) []*StuckLock {
	var stuck []*StuckLock
	for _, ref := range refs {
		if ref != 0 {
			t := (*lockTracker)(unsafe.Pointer(ref))
			ts := atomic.LoadUint32(&t.ts)
			waiting := atomic.LoadInt32(&t.waiting)
			if ts <= minTS && waiting > 0 {
				var gids []int64
				var held int32
				if gid := t.gid; gid != 0 {
					gids = []int64{gid}
					held = 1
				}
				stuck = append(stuck, &StuckLock{
					stack:   slices.Clone(t.stack),
					ts:      ts,
					waiting: waiting,
					held:    held,
					gids:    gids,
				})
			}
		}
	}
	return stuck
}

// HolderStrength describes how a stuck lock is held.
type HolderStrength int32

const (
	HolderExclusive HolderStrength = iota // Mutex, RWMutex write lock, Synchronized
	HolderShared                          // RWMutex read locks
)

func (s HolderStrength) String() string {
	if s == HolderShared {
		return "shared"
	}
	return "exclusive"
}

type StuckLock struct {
	stack          []uintptr
	ts             uint32
	waiting        int32
	held           int32
	holderStrength HolderStrength
	gids           []int64
	holderStacks   []string
}

func (d *StuckLock) FirstLockedAtStack() string {
	fs := runtime.CallersFrames(d.stack)
	var b strings.Builder
	for {
		f, ok := fs.Next()
		if !ok {
			break
		}
		if f.Function != "" {
			b.WriteString(f.Function)
			b.WriteByte('\n')
		}
		if f.File != "" {
			b.WriteByte('\t')
			b.WriteString(f.File)
			b.WriteByte(':')
			b.WriteString(strconv.Itoa(f.Line))
			b.WriteByte('\n')
		}
	}
	return b.String()
}

func (d *StuckLock) HeldSince() time.Time {
	return time.Unix(int64(d.ts), 0)
}

func (d *StuckLock) NumGoroutineHeld() int {
	return int(d.held)
}

func (d *StuckLock) NumGoroutineWaiting() int {
	return int(d.waiting)
}

// HolderGoroutineIDs returns the ids of the goroutines holding the lock: the
// single holder for a Mutex, the writer or up to 8 concurrent readers for an
// RWMutex (compare with NumGoroutineHeld to detect overflow).
func (d *StuckLock) HolderGoroutineIDs() []int64 {
	return d.gids
}

// HolderStrength reports whether the lock is held exclusively or shared by
// readers.
func (d *StuckLock) HolderStrength() HolderStrength {
	return d.holderStrength
}

// HolderStacks returns the holder goroutines' stacks as resolved by
// PopulateHolderStacks, or "" if not resolved.
func (d *StuckLock) HolderStacks() string {
	return strings.Join(d.holderStacks, "\n\n")
}

// PopulateHolderStacks resolves the current stack of each stuck lock's holder
// goroutines from a single snapshot of all goroutine stacks. The snapshot
// stops the world; call it once per detection, not per lock.
func PopulateHolderStacks(stuck []*StuckLock) {
	byGID := make(map[int64][]*StuckLock, len(stuck))
	for _, d := range stuck {
		for _, gid := range d.gids {
			byGID[gid] = append(byGID[gid], d)
		}
	}
	if len(byGID) == 0 {
		return
	}

	buf := make([]byte, 1<<20)
	for {
		n := runtime.Stack(buf, true)
		if n < len(buf) || len(buf) >= 1<<26 {
			buf = buf[:n]
			break
		}
		buf = make([]byte, len(buf)*2)
	}

	for _, g := range strings.Split(string(buf), "\n\n") {
		if gid, ok := parseGoroutineHeader(g); ok {
			for _, d := range byGID[gid] {
				d.holderStacks = append(d.holderStacks, g)
			}
		}
	}
}

// parseGoroutineHeader extracts the goroutine id from a stack block formatted
// like "goroutine 123 [chan receive]:\n..."
func parseGoroutineHeader(g string) (int64, bool) {
	const prefix = "goroutine "
	if !strings.HasPrefix(g, prefix) {
		return 0, false
	}
	rest := g[len(prefix):]
	sp := strings.IndexByte(rest, ' ')
	if sp <= 0 {
		return 0, false
	}
	gid, err := strconv.ParseInt(rest[:sp], 10, 64)
	return gid, err == nil
}

// lockTracker tracks an exclusive lock. trackLock and trackUnlock run while
// the lock is held, so the single holder's bookkeeping needs no atomics;
// only the waiter count, maintained outside the lock, does. The scanner
// reads holder state without the lock and tolerates racy reads (go:norace).
type lockTracker struct {
	stack   []uintptr
	ts      uint32
	waiting int32
	gid     int64
	ref     int
}

func (t *lockTracker) trackWait() {
	atomic.AddInt32(&t.waiting, 1)
}

func (t *lockTracker) trackLock() {
	atomic.AddInt32(&t.waiting, -1)
	t.gid = goid.Get()
	atomic.StoreUint32(&t.ts, atomic.LoadUint32(&lowResTime))

	if atomic.LoadUint32(&enableLockTrackerStackTrace) == 1 {
		n := runtime.Callers(2, t.stack[:lockTrackerMaxStackDepth])
		t.stack = t.stack[:n]
	}
}

func (t *lockTracker) trackUnlock() {
	t.gid = 0
	atomic.StoreUint32(&t.ts, math.MaxUint32)
}

func newLockTracker() *lockTracker {
	t := &lockTracker{
		stack: make([]uintptr, lockTrackerMaxStackDepth),
		ts:    math.MaxUint32,
	}
	t.ref = mutexRefs.add(unsafe.Pointer(t))
	runtime.SetFinalizer(t, func(t *lockTracker) {
		mutexRefs.remove(t.ref)
	})
	return t
}

//go:linkname sync_runtime_canSpin sync.runtime_canSpin
func sync_runtime_canSpin(int) bool

//go:linkname sync_runtime_doSpin sync.runtime_doSpin
func sync_runtime_doSpin()

// trackerInitSentinel marks a tracker pointer field as mid-initialization.
var trackerInitSentinel = unsafe.Pointer(new(int))

// loadTracker returns the tracker if its initialization completed, else nil.
// Unlock paths use it instead of lazyInitTracker: a lock being unlocked was
// necessarily locked first, so the tracker either exists or tracking was
// disabled at Lock time.
func loadTracker[T any](p **T) *T {
	t := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(p)))
	if t == trackerInitSentinel {
		return nil
	}
	return (*T)(t)
}

func lazyInitTracker[T any](p **T, construct func() *T) *T {
	if !lockTrackerEnabled {
		return nil
	}
	up := (*unsafe.Pointer)(unsafe.Pointer(p))
	iter := 0
	for {
		if t := atomic.LoadPointer(up); t == nil {
			if atomic.CompareAndSwapPointer(up, nil, trackerInitSentinel) {
				atomic.StorePointer(up, unsafe.Pointer(construct()))
			}
		} else if t == trackerInitSentinel {
			if sync_runtime_canSpin(iter) {
				sync_runtime_doSpin()
				iter++
			} else {
				runtime.Gosched()
			}
		} else {
			return (*T)(t)
		}
	}
}

type Mutex struct {
	sync.Mutex
	t *lockTracker
}

func (m *Mutex) Lock() {
	t := lazyInitTracker(&m.t, newLockTracker)
	if t != nil {
		t.trackWait()
		defer t.trackLock()
	}
	m.Mutex.Lock()
}

func (m *Mutex) Unlock() {
	if t := loadTracker(&m.t); t != nil {
		t.trackUnlock()
	}
	m.Mutex.Unlock()
}
