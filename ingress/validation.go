package ingress

import (
	"github.com/livekit/protocol/livekit"
)

func ValidateVideoOptionsConsistency(options *livekit.IngressVideoOptions) error {
	layersByQuality := make(map[livekit.VideoQuality]*livekit.VideoLayer)

	for _, layer := range options.Layers {
		if layer.Height == 0 || layer.Width == 0 {
			return ErrInvalidOutputDimensions
		}

		if layer.Bitrate == 0 {
			return NewInvalidVideoParamsError("invalid bitrate")
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
