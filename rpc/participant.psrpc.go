// Code generated by protoc-gen-psrpc v0.5.0, DO NOT EDIT.
// source: rpc/participant.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/rand"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import livekit "github.com/livekit/protocol/livekit"
import livekit4 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_5

// ============================
// Participant Client Interface
// ============================

type ParticipantClient[ParticipantTopicType ~string] interface {
	RemoveParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit4.RoomParticipantIdentity, opts ...psrpc.RequestOption) (*livekit4.RemoveParticipantResponse, error)

	MutePublishedTrack(ctx context.Context, participant ParticipantTopicType, req *livekit4.MuteRoomTrackRequest, opts ...psrpc.RequestOption) (*livekit4.MuteRoomTrackResponse, error)

	UpdateParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit4.UpdateParticipantRequest, opts ...psrpc.RequestOption) (*livekit.ParticipantInfo, error)

	UpdateSubscriptions(ctx context.Context, participant ParticipantTopicType, req *livekit4.UpdateSubscriptionsRequest, opts ...psrpc.RequestOption) (*livekit4.UpdateSubscriptionsResponse, error)
}

// ================================
// Participant ServerImpl Interface
// ================================

type ParticipantServerImpl interface {
	RemoveParticipant(context.Context, *livekit4.RoomParticipantIdentity) (*livekit4.RemoveParticipantResponse, error)

	MutePublishedTrack(context.Context, *livekit4.MuteRoomTrackRequest) (*livekit4.MuteRoomTrackResponse, error)

	UpdateParticipant(context.Context, *livekit4.UpdateParticipantRequest) (*livekit.ParticipantInfo, error)

	UpdateSubscriptions(context.Context, *livekit4.UpdateSubscriptionsRequest) (*livekit4.UpdateSubscriptionsResponse, error)
}

// ============================
// Participant Server Interface
// ============================

type ParticipantServer[ParticipantTopicType ~string] interface {
	RegisterRemoveParticipantTopic(participant ParticipantTopicType) error
	DeregisterRemoveParticipantTopic(participant ParticipantTopicType)
	RegisterMutePublishedTrackTopic(participant ParticipantTopicType) error
	DeregisterMutePublishedTrackTopic(participant ParticipantTopicType)
	RegisterUpdateParticipantTopic(participant ParticipantTopicType) error
	DeregisterUpdateParticipantTopic(participant ParticipantTopicType)
	RegisterUpdateSubscriptionsTopic(participant ParticipantTopicType) error
	DeregisterUpdateSubscriptionsTopic(participant ParticipantTopicType)
	RegisterAllParticipantTopics(participant ParticipantTopicType) error
	DeregisterAllParticipantTopics(participant ParticipantTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ==================
// Participant Client
// ==================

type participantClient[ParticipantTopicType ~string] struct {
	client *client.RPCClient
}

// NewParticipantClient creates a psrpc client that implements the ParticipantClient interface.
func NewParticipantClient[ParticipantTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (ParticipantClient[ParticipantTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Participant",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("RemoveParticipant", false, false, true, true)
	sd.RegisterMethod("MutePublishedTrack", false, false, true, true)
	sd.RegisterMethod("UpdateParticipant", false, false, true, true)
	sd.RegisterMethod("UpdateSubscriptions", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &participantClient[ParticipantTopicType]{
		client: rpcClient,
	}, nil
}

func (c *participantClient[ParticipantTopicType]) RemoveParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit4.RoomParticipantIdentity, opts ...psrpc.RequestOption) (*livekit4.RemoveParticipantResponse, error) {
	return client.RequestSingle[*livekit4.RemoveParticipantResponse](ctx, c.client, "RemoveParticipant", []string{string(participant)}, req, opts...)
}

func (c *participantClient[ParticipantTopicType]) MutePublishedTrack(ctx context.Context, participant ParticipantTopicType, req *livekit4.MuteRoomTrackRequest, opts ...psrpc.RequestOption) (*livekit4.MuteRoomTrackResponse, error) {
	return client.RequestSingle[*livekit4.MuteRoomTrackResponse](ctx, c.client, "MutePublishedTrack", []string{string(participant)}, req, opts...)
}

func (c *participantClient[ParticipantTopicType]) UpdateParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit4.UpdateParticipantRequest, opts ...psrpc.RequestOption) (*livekit.ParticipantInfo, error) {
	return client.RequestSingle[*livekit.ParticipantInfo](ctx, c.client, "UpdateParticipant", []string{string(participant)}, req, opts...)
}

func (c *participantClient[ParticipantTopicType]) UpdateSubscriptions(ctx context.Context, participant ParticipantTopicType, req *livekit4.UpdateSubscriptionsRequest, opts ...psrpc.RequestOption) (*livekit4.UpdateSubscriptionsResponse, error) {
	return client.RequestSingle[*livekit4.UpdateSubscriptionsResponse](ctx, c.client, "UpdateSubscriptions", []string{string(participant)}, req, opts...)
}

// ==================
// Participant Server
// ==================

type participantServer[ParticipantTopicType ~string] struct {
	svc ParticipantServerImpl
	rpc *server.RPCServer
}

// NewParticipantServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewParticipantServer[ParticipantTopicType ~string](svc ParticipantServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (ParticipantServer[ParticipantTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Participant",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("RemoveParticipant", false, false, true, true)
	sd.RegisterMethod("MutePublishedTrack", false, false, true, true)
	sd.RegisterMethod("UpdateParticipant", false, false, true, true)
	sd.RegisterMethod("UpdateSubscriptions", false, false, true, true)
	return &participantServer[ParticipantTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *participantServer[ParticipantTopicType]) RegisterRemoveParticipantTopic(participant ParticipantTopicType) error {
	return server.RegisterHandler(s.rpc, "RemoveParticipant", []string{string(participant)}, s.svc.RemoveParticipant, nil)
}

func (s *participantServer[ParticipantTopicType]) DeregisterRemoveParticipantTopic(participant ParticipantTopicType) {
	s.rpc.DeregisterHandler("RemoveParticipant", []string{string(participant)})
}

func (s *participantServer[ParticipantTopicType]) RegisterMutePublishedTrackTopic(participant ParticipantTopicType) error {
	return server.RegisterHandler(s.rpc, "MutePublishedTrack", []string{string(participant)}, s.svc.MutePublishedTrack, nil)
}

func (s *participantServer[ParticipantTopicType]) DeregisterMutePublishedTrackTopic(participant ParticipantTopicType) {
	s.rpc.DeregisterHandler("MutePublishedTrack", []string{string(participant)})
}

func (s *participantServer[ParticipantTopicType]) RegisterUpdateParticipantTopic(participant ParticipantTopicType) error {
	return server.RegisterHandler(s.rpc, "UpdateParticipant", []string{string(participant)}, s.svc.UpdateParticipant, nil)
}

func (s *participantServer[ParticipantTopicType]) DeregisterUpdateParticipantTopic(participant ParticipantTopicType) {
	s.rpc.DeregisterHandler("UpdateParticipant", []string{string(participant)})
}

func (s *participantServer[ParticipantTopicType]) RegisterUpdateSubscriptionsTopic(participant ParticipantTopicType) error {
	return server.RegisterHandler(s.rpc, "UpdateSubscriptions", []string{string(participant)}, s.svc.UpdateSubscriptions, nil)
}

func (s *participantServer[ParticipantTopicType]) DeregisterUpdateSubscriptionsTopic(participant ParticipantTopicType) {
	s.rpc.DeregisterHandler("UpdateSubscriptions", []string{string(participant)})
}

func (s *participantServer[ParticipantTopicType]) allParticipantTopicRegisterers() server.RegistererSlice {
	return server.RegistererSlice{
		server.NewRegisterer(s.RegisterRemoveParticipantTopic, s.DeregisterRemoveParticipantTopic),
		server.NewRegisterer(s.RegisterMutePublishedTrackTopic, s.DeregisterMutePublishedTrackTopic),
		server.NewRegisterer(s.RegisterUpdateParticipantTopic, s.DeregisterUpdateParticipantTopic),
		server.NewRegisterer(s.RegisterUpdateSubscriptionsTopic, s.DeregisterUpdateSubscriptionsTopic),
	}
}

func (s *participantServer[ParticipantTopicType]) RegisterAllParticipantTopics(participant ParticipantTopicType) error {
	return s.allParticipantTopicRegisterers().Register(participant)
}

func (s *participantServer[ParticipantTopicType]) DeregisterAllParticipantTopics(participant ParticipantTopicType) {
	s.allParticipantTopicRegisterers().Deregister(participant)
}

func (s *participantServer[ParticipantTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *participantServer[ParticipantTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor5 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x09, 0x8a, 0x87, 0x2d, 0x82, 0x5d, 0x15, 0x4a, 0x50, 0xa9, 0xb5, 0x27, 0x91, 0x04,
	0xf4, 0x0d, 0xbc, 0x79, 0x10, 0x4a, 0xd4, 0x8b, 0x17, 0x49, 0x36, 0x63, 0xbb, 0xe4, 0xcf, 0x8c,
	0xbb, 0x93, 0x40, 0x4f, 0xde, 0x04, 0x7d, 0x1c, 0x9f, 0x50, 0xda, 0x34, 0x69, 0x5a, 0x51, 0xda,
	0x53, 0xc8, 0xef, 0x9b, 0xf9, 0xbe, 0x9d, 0xd9, 0x15, 0xc7, 0x86, 0x94, 0x4f, 0xa1, 0x61, 0xad,
	0x34, 0x85, 0x39, 0x7b, 0x64, 0x90, 0x51, 0xee, 0x18, 0x52, 0xee, 0x3e, 0x12, 0x6b, 0xcc, 0x6d,
	0xc5, 0xdc, 0xa3, 0x54, 0x97, 0x90, 0x68, 0x7e, 0xc9, 0x30, 0x86, 0xb4, 0xa6, 0xb2, 0xa6, 0x06,
	0x31, 0xab, 0xd8, 0xf5, 0xd7, 0xae, 0xe8, 0x8c, 0x96, 0x9e, 0xf2, 0x5d, 0x74, 0x03, 0xc8, 0xb0,
	0x84, 0x36, 0xec, 0x7b, 0x8b, 0x4e, 0x2f, 0x40, 0xcc, 0x5a, 0xca, 0x5d, 0x0c, 0x39, 0x6b, 0x9e,
	0xba, 0x83, 0x65, 0xc5, 0x7a, 0x77, 0x00, 0x96, 0x30, 0xb7, 0x30, 0x18, 0x7e, 0x7f, 0x3a, 0xfd,
	0x03, 0xc7, 0x3d, 0x11, 0x9d, 0xd6, 0x14, 0xb2, 0xfd, 0xd3, 0x73, 0xe4, 0x54, 0xc8, 0xfb, 0x82,
	0x61, 0x54, 0x44, 0xa9, 0xb6, 0x13, 0x88, 0x1f, 0x4d, 0xa8, 0x12, 0x79, 0xda, 0xf8, 0xcf, 0xc4,
	0xd9, 0x29, 0xe6, 0x3c, 0x80, 0xb7, 0x02, 0x2c, 0xbb, 0x67, 0x7f, 0xc9, 0x5b, 0x45, 0x97, 0xa2,
	0xfb, 0x44, 0x71, 0xc8, 0x2b, 0xb3, 0x9f, 0x37, 0xd6, 0xbf, 0xb4, 0x3a, 0xbd, 0xd7, 0x94, 0xb4,
	0x57, 0x93, 0xbf, 0xe2, 0x86, 0xb9, 0x1f, 0x8e, 0x38, 0xac, 0xcc, 0x1f, 0x8a, 0xc8, 0x2a, 0xa3,
	0xab, 0xbb, 0x94, 0x17, 0x6b, 0xd1, 0x2b, 0x6a, 0x1d, 0x3e, 0xfc, 0xbf, 0x68, 0x9b, 0x05, 0xdc,
	0x5e, 0x3d, 0x5f, 0x8e, 0x35, 0x4f, 0x8a, 0xc8, 0x53, 0x98, 0xf9, 0x0b, 0xdf, 0xe6, 0x4b, 0xc9,
	0xd8, 0xb7, 0x60, 0x4a, 0xad, 0xc0, 0x37, 0xa4, 0xa2, 0xbd, 0xf9, 0x0b, 0xba, 0xf9, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x0d, 0x71, 0xfb, 0x47, 0x98, 0x02, 0x00, 0x00,
}
