package rpc

import (
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/psrpc"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type TypedSignalClient = SignalClient[livekit.NodeID]
type TypedSignalServer = SignalServer[livekit.NodeID]

func NewTypedSignalClient(nodeID livekit.NodeID, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedSignalClient, error) {
	return NewSignalClient[livekit.NodeID](string(nodeID), bus, opts...)
}

func NewTypedSignalServer(nodeID livekit.NodeID, svc SignalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedSignalServer, error) {
	return NewSignalServer[livekit.NodeID](string(nodeID), svc, bus, opts...)
}

//counterfeiter:generate . TypedRoomClient
type TypedRoomClient = RoomClient[livekit.ParticipantTopic, livekit.RoomTopic]
type TypedRoomServer = RoomServer[livekit.ParticipantTopic, livekit.RoomTopic]

func NewTypedRoomClient(nodeID livekit.NodeID, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedRoomClient, error) {
	return NewRoomClient[livekit.ParticipantTopic, livekit.RoomTopic](string(nodeID), bus, opts...)
}

func NewTypedRoomServer(nodeID livekit.NodeID, svc RoomServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedRoomServer, error) {
	return NewRoomServer[livekit.ParticipantTopic, livekit.RoomTopic](string(nodeID), svc, bus, opts...)
}
