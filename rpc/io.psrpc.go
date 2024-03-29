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
import livekit3 "github.com/livekit/protocol/livekit"

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
	CreateIngress(ctx context.Context, req *livekit3.IngressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

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
	CreateIngress(context.Context, *livekit3.IngressInfo) (*google_protobuf.Empty, error)

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
	sd.RegisterMethod("CreateIngress", false, false, true, true)
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

func (c *iOInfoClient) CreateIngress(ctx context.Context, req *livekit3.IngressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "CreateIngress", nil, req, opts...)
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

	sd.RegisterMethod("CreateIngress", false, false, true, true)
	err = server.RegisterHandler(s, "CreateIngress", nil, svc.CreateIngress, nil)
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
	// 1004 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x53, 0xe4, 0x44,
	0x14, 0x75, 0x3e, 0x97, 0xb9, 0x03, 0x38, 0xf6, 0x0e, 0x38, 0x84, 0xd2, 0xc5, 0x59, 0x51, 0x6a,
	0xdd, 0xca, 0x94, 0xf8, 0xa0, 0xa5, 0x65, 0x95, 0x30, 0x1b, 0x31, 0x8a, 0x30, 0x06, 0xe6, 0x41,
	0x5f, 0x62, 0x48, 0x9a, 0xd0, 0x92, 0x49, 0xb7, 0xdd, 0x1d, 0x58, 0xde, 0x7c, 0x52, 0xff, 0xaf,
	0xfe, 0x00, 0xab, 0x3f, 0x32, 0x0c, 0xcb, 0xcc, 0x62, 0xed, 0x53, 0xd2, 0xe7, 0x9e, 0x7b, 0x73,
	0xee, 0x49, 0xf7, 0x6d, 0x58, 0xe6, 0x2c, 0x1e, 0x10, 0xea, 0x32, 0x4e, 0x25, 0x45, 0x35, 0xce,
	0x62, 0xa7, 0x9b, 0x91, 0x2b, 0x7c, 0x49, 0x64, 0x88, 0x53, 0x8e, 0x85, 0x30, 0x21, 0x67, 0xad,
	0x44, 0x49, 0x3e, 0x0b, 0x6f, 0xa6, 0x94, 0xa6, 0x19, 0x1e, 0xe8, 0xd5, 0x59, 0x71, 0x3e, 0xc0,
	0x13, 0x26, 0x6f, 0x4c, 0xb0, 0x3f, 0x80, 0xce, 0x01, 0x96, 0x9e, 0xe6, 0x07, 0xf8, 0xf7, 0x02,
	0x0b, 0x89, 0x36, 0xa1, 0x65, 0xea, 0x86, 0x24, 0xe9, 0x55, 0xb6, 0x2a, 0x3b, 0xad, 0x60, 0xc9,
	0x00, 0x7e, 0xd2, 0xff, 0xab, 0x02, 0xdd, 0x31, 0x4b, 0x22, 0x89, 0x7f, 0xc4, 0x92, 0x93, 0x78,
	0x9a, 0xf5, 0x31, 0xd4, 0x49, 0x7e, 0x4e, 0x75, 0x42, 0x7b, 0xf7, 0xb1, 0x6b, 0xc5, 0xb8, 0xa6,
	0xb6, 0x9f, 0x9f, 0xd3, 0x40, 0x13, 0x50, 0x1f, 0x56, 0xa2, 0xab, 0x34, 0x8c, 0x59, 0x11, 0x16,
	0x22, 0x4a, 0x71, 0xaf, 0xb6, 0x55, 0xd9, 0xa9, 0x06, 0xed, 0xe8, 0x2a, 0x1d, 0xb2, 0x62, 0xac,
	0x20, 0xc5, 0x99, 0x44, 0x2f, 0x67, 0x38, 0x75, 0xc3, 0x99, 0x44, 0x2f, 0x4b, 0x4e, 0x7f, 0x0c,
	0x6b, 0x07, 0x58, 0xfa, 0xf9, 0x6d, 0x7d, 0xab, 0xe4, 0x3d, 0x00, 0xeb, 0xc0, 0x6d, 0x03, 0x2d,
	0x8b, 0xf8, 0x89, 0x0a, 0x0b, 0xc9, 0x71, 0x34, 0x09, 0x2f, 0xf1, 0x4d, 0xaf, 0x6a, 0xc2, 0x06,
	0xf9, 0x01, 0xdf, 0xf4, 0xff, 0xae, 0xc2, 0xfa, 0xab, 0x75, 0x05, 0xa3, 0xb9, 0xc0, 0x68, 0xe7,
	0x4e, 0x8b, 0xdd, 0x69, 0x8b, 0xb3, 0x5c, 0xd3, 0x63, 0x17, 0x1a, 0x92, 0x5e, 0xe2, 0xdc, 0x96,
	0x37, 0x0b, 0xb4, 0x06, 0xcd, 0x6b, 0x11, 0x16, 0x3c, 0xd3, 0x2d, 0xb7, 0x82, 0xc6, 0xb5, 0x18,
	0xf3, 0x0c, 0x8d, 0x61, 0x35, 0xa3, 0x69, 0x4a, 0xf2, 0x34, 0x3c, 0x27, 0x38, 0x4b, 0x44, 0xaf,
	0xbe, 0x55, 0xdb, 0x69, 0xef, 0xba, 0x2e, 0x67, 0xb1, 0x3b, 0x5f, 0x8b, 0x7b, 0x68, 0x32, 0xbe,
	0xd5, 0x09, 0x5e, 0x2e, 0xf9, 0x4d, 0xb0, 0x92, 0xcd, 0x62, 0xce, 0x37, 0x80, 0xee, 0x93, 0x50,
	0x07, 0x6a, 0xaa, 0x6d, 0xe3, 0x8a, 0x7a, 0x55, 0x5a, 0xaf, 0xa2, 0xac, 0xc0, 0xa5, 0x56, 0xbd,
	0xf8, 0xb2, 0xfa, 0x45, 0xa5, 0x9f, 0xc2, 0x86, 0xf9, 0xd5, 0x56, 0xc0, 0x89, 0x8c, 0x24, 0xfe,
	0x9f, 0x2e, 0x7f, 0x02, 0x0d, 0xa1, 0xe8, 0xba, 0x6a, 0x7b, 0x77, 0xed, 0x55, 0xb3, 0x4c, 0x2d,
	0xc3, 0xe9, 0xff, 0x51, 0x81, 0xad, 0x03, 0x2c, 0x4f, 0xfc, 0xd1, 0x29, 0x2f, 0xf2, 0xcb, 0xbd,
	0x42, 0x5e, 0xe0, 0x5c, 0x92, 0x38, 0x92, 0x84, 0xe6, 0xe5, 0x07, 0x11, 0xd4, 0xcf, 0x39, 0x9d,
	0x58, 0x99, 0xfa, 0x1d, 0xad, 0x42, 0x55, 0x52, 0xeb, 0x66, 0x55, 0x52, 0xf4, 0x04, 0xda, 0x82,
	0xc7, 0x61, 0x94, 0x24, 0xea, 0x1b, 0x7a, 0xd7, 0xb4, 0x02, 0x10, 0x3c, 0xde, 0x33, 0x08, 0x7a,
	0x17, 0x1e, 0x49, 0x1a, 0x5e, 0x50, 0x21, 0x7b, 0x0d, 0x1d, 0x6c, 0x4a, 0xfa, 0x1d, 0x15, 0xb2,
	0x4f, 0xe1, 0x83, 0xd7, 0x28, 0xb0, 0x1b, 0xc0, 0x81, 0xa5, 0x42, 0x60, 0x9e, 0x47, 0x13, 0x5c,
	0x1e, 0x8c, 0x72, 0xad, 0x62, 0x2c, 0x12, 0xe2, 0x9a, 0xf2, 0xc4, 0x4a, 0x9c, 0xae, 0x95, 0xf4,
	0x84, 0x53, 0xa6, 0x85, 0x2e, 0x05, 0xfa, 0xbd, 0xff, 0x67, 0x15, 0x9e, 0x78, 0xca, 0xeb, 0x48,
	0xe2, 0x13, 0x7f, 0xf4, 0x82, 0x08, 0x16, 0xc9, 0xf8, 0x22, 0x28, 0x32, 0x3c, 0x3d, 0x53, 0xcf,
	0x01, 0x09, 0xc2, 0x42, 0x16, 0x71, 0x49, 0x62, 0xc2, 0xa2, 0x5c, 0xde, 0x7a, 0xdd, 0x11, 0x84,
	0x8d, 0x6e, 0x03, 0x7e, 0x82, 0xb6, 0x61, 0x35, 0x8e, 0xb2, 0x4c, 0xed, 0xa3, 0xbc, 0x98, 0x9c,
	0x61, 0x6e, 0x75, 0xac, 0x58, 0xf4, 0x48, 0x83, 0xe8, 0x29, 0x68, 0x00, 0x27, 0x25, 0xcb, 0xd8,
	0xb7, 0x6c, 0x40, 0x4b, 0x7a, 0xd0, 0xc8, 0x0e, 0xd4, 0x18, 0xc9, 0xad, 0x89, 0xea, 0x55, 0xed,
	0xee, 0x9c, 0x86, 0x0a, 0x6c, 0xea, 0x36, 0x1b, 0x39, 0x1d, 0x91, 0x5c, 0x55, 0xb2, 0x9f, 0xd3,
	0xae, 0x3f, 0x32, 0x95, 0x0c, 0xa4, 0x9d, 0xff, 0xb7, 0x02, 0x5b, 0x8b, 0x8d, 0xb0, 0xce, 0x6f,
	0x42, 0x8b, 0x53, 0x3a, 0x09, 0x67, 0xad, 0x57, 0xc0, 0x91, 0xb2, 0xfe, 0x53, 0xe8, 0xde, 0xb5,
	0x48, 0xfd, 0x3a, 0x59, 0x9e, 0xed, 0xc7, 0x6c, 0xd6, 0x25, 0x13, 0x42, 0x4f, 0xa1, 0xcd, 0x8d,
	0xc9, 0x5a, 0xb1, 0xfe, 0x31, 0xfb, 0xd5, 0x5e, 0x25, 0x00, 0x0b, 0x2b, 0xe9, 0xd3, 0x53, 0x5c,
	0x9f, 0x7f, 0x8a, 0x1b, 0xb3, 0xa7, 0xd8, 0x85, 0x26, 0xc7, 0xa2, 0xc8, 0xa4, 0x6e, 0x7f, 0x75,
	0x77, 0x5d, 0x9f, 0xde, 0xd9, 0x86, 0x74, 0x34, 0xb0, 0xac, 0x67, 0xbf, 0xc2, 0x3b, 0xf7, 0x82,
	0xa8, 0x07, 0xdd, 0x43, 0xef, 0x60, 0x6f, 0xf8, 0x73, 0xb8, 0x37, 0x1c, 0x7a, 0xa3, 0xd3, 0xf0,
	0x38, 0x08, 0x47, 0xfe, 0x51, 0xe7, 0x2d, 0x04, 0xd0, 0x34, 0x50, 0xa7, 0x82, 0xde, 0x86, 0x76,
	0xe0, 0xfd, 0x34, 0xf6, 0x4e, 0x4e, 0x75, 0xb0, 0xaa, 0x82, 0x81, 0xf7, 0xbd, 0x37, 0x3c, 0xed,
	0xd4, 0xd0, 0x12, 0xd4, 0x5f, 0x04, 0xc7, 0xa3, 0x4e, 0x7d, 0xf7, 0x9f, 0x06, 0x34, 0xfd, 0x63,
	0x35, 0x35, 0xd0, 0x57, 0xb0, 0x3c, 0xe4, 0x38, 0x92, 0xd8, 0x4c, 0x63, 0x34, 0x6f, 0x3c, 0x3b,
	0xeb, 0xae, 0xb9, 0x29, 0xdc, 0xf2, 0xa6, 0x70, 0x3d, 0x75, 0x53, 0xa8, 0x64, 0x33, 0x06, 0xde,
	0x24, 0xf9, 0x73, 0x68, 0x4d, 0x2f, 0x18, 0xb4, 0x56, 0x4e, 0xb4, 0x3b, 0x17, 0x8e, 0x33, 0xaf,
	0x20, 0xf2, 0x00, 0x0e, 0x89, 0x28, 0x33, 0x9d, 0x29, 0xe5, 0x16, 0x2c, 0xd3, 0x37, 0xe7, 0xc6,
	0xec, 0xc6, 0xd9, 0x87, 0x95, 0x3b, 0xd7, 0x15, 0xda, 0xd0, 0x1a, 0xe6, 0x5d, 0x61, 0x0b, 0x7b,
	0xf8, 0x1a, 0x56, 0x8c, 0x7b, 0x76, 0x76, 0xa1, 0xb9, 0xa3, 0x7f, 0x61, 0xba, 0x0f, 0xab, 0x77,
	0x87, 0x38, 0x72, 0xe6, 0x4e, 0xf6, 0xb2, 0x9b, 0xc5, 0x53, 0x1f, 0x1d, 0x02, 0xba, 0x3f, 0x91,
	0xd1, 0xfb, 0x33, 0x2d, 0xcd, 0x19, 0xd5, 0x0b, 0x85, 0xfd, 0x06, 0x1b, 0x0b, 0x67, 0x1e, 0xda,
	0x2e, 0x75, 0xbc, 0x76, 0x2a, 0x3b, 0x1f, 0x3d, 0x44, 0xb3, 0xca, 0x53, 0xe8, 0x2d, 0x3a, 0xe4,
	0xe8, 0x43, 0x5d, 0xe3, 0x81, 0x61, 0xe8, 0x6c, 0x3f, 0xc0, 0x32, 0x1f, 0xda, 0x7f, 0xfe, 0xcb,
	0xb3, 0x94, 0xc8, 0x8b, 0xe2, 0xcc, 0x8d, 0xe9, 0x64, 0x60, 0xff, 0xd3, 0xf4, 0xc9, 0x2e, 0xd3,
	0x81, 0xc0, 0xfc, 0x8a, 0xc4, 0x78, 0xc0, 0x59, 0x7c, 0xd6, 0xd4, 0x96, 0x7c, 0xf6, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x67, 0xb7, 0xc7, 0x3b, 0x65, 0x09, 0x00, 0x00,
}
