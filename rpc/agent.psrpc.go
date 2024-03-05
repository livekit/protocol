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

	JobRequest(ctx context.Context, namespace string, jobType string, req *livekit1.Job, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	SubscribeWorkerRegistered(ctx context.Context, projectId string) (psrpc.Subscription[*google_protobuf.Empty], error)
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
	RegisterJobRequestTopic(namespace string, jobType string) error
	DeregisterJobRequestTopic(namespace string, jobType string)
	PublishWorkerRegistered(ctx context.Context, projectId string, msg *google_protobuf.Empty) error

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
	sd.RegisterMethod("WorkerRegistered", false, true, false, false)

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

func (c *agentInternalClient) JobRequest(ctx context.Context, namespace string, jobType string, req *livekit1.Job, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "JobRequest", []string{namespace, jobType}, req, opts...)
}

func (c *agentInternalClient) SubscribeWorkerRegistered(ctx context.Context, projectId string) (psrpc.Subscription[*google_protobuf.Empty], error) {
	return client.Join[*google_protobuf.Empty](ctx, c.client, "WorkerRegistered", []string{projectId})
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
	sd.RegisterMethod("WorkerRegistered", false, true, false, false)
	return &agentInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *agentInternalServer) RegisterJobRequestTopic(namespace string, jobType string) error {
	return server.RegisterHandler(s.rpc, "JobRequest", []string{namespace, jobType}, s.svc.JobRequest, s.svc.JobRequestAffinity)
}

func (s *agentInternalServer) DeregisterJobRequestTopic(namespace string, jobType string) {
	s.rpc.DeregisterHandler("JobRequest", []string{namespace, jobType})
}

func (s *agentInternalServer) PublishWorkerRegistered(ctx context.Context, projectId string, msg *google_protobuf.Empty) error {
	return s.rpc.Publish(ctx, "WorkerRegistered", []string{projectId}, msg)
}

func (s *agentInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *agentInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xdf, 0xea, 0xd3, 0x30,
	0x1c, 0xc5, 0xc9, 0x26, 0xd2, 0xc5, 0x0e, 0x6b, 0xe6, 0xc6, 0xac, 0x28, 0x5b, 0xaf, 0x86, 0x4a,
	0x2a, 0xfa, 0x02, 0x53, 0xd9, 0x85, 0x03, 0x6f, 0x0a, 0x22, 0x78, 0x61, 0x69, 0xb2, 0xaf, 0x5d,
	0xd6, 0x3f, 0x89, 0x49, 0x3a, 0xd8, 0x0b, 0x08, 0x7b, 0x00, 0x5f, 0x64, 0x4f, 0x28, 0x5d, 0xbb,
	0xb2, 0x1f, 0x6c, 0x57, 0x21, 0xe7, 0x7b, 0x72, 0xf2, 0xf9, 0x1e, 0xfc, 0x54, 0x2b, 0x1e, 0x26,
	0x29, 0x94, 0x96, 0x2a, 0x2d, 0xad, 0x24, 0x7d, 0xad, 0xb8, 0xff, 0x32, 0x95, 0x32, 0xcd, 0x21,
	0x3c, 0x4b, 0xac, 0xfa, 0x1d, 0x42, 0xa1, 0xec, 0xa1, 0x71, 0xf8, 0x43, 0xa9, 0xac, 0x90, 0xa5,
	0x69, 0xaf, 0xa3, 0x5c, 0xec, 0x21, 0x13, 0x36, 0xbe, 0x4a, 0x09, 0xc6, 0x78, 0xf4, 0x65, 0x0b,
	0x3c, 0x5b, 0x95, 0x09, 0xcb, 0x61, 0x13, 0xc1, 0x9f, 0x0a, 0x8c, 0x0d, 0xfe, 0x22, 0xfc, 0xfc,
	0xa1, 0x6e, 0x94, 0x2c, 0x0d, 0x90, 0x39, 0x76, 0xb5, 0x94, 0x45, 0x0c, 0x8d, 0x3e, 0x45, 0x33,
	0xb4, 0x70, 0xa2, 0x27, 0xb5, 0xd6, 0x5a, 0xc9, 0x5b, 0xfc, 0x4c, 0x55, 0x2c, 0x17, 0x66, 0x0b,
	0xba, 0xf3, 0xf5, 0xce, 0x3e, 0xaf, 0x1b, 0x5c, 0xcc, 0xaf, 0x31, 0x2e, 0x93, 0x02, 0x8c, 0x4a,
	0x38, 0x98, 0x69, 0x7f, 0xd6, 0x5f, 0x0c, 0xa2, 0x2b, 0xe5, 0xc3, 0xbf, 0x1e, 0x1e, 0x7e, 0xaa,
	0x79, 0xbf, 0x96, 0x16, 0x74, 0x99, 0xe4, 0xe4, 0x1b, 0x76, 0xaf, 0xc9, 0xc8, 0x94, 0x6a, 0xc5,
	0xe9, 0x8d, 0x25, 0xfc, 0x17, 0x37, 0x26, 0xcd, 0x1a, 0x81, 0x73, 0x3a, 0xa2, 0x47, 0xcb, 0xde,
	0x02, 0x91, 0xef, 0x18, 0xaf, 0x25, 0x6b, 0x9f, 0x10, 0x97, 0xb6, 0x25, 0xd1, 0xb5, 0x64, 0xfe,
	0x84, 0x36, 0xf5, 0xd2, 0x4b, 0xbd, 0x74, 0x55, 0xd7, 0x1b, 0xcc, 0x4f, 0x47, 0xf4, 0xca, 0x43,
	0xfe, 0x98, 0x0c, 0x3a, 0x54, 0xe2, 0xec, 0x24, 0x8b, 0xed, 0x41, 0xc1, 0x12, 0xbd, 0x47, 0xe4,
	0x17, 0xf6, 0x7e, 0x48, 0x9d, 0x81, 0x8e, 0x20, 0x15, 0xc6, 0x82, 0x86, 0x0d, 0xb9, 0x13, 0x77,
	0xf7, 0x1b, 0xff, 0x74, 0x44, 0x13, 0x07, 0x79, 0xc8, 0x77, 0x09, 0x56, 0x5a, 0xee, 0x80, 0xdb,
	0x58, 0x6c, 0x6a, 0xec, 0xcf, 0xef, 0x7e, 0xbe, 0x49, 0x85, 0xdd, 0x56, 0x8c, 0x72, 0x59, 0x84,
	0x2d, 0x74, 0x77, 0xaa, 0x2c, 0x0d, 0x0d, 0xe8, 0xbd, 0xe0, 0x10, 0x6a, 0xc5, 0xd9, 0xe3, 0x73,
	0xf2, 0xc7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x9b, 0xd1, 0x3e, 0x45, 0x02, 0x00, 0x00,
}
