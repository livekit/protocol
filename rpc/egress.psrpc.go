// Code generated by protoc-gen-psrpc v0.3.2, DO NOT EDIT.
// source: rpc/egress.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import livekit1 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_3_2

// ===============================
// EgressInternal Client Interface
// ===============================

type EgressInternalClient interface {
	StartEgress(ctx context.Context, topic string, req *StartEgressRequest, opts ...psrpc.RequestOption) (*livekit1.EgressInfo, error)

	ListActiveEgress(ctx context.Context, req *ListActiveEgressRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*ListActiveEgressResponse], error)
}

// ===================================
// EgressInternal ServerImpl Interface
// ===================================

type EgressInternalServerImpl interface {
	StartEgress(context.Context, *StartEgressRequest) (*livekit1.EgressInfo, error)
	StartEgressAffinity(*StartEgressRequest) float32

	ListActiveEgress(context.Context, *ListActiveEgressRequest) (*ListActiveEgressResponse, error)
}

// ===============================
// EgressInternal Server Interface
// ===============================

type EgressInternalServer interface {
	RegisterStartEgressTopic(topic string) error
	DeregisterStartEgressTopic(topic string)

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
func NewEgressInternalClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (EgressInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressInternal",
		ID:   clientID,
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

func (c *egressInternalClient) StartEgress(ctx context.Context, topic string, req *StartEgressRequest, opts ...psrpc.RequestOption) (*livekit1.EgressInfo, error) {
	return client.RequestSingle[*livekit1.EgressInfo](ctx, c.client, "StartEgress", []string{topic}, req, opts...)
}

func (c *egressInternalClient) ListActiveEgress(ctx context.Context, req *ListActiveEgressRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*ListActiveEgressResponse], error) {
	return client.RequestMulti[*ListActiveEgressResponse](ctx, c.client, "ListActiveEgress", nil, req, opts...)
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
func NewEgressInternalServer(serverID string, svc EgressInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (EgressInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressInternal",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("StartEgress", true, false, true, false)
	sd.RegisterMethod("ListActiveEgress", false, true, false, false)
	var err error
	err = server.RegisterHandler(s, "ListActiveEgress", nil, svc.ListActiveEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

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
	UpdateStream(ctx context.Context, topic string, req *livekit1.UpdateStreamRequest, opts ...psrpc.RequestOption) (*livekit1.EgressInfo, error)

	StopEgress(ctx context.Context, topic string, req *livekit1.StopEgressRequest, opts ...psrpc.RequestOption) (*livekit1.EgressInfo, error)
}

// ==================================
// EgressHandler ServerImpl Interface
// ==================================

type EgressHandlerServerImpl interface {
	UpdateStream(context.Context, *livekit1.UpdateStreamRequest) (*livekit1.EgressInfo, error)

	StopEgress(context.Context, *livekit1.StopEgressRequest) (*livekit1.EgressInfo, error)
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
func NewEgressHandlerClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (EgressHandlerClient, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressHandler",
		ID:   clientID,
	}

	sd.RegisterMethod("UpdateStream", false, false, true, false)
	sd.RegisterMethod("StopEgress", false, false, true, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &egressHandlerClient{
		client: rpcClient,
	}, nil
}

func (c *egressHandlerClient) UpdateStream(ctx context.Context, topic string, req *livekit1.UpdateStreamRequest, opts ...psrpc.RequestOption) (*livekit1.EgressInfo, error) {
	return client.RequestSingle[*livekit1.EgressInfo](ctx, c.client, "UpdateStream", []string{topic}, req, opts...)
}

func (c *egressHandlerClient) StopEgress(ctx context.Context, topic string, req *livekit1.StopEgressRequest, opts ...psrpc.RequestOption) (*livekit1.EgressInfo, error) {
	return client.RequestSingle[*livekit1.EgressInfo](ctx, c.client, "StopEgress", []string{topic}, req, opts...)
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
func NewEgressHandlerServer(serverID string, svc EgressHandlerServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (EgressHandlerServer, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressHandler",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("UpdateStream", false, false, true, false)
	sd.RegisterMethod("StopEgress", false, false, true, false)
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

var psrpcFileDescriptor0 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6e, 0x12, 0x4d,
	0x18, 0xfe, 0xa6, 0x14, 0x0a, 0x2f, 0xa5, 0x1f, 0x19, 0x6b, 0x3a, 0xdd, 0xb6, 0x09, 0xa2, 0x26,
	0xa4, 0xd1, 0xc5, 0xd0, 0x13, 0xf5, 0xac, 0x35, 0x24, 0x25, 0xa9, 0x51, 0xb7, 0x36, 0x4d, 0x3c,
	0x21, 0xc3, 0xee, 0x14, 0x27, 0x2c, 0x3b, 0xe3, 0xcc, 0x40, 0xc3, 0x25, 0xf4, 0x32, 0xbc, 0x03,
	0xd3, 0x4b, 0xf2, 0x4a, 0x0c, 0x33, 0xcb, 0x76, 0xa1, 0xc1, 0x78, 0x04, 0xfb, 0xfc, 0xcd, 0xfb,
	0x07, 0x75, 0x25, 0xc3, 0x36, 0x1b, 0x2a, 0xa6, 0xb5, 0x2f, 0x95, 0x30, 0x02, 0x17, 0x94, 0x0c,
	0xbd, 0xfd, 0xa1, 0x10, 0xc3, 0x98, 0xb5, 0x2d, 0x34, 0x98, 0xdc, 0xb4, 0x69, 0x32, 0x73, 0xbc,
	0x57, 0x13, 0xd2, 0x70, 0x91, 0xa4, 0x72, 0x6f, 0x37, 0xe6, 0x53, 0x36, 0xe2, 0xa6, 0x9f, 0x0f,
	0x69, 0xfe, 0xde, 0x04, 0x7c, 0x69, 0xa8, 0x32, 0x5d, 0x8b, 0x06, 0xec, 0xc7, 0x84, 0x69, 0x83,
	0x0f, 0xa0, 0xe2, 0x64, 0x7d, 0x1e, 0x11, 0xd4, 0x40, 0xad, 0x4a, 0x50, 0x76, 0x40, 0x2f, 0xc2,
	0x17, 0xb0, 0xa3, 0x84, 0x18, 0xf7, 0x43, 0x31, 0x96, 0x42, 0x73, 0xc3, 0x48, 0xb1, 0x81, 0x5a,
	0xd5, 0xce, 0x73, 0x3f, 0x7d, 0xc2, 0x0f, 0x84, 0x18, 0x7f, 0x58, 0xb0, 0x4b, 0xc9, 0xe7, 0xff,
	0x05, 0x35, 0x95, 0x67, 0xf1, 0x6b, 0x28, 0xdc, 0xb2, 0x01, 0xa9, 0xda, 0x88, 0xfd, 0x2c, 0xe2,
	0x9a, 0x0d, 0x56, 0x8d, 0x73, 0x1d, 0xee, 0x42, 0x55, 0x52, 0x65, 0x78, 0xc8, 0x25, 0x4d, 0x0c,
	0xa9, 0x59, 0xdb, 0xb3, 0xcc, 0xf6, 0xf9, 0x81, 0x5b, 0xb5, 0xe7, 0x7d, 0xf8, 0x13, 0xfc, 0x6f,
	0x14, 0x0d, 0x47, 0xb9, 0x26, 0x4a, 0x36, 0xea, 0x45, 0x16, 0xf5, 0x75, 0xce, 0xaf, 0xed, 0x62,
	0xc7, 0x2c, 0xd1, 0xf8, 0x04, 0x8a, 0x16, 0x21, 0x5b, 0x36, 0xe6, 0x60, 0x39, 0x66, 0xd5, 0xed,
	0xb4, 0x78, 0x0f, 0xb6, 0xec, 0x24, 0x79, 0x44, 0x0a, 0x76, 0xc8, 0xa5, 0xf9, 0x67, 0x2f, 0xc2,
	0xbb, 0x50, 0x34, 0x62, 0xc4, 0x12, 0x52, 0xb6, 0xb0, 0xfb, 0xc0, 0x4f, 0xa1, 0x74, 0xab, 0xfb,
	0x13, 0x15, 0x93, 0x8a, 0x83, 0x6f, 0xf5, 0x95, 0x8a, 0xf1, 0x29, 0x94, 0xc7, 0xcc, 0xd0, 0x88,
	0x1a, 0x4a, 0xb6, 0x1b, 0x85, 0x56, 0xb5, 0xf3, 0xd2, 0x57, 0x32, 0xf4, 0x1f, 0xef, 0xd5, 0xff,
	0x98, 0xea, 0xba, 0x89, 0x51, 0xb3, 0x20, 0xb3, 0x79, 0x5f, 0xa0, 0xb6, 0x44, 0xe1, 0x3a, 0x14,
	0x46, 0x6c, 0x96, 0xae, 0x7e, 0xfe, 0x17, 0x1f, 0x43, 0x71, 0x4a, 0xe3, 0x09, 0x23, 0x1b, 0xb6,
	0xc1, 0x5d, 0xdf, 0x5d, 0x9e, 0xbf, 0xb8, 0x3c, 0xff, 0x34, 0x99, 0x05, 0x4e, 0xf2, 0x7e, 0xe3,
	0x2d, 0x3a, 0xab, 0xc0, 0x96, 0x72, 0xaf, 0x36, 0xf7, 0x61, 0xef, 0x82, 0x6b, 0x73, 0x1a, 0x1a,
	0x3e, 0x5d, 0x1e, 0x64, 0xf3, 0x1d, 0x90, 0xc7, 0x94, 0x96, 0x22, 0xd1, 0x0c, 0x1f, 0x01, 0x64,
	0x47, 0xa8, 0x09, 0x6a, 0x14, 0x5a, 0x95, 0xa0, 0xb2, 0xb8, 0x42, 0xdd, 0xf9, 0x85, 0x60, 0xc7,
	0x39, 0x7a, 0x89, 0x61, 0x2a, 0xa1, 0xf1, 0xfc, 0x38, 0x72, 0x4d, 0xe3, 0xbd, 0x35, 0x63, 0xf0,
	0x9e, 0x64, 0xdb, 0x59, 0x04, 0xdc, 0x88, 0x66, 0xf9, 0xfe, 0x0e, 0x6d, 0xd6, 0xd1, 0x1b, 0x84,
	0xaf, 0xa1, 0xbe, 0x5a, 0x14, 0x3e, 0xb4, 0x59, 0x6b, 0xda, 0xf0, 0x8e, 0xd6, 0xb0, 0xae, 0x93,
	0x66, 0xe9, 0xfe, 0x0e, 0x6d, 0xb4, 0x50, 0xe7, 0x27, 0x82, 0x9a, 0xa3, 0xce, 0x69, 0x12, 0xc5,
	0x4c, 0xe1, 0x1e, 0x6c, 0x5f, 0xc9, 0x88, 0x1a, 0x76, 0x69, 0x14, 0xa3, 0x63, 0x7c, 0x98, 0x55,
	0x96, 0x87, 0xff, 0x5a, 0xb7, 0x0d, 0xaf, 0x23, 0xdc, 0x05, 0xb8, 0x34, 0x42, 0xa6, 0xf5, 0x7a,
	0x99, 0xf4, 0x01, 0xfc, 0x97, 0x98, 0xb3, 0x57, 0xdf, 0x8e, 0x87, 0xdc, 0x7c, 0x9f, 0x0c, 0xfc,
	0x50, 0x8c, 0xdb, 0xa9, 0x30, 0xfb, 0x95, 0xa3, 0x61, 0x5b, 0x33, 0x35, 0xe5, 0x21, 0x6b, 0x2b,
	0x19, 0x0e, 0x4a, 0x76, 0xfd, 0x27, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x04, 0x2a, 0x96, 0x18,
	0x9f, 0x04, 0x00, 0x00,
}
