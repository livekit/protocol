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

package utils

import (
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils/guid"
)

const (
	GuidSize = guid.Size
)

const (
	RoomPrefix            = guid.RoomPrefix
	NodePrefix            = guid.NodePrefix
	ParticipantPrefix     = guid.ParticipantPrefix
	TrackPrefix           = guid.TrackPrefix
	APIKeyPrefix          = guid.APIKeyPrefix
	EgressPrefix          = guid.EgressPrefix
	IngressPrefix         = guid.IngressPrefix
	SIPTrunkPrefix        = guid.SIPTrunkPrefix
	SIPDispatchRulePrefix = guid.SIPDispatchRulePrefix
	SIPCallPrefix         = guid.SIPCallPrefix
	RPCPrefix             = guid.RPCPrefix
	WHIPResourcePrefix    = guid.WHIPResourcePrefix
	RTMPResourcePrefix    = guid.RTMPResourcePrefix
	URLResourcePrefix     = guid.URLResourcePrefix
	AgentWorkerPrefix     = guid.AgentWorkerPrefix
	AgentJobPrefix        = guid.AgentJobPrefix
	AgentDispatchPrefix   = guid.AgentDispatchPrefix
)

func NewGuid(prefix string) string {
	return guid.New(prefix)
}

// HashedID creates a hashed ID from a unique string
func HashedID(id string) string {
	return guid.HashedID(id)
}

func LocalNodeID() (string, error) {
	return guid.LocalNodeID()
}

func MarshalGuid[T livekit.Guid](id T) livekit.GuidBlock {
	return guid.Marshal(id)
}

func UnmarshalGuid[T livekit.Guid](b livekit.GuidBlock) T {
	return guid.Unmarshal[T](b)
}
