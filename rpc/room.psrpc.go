// Code generated by protoc-gen-psrpc v0.4.0, DO NOT EDIT.
// source: rpc/room.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import livekit "github.com/livekit/protocol/livekit"
import livekit3 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_4

// =====================
// Room Client Interface
// =====================

type RoomClient[ParticipantTopicType, RoomTopicType ~string] interface {
	RemoveParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit3.RoomParticipantIdentity, opts ...psrpc.RequestOption) (*livekit3.RemoveParticipantResponse, error)

	MutePublishedTrack(ctx context.Context, participant ParticipantTopicType, req *livekit3.MuteRoomTrackRequest, opts ...psrpc.RequestOption) (*livekit3.MuteRoomTrackResponse, error)

	UpdateParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit3.UpdateParticipantRequest, opts ...psrpc.RequestOption) (*livekit.ParticipantInfo, error)

	UpdateSubscriptions(ctx context.Context, participant ParticipantTopicType, req *livekit3.UpdateSubscriptionsRequest, opts ...psrpc.RequestOption) (*livekit3.UpdateSubscriptionsResponse, error)

	DeleteRoom(ctx context.Context, room RoomTopicType, req *livekit3.DeleteRoomRequest, opts ...psrpc.RequestOption) (*livekit3.DeleteRoomResponse, error)

	SendData(ctx context.Context, room RoomTopicType, req *livekit3.SendDataRequest, opts ...psrpc.RequestOption) (*livekit3.SendDataResponse, error)

	UpdateRoomMetadata(ctx context.Context, room RoomTopicType, req *livekit3.UpdateRoomMetadataRequest, opts ...psrpc.RequestOption) (*livekit.Room, error)
}

// =========================
// Room ServerImpl Interface
// =========================

type RoomServerImpl interface {
	RemoveParticipant(context.Context, *livekit3.RoomParticipantIdentity) (*livekit3.RemoveParticipantResponse, error)

	MutePublishedTrack(context.Context, *livekit3.MuteRoomTrackRequest) (*livekit3.MuteRoomTrackResponse, error)

	UpdateParticipant(context.Context, *livekit3.UpdateParticipantRequest) (*livekit.ParticipantInfo, error)

	UpdateSubscriptions(context.Context, *livekit3.UpdateSubscriptionsRequest) (*livekit3.UpdateSubscriptionsResponse, error)

	DeleteRoom(context.Context, *livekit3.DeleteRoomRequest) (*livekit3.DeleteRoomResponse, error)

	SendData(context.Context, *livekit3.SendDataRequest) (*livekit3.SendDataResponse, error)

	UpdateRoomMetadata(context.Context, *livekit3.UpdateRoomMetadataRequest) (*livekit.Room, error)
}

// =====================
// Room Server Interface
// =====================

type RoomServer[ParticipantTopicType, RoomTopicType ~string] interface {
	RegisterRemoveParticipantTopic(participant ParticipantTopicType) error
	DeregisterRemoveParticipantTopic(participant ParticipantTopicType)
	RegisterMutePublishedTrackTopic(participant ParticipantTopicType) error
	DeregisterMutePublishedTrackTopic(participant ParticipantTopicType)
	RegisterUpdateParticipantTopic(participant ParticipantTopicType) error
	DeregisterUpdateParticipantTopic(participant ParticipantTopicType)
	RegisterUpdateSubscriptionsTopic(participant ParticipantTopicType) error
	DeregisterUpdateSubscriptionsTopic(participant ParticipantTopicType)
	RegisterDeleteRoomTopic(room RoomTopicType) error
	DeregisterDeleteRoomTopic(room RoomTopicType)
	RegisterSendDataTopic(room RoomTopicType) error
	DeregisterSendDataTopic(room RoomTopicType)
	RegisterUpdateRoomMetadataTopic(room RoomTopicType) error
	DeregisterUpdateRoomMetadataTopic(room RoomTopicType)
	RegisterAllParticipantTopics(participant ParticipantTopicType) error
	DeregisterAllParticipantTopics(participant ParticipantTopicType)
	RegisterAllRoomTopics(room RoomTopicType) error
	DeregisterAllRoomTopics(room RoomTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ===========
// Room Client
// ===========

type roomClient[ParticipantTopicType, RoomTopicType ~string] struct {
	client *client.RPCClient
}

// NewRoomClient creates a psrpc client that implements the RoomClient interface.
func NewRoomClient[ParticipantTopicType, RoomTopicType ~string](clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (RoomClient[ParticipantTopicType, RoomTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Room",
		ID:   clientID,
	}

	sd.RegisterMethod("RemoveParticipant", false, false, true, true)
	sd.RegisterMethod("MutePublishedTrack", false, false, true, true)
	sd.RegisterMethod("UpdateParticipant", false, false, true, true)
	sd.RegisterMethod("UpdateSubscriptions", false, false, true, true)
	sd.RegisterMethod("DeleteRoom", false, false, true, true)
	sd.RegisterMethod("SendData", false, false, true, true)
	sd.RegisterMethod("UpdateRoomMetadata", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &roomClient[ParticipantTopicType, RoomTopicType]{
		client: rpcClient,
	}, nil
}

func (c *roomClient[ParticipantTopicType, RoomTopicType]) RemoveParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit3.RoomParticipantIdentity, opts ...psrpc.RequestOption) (*livekit3.RemoveParticipantResponse, error) {
	return client.RequestSingle[*livekit3.RemoveParticipantResponse](ctx, c.client, "RemoveParticipant", []string{string(participant)}, req, opts...)
}

func (c *roomClient[ParticipantTopicType, RoomTopicType]) MutePublishedTrack(ctx context.Context, participant ParticipantTopicType, req *livekit3.MuteRoomTrackRequest, opts ...psrpc.RequestOption) (*livekit3.MuteRoomTrackResponse, error) {
	return client.RequestSingle[*livekit3.MuteRoomTrackResponse](ctx, c.client, "MutePublishedTrack", []string{string(participant)}, req, opts...)
}

func (c *roomClient[ParticipantTopicType, RoomTopicType]) UpdateParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit3.UpdateParticipantRequest, opts ...psrpc.RequestOption) (*livekit.ParticipantInfo, error) {
	return client.RequestSingle[*livekit.ParticipantInfo](ctx, c.client, "UpdateParticipant", []string{string(participant)}, req, opts...)
}

func (c *roomClient[ParticipantTopicType, RoomTopicType]) UpdateSubscriptions(ctx context.Context, participant ParticipantTopicType, req *livekit3.UpdateSubscriptionsRequest, opts ...psrpc.RequestOption) (*livekit3.UpdateSubscriptionsResponse, error) {
	return client.RequestSingle[*livekit3.UpdateSubscriptionsResponse](ctx, c.client, "UpdateSubscriptions", []string{string(participant)}, req, opts...)
}

func (c *roomClient[ParticipantTopicType, RoomTopicType]) DeleteRoom(ctx context.Context, room RoomTopicType, req *livekit3.DeleteRoomRequest, opts ...psrpc.RequestOption) (*livekit3.DeleteRoomResponse, error) {
	return client.RequestSingle[*livekit3.DeleteRoomResponse](ctx, c.client, "DeleteRoom", []string{string(room)}, req, opts...)
}

func (c *roomClient[ParticipantTopicType, RoomTopicType]) SendData(ctx context.Context, room RoomTopicType, req *livekit3.SendDataRequest, opts ...psrpc.RequestOption) (*livekit3.SendDataResponse, error) {
	return client.RequestSingle[*livekit3.SendDataResponse](ctx, c.client, "SendData", []string{string(room)}, req, opts...)
}

func (c *roomClient[ParticipantTopicType, RoomTopicType]) UpdateRoomMetadata(ctx context.Context, room RoomTopicType, req *livekit3.UpdateRoomMetadataRequest, opts ...psrpc.RequestOption) (*livekit.Room, error) {
	return client.RequestSingle[*livekit.Room](ctx, c.client, "UpdateRoomMetadata", []string{string(room)}, req, opts...)
}

// ===========
// Room Server
// ===========

type roomServer[ParticipantTopicType, RoomTopicType ~string] struct {
	svc RoomServerImpl
	rpc *server.RPCServer
}

// NewRoomServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewRoomServer[ParticipantTopicType, RoomTopicType ~string](serverID string, svc RoomServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (RoomServer[ParticipantTopicType, RoomTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Room",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("RemoveParticipant", false, false, true, true)
	sd.RegisterMethod("MutePublishedTrack", false, false, true, true)
	sd.RegisterMethod("UpdateParticipant", false, false, true, true)
	sd.RegisterMethod("UpdateSubscriptions", false, false, true, true)
	sd.RegisterMethod("DeleteRoom", false, false, true, true)
	sd.RegisterMethod("SendData", false, false, true, true)
	sd.RegisterMethod("UpdateRoomMetadata", false, false, true, true)
	return &roomServer[ParticipantTopicType, RoomTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterRemoveParticipantTopic(participant ParticipantTopicType) error {
	return server.RegisterHandler(s.rpc, "RemoveParticipant", []string{string(participant)}, s.svc.RemoveParticipant, nil)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterRemoveParticipantTopic(participant ParticipantTopicType) {
	s.rpc.DeregisterHandler("RemoveParticipant", []string{string(participant)})
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterMutePublishedTrackTopic(participant ParticipantTopicType) error {
	return server.RegisterHandler(s.rpc, "MutePublishedTrack", []string{string(participant)}, s.svc.MutePublishedTrack, nil)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterMutePublishedTrackTopic(participant ParticipantTopicType) {
	s.rpc.DeregisterHandler("MutePublishedTrack", []string{string(participant)})
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterUpdateParticipantTopic(participant ParticipantTopicType) error {
	return server.RegisterHandler(s.rpc, "UpdateParticipant", []string{string(participant)}, s.svc.UpdateParticipant, nil)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterUpdateParticipantTopic(participant ParticipantTopicType) {
	s.rpc.DeregisterHandler("UpdateParticipant", []string{string(participant)})
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterUpdateSubscriptionsTopic(participant ParticipantTopicType) error {
	return server.RegisterHandler(s.rpc, "UpdateSubscriptions", []string{string(participant)}, s.svc.UpdateSubscriptions, nil)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterUpdateSubscriptionsTopic(participant ParticipantTopicType) {
	s.rpc.DeregisterHandler("UpdateSubscriptions", []string{string(participant)})
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterDeleteRoomTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "DeleteRoom", []string{string(room)}, s.svc.DeleteRoom, nil)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterDeleteRoomTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("DeleteRoom", []string{string(room)})
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterSendDataTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "SendData", []string{string(room)}, s.svc.SendData, nil)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterSendDataTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("SendData", []string{string(room)})
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterUpdateRoomMetadataTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "UpdateRoomMetadata", []string{string(room)}, s.svc.UpdateRoomMetadata, nil)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterUpdateRoomMetadataTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("UpdateRoomMetadata", []string{string(room)})
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) allParticipantTopicRegisterers() server.RegistererSlice {
	return server.RegistererSlice{
		server.NewRegisterer(s.RegisterRemoveParticipantTopic, s.DeregisterRemoveParticipantTopic),
		server.NewRegisterer(s.RegisterMutePublishedTrackTopic, s.DeregisterMutePublishedTrackTopic),
		server.NewRegisterer(s.RegisterUpdateParticipantTopic, s.DeregisterUpdateParticipantTopic),
		server.NewRegisterer(s.RegisterUpdateSubscriptionsTopic, s.DeregisterUpdateSubscriptionsTopic),
	}
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterAllParticipantTopics(participant ParticipantTopicType) error {
	return s.allParticipantTopicRegisterers().Register(participant)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterAllParticipantTopics(participant ParticipantTopicType) {
	s.allParticipantTopicRegisterers().Deregister(participant)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) allRoomTopicRegisterers() server.RegistererSlice {
	return server.RegistererSlice{
		server.NewRegisterer(s.RegisterDeleteRoomTopic, s.DeregisterDeleteRoomTopic),
		server.NewRegisterer(s.RegisterSendDataTopic, s.DeregisterSendDataTopic),
		server.NewRegisterer(s.RegisterUpdateRoomMetadataTopic, s.DeregisterUpdateRoomMetadataTopic),
	}
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) RegisterAllRoomTopics(room RoomTopicType) error {
	return s.allRoomTopicRegisterers().Register(room)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) DeregisterAllRoomTopics(room RoomTopicType) {
	s.allRoomTopicRegisterers().Deregister(room)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *roomServer[ParticipantTopicType, RoomTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor3 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x6a, 0xf2, 0x30,
	0x18, 0xc6, 0x29, 0x9f, 0xc8, 0x47, 0x3e, 0x14, 0xcd, 0x37, 0x86, 0xeb, 0xfe, 0xe0, 0x9c, 0x47,
	0x63, 0xb4, 0xb0, 0xdd, 0xc1, 0xf0, 0x64, 0x07, 0x82, 0xd4, 0x8d, 0xc1, 0x60, 0x48, 0x9b, 0xbe,
	0xd3, 0x60, 0xdb, 0x64, 0x49, 0x5a, 0xf0, 0x68, 0x67, 0x03, 0xef, 0x61, 0x57, 0xe1, 0x15, 0x8e,
	0x5a, 0xd3, 0x56, 0xa7, 0xdb, 0x3c, 0x69, 0xe9, 0xf3, 0xbc, 0xef, 0xf3, 0x4b, 0xde, 0x26, 0xa8,
	0x2e, 0x38, 0xb1, 0x05, 0x63, 0xa1, 0xc5, 0x05, 0x53, 0x0c, 0xff, 0x11, 0x9c, 0x98, 0x35, 0xc6,
	0x15, 0x65, 0x91, 0xcc, 0x34, 0xf3, 0x20, 0xa0, 0x09, 0x4c, 0xa9, 0x1a, 0x85, 0xcc, 0x87, 0x40,
	0xab, 0x58, 0xab, 0x45, 0xf7, 0xf5, 0x47, 0x15, 0x55, 0x1c, 0xc6, 0x42, 0xfc, 0x86, 0x9a, 0x0e,
	0x84, 0x2c, 0x81, 0x81, 0x2b, 0x14, 0x25, 0x94, 0xbb, 0x91, 0xc2, 0x6d, 0x6b, 0xd5, 0x62, 0xa5,
	0x35, 0x25, 0xe7, 0xce, 0x87, 0x48, 0x51, 0x35, 0x33, 0x3b, 0x45, 0xc5, 0x66, 0xb7, 0x03, 0x92,
	0xb3, 0x48, 0x42, 0xa7, 0xbb, 0x98, 0x1b, 0xed, 0x86, 0x61, 0x9e, 0xa0, 0x7f, 0xbc, 0x14, 0x5e,
	0xfe, 0x68, 0x19, 0x78, 0x86, 0x70, 0x3f, 0x56, 0x30, 0x88, 0xbd, 0x80, 0xca, 0x09, 0xf8, 0xf7,
	0xc2, 0x25, 0x53, 0x7c, 0x9a, 0xe7, 0xa7, 0x66, 0xba, 0x8a, 0xa5, 0xee, 0xc0, 0x6b, 0x0c, 0x52,
	0x99, 0x67, 0xbb, 0xec, 0xbd, 0xd0, 0x09, 0x6a, 0x3e, 0x70, 0xdf, 0x55, 0x6b, 0x7b, 0x3f, 0xcf,
	0xa3, 0xbf, 0x78, 0x9a, 0xde, 0xca, 0x4b, 0xca, 0xa3, 0x89, 0x5e, 0xd8, 0x2f, 0xb9, 0xef, 0x06,
	0xfa, 0x9f, 0x85, 0x0f, 0x63, 0x4f, 0x12, 0x41, 0xb3, 0x9f, 0x88, 0x2f, 0x36, 0xd0, 0x6b, 0xae,
	0x86, 0x77, 0xbf, 0x2f, 0xda, 0x6b, 0x00, 0xcf, 0x08, 0xf5, 0x20, 0x80, 0x6c, 0x82, 0xd8, 0xcc,
	0x93, 0x0b, 0x51, 0x53, 0x8f, 0xb7, 0x7a, 0x2b, 0xd8, 0xe1, 0x62, 0x6e, 0xe0, 0x86, 0x61, 0xd6,
	0x51, 0x25, 0x3d, 0x62, 0x78, 0xf9, 0x6c, 0x19, 0xf8, 0x11, 0xfd, 0x1d, 0x42, 0xe4, 0xf7, 0x5c,
	0xe5, 0xe2, 0x62, 0x66, 0x5a, 0xd2, 0xd1, 0x47, 0x5b, 0x9c, 0x1f, 0x82, 0x47, 0x08, 0x67, 0x9b,
	0x4f, 0x97, 0xd1, 0x07, 0xe5, 0xfa, 0x29, 0xa2, 0xb3, 0x31, 0x99, 0xb2, 0xa9, 0x61, 0xb5, 0xb5,
	0x93, 0xbd, 0x0b, 0x70, 0x7b, 0xf5, 0x74, 0x39, 0xa6, 0x6a, 0x12, 0x7b, 0x16, 0x61, 0xa1, 0xbd,
	0x6a, 0xc9, 0xdf, 0x7c, 0x3a, 0xb6, 0x25, 0x88, 0x84, 0x12, 0xb0, 0x05, 0x27, 0x5e, 0x75, 0x79,
	0xa7, 0x6e, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xc1, 0x20, 0xb0, 0xa3, 0x03, 0x00, 0x00,
}
