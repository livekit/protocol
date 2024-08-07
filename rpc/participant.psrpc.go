// Code generated by protoc-gen-psrpc v0.5.1, DO NOT EDIT.
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
import livekit1 "github.com/livekit/protocol/livekit"
import livekit6 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_5

// ============================
// Participant Client Interface
// ============================

type ParticipantClient[ParticipantTopicType ~string] interface {
	RemoveParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit6.RoomParticipantIdentity, opts ...psrpc.RequestOption) (*livekit6.RemoveParticipantResponse, error)

	MutePublishedTrack(ctx context.Context, participant ParticipantTopicType, req *livekit6.MuteRoomTrackRequest, opts ...psrpc.RequestOption) (*livekit6.MuteRoomTrackResponse, error)

	UpdateParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit6.UpdateParticipantRequest, opts ...psrpc.RequestOption) (*livekit1.ParticipantInfo, error)

	UpdateSubscriptions(ctx context.Context, participant ParticipantTopicType, req *livekit6.UpdateSubscriptionsRequest, opts ...psrpc.RequestOption) (*livekit6.UpdateSubscriptionsResponse, error)
}

// ================================
// Participant ServerImpl Interface
// ================================

type ParticipantServerImpl interface {
	RemoveParticipant(context.Context, *livekit6.RoomParticipantIdentity) (*livekit6.RemoveParticipantResponse, error)

	MutePublishedTrack(context.Context, *livekit6.MuteRoomTrackRequest) (*livekit6.MuteRoomTrackResponse, error)

	UpdateParticipant(context.Context, *livekit6.UpdateParticipantRequest) (*livekit1.ParticipantInfo, error)

	UpdateSubscriptions(context.Context, *livekit6.UpdateSubscriptionsRequest) (*livekit6.UpdateSubscriptionsResponse, error)
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

func (c *participantClient[ParticipantTopicType]) RemoveParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit6.RoomParticipantIdentity, opts ...psrpc.RequestOption) (*livekit6.RemoveParticipantResponse, error) {
	return client.RequestSingle[*livekit6.RemoveParticipantResponse](ctx, c.client, "RemoveParticipant", []string{string(participant)}, req, opts...)
}

func (c *participantClient[ParticipantTopicType]) MutePublishedTrack(ctx context.Context, participant ParticipantTopicType, req *livekit6.MuteRoomTrackRequest, opts ...psrpc.RequestOption) (*livekit6.MuteRoomTrackResponse, error) {
	return client.RequestSingle[*livekit6.MuteRoomTrackResponse](ctx, c.client, "MutePublishedTrack", []string{string(participant)}, req, opts...)
}

func (c *participantClient[ParticipantTopicType]) UpdateParticipant(ctx context.Context, participant ParticipantTopicType, req *livekit6.UpdateParticipantRequest, opts ...psrpc.RequestOption) (*livekit1.ParticipantInfo, error) {
	return client.RequestSingle[*livekit1.ParticipantInfo](ctx, c.client, "UpdateParticipant", []string{string(participant)}, req, opts...)
}

func (c *participantClient[ParticipantTopicType]) UpdateSubscriptions(ctx context.Context, participant ParticipantTopicType, req *livekit6.UpdateSubscriptionsRequest, opts ...psrpc.RequestOption) (*livekit6.UpdateSubscriptionsResponse, error) {
	return client.RequestSingle[*livekit6.UpdateSubscriptionsResponse](ctx, c.client, "UpdateSubscriptions", []string{string(participant)}, req, opts...)
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
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x09, 0x8a, 0x87, 0x2d, 0x82, 0x5d, 0x15, 0x4a, 0xf0, 0x4f, 0x5b, 0x7b, 0x4e, 0x40,
	0xdf, 0xc0, 0x9b, 0x07, 0xa1, 0x44, 0xbd, 0x78, 0x91, 0x64, 0x33, 0xda, 0xa5, 0xc9, 0xce, 0xba,
	0x3b, 0x5b, 0xe8, 0xc9, 0x9b, 0xa0, 0x8f, 0xe3, 0x13, 0x4a, 0x93, 0x26, 0xdd, 0x56, 0x14, 0x7b,
	0x9c, 0xdf, 0x37, 0xf3, 0x7d, 0x99, 0xc9, 0xb2, 0x63, 0xa3, 0x45, 0xac, 0x53, 0x43, 0x52, 0x48,
	0x9d, 0x2a, 0x8a, 0xb4, 0x41, 0x42, 0xbe, 0x63, 0xb4, 0x08, 0xf7, 0x51, 0x93, 0x44, 0x65, 0x6b,
	0x16, 0x1e, 0x15, 0x72, 0x06, 0x53, 0x49, 0x4f, 0x25, 0xe6, 0x50, 0x34, 0x94, 0x37, 0xd4, 0x20,
	0x96, 0x35, 0xbb, 0xfc, 0xdc, 0x65, 0x9d, 0xf1, 0xca, 0x93, 0xbf, 0xb1, 0x6e, 0x02, 0x25, 0xce,
	0xc0, 0x87, 0xfd, 0x68, 0x39, 0x19, 0x25, 0x88, 0xa5, 0xa7, 0xdc, 0xe4, 0xa0, 0x48, 0xd2, 0x3c,
	0x1c, 0xae, 0x3a, 0x36, 0xa7, 0x13, 0xb0, 0x1a, 0x95, 0x85, 0xe1, 0xe8, 0xeb, 0x23, 0xe8, 0x1f,
	0x04, 0xe1, 0x09, 0xeb, 0x78, 0x5b, 0x70, 0xbf, 0xe8, 0x05, 0x7c, 0xce, 0xf8, 0xad, 0x23, 0x18,
	0xbb, 0xac, 0x90, 0x76, 0x02, 0xf9, 0xbd, 0x49, 0xc5, 0x94, 0x9f, 0xb6, 0xfe, 0x0b, 0x71, 0xf1,
	0x15, 0x15, 0x4f, 0xe0, 0xd5, 0x81, 0xa5, 0xf0, 0xec, 0x37, 0x79, 0xab, 0xe8, 0x19, 0xeb, 0x3e,
	0xe8, 0x3c, 0xa5, 0xb5, 0xdd, 0x07, 0xad, 0xf5, 0x0f, 0xad, 0x49, 0xef, 0xb5, 0x2d, 0xfe, 0x69,
	0xd4, 0x33, 0xfe, 0x33, 0xf7, 0x3d, 0x60, 0x87, 0xb5, 0xf9, 0x9d, 0xcb, 0xac, 0x30, 0xb2, 0xfe,
	0x97, 0xfc, 0x62, 0x23, 0x7a, 0x4d, 0x6d, 0xc2, 0x47, 0x7f, 0x37, 0x6d, 0x73, 0x80, 0xeb, 0xc1,
	0xe3, 0xf9, 0x8b, 0xa4, 0x89, 0xcb, 0x22, 0x81, 0x65, 0xbc, 0xf4, 0x8d, 0xab, 0x87, 0x22, 0xb0,
	0x88, 0x8d, 0x16, 0xd9, 0x5e, 0x55, 0x5d, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xd9, 0xcf,
	0x9a, 0x8d, 0x02, 0x00, 0x00,
}
