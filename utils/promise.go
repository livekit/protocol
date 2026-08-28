package utils

import (
	"context"
	"sync"
)

var closedPromiseChan = make(chan struct{})

func init() {
	close(closedPromiseChan)
}

type Promise[T any] struct {
	mu     sync.Mutex
	done   chan struct{}
	thens  []func(T, error)
	Result T
	Err    error
}

func NewPromise[T any]() *Promise[T] {
	return &Promise[T]{}
}

func GoPromise[T any](f func() (T, error)) *Promise[T] {
	p := NewPromise[T]()
	go func() { p.Resolve(f()) }()
	return p
}

func NewResolvedPromise[T any](result T, err error) *Promise[T] {
	return &Promise[T]{
		Result: result,
		Err:    err,
		done:   closedPromiseChan,
	}
}

func (p *Promise[T]) Resolve(result T, err error) {
	p.mu.Lock()

	if p.done == closedPromiseChan {
		p.mu.Unlock()
		return
	}

	p.Result = result
	p.Err = err

	if p.done != nil {
		close(p.done)
	}
	p.done = closedPromiseChan

	thens := p.thens
	p.thens = nil
	p.mu.Unlock()

	// invoked unlocked so a callback can re-enter the promise
	for _, f := range thens {
		f(result, err)
	}
}

// Then registers f to run when p resolves, on the goroutine that calls Resolve. If p is
// already resolved, f runs inline on the calling goroutine.
func (p *Promise[T]) Then(f func(T, error)) {
	p.mu.Lock()

	if p.done == closedPromiseChan {
		result, err := p.Result, p.Err
		p.mu.Unlock()
		f(result, err)
		return
	}

	p.thens = append(p.thens, f)
	p.mu.Unlock()
}

func (p *Promise[T]) Done() <-chan struct{} {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.done == nil {
		p.done = make(chan struct{})
	}
	return p.done
}

func (p *Promise[T]) Resolved() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.done == closedPromiseChan
}

func AwaitPromise[T any](ctx context.Context, p *Promise[T]) (T, error) {
	select {
	case <-ctx.Done():
		var v T
		return v, ctx.Err()
	case <-p.Done():
		return p.Result, p.Err
	}
}
