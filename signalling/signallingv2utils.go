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
	"github.com/livekit/protocol/livekit"
)

func GetConnectRequest(wireMessage *livekit.Signalv2WireMessage) *livekit.ConnectRequest {
	if wireMessage != nil {
		switch msg := wireMessage.GetMessage().(type) {
		case *livekit.Signalv2WireMessage_Envelope:
			for _, innerMsg := range msg.Envelope.GetClientMessages() {
				switch clientMessage := innerMsg.GetMessage().(type) {
				case *livekit.Signalv2ClientMessage_ConnectRequest:
					return clientMessage.ConnectRequest
				}
			}
		}
	}

	return nil
}

func WithoutConnectRequest(wireMessage *livekit.Signalv2WireMessage) *livekit.Signalv2WireMessage {
	var strippedWireMessage *livekit.Signalv2WireMessage
	if wireMessage != nil {
		switch msg := wireMessage.GetMessage().(type) {
		case *livekit.Signalv2WireMessage_Envelope:
			clientMessages := make(
				[]*livekit.Signalv2ClientMessage,
				0,
				len(msg.Envelope.GetClientMessages()),
			)
			for _, clientMessage := range msg.Envelope.GetClientMessages() {
				switch clientMessage.GetMessage().(type) {
				case *livekit.Signalv2ClientMessage_ConnectRequest:
				default:
					clientMessages = append(clientMessages, clientMessage)
				}
			}

			if len(clientMessages) != 0 {
				strippedWireMessage = &livekit.Signalv2WireMessage{
					Message: &livekit.Signalv2WireMessage_Envelope{
						Envelope: &livekit.Envelope{
							ClientMessages: clientMessages,
						},
					},
				}
			}
		}
	}

	return strippedWireMessage
}

func GetConnectResponse(wireMessage *livekit.Signalv2WireMessage) *livekit.ConnectResponse {
	if wireMessage != nil {
		switch msg := wireMessage.GetMessage().(type) {
		case *livekit.Signalv2WireMessage_Envelope:
			for _, innerMsg := range msg.Envelope.GetServerMessages() {
				switch serverMessage := innerMsg.GetMessage().(type) {
				case *livekit.Signalv2ServerMessage_ConnectResponse:
					return serverMessage.ConnectResponse
				}
			}
		}
	}

	return nil
}
