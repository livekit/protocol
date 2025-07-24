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
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
)

func TestSignalCache(t *testing.T) {
	firstMessageId := uint32(10)
	lastProcessedRemoteMessageId := uint32(2345)
	cache := NewSignalCache(SignalCacheParams{
		FirstMessageId: firstMessageId,
	})

	inputMessages := []*livekit.Signalv2ServerMessage{
		&livekit.Signalv2ServerMessage{
			Message: &livekit.Signalv2ServerMessage_ConnectResponse{},
		},
		&livekit.Signalv2ServerMessage{
			Message: &livekit.Signalv2ServerMessage_PublisherSdp{},
		},
		&livekit.Signalv2ServerMessage{
			Message: &livekit.Signalv2ServerMessage_SubscriberSdp{},
		},
		&livekit.Signalv2ServerMessage{
			Message: &livekit.Signalv2ServerMessage_RoomUpdate{},
		},
	}

	expectedOutputMessages := []*livekit.Signalv2ServerMessage{
		&livekit.Signalv2ServerMessage{
			Sequencer: &livekit.Sequencer{
				MessageId:                    firstMessageId,
				LastProcessedRemoteMessageId: lastProcessedRemoteMessageId,
			},
			Message: &livekit.Signalv2ServerMessage_ConnectResponse{},
		},
		&livekit.Signalv2ServerMessage{
			Sequencer: &livekit.Sequencer{
				MessageId:                    firstMessageId + 1,
				LastProcessedRemoteMessageId: lastProcessedRemoteMessageId,
			},
			Message: &livekit.Signalv2ServerMessage_PublisherSdp{},
		},
		&livekit.Signalv2ServerMessage{
			Sequencer: &livekit.Sequencer{
				MessageId:                    firstMessageId + 2,
				LastProcessedRemoteMessageId: lastProcessedRemoteMessageId,
			},
			Message: &livekit.Signalv2ServerMessage_SubscriberSdp{},
		},
		&livekit.Signalv2ServerMessage{
			Sequencer: &livekit.Sequencer{
				MessageId:                    firstMessageId + 3,
				LastProcessedRemoteMessageId: lastProcessedRemoteMessageId,
			},
			Message: &livekit.Signalv2ServerMessage_RoomUpdate{},
		},
	}

	cache.SetLastProcessedRemoteMessageId(lastProcessedRemoteMessageId)

	// add messages to cache
	for _, inputMessage := range inputMessages {
		messageId := cache.NextMessageId()
		inputMessage.Sequencer = &livekit.Sequencer{
			MessageId: messageId,
		}
		cache.Add(inputMessage, messageId)
	}

	// get all messages in cache
	protoMessages, lprmi := cache.GetFromFront()
	outputMessages := make([]*livekit.Signalv2ServerMessage, 0, len(protoMessages))
	for _, protoMessage := range protoMessages {
		if sm, ok := protoMessage.(*livekit.Signalv2ServerMessage); ok {
			sm.Sequencer.LastProcessedRemoteMessageId = lprmi
			outputMessages = append(outputMessages, sm)
		}
	}
	require.True(t, compareProtoSlices(expectedOutputMessages, outputMessages))
	require.Equal(t, lastProcessedRemoteMessageId, lprmi)

	// clear one and get again
	cache.Clear(firstMessageId)

	protoMessages, lprmi = cache.GetFromFront()
	outputMessages = make([]*livekit.Signalv2ServerMessage, 0, len(protoMessages))
	for _, protoMessage := range protoMessages {
		if sm, ok := protoMessage.(*livekit.Signalv2ServerMessage); ok {
			sm.Sequencer.LastProcessedRemoteMessageId = lprmi
			outputMessages = append(outputMessages, sm)
		}
	}
	require.True(t, compareProtoSlices(expectedOutputMessages[1:], outputMessages))
	require.Equal(t, lastProcessedRemoteMessageId, lprmi)

	// clearing some evicted messages should not clear anything
	cache.Clear(firstMessageId) // firstMessageId has been cleared already at this point

	protoMessages, lprmi = cache.GetFromFront()
	outputMessages = make([]*livekit.Signalv2ServerMessage, 0, len(protoMessages))
	for _, protoMessage := range protoMessages {
		if sm, ok := protoMessage.(*livekit.Signalv2ServerMessage); ok {
			sm.Sequencer.LastProcessedRemoteMessageId = lprmi
			outputMessages = append(outputMessages, sm)
		}
	}
	require.True(t, compareProtoSlices(expectedOutputMessages[1:], outputMessages))
	require.Equal(t, lastProcessedRemoteMessageId, lprmi)

	// clear some and get rest in one go
	protoMessages, lprmi = cache.ClearAndGetFrom(firstMessageId + 3)
	outputMessages = make([]*livekit.Signalv2ServerMessage, 0, len(protoMessages))
	for _, protoMessage := range protoMessages {
		if sm, ok := protoMessage.(*livekit.Signalv2ServerMessage); ok {
			sm.Sequencer.LastProcessedRemoteMessageId = lprmi
			outputMessages = append(outputMessages, sm)
		}
	}
	require.Equal(t, 1, len(outputMessages))
	require.True(t, compareProtoSlices(expectedOutputMessages[3:], outputMessages))
	require.Equal(t, lastProcessedRemoteMessageId, lprmi)

	// getting again should get the same messages again as they sill should in cache
	protoMessages, lprmi = cache.GetFromFront()
	outputMessages = make([]*livekit.Signalv2ServerMessage, 0, len(protoMessages))
	for _, protoMessage := range protoMessages {
		if sm, ok := protoMessage.(*livekit.Signalv2ServerMessage); ok {
			sm.Sequencer.LastProcessedRemoteMessageId = lprmi
			outputMessages = append(outputMessages, sm)
		}
	}
	require.True(t, compareProtoSlices(expectedOutputMessages[3:], outputMessages))
	require.Equal(t, lastProcessedRemoteMessageId, lprmi)

	// clearing all and getting should return nil
	protoMessages, lprmi = cache.ClearAndGetFrom(firstMessageId + uint32(len(inputMessages)))
	require.Nil(t, protoMessages)

	// getting again should return nil as the cache is fully cleared above
	protoMessages, lprmi = cache.GetFromFront()
	require.Nil(t, protoMessages)

	lastProcessedRemoteMessageId = 4567
	cache.SetLastProcessedRemoteMessageId(lastProcessedRemoteMessageId)

	expectedOutputMessages = []*livekit.Signalv2ServerMessage{
		&livekit.Signalv2ServerMessage{
			Sequencer: &livekit.Sequencer{
				MessageId:                    firstMessageId + 4,
				LastProcessedRemoteMessageId: lastProcessedRemoteMessageId,
			},
			Message: &livekit.Signalv2ServerMessage_ConnectResponse{},
		},
		&livekit.Signalv2ServerMessage{
			Sequencer: &livekit.Sequencer{
				MessageId:                    firstMessageId + 1 + 4,
				LastProcessedRemoteMessageId: lastProcessedRemoteMessageId,
			},
			Message: &livekit.Signalv2ServerMessage_PublisherSdp{},
		},
		&livekit.Signalv2ServerMessage{
			Sequencer: &livekit.Sequencer{
				MessageId:                    firstMessageId + 2 + 4,
				LastProcessedRemoteMessageId: lastProcessedRemoteMessageId,
			},
			Message: &livekit.Signalv2ServerMessage_SubscriberSdp{},
		},
		&livekit.Signalv2ServerMessage{
			Sequencer: &livekit.Sequencer{
				MessageId:                    firstMessageId + 3 + 4,
				LastProcessedRemoteMessageId: lastProcessedRemoteMessageId,
			},
			Message: &livekit.Signalv2ServerMessage_RoomUpdate{},
		},
	}

	// add more messages which should increment the message id and add them in scrambled message id order
	for _, inputMessage := range inputMessages {
		messageId := cache.NextMessageId()
		inputMessage.Sequencer = &livekit.Sequencer{
			MessageId: messageId,
		}
	}
	for _, idx := range []int{1, 3, 2, 0} {
		cache.Add(inputMessages[idx], firstMessageId+4+uint32(idx))
	}

	// get all messages in cache
	protoMessages, lprmi = cache.GetFromFront()
	outputMessages = make([]*livekit.Signalv2ServerMessage, 0, len(protoMessages))
	for _, protoMessage := range protoMessages {
		if sm, ok := protoMessage.(*livekit.Signalv2ServerMessage); ok {
			sm.Sequencer.LastProcessedRemoteMessageId = lprmi
			outputMessages = append(outputMessages, sm)
		}
	}
	require.True(t, compareProtoSlices(expectedOutputMessages, outputMessages))
}

func compareProtoSlices(a []*livekit.Signalv2ServerMessage, b []*livekit.Signalv2ServerMessage) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !proto.Equal(a[i], b[i]) {
			return false
		}
	}

	return true
}
