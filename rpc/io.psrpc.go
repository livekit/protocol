// Code generated by protoc-gen-psrpc v0.5.0, DO NOT EDIT.
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

	UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

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

	UpdateMetrics(context.Context, *UpdateMetricsRequest) (*google_protobuf.Empty, error)

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
	sd.RegisterMethod("UpdateMetrics", false, false, true, true)
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

func (c *iOInfoClient) UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateMetrics", nil, req, opts...)
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

	sd.RegisterMethod("UpdateMetrics", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateMetrics", nil, svc.UpdateMetrics, nil)
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

var psrpcFileDescriptor3 = []byte{
	// 893 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5b, 0x73, 0x1b, 0x35,
	0x18, 0x1d, 0x5f, 0x1b, 0x7f, 0x4e, 0x32, 0x19, 0xd5, 0x2e, 0xce, 0x66, 0xa0, 0xc6, 0x10, 0xf0,
	0x40, 0x67, 0x3d, 0x84, 0x07, 0x18, 0x78, 0xa1, 0x2d, 0x26, 0x78, 0x08, 0x25, 0xe3, 0xe0, 0x17,
	0x78, 0xd8, 0x51, 0x76, 0xe5, 0x8d, 0xf0, 0x5a, 0x12, 0x92, 0xd6, 0x69, 0xde, 0x78, 0x02, 0x7e,
	0x10, 0x3f, 0x80, 0x9f, 0xc6, 0xe8, 0xb2, 0xbe, 0xb4, 0x76, 0xc3, 0xf4, 0xc9, 0xab, 0xf3, 0x5d,
	0x74, 0x74, 0xa4, 0xef, 0x18, 0xf6, 0xa5, 0x88, 0x07, 0x94, 0x87, 0x42, 0x72, 0xcd, 0x51, 0x45,
	0x8a, 0x38, 0x68, 0x65, 0x74, 0x41, 0x66, 0x54, 0x47, 0x24, 0x95, 0x44, 0x29, 0x17, 0x0a, 0xda,
	0x05, 0x4a, 0xd9, 0x3a, 0x7c, 0x92, 0x72, 0x9e, 0x66, 0x64, 0x60, 0x57, 0xd7, 0xf9, 0x74, 0x40,
	0xe6, 0x42, 0xdf, 0xb9, 0x60, 0x6f, 0x00, 0x47, 0xe7, 0x44, 0x0f, 0x6d, 0xfe, 0x98, 0xfc, 0x9e,
	0x13, 0xa5, 0xd1, 0x09, 0x34, 0x5c, 0xdf, 0x88, 0x26, 0x9d, 0x52, 0xb7, 0xd4, 0x6f, 0x8c, 0xf7,
	0x1c, 0x30, 0x4a, 0x7a, 0x7f, 0x95, 0xa0, 0x35, 0x11, 0x09, 0xd6, 0xe4, 0x47, 0xa2, 0x25, 0x8d,
	0x97, 0x55, 0x1f, 0x43, 0x95, 0xb2, 0x29, 0xb7, 0x05, 0xcd, 0xb3, 0x87, 0xa1, 0x27, 0x13, 0xba,
	0xde, 0x23, 0x36, 0xe5, 0x63, 0x9b, 0x80, 0x7a, 0x70, 0x80, 0x17, 0x69, 0x14, 0x8b, 0x3c, 0xca,
	0x15, 0x4e, 0x49, 0xa7, 0xd2, 0x2d, 0xf5, 0xcb, 0xe3, 0x26, 0x5e, 0xa4, 0xcf, 0x45, 0x3e, 0x31,
	0x90, 0xc9, 0x99, 0xe3, 0x97, 0x6b, 0x39, 0x55, 0x97, 0x33, 0xc7, 0x2f, 0x8b, 0x9c, 0xde, 0x04,
	0xda, 0xe7, 0x44, 0x8f, 0xd8, 0xaa, 0xbf, 0x67, 0xf2, 0x2e, 0x80, 0x57, 0x60, 0x75, 0x80, 0x86,
	0x47, 0x46, 0x89, 0x09, 0x2b, 0x2d, 0x09, 0x9e, 0x47, 0x33, 0x72, 0xd7, 0x29, 0xbb, 0xb0, 0x43,
	0x7e, 0x20, 0x77, 0xbd, 0xbf, 0xcb, 0xf0, 0xe8, 0xd5, 0xbe, 0x4a, 0x70, 0xa6, 0x08, 0xea, 0x6f,
	0x1c, 0xb1, 0xb5, 0x3c, 0xe2, 0x7a, 0xae, 0x3b, 0x63, 0x0b, 0x6a, 0x9a, 0xcf, 0x08, 0xf3, 0xed,
	0xdd, 0x02, 0xb5, 0xa1, 0x7e, 0xab, 0xa2, 0x5c, 0x66, 0xf6, 0xc8, 0x8d, 0x71, 0xed, 0x56, 0x4d,
	0x64, 0x86, 0x26, 0x70, 0x98, 0xf1, 0x34, 0xa5, 0x2c, 0x8d, 0xa6, 0x94, 0x64, 0x89, 0xea, 0x54,
	0xbb, 0x95, 0x7e, 0xf3, 0x2c, 0x0c, 0xa5, 0x88, 0xc3, 0xed, 0x5c, 0xc2, 0x0b, 0x57, 0xf1, 0x9d,
	0x2d, 0x18, 0x32, 0x2d, 0xef, 0xc6, 0x07, 0xd9, 0x3a, 0x16, 0x7c, 0x03, 0xe8, 0xf5, 0x24, 0x74,
	0x04, 0x15, 0x73, 0x6c, 0xa7, 0x8a, 0xf9, 0x34, 0x5c, 0x17, 0x38, 0xcb, 0x49, 0xc1, 0xd5, 0x2e,
	0xbe, 0x2a, 0x7f, 0x59, 0xea, 0xa5, 0x70, 0xec, 0xae, 0xda, 0x13, 0xb8, 0xd2, 0x58, 0x93, 0xff,
	0xa9, 0xf2, 0xa7, 0x50, 0x53, 0x26, 0xdd, 0x76, 0x6d, 0x9e, 0xb5, 0x5f, 0x15, 0xcb, 0xf5, 0x72,
	0x39, 0xbd, 0x3f, 0x4a, 0xd0, 0x3d, 0x27, 0xfa, 0x6a, 0x74, 0xf9, 0xb3, 0xcc, 0xd9, 0xec, 0x69,
	0xae, 0x6f, 0x08, 0xd3, 0x34, 0xc6, 0x9a, 0x72, 0x56, 0x6c, 0x88, 0xa0, 0x3a, 0x95, 0x7c, 0xee,
	0x69, 0xda, 0x6f, 0x74, 0x08, 0x65, 0xcd, 0xbd, 0x9a, 0x65, 0xcd, 0xd1, 0x63, 0x68, 0x2a, 0x19,
	0x47, 0x38, 0x49, 0xcc, 0x1e, 0xf6, 0xd5, 0x34, 0xc6, 0xa0, 0x64, 0xfc, 0xd4, 0x21, 0xe8, 0x1d,
	0x78, 0xa0, 0x79, 0x74, 0xc3, 0x95, 0xee, 0xd4, 0x6c, 0xb0, 0xae, 0xf9, 0xf7, 0x5c, 0xe9, 0xde,
	0xaf, 0xf0, 0xfe, 0x1b, 0x18, 0xf8, 0x07, 0x10, 0xc0, 0x5e, 0xae, 0x88, 0x64, 0x78, 0x4e, 0x8a,
	0xc1, 0x28, 0xd6, 0x26, 0x26, 0xb0, 0x52, 0xb7, 0x5c, 0x26, 0x9e, 0xe2, 0x72, 0xdd, 0xfb, 0xb3,
	0x0c, 0x8f, 0x87, 0x46, 0x57, 0xac, 0xc9, 0xd5, 0xe8, 0xf2, 0x5b, 0xaa, 0x04, 0xd6, 0xf1, 0xcd,
	0x38, 0xcf, 0xc8, 0x72, 0x7e, 0x9e, 0x00, 0x52, 0x54, 0x44, 0x02, 0x4b, 0x4d, 0x63, 0x2a, 0x30,
	0xd3, 0x2b, 0x5d, 0x8f, 0x14, 0x15, 0x97, 0xab, 0xc0, 0x28, 0x41, 0xa7, 0x70, 0x18, 0xe3, 0x2c,
	0x33, 0x6f, 0x86, 0xe5, 0xf3, 0x6b, 0x22, 0xfd, 0x9e, 0x07, 0x1e, 0x7d, 0x61, 0x41, 0xf4, 0x01,
	0x58, 0x80, 0x24, 0x45, 0x96, 0x93, 0x6a, 0xdf, 0x81, 0x3e, 0xe9, 0x5e, 0xd1, 0x8e, 0xa0, 0x22,
	0x28, 0xf3, 0x82, 0x99, 0x4f, 0xf3, 0x92, 0x19, 0x8f, 0x0c, 0x58, 0xef, 0x96, 0xfa, 0x7b, 0xe3,
	0x1a, 0xe3, 0x97, 0x94, 0x99, 0x4e, 0x7e, 0x3b, 0xab, 0xf0, 0x03, 0xd7, 0xc9, 0x41, 0x56, 0xe5,
	0x7f, 0x4b, 0xd0, 0xdd, 0x2d, 0x84, 0x57, 0xf9, 0x04, 0x1a, 0x92, 0xf3, 0x79, 0xb4, 0x2e, 0xb3,
	0x01, 0x5e, 0x18, 0x99, 0x3f, 0x83, 0xd6, 0xa6, 0x44, 0xe6, 0x9a, 0x74, 0x31, 0xc7, 0x0f, 0xc5,
	0xba, 0x4a, 0x2e, 0x64, 0x58, 0x49, 0x27, 0xb2, 0x65, 0x5c, 0xb1, 0x8c, 0xc1, 0x43, 0x86, 0xf6,
	0x72, 0x5a, 0xab, 0xdb, 0xa7, 0xb5, 0xb6, 0x36, 0xad, 0x67, 0xff, 0xd4, 0xa0, 0x3e, 0xfa, 0xc9,
	0xcc, 0x22, 0xfa, 0x1a, 0xf6, 0x9f, 0x4b, 0x82, 0x35, 0x71, 0x1e, 0x87, 0xb6, 0x99, 0x5e, 0xf0,
	0x28, 0x74, 0xfe, 0x1b, 0x16, 0xfe, 0x1b, 0x0e, 0x8d, 0xff, 0x9a, 0x62, 0x37, 0x5c, 0x6f, 0x53,
	0xfc, 0x05, 0x34, 0x96, 0xb6, 0x8d, 0xda, 0x85, 0x4f, 0x6c, 0xd8, 0x78, 0xb0, 0xad, 0x21, 0x1a,
	0x02, 0x5c, 0x50, 0x55, 0x54, 0x06, 0xcb, 0x94, 0x15, 0x58, 0x94, 0x9f, 0x6c, 0x8d, 0xf9, 0x2b,
	0x7a, 0x06, 0x07, 0x1b, 0x7f, 0x02, 0xe8, 0xd8, 0x72, 0xd8, 0xf6, 0xc7, 0xb0, 0xf3, 0x0c, 0x23,
	0x38, 0xdc, 0xf4, 0x36, 0x14, 0x6c, 0x35, 0xbc, 0x82, 0xce, 0x6e, 0x33, 0x44, 0x17, 0x80, 0x5e,
	0x37, 0x2a, 0xf4, 0xde, 0x1a, 0xa7, 0x2d, 0x0e, 0xb6, 0x93, 0xd8, 0x6f, 0x70, 0xbc, 0xd3, 0x0a,
	0xd0, 0x69, 0xc1, 0xe3, 0x8d, 0x66, 0x15, 0x7c, 0x74, 0x5f, 0x9a, 0x67, 0x9e, 0x42, 0x67, 0xd7,
	0x3c, 0xa0, 0x0f, 0x6d, 0x8f, 0x7b, 0x7c, 0x23, 0x38, 0xbd, 0x27, 0xcb, 0x6d, 0xf4, 0xec, 0xc9,
	0x2f, 0x9f, 0xa4, 0x54, 0xdf, 0xe4, 0xd7, 0x61, 0xcc, 0xe7, 0x03, 0x7f, 0xb5, 0xcb, 0x5f, 0x31,
	0x4b, 0x07, 0x8a, 0xc8, 0x05, 0x8d, 0xc9, 0x40, 0x8a, 0xf8, 0xba, 0x6e, 0x25, 0xf9, 0xfc, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xfa, 0x32, 0x95, 0x7c, 0x08, 0x00, 0x00,
}
