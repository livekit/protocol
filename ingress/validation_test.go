package ingress

import (
	"testing"

	"github.com/livekit/protocol/livekit"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	info := &livekit.IngressInfo{}

	err := Validate(info)
	require.Error(t, err)

	info.StreamKey = "stream_key"
	err = Validate(info)
	require.Error(t, err)

	info.RoomName = "room_name"
	err = Validate(info)
	require.Error(t, err)

	info.ParticipantIdentity = "participant_identity"
	err = Validate(info)
	require.NoError(t, err)

	// make sure video parameters are validated. Full validation logic tested in the next test
	info.Video = &livekit.IngressVideoOptions{}
	err = Validate(info)
	require.NoError(t, err)

	info.Video.Source = livekit.TrackSource_MICROPHONE
	err = Validate(info)
	require.Error(t, err)

	info.Video.Source = livekit.TrackSource_CAMERA

	// make sure audio parameters are validated. Full validation logic tested in the next test
	info.Audio = &livekit.IngressAudioOptions{}
	err = Validate(info)
	require.NoError(t, err)

	info.Audio.Source = livekit.TrackSource_CAMERA
	err = Validate(info)
	require.Error(t, err)

	info.Audio.Source = livekit.TrackSource_SCREEN_SHARE_AUDIO
	err = Validate(info)
	require.NoError(t, err)
}

func TestValidateVideoOptionsConsistency(t *testing.T) {
	video := &livekit.IngressVideoOptions{}
	err := ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.Source = livekit.TrackSource_MICROPHONE
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.Source = livekit.TrackSource_CAMERA
	video.EncodingOptions = &livekit.IngressVideoOptions_Preset{
		Preset: livekit.IngressVideoEncodingPreset(42),
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions = &livekit.IngressVideoOptions_Preset{
		Preset: livekit.IngressVideoEncodingPreset_H264_1080P_30FPS_1_LAYER,
	}
	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.EncodingOptions = &livekit.IngressVideoOptions_Options{
		Options: &livekit.IngressVideoEncodingOptions{
			VideoCodec: livekit.VideoCodec_H264_HIGH,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions = &livekit.IngressVideoOptions_Options{
		Options: &livekit.IngressVideoEncodingOptions{
			VideoCodec: livekit.VideoCodec_DEFAULT_VC,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.EncodingOptions.(*livekit.IngressVideoOptions_Options).Options.Layers = []*livekit.VideoLayer{
		&livekit.VideoLayer{},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*livekit.IngressVideoOptions_Options).Options.Layers = []*livekit.VideoLayer{
		&livekit.VideoLayer{
			Width:  640,
			Height: 480,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.EncodingOptions.(*livekit.IngressVideoOptions_Options).Options.Layers = []*livekit.VideoLayer{
		&livekit.VideoLayer{
			Width:  641,
			Height: 480,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*livekit.IngressVideoOptions_Options).Options.Layers = []*livekit.VideoLayer{
		&livekit.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: livekit.VideoQuality_HIGH,
		},
		&livekit.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: livekit.VideoQuality_LOW,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*livekit.IngressVideoOptions_Options).Options.Layers = []*livekit.VideoLayer{
		&livekit.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: livekit.VideoQuality_HIGH,
		},
		&livekit.VideoLayer{
			Width:   1280,
			Height:  720,
			Quality: livekit.VideoQuality_HIGH,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*livekit.IngressVideoOptions_Options).Options.Layers = []*livekit.VideoLayer{
		&livekit.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: livekit.VideoQuality_HIGH,
		},
		&livekit.VideoLayer{
			Width:   1280,
			Height:  720,
			Quality: livekit.VideoQuality_LOW,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*livekit.IngressVideoOptions_Options).Options.Layers = []*livekit.VideoLayer{
		&livekit.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: livekit.VideoQuality_LOW,
		},
		&livekit.VideoLayer{
			Width:   1280,
			Height:  720,
			Quality: livekit.VideoQuality_HIGH,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)
}

func TestValidateAudioOptionsConsistency(t *testing.T) {
	audio := &livekit.IngressAudioOptions{}
	err := ValidateAudioOptionsConsistency(audio)
	require.NoError(t, err)

	audio.Source = livekit.TrackSource_CAMERA
	err = ValidateAudioOptionsConsistency(audio)
	require.Error(t, err)

	audio.Source = livekit.TrackSource_SCREEN_SHARE_AUDIO
	audio.EncodingOptions = &livekit.IngressAudioOptions_Preset{
		Preset: livekit.IngressAudioEncodingPreset(42),
	}
	err = ValidateAudioOptionsConsistency(audio)
	require.Error(t, err)

	audio.EncodingOptions = &livekit.IngressAudioOptions_Preset{
		Preset: livekit.IngressAudioEncodingPreset_OPUS_MONO_64KBS,
	}
	err = ValidateAudioOptionsConsistency(audio)
	require.NoError(t, err)

	audio.EncodingOptions = &livekit.IngressAudioOptions_Options{
		Options: &livekit.IngressAudioEncodingOptions{
			AudioCodec: livekit.AudioCodec_AAC,
		},
	}
	err = ValidateAudioOptionsConsistency(audio)
	require.Error(t, err)

	audio.EncodingOptions = &livekit.IngressAudioOptions_Options{
		Options: &livekit.IngressAudioEncodingOptions{
			AudioCodec: livekit.AudioCodec_OPUS,
			Channels:   3,
		},
	}
	err = ValidateAudioOptionsConsistency(audio)
	require.Error(t, err)

	audio.EncodingOptions.(*livekit.IngressAudioOptions_Options).Options.Channels = 2
	err = ValidateAudioOptionsConsistency(audio)
	require.NoError(t, err)
}
