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

package sip

import (
	"time"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
)

func BuildSIPToken(
	apiKey, secret, roomName string,
	participantIdentity, participantName, participantMeta string,
	participantAttrs map[string]string,
) (string, error) {
	t := true
	at := auth.NewAccessToken(apiKey, secret).
		AddGrant(&auth.VideoGrant{
			RoomJoin:     true,
			Room:         roomName,
			CanSubscribe: &t,
			CanPublish:   &t,
		}).
		SetIdentity(participantIdentity).
		SetName(participantName).
		SetMetadata(participantMeta).
		SetAttributes(participantAttrs).
		SetKind(livekit.ParticipantInfo_SIP).
		SetValidFor(24 * time.Hour)

	return at.ToJWT()
}
