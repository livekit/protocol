package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestDeferredLogger(t *testing.T) {
	c := &testCore{Core: zap.NewExample().Core()}
	dc, resolve := newDeferredValueCore(c)
	s := zap.New(dc).Sugar()

	s.Infow("test")
	require.Equal(t, 0, c.WriteCount())

	s.With("a", "1").Infow("test")
	require.Equal(t, 0, c.WriteCount())

	resolve("b", "2")
	require.Equal(t, 2, c.WriteCount())

	s.With("c", "3").Infow("test")
	require.Equal(t, 3, c.WriteCount())
}
