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

package signalling

import (
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
)

type Signalv2ServerMessageCache struct {
	*SignalCache
}

func NewSignalv2ServerMessageCache(params SignalCacheParams) *Signalv2ServerMessageCache {
	return &Signalv2ServerMessageCache{
		SignalCache: NewSignalCache(params),
	}
}

func (s *Signalv2ServerMessageCache) Add(msg *livekit.Signalv2ServerMessage) *livekit.Signalv2ServerMessage {
	messageId := s.SignalCache.NextMessageId()
	msg.Sequencer = &livekit.Sequencer{
		MessageId: messageId,
	}
	lprmi := s.SignalCache.Add(msg, messageId)
	msg.Sequencer.LastProcessedRemoteMessageId = lprmi
	return msg
}

func (s *Signalv2ServerMessageCache) GetFromFront() []*livekit.Signalv2ServerMessage {
	return s.recast(s.SignalCache.GetFromFront())
}

func (s *Signalv2ServerMessageCache) ClearAndGetFrom(from uint32) []*livekit.Signalv2ServerMessage {
	return s.recast(s.SignalCache.ClearAndGetFrom(from))
}

func (s *Signalv2ServerMessageCache) recast(protoMessages []proto.Message, lprmi uint32) []*livekit.Signalv2ServerMessage {
	msgs := make([]*livekit.Signalv2ServerMessage, 0, len(protoMessages))
	for _, protoMessage := range protoMessages {
		sm := protoMessage.(*livekit.Signalv2ServerMessage)
		sm.Sequencer.LastProcessedRemoteMessageId = lprmi
		msgs = append(msgs, sm)
	}

	return msgs
}
