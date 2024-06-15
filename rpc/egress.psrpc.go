// Code generated by protoc-gen-psrpc v0.5.0, DO NOT EDIT.
// source: rpc/egress.proto

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
import livekit2 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_5

// ===============================
// EgressInternal Client Interface
// ===============================

type EgressInternalClient interface {
	StartEgress(ctx context.Context, topic string, req *StartEgressRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error)

	ListActiveEgress(ctx context.Context, topic string, req *ListActiveEgressRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*ListActiveEgressResponse], error)
}

// ===================================
// EgressInternal ServerImpl Interface
// ===================================

type EgressInternalServerImpl interface {
	StartEgress(context.Context, *StartEgressRequest) (*livekit2.EgressInfo, error)
	StartEgressAffinity(context.Context, *StartEgressRequest) float32

	ListActiveEgress(context.Context, *ListActiveEgressRequest) (*ListActiveEgressResponse, error)
}

// ===============================
// EgressInternal Server Interface
// ===============================

type EgressInternalServer interface {
	RegisterStartEgressTopic(topic string) error
	DeregisterStartEgressTopic(topic string)
	RegisterListActiveEgressTopic(topic string) error
	DeregisterListActiveEgressTopic(topic string)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =====================
// EgressInternal Client
// =====================

type egressInternalClient struct {
	client *client.RPCClient
}

// NewEgressInternalClient creates a psrpc client that implements the EgressInternalClient interface.
func NewEgressInternalClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (EgressInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressInternal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("StartEgress", true, false, true, false)
	sd.RegisterMethod("ListActiveEgress", false, true, false, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &egressInternalClient{
		client: rpcClient,
	}, nil
}

func (c *egressInternalClient) StartEgress(ctx context.Context, topic string, req *StartEgressRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error) {
	return client.RequestSingle[*livekit2.EgressInfo](ctx, c.client, "StartEgress", []string{topic}, req, opts...)
}

func (c *egressInternalClient) ListActiveEgress(ctx context.Context, topic string, req *ListActiveEgressRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*ListActiveEgressResponse], error) {
	return client.RequestMulti[*ListActiveEgressResponse](ctx, c.client, "ListActiveEgress", []string{topic}, req, opts...)
}

// =====================
// EgressInternal Server
// =====================

type egressInternalServer struct {
	svc EgressInternalServerImpl
	rpc *server.RPCServer
}

// NewEgressInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewEgressInternalServer(svc EgressInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (EgressInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressInternal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("StartEgress", true, false, true, false)
	sd.RegisterMethod("ListActiveEgress", false, true, false, false)
	return &egressInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *egressInternalServer) RegisterStartEgressTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "StartEgress", []string{topic}, s.svc.StartEgress, s.svc.StartEgressAffinity)
}

func (s *egressInternalServer) DeregisterStartEgressTopic(topic string) {
	s.rpc.DeregisterHandler("StartEgress", []string{topic})
}

func (s *egressInternalServer) RegisterListActiveEgressTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "ListActiveEgress", []string{topic}, s.svc.ListActiveEgress, nil)
}

func (s *egressInternalServer) DeregisterListActiveEgressTopic(topic string) {
	s.rpc.DeregisterHandler("ListActiveEgress", []string{topic})
}

func (s *egressInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *egressInternalServer) Kill() {
	s.rpc.Close(true)
}

// ==============================
// EgressHandler Client Interface
// ==============================

type EgressHandlerClient interface {
	UpdateStream(ctx context.Context, topic string, req *livekit2.UpdateStreamRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error)

	StopEgress(ctx context.Context, topic string, req *livekit2.StopEgressRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error)
}

// ==================================
// EgressHandler ServerImpl Interface
// ==================================

type EgressHandlerServerImpl interface {
	UpdateStream(context.Context, *livekit2.UpdateStreamRequest) (*livekit2.EgressInfo, error)

	StopEgress(context.Context, *livekit2.StopEgressRequest) (*livekit2.EgressInfo, error)
}

// ==============================
// EgressHandler Server Interface
// ==============================

type EgressHandlerServer interface {
	RegisterUpdateStreamTopic(topic string) error
	DeregisterUpdateStreamTopic(topic string)
	RegisterStopEgressTopic(topic string) error
	DeregisterStopEgressTopic(topic string)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ====================
// EgressHandler Client
// ====================

type egressHandlerClient struct {
	client *client.RPCClient
}

// NewEgressHandlerClient creates a psrpc client that implements the EgressHandlerClient interface.
func NewEgressHandlerClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (EgressHandlerClient, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressHandler",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("UpdateStream", false, false, true, true)
	sd.RegisterMethod("StopEgress", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &egressHandlerClient{
		client: rpcClient,
	}, nil
}

func (c *egressHandlerClient) UpdateStream(ctx context.Context, topic string, req *livekit2.UpdateStreamRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error) {
	return client.RequestSingle[*livekit2.EgressInfo](ctx, c.client, "UpdateStream", []string{topic}, req, opts...)
}

func (c *egressHandlerClient) StopEgress(ctx context.Context, topic string, req *livekit2.StopEgressRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error) {
	return client.RequestSingle[*livekit2.EgressInfo](ctx, c.client, "StopEgress", []string{topic}, req, opts...)
}

// ====================
// EgressHandler Server
// ====================

type egressHandlerServer struct {
	svc EgressHandlerServerImpl
	rpc *server.RPCServer
}

// NewEgressHandlerServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewEgressHandlerServer(svc EgressHandlerServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (EgressHandlerServer, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressHandler",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("UpdateStream", false, false, true, true)
	sd.RegisterMethod("StopEgress", false, false, true, true)
	return &egressHandlerServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *egressHandlerServer) RegisterUpdateStreamTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "UpdateStream", []string{topic}, s.svc.UpdateStream, nil)
}

func (s *egressHandlerServer) DeregisterUpdateStreamTopic(topic string) {
	s.rpc.DeregisterHandler("UpdateStream", []string{topic})
}

func (s *egressHandlerServer) RegisterStopEgressTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "StopEgress", []string{topic}, s.svc.StopEgress, nil)
}

func (s *egressHandlerServer) DeregisterStopEgressTopic(topic string) {
	s.rpc.DeregisterHandler("StopEgress", []string{topic})
}

func (s *egressHandlerServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *egressHandlerServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor1 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0x12, 0x41,
	0x14, 0x86, 0x9d, 0x12, 0x96, 0xee, 0x41, 0x90, 0x8c, 0x35, 0x4c, 0x69, 0x1b, 0x29, 0x7a, 0xc1,
	0x8d, 0x8b, 0xa1, 0x57, 0xde, 0x69, 0x0d, 0xb1, 0x24, 0x4d, 0x34, 0x60, 0x63, 0xa2, 0x17, 0x64,
	0x99, 0x1d, 0xeb, 0x84, 0xdd, 0x9d, 0x71, 0xe6, 0x50, 0x9e, 0xa1, 0x8f, 0xe1, 0x2b, 0xf4, 0xc2,
	0xa7, 0xf0, 0xa1, 0x0c, 0xb3, 0xb0, 0x5d, 0x68, 0x30, 0x5e, 0xce, 0xff, 0x9f, 0xff, 0xdb, 0x73,
	0xce, 0xcc, 0x42, 0xc3, 0x68, 0xde, 0x13, 0xd7, 0x46, 0x58, 0x1b, 0x68, 0xa3, 0x50, 0xd1, 0x92,
	0xd1, 0xbc, 0x55, 0x53, 0x1a, 0xa5, 0x4a, 0x57, 0x5a, 0xeb, 0x20, 0x96, 0x37, 0x62, 0x26, 0x71,
	0x52, 0xac, 0xec, 0xfc, 0x29, 0x01, 0x1d, 0x63, 0x68, 0x70, 0xe0, 0xd4, 0x91, 0xf8, 0x39, 0x17,
	0x16, 0xe9, 0x11, 0xf8, 0x59, 0xd9, 0x44, 0x46, 0x8c, 0xb4, 0x49, 0xd7, 0x1f, 0xed, 0x67, 0xc2,
	0x30, 0xa2, 0x97, 0x50, 0x37, 0x4a, 0x25, 0x13, 0xae, 0x12, 0xad, 0xac, 0x44, 0xc1, 0xca, 0x6d,
	0xd2, 0xad, 0xf6, 0x5f, 0x04, 0xab, 0x4f, 0x04, 0x23, 0xa5, 0x92, 0xf7, 0x6b, 0x77, 0x83, 0x7c,
	0xf1, 0x68, 0x54, 0x33, 0x45, 0x97, 0xbe, 0x82, 0xd2, 0x42, 0x4c, 0x59, 0xd5, 0x21, 0x0e, 0x73,
	0xc4, 0x17, 0x31, 0xdd, 0x0e, 0x2e, 0xeb, 0xe8, 0x00, 0xaa, 0x3a, 0x34, 0x28, 0xb9, 0xd4, 0x61,
	0x8a, 0xac, 0xe6, 0x62, 0xa7, 0x79, 0xec, 0xd3, 0xbd, 0xb7, 0x1d, 0x2f, 0xe6, 0xe8, 0x47, 0x78,
	0x82, 0x26, 0xe4, 0xb3, 0xc2, 0x10, 0x9e, 0x43, 0xbd, 0xcc, 0x51, 0x9f, 0x97, 0xfe, 0xce, 0x29,
	0xea, 0xb8, 0x61, 0xd3, 0x33, 0x28, 0x3b, 0x85, 0x55, 0x1c, 0xe6, 0x68, 0x13, 0xb3, 0x9d, 0xce,
	0x6a, 0x69, 0x13, 0x2a, 0x6e, 0x93, 0x32, 0x62, 0x25, 0xb7, 0x64, 0x6f, 0x79, 0x1c, 0x46, 0xf4,
	0x00, 0xca, 0xa8, 0x66, 0x22, 0x65, 0xfb, 0x4e, 0xce, 0x0e, 0xf4, 0x19, 0x78, 0x0b, 0x3b, 0x99,
	0x9b, 0x98, 0xf9, 0x99, 0xbc, 0xb0, 0x57, 0x26, 0x3e, 0xf7, 0xa1, 0x62, 0x32, 0x72, 0xe7, 0x10,
	0x9a, 0x97, 0xd2, 0xe2, 0x3b, 0x8e, 0xf2, 0x66, 0xb3, 0xe5, 0xce, 0x1b, 0x60, 0x0f, 0x2d, 0xab,
	0x55, 0x6a, 0x05, 0x3d, 0x01, 0xc8, 0xaf, 0xdb, 0x32, 0xd2, 0x2e, 0x75, 0xfd, 0x91, 0xbf, 0xbe,
	0x6f, 0xdb, 0xff, 0x4d, 0xa0, 0x9e, 0x25, 0x86, 0x29, 0x0a, 0x93, 0x86, 0x31, 0xfd, 0x00, 0xd5,
	0xc2, 0xb3, 0xa1, 0xcd, 0xc0, 0x68, 0x1e, 0x3c, 0x7c, 0x48, 0xad, 0xa7, 0xf9, 0x1e, 0xd6, 0x80,
	0xef, 0xaa, 0x03, 0x77, 0xb7, 0xc4, 0x6b, 0x90, 0xb7, 0xe4, 0x35, 0xa1, 0xdf, 0xa0, 0xb1, 0xdd,
	0x16, 0x3d, 0x76, 0xb4, 0x1d, 0x83, 0xb4, 0x4e, 0x76, 0xb8, 0xd9, 0x2c, 0x39, 0x7c, 0xaf, 0x4b,
	0xfa, 0xbf, 0x08, 0xd4, 0x32, 0xfb, 0x22, 0x4c, 0xa3, 0x58, 0x18, 0x3a, 0x84, 0xc7, 0x57, 0x3a,
	0x0a, 0x51, 0x8c, 0xd1, 0x88, 0x30, 0xa1, 0xc7, 0x79, 0x7f, 0x45, 0xf9, 0x9f, 0xdd, 0x7b, 0x77,
	0xb7, 0x64, 0xaf, 0x41, 0xe8, 0x00, 0x60, 0x8c, 0x4a, 0xaf, 0x7a, 0x6e, 0xe5, 0xa5, 0xf7, 0xe2,
	0xff, 0x60, 0xce, 0x4f, 0xbf, 0x3e, 0xbf, 0x96, 0xf8, 0x63, 0x3e, 0x0d, 0xb8, 0x4a, 0x7a, 0xab,
	0xc2, 0x9e, 0xfb, 0x3b, 0xb9, 0x8a, 0x7b, 0x46, 0xf3, 0xa9, 0xe7, 0x4e, 0x67, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x47, 0x4e, 0x8e, 0x31, 0xe9, 0x03, 0x00, 0x00,
}
