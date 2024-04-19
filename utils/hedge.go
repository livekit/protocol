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
