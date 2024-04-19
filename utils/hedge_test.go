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
