package utils

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils/promise"
)

// A defined type here would compile but break every caller that mixes the two
// spellings, which no test of behaviour would catch.
var _ *Promise[int] = (*promise.Promise[int])(nil)

func TestPromiseAlias(t *testing.T) {
	p := NewPromise[int]()
	promise.Await(t.Context(), promise.NewResolved(0, nil))
	p.Resolve(7, nil)

	result, err := AwaitPromise(t.Context(), p)
	require.NoError(t, err)
	require.Equal(t, 7, result)
}
