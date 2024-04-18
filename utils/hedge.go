package utils

import (
	"context"
	"time"

	"go.uber.org/multierr"
)

type HedgeParams[T any] struct {
	Timeout       time.Duration
	RetryDelay    time.Duration
	MaxAttempts   int
	IsRecoverable func(err error) bool
	Func          func(context.Context) (T, error)
}

// race retries if the function takes to long to return
// |---------------- attempt 1 ----------------|
// |    delay    |--------- attempt 2 ---------|
func HedgeCall[T any](ctx context.Context, params HedgeParams[T]) (v T, err error) {
	ctx, cancel := context.WithTimeout(ctx, params.Timeout)
	defer cancel()

	type result struct {
		value T
		err   error
	}
	ch := make(chan result, params.MaxAttempts)

	race := func() {
		value, err := params.Func(ctx)
		ch <- result{value, err}
	}

	var attempt int
	delay := time.NewTimer(0)
	defer delay.Stop()

	for {
		select {
		case <-delay.C:
			go race()
			if attempt++; attempt < params.MaxAttempts {
				delay.Reset(params.RetryDelay)
			}
		case res := <-ch:
			if res.err == nil || params.IsRecoverable == nil || !params.IsRecoverable(res.err) {
				return res.value, res.err
			}
			err = multierr.Append(err, res.err)
		case <-ctx.Done():
			err = multierr.Append(err, ctx.Err())
			return
		}
	}
}
