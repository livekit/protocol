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

package utils_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils"
)

func TestSynchronizedBasic(t *testing.T) {
	t.Cleanup(cleanupTest)

	s := utils.NewSynchronized(map[string]int{})
	v, h := s.Lock()
	(*v)["a"] = 1
	h.Unlock()

	v, h = s.Lock()
	require.Equal(t, 1, (*v)["a"])
	h.Unlock()
	require.Panics(t, func() { h.Unlock() })
}

func TestSynchronizedStuck(t *testing.T) {
	t.Cleanup(cleanupTest)
	require.Nil(t, utils.ScanTrackedLocks(time.Millisecond))

	s := utils.NewSynchronized(0)
	release := make(chan struct{})
	handoff := make(chan utils.Holder, 1)
	done := make(chan struct{})

	go func() {
		_, h := s.Lock()
		handoff <- h
		parkHoldingLock(release)
	}()
	h := <-handoff
	go func() {
		v, h := s.Lock()
		*v++
		h.Unlock()
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)
	locks := utils.ScanTrackedLocks(time.Millisecond)
	require.NotNil(t, locks)
	require.Len(t, locks[0].HolderGoroutineIDs(), 1)
	require.Equal(t, 1, locks[0].NumGoroutineHeld())
	require.Equal(t, 1, locks[0].NumGoroutineWaiting())

	utils.PopulateHolderStacks(locks)
	require.Contains(t, locks[0].HolderStacks(), "parkHoldingLock(")

	// handle-based release: unlocking from a goroutine other than the locker
	// is exact, and the waiter proceeds
	h.Unlock()
	close(release)
	<-done
	require.Nil(t, utils.ScanTrackedLocks(time.Millisecond))
}

func BenchmarkSynchronized(b *testing.B) {
	s := utils.NewSynchronized(0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v, h := s.Lock()
		*v++
		h.Unlock()
	}
}
