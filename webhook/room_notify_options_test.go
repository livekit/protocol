// Copyright 2026 LiveKit, Inc.
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

package webhook

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestGetRoomNotifyOptions(t *testing.T) {
	whs := []*livekit.WebhookConfig{{Url: "https://example.com/hook"}}

	roomLevel := []string{
		EventRoomStarted,
		EventRoomFinished,
		EventParticipantJoined,
		EventParticipantLeft,
		EventParticipantConnectionAborted,
		EventTrackPublished,
		EventTrackUnpublished,
	}
	for _, ev := range roomLevel {
		opts := GetRoomNotifyOptions(ev, whs)
		require.Len(t, opts, 1, ev)

		p := &NotifyParams{}
		opts[0](p)
		require.Equal(t, whs, p.ExtraWebhooks, ev)
	}

	// egress, ingress and agent events never reach a room webhook
	notRoomLevel := []string{
		EventEgressStarted,
		EventEgressUpdated,
		EventEgressEnded,
		EventIngressStarted,
		EventIngressEnded,
		EventAgentJobStarted,
		EventAgentJobEnded,
	}
	for _, ev := range notRoomLevel {
		require.Nil(t, GetRoomNotifyOptions(ev, whs), ev)
	}

	// a room with no webhooks configured adds no options
	require.Nil(t, GetRoomNotifyOptions(EventRoomStarted, nil))
}
