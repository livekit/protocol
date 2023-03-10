package utils

import (
	"errors"
	"fmt"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/pion/webrtc/v3"
)

func GetMimeTypeForVideoCodec(codec livekit.VideoCodec) string {
	switch codec {
	case livekit.VideoCodec_H264_BASELINE,
		livekit.VideoCodec_H264_MAIN,
		livekit.VideoCodec_H264_HIGH:
		return webrtc.MimeTypeH264
	case livekit.VideoCodec_VP8:
		return webrtc.MimeTypeVP8
	default:
		logger.Errorw(fmt.Sprintf("GetMimeTypeForVideoCodec unimplemented for %s", codec.String()), errors.New("unimplemented"))
		return ""
	}
}

func GetMimeTypeForAudioCodec(codec livekit.AudioCodec) string {
	// AAC is not supported in webrtc

	switch codec {
	case livekit.AudioCodec_OPUS:
		return webrtc.MimeTypeOpus
	default:
		logger.Errorw(fmt.Sprintf("GetMimeTypeForAudioCodec unimplemented for %s", codec.String()), errors.New("unimplemented"))
		return ""
	}
}
