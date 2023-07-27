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
			return EgressTypeTrack, OutputTypeStream
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
