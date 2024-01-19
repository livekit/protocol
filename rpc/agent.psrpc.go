// Code generated by protoc-gen-psrpc v0.5.1, DO NOT EDIT.
// source: rpc/agent.proto

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
import livekit1 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_5

// ==============================
// AgentInternal Client Interface
// ==============================

type AgentInternalClient interface {
	CheckEnabled(ctx context.Context, req *CheckEnabledRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*CheckEnabledResponse], error)

	JobRequest(ctx context.Context, topic string, req *livekit1.Job, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	SubscribeAgentRegistered(ctx context.Context, projectId string) (psrpc.Subscription[*CheckEnabledResponse], error)
}

// ==================================
// AgentInternal ServerImpl Interface
// ==================================

type AgentInternalServerImpl interface {
	CheckEnabled(context.Context, *CheckEnabledRequest) (*CheckEnabledResponse, error)

	JobRequest(context.Context, *livekit1.Job) (*google_protobuf.Empty, error)
	JobRequestAffinity(context.Context, *livekit1.Job) float32
}

// ==============================
// AgentInternal Server Interface
// ==============================

type AgentInternalServer interface {
	RegisterJobRequestTopic(topic string) error
	DeregisterJobRequestTopic(topic string)
	PublishAgentRegistered(ctx context.Context, projectId string, msg *CheckEnabledResponse) error

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ====================
// AgentInternal Client
// ====================

type agentInternalClient struct {
	client *client.RPCClient
}

// NewAgentInternalClient creates a psrpc client that implements the AgentInternalClient interface.
func NewAgentInternalClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (AgentInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "AgentInternal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CheckEnabled", false, true, false, false)
	sd.RegisterMethod("JobRequest", true, false, true, false)
	sd.RegisterMethod("AgentRegistered", false, true, false, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &agentInternalClient{
		client: rpcClient,
	}, nil
}

func (c *agentInternalClient) CheckEnabled(ctx context.Context, req *CheckEnabledRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*CheckEnabledResponse], error) {
	return client.RequestMulti[*CheckEnabledResponse](ctx, c.client, "CheckEnabled", nil, req, opts...)
}

func (c *agentInternalClient) JobRequest(ctx context.Context, topic string, req *livekit1.Job, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "JobRequest", []string{topic}, req, opts...)
}

func (c *agentInternalClient) SubscribeAgentRegistered(ctx context.Context, projectId string) (psrpc.Subscription[*CheckEnabledResponse], error) {
	return client.Join[*CheckEnabledResponse](ctx, c.client, "AgentRegistered", []string{projectId})
}

// ====================
// AgentInternal Server
// ====================

type agentInternalServer struct {
	svc AgentInternalServerImpl
	rpc *server.RPCServer
}

// NewAgentInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewAgentInternalServer(svc AgentInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (AgentInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "AgentInternal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CheckEnabled", false, true, false, false)
	var err error
	err = server.RegisterHandler(s, "CheckEnabled", nil, svc.CheckEnabled, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("JobRequest", true, false, true, false)
	sd.RegisterMethod("AgentRegistered", false, true, false, false)
	return &agentInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *agentInternalServer) RegisterJobRequestTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "JobRequest", []string{topic}, s.svc.JobRequest, s.svc.JobRequestAffinity)
}

func (s *agentInternalServer) DeregisterJobRequestTopic(topic string) {
	s.rpc.DeregisterHandler("JobRequest", []string{topic})
}

func (s *agentInternalServer) PublishAgentRegistered(ctx context.Context, projectId string, msg *CheckEnabledResponse) error {
	return s.rpc.Publish(ctx, "AgentRegistered", []string{projectId}, msg)
}

func (s *agentInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *agentInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xb3, 0xbc, 0x6f, 0x08, 0x59, 0x4b, 0xc0, 0x45, 0x09, 0xd6, 0x8b, 0xf6, 0x44, 0xd4,
	0x6c, 0x8d, 0xde, 0x0d, 0x6a, 0x38, 0x48, 0xe2, 0x85, 0xa3, 0x17, 0x64, 0x97, 0xa1, 0xac, 0x94,
	0xce, 0xba, 0xbb, 0x25, 0xf1, 0xe6, 0x95, 0xaf, 0xc3, 0x27, 0x34, 0xfd, 0x63, 0x83, 0x09, 0x7a,
	0x6a, 0xfa, 0xcc, 0x6f, 0x76, 0x7e, 0x79, 0x68, 0xcb, 0x68, 0x19, 0x4e, 0x23, 0x48, 0x1c, 0xd7,
	0x06, 0x1d, 0xb2, 0x7f, 0x46, 0x4b, 0xff, 0x34, 0x42, 0x8c, 0x62, 0x08, 0xf3, 0x48, 0xa4, 0xf3,
	0x10, 0x56, 0xda, 0x7d, 0x14, 0x84, 0xdf, 0x44, 0xed, 0x14, 0x26, 0xb6, 0xfc, 0xed, 0xc4, 0x6a,
	0x0d, 0x4b, 0xe5, 0x26, 0x3b, 0xaf, 0x04, 0xc7, 0xb4, 0xf3, 0xb8, 0x00, 0xb9, 0x1c, 0x26, 0x53,
	0x11, 0xc3, 0x6c, 0x0c, 0xef, 0x29, 0x58, 0x17, 0xcc, 0xe9, 0xd1, 0xcf, 0xd8, 0x6a, 0x4c, 0x2c,
	0xb0, 0x73, 0xea, 0x19, 0xc4, 0xd5, 0x04, 0x8a, 0xbc, 0x47, 0xce, 0x48, 0xbf, 0x31, 0x3e, 0xc8,
	0xb2, 0x12, 0x65, 0x97, 0xf4, 0x50, 0xa7, 0x22, 0x56, 0x76, 0x01, 0xa6, 0xe2, 0x6a, 0x39, 0xd7,
	0xae, 0x06, 0x25, 0x7c, 0xf3, 0x59, 0xa3, 0xcd, 0xfb, 0x4c, 0xe7, 0x29, 0x71, 0x60, 0x92, 0x69,
	0xcc, 0x9e, 0xa9, 0xb7, 0x7b, 0x99, 0xf5, 0xb8, 0xd1, 0x92, 0xef, 0x71, 0xf4, 0x4f, 0xf6, 0x4c,
	0x0a, 0xcd, 0xa0, 0xb1, 0xdd, 0x90, 0xff, 0x83, 0x5a, 0x9f, 0xb0, 0x3b, 0x4a, 0x47, 0x28, 0xca,
	0x15, 0xe6, 0xf1, 0xb2, 0x03, 0x3e, 0x42, 0xe1, 0x77, 0x79, 0xd1, 0x1e, 0xff, 0x6e, 0x8f, 0x0f,
	0xb3, 0xf6, 0x02, 0xba, 0xdd, 0x90, 0x7a, 0x9b, 0x0c, 0xc8, 0x35, 0x61, 0xaf, 0xb4, 0x95, 0xfb,
	0x8d, 0x21, 0x52, 0xd6, 0x81, 0x81, 0x19, 0xfb, 0x65, 0xed, 0x2f, 0x1f, 0x7f, 0xbb, 0x21, 0xdd,
	0x06, 0x69, 0x13, 0xdf, 0x63, 0x54, 0x1b, 0x7c, 0x03, 0xe9, 0x26, 0x6a, 0x96, 0x19, 0x3e, 0x5c,
	0xbd, 0x5c, 0x44, 0xca, 0x2d, 0x52, 0xc1, 0x25, 0xae, 0xc2, 0xd2, 0xaf, 0xfa, 0xea, 0x65, 0x14,
	0x5a, 0x30, 0x6b, 0x25, 0x21, 0x34, 0x5a, 0x8a, 0x7a, 0x7e, 0xf4, 0xf6, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0xd5, 0xcb, 0x7e, 0x13, 0x0f, 0x02, 0x00, 0x00,
}
