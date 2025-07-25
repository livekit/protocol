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

type Signalv2ClientMessageCache struct {
	*SignalCache
}

func NewSignalv2ClientMessageCache(params SignalCacheParams) *Signalv2ClientMessageCache {
	return &Signalv2ClientMessageCache{
		SignalCache: NewSignalCache(params),
	}
}

func (s *Signalv2ClientMessageCache) Add(msg *livekit.Signalv2ClientMessage) *livekit.Signalv2ClientMessage {
	messageId := s.SignalCache.NextMessageId()
	msg.Sequencer = &livekit.Sequencer{
		MessageId: messageId,
	}
	msg.Sequencer.LastProcessedRemoteMessageId = s.SignalCache.Add(msg, messageId)
	return msg
}

func (s *Signalv2ClientMessageCache) GetFromFront() []*livekit.Signalv2ClientMessage {
	return s.recast(s.SignalCache.GetFromFront())
}

func (s *Signalv2ClientMessageCache) ClearAndGetFrom(from uint32) []*livekit.Signalv2ClientMessage {
	return s.recast(s.SignalCache.ClearAndGetFrom(from))
}

func (s *Signalv2ClientMessageCache) recast(protoMessages []proto.Message, lprmi uint32) []*livekit.Signalv2ClientMessage {
	msgs := make([]*livekit.Signalv2ClientMessage, 0, len(protoMessages))
	for _, protoMessage := range protoMessages {
		sm := protoMessage.(*livekit.Signalv2ClientMessage)
		sm.Sequencer.LastProcessedRemoteMessageId = lprmi
		msgs = append(msgs, sm)
	}

	return msgs
}
