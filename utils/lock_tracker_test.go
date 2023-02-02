package utils_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/livekit/protocol/utils"
	"github.com/stretchr/testify/assert"
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
	assert.False(t, utils.ScanTrackedLocks(time.Millisecond))

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

	time.Sleep(5 * time.Millisecond)
	assert.True(t, utils.ScanTrackedLocks(time.Millisecond))

	ms[50].Unlock()
}

func TestMutexFinalizer(t *testing.T) {
	cleanupTest()
	assert.Equal(t, 0, utils.NumMutexes())

	go func() {
		m := &utils.Mutex{}
		m.Lock()
		go func() {
			m.Unlock()
		}()
		assert.Equal(t, 1, utils.NumMutexes())
	}()

	time.Sleep(time.Millisecond)
	cleanupTest()

	assert.Equal(t, 0, utils.NumMutexes())
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
	b.Run("native mutex", func(b *testing.B) {
		var m sync.Mutex
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
	b.Run("native rwmutex", func(b *testing.B) {
		var m sync.RWMutex
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
	b.Run("native mutex init", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m sync.Mutex
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
	b.Run("native rwmutex init", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m sync.RWMutex
			m.Lock()
			noop()
			m.Unlock()
		}
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
