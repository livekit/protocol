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
	"sync"
	"time"

	"go.uber.org/atomic"
)

type LeakyBucket struct {
	sync.Mutex
	last       time.Time
	sleepFor   time.Duration
	perRequest atomic.Duration
	maxSlack   time.Duration
	clock      Clock
}

func NewLeakyBucket(rateLimit int, slack time.Duration, clock Clock) *LeakyBucket {
	var lb LeakyBucket
	lb.SetRateLimit(rateLimit)
	lb.maxSlack = -1 * time.Duration(slack) * lb.perRequest.Load()
	lb.clock = clock
	return &lb
}

// SetRateLimit sets the underlying rate limit.
// The setting may not be applied immediately.
// SetRateLimit is THREAD SAFE and NON-BLOCKING.
func (lb *LeakyBucket) SetRateLimit(rateLimit int) {
	lb.perRequest.Store(time.Second / time.Duration(rateLimit))
}

// Take blocks to ensure that the time spent between multiple Take calls
// is on average time.Second/rate.
// Take is THREAD SAFE and BLOCKING.
func (lb *LeakyBucket) Take() time.Time {
	lb.Lock()
	defer lb.Unlock()

	now := lb.clock.Now()

	// If this is our first request, then we allow it.
	if lb.last.IsZero() {
		lb.last = now
		return lb.last
	}

	// sleepFor calculates how much time we should sleep based on
	// the perRequest budget and how long the last request took.
	// Since the request may take longer than the budget, this number
	// can get negative, and is summed across requests.
	lb.sleepFor += lb.perRequest.Load() - now.Sub(lb.last)

	// We shouldn't allow sleepFor to get too negative, since it would mean that
	// a service that slowed down a lot for a short period of time would get
	// a much higher RPS following that.
	if lb.sleepFor < lb.maxSlack {
		lb.sleepFor = lb.maxSlack
	}

	// If sleepFor is positive, then we should sleep now.
	if lb.sleepFor > 0 {
		lb.clock.Sleep(lb.sleepFor)
		lb.last = now.Add(lb.sleepFor)
		lb.sleepFor = 0
	} else {
		lb.last = now
	}

	return lb.last
}
