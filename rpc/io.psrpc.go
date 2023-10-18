// Code generated by protoc-gen-psrpc v0.3.3, DO NOT EDIT.
// source: rpc/io.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import google_protobuf3 "google.golang.org/protobuf/types/known/emptypb"
import livekit1 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_3_3

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	// egress
	CreateEgress(ctx context.Context, req *livekit1.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error)

	UpdateEgress(ctx context.Context, req *livekit1.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error)

	GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit1.EgressInfo, error)

	ListEgress(ctx context.Context, req *livekit1.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit1.ListEgressResponse, error)

	// ingress
	GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error)

	UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error)

	// deprecated
	UpdateEgressInfo(ctx context.Context, req *livekit1.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error)
}

// ===========================
// IOInfo ServerImpl Interface
// ===========================

type IOInfoServerImpl interface {
	// egress
	CreateEgress(context.Context, *livekit1.EgressInfo) (*google_protobuf3.Empty, error)

	UpdateEgress(context.Context, *livekit1.EgressInfo) (*google_protobuf3.Empty, error)

	GetEgress(context.Context, *GetEgressRequest) (*livekit1.EgressInfo, error)

	ListEgress(context.Context, *livekit1.ListEgressRequest) (*livekit1.ListEgressResponse, error)

	// ingress
	GetIngressInfo(context.Context, *GetIngressInfoRequest) (*GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *UpdateIngressStateRequest) (*google_protobuf3.Empty, error)

	// deprecated
	UpdateEgressInfo(context.Context, *livekit1.EgressInfo) (*google_protobuf3.Empty, error)
}

// =======================
// IOInfo Server Interface
// =======================

type IOInfoServer interface {

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// IOInfo Client
// =============

type iOInfoClient struct {
	client *client.RPCClient
}

// NewIOInfoClient creates a psrpc client that implements the IOInfoClient interface.
func NewIOInfoClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IOInfoClient, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   clientID,
	}

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	sd.RegisterMethod("GetEgress", false, false, true, true)
	sd.RegisterMethod("ListEgress", false, false, true, true)
	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	sd.RegisterMethod("UpdateIngressState", false, false, true, true)
	sd.RegisterMethod("UpdateEgressInfo", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &iOInfoClient{
		client: rpcClient,
	}, nil
}

func (c *iOInfoClient) CreateEgress(ctx context.Context, req *livekit1.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error) {
	return client.RequestSingle[*google_protobuf3.Empty](ctx, c.client, "CreateEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateEgress(ctx context.Context, req *livekit1.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error) {
	return client.RequestSingle[*google_protobuf3.Empty](ctx, c.client, "UpdateEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit1.EgressInfo, error) {
	return client.RequestSingle[*livekit1.EgressInfo](ctx, c.client, "GetEgress", nil, req, opts...)
}

func (c *iOInfoClient) ListEgress(ctx context.Context, req *livekit1.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit1.ListEgressResponse, error) {
	return client.RequestSingle[*livekit1.ListEgressResponse](ctx, c.client, "ListEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error) {
	return client.RequestSingle[*GetIngressInfoResponse](ctx, c.client, "GetIngressInfo", nil, req, opts...)
}

func (c *iOInfoClient) UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error) {
	return client.RequestSingle[*google_protobuf3.Empty](ctx, c.client, "UpdateIngressState", nil, req, opts...)
}

func (c *iOInfoClient) UpdateEgressInfo(ctx context.Context, req *livekit1.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error) {
	return client.RequestSingle[*google_protobuf3.Empty](ctx, c.client, "UpdateEgressInfo", nil, req, opts...)
}

// =============
// IOInfo Server
// =============

type iOInfoServer struct {
	svc IOInfoServerImpl
	rpc *server.RPCServer
}

// NewIOInfoServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIOInfoServer(serverID string, svc IOInfoServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IOInfoServer, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	var err error
	err = server.RegisterHandler(s, "CreateEgress", nil, svc.CreateEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateEgress", nil, svc.UpdateEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetEgress", false, false, true, true)
	err = server.RegisterHandler(s, "GetEgress", nil, svc.GetEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("ListEgress", false, false, true, true)
	err = server.RegisterHandler(s, "ListEgress", nil, svc.ListEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	err = server.RegisterHandler(s, "GetIngressInfo", nil, svc.GetIngressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateIngressState", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateIngressState", nil, svc.UpdateIngressState, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateEgressInfo", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateEgressInfo", nil, svc.UpdateEgressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &iOInfoServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *iOInfoServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *iOInfoServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor2 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0xa7, 0xae, 0x5d, 0xcc, 0x6b, 0x91, 0x32, 0x6e, 0xca, 0x9a, 0xa0, 0x48, 0x4e, 0x45, 0x65,
	0x02, 0xeb, 0xc1, 0x83, 0x07, 0x41, 0x09, 0x25, 0x58, 0x10, 0x2a, 0x7b, 0xf1, 0x12, 0xb2, 0xc9,
	0xdb, 0x38, 0x24, 0x9b, 0x19, 0x67, 0x26, 0x2d, 0xfb, 0x5d, 0xfd, 0x30, 0x92, 0x99, 0x49, 0xba,
	0x6e, 0xb7, 0xa0, 0x9e, 0xc2, 0xfc, 0xfe, 0xcd, 0x9b, 0xf7, 0x0b, 0x9c, 0x4a, 0x51, 0xc4, 0x8c,
	0x53, 0x21, 0xb9, 0xe6, 0x64, 0x22, 0x45, 0x11, 0xcc, 0x1a, 0x76, 0x83, 0x35, 0xd3, 0x19, 0x56,
	0x12, 0x95, 0xb2, 0x54, 0xe0, 0x0f, 0x28, 0x6b, 0x77, 0xe1, 0xb0, 0xe2, 0xbc, 0x6a, 0x30, 0x36,
	0xa7, 0x55, 0xb7, 0x8e, 0x71, 0x23, 0xf4, 0xd6, 0x92, 0x51, 0x0c, 0x67, 0x97, 0xa8, 0x13, 0xa3,
	0xbf, 0xc6, 0x9f, 0x1d, 0x2a, 0x4d, 0x42, 0xf0, 0x6c, 0x6e, 0xc6, 0xca, 0xf9, 0xd1, 0xab, 0xa3,
	0x0b, 0xef, 0xfa, 0x89, 0x05, 0xd2, 0x32, 0x5a, 0x82, 0x7f, 0x89, 0x3a, 0xb5, 0x37, 0xa4, 0xed,
	0x9a, 0x0f, 0xae, 0x17, 0x00, 0xee, 0xde, 0x3b, 0x9b, 0xe7, 0x90, 0xb4, 0xec, 0x69, 0xa5, 0x25,
	0xe6, 0x9b, 0xac, 0xc6, 0xed, 0xfc, 0x91, 0xa5, 0x2d, 0xf2, 0x05, 0xb7, 0x11, 0x87, 0xf3, 0xfd,
	0x58, 0x25, 0x78, 0xab, 0x90, 0x5c, 0xc0, 0x63, 0xd6, 0xae, 0xb9, 0x49, 0x3c, 0x59, 0xcc, 0xa8,
	0x7b, 0x24, 0xdd, 0xd5, 0x1a, 0x05, 0x99, 0xc1, 0xb1, 0xe6, 0x35, 0xb6, 0x2e, 0xdd, 0x1e, 0x88,
	0x0f, 0xd3, 0x5b, 0x95, 0x75, 0xb2, 0x99, 0x4f, 0x2c, 0x7c, 0xab, 0x96, 0xb2, 0x89, 0x2a, 0x78,
	0xbe, 0x14, 0x65, 0xae, 0xd1, 0xe5, 0x7c, 0xd3, 0xb9, 0xc6, 0xbf, 0x7c, 0xcb, 0x1b, 0x38, 0x56,
	0xbd, 0xdc, 0x5c, 0x74, 0xb2, 0xf0, 0xf7, 0x67, 0xb2, 0x59, 0x56, 0xb3, 0xf8, 0x35, 0x81, 0x69,
	0xfa, 0xb5, 0x1f, 0x93, 0x7c, 0x80, 0xd3, 0xcf, 0x12, 0x73, 0x8d, 0x76, 0xdf, 0xe4, 0xd9, 0x68,
	0x4c, 0xc6, 0xb7, 0x04, 0xe7, 0xd4, 0xf6, 0x45, 0x87, 0xbe, 0x68, 0xd2, 0xf7, 0xd5, 0x9b, 0xed,
	0xc0, 0xff, 0x63, 0x7e, 0x0f, 0xde, 0x58, 0x33, 0xf1, 0xa9, 0x14, 0x05, 0xdd, 0xaf, 0x3d, 0x38,
	0x14, 0x48, 0x12, 0x80, 0x2b, 0xa6, 0x06, 0x67, 0x30, 0x4a, 0xee, 0xc0, 0xc1, 0x1e, 0x1e, 0xe4,
	0x5c, 0x89, 0x29, 0x3c, 0xfd, 0xb3, 0x5e, 0x12, 0x0c, 0x43, 0xdc, 0xff, 0x95, 0x82, 0xf0, 0x20,
	0xe7, 0xa2, 0xae, 0x80, 0xdc, 0x2f, 0x8e, 0xbc, 0x34, 0x96, 0x07, 0x1b, 0x7d, 0x70, 0x31, 0x1f,
	0xe1, 0x6c, 0x77, 0xab, 0x66, 0xb4, 0x7f, 0xd9, 0xec, 0xa7, 0xb7, 0xdf, 0x5f, 0x57, 0x4c, 0xff,
	0xe8, 0x56, 0xb4, 0xe0, 0x9b, 0xd8, 0x19, 0xc7, 0xaf, 0xa8, 0xab, 0x58, 0xa1, 0xbc, 0x61, 0x05,
	0xc6, 0x52, 0x14, 0xab, 0xa9, 0x71, 0xbf, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x75, 0x83, 0xf7,
	0xca, 0xd4, 0x03, 0x00, 0x00,
}
