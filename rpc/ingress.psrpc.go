// Code generated by protoc-gen-psrpc v0.3.2, DO NOT EDIT.
// source: rpc/ingress.proto

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
import livekit2 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_3_2

// ================================
// IngressInternal Client Interface
// ================================

type IngressInternalClient interface {
	ListActiveIngress(ctx context.Context, req *ListActiveIngressRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*ListActiveIngressResponse], error)
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
	client *client.RPCClient
}

// NewIngressInternalClient creates a psrpc client that implements the IngressInternalClient interface.
func NewIngressInternalClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IngressInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "IngressInternal",
		ID:   clientID,
	}

	sd.RegisterMethod("ListActiveIngress", false, true, false, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &ingressInternalClient{
		client: rpcClient,
	}, nil
}

func (c *ingressInternalClient) ListActiveIngress(ctx context.Context, req *ListActiveIngressRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*ListActiveIngressResponse], error) {
	return client.RequestMulti[*ListActiveIngressResponse](ctx, c.client, "ListActiveIngress", nil, req, opts...)
}

// ======================
// IngressInternal Server
// ======================

type ingressInternalServer struct {
	svc IngressInternalServerImpl
	rpc *server.RPCServer
}

// NewIngressInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIngressInternalServer(serverID string, svc IngressInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IngressInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "IngressInternal",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("ListActiveIngress", false, true, false, false)
	var err error
	err = server.RegisterHandler(s, "ListActiveIngress", nil, svc.ListActiveIngress, nil)
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
	UpdateIngress(ctx context.Context, topic string, req *livekit2.UpdateIngressRequest, opts ...psrpc.RequestOption) (*livekit2.IngressState, error)

	DeleteIngress(ctx context.Context, topic string, req *livekit2.DeleteIngressRequest, opts ...psrpc.RequestOption) (*livekit2.IngressState, error)

	DeleteWHIPResource(ctx context.Context, topic string, req *DeleteWHIPResourceRequest, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error)
}

// ===================================
// IngressHandler ServerImpl Interface
// ===================================

type IngressHandlerServerImpl interface {
	UpdateIngress(context.Context, *livekit2.UpdateIngressRequest) (*livekit2.IngressState, error)

	DeleteIngress(context.Context, *livekit2.DeleteIngressRequest) (*livekit2.IngressState, error)

	DeleteWHIPResource(context.Context, *DeleteWHIPResourceRequest) (*google_protobuf3.Empty, error)
}

// ===============================
// IngressHandler Server Interface
// ===============================

type IngressHandlerServer interface {
	RegisterUpdateIngressTopic(topic string) error
	DeregisterUpdateIngressTopic(topic string)
	RegisterDeleteIngressTopic(topic string) error
	DeregisterDeleteIngressTopic(topic string)
	RegisterDeleteWHIPResourceTopic(topic string) error
	DeregisterDeleteWHIPResourceTopic(topic string)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =====================
// IngressHandler Client
// =====================

type ingressHandlerClient struct {
	client *client.RPCClient
}

// NewIngressHandlerClient creates a psrpc client that implements the IngressHandlerClient interface.
func NewIngressHandlerClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IngressHandlerClient, error) {
	sd := &info.ServiceDefinition{
		Name: "IngressHandler",
		ID:   clientID,
	}

	sd.RegisterMethod("UpdateIngress", false, false, true, false)
	sd.RegisterMethod("DeleteIngress", false, false, true, false)
	sd.RegisterMethod("DeleteWHIPResource", false, false, true, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &ingressHandlerClient{
		client: rpcClient,
	}, nil
}

func (c *ingressHandlerClient) UpdateIngress(ctx context.Context, topic string, req *livekit2.UpdateIngressRequest, opts ...psrpc.RequestOption) (*livekit2.IngressState, error) {
	return client.RequestSingle[*livekit2.IngressState](ctx, c.client, "UpdateIngress", []string{topic}, req, opts...)
}

func (c *ingressHandlerClient) DeleteIngress(ctx context.Context, topic string, req *livekit2.DeleteIngressRequest, opts ...psrpc.RequestOption) (*livekit2.IngressState, error) {
	return client.RequestSingle[*livekit2.IngressState](ctx, c.client, "DeleteIngress", []string{topic}, req, opts...)
}

func (c *ingressHandlerClient) DeleteWHIPResource(ctx context.Context, topic string, req *DeleteWHIPResourceRequest, opts ...psrpc.RequestOption) (*google_protobuf3.Empty, error) {
	return client.RequestSingle[*google_protobuf3.Empty](ctx, c.client, "DeleteWHIPResource", []string{topic}, req, opts...)
}

// =====================
// IngressHandler Server
// =====================

type ingressHandlerServer struct {
	svc IngressHandlerServerImpl
	rpc *server.RPCServer
}

// NewIngressHandlerServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIngressHandlerServer(serverID string, svc IngressHandlerServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IngressHandlerServer, error) {
	sd := &info.ServiceDefinition{
		Name: "IngressHandler",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("UpdateIngress", false, false, true, false)
	sd.RegisterMethod("DeleteIngress", false, false, true, false)
	sd.RegisterMethod("DeleteWHIPResource", false, false, true, false)
	return &ingressHandlerServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *ingressHandlerServer) RegisterUpdateIngressTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "UpdateIngress", []string{topic}, s.svc.UpdateIngress, nil)
}

func (s *ingressHandlerServer) DeregisterUpdateIngressTopic(topic string) {
	s.rpc.DeregisterHandler("UpdateIngress", []string{topic})
}

func (s *ingressHandlerServer) RegisterDeleteIngressTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "DeleteIngress", []string{topic}, s.svc.DeleteIngress, nil)
}

func (s *ingressHandlerServer) DeregisterDeleteIngressTopic(topic string) {
	s.rpc.DeregisterHandler("DeleteIngress", []string{topic})
}

func (s *ingressHandlerServer) RegisterDeleteWHIPResourceTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "DeleteWHIPResource", []string{topic}, s.svc.DeleteWHIPResource, nil)
}

func (s *ingressHandlerServer) DeregisterDeleteWHIPResourceTopic(topic string) {
	s.rpc.DeregisterHandler("DeleteWHIPResource", []string{topic})
}

func (s *ingressHandlerServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *ingressHandlerServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor1 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4b, 0xe3, 0x40,
	0x14, 0xc6, 0x49, 0x0b, 0x85, 0xce, 0xd2, 0xdd, 0xed, 0x40, 0x97, 0x34, 0x4b, 0xb5, 0xe4, 0x54,
	0x44, 0x26, 0x50, 0xaf, 0x5e, 0x14, 0x85, 0x06, 0x15, 0xa4, 0x2a, 0x42, 0x3d, 0x84, 0x74, 0xf2,
	0x8c, 0x43, 0x93, 0xcc, 0x38, 0x33, 0x29, 0xf4, 0xee, 0xc5, 0x7f, 0xc7, 0xbf, 0x50, 0xd2, 0x99,
	0x14, 0x4b, 0x2d, 0x78, 0x0a, 0xf9, 0x7e, 0x79, 0xdf, 0xf7, 0xde, 0xcb, 0x43, 0x5d, 0x29, 0x68,
	0xc0, 0x8a, 0x54, 0x82, 0x52, 0x44, 0x48, 0xae, 0x39, 0x6e, 0x4a, 0x41, 0xbd, 0x0e, 0x17, 0x9a,
	0xf1, 0xc2, 0x6a, 0x5e, 0x2f, 0x63, 0x4b, 0x58, 0x30, 0x1d, 0x6d, 0x7d, 0xea, 0xfd, 0x4f, 0x39,
	0x4f, 0x33, 0x08, 0xd6, 0x6f, 0xf3, 0xf2, 0x39, 0x80, 0x5c, 0xe8, 0x95, 0x81, 0xbe, 0x87, 0xdc,
	0x6b, 0xa6, 0xf4, 0x19, 0xd5, 0x6c, 0x09, 0xa1, 0xa9, 0x9b, 0xc2, 0x6b, 0x09, 0x4a, 0xfb, 0xa7,
	0xa8, 0xff, 0x0d, 0x53, 0x82, 0x17, 0x0a, 0xf0, 0x21, 0xfa, 0x65, 0x63, 0x22, 0x96, 0x28, 0xd7,
	0x19, 0x36, 0x47, 0xed, 0x29, 0xb2, 0x52, 0x98, 0x28, 0xff, 0x09, 0xf5, 0x2f, 0x20, 0x03, 0x0d,
	0x8f, 0x93, 0xf0, 0x76, 0x0a, 0x8a, 0x97, 0x92, 0x82, 0xb5, 0xae, 0xaa, 0xa5, 0x95, 0x22, 0x96,
	0xb8, 0xce, 0xd0, 0xa9, 0xaa, 0x6b, 0x29, 0x4c, 0xf0, 0x00, 0x21, 0xa5, 0x25, 0xc4, 0x79, 0xb4,
	0x80, 0x95, 0xdb, 0x58, 0xf3, 0xb6, 0x51, 0xae, 0x60, 0x35, 0xce, 0xd1, 0x1f, 0xdb, 0x50, 0x58,
	0x68, 0x90, 0x45, 0x9c, 0xe1, 0x19, 0xea, 0xee, 0x74, 0x8b, 0x07, 0x44, 0x0a, 0x4a, 0xf6, 0x4d,
	0xe8, 0x1d, 0xec, 0xc3, 0x66, 0x48, 0xbf, 0xf5, 0xf1, 0xee, 0x34, 0x46, 0xce, 0xf8, 0xad, 0x81,
	0x7e, 0x5b, 0x36, 0x89, 0x8b, 0x24, 0x03, 0x89, 0x6f, 0x50, 0xe7, 0x41, 0x24, 0xb1, 0xfe, 0x12,
	0x65, 0xd7, 0x4f, 0xb6, 0xf4, 0x3a, 0xaa, 0xb7, 0xc1, 0x16, 0xdc, 0xe9, 0x58, 0xdb, 0x84, 0xbf,
	0x4e, 0x65, 0x67, 0xb6, 0xb5, 0x6b, 0xb7, 0xa5, 0xff, 0xd0, 0xee, 0x1e, 0xe1, 0xdd, 0xe5, 0x63,
	0x33, 0xee, 0xde, 0xbf, 0xe2, 0xfd, 0x23, 0xe6, 0x54, 0x48, 0x7d, 0x2a, 0xe4, 0xb2, 0x3a, 0x95,
	0xda, 0xf5, 0xfc, 0x78, 0x76, 0x94, 0x32, 0xfd, 0x52, 0xce, 0x09, 0xe5, 0x79, 0x60, 0x1b, 0xd8,
	0x3c, 0xc5, 0x22, 0x0d, 0x14, 0xc8, 0x25, 0xa3, 0x10, 0x48, 0x41, 0xe7, 0xad, 0xb5, 0xcb, 0xc9,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xac, 0x21, 0x84, 0xbe, 0x02, 0x00, 0x00,
}
