package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testItem struct {
	Hook[testItem]
	value string
}

type testList = List[testItem, *testItem]

func collectValues(l *testList) []string {
	var out []string
	for it := l.Front(); it != nil; it = it.Next() {
		out = append(out, it.value)
	}
	return out
}

func TestTLZeroValue(t *testing.T) {
	var l testList
	require.True(t, l.Empty())
	require.Nil(t, l.Front())
	require.Nil(t, l.Back())
}

func TestTLPushFrontBackAndRemove(t *testing.T) {
	var l testList
	a := &testItem{value: "a"}
	b := &testItem{value: "b"}
	c := &testItem{value: "c"}

	l.PushFront(b)
	l.PushFront(a)
	l.PushBack(c)

	require.Equal(t, []string{"a", "b", "c"}, collectValues(&l))
	require.Same(t, a, l.Front())
	require.Same(t, c, l.Back())
	require.Same(t, b, a.Next())
	require.Same(t, b, c.Prev())

	l.Remove(b)
	require.Equal(t, []string{"a", "c"}, collectValues(&l))
	require.Nil(t, b.Next())
	require.Nil(t, b.Prev())

	l.Remove(a)
	l.Remove(c)
	require.True(t, l.Empty())
	require.Nil(t, l.Front())
	require.Nil(t, l.Back())
}

func TestTLInsertAndMove(t *testing.T) {
	var l testList
	a := &testItem{value: "a"}
	b := &testItem{value: "b"}
	c := &testItem{value: "c"}
	d := &testItem{value: "d"}

	l.PushBack(a)
	l.PushBack(c)
	l.InsertAfter(b, a)
	l.InsertBefore(d, a)

	require.Equal(t, []string{"d", "a", "b", "c"}, collectValues(&l))

	l.MoveToFront(c)
	require.Equal(t, []string{"c", "d", "a", "b"}, collectValues(&l))

	l.MoveToBack(d)
	require.Equal(t, []string{"c", "a", "b", "d"}, collectValues(&l))
	require.Same(t, c, l.Front())
	require.Same(t, d, l.Back())
}
