package ingress

import (
	"github.com/livekit/protocol/livekit"
)

// This validates that advanced encoding options have no consistency issues and provide enough parameters
// to be usable. Options that pass this test may still need some fields to be poulated with default values
// before being used in a media pipeline.
func ValidateVideoOptionsConsistency(options *livekit.IngressVideoOptions) error {
	if options == nil {
		return nil
	}

	switch options.Source {
	case livekit.TrackSource_UNKNOWN,
		livekit.TrackSource_CAMERA,
		livekit.TrackSource_SCREEN_SHARE:
		// continue
	default:
		return NewInvalidVideoParamsError("invalid track source")
	}

	switch o := options.EncodingOptions.(type) {
	case nil:
		// Default, continue
	case *livekit.IngressVideoOptions_Preset:
		_, ok := livekit.IngressVideoEncodingPreset_name[int32(o.Preset)]
		if !ok {
			return NewInvalidVideoParamsError("invalid preset")
		}

	case *livekit.IngressVideoOptions_Options:
		err := ValidateVideoEncodingOptionsConsistency(o.Options)
		if err != nil {
			return err
		}
	}

	return nil
}

func ValidateVideoEncodingOptionsConsistency(options *livekit.IngressVideoEncodingOptions) error {
	if options == nil {
		return NewInvalidVideoParamsError("empty options")
	}

	if (len(options.Layers)) == 0 {
		return NewInvalidVideoParamsError("no video layer")
	}

	if options.FrameRate < 0 {
		return NewInvalidVideoParamsError("invalid framerate")
	}

	switch options.VideoCodec {
	case livekit.VideoCodec_DEFAULT_VC,
		livekit.VideoCodec_H264_BASELINE,
		livekit.VideoCodec_VP8:
		// continue
	default:
		return NewInvalidVideoParamsError("video codec unsupported")
	}

	layersByQuality := make(map[livekit.VideoQuality]*livekit.VideoLayer)

	for _, layer := range options.Layers {
		if layer.Height == 0 || layer.Width == 0 {
			return ErrInvalidOutputDimensions
		}

		if _, ok := layersByQuality[layer.Quality]; ok {
			return NewInvalidVideoParamsError("more than one layer with the same quality level")
		}
		layersByQuality[layer.Quality] = layer
	}

	var oldLayerArea uint32
	for q := livekit.VideoQuality_LOW; q <= livekit.VideoQuality_HIGH; q++ {
		layer, ok := layersByQuality[q]
		if !ok {
			continue
		}
		layerArea := layer.Width * layer.Height

		if layerArea <= oldLayerArea {
			return NewInvalidVideoParamsError("video layers do not have increasing pixel count with increasing quality")
		}
		oldLayerArea = layerArea
	}

	return nil
}

func ValidateAudioOptionsConsistency(options *livekit.IngressAudioOptions) error {
	if options == nil {
		return nil
	}

	switch options.Source {
	case livekit.TrackSource_UNKNOWN,
		livekit.TrackSource_MICROPHONE,
		livekit.TrackSource_SCREEN_SHARE_AUDIO:
		// continue
	default:
		return NewInvalidAudioParamsError("invalid track source")
	}

	switch o := options.EncodingOptions.(type) {
	case nil:
		// Default, continue
	case *livekit.IngressAudioOptions_Preset:
		_, ok := livekit.IngressAudioEncodingPreset_name[int32(o.Preset)]
		if !ok {
			return NewInvalidAudioParamsError("invalid preset")
		}

	case *livekit.IngressAudioOptions_Options:
		err := ValidateAudioEncodingOptionsConsistency(o.Options)
		if err != nil {
			return err
		}
	}

	return nil
}

func ValidateAudioEncodingOptionsConsistency(options *livekit.IngressAudioEncodingOptions) error {
	if options == nil {
		return NewInvalidAudioParamsError("empty options")
	}

	switch options.AudioCodec {
	case livekit.AudioCodec_DEFAULT_AC,
		livekit.AudioCodec_OPUS:
		// continue
	default:
		return NewInvalidAudioParamsError("audio codec unsupported")
	}

	if options.Channels > 2 {
		return NewInvalidAudioParamsError("invalid audio channel count")
	}

	return nil
}
