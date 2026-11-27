// Copyright 2025 LiveKit, Inc.
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

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/rpc"
	"github.com/livekit/psrpc"
)

// AutoRoomRequest builds a room's auto room egress request from its v1 or v2 config, or nil if it has none
func AutoRoomRequest(cfg *livekit.RoomEgress, roomName string) *rpc.StartEgressRequest {
	if v2 := cfg.GetRoomEgress(); v2 != nil {
		req := proto.Clone(v2).(*livekit.StartEgressRequest)
		req.RoomName = roomName
		return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_Egress{Egress: req}}
	}
	if v1 := cfg.GetRoom(); v1 != nil {
		req := proto.Clone(v1).(*livekit.RoomCompositeEgressRequest)
		req.RoomName = roomName
		return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_RoomComposite{RoomComposite: req}}
	}
	return nil
}

// AutoParticipantRequest builds a participant's auto-egress request from the room's v1 or v2 config, or nil if it has none
func AutoParticipantRequest(internal *livekit.RoomInternal, roomName, identity string) *rpc.StartEgressRequest {
	if v2 := internal.GetParticipantEgressConfig(); v2 != nil {
		c := proto.Clone(v2).(*livekit.ParticipantEgressConfig)
		req := &livekit.StartEgressRequest{
			RoomName: roomName,
			Source:   participantSource(identity, c.PreferScreenShare, c.CaptureData),
			Outputs:  c.Outputs,
			Storage:  c.Storage,
			Webhooks: c.Webhooks,
		}
		switch e := c.Encoding.(type) {
		case *livekit.ParticipantEgressConfig_Preset:
			req.Encoding = &livekit.StartEgressRequest_Preset{Preset: e.Preset}
		case *livekit.ParticipantEgressConfig_Advanced:
			req.Encoding = &livekit.StartEgressRequest_Advanced{Advanced: e.Advanced}
		}
		return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_Egress{Egress: req}}
	}

	if v1 := internal.GetParticipantEgress(); v1 != nil {
		c := proto.Clone(v1).(*livekit.AutoParticipantEgress)
		req := &livekit.ParticipantEgressRequest{
			RoomName:       roomName,
			Identity:       identity,
			FileOutputs:    c.FileOutputs,
			SegmentOutputs: c.SegmentOutputs,
		}
		switch o := c.Options.(type) {
		case *livekit.AutoParticipantEgress_Preset:
			req.Options = &livekit.ParticipantEgressRequest_Preset{Preset: o.Preset}
		case *livekit.AutoParticipantEgress_Advanced:
			req.Options = &livekit.ParticipantEgressRequest_Advanced{Advanced: o.Advanced}
		}
		return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_Participant{Participant: req}}
	}

	return nil
}

func participantSource(identity string, preferScreenShare, captureData bool) *livekit.StartEgressRequest_Media {
	media := &livekit.MediaSource{
		Video: &livekit.MediaSource_ParticipantVideo{ParticipantVideo: &livekit.ParticipantVideo{
			Identity:          identity,
			PreferScreenShare: preferScreenShare,
		}},
		Audio: &livekit.AudioConfig{Routes: []*livekit.AudioRoute{
			{Match: &livekit.AudioRoute_ParticipantIdentity{ParticipantIdentity: identity}},
		}},
	}
	if captureData {
		media.Data = &livekit.DataConfig{Selectors: []*livekit.DataSelector{
			{Match: &livekit.DataSelector_ParticipantIdentity{ParticipantIdentity: identity}},
		}}
	}
	return &livekit.StartEgressRequest_Media{Media: media}
}

// AutoTrackRequest builds a published track's auto-egress request from the room's v1 or v2 config, or nil if it has none
func AutoTrackRequest(internal *livekit.RoomInternal, roomName, trackID string, trackType livekit.TrackType) *rpc.StartEgressRequest {
	if v2 := internal.GetTrackEgressConfig(); v2 != nil {
		c := proto.Clone(v2).(*livekit.TrackEgressConfig)
		media := &livekit.MediaSource{}
		if trackType == livekit.TrackType_AUDIO {
			media.Audio = &livekit.AudioConfig{Routes: []*livekit.AudioRoute{
				{Match: &livekit.AudioRoute_TrackId{TrackId: trackID}},
			}}
		} else {
			media.Video = &livekit.MediaSource_VideoTrackId{VideoTrackId: trackID}
		}
		req := &livekit.StartEgressRequest{
			RoomName: roomName,
			Source:   &livekit.StartEgressRequest_Media{Media: media},
			Outputs:  c.Outputs,
			Storage:  c.Storage,
			Webhooks: c.Webhooks,
		}
		switch e := c.Encoding.(type) {
		case *livekit.TrackEgressConfig_Preset:
			req.Encoding = &livekit.StartEgressRequest_Preset{Preset: e.Preset}
		case *livekit.TrackEgressConfig_Advanced:
			req.Encoding = &livekit.StartEgressRequest_Advanced{Advanced: e.Advanced}
		}
		return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_Egress{Egress: req}}
	}

	if v1 := internal.GetTrackEgress(); v1 != nil {
		c := proto.Clone(v1).(*livekit.AutoTrackEgress)
		output := &livekit.DirectFileOutput{
			Filepath:        trackFilepath(c.Filepath),
			DisableManifest: c.DisableManifest,
		}
		switch o := c.Output.(type) {
		case *livekit.AutoTrackEgress_S3:
			output.Output = &livekit.DirectFileOutput_S3{S3: o.S3}
		case *livekit.AutoTrackEgress_Gcp:
			output.Output = &livekit.DirectFileOutput_Gcp{Gcp: o.Gcp}
		case *livekit.AutoTrackEgress_Azure:
			output.Output = &livekit.DirectFileOutput_Azure{Azure: o.Azure}
		case *livekit.AutoTrackEgress_AliOSS:
			output.Output = &livekit.DirectFileOutput_AliOSS{AliOSS: o.AliOSS}
		}
		req := &livekit.TrackEgressRequest{
			RoomName: roomName,
			TrackId:  trackID,
			Output:   &livekit.TrackEgressRequest_File{File: output},
		}
		return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_Track{Track: req}}
	}

	return nil
}

func trackFilepath(filepath string) string {
	if filepath == "" || strings.HasSuffix(filepath, "/") || strings.Contains(filepath, "{track_id}") {
		return filepath
	}

	idx := strings.Index(filepath, ".")
	if idx == -1 {
		return fmt.Sprintf("%s-{track_id}", filepath)
	} else {
		return fmt.Sprintf("%s-%s%s", filepath[:idx], "{track_id}", filepath[idx:])
	}
}

// ValidateRoomEgress rejects v1 and v2 config for the same kind, and v2 config that egress would reject
func ValidateRoomEgress(cfg *livekit.RoomEgress) error {
	switch {
	case cfg.GetRoom() != nil && cfg.GetRoomEgress() != nil:
		return psrpc.NewErrorf(psrpc.InvalidArgument, "room and room_egress are mutually exclusive")
	case cfg.GetParticipant() != nil && cfg.GetParticipantEgress() != nil:
		return psrpc.NewErrorf(psrpc.InvalidArgument, "participant and participant_egress are mutually exclusive")
	case cfg.GetTracks() != nil && cfg.GetTrackEgress() != nil:
		return psrpc.NewErrorf(psrpc.InvalidArgument, "tracks and track_egress are mutually exclusive")
	}

	if re := cfg.GetRoomEgress(); re != nil {
		if re.Source == nil {
			return psrpc.NewErrorf(psrpc.InvalidArgument, "room_egress requires a source")
		}
		if len(re.Outputs) == 0 {
			return psrpc.NewErrorf(psrpc.InvalidArgument, "room_egress requires an output")
		}
	}
	if pe := cfg.GetParticipantEgress(); pe != nil {
		if len(pe.Outputs) == 0 {
			return psrpc.NewErrorf(psrpc.InvalidArgument, "participant_egress requires an output")
		}
		if pe.CaptureData && (len(pe.Outputs) != 1 || pe.Outputs[0].GetFile().GetFileType() != livekit.EncodedFileType_MCAP) {
			return psrpc.NewErrorf(psrpc.InvalidArgument, "participant_egress capture_data requires a single MCAP file output")
		}
	}
	if te := cfg.GetTrackEgress(); te != nil && len(te.Outputs) == 0 {
		return psrpc.NewErrorf(psrpc.InvalidArgument, "track_egress requires an output")
	}
	return nil
}
