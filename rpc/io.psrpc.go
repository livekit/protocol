// Code generated by protoc-gen-psrpc v0.5.1, DO NOT EDIT.
// source: rpc/io.proto

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

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	// egress
	CreateEgress(ctx context.Context, req *livekit2.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	UpdateEgress(ctx context.Context, req *livekit2.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error)

	ListEgress(ctx context.Context, req *livekit2.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit2.ListEgressResponse, error)

	// ingress
	GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error)

	UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// sip
	GetSIPTrunkAuthentication(ctx context.Context, req *GetSIPTrunkAuthenticationRequest, opts ...psrpc.RequestOption) (*GetSIPTrunkAuthenticationResponse, error)

	EvaluateSIPDispatchRules(ctx context.Context, req *EvaluateSIPDispatchRulesRequest, opts ...psrpc.RequestOption) (*EvaluateSIPDispatchRulesResponse, error)
}

// ===========================
// IOInfo ServerImpl Interface
// ===========================

type IOInfoServerImpl interface {
	// egress
	CreateEgress(context.Context, *livekit2.EgressInfo) (*google_protobuf.Empty, error)

	UpdateEgress(context.Context, *livekit2.EgressInfo) (*google_protobuf.Empty, error)

	GetEgress(context.Context, *GetEgressRequest) (*livekit2.EgressInfo, error)

	ListEgress(context.Context, *livekit2.ListEgressRequest) (*livekit2.ListEgressResponse, error)

	// ingress
	GetIngressInfo(context.Context, *GetIngressInfoRequest) (*GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *UpdateIngressStateRequest) (*google_protobuf.Empty, error)

	// sip
	GetSIPTrunkAuthentication(context.Context, *GetSIPTrunkAuthenticationRequest) (*GetSIPTrunkAuthenticationResponse, error)

	EvaluateSIPDispatchRules(context.Context, *EvaluateSIPDispatchRulesRequest) (*EvaluateSIPDispatchRulesResponse, error)
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
func NewIOInfoClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IOInfoClient, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	sd.RegisterMethod("GetEgress", false, false, true, true)
	sd.RegisterMethod("ListEgress", false, false, true, true)
	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	sd.RegisterMethod("UpdateIngressState", false, false, true, true)
	sd.RegisterMethod("GetSIPTrunkAuthentication", false, false, true, true)
	sd.RegisterMethod("EvaluateSIPDispatchRules", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &iOInfoClient{
		client: rpcClient,
	}, nil
}

func (c *iOInfoClient) CreateEgress(ctx context.Context, req *livekit2.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "CreateEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateEgress(ctx context.Context, req *livekit2.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error) {
	return client.RequestSingle[*livekit2.EgressInfo](ctx, c.client, "GetEgress", nil, req, opts...)
}

func (c *iOInfoClient) ListEgress(ctx context.Context, req *livekit2.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit2.ListEgressResponse, error) {
	return client.RequestSingle[*livekit2.ListEgressResponse](ctx, c.client, "ListEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error) {
	return client.RequestSingle[*GetIngressInfoResponse](ctx, c.client, "GetIngressInfo", nil, req, opts...)
}

func (c *iOInfoClient) UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateIngressState", nil, req, opts...)
}

func (c *iOInfoClient) GetSIPTrunkAuthentication(ctx context.Context, req *GetSIPTrunkAuthenticationRequest, opts ...psrpc.RequestOption) (*GetSIPTrunkAuthenticationResponse, error) {
	return client.RequestSingle[*GetSIPTrunkAuthenticationResponse](ctx, c.client, "GetSIPTrunkAuthentication", nil, req, opts...)
}

func (c *iOInfoClient) EvaluateSIPDispatchRules(ctx context.Context, req *EvaluateSIPDispatchRulesRequest, opts ...psrpc.RequestOption) (*EvaluateSIPDispatchRulesResponse, error) {
	return client.RequestSingle[*EvaluateSIPDispatchRulesResponse](ctx, c.client, "EvaluateSIPDispatchRules", nil, req, opts...)
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
func NewIOInfoServer(svc IOInfoServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IOInfoServer, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   rand.NewServerID(),
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

	sd.RegisterMethod("GetSIPTrunkAuthentication", false, false, true, true)
	err = server.RegisterHandler(s, "GetSIPTrunkAuthentication", nil, svc.GetSIPTrunkAuthentication, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("EvaluateSIPDispatchRules", false, false, true, true)
	err = server.RegisterHandler(s, "EvaluateSIPDispatchRules", nil, svc.EvaluateSIPDispatchRules, nil)
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

var psrpcFileDescriptor4 = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xf3, 0x44,
	0x18, 0x55, 0xae, 0x7f, 0xf3, 0xb5, 0x54, 0xd5, 0xfc, 0x49, 0x49, 0x1d, 0x41, 0x43, 0xa0, 0xa8,
	0x82, 0xca, 0x11, 0x65, 0xc1, 0x82, 0x55, 0x81, 0xa8, 0x58, 0x54, 0x25, 0x4a, 0xe9, 0x06, 0x16,
	0x96, 0x63, 0x4f, 0x9c, 0x21, 0xce, 0xcc, 0x30, 0x33, 0x6e, 0xd5, 0x1d, 0x2b, 0x9e, 0x8b, 0x47,
	0xe2, 0x11, 0xd0, 0x5c, 0xec, 0xa4, 0x6d, 0xd2, 0x20, 0x56, 0xc9, 0x9c, 0xef, 0x76, 0x7c, 0x66,
	0xbe, 0x03, 0x07, 0x82, 0xc7, 0x43, 0xc2, 0x7c, 0x2e, 0x98, 0x62, 0xa8, 0x26, 0x78, 0xec, 0xb5,
	0x33, 0xf2, 0x80, 0x17, 0x44, 0x85, 0x38, 0x15, 0x58, 0x4a, 0x1b, 0xf2, 0x3a, 0x05, 0x4a, 0xe8,
	0x3a, 0xdc, 0x4b, 0x19, 0x4b, 0x33, 0x3c, 0x34, 0xa7, 0x69, 0x3e, 0x1b, 0xe2, 0x25, 0x57, 0x4f,
	0x36, 0x38, 0x18, 0xc2, 0xd1, 0x35, 0x56, 0x23, 0x93, 0x3f, 0xc1, 0x7f, 0xe4, 0x58, 0x2a, 0xd4,
	0x83, 0x96, 0xed, 0x1b, 0x92, 0xa4, 0x5b, 0xe9, 0x57, 0xce, 0x5b, 0x93, 0x3d, 0x0b, 0x04, 0xc9,
	0xe0, 0x1e, 0x3a, 0xd7, 0x58, 0x05, 0x76, 0x42, 0x40, 0x67, 0xac, 0xa8, 0xfa, 0x08, 0xc0, 0xcd,
	0x5d, 0x95, 0xb5, 0x1c, 0x12, 0x24, 0x3a, 0x2c, 0x95, 0xc0, 0xd1, 0x32, 0x5c, 0xe0, 0xa7, 0x6e,
	0xd5, 0x86, 0x2d, 0xf2, 0x13, 0x7e, 0x1a, 0x30, 0x38, 0x7e, 0xd9, 0x56, 0x72, 0x46, 0x25, 0x46,
	0xe7, 0x50, 0x27, 0x74, 0xc6, 0x4c, 0xc7, 0xfd, 0xcb, 0xb6, 0xef, 0x3e, 0xd2, 0x5f, 0xcf, 0x35,
	0x19, 0xa8, 0x0d, 0x0d, 0xc5, 0x16, 0x98, 0xba, 0xee, 0xf6, 0x80, 0x3a, 0xd0, 0x7c, 0x94, 0x61,
	0x2e, 0xb2, 0x6e, 0xcd, 0xc2, 0x8f, 0xf2, 0x5e, 0x64, 0x83, 0x14, 0x4e, 0xee, 0x79, 0x12, 0x29,
	0xec, 0xfa, 0xdc, 0xa9, 0x48, 0xe1, 0xff, 0xf8, 0x2d, 0x5f, 0x42, 0x43, 0xea, 0x74, 0x33, 0x68,
	0xff, 0xb2, 0xf3, 0x92, 0x93, 0xed, 0x65, 0x73, 0x06, 0x7f, 0x56, 0xa0, 0x7f, 0x8d, 0xd5, 0x5d,
	0x30, 0xfe, 0x45, 0xe4, 0x74, 0x71, 0x95, 0xab, 0x39, 0xa6, 0x8a, 0xc4, 0x91, 0x22, 0x8c, 0x16,
	0x03, 0x11, 0xd4, 0x67, 0x82, 0x2d, 0x1d, 0x73, 0xf3, 0x1f, 0x1d, 0x42, 0x55, 0x31, 0x47, 0xba,
	0xaa, 0x18, 0x3a, 0x85, 0x7d, 0x29, 0xe2, 0x30, 0x4a, 0x12, 0x3d, 0xa3, 0x5b, 0x37, 0x01, 0x90,
	0x22, 0xbe, 0xb2, 0x08, 0xfa, 0x10, 0xde, 0x29, 0x16, 0xce, 0x99, 0x54, 0xdd, 0x86, 0x09, 0x36,
	0x15, 0xfb, 0x91, 0x49, 0x35, 0xf8, 0x0d, 0x3e, 0x79, 0x83, 0x81, 0xd3, 0xd9, 0x83, 0xbd, 0x5c,
	0x62, 0x41, 0xa3, 0x25, 0x2e, 0x2e, 0xbd, 0x38, 0xeb, 0x18, 0x8f, 0xa4, 0x7c, 0x64, 0x22, 0x71,
	0x14, 0xcb, 0xf3, 0xe0, 0xaf, 0x2a, 0x9c, 0x8e, 0x1e, 0xa2, 0x2c, 0x8f, 0x14, 0xbe, 0x0b, 0xc6,
	0x3f, 0x10, 0xc9, 0x23, 0x15, 0xcf, 0x27, 0x79, 0x86, 0xcb, 0x17, 0x75, 0x01, 0x48, 0x12, 0x1e,
	0xf2, 0x48, 0x28, 0x12, 0x13, 0x1e, 0x51, 0xb5, 0xd2, 0xf5, 0x48, 0x12, 0x3e, 0x5e, 0x05, 0x82,
	0x04, 0x9d, 0xc1, 0x61, 0x1c, 0x65, 0x19, 0xa1, 0x69, 0x48, 0xf3, 0xe5, 0x14, 0x0b, 0x37, 0xf3,
	0x03, 0x87, 0xde, 0x1a, 0x10, 0x7d, 0x0a, 0x06, 0xc0, 0x49, 0x91, 0x65, 0xa5, 0x3a, 0xb0, 0xa0,
	0x4b, 0xda, 0x29, 0xda, 0x11, 0xd4, 0x38, 0xa1, 0x4e, 0x30, 0xfd, 0x57, 0x3f, 0x18, 0xca, 0x42,
	0x0d, 0x36, 0xfb, 0x95, 0xf3, 0xbd, 0x49, 0x83, 0xb2, 0x31, 0xa1, 0xba, 0x93, 0x1b, 0x67, 0x14,
	0x7e, 0x67, 0x3b, 0x59, 0xc8, 0xa8, 0xfc, 0x77, 0x05, 0xfa, 0xdb, 0x85, 0x70, 0x2a, 0xf7, 0xa0,
	0x25, 0x18, 0x5b, 0x86, 0xeb, 0x32, 0x6b, 0xe0, 0x56, 0xcb, 0xfc, 0x15, 0xb4, 0x9f, 0x4b, 0xa4,
	0xaf, 0x49, 0x15, 0xdb, 0xf2, 0x9e, 0xaf, 0xab, 0x64, 0x43, 0x9a, 0x95, 0xb0, 0x22, 0x1b, 0xc6,
	0x35, 0xc3, 0x18, 0x1c, 0xa4, 0x69, 0x97, 0x4b, 0x51, 0xdf, 0xbc, 0x14, 0x8d, 0xb5, 0xa5, 0xb8,
	0xfc, 0xa7, 0x0e, 0xcd, 0xe0, 0x67, 0xbd, 0x52, 0xe8, 0x5b, 0x38, 0xf8, 0x5e, 0xe0, 0x48, 0x61,
	0xeb, 0x0d, 0xe8, 0x7d, 0xf9, 0xc8, 0x47, 0xe5, 0xde, 0x79, 0xc7, 0xbe, 0xf5, 0x16, 0xbf, 0xf0,
	0x16, 0x7f, 0xa4, 0xbd, 0x45, 0x17, 0xdb, 0xe5, 0xfa, 0x3f, 0xc5, 0xdf, 0x40, 0xab, 0xb4, 0x24,
	0xd4, 0xf1, 0x05, 0x8f, 0xfd, 0x97, 0x16, 0xe5, 0x6d, 0x6a, 0x88, 0x46, 0x00, 0x37, 0x44, 0x16,
	0x95, 0x5e, 0x99, 0xb2, 0x02, 0x8b, 0xf2, 0xde, 0xc6, 0x98, 0xbb, 0xa2, 0x00, 0x0e, 0x9f, 0x5b,
	0x11, 0xf2, 0x0a, 0x12, 0xaf, 0x6d, 0xcf, 0xeb, 0x6d, 0x8c, 0xb9, 0x56, 0x37, 0x80, 0x5e, 0x9b,
	0x0c, 0xfa, 0xd8, 0x94, 0x6c, 0x75, 0x9f, 0xad, 0xc2, 0xfc, 0x0e, 0x27, 0x5b, 0xd7, 0x18, 0x9d,
	0x15, 0x3c, 0xde, 0x34, 0x1a, 0xef, 0xf3, 0x5d, 0x69, 0x8e, 0x79, 0x0a, 0xdd, 0x6d, 0x6f, 0x19,
	0x7d, 0x66, 0x7a, 0xec, 0xd8, 0x79, 0xef, 0x6c, 0x47, 0x96, 0x1d, 0xf4, 0xdd, 0xc5, 0xaf, 0x5f,
	0xa4, 0x44, 0xcd, 0xf3, 0xa9, 0x1f, 0xb3, 0xe5, 0xd0, 0x5d, 0x4b, 0xf9, 0xcb, 0x17, 0xe9, 0x50,
	0x62, 0xf1, 0x40, 0x62, 0x3c, 0x14, 0x3c, 0x9e, 0x36, 0x8d, 0x24, 0x5f, 0xff, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x80, 0x72, 0x79, 0x3f, 0x14, 0x07, 0x00, 0x00,
}
