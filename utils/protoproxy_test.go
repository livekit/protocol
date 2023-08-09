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
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestProtoProxy(t *testing.T) {
	numGoRoutines := runtime.NumGoroutine()
	proxy, numParticipants, freeze := createTestProxy()

	select {
	case <-proxy.Updated():
		t.Fatal("should not have received an update")
	default:
	}

	// should not have changed, initial value should persist
	require.EqualValues(t, 0, proxy.Get().NumParticipants)

	// immediate change
	proxy.MarkDirty(true)
	time.Sleep(100 * time.Millisecond)

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
	// freezing and consuming the previous notification to ensure counter does not increase in updateFn
	select {
	case <-proxy.Updated():
	case <-time.After(100 * time.Millisecond):
		t.Fatal("should have received an update")
	}
	// possible that ticker was updated while markDirty queued another update
	require.GreaterOrEqual(t, int(proxy.Get().NumParticipants), 2)

	// trigger another update, but should not get notification as freeze is in place and the model should not have changed
	proxy.MarkDirty(false)
	time.Sleep(100 * time.Millisecond)
	select {
	case <-proxy.Updated():
		t.Fatal("should not have received an update")
	default:
	}
	require.EqualValues(t, 2, proxy.Get().NumParticipants)

	// ensure we didn't leak
	proxy.Stop()

	for i := 0; i < 10; i++ {
		if numGoRoutines <= runtime.NumGoroutine() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	require.LessOrEqual(t, runtime.NumGoroutine(), numGoRoutines)
}

func createTestProxy() (*ProtoProxy[*livekit.Room], *atomic.Uint32, *atomic.Bool) {
	// uses an update func that increments numParticipants each time
	var numParticipants atomic.Uint32
	var freeze atomic.Bool
	return NewProtoProxy[*livekit.Room](
		10*time.Millisecond,
		func() *livekit.Room {
			if !freeze.Load() {
				defer numParticipants.Add(1)
			}
			return &livekit.Room{
				NumParticipants: numParticipants.Load(),
			}
		},
	), &numParticipants, &freeze
}
