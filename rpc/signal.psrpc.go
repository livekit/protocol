// Code generated by protoc-gen-psrpc v0.5.0, DO NOT EDIT.
// source: rpc/signal.proto

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

var _ = version.PsrpcVersion_0_5

// =======================
// Signal Client Interface
// =======================

type SignalClient[NodeIdTopicType ~string] interface {
	RelaySignal(ctx context.Context, nodeId NodeIdTopicType, opts ...psrpc.RequestOption) (psrpc.ClientStream[*RelaySignalRequest, *RelaySignalResponse], error)
}

// ===========================
// Signal ServerImpl Interface
// ===========================

type SignalServerImpl interface {
	RelaySignal(psrpc.ServerStream[*RelaySignalResponse, *RelaySignalRequest]) error
}

// =======================
// Signal Server Interface
// =======================

type SignalServer[NodeIdTopicType ~string] interface {
	RegisterRelaySignalTopic(nodeId NodeIdTopicType) error
	DeregisterRelaySignalTopic(nodeId NodeIdTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// Signal Client
// =============

type signalClient[NodeIdTopicType ~string] struct {
	client *client.RPCClient
}

// NewSignalClient creates a psrpc client that implements the SignalClient interface.
func NewSignalClient[NodeIdTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (SignalClient[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Signal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("RelaySignal", false, false, false, true)

	rpcClient, err := client.NewRPCClientWithStreams(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &signalClient[NodeIdTopicType]{
		client: rpcClient,
	}, nil
}

func (c *signalClient[NodeIdTopicType]) RelaySignal(ctx context.Context, nodeId NodeIdTopicType, opts ...psrpc.RequestOption) (psrpc.ClientStream[*RelaySignalRequest, *RelaySignalResponse], error) {
	return client.OpenStream[*RelaySignalRequest, *RelaySignalResponse](ctx, c.client, "RelaySignal", []string{string(nodeId)}, opts...)
}

// =============
// Signal Server
// =============

type signalServer[NodeIdTopicType ~string] struct {
	svc SignalServerImpl
	rpc *server.RPCServer
}

// NewSignalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewSignalServer[NodeIdTopicType ~string](svc SignalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (SignalServer[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Signal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("RelaySignal", false, false, false, true)
	return &signalServer[NodeIdTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *signalServer[NodeIdTopicType]) RegisterRelaySignalTopic(nodeId NodeIdTopicType) error {
	return server.RegisterStreamHandler(s.rpc, "RelaySignal", []string{string(nodeId)}, s.svc.RelaySignal, nil)
}

func (s *signalServer[NodeIdTopicType]) DeregisterRelaySignalTopic(nodeId NodeIdTopicType) {
	s.rpc.DeregisterHandler("RelaySignal", []string{string(nodeId)})
}

func (s *signalServer[NodeIdTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *signalServer[NodeIdTopicType]) Kill() {
	s.rpc.Close(true)
}

<<<<<<< Updated upstream
var psrpcFileDescriptor4 = []byte{
=======
var psrpcFileDescriptor5 = []byte{
>>>>>>> Stashed changes
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0xc6, 0x53, 0x17, 0x10, 0x8b, 0x24, 0x58, 0x14, 0x9a, 0x3d, 0x6d, 0x38, 0x6d, 0x8c, 0xe9,
	0x26, 0x6b, 0xbc, 0x78, 0xf4, 0x11, 0xca, 0x49, 0x2f, 0x04, 0xca, 0x04, 0x1b, 0xd6, 0xb6, 0x74,
	0x0a, 0x89, 0x8f, 0xe0, 0xbb, 0x78, 0xf2, 0x09, 0x0d, 0x94, 0x3f, 0x6b, 0xf0, 0xb4, 0x33, 0xf3,
	0x7d, 0xfb, 0xfd, 0x26, 0x53, 0xda, 0xf3, 0x4e, 0x15, 0xa8, 0x17, 0x66, 0x5a, 0x09, 0xe7, 0x6d,
	0xb0, 0x2c, 0xf1, 0x4e, 0xa5, 0x5d, 0xeb, 0x82, 0xb6, 0x06, 0xe3, 0x2c, 0x1d, 0x54, 0x7a, 0x03,
	0x4b, 0x1d, 0x26, 0xda, 0x04, 0xf0, 0x47, 0x6f, 0x7a, 0x73, 0x98, 0xfb, 0xa0, 0xe2, 0x68, 0xf4,
	0x4d, 0x28, 0x93, 0x50, 0x4d, 0x3f, 0xc7, 0xbb, 0x50, 0x09, 0xab, 0x35, 0x60, 0x60, 0xcf, 0xb4,
	0x8b, 0x61, 0xea, 0xc3, 0x04, 0x01, 0x51, 0x5b, 0xc3, 0x49, 0x46, 0xf2, 0x4e, 0x79, 0x27, 0xf6,
	0x09, 0x62, 0xbc, 0x55, 0xc7, 0x51, 0x94, 0xd7, 0x58, 0xeb, 0x58, 0x49, 0xdb, 0x3e, 0xc6, 0x20,
	0x4f, 0xb2, 0x24, 0xef, 0x94, 0x83, 0xd3, 0x6f, 0x75, 0x8a, 0x3c, 0xfa, 0x58, 0x8f, 0x26, 0x08,
	0x2b, 0xde, 0xc8, 0x48, 0xde, 0x90, 0xdb, 0x92, 0xdd, 0xd2, 0xa6, 0xaa, 0x2c, 0x02, 0x6f, 0x66,
	0x24, 0x6f, 0xcb, 0xd8, 0x8c, 0x02, 0xed, 0xff, 0xd9, 0x16, 0x9d, 0x35, 0x08, 0xec, 0x89, 0x5e,
	0xf9, 0x7d, 0x8d, 0xfc, 0x62, 0xc7, 0x1c, 0x9e, 0x31, 0xa3, 0x2e, 0x4f, 0xce, 0x03, 0x35, 0xf9,
	0x87, 0xda, 0xa8, 0x51, 0x4b, 0x45, 0x5b, 0x31, 0x84, 0xbd, 0xd2, 0x4e, 0x8d, 0xcf, 0x86, 0xc2,
	0x3b, 0x25, 0xce, 0xef, 0x97, 0xf2, 0x73, 0x21, 0x42, 0x47, 0xc3, 0x9f, 0x2f, 0xd2, 0xef, 0x91,
	0xb4, 0xcb, 0x2e, 0x8d, 0x9d, 0xc3, 0x44, 0xcf, 0xb7, 0xb7, 0xcd, 0xc8, 0xcb, 0xc3, 0xdb, 0xfd,
	0x42, 0x87, 0xf7, 0xf5, 0x4c, 0x28, 0xfb, 0x51, 0xec, 0x97, 0x3f, 0x7e, 0xdd, 0x72, 0x51, 0x20,
	0xf8, 0x8d, 0x56, 0x50, 0x78, 0xa7, 0x66, 0xad, 0xdd, 0xf3, 0x3d, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x44, 0x2f, 0x12, 0x62, 0x11, 0x02, 0x00, 0x00,
}
