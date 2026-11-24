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

// AutoRoomRequest builds the request for a room's auto room egress, or nil if it has none
func AutoRoomRequest(cfg *livekit.RoomEgress, roomName string) *rpc.StartEgressRequest {
	if v2 := cfg.GetRoomEgress(); v2 != nil {
		req := proto.Clone(v2).(*livekit.StartEgressRequest)
		req.RoomName = roomName
		return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_Egress{Egress: req}}
	}
	if v1 := cfg.GetRoom(); v1 != nil {
		req := proto.Clone(v1).(*livekit.RoomCompositeEgressRequest)
		req.RoomName = roomName
		// TemplateSource has no audio mixing
		if req.AudioMixing != livekit.AudioMixing_DEFAULT_MIXING {
			return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_RoomComposite{RoomComposite: req}}
		}
		return &rpc.StartEgressRequest{Request: &rpc.StartEgressRequest_Egress{Egress: roomRequestFromV1(req)}}
	}
	return nil
}

func roomRequestFromV1(v1 *livekit.RoomCompositeEgressRequest) *livekit.StartEgressRequest {
	req := &livekit.StartEgressRequest{
		RoomName: v1.RoomName,
		Source: &livekit.StartEgressRequest_Template{Template: &livekit.TemplateSource{
			Layout:        v1.Layout,
			AudioOnly:     v1.AudioOnly,
			VideoOnly:     v1.VideoOnly,
			CustomBaseUrl: v1.CustomBaseUrl,
		}},
		Outputs:  outputsFromV1(v1.FileOutputs, v1.SegmentOutputs),
		Webhooks: v1.Webhooks,
	}

	switch o := v1.Options.(type) {
	case *livekit.RoomCompositeEgressRequest_Preset:
		req.Encoding = &livekit.StartEgressRequest_Preset{Preset: o.Preset}
	case *livekit.RoomCompositeEgressRequest_Advanced:
		req.Encoding = &livekit.StartEgressRequest_Advanced{Advanced: o.Advanced}
	}

	for _, s := range v1.StreamOutputs {
		req.Outputs = append(req.Outputs, &livekit.Output{Config: &livekit.Output_Stream{Stream: s}})
	}
	for _, i := range v1.ImageOutputs {
		storage := storageFrom(i)
		i.Output = nil
		req.Outputs = append(req.Outputs, &livekit.Output{Config: &livekit.Output_Images{Images: i}, Storage: storage})
	}

	switch o := v1.Output.(type) {
	case *livekit.RoomCompositeEgressRequest_File:
		req.Outputs = append(req.Outputs, fileOutput(o.File))
	case *livekit.RoomCompositeEgressRequest_Stream:
		req.Outputs = append(req.Outputs, &livekit.Output{Config: &livekit.Output_Stream{Stream: o.Stream}})
	case *livekit.RoomCompositeEgressRequest_Segments:
		req.Outputs = append(req.Outputs, segmentsOutput(o.Segments))
	}

	return req
}

// AutoParticipantRequest builds the auto-egress request for a participant, or nil if the room has none
func AutoParticipantRequest(internal *livekit.RoomInternal, roomName, identity string) *livekit.StartEgressRequest {
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
		return req
	}

	if v1 := internal.GetParticipantEgress(); v1 != nil {
		c := proto.Clone(v1).(*livekit.AutoParticipantEgress)
		req := &livekit.StartEgressRequest{
			RoomName: roomName,
			Source:   participantSource(identity, false, false),
			Outputs:  outputsFromV1(c.FileOutputs, c.SegmentOutputs),
		}
		switch o := c.Options.(type) {
		case *livekit.AutoParticipantEgress_Preset:
			req.Encoding = &livekit.StartEgressRequest_Preset{Preset: o.Preset}
		case *livekit.AutoParticipantEgress_Advanced:
			req.Encoding = &livekit.StartEgressRequest_Advanced{Advanced: o.Advanced}
		}
		return req
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

// AutoTrackRequest builds the auto-egress request for a published track, or nil if the room has none
func AutoTrackRequest(internal *livekit.RoomInternal, roomName, trackID string, trackType livekit.TrackType) *livekit.StartEgressRequest {
	media := &livekit.MediaSource{}
	if trackType == livekit.TrackType_AUDIO {
		media.Audio = &livekit.AudioConfig{Routes: []*livekit.AudioRoute{
			{Match: &livekit.AudioRoute_TrackId{TrackId: trackID}},
		}}
	} else {
		media.Video = &livekit.MediaSource_VideoTrackId{VideoTrackId: trackID}
	}
	source := &livekit.StartEgressRequest_Media{Media: media}

	if v2 := internal.GetTrackEgressConfig(); v2 != nil {
		c := proto.Clone(v2).(*livekit.TrackEgressConfig)
		req := &livekit.StartEgressRequest{
			RoomName: roomName,
			Source:   source,
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
		return req
	}

	if v1 := internal.GetTrackEgress(); v1 != nil {
		c := proto.Clone(v1).(*livekit.AutoTrackEgress)
		return &livekit.StartEgressRequest{
			RoomName: roomName,
			Source:   source,
			Encoding: &livekit.StartEgressRequest_Preset{Preset: livekit.EncodingOptionsPreset_PASSTHROUGH},
			Outputs: []*livekit.Output{{Config: &livekit.Output_File{File: &livekit.FileOutput{
				Filepath:        trackFilepath(c.Filepath),
				DisableManifest: c.DisableManifest,
			}}}},
			Storage: storageFrom(c),
		}
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

// ValidateRoomEgress rejects v1 and v2 config for the same kind, and v2 config missing a source or outputs
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
	if pe := cfg.GetParticipantEgress(); pe != nil && len(pe.Outputs) == 0 {
		return psrpc.NewErrorf(psrpc.InvalidArgument, "participant_egress requires an output")
	}
	if te := cfg.GetTrackEgress(); te != nil && len(te.Outputs) == 0 {
		return psrpc.NewErrorf(psrpc.InvalidArgument, "track_egress requires an output")
	}
	return nil
}

func outputsFromV1(files []*livekit.EncodedFileOutput, segments []*livekit.SegmentedFileOutput) []*livekit.Output {
	outputs := make([]*livekit.Output, 0, len(files)+len(segments))
	for _, f := range files {
		outputs = append(outputs, fileOutput(f))
	}
	for _, s := range segments {
		outputs = append(outputs, segmentsOutput(s))
	}
	return outputs
}

func fileOutput(f *livekit.EncodedFileOutput) *livekit.Output {
	return &livekit.Output{
		Config: &livekit.Output_File{File: &livekit.FileOutput{
			FileType:        f.FileType,
			Filepath:        f.Filepath,
			DisableManifest: f.DisableManifest,
		}},
		Storage: storageFrom(f),
	}
}

// segmentsOutput clears s's inline upload, so s must belong to a cloned config
func segmentsOutput(s *livekit.SegmentedFileOutput) *livekit.Output {
	storage := storageFrom(s)
	s.Output = nil
	return &livekit.Output{Config: &livekit.Output_Segments{Segments: s}, Storage: storage}
}

func storageFrom(u UploadRequest) *livekit.StorageConfig {
	if s3 := u.GetS3(); s3 != nil {
		return &livekit.StorageConfig{Provider: &livekit.StorageConfig_S3{S3: s3}}
	}
	if gcp := u.GetGcp(); gcp != nil {
		return &livekit.StorageConfig{Provider: &livekit.StorageConfig_Gcp{Gcp: gcp}}
	}
	if azure := u.GetAzure(); azure != nil {
		return &livekit.StorageConfig{Provider: &livekit.StorageConfig_Azure{Azure: azure}}
	}
	if aliOSS := u.GetAliOSS(); aliOSS != nil {
		return &livekit.StorageConfig{Provider: &livekit.StorageConfig_AliOSS{AliOSS: aliOSS}}
	}
	return nil
}
