package egress

import "github.com/livekit/protocol/livekit"

const (
	EgressTypeRoomComposite  = "room_composite"
	EgressTypeWeb            = "web"
	EgressTypeTrackComposite = "track_composite"
	EgressTypeTrack          = "track"

	OutputTypeFile      = "file"
	OutputTypeStream    = "stream"
	OutputTypeWebsocket = "websocket"
	OutputTypeSegments  = "segments"
	OutputTypeMultiple  = "multiple"

	Unknown = "unknown"
)

func GetTypes(info *livekit.EgressInfo) (string, string) {
	switch req := info.Request.(type) {
	case *livekit.EgressInfo_RoomComposite:
		return EgressTypeRoomComposite, getOutputType(req.RoomComposite)

	case *livekit.EgressInfo_Web:
		return EgressTypeWeb, getOutputType(req.Web)

	case *livekit.EgressInfo_TrackComposite:
		return EgressTypeTrackComposite, getOutputType(req.TrackComposite)

	case *livekit.EgressInfo_Track:
		switch req.Track.Output.(type) {
		case *livekit.TrackEgressRequest_File:
			return EgressTypeTrack, OutputTypeFile
		case *livekit.TrackEgressRequest_WebsocketUrl:
			return EgressTypeTrack, OutputTypeWebsocket
		}
	}

	return Unknown, Unknown
}

func getOutputType(req interface {
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
