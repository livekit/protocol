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

package rpc

import (
	"context"
	"fmt"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/psrpc"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type TypedSignalClient = SignalClient[livekit.NodeID]
type TypedSignalServer = SignalServer[livekit.NodeID]

func NewTypedSignalClient(nodeID livekit.NodeID, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedSignalClient, error) {
	return NewSignalClient[livekit.NodeID](bus, append(opts[:len(opts):len(opts)], psrpc.WithClientID(string(nodeID)))...)
}

func NewTypedSignalServer(nodeID livekit.NodeID, svc SignalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedSignalServer, error) {
	return NewSignalServer[livekit.NodeID](svc, bus, append(opts[:len(opts):len(opts)], psrpc.WithServerID(string(nodeID)))...)
}

type ParticipantTopic string
type RoomTopic string

func FormatParticipantTopic(roomName livekit.RoomName, identity livekit.ParticipantIdentity) ParticipantTopic {
	return ParticipantTopic(fmt.Sprintf("%s_%s", roomName, identity))
}

func FormatRoomTopic(roomName livekit.RoomName) RoomTopic {
	return RoomTopic(roomName)
}

type TopicFormatter interface {
	ParticipantTopic(ctx context.Context, roomName livekit.RoomName, identity livekit.ParticipantIdentity) ParticipantTopic
	RoomTopic(ctx context.Context, roomName livekit.RoomName) RoomTopic
}

//counterfeiter:generate . TypedRoomClient
type TypedRoomClient = RoomClient[ParticipantTopic, RoomTopic]
type TypedRoomServer = RoomServer[ParticipantTopic, RoomTopic]

func NewTypedRoomClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedRoomClient, error) {
	return NewRoomClient[ParticipantTopic, RoomTopic](bus, opts...)
}

func NewTypedRoomServer(svc RoomServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedRoomServer, error) {
	return NewRoomServer[ParticipantTopic, RoomTopic](svc, bus, opts...)
}
