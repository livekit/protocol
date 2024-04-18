package utils

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
)

func TestHedgeCall(t *testing.T) {
	var attempts atomic.Uint32
	res, err := HedgeCall(context.Background(), HedgeParams[uint32]{
		Timeout:     200 * time.Millisecond,
		RetryDelay:  50 * time.Millisecond,
		MaxAttempts: 2,
		Func: func(context.Context) (uint32, error) {
			n := attempts.Add(1)
			time.Sleep(75 * time.Millisecond)
			return n, nil
		},
	})
	require.NoError(t, err)
	require.EqualValues(t, 1, res)
	require.EqualValues(t, 2, attempts.Load())
}
