package egress

import "github.com/livekit/protocol/livekit"

const (
	EgressTypeRoomComposite  = "room_composite"
	EgressTypeWeb            = "web"
	EgressTypeParticipant    = "participant"
	EgressTypeTrackComposite = "track_composite"
	EgressTypeTrack          = "track"

	OutputTypeFile     = "file"
	OutputTypeStream   = "stream"
	OutputTypeSegments = "segments"
	OutputTypeMultiple = "multiple"

	Unknown = "unknown"
)

func GetTypes(info *livekit.EgressInfo) (string, string) {
	switch req := info.Request.(type) {
	case *livekit.EgressInfo_RoomComposite:
		return EgressTypeRoomComposite, getOutputTypeDeprecated(req.RoomComposite)

	case *livekit.EgressInfo_Web:
		return EgressTypeWeb, getOutputTypeDeprecated(req.Web)

	case *livekit.EgressInfo_Participant:
		return EgressTypeParticipant, getOutputType(req.Participant)

	case *livekit.EgressInfo_TrackComposite:
		return EgressTypeTrackComposite, getOutputTypeDeprecated(req.TrackComposite)

	case *livekit.EgressInfo_Track:
		switch req.Track.Output.(type) {
		case *livekit.TrackEgressRequest_File:
			return EgressTypeTrack, OutputTypeFile
		case *livekit.TrackEgressRequest_WebsocketUrl:
			return EgressTypeTrack, OutputTypeStream
		}
	}

	return Unknown, Unknown
}

func getOutputTypeDeprecated(req interface {
	GetFile() *livekit.EncodedFileOutput
	GetStream() *livekit.StreamOutput
	GetSegments() *livekit.SegmentedFileOutput
	GetFileOutputs() []*livekit.EncodedFileOutput
	GetStreamOutputs() []*livekit.StreamOutput
	GetSegmentOutputs() []*livekit.SegmentedFileOutput
}) string {
	hasFile := len(req.GetFileOutputs()) > 0
	hasStream := len(req.GetStreamOutputs()) > 0
	hasSegments := len(req.GetSegmentOutputs()) > 0

	switch {
	case (hasFile && (hasStream || hasSegments)) || (hasStream && hasSegments):
		return OutputTypeMultiple
	case hasFile || req.GetFile() != nil:
		return OutputTypeFile
	case hasStream || req.GetStream() != nil:
		return OutputTypeStream
	case hasSegments || req.GetSegments() != nil:
		return OutputTypeSegments
	}

	return Unknown
}

func getOutputType(req interface {
	GetFileOutputs() []*livekit.EncodedFileOutput
	GetStreamOutputs() []*livekit.StreamOutput
	GetSegmentOutputs() []*livekit.SegmentedFileOutput
}) string {
	hasFile := len(req.GetFileOutputs()) > 0
	hasStream := len(req.GetStreamOutputs()) > 0
	hasSegments := len(req.GetSegmentOutputs()) > 0

	switch {
	case (hasFile && (hasStream || hasSegments)) || (hasStream && hasSegments):
		return OutputTypeMultiple
	case hasFile:
		return OutputTypeFile
	case hasStream:
		return OutputTypeStream
	case hasSegments:
		return OutputTypeSegments
	}

	return Unknown
}
