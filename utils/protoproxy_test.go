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
	proxy, numParticipants := createTestProxy()

	select {
	case <-proxy.Updated():
		t.Fatal("should not have received an update")
	default:
	}

	// should not have changed, initial value should persist
	require.EqualValues(t, 0, proxy.Get().NumParticipants)
	require.EqualValues(t, 0, proxy.Get().NumParticipants)

	// immediate change
	proxy.MarkDirty(true)
	time.Sleep(10 * time.Millisecond)

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

	select {
	case <-proxy.Updated():
	case <-time.After(100 * time.Millisecond):
		t.Fatal("should have received an update")
	}
	require.EqualValues(t, 2, proxy.Get().NumParticipants)

	// ensure we didn't leak
	proxy.Stop()
	require.Equal(t, numGoRoutines, runtime.NumGoroutine())
}

func createTestProxy() (*ProtoProxy[*livekit.Room], *atomic.Uint32) {
	// uses an update func that increments numParticipants each time
	var numParticipants atomic.Uint32
	return NewProtoProxy[*livekit.Room](10*time.Millisecond, func() *livekit.Room {
		defer numParticipants.Add(1)
		return &livekit.Room{
			NumParticipants: numParticipants.Load(),
		}
	}), &numParticipants
}
