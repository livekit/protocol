// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/agent_dispatch.proto

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
import livekit3 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_6

// ======================================
// AgentDispatchInternal Client Interface
// ======================================

type AgentDispatchInternalClient[RoomTopicType ~string] interface {
	CreateDispatch(ctx context.Context, room RoomTopicType, req *livekit3.CreateAgentDispatchRequest, opts ...psrpc.RequestOption) (*livekit3.AgentDispatch, error)

	DeleteDispatch(ctx context.Context, room RoomTopicType, req *livekit3.DeleteAgentDispatchRequest, opts ...psrpc.RequestOption) (*livekit3.AgentDispatch, error)

	ListDispatch(ctx context.Context, room RoomTopicType, req *livekit3.ListAgentDispatchRequest, opts ...psrpc.RequestOption) (*livekit3.ListAgentDispatchResponse, error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ==========================================
// AgentDispatchInternal ServerImpl Interface
// ==========================================

type AgentDispatchInternalServerImpl interface {
	CreateDispatch(context.Context, *livekit3.CreateAgentDispatchRequest) (*livekit3.AgentDispatch, error)

	DeleteDispatch(context.Context, *livekit3.DeleteAgentDispatchRequest) (*livekit3.AgentDispatch, error)

	ListDispatch(context.Context, *livekit3.ListAgentDispatchRequest) (*livekit3.ListAgentDispatchResponse, error)
}

// ======================================
// AgentDispatchInternal Server Interface
// ======================================

type AgentDispatchInternalServer[RoomTopicType ~string] interface {
	RegisterCreateDispatchTopic(room RoomTopicType) error
	DeregisterCreateDispatchTopic(room RoomTopicType)
	RegisterDeleteDispatchTopic(room RoomTopicType) error
	DeregisterDeleteDispatchTopic(room RoomTopicType)
	RegisterListDispatchTopic(room RoomTopicType) error
	DeregisterListDispatchTopic(room RoomTopicType)
	RegisterAllRoomTopics(room RoomTopicType) error
	DeregisterAllRoomTopics(room RoomTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ============================
// AgentDispatchInternal Client
// ============================

type agentDispatchInternalClient[RoomTopicType ~string] struct {
	client *client.RPCClient
}

// NewAgentDispatchInternalClient creates a psrpc client that implements the AgentDispatchInternalClient interface.
func NewAgentDispatchInternalClient[RoomTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (AgentDispatchInternalClient[RoomTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "AgentDispatchInternal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CreateDispatch", false, false, true, true)
	sd.RegisterMethod("DeleteDispatch", false, false, true, true)
	sd.RegisterMethod("ListDispatch", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &agentDispatchInternalClient[RoomTopicType]{
		client: rpcClient,
	}, nil
}

func (c *agentDispatchInternalClient[RoomTopicType]) CreateDispatch(ctx context.Context, room RoomTopicType, req *livekit3.CreateAgentDispatchRequest, opts ...psrpc.RequestOption) (*livekit3.AgentDispatch, error) {
	return client.RequestSingle[*livekit3.AgentDispatch](ctx, c.client, "CreateDispatch", []string{string(room)}, req, opts...)
}

func (c *agentDispatchInternalClient[RoomTopicType]) DeleteDispatch(ctx context.Context, room RoomTopicType, req *livekit3.DeleteAgentDispatchRequest, opts ...psrpc.RequestOption) (*livekit3.AgentDispatch, error) {
	return client.RequestSingle[*livekit3.AgentDispatch](ctx, c.client, "DeleteDispatch", []string{string(room)}, req, opts...)
}

func (c *agentDispatchInternalClient[RoomTopicType]) ListDispatch(ctx context.Context, room RoomTopicType, req *livekit3.ListAgentDispatchRequest, opts ...psrpc.RequestOption) (*livekit3.ListAgentDispatchResponse, error) {
	return client.RequestSingle[*livekit3.ListAgentDispatchResponse](ctx, c.client, "ListDispatch", []string{string(room)}, req, opts...)
}

func (s *agentDispatchInternalClient[RoomTopicType]) Close() {
	s.client.Close()
}

// ============================
// AgentDispatchInternal Server
// ============================

type agentDispatchInternalServer[RoomTopicType ~string] struct {
	svc AgentDispatchInternalServerImpl
	rpc *server.RPCServer
}

// NewAgentDispatchInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewAgentDispatchInternalServer[RoomTopicType ~string](svc AgentDispatchInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (AgentDispatchInternalServer[RoomTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "AgentDispatchInternal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CreateDispatch", false, false, true, true)
	sd.RegisterMethod("DeleteDispatch", false, false, true, true)
	sd.RegisterMethod("ListDispatch", false, false, true, true)
	return &agentDispatchInternalServer[RoomTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *agentDispatchInternalServer[RoomTopicType]) RegisterCreateDispatchTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "CreateDispatch", []string{string(room)}, s.svc.CreateDispatch, nil)
}

func (s *agentDispatchInternalServer[RoomTopicType]) DeregisterCreateDispatchTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("CreateDispatch", []string{string(room)})
}

func (s *agentDispatchInternalServer[RoomTopicType]) RegisterDeleteDispatchTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "DeleteDispatch", []string{string(room)}, s.svc.DeleteDispatch, nil)
}

func (s *agentDispatchInternalServer[RoomTopicType]) DeregisterDeleteDispatchTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("DeleteDispatch", []string{string(room)})
}

func (s *agentDispatchInternalServer[RoomTopicType]) RegisterListDispatchTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "ListDispatch", []string{string(room)}, s.svc.ListDispatch, nil)
}

func (s *agentDispatchInternalServer[RoomTopicType]) DeregisterListDispatchTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("ListDispatch", []string{string(room)})
}

func (s *agentDispatchInternalServer[RoomTopicType]) allRoomTopicRegisterers() server.RegistererSlice {
	return server.RegistererSlice{
		server.NewRegisterer(s.RegisterCreateDispatchTopic, s.DeregisterCreateDispatchTopic),
		server.NewRegisterer(s.RegisterDeleteDispatchTopic, s.DeregisterDeleteDispatchTopic),
		server.NewRegisterer(s.RegisterListDispatchTopic, s.DeregisterListDispatchTopic),
	}
}

func (s *agentDispatchInternalServer[RoomTopicType]) RegisterAllRoomTopics(room RoomTopicType) error {
	return s.allRoomTopicRegisterers().Register(room)
}

func (s *agentDispatchInternalServer[RoomTopicType]) DeregisterAllRoomTopics(room RoomTopicType) {
	s.allRoomTopicRegisterers().Deregister(room)
}

func (s *agentDispatchInternalServer[RoomTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *agentDispatchInternalServer[RoomTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor1 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2a, 0x48, 0xd6,
	0x4f, 0x4c, 0x4f, 0xcd, 0x2b, 0x89, 0x4f, 0xc9, 0x2c, 0x2e, 0x48, 0x2c, 0x49, 0xce, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0xe2, 0xcd, 0x2f, 0x28, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0x88, 0x49, 0xc9, 0xe4, 0x64, 0x96, 0xa5, 0x66, 0x67, 0x96, 0xc4, 0x63, 0xd3,
	0x61, 0x74, 0x85, 0x89, 0x4b, 0xd4, 0x11, 0x24, 0xe1, 0x02, 0x15, 0xf7, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0x11, 0x4a, 0xe5, 0xe2, 0x73, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x85, 0xc9, 0x08,
	0x29, 0xeb, 0x41, 0x8d, 0xd2, 0x83, 0x48, 0xa0, 0xe8, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x91, 0x12, 0x83, 0x2b, 0x42, 0x91, 0x56, 0x12, 0xdb, 0xd4, 0xc9, 0x28, 0x24, 0xc0, 0x28, 0xc5,
	0xc7, 0xc5, 0x52, 0x94, 0x9f, 0x9f, 0x2b, 0x04, 0x26, 0x25, 0x18, 0x41, 0xd6, 0xb8, 0xa4, 0xe6,
	0xa4, 0x62, 0xb5, 0x06, 0x22, 0x41, 0x15, 0x6b, 0x72, 0xb9, 0x78, 0x7c, 0x32, 0x8b, 0xe1, 0xea,
	0x84, 0x14, 0xe1, 0xfa, 0x41, 0xc2, 0x58, 0xad, 0x50, 0xc2, 0xa7, 0xa4, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x15, 0x97, 0x75, 0x4e, 0x8a, 0x51, 0xf2, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x50, 0x73, 0xf4, 0xc1, 0x41, 0x9e, 0x9c, 0x9f, 0xa3, 0x5f, 0x54, 0x90, 0x9c,
	0xc4, 0x06, 0xe6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd2, 0x29, 0xe1, 0xce, 0x01,
	0x00, 0x00,
}
