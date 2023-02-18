package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, tv1.v.Load(), tv2.v.Load())
	})
}
