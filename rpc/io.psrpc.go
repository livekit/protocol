// Code generated by protoc-gen-psrpc v0.2.5, DO NOT EDIT.
// source: rpc/io.proto

package rpc

import context "context"
import psrpc "github.com/livekit/psrpc"
import version "github.com/livekit/psrpc/version"
import google_protobuf2 "google.golang.org/protobuf/types/known/emptypb"
import livekit "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_2_5

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	UpdateEgressInfo(context.Context, *livekit.EgressInfo, ...psrpc.RequestOption) (*google_protobuf2.Empty, error)

	GetIngressInfo(context.Context, *GetIngressInfoRequest, ...psrpc.RequestOption) (*GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *UpdateIngressStateRequest, ...psrpc.RequestOption) (*google_protobuf2.Empty, error)
}

// ===========================
// IOInfo ServerImpl Interface
// ===========================

type IOInfoServerImpl interface {
	UpdateEgressInfo(context.Context, *livekit.EgressInfo) (*google_protobuf2.Empty, error)

	GetIngressInfo(context.Context, *GetIngressInfoRequest) (*GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *UpdateIngressStateRequest) (*google_protobuf2.Empty, error)
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
	client *psrpc.RPCClient
}

// NewIOInfoClient creates a psrpc client that implements the IOInfoClient interface.
func NewIOInfoClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IOInfoClient, error) {
	rpcClient, err := psrpc.NewRPCClient("IOInfo", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &iOInfoClient{
		client: rpcClient,
	}, nil
}

func (c *iOInfoClient) UpdateEgressInfo(ctx context.Context, req *livekit.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf2.Empty, error) {
	return psrpc.RequestSingle[*google_protobuf2.Empty](ctx, c.client, "UpdateEgressInfo", "", req, opts...)
}

func (c *iOInfoClient) GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error) {
	return psrpc.RequestSingle[*GetIngressInfoResponse](ctx, c.client, "GetIngressInfo", "", req, opts...)
}

func (c *iOInfoClient) UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf2.Empty, error) {
	return psrpc.RequestSingle[*google_protobuf2.Empty](ctx, c.client, "UpdateIngressState", "", req, opts...)
}

// =============
// IOInfo Server
// =============

type iOInfoServer struct {
	svc IOInfoServerImpl
	rpc *psrpc.RPCServer
}

// NewIOInfoServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIOInfoServer(serverID string, svc IOInfoServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IOInfoServer, error) {
	s := psrpc.NewRPCServer("IOInfo", serverID, bus, opts...)

	var err error
	err = psrpc.RegisterHandler(s, "UpdateEgressInfo", "", svc.UpdateEgressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	err = psrpc.RegisterHandler(s, "GetIngressInfo", "", svc.GetIngressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	err = psrpc.RegisterHandler(s, "UpdateIngressState", "", svc.UpdateIngressState, nil)
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
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x6d, 0x4b, 0xf3, 0x30,
	0x14, 0xa5, 0xcf, 0x5e, 0x60, 0xd9, 0xc3, 0xc3, 0x43, 0x5c, 0x47, 0xed, 0x50, 0x64, 0x9f, 0x86,
	0x4a, 0x0a, 0xf3, 0x07, 0x08, 0xc2, 0x90, 0xa2, 0x20, 0x4c, 0xf6, 0xc5, 0x2f, 0x65, 0x6b, 0xef,
	0x6a, 0x68, 0xd7, 0xc4, 0x24, 0xdd, 0xd8, 0xff, 0xf1, 0x97, 0xf9, 0x4b, 0xa4, 0x49, 0x36, 0xe7,
	0xdc, 0xc0, 0x4f, 0xe1, 0x9e, 0x7b, 0xef, 0xc9, 0x39, 0xf7, 0xa0, 0xbf, 0x82, 0xc7, 0x01, 0x65,
	0x84, 0x0b, 0xa6, 0x18, 0xae, 0x09, 0x1e, 0xfb, 0x9d, 0x9c, 0x2e, 0x21, 0xa3, 0x2a, 0x82, 0x54,
	0x80, 0x94, 0xa6, 0xe5, 0xbb, 0x1b, 0x94, 0x16, 0xbb, 0x70, 0x2f, 0x65, 0x2c, 0xcd, 0x21, 0xd0,
	0xd5, 0xac, 0x9c, 0x07, 0xb0, 0xe0, 0x6a, 0x6d, 0x9a, 0xfd, 0x09, 0x72, 0xef, 0x41, 0x85, 0x66,
	0x21, 0x2c, 0xe6, 0x6c, 0x0c, 0x6f, 0x25, 0x48, 0x85, 0xcf, 0x10, 0xb2, 0x34, 0x11, 0x4d, 0x3c,
	0xe7, 0xc2, 0x19, 0xb4, 0xc6, 0x2d, 0x8b, 0x84, 0x49, 0xd5, 0x96, 0x4a, 0xc0, 0x74, 0x11, 0x65,
	0xb0, 0xf6, 0xfe, 0x98, 0xb6, 0x41, 0x1e, 0x60, 0xdd, 0x7f, 0x77, 0x50, 0x77, 0x9f, 0x57, 0x72,
	0x56, 0x48, 0xc0, 0x03, 0x54, 0xa7, 0xc5, 0x9c, 0x69, 0xca, 0xf6, 0xb0, 0x43, 0xac, 0x68, 0xb2,
	0x3b, 0xab, 0x27, 0x70, 0x07, 0x35, 0x14, 0xcb, 0xa0, 0xb0, 0xf4, 0xa6, 0xc0, 0x2e, 0x6a, 0xae,
	0x64, 0x54, 0x8a, 0xdc, 0xab, 0x19, 0x78, 0x25, 0x27, 0x22, 0xaf, 0x86, 0x41, 0x08, 0x26, 0xbc,
	0xba, 0x41, 0x75, 0x51, 0xc9, 0x14, 0xc6, 0x50, 0xe5, 0xa2, 0x61, 0x64, 0x5a, 0x24, 0x4c, 0xfa,
	0x29, 0x3a, 0x9d, 0xf0, 0x64, 0xaa, 0xc0, 0x7e, 0xfe, 0xac, 0xa6, 0x0a, 0x7e, 0x79, 0x81, 0x2b,
	0xd4, 0x90, 0xd5, 0xb8, 0x56, 0xd7, 0x1e, 0xba, 0xfb, 0x46, 0x0c, 0x97, 0x99, 0x19, 0x7e, 0x38,
	0xa8, 0x19, 0x3e, 0x55, 0xde, 0xf0, 0x2d, 0xfa, 0x6f, 0xfe, 0x1c, 0x6d, 0xfd, 0xe2, 0x93, 0xed,
	0xf2, 0x17, 0xe8, 0x77, 0x89, 0x09, 0x8e, 0x6c, 0x82, 0x23, 0xa3, 0x2a, 0x38, 0x1c, 0xa2, 0x7f,
	0xdf, 0x4f, 0x8b, 0x7d, 0x22, 0x78, 0x4c, 0x0e, 0xe6, 0xe8, 0xf7, 0x0e, 0xf6, 0x6c, 0x16, 0x8f,
	0x08, 0xff, 0xf4, 0x8f, 0xcf, 0xf5, 0xca, 0xd1, 0xc3, 0x1c, 0x13, 0x76, 0x77, 0xfd, 0x72, 0x99,
	0x52, 0xf5, 0x5a, 0xce, 0x48, 0xcc, 0x16, 0x81, 0x75, 0xb4, 0x7d, 0x79, 0x96, 0x06, 0x12, 0xc4,
	0x92, 0xc6, 0x10, 0x08, 0x1e, 0xcf, 0x9a, 0x7a, 0xfb, 0xe6, 0x33, 0x00, 0x00, 0xff, 0xff, 0xca,
	0x7b, 0x14, 0x3c, 0xdf, 0x02, 0x00, 0x00,
}
