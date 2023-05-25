package utils

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestMarshalUnmarshalGuid(t *testing.T) {
	id0 := livekit.TrackID(NewGuid(TrackPrefix))
	b0 := MarshalGuid(id0)
	id1 := UnmarshalGuid[livekit.TrackID](b0)
	b1 := MarshalGuid(id1)
	require.EqualValues(t, id0, id1)
	require.EqualValues(t, b0, b1)
}
