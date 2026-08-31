package utils

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPromise(t *testing.T) {
	t.Run("zero value is usable", func(t *testing.T) {
		var p Promise[bool]

		require.False(t, p.Resolved())
		done := p.Done()
		select {
		case <-done:
			require.FailNow(t, "unresolved done channel should block")
		default:
		}

		p.Resolve(true, nil)

		require.True(t, p.Resolved())
		select {
		case <-done:
		default:
			require.FailNow(t, "resolved done channel should not block")
		}

		require.True(t, p.Result)
	})

	t.Run("promise cannot be resolved twice", func(t *testing.T) {
		p := NewPromise[bool]()
		p.Resolve(false, errors.New("fail"))
		p.Resolve(true, nil)

		require.False(t, p.Result)
		require.Error(t, p.Err)
	})
}

func TestPromiseThen(t *testing.T) {
	t.Run("callback runs on the resolving goroutine", func(t *testing.T) {
		p := NewPromise[int]()

		var (
			gotResult int
			gotErr    error
			called    bool
		)
		p.Then(func(result int, err error) {
			gotResult, gotErr, called = result, err, true
		})
		require.False(t, called, "callback should not run before resolution")

		expErr := errors.New("fail")
		var calledBeforeResolveReturned bool
		resolved := make(chan struct{})
		go func() {
			defer close(resolved)
			p.Resolve(7, expErr)
			calledBeforeResolveReturned = called
		}()

		<-resolved
		require.True(t, calledBeforeResolveReturned, "callback should have run before Resolve returned")
		require.Equal(t, 7, gotResult)
		require.Equal(t, expErr, gotErr)
	})

	t.Run("callback registered after resolution runs inline", func(t *testing.T) {
		p := NewPromise[int]()
		p.Resolve(7, nil)

		var called bool
		p.Then(func(result int, err error) {
			called = true
			require.Equal(t, 7, result)
			require.NoError(t, err)
		})
		require.True(t, called, "callback should run before Then returned")
	})

	t.Run("callbacks run in registration order", func(t *testing.T) {
		p := NewPromise[int]()

		var order []int
		for i := range 3 {
			p.Then(func(int, error) { order = append(order, i) })
		}
		p.Resolve(0, nil)

		require.Equal(t, []int{0, 1, 2}, order)
	})

	t.Run("callbacks run once when resolved twice", func(t *testing.T) {
		p := NewPromise[int]()

		var calls int
		p.Then(func(int, error) { calls++ })

		p.Resolve(1, nil)
		p.Resolve(2, nil)

		require.Equal(t, 1, calls)
	})

	t.Run("zero value and resolved promise accept callbacks", func(t *testing.T) {
		var zero Promise[int]
		var zeroCalled bool
		zero.Then(func(result int, err error) { zeroCalled = true })
		require.False(t, zeroCalled)
		zero.Resolve(7, nil)
		require.True(t, zeroCalled)

		var resolvedCalled bool
		NewResolvedPromise(7, nil).Then(func(result int, err error) {
			resolvedCalled = true
			require.Equal(t, 7, result)
		})
		require.True(t, resolvedCalled)
	})

	t.Run("callback can re-enter the promise", func(t *testing.T) {
		p := NewPromise[int]()

		var (
			resolvedDuringCallback bool
			awaitResult            int
			awaitErr               error
			nested                 bool
		)
		p.Then(func(int, error) {
			resolvedDuringCallback = p.Resolved()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			awaitResult, awaitErr = AwaitPromise(ctx, p)

			p.Then(func(int, error) { nested = true })
		})

		done := make(chan struct{})
		go func() {
			defer close(done)
			p.Resolve(7, nil)
		}()

		select {
		case <-done:
		case <-time.After(5 * time.Second):
			require.FailNow(t, "re-entrant callback deadlocked")
		}
		require.True(t, resolvedDuringCallback)
		require.Equal(t, 7, awaitResult)
		require.NoError(t, awaitErr)
		require.True(t, nested, "callback registered during dispatch should run inline")
	})
}
