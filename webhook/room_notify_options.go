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

import "github.com/livekit/protocol/livekit"

// roomLevelEvents are the events a room-scoped webhook may receive. Egress, ingress
// and agent events are deliberately excluded: those resources carry their own webhook
// config on the request that created them (see egress.GetEgressNotifyOptions), so
// delivering them here would both duplicate and leak beyond the room's scope.
var roomLevelEvents = map[string]struct{}{
	EventRoomStarted:                  {},
	EventRoomFinished:                 {},
	EventParticipantJoined:            {},
	EventParticipantLeft:              {},
	EventParticipantConnectionAborted: {},
	EventTrackPublished:               {},
	EventTrackUnpublished:             {},
}

// IsRoomLevelEvent reports whether event may be delivered to a room-scoped webhook.
func IsRoomLevelEvent(event string) bool {
	_, ok := roomLevelEvents[event]
	return ok
}

// GetRoomNotifyOptions returns the notify options carrying a room's own webhooks
// (RoomInternal.GetWebhooks()), or nil when the room configured none or the event is
// not one a room-scoped webhook may receive.
func GetRoomNotifyOptions(event string, roomWebhooks []*livekit.WebhookConfig) []NotifyOption {
	if len(roomWebhooks) == 0 || !IsRoomLevelEvent(event) {
		return nil
	}

	return []NotifyOption{WithExtraWebhooks(roomWebhooks)}
}
