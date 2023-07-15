package utils

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

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
	require.EqualValues(t, 2, proxy.Get().NumParticipants)

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
		func(lhs *livekit.Room, rhs *livekit.Room) bool {
			return proto.Equal(lhs, rhs)
		},
	), &numParticipants, &freeze
}
