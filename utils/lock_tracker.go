package utils

import (
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var lockTrackerEnabled = false
var enableLockTrackerOnce sync.Once
var lowResTime uint32 = uint32(time.Now().Unix())

// EnableLockTracker enable lock tracking background worker. This should be
// called during init
func EnableLockTracker() {
	enableLockTrackerOnce.Do(func() {
		lockTrackerEnabled = true
		go updateLowResTime()
	})
}

func updateLowResTime() {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		atomic.StoreUint32(&lowResTime, uint32(t.Unix()))
	}
}

var weakRefs []uintptr
var weakRefFree []int
var weakRefLock sync.Mutex

func NumMutexes() int {
	weakRefLock.Lock()
	defer weakRefLock.Unlock()
	return len(weakRefs) - len(weakRefFree)
}

// ScanTrackedLocks check all lock trackers
func ScanTrackedLocks(threshold time.Duration) bool {
	minTS := uint32(time.Now().Add(-threshold).Unix())

	weakRefLock.Lock()
	defer weakRefLock.Unlock()
	return scanTrackedLocks(weakRefs, minTS)
}

// ScanTrackedLocksP check all lock trackers in parallel
func ScanTrackedLocksP(threshold time.Duration) bool {
	minTS := uint32(time.Now().Add(-threshold).Unix())

	weakRefLock.Lock()
	defer weakRefLock.Unlock()

	var found uint32

	k := runtime.NumCPU()
	n := (len(weakRefs) + (k - 1)) / k

	var wg sync.WaitGroup
	wg.Add(k)
	for i := 0; i < k; i++ {
		min := n * i
		max := n * (i + 1)
		go func() {
			for _, ref := range weakRefs[min:max] {
				if ref != 0 {
					t := (*lockTracker)(unsafe.Pointer(ref))
					ts := atomic.LoadUint32(&t.ts)
					if ts <= minTS {
						atomic.StoreUint32(&found, 1)
						return
					}
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return found == 1
}

var nextScanMin int

// ScanTrackedLocksI check lock trackers incrementally n at a time
func ScanTrackedLocksI(threshold time.Duration, n int) bool {
	minTS := uint32(time.Now().Add(-threshold).Unix())
	if n <= 0 {
		n = 10000
	}

	weakRefLock.Lock()
	defer weakRefLock.Unlock()

	min := nextScanMin
	max := nextScanMin + n
	if rl := len(weakRefs); rl <= max {
		max = rl
		nextScanMin = 0
	} else {
		nextScanMin = max
	}

	return scanTrackedLocks(weakRefs[min:max], minTS)
}

func scanTrackedLocks(refs []uintptr, minTS uint32) bool {
	for _, ref := range weakRefs {
		if ref != 0 {
			t := (*lockTracker)(unsafe.Pointer(ref))
			ts := atomic.LoadUint32(&t.ts)
			if ts <= minTS {
				return true
			}
		}
	}
	return false
}

type lockTracker struct {
	ts  uint32
	ref int
}

func (t *lockTracker) trackWait() {
	if t != nil {
		atomic.StoreUint32(&t.ts, atomic.LoadUint32(&lowResTime))
	}
}

func (t *lockTracker) trackLock() {
	if t != nil {
		atomic.StoreUint32(&t.ts, math.MaxUint32)
	}
}

func newLockTracker() *lockTracker {
	t := &lockTracker{}

	runtime.SetFinalizer(t, finalizeLockTracker)

	weakRefLock.Lock()
	defer weakRefLock.Unlock()

	if fi := len(weakRefFree) - 1; fi >= 0 {
		t.ref = weakRefFree[fi]
		weakRefs[t.ref] = uintptr((unsafe.Pointer)(t))
		weakRefFree = weakRefFree[:fi]
	} else {
		t.ref = len(weakRefs)
		weakRefs = append(weakRefs, uintptr((unsafe.Pointer)(t)))
	}

	return t
}

func finalizeLockTracker(t *lockTracker) {
	weakRefLock.Lock()
	defer weakRefLock.Unlock()

	weakRefs[t.ref] = 0
	weakRefFree = append(weakRefFree, t.ref)
}

var waiting lockTracker

//go:linkname sync_runtime_canSpin sync.runtime_canSpin
func sync_runtime_canSpin(int) bool

//go:linkname sync_runtime_doSpin sync.runtime_doSpin
func sync_runtime_doSpin()

func lazyInitLockTracker(p **lockTracker) {
	if !lockTrackerEnabled {
		return
	}
	up := (*unsafe.Pointer)(unsafe.Pointer(p))
	iter := 0
	for {
		if t := atomic.LoadPointer(up); t == nil {
			if atomic.CompareAndSwapPointer(up, nil, (unsafe.Pointer)(&waiting)) {
				atomic.StorePointer(up, (unsafe.Pointer)(newLockTracker()))
				return
			}
		} else if t == (unsafe.Pointer)(&waiting) {
			if sync_runtime_canSpin(iter) {
				sync_runtime_doSpin()
				iter++
			} else {
				runtime.Gosched()
			}
		} else {
			return
		}
	}
}

type Mutex struct {
	sync.Mutex
	t *lockTracker
}

func (m *Mutex) Lock() {
	lazyInitLockTracker(&m.t)
	m.t.trackWait()
	m.Mutex.Lock()
	m.t.trackLock()
}

type RWMutex struct {
	sync.RWMutex
	t *lockTracker
}

func (m *RWMutex) Lock() {
	lazyInitLockTracker(&m.t)
	m.t.trackWait()
	m.RWMutex.Lock()
	m.t.trackLock()
}

func (m *RWMutex) RLock() {
	lazyInitLockTracker(&m.t)
	m.t.trackWait()
	m.RWMutex.RLock()
	m.t.trackLock()
}
