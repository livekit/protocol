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
import livekit4 "github.com/livekit/protocol/livekit"
import livekit5 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_5

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	// egress
	CreateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	UpdateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit4.EgressInfo, error)

	ListEgress(ctx context.Context, req *livekit4.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit4.ListEgressResponse, error)

	UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// ingress
	CreateIngress(ctx context.Context, req *livekit5.IngressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

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
	CreateEgress(context.Context, *livekit4.EgressInfo) (*google_protobuf.Empty, error)

	UpdateEgress(context.Context, *livekit4.EgressInfo) (*google_protobuf.Empty, error)

	GetEgress(context.Context, *GetEgressRequest) (*livekit4.EgressInfo, error)

	ListEgress(context.Context, *livekit4.ListEgressRequest) (*livekit4.ListEgressResponse, error)

	UpdateMetrics(context.Context, *UpdateMetricsRequest) (*google_protobuf.Empty, error)

	// ingress
	CreateIngress(context.Context, *livekit5.IngressInfo) (*google_protobuf.Empty, error)

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

func (c *iOInfoClient) CreateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "CreateEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit4.EgressInfo, error) {
	return client.RequestSingle[*livekit4.EgressInfo](ctx, c.client, "GetEgress", nil, req, opts...)
}

func (c *iOInfoClient) ListEgress(ctx context.Context, req *livekit4.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit4.ListEgressResponse, error) {
	return client.RequestSingle[*livekit4.ListEgressResponse](ctx, c.client, "ListEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateMetrics", nil, req, opts...)
}

func (c *iOInfoClient) CreateIngress(ctx context.Context, req *livekit5.IngressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
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

var psrpcFileDescriptor4 = []byte{
	// 1178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x72, 0xdb, 0x44,
	0x18, 0xc6, 0xc7, 0xc6, 0xbf, 0x93, 0xd4, 0xdd, 0x3a, 0xc1, 0x55, 0x86, 0x36, 0x75, 0x29, 0x04,
	0x98, 0x91, 0x21, 0x5c, 0x70, 0x1a, 0x66, 0x9a, 0xb8, 0x26, 0x15, 0xa4, 0xad, 0x51, 0x92, 0x0b,
	0xb8, 0x11, 0x1b, 0x69, 0xe3, 0x2c, 0x91, 0xb5, 0x62, 0x77, 0x95, 0x34, 0x4f, 0x00, 0xb7, 0xbc,
	0x01, 0x0f, 0xc5, 0x43, 0x70, 0xcf, 0x0b, 0x30, 0x7b, 0x90, 0xa3, 0x24, 0x72, 0xd3, 0x72, 0xa7,
	0xfd, 0xfe, 0x83, 0xbe, 0xff, 0xb4, 0xfb, 0xc3, 0x22, 0x4f, 0xc3, 0x01, 0x65, 0x6e, 0xca, 0x99,
	0x64, 0xa8, 0xc6, 0xd3, 0xd0, 0xe9, 0xc6, 0xf4, 0x94, 0x9c, 0x50, 0x19, 0x90, 0x09, 0x27, 0x42,
	0x18, 0x91, 0xb3, 0x92, 0xa3, 0x34, 0x29, 0xc2, 0x6b, 0x13, 0xc6, 0x26, 0x31, 0x19, 0xe8, 0xd3,
	0x61, 0x76, 0x34, 0x20, 0xd3, 0x54, 0x9e, 0x1b, 0x61, 0x7f, 0x00, 0x9d, 0x1d, 0x22, 0x47, 0x5a,
	0xdf, 0x27, 0xbf, 0x65, 0x44, 0x48, 0xb4, 0x06, 0x2d, 0xe3, 0x37, 0xa0, 0x51, 0xaf, 0xb2, 0x5e,
	0xd9, 0x68, 0xf9, 0x0b, 0x06, 0xf0, 0xa2, 0xfe, 0xef, 0x15, 0xe8, 0x1e, 0xa4, 0x11, 0x96, 0xe4,
	0x39, 0x91, 0x9c, 0x86, 0x33, 0xab, 0x0f, 0xa1, 0x4e, 0x93, 0x23, 0xa6, 0x0d, 0xda, 0x9b, 0x77,
	0x5d, 0x4b, 0xc6, 0x35, 0xbe, 0xbd, 0xe4, 0x88, 0xf9, 0x5a, 0x01, 0xf5, 0x61, 0x09, 0x9f, 0x4e,
	0x82, 0x30, 0xcd, 0x82, 0x4c, 0xe0, 0x09, 0xe9, 0xd5, 0xd6, 0x2b, 0x1b, 0x55, 0xbf, 0x8d, 0x4f,
	0x27, 0xc3, 0x34, 0x3b, 0x50, 0x90, 0xd2, 0x99, 0xe2, 0x57, 0x05, 0x9d, 0xba, 0xd1, 0x99, 0xe2,
	0x57, 0xb9, 0x4e, 0xff, 0x00, 0x56, 0x76, 0x88, 0xf4, 0x92, 0x0b, 0xff, 0x96, 0xc9, 0x7b, 0x00,
	0x36, 0x03, 0x17, 0x01, 0xb4, 0x2c, 0xe2, 0x45, 0x4a, 0x2c, 0x24, 0x27, 0x78, 0x1a, 0x9c, 0x90,
	0xf3, 0x5e, 0xd5, 0x88, 0x0d, 0xf2, 0x03, 0x39, 0xef, 0xff, 0x51, 0x85, 0xd5, 0xab, 0x7e, 0x45,
	0xca, 0x12, 0x41, 0xd0, 0xc6, 0xa5, 0x10, 0xbb, 0xb3, 0x10, 0x8b, 0xba, 0x26, 0xc6, 0x2e, 0x34,
	0x24, 0x3b, 0x21, 0x89, 0x75, 0x6f, 0x0e, 0x68, 0x05, 0x9a, 0x67, 0x22, 0xc8, 0x78, 0xac, 0x43,
	0x6e, 0xf9, 0x8d, 0x33, 0x71, 0xc0, 0x63, 0x74, 0x00, 0xcb, 0x31, 0x9b, 0x4c, 0x68, 0x32, 0x09,
	0x8e, 0x28, 0x89, 0x23, 0xd1, 0xab, 0xaf, 0xd7, 0x36, 0xda, 0x9b, 0xae, 0xcb, 0xd3, 0xd0, 0x2d,
	0xe7, 0xe2, 0xee, 0x1a, 0x8b, 0xef, 0xb4, 0xc1, 0x28, 0x91, 0xfc, 0xdc, 0x5f, 0x8a, 0x8b, 0x98,
	0xf3, 0x04, 0xd0, 0x75, 0x25, 0xd4, 0x81, 0x9a, 0x0a, 0xdb, 0x64, 0x45, 0x7d, 0x2a, 0xae, 0xa7,
	0x38, 0xce, 0x48, 0xce, 0x55, 0x1f, 0xbe, 0xae, 0x7e, 0x59, 0xe9, 0x4f, 0xe0, 0x9e, 0x29, 0xb5,
	0x25, 0xb0, 0x27, 0xb1, 0x24, 0x6f, 0x98, 0xe5, 0x4f, 0xa0, 0x21, 0x94, 0xba, 0xf6, 0xda, 0xde,
	0x5c, 0xb9, 0x9a, 0x2c, 0xe3, 0xcb, 0xe8, 0xf4, 0xff, 0xaa, 0xc0, 0xfa, 0x0e, 0x91, 0x7b, 0xde,
	0x78, 0x9f, 0x67, 0xc9, 0xc9, 0x56, 0x26, 0x8f, 0x49, 0x22, 0x69, 0x88, 0x25, 0x65, 0x49, 0xfe,
	0xc3, 0xfb, 0xd0, 0x16, 0x34, 0x0d, 0x42, 0x1c, 0xc7, 0xea, 0x8f, 0x4d, 0x5b, 0x38, 0x9a, 0x0e,
	0x71, 0x1c, 0x7b, 0x11, 0x42, 0x50, 0x3f, 0xe2, 0x6c, 0x6a, 0xc3, 0xd0, 0xdf, 0x68, 0x19, 0xaa,
	0x92, 0xd9, 0x6c, 0x57, 0x25, 0x43, 0x0f, 0xa0, 0x2d, 0x78, 0x18, 0xe0, 0x28, 0x52, 0x1c, 0x74,
	0x57, 0xb5, 0x7c, 0x10, 0x3c, 0xdc, 0x32, 0x08, 0x7a, 0x17, 0x6e, 0x49, 0x16, 0x1c, 0x33, 0x21,
	0x7b, 0x0d, 0x2d, 0x6c, 0x4a, 0xf6, 0x8c, 0x09, 0xd9, 0xff, 0xb3, 0x02, 0x0f, 0x5f, 0x43, 0xd1,
	0x76, 0x88, 0x03, 0x0b, 0x99, 0x20, 0x3c, 0xc1, 0x53, 0x92, 0x4f, 0x4e, 0x7e, 0x56, 0xb2, 0x14,
	0x0b, 0x71, 0xc6, 0x78, 0x64, 0x39, 0xce, 0xce, 0x8a, 0x7b, 0xc4, 0x59, 0xaa, 0x99, 0x2e, 0xf8,
	0xfa, 0x1b, 0xad, 0xc3, 0xa2, 0x8a, 0x57, 0xaa, 0xdf, 0xa9, 0x80, 0x73, 0xb2, 0x34, 0xd5, 0x0c,
	0xbc, 0xa8, 0xff, 0x77, 0x0d, 0x1e, 0x8c, 0x54, 0xb9, 0xb0, 0x24, 0x7b, 0xde, 0xf8, 0x29, 0x15,
	0x29, 0x96, 0xe1, 0xb1, 0x9f, 0xc5, 0x44, 0xcc, 0xc9, 0xda, 0xc2, 0xd5, 0xac, 0x7d, 0x0a, 0x48,
	0xc9, 0x53, 0xcc, 0x25, 0x0d, 0x69, 0x8a, 0x13, 0x39, 0x2b, 0xe7, 0x76, 0xb5, 0x57, 0xf1, 0x3b,
	0x82, 0xa6, 0xe3, 0x0b, 0xa1, 0x17, 0xa1, 0xc7, 0xb0, 0xac, 0xbc, 0xa9, 0x76, 0x4d, 0xb2, 0xe9,
	0x21, 0xe1, 0x36, 0x9a, 0x25, 0x8b, 0xbe, 0xd0, 0x20, 0x7a, 0x04, 0x1a, 0x20, 0x51, 0xae, 0x65,
	0xaa, 0xb0, 0x68, 0x40, 0xab, 0x74, 0x63, 0x3d, 0x3a, 0x50, 0x4b, 0x69, 0x62, 0x6b, 0xa1, 0x3e,
	0xd5, 0x10, 0x25, 0x2c, 0x50, 0x60, 0x53, 0x27, 0xab, 0x91, 0xb0, 0x31, 0x4d, 0x94, 0x27, 0xfb,
	0x3b, 0x5d, 0xbc, 0x5b, 0xc6, 0x93, 0x81, 0x54, 0x01, 0x51, 0x04, 0x1d, 0xf2, 0x4a, 0x72, 0x1c,
	0x60, 0x29, 0x39, 0x3d, 0xcc, 0x24, 0x11, 0xbd, 0x96, 0x9e, 0xb3, 0xaf, 0xf4, 0x9c, 0xdd, 0x90,
	0x48, 0x77, 0xa4, 0x8c, 0xb7, 0x66, 0xb6, 0x66, 0xe4, 0x6e, 0x93, 0xcb, 0xa8, 0xb3, 0x0d, 0xdd,
	0x32, 0xc5, 0xb7, 0x1a, 0xbb, 0x7f, 0xea, 0xb0, 0x3e, 0x9f, 0x8d, 0xed, 0xb4, 0x35, 0x68, 0x71,
	0xc6, 0xa6, 0x41, 0xb1, 0xd5, 0x14, 0xf0, 0x42, 0xb5, 0xda, 0x67, 0xd0, 0xbd, 0x5c, 0x50, 0xd5,
	0xaa, 0x32, 0xbf, 0xec, 0xee, 0xa6, 0xc5, 0x7a, 0x1a, 0x11, 0xfa, 0x08, 0x3a, 0x45, 0x13, 0xed,
	0xd6, 0x24, 0xf1, 0x76, 0x01, 0x2f, 0xf3, 0x3e, 0x25, 0x12, 0x47, 0x58, 0x62, 0xdb, 0x5b, 0x45,
	0xef, 0xcf, 0xad, 0x08, 0x9d, 0xc1, 0x6a, 0xd1, 0xa4, 0x50, 0x82, 0xb6, 0x2e, 0xc1, 0x93, 0x1b,
	0x4a, 0x60, 0x2f, 0xbd, 0x42, 0x23, 0x5e, 0xad, 0xc4, 0x4a, 0x5a, 0x26, 0x43, 0x8f, 0xa0, 0xcd,
	0x4d, 0x01, 0x75, 0xcb, 0xe8, 0xf9, 0xd2, 0x7d, 0x0d, 0x16, 0x56, 0xbd, 0x33, 0xbb, 0xad, 0xeb,
	0xe5, 0xb7, 0x75, 0xa3, 0x78, 0x5b, 0xbb, 0xd0, 0xe4, 0x44, 0x64, 0xb1, 0xd4, 0xfd, 0xb7, 0xbc,
	0xb9, 0xaa, 0xa9, 0x17, 0x29, 0x6b, 0xa9, 0x6f, 0xb5, 0xae, 0x8d, 0x71, 0xeb, 0xea, 0x18, 0xa3,
	0x01, 0x74, 0x95, 0x46, 0x64, 0xed, 0x03, 0x9e, 0xc5, 0x44, 0x69, 0x82, 0xd6, 0xbc, 0x23, 0x68,
	0x5a, 0xcc, 0x86, 0x17, 0x39, 0xcf, 0xc0, 0x99, 0x9f, 0x89, 0xb7, 0x69, 0xb5, 0x8f, 0x7f, 0x81,
	0x3b, 0xd7, 0x98, 0xa3, 0x1e, 0x74, 0x77, 0x47, 0x3b, 0x5b, 0xc3, 0x9f, 0x82, 0xad, 0xe1, 0x70,
	0x34, 0xde, 0x0f, 0x5e, 0xfa, 0xc1, 0xd8, 0x7b, 0xd1, 0x79, 0x07, 0x01, 0x34, 0x0d, 0xd4, 0xa9,
	0xa0, 0xdb, 0xd0, 0xf6, 0x47, 0x3f, 0x1e, 0x8c, 0xf6, 0xf6, 0xb5, 0xb0, 0xaa, 0x84, 0xfe, 0xe8,
	0xfb, 0xd1, 0x70, 0xbf, 0x53, 0x43, 0x0b, 0x50, 0x7f, 0xea, 0xbf, 0x1c, 0x77, 0xea, 0x9b, 0xff,
	0x36, 0xa0, 0xe9, 0xbd, 0x54, 0x4f, 0x17, 0xfa, 0x06, 0x16, 0x87, 0x9c, 0x60, 0x49, 0xcc, 0x4a,
	0x80, 0xca, 0x76, 0x04, 0x67, 0xd5, 0x35, 0xeb, 0x8a, 0x9b, 0xaf, 0x2b, 0xee, 0x48, 0xad, 0x2b,
	0xca, 0xd8, 0xbc, 0x45, 0xff, 0xc7, 0xf8, 0x0b, 0x68, 0xcd, 0xb6, 0x1c, 0xb4, 0x92, 0x3f, 0xab,
	0x97, 0xb6, 0x1e, 0xa7, 0xcc, 0x21, 0x1a, 0x01, 0xec, 0x52, 0x91, 0x5b, 0x3a, 0x33, 0x95, 0x0b,
	0x30, 0x37, 0x5f, 0x2b, 0x95, 0xd9, 0x61, 0xdd, 0x86, 0xa5, 0x4b, 0x3b, 0x13, 0xba, 0xa7, 0x39,
	0x94, 0xed, 0x51, 0x73, 0x63, 0xf8, 0x16, 0x96, 0x4c, 0xf6, 0xec, 0x03, 0x8a, 0x4a, 0xf7, 0x8f,
	0xb9, 0xe6, 0x1e, 0x2c, 0x5f, 0xde, 0x24, 0x90, 0x53, 0xba, 0x5e, 0xe4, 0xd1, 0xcc, 0x5f, 0x3d,
	0xd0, 0x2e, 0xa0, 0xeb, 0x6b, 0x01, 0xba, 0x5f, 0x08, 0xa9, 0x64, 0x5f, 0x98, 0x4b, 0xec, 0x57,
	0xb8, 0x37, 0xf7, 0x5d, 0x45, 0x8f, 0x73, 0x1e, 0xaf, 0x5d, 0x0d, 0x9c, 0x0f, 0x6e, 0x52, 0xb3,
	0xcc, 0x27, 0xd0, 0x9b, 0x77, 0xc7, 0xa0, 0xf7, 0xdf, 0xe4, 0x15, 0x70, 0x1e, 0xbf, 0xd1, 0x45,
	0xb5, 0xfd, 0xf0, 0xe7, 0x07, 0x13, 0x2a, 0x8f, 0xb3, 0x43, 0x37, 0x64, 0xd3, 0x81, 0xad, 0x93,
	0xd9, 0xc0, 0x43, 0x16, 0x0f, 0x78, 0x1a, 0x1e, 0x36, 0xf5, 0xe9, 0xf3, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x72, 0xad, 0xd4, 0x68, 0xdf, 0x0b, 0x00, 0x00,
}
