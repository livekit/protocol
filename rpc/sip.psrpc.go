// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
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
import google_protobuf "google.golang.org/protobuf/types/known/emptypb"

var _ = version.PsrpcVersion_0_6

// ============================
// SIPInternal Client Interface
// ============================

type SIPInternalClient interface {
	CreateSIPParticipant(ctx context.Context, topic string, req *InternalCreateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalCreateSIPParticipantResponse, error)

	TransferSIPParticipant(ctx context.Context, sipCallId string, req *InternalTransferSIPParticipantRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ================================
// SIPInternal ServerImpl Interface
// ================================

type SIPInternalServerImpl interface {
	CreateSIPParticipant(context.Context, *InternalCreateSIPParticipantRequest) (*InternalCreateSIPParticipantResponse, error)
	CreateSIPParticipantAffinity(context.Context, *InternalCreateSIPParticipantRequest) float32

	TransferSIPParticipant(context.Context, *InternalTransferSIPParticipantRequest) (*google_protobuf.Empty, error)
}

// ============================
// SIPInternal Server Interface
// ============================

type SIPInternalServer interface {
	RegisterCreateSIPParticipantTopic(topic string) error
	DeregisterCreateSIPParticipantTopic(topic string)
	RegisterTransferSIPParticipantTopic(sipCallId string) error
	DeregisterTransferSIPParticipantTopic(sipCallId string)

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
	sd.RegisterMethod("TransferSIPParticipant", false, false, true, true)

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

func (c *sIPInternalClient) TransferSIPParticipant(ctx context.Context, sipCallId string, req *InternalTransferSIPParticipantRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "TransferSIPParticipant", []string{sipCallId}, req, opts...)
}

func (s *sIPInternalClient) Close() {
	s.client.Close()
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
	sd.RegisterMethod("TransferSIPParticipant", false, false, true, true)
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

func (s *sIPInternalServer) RegisterTransferSIPParticipantTopic(sipCallId string) error {
	return server.RegisterHandler(s.rpc, "TransferSIPParticipant", []string{sipCallId}, s.svc.TransferSIPParticipant, nil)
}

func (s *sIPInternalServer) DeregisterTransferSIPParticipantTopic(sipCallId string) {
	s.rpc.DeregisterHandler("TransferSIPParticipant", []string{sipCallId})
}

func (s *sIPInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *sIPInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor10 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xc1, 0x6e, 0x1b, 0x37,
	0x10, 0x05, 0xad, 0x44, 0xb6, 0x46, 0x96, 0x2d, 0xd3, 0x96, 0xc2, 0x6e, 0xd0, 0x44, 0x55, 0x1a,
	0x40, 0xe9, 0x61, 0xd5, 0x3a, 0x28, 0x50, 0xe4, 0xd4, 0xc4, 0x75, 0x11, 0x1d, 0xda, 0x1a, 0xb2,
	0x7a, 0xe9, 0x65, 0x41, 0xed, 0xd2, 0x32, 0xeb, 0xdd, 0x25, 0x4b, 0x72, 0xe3, 0xa8, 0x7f, 0x10,
	0x14, 0xfd, 0x87, 0x7e, 0x43, 0x7e, 0xaf, 0x97, 0x82, 0xe4, 0xae, 0x23, 0x4b, 0xb2, 0x6b, 0xdf,
	0x38, 0xef, 0xcd, 0xbc, 0x9d, 0x19, 0xee, 0x0c, 0xa1, 0xa5, 0x64, 0x3c, 0xd4, 0x5c, 0x86, 0x52,
	0x09, 0x23, 0x70, 0x4d, 0xc9, 0x38, 0x78, 0x3c, 0x13, 0x62, 0x96, 0xb2, 0xa1, 0x83, 0xa6, 0xc5,
	0xd9, 0x90, 0x65, 0xd2, 0xcc, 0xbd, 0x47, 0xf0, 0x64, 0x99, 0x4c, 0x0a, 0x45, 0x0d, 0x17, 0x79,
	0xc9, 0xb7, 0x84, 0xb4, 0x96, 0x2e, 0xcd, 0xbd, 0x94, 0xbf, 0x63, 0x17, 0xdc, 0x44, 0x57, 0xdf,
	0xe8, 0xff, 0x05, 0xf0, 0x6c, 0x94, 0x1b, 0xa6, 0x72, 0x9a, 0x1e, 0x29, 0x46, 0x0d, 0x3b, 0x1d,
	0x9d, 0x9c, 0x50, 0x65, 0x78, 0xcc, 0x25, 0xcd, 0xcd, 0x98, 0xfd, 0x51, 0x30, 0x6d, 0xf0, 0xe7,
	0x00, 0x52, 0x89, 0xdf, 0x59, 0x6c, 0x22, 0x9e, 0x10, 0xdc, 0x43, 0x83, 0xc6, 0xb8, 0x51, 0x22,
	0xa3, 0x04, 0x3f, 0x81, 0xa6, 0xe6, 0x32, 0x8a, 0x69, 0x9a, 0x5a, 0xbe, 0xe5, 0x79, 0xcd, 0xe5,
	0x11, 0x4d, 0xd3, 0x51, 0x82, 0x7b, 0xb0, 0x6d, 0x79, 0xa3, 0x8a, 0xfc, 0xc2, 0x3a, 0xec, 0x3b,
	0x07, 0xd0, 0x5c, 0x4e, 0x2c, 0x34, 0x4a, 0x30, 0x81, 0x4d, 0x9a, 0x24, 0x8a, 0x69, 0x4d, 0x36,
	0x1c, 0x59, 0x99, 0x38, 0x80, 0xad, 0x73, 0xa1, 0x4d, 0x4e, 0x33, 0x46, 0x0e, 0x1c, 0x75, 0x65,
	0xe3, 0x97, 0xd0, 0x30, 0x8a, 0xe6, 0x5a, 0x0a, 0x65, 0x48, 0xbb, 0x87, 0x06, 0x3b, 0x87, 0x9d,
	0xb0, 0xac, 0x32, 0x3c, 0x1d, 0x9d, 0x4c, 0x2a, 0x72, 0xfc, 0xc9, 0x0f, 0x77, 0xa1, 0x9e, 0x17,
	0xd9, 0x94, 0x29, 0x52, 0x73, 0x72, 0xa5, 0x85, 0x1f, 0xc1, 0xa6, 0x2b, 0xc0, 0x08, 0xf2, 0xc0,
	0x13, 0xd6, 0x9c, 0x08, 0x9b, 0x41, 0xa1, 0x6d, 0x8b, 0x32, 0x46, 0x1e, 0xfa, 0x0c, 0x2a, 0xdb,
	0x72, 0x92, 0x6a, 0x7d, 0x29, 0x54, 0x42, 0xea, 0x9e, 0xab, 0x6c, 0xfc, 0x18, 0x1a, 0x4a, 0x88,
	0x2c, 0x72, 0x81, 0x9b, 0x9e, 0xb4, 0xc0, 0xcf, 0x36, 0xf0, 0x1b, 0x38, 0x90, 0x9f, 0xfa, 0x1c,
	0xf1, 0x84, 0xe5, 0x86, 0x9b, 0x39, 0xd9, 0x72, 0x7e, 0xfb, 0x0b, 0xdc, 0xa8, 0xa4, 0xf0, 0x0b,
	0x68, 0x2f, 0x86, 0x38, 0xd9, 0x1d, 0xe7, 0xbe, 0xbb, 0x80, 0xaf, 0x53, 0xcf, 0x98, 0xa1, 0x09,
	0x35, 0x94, 0xec, 0xae, 0xa8, 0xff, 0x54, 0x52, 0xf8, 0x4f, 0xe8, 0x2e, 0x86, 0x50, 0x63, 0x14,
	0x9f, 0x16, 0x86, 0x69, 0xb2, 0xd7, 0xab, 0x0d, 0x9a, 0x87, 0x47, 0xa1, 0x92, 0x71, 0x78, 0x87,
	0x9f, 0x25, 0x5c, 0x80, 0x5e, 0x5f, 0xa9, 0x1c, 0xe7, 0x46, 0xcd, 0xc7, 0x1d, 0xb9, 0x8e, 0xc3,
	0x07, 0xf0, 0xd0, 0x88, 0x0b, 0x96, 0x93, 0x86, 0xcb, 0xcf, 0x1b, 0xb8, 0x03, 0xf5, 0x4b, 0x1d,
	0x15, 0x2a, 0x25, 0xe0, 0xe1, 0x4b, 0xfd, 0xab, 0x4a, 0x31, 0x86, 0x07, 0x89, 0xc9, 0xce, 0x48,
	0xd3, 0x81, 0xee, 0x8c, 0x9f, 0x41, 0x4b, 0xa6, 0x74, 0x1e, 0x29, 0x9e, 0xcf, 0x8c, 0xc8, 0x19,
	0xd9, 0xee, 0xa1, 0xc1, 0xd6, 0x78, 0xdb, 0x82, 0xe3, 0x12, 0xc3, 0xbf, 0xc0, 0xe6, 0x39, 0xa3,
	0x09, 0x53, 0x9a, 0x74, 0x5c, 0x49, 0xdf, 0xde, 0xb9, 0xa4, 0xb7, 0x3e, 0xce, 0x17, 0x51, 0xa9,
	0xe0, 0x02, 0x3a, 0xe5, 0x31, 0x32, 0x62, 0xb1, 0x63, 0x5d, 0x27, 0xff, 0xfa, 0xbe, 0xf2, 0x13,
	0xb1, 0xdc, 0xaf, 0xfd, 0xf3, 0x55, 0x06, 0xbf, 0x81, 0x5d, 0x5b, 0x27, 0xcf, 0x67, 0x91, 0xe1,
	0x19, 0x13, 0x85, 0x21, 0x8f, 0x7a, 0x68, 0xd0, 0x3c, 0xfc, 0x2c, 0xf4, 0x0b, 0x21, 0xac, 0x16,
	0x42, 0xf8, 0x43, 0xb9, 0x10, 0xc6, 0x3b, 0x65, 0xc4, 0xc4, 0x07, 0xe0, 0x63, 0xd8, 0xcb, 0xe8,
	0x7b, 0x3f, 0xb1, 0xd5, 0xd6, 0x20, 0xe4, 0xff, 0x54, 0x76, 0x33, 0xfa, 0xde, 0x8e, 0x74, 0x05,
	0x04, 0x6f, 0x21, 0xb8, 0xf9, 0xb6, 0x71, 0x1b, 0x6a, 0x17, 0x6c, 0x4e, 0x90, 0xbb, 0x28, 0x7b,
	0xb4, 0x17, 0xfd, 0x8e, 0xa6, 0x05, 0x2b, 0x87, 0xdc, 0x1b, 0xaf, 0x36, 0xbe, 0x43, 0xc1, 0x2b,
	0xd8, 0x5e, 0x6c, 0xf2, 0xbd, 0x62, 0x7f, 0x04, 0x72, 0x53, 0x07, 0xef, 0xa3, 0xd3, 0xff, 0x07,
	0xc1, 0x97, 0xb7, 0x5f, 0x97, 0x96, 0x22, 0xd7, 0x0c, 0x3f, 0x87, 0x9d, 0xeb, 0xc3, 0x5b, 0xea,
	0xb7, 0xae, 0x8d, 0xed, 0x8d, 0x33, 0xbe, 0x71, 0xf3, 0x8c, 0x2f, 0x6d, 0xd2, 0xda, 0xd2, 0x26,
	0xed, 0xff, 0x8d, 0xe0, 0x79, 0x95, 0xa2, 0xdb, 0x6e, 0x67, 0x4c, 0xad, 0x5f, 0xd9, 0x4b, 0x4a,
	0x68, 0x79, 0x27, 0x3f, 0x85, 0xa6, 0x29, 0x05, 0xec, 0xca, 0xf3, 0x39, 0x41, 0x05, 0x4d, 0xc4,
	0xea, 0x4c, 0xd5, 0x56, 0x67, 0xea, 0xf0, 0x5f, 0x04, 0xcd, 0xd3, 0xd1, 0x49, 0x95, 0x12, 0xbe,
	0x84, 0x83, 0x75, 0x9d, 0xc3, 0x83, 0xbb, 0xce, 0x42, 0xf0, 0xe2, 0x0e, 0x9e, 0xfe, 0x1a, 0xfa,
	0xf0, 0xf1, 0x03, 0xaa, 0xb7, 0xd1, 0xf7, 0xe8, 0x6b, 0x84, 0x35, 0x74, 0xd7, 0xf7, 0x03, 0x7f,
	0x75, 0x4d, 0xf0, 0xd6, 0xa6, 0x05, 0xdd, 0x95, 0x7f, 0xff, 0xd8, 0xbe, 0xb7, 0xfd, 0xce, 0xc7,
	0x0f, 0x68, 0xaf, 0x8d, 0x82, 0x16, 0x5e, 0x6c, 0xea, 0x9b, 0x2f, 0x7e, 0x7b, 0x3a, 0xe3, 0xe6,
	0xbc, 0x98, 0x86, 0xb1, 0xc8, 0x86, 0xe5, 0xc3, 0xe3, 0x9f, 0xe3, 0x58, 0xa4, 0x43, 0x25, 0xe3,
	0x69, 0xdd, 0x59, 0x2f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xc5, 0x6f, 0x55, 0xdd, 0x07,
	0x00, 0x00,
}
