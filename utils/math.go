package utils

import (
	"time"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Signed | constraints.Unsigned | time.Duration
}

func Max[T Numeric](vs ...T) T {
	return Least(func(a, b T) bool { return a > b }, vs...)
}

func Min[T Numeric](vs ...T) T {
	return Least(func(a, b T) bool { return a < b }, vs...)
}

func Most[T Numeric](less func(a, b T) bool, vs ...T) T {
	return Least(func(a, b T) bool { return !less(a, b) }, vs...)
}

func Least[T Numeric](less func(a, b T) bool, vs ...T) T {
	if len(vs) == 0 {
		return 0
	}
	v := vs[0]
	for i := 1; i < len(vs); i++ {
		if less(vs[i], v) {
			v = vs[i]
		}
	}
	return v
}
