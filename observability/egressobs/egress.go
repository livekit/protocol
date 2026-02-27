package egressobs

import (
	"encoding/json"

	"google.golang.org/protobuf/encoing/protojson"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
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
	logger := logger.GetLogger().WithValues("egressID", info.EgressId)

	switch req := info.Request.(type) {
	case *livekit.EgressInfo_RoomComposite:
		b, err := protojson.Marshal(req.RoomComposite)
		if err != nil {
			logger.Warnw("failed serializing RoomComposite request", err, "request", info.Request)
		}
		return string(b)
	case *livekit.EgressInfo_Web:
		b, err := protojson.Marshal(req.Web)
		if err != nil {
			logger.Warnw("failed serializing Web request", err, "request", info.Request)
		}
		return string(b)
	case *livekit.EgressInfo_Participant:
		b, err := protojson.Marshal(req.Participant)
		if err != nil {
			logger.Warnw("failed serializing Participant request", err, "request", info.Request)
		}
		return string(b)
	case *livekit.EgressInfo_TrackComposite:
		b, err := protojson.Marshal(req.TrackComposite)
		if err != nil {
			logger.Warnw("failed serializing TrackComposite request", err, "request", info.Request)
		}
		return string(b)
	case *livekit.EgressInfo_Track:
		b, err := protojson.Marshal(req.Track)
		if err != nil {
			logger.Warnw("failed serializing Track request", err, "request", info.Request)
		}
		return string(b)
	default:
		return ""
	}
}

func GetResult(info *livekit.EgressInfo) string {
	logger := logger.GetLogger().WithValues("egressID", info.EgressId)

	if file := info.GetFile(); file != nil {
		b, err := protojson.Marshal(file)
		if err != nil {
			logger.Warnw("failed serializing File result", err, "result", info.GetFile())
		}
		return string(b)
	} else if stream := info.GetStream(); stream != nil {
		b, err := protojson.Marshal(stream)
		if err != nil {
			logger.Warnw("failed serializing Stream result", err, "result", info.GetStream())
		}
		return string(b)
	} else if segments := info.GetSegments(); segments != nil {
		b, err := protojson.Marshal(segments)
		if err != nil {
			logger.Warnw("failed serializing File result", err, "result", info.GetSegments())
		}
		return string(b)
	} else {
		results := &EgressResults{
			FileResults:    info.FileResults,
			StreamResults:  info.StreamResults,
			SegmentResults: info.SegmentResults,
			ImageResults:   info.ImageResults,
		}
		b, err := json.Marshal(results)
		if err != nil {
			logger.Warnw("failed serializing Multiple result", err, "result", results)
		}
		return string(b)
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
