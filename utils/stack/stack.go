package stack

type Stack[T any] []T

func (s Stack[T]) Empty() bool {
	return len(s) == 0
}

func (s *Stack[T]) Reset() {
	*s = (*s)[:0]
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s Stack[T]) Peek() T {
	return s[len(s)-1]
}
