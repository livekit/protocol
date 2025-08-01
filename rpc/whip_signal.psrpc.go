// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/whip_signal.proto

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
import google_protobuf "google.golang.org/protobuf/types/known/emptypb"

var _ = version.PsrpcVersion_0_6

// =====================
// WHIP Client Interface
// =====================

type WHIPClient[TopicTopicType ~string] interface {
	Create(ctx context.Context, topic TopicTopicType, req *WHIPCreateRequest, opts ...psrpc.RequestOption) (*WHIPCreateResponse, error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// =========================
// WHIP ServerImpl Interface
// =========================

type WHIPServerImpl interface {
	Create(context.Context, *WHIPCreateRequest) (*WHIPCreateResponse, error)
}

// =====================
// WHIP Server Interface
// =====================

type WHIPServer[TopicTopicType ~string] interface {
	RegisterCreateTopic(topic TopicTopicType) error
	DeregisterCreateTopic(topic TopicTopicType)
	RegisterAllCommonTopics(topic TopicTopicType) error
	DeregisterAllCommonTopics(topic TopicTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ===========
// WHIP Client
// ===========

type wHIPClient[TopicTopicType ~string] struct {
	client *client.RPCClient
}

// NewWHIPClient creates a psrpc client that implements the WHIPClient interface.
func NewWHIPClient[TopicTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (WHIPClient[TopicTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "WHIP",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("Create", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &wHIPClient[TopicTopicType]{
		client: rpcClient,
	}, nil
}

func (c *wHIPClient[TopicTopicType]) Create(ctx context.Context, topic TopicTopicType, req *WHIPCreateRequest, opts ...psrpc.RequestOption) (*WHIPCreateResponse, error) {
	return client.RequestSingle[*WHIPCreateResponse](ctx, c.client, "Create", []string{string(topic)}, req, opts...)
}

func (s *wHIPClient[TopicTopicType]) Close() {
	s.client.Close()
}

// ===========
// WHIP Server
// ===========

type wHIPServer[TopicTopicType ~string] struct {
	svc WHIPServerImpl
	rpc *server.RPCServer
}

// NewWHIPServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewWHIPServer[TopicTopicType ~string](svc WHIPServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (WHIPServer[TopicTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "WHIP",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("Create", false, false, true, true)
	return &wHIPServer[TopicTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *wHIPServer[TopicTopicType]) RegisterCreateTopic(topic TopicTopicType) error {
	return server.RegisterHandler(s.rpc, "Create", []string{string(topic)}, s.svc.Create, nil)
}

func (s *wHIPServer[TopicTopicType]) DeregisterCreateTopic(topic TopicTopicType) {
	s.rpc.DeregisterHandler("Create", []string{string(topic)})
}

func (s *wHIPServer[TopicTopicType]) allCommonTopicRegisterers() server.RegistererSlice {
	return server.RegistererSlice{
		server.NewRegisterer(s.RegisterCreateTopic, s.DeregisterCreateTopic),
	}
}

func (s *wHIPServer[TopicTopicType]) RegisterAllCommonTopics(topic TopicTopicType) error {
	return s.allCommonTopicRegisterers().Register(topic)
}

func (s *wHIPServer[TopicTopicType]) DeregisterAllCommonTopics(topic TopicTopicType) {
	s.allCommonTopicRegisterers().Deregister(topic)
}

func (s *wHIPServer[TopicTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *wHIPServer[TopicTopicType]) Kill() {
	s.rpc.Close(true)
}

// ================================
// WHIPParticipant Client Interface
// ================================

type WHIPParticipantClient[TopicTopicType ~string] interface {
	ICETrickle(ctx context.Context, topic TopicTopicType, req *WHIPParticipantICETrickleRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	ICERestart(ctx context.Context, topic TopicTopicType, req *WHIPParticipantICERestartRequest, opts ...psrpc.RequestOption) (*WHIPParticipantICERestartResponse, error)

	DeleteSession(ctx context.Context, topic TopicTopicType, req *WHIPParticipantDeleteSessionRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ====================================
// WHIPParticipant ServerImpl Interface
// ====================================

type WHIPParticipantServerImpl interface {
	ICETrickle(context.Context, *WHIPParticipantICETrickleRequest) (*google_protobuf.Empty, error)

	ICERestart(context.Context, *WHIPParticipantICERestartRequest) (*WHIPParticipantICERestartResponse, error)

	DeleteSession(context.Context, *WHIPParticipantDeleteSessionRequest) (*google_protobuf.Empty, error)
}

// ================================
// WHIPParticipant Server Interface
// ================================

type WHIPParticipantServer[TopicTopicType ~string] interface {
	RegisterICETrickleTopic(topic TopicTopicType) error
	DeregisterICETrickleTopic(topic TopicTopicType)
	RegisterICERestartTopic(topic TopicTopicType) error
	DeregisterICERestartTopic(topic TopicTopicType)
	RegisterDeleteSessionTopic(topic TopicTopicType) error
	DeregisterDeleteSessionTopic(topic TopicTopicType)
	RegisterAllCommonTopics(topic TopicTopicType) error
	DeregisterAllCommonTopics(topic TopicTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ======================
// WHIPParticipant Client
// ======================

type wHIPParticipantClient[TopicTopicType ~string] struct {
	client *client.RPCClient
}

// NewWHIPParticipantClient creates a psrpc client that implements the WHIPParticipantClient interface.
func NewWHIPParticipantClient[TopicTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (WHIPParticipantClient[TopicTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "WHIPParticipant",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("ICETrickle", false, false, true, true)
	sd.RegisterMethod("ICERestart", false, false, true, true)
	sd.RegisterMethod("DeleteSession", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &wHIPParticipantClient[TopicTopicType]{
		client: rpcClient,
	}, nil
}

func (c *wHIPParticipantClient[TopicTopicType]) ICETrickle(ctx context.Context, topic TopicTopicType, req *WHIPParticipantICETrickleRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "ICETrickle", []string{string(topic)}, req, opts...)
}

func (c *wHIPParticipantClient[TopicTopicType]) ICERestart(ctx context.Context, topic TopicTopicType, req *WHIPParticipantICERestartRequest, opts ...psrpc.RequestOption) (*WHIPParticipantICERestartResponse, error) {
	return client.RequestSingle[*WHIPParticipantICERestartResponse](ctx, c.client, "ICERestart", []string{string(topic)}, req, opts...)
}

func (c *wHIPParticipantClient[TopicTopicType]) DeleteSession(ctx context.Context, topic TopicTopicType, req *WHIPParticipantDeleteSessionRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "DeleteSession", []string{string(topic)}, req, opts...)
}

func (s *wHIPParticipantClient[TopicTopicType]) Close() {
	s.client.Close()
}

// ======================
// WHIPParticipant Server
// ======================

type wHIPParticipantServer[TopicTopicType ~string] struct {
	svc WHIPParticipantServerImpl
	rpc *server.RPCServer
}

// NewWHIPParticipantServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewWHIPParticipantServer[TopicTopicType ~string](svc WHIPParticipantServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (WHIPParticipantServer[TopicTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "WHIPParticipant",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("ICETrickle", false, false, true, true)
	sd.RegisterMethod("ICERestart", false, false, true, true)
	sd.RegisterMethod("DeleteSession", false, false, true, true)
	return &wHIPParticipantServer[TopicTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *wHIPParticipantServer[TopicTopicType]) RegisterICETrickleTopic(topic TopicTopicType) error {
	return server.RegisterHandler(s.rpc, "ICETrickle", []string{string(topic)}, s.svc.ICETrickle, nil)
}

func (s *wHIPParticipantServer[TopicTopicType]) DeregisterICETrickleTopic(topic TopicTopicType) {
	s.rpc.DeregisterHandler("ICETrickle", []string{string(topic)})
}

func (s *wHIPParticipantServer[TopicTopicType]) RegisterICERestartTopic(topic TopicTopicType) error {
	return server.RegisterHandler(s.rpc, "ICERestart", []string{string(topic)}, s.svc.ICERestart, nil)
}

func (s *wHIPParticipantServer[TopicTopicType]) DeregisterICERestartTopic(topic TopicTopicType) {
	s.rpc.DeregisterHandler("ICERestart", []string{string(topic)})
}

func (s *wHIPParticipantServer[TopicTopicType]) RegisterDeleteSessionTopic(topic TopicTopicType) error {
	return server.RegisterHandler(s.rpc, "DeleteSession", []string{string(topic)}, s.svc.DeleteSession, nil)
}

func (s *wHIPParticipantServer[TopicTopicType]) DeregisterDeleteSessionTopic(topic TopicTopicType) {
	s.rpc.DeregisterHandler("DeleteSession", []string{string(topic)})
}

func (s *wHIPParticipantServer[TopicTopicType]) allCommonTopicRegisterers() server.RegistererSlice {
	return server.RegistererSlice{
		server.NewRegisterer(s.RegisterICETrickleTopic, s.DeregisterICETrickleTopic),
		server.NewRegisterer(s.RegisterICERestartTopic, s.DeregisterICERestartTopic),
		server.NewRegisterer(s.RegisterDeleteSessionTopic, s.DeregisterDeleteSessionTopic),
	}
}

func (s *wHIPParticipantServer[TopicTopicType]) RegisterAllCommonTopics(topic TopicTopicType) error {
	return s.allCommonTopicRegisterers().Register(topic)
}

func (s *wHIPParticipantServer[TopicTopicType]) DeregisterAllCommonTopics(topic TopicTopicType) {
	s.allCommonTopicRegisterers().Deregister(topic)
}

func (s *wHIPParticipantServer[TopicTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *wHIPParticipantServer[TopicTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor11 = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x6e, 0x1a, 0x3b,
	0x14, 0x96, 0x81, 0x44, 0x97, 0x43, 0xc8, 0x4d, 0x7c, 0x6f, 0x12, 0xee, 0xa0, 0x28, 0x84, 0xdb,
	0x54, 0x2c, 0xaa, 0x41, 0x25, 0x52, 0x1b, 0x65, 0xd9, 0x94, 0xaa, 0x48, 0x55, 0x15, 0x0d, 0x51,
	0x2b, 0x75, 0xd1, 0xd1, 0xe0, 0x31, 0xc4, 0x62, 0x66, 0xec, 0xda, 0x86, 0x08, 0xf5, 0x09, 0xb2,
	0xee, 0x9b, 0xa4, 0xea, 0xc3, 0x74, 0xd5, 0x57, 0xa9, 0xc6, 0x33, 0x10, 0x08, 0x93, 0x1f, 0x75,
	0x93, 0x9d, 0xe7, 0xf3, 0xf1, 0x39, 0xdf, 0xf7, 0x9d, 0x63, 0x0f, 0x6c, 0x49, 0x41, 0x9a, 0x17,
	0xe7, 0x4c, 0xb8, 0x8a, 0x0d, 0x22, 0x2f, 0xb0, 0x85, 0xe4, 0x9a, 0xe3, 0xbc, 0x14, 0xc4, 0xaa,
	0x0e, 0x38, 0x1f, 0x04, 0xb4, 0x69, 0xa0, 0xde, 0xa8, 0xdf, 0xa4, 0xa1, 0xd0, 0x93, 0x24, 0xc2,
	0x2a, 0x73, 0xa1, 0x19, 0x8f, 0x54, 0xfa, 0xb9, 0x1d, 0xb0, 0x31, 0x1d, 0x32, 0xed, 0xb2, 0x48,
	0x53, 0x39, 0x4b, 0x64, 0x6d, 0x4e, 0x71, 0xa9, 0x49, 0x02, 0xd5, 0x2f, 0xf3, 0xb0, 0xf9, 0xf1,
	0x6d, 0xe7, 0xf4, 0x44, 0x52, 0x4f, 0x53, 0x87, 0x7e, 0x19, 0x51, 0xa5, 0xf1, 0x31, 0x94, 0x95,
	0xf6, 0xa4, 0x76, 0x15, 0x55, 0x8a, 0xf1, 0xa8, 0x82, 0x6a, 0xa8, 0x51, 0x6a, 0x6d, 0xd9, 0x69,
	0x02, 0xbb, 0x1b, 0xef, 0x76, 0x93, 0x4d, 0x67, 0x4d, 0xcd, 0x7d, 0xe1, 0x2a, 0x14, 0x79, 0xbf,
	0x4f, 0xa5, 0xab, 0x7c, 0x51, 0xc9, 0xd5, 0x50, 0xa3, 0xe8, 0xfc, 0x65, 0x80, 0xae, 0x2f, 0xf0,
	0x57, 0xd8, 0x55, 0xa3, 0x9e, 0x22, 0x92, 0xf5, 0xa8, 0xef, 0x0a, 0x4f, 0x6a, 0x46, 0x98, 0xf0,
	0x22, 0xed, 0x6a, 0xe9, 0x91, 0xa1, 0xaa, 0xe4, 0x6b, 0xf9, 0x46, 0xa9, 0xf5, 0xd2, 0x96, 0x82,
	0xd8, 0x4b, 0xbc, 0xec, 0xee, 0xec, 0xec, 0xe9, 0xf5, 0xd1, 0x33, 0x73, 0xb2, 0x1d, 0x69, 0x39,
	0x71, 0xaa, 0xea, 0xf6, 0x08, 0xeb, 0x19, 0x14, 0xcd, 0xea, 0x1d, 0x53, 0x1a, 0xef, 0x41, 0xc9,
	0x94, 0x74, 0x23, 0x2f, 0xa4, 0xaa, 0x82, 0x6a, 0xf9, 0x46, 0xd1, 0x01, 0x03, 0xbd, 0x8f, 0x11,
	0x4b, 0x40, 0xed, 0xbe, 0x72, 0x78, 0x03, 0xf2, 0x43, 0x3a, 0x31, 0xee, 0x14, 0x9d, 0x78, 0x89,
	0x5f, 0xc0, 0xca, 0xd8, 0x0b, 0x46, 0xd4, 0x28, 0x2f, 0xb5, 0x6a, 0xb7, 0x08, 0x99, 0xf1, 0x70,
	0x92, 0xf0, 0xe3, 0xdc, 0x11, 0xaa, 0xff, 0x40, 0x80, 0xe7, 0x43, 0x95, 0xe0, 0x91, 0xa2, 0x78,
	0x17, 0xc0, 0x8b, 0xd4, 0x45, 0xea, 0x68, 0x52, 0xab, 0x98, 0x20, 0xb1, 0xa5, 0x07, 0xb0, 0x3e,
	0xef, 0x23, 0xf3, 0x53, 0xd3, 0xcb, 0x73, 0x68, 0xc7, 0xc7, 0x87, 0x50, 0x62, 0x84, 0xba, 0x8a,
	0xca, 0x31, 0x95, 0x53, 0x9f, 0xf1, 0xac, 0xa1, 0x9d, 0x93, 0x76, 0xd7, 0x6c, 0x39, 0xc0, 0x08,
	0x4d, 0x96, 0x0a, 0x3f, 0x81, 0xf5, 0xe4, 0x90, 0x69, 0x6d, 0x9c, 0xbb, 0x60, 0x72, 0xaf, 0x99,
	0x18, 0x03, 0x76, 0xfc, 0xfa, 0x4f, 0x04, 0xb5, 0x98, 0xf7, 0x9c, 0x49, 0x9d, 0x93, 0xf6, 0x99,
	0x64, 0x64, 0x18, 0xcc, 0x46, 0x0a, 0x43, 0x41, 0x72, 0x1e, 0xa6, 0xfc, 0xcd, 0x1a, 0x3f, 0x87,
	0x7f, 0x17, 0xa9, 0xd3, 0x48, 0x33, 0x3d, 0x49, 0x05, 0xfc, 0xb3, 0x20, 0x20, 0xd9, 0xca, 0x50,
	0x9b, 0xcf, 0x52, 0xfb, 0x20, 0xe2, 0x78, 0x1f, 0xd6, 0x94, 0x2f, 0xdc, 0xbe, 0xf4, 0x06, 0x21,
	0x8d, 0x74, 0x65, 0xc5, 0xc4, 0x94, 0x94, 0x2f, 0xde, 0xa4, 0x50, 0xfd, 0x7b, 0xa6, 0x36, 0x87,
	0x9a, 0x91, 0x7f, 0x1c, 0x6d, 0x37, 0x59, 0x17, 0x96, 0x59, 0x07, 0xb0, 0x7f, 0x07, 0xe9, 0x74,
	0xae, 0x96, 0x3d, 0x42, 0x0f, 0xf0, 0x28, 0xb7, 0x5c, 0xed, 0x1b, 0x82, 0xff, 0x6f, 0x94, 0x7b,
	0x4d, 0x03, 0xaa, 0xa7, 0x59, 0x1e, 0xc5, 0xa6, 0xd6, 0x67, 0x28, 0xc4, 0xa4, 0xf0, 0x07, 0x58,
	0x4d, 0x2e, 0x14, 0xde, 0xce, 0xbe, 0x8c, 0xd6, 0xce, 0x12, 0x9e, 0x38, 0x54, 0xaf, 0x5e, 0x5d,
	0xa2, 0x9d, 0x0d, 0x64, 0x6d, 0xc2, 0x2a, 0xe1, 0x61, 0xc8, 0x23, 0xbc, 0xa2, 0xb9, 0x60, 0xa4,
	0x82, 0x8e, 0x50, 0xeb, 0x57, 0x0e, 0xfe, 0xbe, 0xa1, 0x1a, 0xf7, 0x01, 0xae, 0x27, 0x1f, 0x1f,
	0xcc, 0xf2, 0xde, 0x75, 0x33, 0xac, 0x6d, 0x3b, 0x79, 0xda, 0xed, 0xe9, 0xd3, 0x6e, 0xb7, 0xe3,
	0xa7, 0xbd, 0xfe, 0xdf, 0xd5, 0x25, 0xda, 0xca, 0xac, 0x8e, 0xc7, 0xa6, 0x4e, 0xda, 0xd0, 0x5b,
	0xeb, 0x2c, 0x4e, 0xa9, 0xf5, 0xf4, 0xbe, 0xb0, 0x54, 0xf5, 0x1d, 0x75, 0x03, 0x28, 0x2f, 0x74,
	0x16, 0x37, 0xb2, 0x72, 0x66, 0x35, 0xff, 0x0f, 0x54, 0xbe, 0xda, 0xff, 0xb4, 0x37, 0x60, 0xfa,
	0x7c, 0xd4, 0xb3, 0x09, 0x0f, 0x9b, 0xe9, 0x4b, 0x95, 0xfc, 0x00, 0x09, 0x0f, 0x9a, 0x52, 0x90,
	0xde, 0xaa, 0xf9, 0x3a, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xf3, 0x21, 0x4d, 0x3a, 0x07,
	0x00, 0x00,
}
