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

package webhook

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/frostbyte73/core"
	"github.com/gammazero/deque"
	"github.com/livekit/protocol/livekit"
)

var (
	errQueueBackedUp = errors.New("queue is backed up")
	errQueueFull     = errors.New("queue is full")
)

type item struct {
	ctx      context.Context
	queuedAt time.Time
	event    *livekit.WebhookEvent
}

type resourceQueueParams struct {
	MaxAge   time.Duration
	MaxDepth int

	Poster poster
}

type resourceQueue struct {
	params resourceQueueParams

	mu    sync.Mutex
	items deque.Deque[*item]
	cond  *sync.Cond

	closed core.Fuse
	drain  atomic.Bool
}

func newResourceQueue(params resourceQueueParams) *resourceQueue {
	r := &resourceQueue{
		params: params,
	}
	r.items.SetBaseCap(int(min(params.MaxDepth, 16)))
	r.cond = sync.NewCond(&r.mu)

	go r.worker()
	return r
}

func (r *resourceQueue) Stop(force bool) {
	r.drain.Store(!force)
	r.closed.Break()
	r.cond.Broadcast()
}

func (r *resourceQueue) Enqueue(ctx context.Context, whEvent *livekit.WebhookEvent) error {
	return r.EnqueueAt(ctx, time.Now(), whEvent)
}

func (r *resourceQueue) EnqueueAt(ctx context.Context, at time.Time, whEvent *livekit.WebhookEvent) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.items.Len() >= r.params.MaxDepth {
		return errQueueFull
	}

	if r.items.Len() > 0 {
		front := r.items.Front()
		if time.Since(front.queuedAt) > r.params.MaxAge {
			return errQueueBackedUp
		}
	}

	r.items.PushBack(&item{ctx, at, whEvent})
	r.cond.Broadcast()
	return nil
}

func (r *resourceQueue) flush() {
	r.mu.Lock()
	for r.items.Len() > 0 {
		item := r.items.PopFront()
		r.mu.Unlock()

		if r.params.Poster != nil && item != nil {
			r.params.Poster.Process(item.ctx, item.queuedAt, item.event)
		}

		r.mu.Lock()
	}
	r.mu.Unlock()
}

func (r *resourceQueue) worker() {
	for {
		select {
		case <-r.closed.Watch():
			if r.drain.Load() {
				r.flush()
			}
			return

		default:
			r.mu.Lock()
			if r.items.Len() == 0 {
				r.cond.Wait()
			}

			if r.closed.IsBroken() {
				r.mu.Unlock()
				continue
			}

			var item *item
			if r.items.Len() > 0 {
				item = r.items.PopFront()
			}
			r.mu.Unlock()

			if r.params.Poster != nil && item != nil {
				r.params.Poster.Process(item.ctx, item.queuedAt, item.event)
			}
		}
	}
}
