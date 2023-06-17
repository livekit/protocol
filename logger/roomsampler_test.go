package logger

import (
	"fmt"
	"math"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/stretchr/testify/require"
)

type testStringer string

func (s testStringer) String() string {
	return string(s)
}

type testStringLike string

func TestRoomSampleRate(t *testing.T) {
	expectedRate := 0.50
	n := 10000

	var writeCount int
	s := NewRoomSampler(nil, expectedRate).(*roomSampler)
	for i := 0; i < n; i++ {
		write, ok := s.checkSampleField([]zapcore.Field{zap.String("room", fmt.Sprintf("room_%d", i))})
		if write && ok {
			writeCount++
		}
	}

	rate := float64(writeCount) / float64(n)
	require.Greater(t, 0.001, math.Abs(rate-expectedRate))
}

type testCore struct {
	zapcore.Core
	writeCount int
}

func (s *testCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	s.writeCount++
	return nil
}

func TestRoomSamplerWith(t *testing.T) {
	c := &testCore{Core: zap.NewExample().Core()}
	s := zap.New(NewRoomSampler(c, 0).(*roomSampler)).Sugar()

	s.Debugw("test", "room", "test")
	require.Equal(t, 0, c.writeCount)

	s.With("room", "test").Debugw("test")
	require.Equal(t, 0, c.writeCount)
}

func TestRoomSampleFindSampleField(t *testing.T) {
	v, ok := (*roomSampler).findSampleField(nil, []zapcore.Field{zap.String("room", "test")})
	require.Equal(t, "test", v)
	require.True(t, ok)

	v, ok = (*roomSampler).findSampleField(nil, []zapcore.Field{zap.Reflect("room", testStringLike("test"))})
	require.Equal(t, "test", v)
	require.True(t, ok)

	v, ok = (*roomSampler).findSampleField(nil, []zapcore.Field{zap.Stringer("room", testStringer("test"))})
	require.Equal(t, "test", v)
	require.True(t, ok)

	_, ok = (*roomSampler).findSampleField(nil, []zapcore.Field{zap.Int("room", 123)})
	require.False(t, ok)

	_, ok = (*roomSampler).findSampleField(nil, []zapcore.Field{})
	require.False(t, ok)
}
