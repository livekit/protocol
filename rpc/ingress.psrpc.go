// Code generated by protoc-gen-psrpc v0.2.4, DO NOT EDIT.
// source: rpc/ingress.proto

package rpc

import context "context"
import psrpc "github.com/livekit/psrpc"
import version "github.com/livekit/psrpc/version"
import livekit2 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_2_4

// ================================
// IngressInternal Client Interface
// ================================

type IngressInternalClient interface {
	ListActiveIngress(context.Context, *ListActiveIngressRequest, ...psrpc.RequestOption) (<-chan *psrpc.Response[*ListActiveIngressResponse], error)
}

// ====================================
// IngressInternal ServerImpl Interface
// ====================================

type IngressInternalServerImpl interface {
	ListActiveIngress(context.Context, *ListActiveIngressRequest) (*ListActiveIngressResponse, error)
}

// ================================
// IngressInternal Server Interface
// ================================

type IngressInternalServer interface {
	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ======================
// IngressInternal Client
// ======================

type ingressInternalClient struct {
	client *psrpc.RPCClient
}

// NewIngressInternalClient creates a psrpc client that implements the IngressInternalClient interface.
func NewIngressInternalClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IngressInternalClient, error) {
	rpcClient, err := psrpc.NewRPCClient("IngressInternal", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &ingressInternalClient{
		client: rpcClient,
	}, nil
}

func (c *ingressInternalClient) ListActiveIngress(ctx context.Context, req *ListActiveIngressRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*ListActiveIngressResponse], error) {
	return psrpc.RequestMulti[*ListActiveIngressResponse](ctx, c.client, "ListActiveIngress", "", req, opts...)
}

// ======================
// IngressInternal Server
// ======================

type ingressInternalServer struct {
	svc IngressInternalServerImpl
	rpc *psrpc.RPCServer
}

// NewIngressInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIngressInternalServer(serverID string, svc IngressInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IngressInternalServer, error) {
	s := psrpc.NewRPCServer("IngressInternal", serverID, bus, opts...)

	var err error
	err = psrpc.RegisterHandler(s, "ListActiveIngress", "", svc.ListActiveIngress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &ingressInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *ingressInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *ingressInternalServer) Kill() {
	s.rpc.Close(true)
}

// ===============================
// IngressHandler Client Interface
// ===============================

type IngressHandlerClient interface {
	UpdateIngress(context.Context, string, *livekit2.UpdateIngressRequest, ...psrpc.RequestOption) (*livekit2.IngressState, error)

	DeleteIngress(context.Context, string, *livekit2.DeleteIngressRequest, ...psrpc.RequestOption) (*livekit2.IngressState, error)
}

// ===================================
// IngressHandler ServerImpl Interface
// ===================================

type IngressHandlerServerImpl interface {
	UpdateIngress(context.Context, *livekit2.UpdateIngressRequest) (*livekit2.IngressState, error)

	DeleteIngress(context.Context, *livekit2.DeleteIngressRequest) (*livekit2.IngressState, error)
}

// ===============================
// IngressHandler Server Interface
// ===============================

type IngressHandlerServer interface {
	RegisterUpdateIngressTopic(string) error
	DeregisterUpdateIngressTopic(string)

	RegisterDeleteIngressTopic(string) error
	DeregisterDeleteIngressTopic(string)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =====================
// IngressHandler Client
// =====================

type ingressHandlerClient struct {
	client *psrpc.RPCClient
}

// NewIngressHandlerClient creates a psrpc client that implements the IngressHandlerClient interface.
func NewIngressHandlerClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IngressHandlerClient, error) {
	rpcClient, err := psrpc.NewRPCClient("IngressHandler", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &ingressHandlerClient{
		client: rpcClient,
	}, nil
}

func (c *ingressHandlerClient) UpdateIngress(ctx context.Context, topic string, req *livekit2.UpdateIngressRequest, opts ...psrpc.RequestOption) (*livekit2.IngressState, error) {
	return psrpc.RequestSingle[*livekit2.IngressState](ctx, c.client, "UpdateIngress", topic, req, opts...)
}

func (c *ingressHandlerClient) DeleteIngress(ctx context.Context, topic string, req *livekit2.DeleteIngressRequest, opts ...psrpc.RequestOption) (*livekit2.IngressState, error) {
	return psrpc.RequestSingle[*livekit2.IngressState](ctx, c.client, "DeleteIngress", topic, req, opts...)
}

// =====================
// IngressHandler Server
// =====================

type ingressHandlerServer struct {
	svc IngressHandlerServerImpl
	rpc *psrpc.RPCServer
}

// NewIngressHandlerServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIngressHandlerServer(serverID string, svc IngressHandlerServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IngressHandlerServer, error) {
	s := psrpc.NewRPCServer("IngressHandler", serverID, bus, opts...)

	return &ingressHandlerServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *ingressHandlerServer) RegisterUpdateIngressTopic(topic string) error {
	return psrpc.RegisterHandler(s.rpc, "UpdateIngress", topic, s.svc.UpdateIngress, nil)
}

func (s *ingressHandlerServer) DeregisterUpdateIngressTopic(topic string) {
	s.rpc.DeregisterHandler("UpdateIngress", topic)
}

func (s *ingressHandlerServer) RegisterDeleteIngressTopic(topic string) error {
	return psrpc.RegisterHandler(s.rpc, "DeleteIngress", topic, s.svc.DeleteIngress, nil)
}

func (s *ingressHandlerServer) DeregisterDeleteIngressTopic(topic string) {
	s.rpc.DeregisterHandler("DeleteIngress", topic)
}

func (s *ingressHandlerServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *ingressHandlerServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor1 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x89, 0x85, 0xa2, 0x91, 0x2a, 0x0d, 0x14, 0xea, 0x82, 0x7f, 0xd8, 0x93, 0x88, 0x64,
	0xa1, 0x5e, 0xbd, 0x28, 0x1e, 0x5c, 0xd0, 0x4b, 0xc5, 0x4b, 0x2f, 0x65, 0x9b, 0x1d, 0xd6, 0xd0,
	0x6d, 0x12, 0x33, 0xd3, 0x7d, 0x07, 0x5f, 0xc6, 0x83, 0x4f, 0x28, 0xda, 0xb1, 0xb8, 0xd8, 0x82,
	0xa7, 0xc0, 0xf7, 0x9b, 0xfc, 0xbe, 0x21, 0x91, 0xfd, 0x18, 0x4c, 0x66, 0x5d, 0x15, 0x01, 0x51,
	0x87, 0xe8, 0xc9, 0xab, 0x4e, 0x0c, 0x26, 0xe9, 0xf9, 0x40, 0xd6, 0x3b, 0xce, 0x92, 0x41, 0x6d,
	0x1b, 0x98, 0x5b, 0x9a, 0xb6, 0x46, 0xd3, 0x44, 0x0e, 0x1f, 0x2c, 0xd2, 0x8d, 0x21, 0xdb, 0x40,
	0xbe, 0x42, 0x63, 0x78, 0x5d, 0x02, 0x52, 0x7a, 0x2d, 0x8f, 0x36, 0x30, 0x0c, 0xde, 0x21, 0xa8,
	0x53, 0xb9, 0xcf, 0xa6, 0xa9, 0x2d, 0x71, 0x28, 0xce, 0x3a, 0xe7, 0x7b, 0x63, 0xc9, 0x51, 0x5e,
	0xe2, 0x68, 0x21, 0x0f, 0xf9, 0x4e, 0xee, 0x08, 0xa2, 0x2b, 0x6a, 0x35, 0x91, 0xfd, 0x3f, 0x42,
	0x75, 0xac, 0x63, 0x30, 0x7a, 0xdb, 0x12, 0xc9, 0xc9, 0x36, 0xbc, 0xda, 0x23, 0xed, 0x7e, 0xbc,
	0x89, 0x9d, 0x5d, 0x31, 0x7a, 0x17, 0xf2, 0x80, 0xd9, 0x7d, 0xe1, 0xca, 0x1a, 0xa2, 0x7a, 0x94,
	0xbd, 0xe7, 0x50, 0x16, 0xf4, 0xab, 0x8a, 0x1f, 0x41, 0xb7, 0xf2, 0x9f, 0xaa, 0xc1, 0x1a, 0x33,
	0x78, 0xa2, 0x82, 0xb8, 0x61, 0x28, 0xbe, 0x74, 0x77, 0x50, 0xc3, 0x26, 0x5d, 0x2b, 0xff, 0x9f,
	0xee, 0xf6, 0x72, 0x72, 0x51, 0x59, 0x7a, 0x59, 0xce, 0xb4, 0xf1, 0x8b, 0x8c, 0x47, 0xd7, 0x67,
	0x98, 0x57, 0x19, 0x42, 0x6c, 0xac, 0x81, 0x2c, 0x06, 0x33, 0xeb, 0x7e, 0x7f, 0xd7, 0xd5, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x7d, 0xaa, 0x66, 0xee, 0x01, 0x00, 0x00,
}
