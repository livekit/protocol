package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimedVersion(t *testing.T) {
	t.Run("timed versions are monotonic and comparable", func(t *testing.T) {
		gen := NewDefaultTimedVersionGenerator()
		tv1 := gen.New()
		tv2 := gen.New()
		tv3 := gen.New()

		assert.True(t, tv3.After(tv1))
		assert.True(t, tv3.After(tv2))
		assert.True(t, tv2.After(tv1))

		tv2.Update(tv3)
		assert.True(t, tv2.After(tv1))
		// tv3 and tv2 are equivalent after update
		assert.False(t, tv3.After(tv2))

		assert.Equal(t, 0, tv1.Compare(tv1))
		assert.Equal(t, -1, tv1.Compare(tv2))
		assert.Equal(t, 1, tv2.Compare(tv1))
	})

	t.Run("protobuf roundtrip", func(t *testing.T) {
		gen := NewDefaultTimedVersionGenerator()
		tv1 := gen.New()
		tv2 := NewTimedVersionFromProto(tv1.ToProto())
		require.Equal(t, *tv1, *tv2)
	})

	t.Run("timed version protobufs are backward compatible", func(t *testing.T) {
		gen := NewDefaultTimedVersionGenerator()
		tv := gen.New().ToProto()
		now := time.Now()
		d := tv.UnixMicro - now.UnixMicro()

		// +/- 1 millisecond
		require.Less(t, -int64(1000), d)
		require.Greater(t, int64(1000), d)
	})
}
