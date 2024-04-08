// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// EDIT: slight modification to allow setting rate limit on the fly
// SCOPE: LeakyBucket
package utils

import (
	"time"

	"go.uber.org/atomic"
)

type LeakyBucket struct {
	//lint:ignore U1000 Padding is unused but it is crucial to maintain performance
	// of this rate limiter in case of collocation with other frequently accessed memory.
	prepadding [64]byte     // cache line size = 64; created to avoid false sharing.
	state      atomic.Int64 // unix nanoseconds of the next permissions issue.
	//lint:ignore U1000 like prepadding.
	postpadding [56]byte // cache line size - state size = 64 - 8; created to avoid false sharing.

	perRequest time.Duration
	maxSlack   time.Duration
	clock      Clock
}

func NewLeakyBucket(rateLimit int, slack time.Duration, clock Clock) *LeakyBucket {
	var lb LeakyBucket
	lb.SetRateLimit(rateLimit)
	lb.maxSlack = slack * lb.perRequest
	lb.clock = clock
	return &lb
}

func (lb *LeakyBucket) SetRateLimit(rateLimit int) {
	lb.perRequest = time.Second / time.Duration(rateLimit)
}

// Take blocks to ensure that the time spent between multiple
// Take calls is on average time.Second/rate.
func (lb *LeakyBucket) Take() time.Time {
	var (
		newTimeOfNextPermissionIssue int64
		now                          int64
	)
	for {
		now = lb.clock.Now().UnixNano()
		timeOfNextPermissionIssue := lb.state.Load()

		switch {
		case timeOfNextPermissionIssue == 0 || (lb.maxSlack == 0 && now-timeOfNextPermissionIssue > int64(lb.perRequest)):
			// if this is our first call or t.maxSlack == 0 we need to shrink issue time to now
			newTimeOfNextPermissionIssue = now
		case lb.maxSlack > 0 && now-timeOfNextPermissionIssue > int64(lb.maxSlack)+int64(lb.perRequest):
			// a lot of nanoseconds passed since the last Take call
			// we will limit max accumulated time to maxSlack
			newTimeOfNextPermissionIssue = now - int64(lb.maxSlack)
		default:
			// calculate the time at which our permission was issued
			newTimeOfNextPermissionIssue = timeOfNextPermissionIssue + int64(lb.perRequest)
		}

		if lb.state.CompareAndSwap(timeOfNextPermissionIssue, newTimeOfNextPermissionIssue) {
			break
		}
	}

	sleepDuration := time.Duration(newTimeOfNextPermissionIssue - now)
	if sleepDuration > 0 {
		lb.clock.Sleep(sleepDuration)
		return time.Unix(0, newTimeOfNextPermissionIssue)
	}
	// return now if we don't sleep as atomicLimiter does
	return time.Unix(0, now)
}
