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

package utils_test

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils"
)

func init() {
	utils.EnableLockTracker()
}

func cleanupTest() {
	runtime.GC()
	time.Sleep(time.Millisecond)
}

func noop() {}

func TestScanTrackedLocks(t *testing.T) {
	t.Cleanup(cleanupTest)
	require.Nil(t, utils.ScanTrackedLocks(time.Millisecond))

	ms := make([]*utils.Mutex, 100)
	for i := range ms {
		m := &utils.Mutex{}
		m.Lock()
		noop()
		m.Unlock()
		ms[i] = m
	}

	go func() {
		ms[50].Lock()
		ms[50].Lock()
	}()

	time.Sleep(100 * time.Millisecond)
	require.NotNil(t, utils.ScanTrackedLocks(time.Millisecond))

	ms[50].Unlock()
}

func TestFirstLockStackTrace(t *testing.T) {
	t.Cleanup(cleanupTest)
	require.Nil(t, utils.ScanTrackedLocks(time.Millisecond))

	utils.ToggleLockTrackerStackTraces(true)
	defer utils.ToggleLockTrackerStackTraces(false)

	m := &utils.Mutex{}

	var deepLock func(n int)
	deepLock = func(n int) {
		if n > 0 {
			deepLock(n - 1)
		} else {
			m.Lock()
		}
	}

	go func() {
		deepLock(5)
		m.Lock()
	}()

	time.Sleep(100 * time.Millisecond)
	locks := utils.ScanTrackedLocks(time.Millisecond)
	require.NotNil(t, locks)
	require.NotEqual(t, "", locks[0].FirstLockedAtStack())

	m.Unlock()
}

func parkHoldingLock(release chan struct{}) {
	<-release
}

func TestHolderStacks(t *testing.T) {
	t.Cleanup(cleanupTest)
	require.Nil(t, utils.ScanTrackedLocks(time.Millisecond))

	m := &utils.RWMutex{}
	release := make(chan struct{})
	done := make(chan struct{})

	var held sync.WaitGroup
	held.Add(2)
	for range 2 {
		go func() {
			m.RLock()
			held.Done()
			parkHoldingLock(release)
			m.RUnlock()
		}()
	}

	held.Wait()
	go func() {
		m.Lock()
		noop()
		m.Unlock()
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)
	locks := utils.ScanTrackedLocks(time.Millisecond)
	require.NotNil(t, locks)
	require.Len(t, locks[0].HolderGoroutineIDs(), 2)
	require.Equal(t, utils.HolderShared, locks[0].HolderStrength())
	require.Equal(t, "", locks[0].HolderStacks())

	utils.PopulateHolderStacks(locks)
	for _, gid := range locks[0].HolderGoroutineIDs() {
		require.Contains(t, locks[0].HolderStacks(), fmt.Sprintf("goroutine %d ", gid))
	}
	require.Equal(t, 2, strings.Count(locks[0].HolderStacks(), "parkHoldingLock("))

	close(release)
	<-done
}

func TestHolderCrossGoroutineUnlock(t *testing.T) {
	t.Cleanup(cleanupTest)
	require.Nil(t, utils.ScanTrackedLocks(time.Millisecond))

	m := &utils.Mutex{}
	locked := make(chan struct{})
	go func() {
		m.Lock()
		close(locked)
	}()
	<-locked
	m.Unlock() // legal: different goroutine than the locker

	release := make(chan struct{})
	held := make(chan struct{})
	done := make(chan struct{})
	go func() {
		m.Lock()
		close(held)
		parkHoldingLock(release)
		m.Unlock()
	}()
	<-held
	go func() {
		m.Lock()
		noop()
		m.Unlock()
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)
	locks := utils.ScanTrackedLocks(time.Millisecond)
	require.NotNil(t, locks)
	// the cross-goroutine unlock must not leak the original locker's slot
	require.Len(t, locks[0].HolderGoroutineIDs(), 1)
	require.Equal(t, utils.HolderExclusive, locks[0].HolderStrength())
	utils.PopulateHolderStacks(locks)
	require.Contains(t, locks[0].HolderStacks(), "parkHoldingLock(")

	close(release)
	<-done
}

func TestHolderOverflow(t *testing.T) {
	t.Cleanup(cleanupTest)
	require.Nil(t, utils.ScanTrackedLocks(time.Millisecond))

	const readers = 10 // more than the 8 holder slots

	m := &utils.RWMutex{}
	release := make(chan struct{})
	done := make(chan struct{})

	var held, finished sync.WaitGroup
	held.Add(readers)
	finished.Add(readers)
	for range readers {
		go func() {
			m.RLock()
			held.Done()
			parkHoldingLock(release)
			m.RUnlock()
			finished.Done()
		}()
	}

	held.Wait()
	go func() {
		m.Lock()
		noop()
		m.Unlock()
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)
	locks := utils.ScanTrackedLocks(time.Millisecond)
	require.NotNil(t, locks)
	require.Len(t, locks[0].HolderGoroutineIDs(), 8)
	require.Equal(t, readers, locks[0].NumGoroutineHeld())

	close(release)
	finished.Wait()
	<-done

	// slot and overflow accounting must drain fully for the next episode
	release2 := make(chan struct{})
	held2 := make(chan struct{})
	done2 := make(chan struct{})
	go func() {
		m.RLock()
		close(held2)
		parkHoldingLock(release2)
		m.RUnlock()
	}()
	<-held2
	go func() {
		m.Lock()
		noop()
		m.Unlock()
		close(done2)
	}()
	time.Sleep(100 * time.Millisecond)
	locks = utils.ScanTrackedLocks(time.Millisecond)
	require.NotNil(t, locks)
	require.Len(t, locks[0].HolderGoroutineIDs(), 1)

	close(release2)
	<-done2
}

func TestRecursiveRLockDedupe(t *testing.T) {
	t.Cleanup(cleanupTest)
	require.Nil(t, utils.ScanTrackedLocks(time.Millisecond))

	m := &utils.RWMutex{}
	release := make(chan struct{})
	held := make(chan struct{})
	done := make(chan struct{})
	go func() {
		m.RLock()
		m.RLock()
		close(held)
		parkHoldingLock(release)
		m.RUnlock()
		m.RUnlock()
	}()
	<-held
	go func() {
		m.Lock()
		noop()
		m.Unlock()
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)
	locks := utils.ScanTrackedLocks(time.Millisecond)
	require.NotNil(t, locks)
	// recursive RLock occupies two slots but is one holder
	require.Len(t, locks[0].HolderGoroutineIDs(), 1)
	require.Equal(t, 2, locks[0].NumGoroutineHeld())

	utils.PopulateHolderStacks(locks)
	require.Equal(t, 1, strings.Count(locks[0].HolderStacks(), "parkHoldingLock("))

	close(release)
	<-done
}

func TestMutexFinalizer(t *testing.T) {
	cleanupTest()
	require.Equal(t, 0, utils.NumMutexes())

	{
		m := &utils.Mutex{}
		m.Lock()
		go func() {
			m.Unlock()
		}()
		require.Equal(t, 1, utils.NumMutexes())
	}

	for range 100 {
		cleanupTest()
		if utils.NumMutexes() == 0 {
			break
		}
	}

	require.Equal(t, 0, utils.NumMutexes())
}

func TestEmbeddedMutex(t *testing.T) {
	t.Cleanup(cleanupTest)

	foo := struct{ m utils.Mutex }{}
	foo.m.Lock()
	noop()
	foo.m.Unlock()

	bar := struct{ utils.Mutex }{}
	bar.Lock()
	noop()
	bar.Unlock()
}

func TestContestedGlobalLock(t *testing.T) {
	t.Cleanup(cleanupTest)

	ms := make([]*utils.Mutex, 100)
	for i := range ms {
		m := &utils.Mutex{}
		m.Lock()
		noop()
		m.Unlock()
		ms[i] = m
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				utils.ScanTrackedLocks(time.Minute)
				wg.Done()
			}()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			var m utils.Mutex
			wg.Add(3)
			for i := 0; i < 3; i++ {
				go func() {
					m.Lock()
					noop()
					m.Unlock()
					wg.Done()
				}()
			}
		}
		wg.Done()
	}()

	wg.Wait()
}

func TestInitRace(t *testing.T) {
	t.Cleanup(cleanupTest)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		var m utils.Mutex
		wg.Add(3)
		done := make(chan struct{})
		for i := 0; i < 3; i++ {
			go func() {
				<-done
				m.Lock()
				noop()
				m.Unlock()
				wg.Done()
			}()
		}
		close(done)
		runtime.Gosched()
	}

	wg.Wait()
}

func BenchmarkLockTracker(b *testing.B) {
	b.Run("wrapped mutex", func(b *testing.B) {
		var m utils.Mutex
		for i := 0; i < b.N; i++ {
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("wrapped rwmutex", func(b *testing.B) {
		var m utils.RWMutex
		for i := 0; i < b.N; i++ {
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("wrapped mutex init", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m utils.Mutex
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("wrapped rwmutex init", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m utils.RWMutex
			m.Lock()
			noop()
			m.Unlock()
		}
	})

	utils.ToggleLockTrackerStackTraces(true)
	b.Run("wrapped mutex + stack trace", func(b *testing.B) {
		var m utils.Mutex
		for i := 0; i < b.N; i++ {
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("wrapped rwmutex + stack trace", func(b *testing.B) {
		var m utils.RWMutex
		for i := 0; i < b.N; i++ {
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("wrapped mutex init + stack trace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m utils.Mutex
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("wrapped rwmutex init + stack trace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m utils.RWMutex
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	utils.ToggleLockTrackerStackTraces(false)

	b.Run("native mutex", func(b *testing.B) {
		var m sync.Mutex
		for i := 0; i < b.N; i++ {
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("native rwmutex", func(b *testing.B) {
		var m sync.RWMutex
		for i := 0; i < b.N; i++ {
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("native mutex init", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m sync.Mutex
			m.Lock()
			noop()
			m.Unlock()
		}
	})
	b.Run("native rwmutex init", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m sync.RWMutex
			m.Lock()
			noop()
			m.Unlock()
		}
	})
}

func BenchmarkLockTrackerParallel(b *testing.B) {
	b.Run("wrapped rwmutex rlock", func(b *testing.B) {
		var m utils.RWMutex
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.RLock()
				noop()
				m.RUnlock()
			}
		})
	})
	b.Run("native rwmutex rlock", func(b *testing.B) {
		var m sync.RWMutex
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.RLock()
				noop()
				m.RUnlock()
			}
		})
	})
}

func BenchmarkGetBlocked(b *testing.B) {
	for n := 100; n <= 1000000; n *= 100 {
		n := n
		b.Run(fmt.Sprintf("serial/%d", n), func(b *testing.B) {
			cleanupTest()

			ms := make([]*utils.Mutex, n)
			for i := range ms {
				m := &utils.Mutex{}
				m.Lock()
				noop()
				m.Unlock()
				ms[i] = m
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				utils.ScanTrackedLocks(time.Minute)
			}
		})
	}
}
