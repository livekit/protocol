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
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x73, 0xe3, 0x34,
	0x14, 0x9d, 0xa4, 0x4d, 0xb6, 0xb9, 0xfd, 0x98, 0x8e, 0x36, 0x59, 0x52, 0x67, 0x60, 0x83, 0xa1,
	0xd0, 0x81, 0x1d, 0x67, 0x28, 0x0f, 0x3c, 0xf0, 0xb4, 0xbb, 0x74, 0x8a, 0x87, 0x65, 0xe9, 0xa4,
	0xf4, 0x05, 0x1e, 0x3c, 0xaa, 0xad, 0xba, 0x22, 0xb6, 0x24, 0x24, 0x39, 0xdd, 0xbe, 0xf1, 0x04,
	0x7f, 0x88, 0x1f, 0xc0, 0x4f, 0x63, 0xf4, 0x61, 0x27, 0xed, 0x26, 0x2d, 0xc3, 0x53, 0xa2, 0x73,
	0x8f, 0xae, 0xce, 0x3d, 0xb2, 0x0e, 0xec, 0x48, 0x91, 0x4e, 0x28, 0x8f, 0x84, 0xe4, 0x9a, 0xa3,
	0x0d, 0x29, 0xd2, 0xa0, 0x5f, 0xd0, 0x39, 0x99, 0x51, 0x9d, 0x90, 0x5c, 0x12, 0xa5, 0x5c, 0x29,
	0x18, 0xd4, 0x28, 0x65, 0xcb, 0xf0, 0x28, 0xe7, 0x3c, 0x2f, 0xc8, 0xc4, 0xae, 0x2e, 0xab, 0xab,
	0x09, 0x29, 0x85, 0xbe, 0x75, 0xc5, 0x70, 0x02, 0xfb, 0xa7, 0x44, 0x9f, 0x58, 0xfe, 0x94, 0xfc,
	0x5e, 0x11, 0xa5, 0xd1, 0x08, 0x7a, 0xae, 0x6f, 0x42, 0xb3, 0x61, 0x6b, 0xdc, 0x3a, 0xea, 0x4d,
	0xb7, 0x1c, 0x10, 0x67, 0xe1, 0x5f, 0x2d, 0xe8, 0x5f, 0x88, 0x0c, 0x6b, 0xf2, 0x23, 0xd1, 0x92,
	0xa6, 0xcd, 0xae, 0xcf, 0x61, 0x93, 0xb2, 0x2b, 0x6e, 0x37, 0x6c, 0x1f, 0x3f, 0x8d, 0xbc, 0x98,
	0xc8, 0xf5, 0x8e, 0xd9, 0x15, 0x9f, 0x5a, 0x02, 0x0a, 0x61, 0x17, 0xcf, 0xf3, 0x24, 0x15, 0x55,
	0x52, 0x29, 0x9c, 0x93, 0xe1, 0xc6, 0xb8, 0x75, 0xd4, 0x9e, 0x6e, 0xe3, 0x79, 0xfe, 0x5a, 0x54,
	0x17, 0x06, 0x32, 0x9c, 0x12, 0xbf, 0x5b, 0xe2, 0x6c, 0x3a, 0x4e, 0x89, 0xdf, 0xd5, 0x9c, 0xf0,
	0x02, 0x06, 0xa7, 0x44, 0xc7, 0x6c, 0xd1, 0xdf, 0x2b, 0xf9, 0x10, 0xc0, 0x3b, 0xb0, 0x18, 0xa0,
	0xe7, 0x91, 0x38, 0x33, 0x65, 0xa5, 0x25, 0xc1, 0x65, 0x32, 0x23, 0xb7, 0xc3, 0xb6, 0x2b, 0x3b,
	0xe4, 0x07, 0x72, 0x1b, 0x72, 0x78, 0x76, 0xbf, 0xad, 0x12, 0x9c, 0x29, 0x82, 0x8e, 0xee, 0x4c,
	0xd8, 0x6f, 0x26, 0x5c, 0xe6, 0xba, 0x11, 0xfb, 0xd0, 0xd1, 0x7c, 0x46, 0x98, 0xef, 0xee, 0x16,
	0x68, 0x00, 0xdd, 0x1b, 0x95, 0x54, 0xb2, 0xb0, 0x13, 0xf7, 0xa6, 0x9d, 0x1b, 0x75, 0x21, 0x8b,
	0x30, 0x87, 0x03, 0x67, 0xa8, 0xef, 0x73, 0xae, 0xb1, 0x26, 0xff, 0x71, 0x96, 0x2f, 0xa1, 0xa3,
	0x0c, 0xdd, 0x1e, 0xb4, 0x7d, 0x3c, 0xb8, 0xaf, 0xc9, 0xf5, 0x72, 0x9c, 0xf0, 0x8f, 0x16, 0x8c,
	0x4f, 0x89, 0x3e, 0x8f, 0xcf, 0x7e, 0x96, 0x15, 0x9b, 0xbd, 0xac, 0xf4, 0x35, 0x61, 0x9a, 0xa6,
	0x58, 0x53, 0xce, 0xea, 0x03, 0x11, 0x6c, 0x5e, 0x49, 0x5e, 0x7a, 0xe5, 0xf6, 0x3f, 0xda, 0x83,
	0xb6, 0xe6, 0x5e, 0x74, 0x5b, 0x73, 0xf4, 0x1c, 0xb6, 0x95, 0x4c, 0x13, 0x9c, 0x65, 0xe6, 0x0c,
	0x7b, 0x37, 0xbd, 0x29, 0x28, 0x99, 0xbe, 0x74, 0x08, 0xfa, 0x00, 0x9e, 0x68, 0x9e, 0x5c, 0x73,
	0xa5, 0x87, 0x1d, 0x5b, 0xec, 0x6a, 0xfe, 0x3d, 0x57, 0x3a, 0xfc, 0x15, 0x3e, 0x7e, 0x40, 0x81,
	0xf7, 0x39, 0x80, 0xad, 0x4a, 0x11, 0xc9, 0x70, 0x49, 0xea, 0xcf, 0xaf, 0x5e, 0x9b, 0x9a, 0xc0,
	0x4a, 0xdd, 0x70, 0x99, 0x79, 0x89, 0xcd, 0x3a, 0xfc, 0xb3, 0x0d, 0xcf, 0x4f, 0xe6, 0xb8, 0xa8,
	0xb0, 0x26, 0xe7, 0xf1, 0xd9, 0x77, 0x54, 0x09, 0xac, 0xd3, 0xeb, 0x69, 0x55, 0x90, 0xe6, 0x2b,
	0x7d, 0x01, 0x48, 0x51, 0x91, 0x08, 0x2c, 0x35, 0x4d, 0xa9, 0xc0, 0x4c, 0x2f, 0x7c, 0xdd, 0x57,
	0x54, 0x9c, 0x2d, 0x0a, 0x71, 0x86, 0x0e, 0x61, 0x2f, 0xc5, 0x45, 0x41, 0x59, 0x9e, 0xb0, 0xaa,
	0xbc, 0x24, 0xd2, 0x9f, 0xb9, 0xeb, 0xd1, 0xb7, 0x16, 0x44, 0x9f, 0x80, 0x05, 0x48, 0x56, 0xb3,
	0x9c, 0x55, 0x3b, 0x0e, 0xf4, 0xa4, 0x47, 0x4d, 0xdb, 0x87, 0x0d, 0x41, 0x99, 0x37, 0xcc, 0xfc,
	0x35, 0x1f, 0x0c, 0xe3, 0x89, 0x01, 0xbb, 0xe3, 0xd6, 0xd1, 0xd6, 0xb4, 0xc3, 0xf8, 0x19, 0x65,
	0xa6, 0x93, 0x3f, 0xce, 0x3a, 0xfc, 0xc4, 0x75, 0x72, 0x90, 0x75, 0xf9, 0x9f, 0x16, 0x8c, 0xd7,
	0x1b, 0xe1, 0x5d, 0x1e, 0x41, 0x4f, 0x72, 0x5e, 0x26, 0xcb, 0x36, 0x1b, 0xe0, 0xad, 0xb1, 0xf9,
	0x2b, 0xe8, 0xdf, 0xb5, 0xc8, 0x5c, 0x93, 0xae, 0x5f, 0xcb, 0x53, 0xb1, 0xec, 0x92, 0x2b, 0x19,
	0x55, 0xd2, 0x99, 0x6c, 0x15, 0x6f, 0x58, 0xc5, 0xe0, 0x21, 0x23, 0xbb, 0x79, 0x14, 0x9b, 0xab,
	0x1f, 0x45, 0x67, 0xe9, 0x51, 0x1c, 0xff, 0xdd, 0x81, 0x6e, 0xfc, 0x93, 0x79, 0x52, 0xe8, 0x5b,
	0xd8, 0x79, 0x2d, 0x09, 0xd6, 0xc4, 0x25, 0x09, 0x5a, 0x15, 0x2d, 0xc1, 0xb3, 0xc8, 0xa5, 0x5c,
	0x54, 0xa7, 0x5c, 0x74, 0x62, 0x52, 0xce, 0x6c, 0x76, 0x8f, 0xeb, 0xff, 0x6c, 0xfe, 0x06, 0x7a,
	0x4d, 0x38, 0xa2, 0x41, 0x24, 0x45, 0x1a, 0xdd, 0x0f, 0xcb, 0x60, 0x55, 0x43, 0x74, 0x02, 0xf0,
	0x86, 0xaa, 0x7a, 0x67, 0xd0, 0x50, 0x16, 0x60, 0xbd, 0x7d, 0xb4, 0xb2, 0xe6, 0xaf, 0xe8, 0x15,
	0xec, 0xde, 0x89, 0x5a, 0x74, 0x60, 0x35, 0xac, 0x8a, 0xdf, 0xb5, 0x33, 0xc4, 0xb0, 0x77, 0x37,
	0xce, 0x50, 0x50, 0x0f, 0xf2, 0x7e, 0x74, 0x06, 0xa3, 0x95, 0x35, 0x2f, 0xe7, 0x0d, 0xa0, 0xf7,
	0x83, 0x0a, 0x7d, 0xb4, 0xa4, 0x69, 0x45, 0x82, 0xad, 0x15, 0xf6, 0x1b, 0x1c, 0xac, 0x8d, 0x02,
	0x74, 0x58, 0xeb, 0x78, 0x30, 0xac, 0x82, 0xcf, 0x1e, 0xa3, 0x79, 0xe5, 0x39, 0x0c, 0xd7, 0xbd,
	0x07, 0xf4, 0xa9, 0xed, 0xf1, 0x48, 0x6e, 0x04, 0x87, 0x8f, 0xb0, 0xdc, 0x41, 0xaf, 0x5e, 0xfc,
	0xf2, 0x45, 0x4e, 0xf5, 0x75, 0x75, 0x19, 0xa5, 0xbc, 0x9c, 0xf8, 0xab, 0x6d, 0x7e, 0xc5, 0x2c,
	0x9f, 0x28, 0x22, 0xe7, 0x34, 0x25, 0x13, 0x29, 0xd2, 0xcb, 0xae, 0xb5, 0xe4, 0xeb, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x29, 0xd5, 0xc5, 0x46, 0xe2, 0x07, 0x00, 0x00,
}
