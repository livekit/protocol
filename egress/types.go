// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package egress

import "github.com/livekit/protocol/livekit"

const (
	EgressTypeRoomComposite  = "room_composite"
	EgressTypeWeb            = "web"
	EgressTypeTrackComposite = "track_composite"
	EgressTypeTrack          = "track"

	OutputTypeFile     = "file"
	OutputTypeStream   = "stream"
	OutputTypeSegments = "segments"
	OutputTypeMultiple = "multiple"

	Unknown = "unknown"
)

type EncodedOutput interface {
	GetFileOutputs() []*livekit.EncodedFileOutput
	GetStreamOutputs() []*livekit.StreamOutput
	GetSegmentOutputs() []*livekit.SegmentedFileOutput
}

type EncodedOutputDeprecated interface {
	GetFile() *livekit.EncodedFileOutput
	GetStream() *livekit.StreamOutput
	GetSegments() *livekit.SegmentedFileOutput
}

func GetTypes(request interface{}) (string, string) {
	switch req := request.(type) {
	case *livekit.EgressInfo_RoomComposite:
		return EgressTypeRoomComposite, GetOutputType(req.RoomComposite)

	case *livekit.EgressInfo_Web:
		return EgressTypeWeb, GetOutputType(req.Web)

	case *livekit.EgressInfo_TrackComposite:
		return EgressTypeTrackComposite, GetOutputType(req.TrackComposite)

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

func GetOutputType(req EncodedOutput) string {
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
	default:
		if r, ok := req.(EncodedOutputDeprecated); ok {
			if r.GetFile() != nil {
				return OutputTypeFile
			}
			if r.GetStream() != nil {
				return OutputTypeStream
			}
			if r.GetSegments() != nil {
				return OutputTypeSegments
			}
		}
	}

	return Unknown
}
