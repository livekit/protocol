package list

type Hook[T any] struct {
	next *T
	prev *T
}

func (h *Hook[T]) getListHook() *Hook[T] { //nolint:unused // implements hookAccessor interface
	return h
}

func (h *Hook[T]) Next() *T {
	if h == nil {
		return nil
	}
	return h.next
}

func (h *Hook[T]) Prev() *T {
	if h == nil {
		return nil
	}
	return h.prev
}

type hookAccessor[T any] interface {
	getListHook() *Hook[T]
}

type Hooked[T any] interface {
	*T
	hookAccessor[T]
}

// List is an intrusive doubly-linked list. Mutations involving a node the list
// does not own are no-ops, as in container/list. A node linked into a DIFFERENT
// list is indistinguishable from a member without a per-node list pointer and
// remains the caller's contract.
type List[T any, P Hooked[T]] struct {
	head P
	tail P
}

// linked reports whether it is a member of this list. A zeroed hook is
// ambiguous between "not linked" and "sole element"; the head check settles it.
func (l *List[T, P]) linked(it P, h *Hook[T]) bool {
	return h.prev != nil || h.next != nil || l.head == it
}

func (l *List[T, P]) Empty() bool {
	return l.head == nil
}

func (l *List[T, P]) Front() P {
	return l.head
}

func (l *List[T, P]) Back() P {
	return l.tail
}

func (l *List[T, P]) PushFront(it P) {
	l.insert(it, nil, l.head)
}

func (l *List[T, P]) PushBack(it P) {
	l.insert(it, l.tail, nil)
}

func (l *List[T, P]) InsertBefore(it, mark P) {
	if !l.linked(mark, mark.getListHook()) {
		return
	}
	l.insert(it, mark.getListHook().prev, mark)
}

func (l *List[T, P]) InsertAfter(it, mark P) {
	if !l.linked(mark, mark.getListHook()) {
		return
	}
	l.insert(it, mark, mark.getListHook().next)
}

func (l *List[T, P]) MoveToFront(it P) {
	if l.head == it {
		return
	}
	h := it.getListHook()
	if !l.linked(it, h) {
		return
	}
	l.unlink(h)
	l.link(it, h, nil, l.head)
}

func (l *List[T, P]) MoveToBack(it P) {
	if l.tail == it {
		return
	}
	h := it.getListHook()
	if !l.linked(it, h) {
		return
	}
	l.unlink(h)
	l.link(it, h, l.tail, nil)
}

func (l *List[T, P]) Remove(it P) {
	h := it.getListHook()
	if !l.linked(it, h) {
		return
	}
	l.unlink(h)
	h.next = nil
	h.prev = nil
}

func (l *List[T, P]) insert(it, prev, next P) {
	h := it.getListHook()
	if l.linked(it, h) {
		return
	}
	l.link(it, h, prev, next)
}

func (l *List[T, P]) link(it P, h *Hook[T], prev, next P) {
	h.prev = prev
	h.next = next

	if prev != nil {
		prev.getListHook().next = it
	} else {
		l.head = it
	}
	if next != nil {
		next.getListHook().prev = it
	} else {
		l.tail = it
	}
}

// unlink requires a linked h: nil prev/next mean head/tail here, so a zeroed
// hook would clear both.
func (l *List[T, P]) unlink(h *Hook[T]) {
	if h.prev != nil {
		P(h.prev).getListHook().next = h.next
	} else {
		l.head = h.next
	}
	if h.next != nil {
		P(h.next).getListHook().prev = h.prev
	} else {
		l.tail = h.prev
	}
}
