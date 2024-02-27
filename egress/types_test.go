package egress

import (
	"testing"

	"github.com/livekit/protocol/livekit"
	"github.com/stretchr/testify/require"
)

func TestGetOutputType(t *testing.T) {
	roomReq := &livekit.RoomCompositeEgressRequest{
		FileOutputs: []*livekit.EncodedFileOutput{
			&livekit.EncodedFileOutput{},
		},
	}

	ot := GetOutputType(roomReq)
	require.Equal(t, OutputTypeFile, ot)

	roomReq = &livekit.RoomCompositeEgressRequest{
		Output: &livekit.RoomCompositeEgressRequest_File{
			File: &livekit.EncodedFileOutput{},
		},
	}

	ot = GetOutputType(roomReq)
	require.Equal(t, OutputTypeFile, ot)

	trackReq := &livekit.TrackCompositeEgressRequest{
		SegmentOutputs: []*livekit.SegmentedFileOutput{
			&livekit.SegmentedFileOutput{},
		},
	}

	ot = GetOutputType(trackReq)
	require.Equal(t, OutputTypeSegments, ot)

	trackReq = &livekit.TrackCompositeEgressRequest{
		Output: &livekit.TrackCompositeEgressRequest_Segments{
			Segments: &livekit.SegmentedFileOutput{},
		},
	}

	ot = GetOutputType(trackReq)
	require.Equal(t, OutputTypeSegments, ot)

	webReq := &livekit.WebEgressRequest{
		StreamOutputs: []*livekit.StreamOutput{
			&livekit.StreamOutput{},
		},
	}

	ot = GetOutputType(webReq)
	require.Equal(t, OutputTypeStream, ot)

	webReq = &livekit.WebEgressRequest{
		Output: &livekit.WebEgressRequest_Stream{
			Stream: &livekit.StreamOutput{},
		},
	}

	ot = GetOutputType(webReq)
	require.Equal(t, OutputTypeStream, ot)

	participantReq := &livekit.ParticipantEgressRequest{
		ImageOutputs: []*livekit.ImageOutput{
			&livekit.ImageOutput{},
		},
	}

	ot = GetOutputType(participantReq)
	require.Equal(t, OutputTypeImages, ot)

	participantReq.SegmentOutputs = []*livekit.SegmentedFileOutput{
		&livekit.SegmentedFileOutput{},
	}

	ot = GetOutputType(participantReq)
	require.Equal(t, OutputTypeMultiple, ot)

}
