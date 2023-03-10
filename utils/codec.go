package utils

import (
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
		logger.Errorw("GetMimeTypeForVideoCodec unimplemented", nil, "codec", codec.String())
		return ""
	}
}

func GetMimeTypeForAudioCodec(codec livekit.AudioCodec) string {
	// AAC is not supported in webrtc

	switch codec {
	case livekit.AudioCodec_OPUS:
		return webrtc.MimeTypeOpus
	default:
		logger.Errorw("GetMimeTypeForAudioCodec unimplemented", nil, "codec", codec.String())
		return ""
	}
}
