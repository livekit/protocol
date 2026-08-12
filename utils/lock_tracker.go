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
// don't keep their owners alive; finalizers clear the slots.
type weakRefList struct {
	refs []uintptr
	free []int
	next int
}

var weakRefLock sync.Mutex
var weakRefs weakRefList

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

// window returns the next n-element window of refs, wrapping to the start
// when the end is reached.
func (l *weakRefList) window(n int) []uintptr {
	min := l.next
	max := min + n
	if len(l.refs) <= max {
		max = len(l.refs)
		l.next = 0
	} else {
		l.next = max
	}
	return l.refs[min:max]
}

func NumMutexes() int {
	weakRefLock.Lock()
	defer weakRefLock.Unlock()
	return weakRefs.count()
}

// ScanTrackedLocks check all lock trackers
func ScanTrackedLocks(threshold time.Duration) []*StuckLock {
	minTS := uint32(time.Now().Add(-threshold).Unix())

	weakRefLock.Lock()
	defer weakRefLock.Unlock()
	return scanTrackedLocks(weakRefs.refs, minTS)
}

// ScanTrackedLocksI check lock trackers incrementally n at a time
func ScanTrackedLocksI(threshold time.Duration, n int) []*StuckLock {
	minTS := uint32(time.Now().Add(-threshold).Unix())
	if n <= 0 {
		n = 10000
	}

	weakRefLock.Lock()
	defer weakRefLock.Unlock()

	return scanTrackedLocks(weakRefs.window(n), minTS)
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
				stuck = append(stuck, t.toStuckLock(ts, waiting))
			}
		}
	}
	return stuck
}

//go:norace
func (t *lockTracker) toStuckLock(ts uint32, waiting int32) *StuckLock {
	d := &StuckLock{
		stack:   slices.Clone(t.stack),
		ts:      ts,
		waiting: waiting,
	}
	if gid := t.gid; gid != 0 {
		d.gids = append(d.gids, gid)
		d.held = 1
	}
	if t.rw {
		r := (*rwLockTracker)(unsafe.Pointer(t))
		d.held += r.rheld.Load()
		if len(d.gids) == 0 && d.held > 0 {
			d.holderStrength = HolderShared
		}
		for i := range r.rgids {
			// dedupe: a recursive RLock occupies multiple slots with one gid
			if gid := r.rgids[i].Load(); gid != 0 && !slices.Contains(d.gids, gid) {
				d.gids = append(d.gids, gid)
			}
		}
	}
	return d
}

// HolderStrength describes how a stuck lock is held.
type HolderStrength int32

const (
	HolderExclusive HolderStrength = iota // Mutex, RWMutex write lock
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

	// size for the live goroutine count up front: each doubling pass repeats
	// the stop-the-world snapshot, worst exactly when diagnosing a
	// goroutine-heavy stuck process
	buf := make([]byte, min(max(1<<20, runtime.NumGoroutine()*4096), 1<<26))
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
// rw marks trackers that are really rwLockTrackers, which extend this with
// read-holder slots.
type lockTracker struct {
	stack   []uintptr
	ts      uint32
	waiting int32
	gid     int64
	rw      bool
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

// newLockTracker allocates a tracker without side effects; lazyInitTracker
// registers whichever allocation wins publication.
func newLockTracker() *lockTracker {
	return &lockTracker{
		stack: make([]uintptr, lockTrackerMaxStackDepth),
		ts:    math.MaxUint32,
	}
}

// loadTracker returns the tracker, or nil if there is none. Unlock paths use
// it instead of lazyInitTracker: a lock being unlocked was necessarily
// locked first, so the tracker either exists or tracking was disabled at
// Lock time.
func loadTracker[T any](p **T) *T {
	return (*T)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(p))))
}

// lazyInitTracker returns the field's tracker, constructing and registering
// it on first use. construct must be free of side effects: racing
// initializers may each allocate, but only the publishing CAS winner
// registers its tracker; losers' allocations are left to the GC.
func lazyInitTracker[T any](p **T, construct func() *T) *T {
	if !lockTrackerEnabled {
		return nil
	}
	up := (*unsafe.Pointer)(unsafe.Pointer(p))
	if t := atomic.LoadPointer(up); t != nil {
		return (*T)(t)
	}
	t := construct()
	if !atomic.CompareAndSwapPointer(up, nil, unsafe.Pointer(t)) {
		return (*T)(atomic.LoadPointer(up))
	}
	ref := weakRefs.add(unsafe.Pointer(t))
	runtime.SetFinalizer(t, func(*T) {
		weakRefs.remove(ref)
	})
	return t
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
