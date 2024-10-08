// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/keepalive.proto

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

var _ = version.PsrpcVersion_0_6

// ==========================
// Keepalive Client Interface
// ==========================

type KeepaliveClient[NodeIDTopicType ~string] interface {
	SubscribePing(ctx context.Context, nodeID NodeIDTopicType) (psrpc.Subscription[*KeepalivePing], error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ==============================
// Keepalive ServerImpl Interface
// ==============================

type KeepaliveServerImpl interface {
}

// ==========================
// Keepalive Server Interface
// ==========================

type KeepaliveServer[NodeIDTopicType ~string] interface {
	PublishPing(ctx context.Context, nodeID NodeIDTopicType, msg *KeepalivePing) error

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ================
// Keepalive Client
// ================

type keepaliveClient[NodeIDTopicType ~string] struct {
	client *client.RPCClient
}

// NewKeepaliveClient creates a psrpc client that implements the KeepaliveClient interface.
func NewKeepaliveClient[NodeIDTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (KeepaliveClient[NodeIDTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Keepalive",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("Ping", false, true, false, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &keepaliveClient[NodeIDTopicType]{
		client: rpcClient,
	}, nil
}

func (c *keepaliveClient[NodeIDTopicType]) SubscribePing(ctx context.Context, nodeID NodeIDTopicType) (psrpc.Subscription[*KeepalivePing], error) {
	return client.Join[*KeepalivePing](ctx, c.client, "Ping", []string{string(nodeID)})
}

func (s *keepaliveClient[NodeIDTopicType]) Close() {
	s.client.Close()
}

// ================
// Keepalive Server
// ================

type keepaliveServer[NodeIDTopicType ~string] struct {
	svc KeepaliveServerImpl
	rpc *server.RPCServer
}

// NewKeepaliveServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewKeepaliveServer[NodeIDTopicType ~string](svc KeepaliveServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (KeepaliveServer[NodeIDTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Keepalive",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("Ping", false, true, false, false)
	return &keepaliveServer[NodeIDTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *keepaliveServer[NodeIDTopicType]) PublishPing(ctx context.Context, nodeID NodeIDTopicType, msg *KeepalivePing) error {
	return s.rpc.Publish(ctx, "Ping", []string{string(nodeID)}, msg)
}

func (s *keepaliveServer[NodeIDTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *keepaliveServer[NodeIDTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor5 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2a, 0x48, 0xd6,
	0xcf, 0x4e, 0x4d, 0x2d, 0x48, 0xcc, 0xc9, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2e, 0x2a, 0x48, 0x96, 0xe2, 0xcd, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x88, 0x29,
	0xe9, 0x72, 0xf1, 0x7a, 0xc3, 0x94, 0x05, 0x64, 0xe6, 0xa5, 0x0b, 0xc9, 0x70, 0x71, 0x96, 0x64,
	0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x21, 0x04,
	0x8c, 0x42, 0xb9, 0x38, 0xe1, 0xca, 0x85, 0x3c, 0xb8, 0x58, 0xc0, 0x5a, 0x84, 0xf4, 0x8a, 0x0a,
	0x92, 0xf5, 0x50, 0x8c, 0x91, 0xc2, 0x22, 0xa6, 0x24, 0xb1, 0xa9, 0x93, 0x51, 0x84, 0x83, 0x51,
	0x80, 0x51, 0x8a, 0x4b, 0x88, 0x2d, 0x2f, 0x3f, 0x25, 0xd5, 0xd3, 0x45, 0x82, 0xd1, 0x81, 0x49,
	0x83, 0xd1, 0x49, 0x31, 0x4a, 0x3e, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x1f, 0xa4, 0x21, 0x3b, 0xb3, 0x44, 0x1f, 0xec, 0xc2, 0xe4, 0xfc, 0x1c, 0xfd, 0xa2, 0x82, 0xe4,
	0x24, 0x36, 0x30, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x1d, 0x8f, 0xd1, 0xda, 0x00,
	0x00, 0x00,
}
