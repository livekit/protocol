package utils

var closedPromiseChan = make(chan struct{})

func init() {
	close(closedPromiseChan)
}

type Promise[T any] struct {
	Result T
	Err    error
	done   chan struct{}
}

func NewPromise[T any]() *Promise[T] {
	return &Promise[T]{
		done: make(chan struct{}),
	}
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
	p.Result = result
	p.Err = err
	close(p.done)
}

func (p *Promise[T]) Done() <-chan struct{} {
	return p.done
}

func (p *Promise[T]) Resolved() bool {
	select {
	case <-p.done:
		return true
	default:
		return false
	}
}
