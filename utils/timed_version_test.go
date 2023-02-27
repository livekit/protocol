package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimedVersion(t *testing.T) {
	t.Run("timed versions are monotonic and comparable", func(t *testing.T) {
		gen := NewDefaultTimedVersionGenerator()
		tv1 := gen.New()
		tv2 := gen.New()
		tv3 := gen.New()

		require.True(t, tv3.After(tv1))
		require.True(t, tv3.After(tv2))
		require.True(t, tv2.After(tv1))

		tv2.Update(tv3)
		require.True(t, tv2.After(tv1))
		// tv3 and tv2 are equivalent after update
		require.False(t, tv3.After(tv2))

		require.Equal(t, 0, tv1.Compare(tv1))
		require.Equal(t, -1, tv1.Compare(tv2))
		require.Equal(t, 1, tv2.Compare(tv1))
	})

	t.Run("protobuf roundtrip", func(t *testing.T) {
		gen := NewDefaultTimedVersionGenerator()
		tv1 := gen.New()
		tv2 := NewTimedVersionFromProto(tv1.ToProto())
		require.Equal(t, tv1.v.Load(), tv2.v.Load())
	})

	t.Run("from zero time yields epoch version", func(t *testing.T) {
		gen := NewDefaultTimedVersionGenerator()
		tv1 := NewTimedVersionFromTime(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC))
		tv2 := gen.New()
		require.True(t, tv2.After(tv1))
	})
}
