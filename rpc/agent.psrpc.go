// Code generated by protoc-gen-psrpc v0.5.0, DO NOT EDIT.
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
import livekit2 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_5

// ==============================
// AgentInternal Client Interface
// ==============================

type AgentInternalClient interface {
	CheckEnabled(ctx context.Context, req *CheckEnabledRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*CheckEnabledResponse], error)

	JobRequest(ctx context.Context, namespace string, jobType string, req *livekit2.Job, opts ...psrpc.RequestOption) (*JobRequestResponse, error)

	JobTerminate(ctx context.Context, jobId string, req *JobTerminateRequest, opts ...psrpc.RequestOption) (*JobTerminateResponse, error)

	SubscribeWorkerRegistered(ctx context.Context, handlerNamespace string) (psrpc.Subscription[*google_protobuf.Empty], error)
}

// ==================================
// AgentInternal ServerImpl Interface
// ==================================

type AgentInternalServerImpl interface {
	CheckEnabled(context.Context, *CheckEnabledRequest) (*CheckEnabledResponse, error)

	JobRequest(context.Context, *livekit2.Job) (*JobRequestResponse, error)
	JobRequestAffinity(context.Context, *livekit2.Job) float32

	JobTerminate(context.Context, *JobTerminateRequest) (*JobTerminateResponse, error)
}

// ==============================
// AgentInternal Server Interface
// ==============================

type AgentInternalServer interface {
	RegisterJobRequestTopic(namespace string, jobType string) error
	DeregisterJobRequestTopic(namespace string, jobType string)
	RegisterJobTerminateTopic(jobId string) error
	DeregisterJobTerminateTopic(jobId string)
	PublishWorkerRegistered(ctx context.Context, handlerNamespace string, msg *google_protobuf.Empty) error

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
	sd.RegisterMethod("JobTerminate", false, false, true, true)
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

func (c *agentInternalClient) JobRequest(ctx context.Context, namespace string, jobType string, req *livekit2.Job, opts ...psrpc.RequestOption) (*JobRequestResponse, error) {
	return client.RequestSingle[*JobRequestResponse](ctx, c.client, "JobRequest", []string{namespace, jobType}, req, opts...)
}

func (c *agentInternalClient) JobTerminate(ctx context.Context, jobId string, req *JobTerminateRequest, opts ...psrpc.RequestOption) (*JobTerminateResponse, error) {
	return client.RequestSingle[*JobTerminateResponse](ctx, c.client, "JobTerminate", []string{jobId}, req, opts...)
}

func (c *agentInternalClient) SubscribeWorkerRegistered(ctx context.Context, handlerNamespace string) (psrpc.Subscription[*google_protobuf.Empty], error) {
	return client.Join[*google_protobuf.Empty](ctx, c.client, "WorkerRegistered", []string{handlerNamespace})
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
	sd.RegisterMethod("JobTerminate", false, false, true, true)
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

func (s *agentInternalServer) RegisterJobTerminateTopic(jobId string) error {
	return server.RegisterHandler(s.rpc, "JobTerminate", []string{jobId}, s.svc.JobTerminate, nil)
}

func (s *agentInternalServer) DeregisterJobTerminateTopic(jobId string) {
	s.rpc.DeregisterHandler("JobTerminate", []string{jobId})
}

func (s *agentInternalServer) PublishWorkerRegistered(ctx context.Context, handlerNamespace string, msg *google_protobuf.Empty) error {
	return s.rpc.Publish(ctx, "WorkerRegistered", []string{handlerNamespace}, msg)
}

func (s *agentInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *agentInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor0 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x6f, 0xd2, 0x50,
	0x14, 0xf7, 0xc2, 0x46, 0xe0, 0xc0, 0x5c, 0xb9, 0x80, 0xd6, 0x1a, 0xb3, 0xd1, 0x17, 0x89, 0x26,
	0xad, 0xc1, 0x67, 0xe3, 0xd8, 0xac, 0x06, 0xe2, 0x20, 0x76, 0x9d, 0x26, 0x26, 0xa6, 0x69, 0xcb,
	0x11, 0x3a, 0x4a, 0x6f, 0xbd, 0xbd, 0x98, 0xec, 0x23, 0xf0, 0x59, 0x7c, 0xe3, 0xb3, 0xf9, 0x01,
	0x4c, 0xff, 0xac, 0xc2, 0x82, 0x0f, 0x3e, 0xde, 0xdf, 0xf9, 0xe5, 0xe4, 0xf7, 0xe7, 0x1e, 0x38,
	0xe6, 0x91, 0xa7, 0x3b, 0x33, 0x0c, 0x85, 0x16, 0x71, 0x26, 0x18, 0x2d, 0xf3, 0xc8, 0x53, 0x9e,
	0xce, 0x18, 0x9b, 0x05, 0xa8, 0xa7, 0x90, 0xbb, 0xfa, 0xae, 0xe3, 0x32, 0x12, 0xb7, 0x19, 0x43,
	0x39, 0x62, 0x91, 0xf0, 0x59, 0x18, 0xe7, 0xcf, 0x56, 0xe0, 0xff, 0xc4, 0x85, 0x2f, 0xec, 0xad,
	0x2d, 0x6a, 0x07, 0x5a, 0x17, 0x73, 0xf4, 0x16, 0x46, 0xe8, 0xb8, 0x01, 0x4e, 0x4d, 0xfc, 0xb1,
	0xc2, 0x58, 0xa8, 0xbf, 0x08, 0xb4, 0x77, 0xf1, 0x38, 0x62, 0x61, 0x8c, 0xb4, 0x0b, 0x0d, 0xce,
	0xd8, 0xd2, 0xc6, 0x0c, 0x97, 0xc9, 0x29, 0xe9, 0x55, 0xcd, 0x7a, 0x82, 0xe5, 0x54, 0xfa, 0x12,
	0x9a, 0xd1, 0xca, 0x0d, 0xfc, 0x78, 0x8e, 0xbc, 0xe0, 0x95, 0x52, 0x9e, 0x54, 0x0c, 0xee, 0xc8,
	0x2a, 0x40, 0xe8, 0x2c, 0x31, 0x8e, 0x1c, 0x0f, 0x63, 0xb9, 0x7c, 0x5a, 0xee, 0xd5, 0xce, 0x4b,
	0x32, 0x31, 0xb7, 0x50, 0x7a, 0x02, 0xf5, 0x54, 0xb2, 0x9d, 0x62, 0xf2, 0x41, 0x42, 0x32, 0x21,
	0x85, 0xc6, 0x09, 0xa2, 0xbe, 0x01, 0x3a, 0x62, 0x6e, 0xae, 0xbd, 0x90, 0xfa, 0x1c, 0x0e, 0x63,
	0xe1, 0x08, 0x4c, 0x35, 0xd6, 0xfb, 0x4d, 0x2d, 0xf7, 0xaf, 0x8d, 0x98, 0x7b, 0x95, 0x0c, 0xcc,
	0x6c, 0xae, 0x7e, 0x83, 0xd6, 0x88, 0xb9, 0x16, 0xf2, 0xa5, 0x1f, 0x26, 0x70, 0xb6, 0x87, 0x76,
	0xa0, 0x72, 0xc3, 0x5c, 0xdb, 0xcf, 0x4c, 0xd6, 0xcc, 0xc3, 0x1b, 0xe6, 0x0e, 0xa7, 0x54, 0x87,
	0x0a, 0x47, 0x27, 0x66, 0x61, 0xea, 0xe9, 0x61, 0xff, 0xb1, 0xc6, 0x23, 0x4f, 0xdb, 0x5d, 0x90,
	0x8c, 0xcd, 0x9c, 0xa6, 0xbe, 0x85, 0xf6, 0xee, 0xf4, 0x3f, 0xf5, 0xbd, 0xb8, 0x48, 0xed, 0xdd,
	0x5b, 0x4f, 0x65, 0x68, 0x5b, 0x86, 0x39, 0x1c, 0x0f, 0xac, 0xe1, 0x64, 0x6c, 0x9b, 0xc6, 0xa7,
	0x6b, 0xe3, 0xca, 0x32, 0xde, 0x49, 0x0f, 0x68, 0x0b, 0x8e, 0x07, 0x1f, 0x8c, 0xb1, 0x65, 0x7f,
	0x34, 0xde, 0x5b, 0xb6, 0x39, 0x99, 0x5c, 0x4a, 0xa4, 0xff, 0xbb, 0x04, 0x47, 0x83, 0x24, 0xb2,
	0x61, 0x28, 0x90, 0x87, 0x4e, 0x40, 0x2f, 0xa1, 0xb1, 0x5d, 0x31, 0x95, 0x53, 0x23, 0x7b, 0x7e,
	0x83, 0xf2, 0x64, 0xcf, 0x24, 0x33, 0xa1, 0x56, 0x37, 0x6b, 0x72, 0x70, 0x56, 0xea, 0x11, 0xfa,
	0x19, 0xe0, 0x6f, 0x09, 0xb4, 0xb1, 0xed, 0x46, 0x29, 0x32, 0xba, 0xd7, 0x91, 0xda, 0xdd, 0xac,
	0xc9, 0x33, 0x89, 0x28, 0x1d, 0x5a, 0x2b, 0x0a, 0xa7, 0xd5, 0x24, 0x76, 0x71, 0x1b, 0xe1, 0x19,
	0x79, 0x45, 0xe8, 0x35, 0x34, 0xb6, 0xdd, 0xe7, 0x32, 0xf7, 0x14, 0x96, 0xcb, 0xdc, 0x97, 0xb5,
	0x2a, 0x6d, 0xd6, 0xa4, 0x21, 0x11, 0xa5, 0x4a, 0xf3, 0x4e, 0x29, 0x82, 0xf4, 0x85, 0xf1, 0x05,
	0x72, 0x13, 0x67, 0x7e, 0x2c, 0x90, 0xe3, 0x94, 0x3e, 0xd2, 0xb2, 0x73, 0xd2, 0xee, 0xce, 0x49,
	0x33, 0x92, 0x73, 0x52, 0xfe, 0x81, 0x67, 0xea, 0xab, 0x44, 0x22, 0x4a, 0x8b, 0x36, 0xe7, 0x4e,
	0x38, 0x0d, 0x90, 0xdb, 0x85, 0x8f, 0x24, 0x95, 0xf3, 0xee, 0xd7, 0x93, 0x99, 0x2f, 0xe6, 0x2b,
	0x57, 0xf3, 0xd8, 0x52, 0xcf, 0x33, 0xc9, 0xce, 0xd5, 0x63, 0x81, 0xce, 0x23, 0xcf, 0xad, 0xa4,
	0xaf, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xb5, 0x90, 0x87, 0xe2, 0x03, 0x00, 0x00,
}
