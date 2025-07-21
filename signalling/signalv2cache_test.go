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

func TestSignalv2Cache(t *testing.T) {
	firstMessageId := uint32(10)
	lastProcessedRemoteMessageId := uint32(2345)
	cache := NewSignalv2Cache(Signalv2CacheParams{
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

	// Add() - add one message at a time
	for _, inputMessage := range inputMessages {
		cache.Add(inputMessage)
	}

	// get all messages in cache
	outputMessages := cache.GetFromFront()
	require.True(t, compareProtoSlices(expectedOutputMessages, outputMessages))

	// clear one and get again
	cache.Clear(firstMessageId)

	outputMessages = cache.GetFromFront()
	require.True(t, compareProtoSlices(expectedOutputMessages[1:], outputMessages))

	// clearing some evicted messages should not clear anything
	cache.Clear(firstMessageId) // firstMessageId has been cleared already at this point

	outputMessages = cache.GetFromFront()
	require.True(t, compareProtoSlices(expectedOutputMessages[1:], outputMessages))

	// clear some and get rest in one go
	outputMessages = cache.ClearAndGetFrom(firstMessageId + 3)
	require.Equal(t, 1, len(outputMessages))
	require.True(t, compareProtoSlices(expectedOutputMessages[3:], outputMessages))

	// getting again should get the same messages again as they sill should in cache
	outputMessages = cache.GetFromFront()
	require.True(t, compareProtoSlices(expectedOutputMessages[3:], outputMessages))

	// clearing all and getting should return nil
	require.Nil(t, cache.ClearAndGetFrom(firstMessageId+uint32(len(inputMessages))))

	// getting again should return nil as the cache is fully cleared above
	require.Nil(t, cache.GetFromFront())

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

	// AddBatch() - add all messages at once
	cache.AddBatch(inputMessages)

	// get all messages in cache
	outputMessages = cache.GetFromFront()
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
