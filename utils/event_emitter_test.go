package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEventEmitter(t *testing.T) {
	t.Run("emitter", func(t *testing.T) {
		emitter := NewDefaultEventEmitter[string, int]()
		ao0 := emitter.Observe("a")
		ao1 := emitter.Observe("a")
		bo := emitter.Observe("b")

		emitter.Emit("a", 1)
		emitter.Emit("b", 2)
		require.Equal(t, 1, <-ao0.Events())
		require.Equal(t, 1, <-ao1.Events())
		require.Equal(t, 2, <-bo.Events())

		ao1.Stop()
		emitter.Emit("a", 3)
		require.Equal(t, 3, <-ao0.Events())
		select {
		case <-ao1.Events():
			require.Fail(t, "expected no event from stopped observer")
		default:
		}

		keys := emitter.ObservedKeys()
		sort.Strings(keys)
		require.Equal(t, []string{"a", "b"}, keys)
	})

	t.Run("observer", func(t *testing.T) {
		var closeCalled bool
		o, emit := NewEventObserver[int](func() { closeCalled = true })

		emit(1)
		require.Equal(t, 1, <-o.Events())

		o.Stop()
		require.True(t, closeCalled)
	})
}
