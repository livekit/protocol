package egressobs

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/pkg/errors"

	"github.com/livekit/protocol/livekit"
)

type EgressResults struct {
	FileResults    []*livekit.FileInfo
	StreamResults  []*livekit.StreamInfo
	SegmentResults []*livekit.SegmentsInfo
	ImageResults   []*livekit.ImagesInfo
}

func GetSourceType(info *livekit.EgressInfo) SessionSourceType {
	switch info.SourceType {
	case livekit.EgressSourceType_EGRESS_SOURCE_TYPE_WEB:
		return SessionSourceTypeWeb
	case livekit.EgressSourceType_EGRESS_SOURCE_TYPE_SDK:
		return SessionSourceTypeSdk
	default:
		return SessionSourceTypeUndefined
	}
}

func GetRequestType(info *livekit.EgressInfo) EgressRequestType {
	switch info.Request.(type) {
	case *livekit.EgressInfo_RoomComposite:
		return EgressRequestTypeRoomComposite
	case *livekit.EgressInfo_Web:
		return EgressRequestTypeWeb
	case *livekit.EgressInfo_Participant:
		return EgressRequestTypeParticipant
	case *livekit.EgressInfo_TrackComposite:
		return EgressRequestTypeTrackComposite
	case *livekit.EgressInfo_Track:
		return EgressRequestTypeTrack
	default:
		return EgressRequestTypeUndefined
	}
}

func GetStatus(info *livekit.EgressInfo) SessionStatus {
	switch info.Status {
	case livekit.EgressStatus_EGRESS_STARTING:
		return SessionStatusStarting
	case livekit.EgressStatus_EGRESS_ACTIVE:
		return SessionStatusActive
	case livekit.EgressStatus_EGRESS_ENDING:
		return SessionStatusEnding
	case livekit.EgressStatus_EGRESS_COMPLETE:
		return SessionStatusComplete
	case livekit.EgressStatus_EGRESS_ABORTED:
		return SessionStatusAborted
	case livekit.EgressStatus_EGRESS_LIMIT_REACHED:
		return SessionStatusLimitReached
	case livekit.EgressStatus_EGRESS_FAILED:
		return SessionStatusFailed
	default:
		return SessionStatusUndefined
	}
}

func GetRequest(info *livekit.EgressInfo) (string, error) {
	switch req := info.Request.(type) {
	case *livekit.EgressInfo_RoomComposite:
		b, err := protojson.Marshal(req.RoomComposite)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing RoomComposite request")
		}
		return string(b), nil
	case *livekit.EgressInfo_Web:
		b, err := protojson.Marshal(req.Web)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing Web request")
		}
		return string(b), nil
	case *livekit.EgressInfo_Participant:
		b, err := protojson.Marshal(req.Participant)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing Participant request")
		}
		return string(b), nil
	case *livekit.EgressInfo_TrackComposite:
		b, err := protojson.Marshal(req.TrackComposite)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing TrackComposite request")
		}
		return string(b), nil
	case *livekit.EgressInfo_Track:
		b, err := protojson.Marshal(req.Track)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing Track request")
		}
		return string(b), nil
	default:
		return "", nil
	}
}

func GetResult(info *livekit.EgressInfo) (string, error) {
	if file := info.GetFile(); file != nil {
		b, err := protojson.Marshal(file)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing File result")
		}
		return string(b), nil
	} else if stream := info.GetStream(); stream != nil {
		b, err := protojson.Marshal(stream)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing Stream result")
		}
		return string(b), nil
	} else if segments := info.GetSegments(); segments != nil {
		b, err := protojson.Marshal(segments)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing Segments result")
		}
		return string(b), nil
	} else {
		results := &EgressResults{
			FileResults:    info.FileResults,
			StreamResults:  info.StreamResults,
			SegmentResults: info.SegmentResults,
			ImageResults:   info.ImageResults,
		}
		b, err := json.Marshal(results)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing Multiple result")
		}
		return string(b), nil
	}
}

func GetAudioOnly(info *livekit.EgressInfo) bool {
	switch req := info.Request.(type) {
	case *livekit.EgressInfo_RoomComposite:
		return req.RoomComposite.AudioOnly
	case *livekit.EgressInfo_Web:
		return req.Web.AudioOnly
	default:
		return false
	}
}
