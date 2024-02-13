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
	// 992 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x57, 0xe3, 0x44,
	0x14, 0xb6, 0x3f, 0x97, 0xde, 0x2e, 0x58, 0x67, 0x0b, 0x96, 0x70, 0x74, 0xb1, 0x2b, 0xca, 0x59,
	0xf7, 0xa4, 0x47, 0x7c, 0xd0, 0xa3, 0x2f, 0x42, 0x37, 0x62, 0x14, 0xa1, 0x06, 0xfa, 0xa0, 0x2f,
	0x31, 0x24, 0x43, 0x18, 0x49, 0x33, 0xe3, 0xcc, 0x04, 0x96, 0x37, 0x9f, 0xd4, 0x3f, 0xc8, 0x3f,
	0xcd, 0x3f, 0xc0, 0x33, 0x3f, 0x52, 0xda, 0xa5, 0x5d, 0x3c, 0x3e, 0x25, 0xf3, 0xdd, 0x6f, 0x6e,
	0xbe, 0xfb, 0x65, 0xee, 0x1d, 0x78, 0xcc, 0x59, 0x3c, 0x20, 0xd4, 0x65, 0x9c, 0x4a, 0x8a, 0x6a,
	0x9c, 0xc5, 0x4e, 0x37, 0x23, 0xd7, 0xf8, 0x8a, 0xc8, 0x10, 0xa7, 0x1c, 0x0b, 0x61, 0x42, 0xce,
	0x7a, 0x89, 0x92, 0x7c, 0x16, 0xde, 0x4a, 0x29, 0x4d, 0x33, 0x3c, 0xd0, 0xab, 0xf3, 0xe2, 0x62,
	0x80, 0x27, 0x4c, 0xde, 0x9a, 0x60, 0x7f, 0x00, 0x9d, 0x43, 0x2c, 0x3d, 0xcd, 0x0f, 0xf0, 0x6f,
	0x05, 0x16, 0x12, 0x6d, 0x41, 0xcb, 0xe4, 0x0d, 0x49, 0xd2, 0xab, 0x6c, 0x57, 0x76, 0x5b, 0xc1,
	0x8a, 0x01, 0xfc, 0xa4, 0xff, 0x67, 0x05, 0xba, 0x63, 0x96, 0x44, 0x12, 0xff, 0x80, 0x25, 0x27,
	0xf1, 0x74, 0xd7, 0xc7, 0x50, 0x27, 0xf9, 0x05, 0xd5, 0x1b, 0xda, 0x7b, 0x4f, 0x5c, 0x2b, 0xc6,
	0x35, 0xb9, 0xfd, 0xfc, 0x82, 0x06, 0x9a, 0x80, 0xfa, 0xb0, 0x1a, 0x5d, 0xa7, 0x61, 0xcc, 0x8a,
	0xb0, 0x10, 0x51, 0x8a, 0x7b, 0xb5, 0xed, 0xca, 0x6e, 0x35, 0x68, 0x47, 0xd7, 0xe9, 0x90, 0x15,
	0x63, 0x05, 0x29, 0xce, 0x24, 0x7a, 0x35, 0xc3, 0xa9, 0x1b, 0xce, 0x24, 0x7a, 0x55, 0x72, 0xfa,
	0x63, 0x58, 0x3f, 0xc4, 0xd2, 0xcf, 0xef, 0xf2, 0x5b, 0x25, 0xef, 0x01, 0x58, 0x07, 0xee, 0x0a,
	0x68, 0x59, 0xc4, 0x4f, 0x54, 0x58, 0x48, 0x8e, 0xa3, 0x49, 0x78, 0x85, 0x6f, 0x7b, 0x55, 0x13,
	0x36, 0xc8, 0xf7, 0xf8, 0xb6, 0xff, 0x57, 0x15, 0x36, 0x5e, 0xcf, 0x2b, 0x18, 0xcd, 0x05, 0x46,
	0xbb, 0x73, 0x25, 0x76, 0xa7, 0x25, 0xce, 0x72, 0x4d, 0x8d, 0x5d, 0x68, 0x48, 0x7a, 0x85, 0x73,
	0x9b, 0xde, 0x2c, 0xd0, 0x3a, 0x34, 0x6f, 0x44, 0x58, 0xf0, 0x4c, 0x97, 0xdc, 0x0a, 0x1a, 0x37,
	0x62, 0xcc, 0x33, 0x34, 0x86, 0xb5, 0x8c, 0xa6, 0x29, 0xc9, 0xd3, 0xf0, 0x82, 0xe0, 0x2c, 0x11,
	0xbd, 0xfa, 0x76, 0x6d, 0xb7, 0xbd, 0xe7, 0xba, 0x9c, 0xc5, 0xee, 0x62, 0x2d, 0xee, 0x91, 0xd9,
	0xf1, 0x8d, 0xde, 0xe0, 0xe5, 0x92, 0xdf, 0x06, 0xab, 0xd9, 0x2c, 0xe6, 0x7c, 0x0d, 0xe8, 0x3e,
	0x09, 0x75, 0xa0, 0xa6, 0xca, 0x36, 0xae, 0xa8, 0x57, 0xa5, 0xf5, 0x3a, 0xca, 0x0a, 0x5c, 0x6a,
	0xd5, 0x8b, 0x2f, 0xab, 0x5f, 0x54, 0xfa, 0x29, 0x6c, 0x9a, 0x5f, 0x6d, 0x05, 0x9c, 0xca, 0x48,
	0xe2, 0xff, 0xe8, 0xf2, 0x27, 0xd0, 0x10, 0x8a, 0xae, 0xb3, 0xb6, 0xf7, 0xd6, 0x5f, 0x37, 0xcb,
	0xe4, 0x32, 0x9c, 0xfe, 0xef, 0x15, 0xd8, 0x3e, 0xc4, 0xf2, 0xd4, 0x1f, 0x9d, 0xf1, 0x22, 0xbf,
	0xda, 0x2f, 0xe4, 0x25, 0xce, 0x25, 0x89, 0x23, 0x49, 0x68, 0x5e, 0x7e, 0x10, 0x41, 0xfd, 0x82,
	0xd3, 0x89, 0x95, 0xa9, 0xdf, 0xd1, 0x1a, 0x54, 0x25, 0xb5, 0x6e, 0x56, 0x25, 0x45, 0x4f, 0xa1,
	0x2d, 0x78, 0x1c, 0x46, 0x49, 0xa2, 0xbe, 0xa1, 0x4f, 0x4d, 0x2b, 0x00, 0xc1, 0xe3, 0x7d, 0x83,
	0xa0, 0x77, 0xe1, 0x91, 0xa4, 0xe1, 0x25, 0x15, 0xb2, 0xd7, 0xd0, 0xc1, 0xa6, 0xa4, 0xdf, 0x52,
	0x21, 0xfb, 0x14, 0x3e, 0x78, 0x83, 0x02, 0x7b, 0x00, 0x1c, 0x58, 0x29, 0x04, 0xe6, 0x79, 0x34,
	0xc1, 0x65, 0x63, 0x94, 0x6b, 0x15, 0x63, 0x91, 0x10, 0x37, 0x94, 0x27, 0x56, 0xe2, 0x74, 0xad,
	0xa4, 0x27, 0x9c, 0x32, 0x2d, 0x74, 0x25, 0xd0, 0xef, 0xfd, 0x3f, 0xaa, 0xf0, 0xd4, 0x53, 0x5e,
	0x47, 0x12, 0x9f, 0xfa, 0xa3, 0x97, 0x44, 0xb0, 0x48, 0xc6, 0x97, 0x41, 0x91, 0xe1, 0x69, 0x4f,
	0xbd, 0x00, 0x24, 0x08, 0x0b, 0x59, 0xc4, 0x25, 0x89, 0x09, 0x8b, 0x72, 0x79, 0xe7, 0x75, 0x47,
	0x10, 0x36, 0xba, 0x0b, 0xf8, 0x09, 0xda, 0x81, 0xb5, 0x38, 0xca, 0x32, 0x75, 0x8e, 0xf2, 0x62,
	0x72, 0x8e, 0xb9, 0xd5, 0xb1, 0x6a, 0xd1, 0x63, 0x0d, 0xa2, 0x67, 0xa0, 0x01, 0x9c, 0x94, 0x2c,
	0x63, 0xdf, 0x63, 0x03, 0x5a, 0xd2, 0x83, 0x46, 0x76, 0xa0, 0xc6, 0x48, 0x6e, 0x4d, 0x54, 0xaf,
	0xea, 0x74, 0xe7, 0x34, 0x54, 0x60, 0x53, 0x97, 0xd9, 0xc8, 0xe9, 0x88, 0xe4, 0x2a, 0x93, 0xfd,
	0x9c, 0x76, 0xfd, 0x91, 0xc9, 0x64, 0x20, 0xed, 0xfc, 0x3f, 0x15, 0xd8, 0x5e, 0x6e, 0x84, 0x75,
	0x7e, 0x0b, 0x5a, 0x9c, 0xd2, 0x49, 0x38, 0x6b, 0xbd, 0x02, 0x8e, 0x95, 0xf5, 0x9f, 0x42, 0x77,
	0xde, 0x22, 0xf5, 0xeb, 0x64, 0xd9, 0xdb, 0x4f, 0xd8, 0xac, 0x4b, 0x26, 0x84, 0x9e, 0x41, 0x9b,
	0x1b, 0x93, 0xb5, 0x62, 0xfd, 0x63, 0x0e, 0xaa, 0xbd, 0x4a, 0x00, 0x16, 0x56, 0xd2, 0xa7, 0x5d,
	0x5c, 0x5f, 0xdc, 0xc5, 0x8d, 0xd9, 0x2e, 0x76, 0xa1, 0xc9, 0xb1, 0x28, 0x32, 0xa9, 0xcb, 0x5f,
	0xdb, 0xdb, 0xd0, 0xdd, 0x3b, 0x5b, 0x90, 0x8e, 0x06, 0x96, 0xf5, 0xfc, 0x17, 0x78, 0xe7, 0x5e,
	0x10, 0xf5, 0xa0, 0x7b, 0xe4, 0x1d, 0xee, 0x0f, 0x7f, 0x0a, 0xf7, 0x87, 0x43, 0x6f, 0x74, 0x16,
	0x9e, 0x04, 0xe1, 0xc8, 0x3f, 0xee, 0xbc, 0x85, 0x00, 0x9a, 0x06, 0xea, 0x54, 0xd0, 0xdb, 0xd0,
	0x0e, 0xbc, 0x1f, 0xc7, 0xde, 0xe9, 0x99, 0x0e, 0x56, 0x55, 0x30, 0xf0, 0xbe, 0xf3, 0x86, 0x67,
	0x9d, 0x1a, 0x5a, 0x81, 0xfa, 0xcb, 0xe0, 0x64, 0xd4, 0xa9, 0xef, 0xfd, 0xdd, 0x80, 0xa6, 0x7f,
	0xa2, 0xa6, 0x06, 0xfa, 0x0a, 0x1e, 0x0f, 0x39, 0x8e, 0x24, 0x36, 0xd3, 0x18, 0x2d, 0x1a, 0xcf,
	0xce, 0x86, 0x6b, 0x6e, 0x0a, 0xb7, 0xbc, 0x29, 0x5c, 0x4f, 0xdd, 0x14, 0x6a, 0xb3, 0x19, 0x03,
	0xff, 0x67, 0xf3, 0xe7, 0xd0, 0x9a, 0x5e, 0x30, 0x68, 0xbd, 0x9c, 0x68, 0x73, 0x17, 0x8e, 0xb3,
	0x28, 0x21, 0xf2, 0x00, 0x8e, 0x88, 0x28, 0x77, 0x3a, 0x53, 0xca, 0x1d, 0x58, 0x6e, 0xdf, 0x5a,
	0x18, 0xb3, 0x07, 0xe7, 0x00, 0x56, 0xe7, 0xae, 0x2b, 0xb4, 0xa9, 0x35, 0x2c, 0xba, 0xc2, 0x96,
	0xd6, 0xe0, 0xc3, 0xda, 0xfc, 0x14, 0x46, 0xce, 0xc2, 0xd1, 0x5c, 0xca, 0x59, 0x3e, 0xb6, 0xd1,
	0x11, 0xa0, 0xfb, 0x23, 0x15, 0xbd, 0x3f, 0xa3, 0x69, 0xc1, 0xac, 0x5d, 0x2a, 0xec, 0x57, 0xd8,
	0x5c, 0x3a, 0xb4, 0xd0, 0x4e, 0xa9, 0xe3, 0x8d, 0x63, 0xd5, 0xf9, 0xe8, 0x21, 0x9a, 0x55, 0x9e,
	0x42, 0x6f, 0x59, 0x97, 0xa2, 0x0f, 0x75, 0x8e, 0x07, 0xa6, 0x99, 0xb3, 0xf3, 0x00, 0xcb, 0x7c,
	0xe8, 0xe0, 0xc5, 0xcf, 0xcf, 0x53, 0x22, 0x2f, 0x8b, 0x73, 0x37, 0xa6, 0x93, 0x81, 0xfd, 0xb5,
	0xd3, 0x27, 0xbb, 0x4a, 0x07, 0x02, 0xf3, 0x6b, 0x12, 0xe3, 0x01, 0x67, 0xf1, 0x79, 0x53, 0x5b,
	0xf2, 0xd9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x60, 0x84, 0xf9, 0x26, 0x09, 0x00, 0x00,
}
