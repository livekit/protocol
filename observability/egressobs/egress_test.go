package egressobs

import (
	"testing"

	"github.com/livekit/protocol/livekit"
	"github.com/stretchr/testify/require"
)

func TestGetSourceType(t *testing.T) {
	tests := []struct {
		sourceType livekit.EgressSourceType
		expected   string
	}{
		{livekit.EgressSourceType_EGRESS_SOURCE_TYPE_WEB, "web"},
		{livekit.EgressSourceType_EGRESS_SOURCE_TYPE_SDK, "sdk"},
		{livekit.EgressSourceType(99), ""}, // Unknown value falls back to undefined (empty string)
	}

	for _, tt := range tests {
		t.Run(tt.sourceType.String(), func(t *testing.T) {
			info := &livekit.EgressInfo{SourceType: tt.sourceType}
			result := getSourceType(info)
			require.Equal(t, tt.expected, string(result))
		})
	}
}

func TestGetRequestType(t *testing.T) {
	tests := []struct {
		name     string
		info     *livekit.EgressInfo
		expected string
	}{
		{
			name: "RoomComposite",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_RoomComposite{
					RoomComposite: &livekit.RoomCompositeEgressRequest{},
				},
			},
			expected: "room_composite",
		},
		{
			name: "Web",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_Web{
					Web: &livekit.WebEgressRequest{},
				},
			},
			expected: "web",
		},
		{
			name: "Participant",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_Participant{
					Participant: &livekit.ParticipantEgressRequest{},
				},
			},
			expected: "participant",
		},
		{
			name: "TrackComposite",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_TrackComposite{
					TrackComposite: &livekit.TrackCompositeEgressRequest{},
				},
			},
			expected: "track_composite",
		},
		{
			name: "Track",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_Track{
					Track: &livekit.TrackEgressRequest{},
				},
			},
			expected: "track",
		},
		{
			name:     "Undefined",
			info:     &livekit.EgressInfo{},
			expected: "", // Undefined is an empty string
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getRequestType(tt.info)
			require.Equal(t, tt.expected, string(result))
		})
	}
}

func TestGetAudioOnly(t *testing.T) {
	tests := []struct {
		name      string
		info      *livekit.EgressInfo
		audioOnly bool
	}{
		{
			name: "RoomComposite audio only",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_RoomComposite{
					RoomComposite: &livekit.RoomCompositeEgressRequest{AudioOnly: true},
				},
			},
			audioOnly: true,
		},
		{
			name: "RoomComposite not audio only",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_RoomComposite{
					RoomComposite: &livekit.RoomCompositeEgressRequest{AudioOnly: false},
				},
			},
			audioOnly: false,
		},
		{
			name: "Web audio only",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_Web{
					Web: &livekit.WebEgressRequest{AudioOnly: true},
				},
			},
			audioOnly: true,
		},
		{
			name: "Track request returns false",
			info: &livekit.EgressInfo{
				Request: &livekit.EgressInfo_Track{
					Track: &livekit.TrackEgressRequest{},
				},
			},
			audioOnly: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.audioOnly, getAudioOnly(tt.info))
		})
	}
}

func TestGetStatus(t *testing.T) {
	tests := []struct {
		status   livekit.EgressStatus
		expected string
	}{
		{livekit.EgressStatus_EGRESS_STARTING, "starting"},
		{livekit.EgressStatus_EGRESS_ACTIVE, "active"},
		{livekit.EgressStatus_EGRESS_ENDING, "ending"},
		{livekit.EgressStatus_EGRESS_COMPLETE, "complete"},
		{livekit.EgressStatus_EGRESS_ABORTED, "aborted"},
		{livekit.EgressStatus_EGRESS_LIMIT_REACHED, "limit_reached"},
		{livekit.EgressStatus_EGRESS_FAILED, "failed"},
	}

	for _, tt := range tests {
		t.Run(tt.status.String(), func(t *testing.T) {
			info := &livekit.EgressInfo{Status: tt.status}
			result := getStatus(info)
			require.Equal(t, tt.expected, string(result))
		})
	}
}
