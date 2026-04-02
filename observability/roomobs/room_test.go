package roomobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrackLayer(t *testing.T) {
	x, y := UnpackTrackLayer(PackTrackLayer(1920, 1080))
	require.Equal(t, 1920, x)
	require.Equal(t, 1080, y)
}

func TestCountryCode(t *testing.T) {
	require.Equal(t, "us", UnpackCountryCode(PackCountryCode("us")))
}

func TestTag(t *testing.T) {
	tag := NewTag("region", "us-west")

	key, value := tag.KeyValue()
	require.Equal(t, "region", key)
	require.Equal(t, "us-west", value)
	require.Equal(t, Tag("region\x1eus-west"), tag)
}
