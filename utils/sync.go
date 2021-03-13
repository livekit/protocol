package utils

import (
	"errors"
	"sync/atomic"
)

var (
	ErrChannelClosed = errors.New("cannot write to closed channel")
)

type AtomicFlag struct {
	val int32
}

// set flag to value if existing flag is different, otherwise return
func (b *AtomicFlag) TrySet(bVal bool) bool {
	var v int32
	if bVal {
		v = 1
	}
	prev := atomic.SwapInt32(&b.val, v)
	// already set. unsuccessful
	if prev == v {
		return false
	}
	return true
}

func (b *AtomicFlag) Get() bool {
	return atomic.LoadInt32(&b.val) == 1
}
