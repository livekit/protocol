// Code generated by protoc-gen-psrpc v0.5.1, DO NOT EDIT.
// source: rpc/sip.proto

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

// ============================
// SIPInternal Client Interface
// ============================

type SIPInternalClient interface {
	CreateSIPParticipant(ctx context.Context, topic string, req *InternalCreateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalCreateSIPParticipantResponse, error)
}

// ================================
// SIPInternal ServerImpl Interface
// ================================

type SIPInternalServerImpl interface {
	CreateSIPParticipant(context.Context, *InternalCreateSIPParticipantRequest) (*InternalCreateSIPParticipantResponse, error)
	CreateSIPParticipantAffinity(context.Context, *InternalCreateSIPParticipantRequest) float32
}

// ============================
// SIPInternal Server Interface
// ============================

type SIPInternalServer interface {
	RegisterCreateSIPParticipantTopic(topic string) error
	DeregisterCreateSIPParticipantTopic(topic string)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ==================
// SIPInternal Client
// ==================

type sIPInternalClient struct {
	client *client.RPCClient
}

// NewSIPInternalClient creates a psrpc client that implements the SIPInternalClient interface.
func NewSIPInternalClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (SIPInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "SIPInternal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CreateSIPParticipant", true, false, true, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &sIPInternalClient{
		client: rpcClient,
	}, nil
}

func (c *sIPInternalClient) CreateSIPParticipant(ctx context.Context, topic string, req *InternalCreateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalCreateSIPParticipantResponse, error) {
	return client.RequestSingle[*InternalCreateSIPParticipantResponse](ctx, c.client, "CreateSIPParticipant", []string{topic}, req, opts...)
}

// ==================
// SIPInternal Server
// ==================

type sIPInternalServer struct {
	svc SIPInternalServerImpl
	rpc *server.RPCServer
}

// NewSIPInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewSIPInternalServer(svc SIPInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (SIPInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "SIPInternal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CreateSIPParticipant", true, false, true, false)
	return &sIPInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *sIPInternalServer) RegisterCreateSIPParticipantTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "CreateSIPParticipant", []string{topic}, s.svc.CreateSIPParticipant, s.svc.CreateSIPParticipantAffinity)
}

func (s *sIPInternalServer) DeregisterCreateSIPParticipantTopic(topic string) {
	s.rpc.DeregisterHandler("CreateSIPParticipant", []string{topic})
}

func (s *sIPInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *sIPInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor9 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0x56, 0xd6, 0xad, 0x5b, 0x4f, 0xd7, 0xae, 0xf3, 0xdf, 0xee, 0xb7, 0x82, 0x04, 0xa5, 0x03,
	0xa9, 0xbb, 0x49, 0xa1, 0xbb, 0x41, 0x5c, 0x01, 0x15, 0x12, 0xb9, 0x00, 0x55, 0xe9, 0xb8, 0xe1,
	0x26, 0x72, 0x13, 0x33, 0xac, 0x26, 0xb6, 0xb1, 0x9d, 0x55, 0xe5, 0x01, 0x90, 0x78, 0x0b, 0x9e,
	0x81, 0x97, 0xe1, 0x75, 0x50, 0x9c, 0xb4, 0xeb, 0x46, 0x8b, 0x7a, 0x97, 0xef, 0xfb, 0xce, 0xf9,
	0xf4, 0xd9, 0xc7, 0x27, 0xd0, 0x50, 0x32, 0x1a, 0x68, 0x26, 0x3d, 0xa9, 0x84, 0x11, 0xa8, 0xa2,
	0x64, 0xe4, 0x36, 0x84, 0x34, 0x4c, 0x70, 0x5d, 0x70, 0xee, 0x69, 0xc2, 0x6e, 0xe8, 0x8c, 0x99,
	0x70, 0x55, 0xd6, 0xfb, 0x7d, 0x00, 0xe7, 0x3e, 0x37, 0x54, 0x71, 0x92, 0x8c, 0x14, 0x25, 0x86,
	0x4e, 0xfc, 0xf1, 0x98, 0x28, 0xc3, 0x22, 0x26, 0x09, 0x37, 0x01, 0xfd, 0x9a, 0x51, 0x6d, 0xd0,
	0x43, 0xa8, 0x6b, 0x26, 0xc3, 0x88, 0x24, 0x49, 0xc8, 0x62, 0xdc, 0xe8, 0x3a, 0xfd, 0x5a, 0x50,
	0xd3, 0x4c, 0x8e, 0x48, 0x92, 0xf8, 0x31, 0xc2, 0x70, 0x48, 0xe2, 0x58, 0x51, 0xad, 0xf1, 0x9e,
	0xd5, 0x96, 0x10, 0x5d, 0x42, 0xcd, 0x28, 0xc2, 0xb5, 0x14, 0xca, 0xe0, 0x56, 0xd7, 0xe9, 0x37,
	0x87, 0x1d, 0xaf, 0x0c, 0xe2, 0x4d, 0xfc, 0xf1, 0xd5, 0x52, 0x0c, 0x6e, 0xeb, 0xd0, 0x19, 0x54,
	0x79, 0x96, 0x4e, 0xa9, 0xc2, 0x15, 0xeb, 0x56, 0x22, 0xf4, 0x3f, 0x1c, 0xda, 0x08, 0x46, 0xe0,
	0xfd, 0x42, 0xc8, 0xe1, 0x95, 0x40, 0x2e, 0x1c, 0x65, 0x3a, 0x3f, 0x45, 0x4a, 0xf1, 0x81, 0x55,
	0x56, 0x38, 0xd7, 0x24, 0xd1, 0x7a, 0x2e, 0x54, 0x8c, 0xab, 0x85, 0xb6, 0xc4, 0xe8, 0x01, 0xd4,
	0x94, 0x10, 0x69, 0x68, 0x1b, 0x0f, 0x0b, 0x31, 0x27, 0x3e, 0xe4, 0x8d, 0xcf, 0xa1, 0x2d, 0x6f,
	0xaf, 0x22, 0x64, 0x31, 0xe5, 0x86, 0x99, 0x05, 0x3e, 0xb2, 0x75, 0xff, 0xad, 0x69, 0x7e, 0x29,
	0xa1, 0x0b, 0x68, 0xad, 0xb7, 0x58, 0xdb, 0xa6, 0x2d, 0x3f, 0x59, 0xe3, 0x37, 0xb9, 0xa7, 0xd4,
	0x90, 0x98, 0x18, 0x82, 0x4f, 0xfe, 0x72, 0x7f, 0x5f, 0x4a, 0xe8, 0x1b, 0x9c, 0xad, 0xb7, 0x10,
	0x63, 0x14, 0x9b, 0x66, 0x86, 0x6a, 0x7c, 0xda, 0xad, 0xf4, 0xeb, 0xc3, 0x91, 0xa7, 0x64, 0xe4,
	0xed, 0x30, 0x4f, 0x6f, 0x8d, 0x7a, 0xbd, 0x72, 0x79, 0xcb, 0x8d, 0x5a, 0x04, 0x1d, 0xb9, 0x49,
	0x43, 0x6d, 0x38, 0x30, 0x62, 0x46, 0x39, 0xae, 0xd9, 0x7c, 0x05, 0x40, 0x1d, 0xa8, 0xce, 0x75,
	0x98, 0xa9, 0x04, 0x43, 0x41, 0xcf, 0xf5, 0x47, 0x95, 0x20, 0x04, 0xfb, 0xb1, 0x49, 0x3f, 0xe3,
	0xba, 0x25, 0xed, 0x37, 0x3a, 0x87, 0x86, 0x4c, 0xc8, 0x22, 0x54, 0x8c, 0x5f, 0x1b, 0xc1, 0x29,
	0x3e, 0xee, 0x3a, 0xfd, 0xa3, 0xe0, 0x38, 0x27, 0x83, 0x92, 0x73, 0xdf, 0x81, 0xbb, 0x3d, 0x1a,
	0x6a, 0x41, 0x65, 0x46, 0x17, 0xd8, 0xb1, 0xae, 0xf9, 0x67, 0x9e, 0xea, 0x86, 0x24, 0x19, 0x2d,
	0x5f, 0x5d, 0x01, 0x5e, 0xee, 0xbd, 0x70, 0x7a, 0x3f, 0x1d, 0x78, 0xf2, 0xef, 0x9b, 0xd0, 0x52,
	0x70, 0x4d, 0xd1, 0x53, 0x68, 0xde, 0x9d, 0x72, 0xe9, 0xdf, 0xb8, 0x33, 0xdf, 0xad, 0x8f, 0x61,
	0x6f, 0xfb, 0x63, 0xb8, 0xb7, 0x34, 0x95, 0x7b, 0x4b, 0x33, 0xfc, 0xee, 0x40, 0x7d, 0xe2, 0x8f,
	0x97, 0x29, 0xd1, 0x1c, 0xda, 0x9b, 0x92, 0xa2, 0xfe, 0xae, 0x63, 0x75, 0x2f, 0x76, 0xa8, 0x2c,
	0x8e, 0xdd, 0x83, 0x5f, 0x3f, 0x9c, 0x6a, 0xcb, 0x79, 0xe5, 0x3c, 0x73, 0xde, 0x3c, 0xfe, 0xf4,
	0xe8, 0x9a, 0x99, 0x2f, 0xd9, 0xd4, 0x8b, 0x44, 0x3a, 0x28, 0x97, 0x73, 0x60, 0xff, 0x10, 0x91,
	0x48, 0x06, 0x4a, 0x46, 0xd3, 0xaa, 0x45, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x57, 0x06,
	0xc6, 0x3b, 0x67, 0x04, 0x00, 0x00,
}
