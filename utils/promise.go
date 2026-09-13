package utils

import (
	"context"

	"github.com/livekit/protocol/utils/promise"
)

// Promise must stay an alias rather than a defined type: a promise built under
// either spelling has to satisfy an interface written against the other.
type Promise[T any] = promise.Promise[T]

func NewPromise[T any]() *Promise[T] { return promise.New[T]() }

func GoPromise[T any](f func() (T, error)) *Promise[T] { return promise.Go(f) }

func NewResolvedPromise[T any](result T, err error) *Promise[T] {
	return promise.NewResolved(result, err)
}

func AwaitPromise[T any](ctx context.Context, p *Promise[T]) (T, error) {
	return promise.Await(ctx, p)
}
