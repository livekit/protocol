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
