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
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestProtoProxy(t *testing.T) {
	t.Run("basics", func(t *testing.T) {
		var (
			refreshInterval = 10 * time.Millisecond
			mu              = &sync.Mutex{}
		)
		proxy, numParticipants, freeze := createTestProxy(refreshInterval, mu)
		defer proxy.Stop()

		mu.Lock()
		{
			select {
			case <-proxy.Updated():
				t.Fatal("should not have received an update")
			default:
			}

			// should not have changed, initial value should persist
			require.EqualValues(t, 0, proxy.Get().NumParticipants)
		}
		mu.Unlock()

		// immediate change
		proxy.MarkDirty(true)
		time.Sleep(refreshInterval)
		mu.Lock()
		{
			require.EqualValues(t, 2, numParticipants.Load())
			require.EqualValues(t, 1, proxy.Get().NumParticipants)

			// queued updates
			proxy.MarkDirty(false)
			select {
			case <-proxy.Updated():
				// consume previous notification
			default:
			}
			require.EqualValues(t, 1, proxy.Get().NumParticipants)

			// freeze and ensure that updates are not triggered
			freeze.Store(true)
		}
		mu.Unlock()

		// freezing and consuming the previous notification to ensure counter does not increase in updateFn
		select {
		case <-proxy.Updated():

		case <-time.After(refreshInterval):
			t.Fatal("should have received an update")
		}

		// possible that ticker was updated while markDirty queued another update
		np := proxy.Get().NumParticipants
		require.GreaterOrEqual(t, int(np), 2)

		// trigger another update, but should not get notification as freeze is in place and the model should not have changed
		mu.Lock()
		{
			proxy.MarkDirty(false)
			select {
			case <-proxy.Updated():
				t.Fatal("should not have received an update")
			default:
				require.EqualValues(t, np, proxy.Get().NumParticipants)
			}
		}
		mu.Unlock()

	})

	t.Run("await next update after marking dirty", func(t *testing.T) {
		var (
			refreshInterval = 10 * time.Millisecond
			mu              = &sync.Mutex{}
		)
		proxy, _, _ := createTestProxy(refreshInterval, mu)
		require.EqualValues(t, 0, proxy.Get().NumParticipants)
		<-proxy.MarkDirty(true)
		require.EqualValues(t, 1, proxy.Get().NumParticipants)
	})

	t.Run("await resolves when proxy is stopped", func(t *testing.T) {
		var (
			refreshInterval = 10 * time.Millisecond
			mu              = &sync.Mutex{}
		)
		proxy, _, _ := createTestProxy(refreshInterval, mu)
		done := proxy.MarkDirty(true)
		proxy.Stop()
		<-done
	})

	t.Run("multiple awaits resolve for one update", func(t *testing.T) {
		var (
			refreshInterval = 10 * time.Millisecond
			mu              = &sync.Mutex{}
		)
		proxy, _, _ := createTestProxy(refreshInterval, mu)
		done0 := proxy.MarkDirty(false)
		done1 := proxy.MarkDirty(true)
		<-done0
		<-done1
		require.EqualValues(t, 1, proxy.Get().NumParticipants)
	})

	t.Run("await resolve when there is no change", func(t *testing.T) {
		proxy := NewProtoProxy(10*time.Millisecond, func() *livekit.Room { return nil })
		done := proxy.MarkDirty(true)
		time.Sleep(100 * time.Millisecond)
		select {
		case <-done:
		default:
			t.FailNow()
		}
	})
}

func createTestProxy(refreshInterval time.Duration, mu *sync.Mutex) (*ProtoProxy[*livekit.Room], *atomic.Uint32, *atomic.Bool) {
	// uses an update func that increments numParticipants each time
	var (
		numParticipants atomic.Uint32
		freeze          atomic.Bool
	)
	updateFn := func() *livekit.Room {
		mu.Lock()
		defer mu.Unlock()

		if !freeze.Load() {
			defer numParticipants.Add(1)
		}
		return &livekit.Room{
			NumParticipants: numParticipants.Load(),
		}
	}
	return NewProtoProxy(refreshInterval, updateFn), &numParticipants, &freeze
}
