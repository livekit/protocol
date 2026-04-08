package egressobs

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/pkg/errors"

	"github.com/livekit/protocol/egress"
	"github.com/livekit/protocol/livekit"
)

type EgressStatus = string

const (
	EgressStatusUndefined    EgressStatus = ""
	EgressStatusStarting     EgressStatus = "starting"
	EgressStatusActive       EgressStatus = "active"
	EgressStatusEnding       EgressStatus = "ending"
	EgressStatusComplete     EgressStatus = "complete"
	EgressStatusFailed       EgressStatus = "failed"
	EgressStatusAborted      EgressStatus = "aborted"
	EgressStatusLimitReached EgressStatus = "limit_reached"
)

type EgressResults struct {
	FileResults    []*livekit.FileInfo
	StreamResults  []*livekit.StreamInfo
	SegmentResults []*livekit.SegmentsInfo
	ImageResults   []*livekit.ImagesInfo
}

func GetSourceType(info *livekit.EgressInfo) SessionSourceType {
	switch r := info.Request.(type) {
	// case *livekit.EgressInfo_Egress:
	// 	return getSourceTypeV2(r.Egress)
	case *livekit.EgressInfo_Replay:
		return getSourceTypeV2(r.Replay)
	default:
		switch info.SourceType {
		case livekit.EgressSourceType_EGRESS_SOURCE_TYPE_WEB:
			return SessionSourceTypeWeb
		case livekit.EgressSourceType_EGRESS_SOURCE_TYPE_SDK:
			return SessionSourceTypeSdk
		default:
			return SessionSourceTypeUndefined
		}
	}
}

func getSourceTypeV2(r egress.EgressRequest) SessionSourceType {
	if r.GetMedia() != nil {
		return SessionSourceTypeMedia
	}
	if r.GetTemplate() != nil {
		return SessionSourceTypeTemplate
	}
	if r.GetWeb() != nil {
		return SessionSourceTypeWeb
	}
	return SessionSourceTypeUndefined
}

func GetRequestType(info *livekit.EgressInfo) EgressRequestType {
	switch info.Request.(type) {
	// case *livekit.EgressInfo_Egress:
	// 	return EgressRequestTypeEgress
	case *livekit.EgressInfo_Replay:
		return EgressRequestTypeReplay
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

func GetStatus(info *livekit.EgressInfo) EgressStatus {
	switch info.Status {
	case livekit.EgressStatus_EGRESS_STARTING:
		return EgressStatusStarting
	case livekit.EgressStatus_EGRESS_ACTIVE:
		return EgressStatusActive
	case livekit.EgressStatus_EGRESS_ENDING:
		return EgressStatusEnding
	case livekit.EgressStatus_EGRESS_COMPLETE:
		return EgressStatusComplete
	case livekit.EgressStatus_EGRESS_ABORTED:
		return EgressStatusAborted
	case livekit.EgressStatus_EGRESS_LIMIT_REACHED:
		return EgressStatusLimitReached
	case livekit.EgressStatus_EGRESS_FAILED:
		return EgressStatusFailed
	default:
		return EgressStatusUndefined
	}
}

func GetRequest(info *livekit.EgressInfo) (string, error) {
	switch req := info.Request.(type) {
	// case *livekit.EgressInfo_Egress:
	// 	b, err := protojson.Marshal(req.Egress)
	// 	if err != nil {
	// 		return "", errors.Wrap(err, "failed to marshal egress request")
	// 	}
	// 	return string(b), nil
	case *livekit.EgressInfo_Replay:
		b, err := protojson.Marshal(req.Replay)
		if err != nil {
			return "", errors.Wrap(err, "failed serializing Replay request")
		}
		return string(b), nil
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
