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
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x72, 0xdb, 0x36,
	0x10, 0x1e, 0x5a, 0x89, 0x7f, 0x56, 0xbf, 0x86, 0x2c, 0x05, 0xa6, 0xa7, 0x89, 0xaa, 0x34, 0x33,
	0x4a, 0x0f, 0x54, 0xab, 0x4c, 0xa7, 0x9d, 0x1c, 0x3a, 0x4d, 0x1c, 0x67, 0xa2, 0x43, 0x1b, 0x57,
	0x56, 0x2f, 0xbd, 0x70, 0x20, 0x12, 0x96, 0x59, 0x93, 0x04, 0x0b, 0x80, 0x71, 0xd4, 0x37, 0xc8,
	0x3b, 0xf4, 0xd0, 0x67, 0xc8, 0x83, 0xf4, 0x85, 0x7a, 0xe9, 0x00, 0x20, 0x2d, 0x5a, 0x3f, 0xa9,
	0x34, 0xbd, 0x71, 0xf7, 0xdb, 0xfd, 0x76, 0x17, 0x5a, 0x7c, 0x10, 0x54, 0x79, 0xe2, 0xf5, 0x45,
	0x90, 0x38, 0x09, 0x67, 0x92, 0xa1, 0x12, 0x4f, 0x3c, 0xfb, 0x64, 0xca, 0xd8, 0x34, 0xa4, 0x7d,
	0xed, 0x9a, 0xa4, 0x97, 0x7d, 0x1a, 0x25, 0x72, 0x66, 0x22, 0xec, 0x87, 0x8b, 0xa0, 0x9f, 0x72,
	0x22, 0x03, 0x16, 0x67, 0x78, 0x95, 0x25, 0xca, 0x12, 0x99, 0x79, 0x18, 0x06, 0xef, 0xe8, 0x75,
	0x20, 0xdd, 0xdb, 0x1a, 0xdd, 0xbf, 0x2b, 0xf0, 0x78, 0x18, 0x4b, 0xca, 0x63, 0x12, 0x9e, 0x72,
	0x4a, 0x24, 0xbd, 0x18, 0x9e, 0x9f, 0x13, 0x2e, 0x03, 0x2f, 0x48, 0x48, 0x2c, 0x47, 0xf4, 0xf7,
	0x94, 0x0a, 0x89, 0x3e, 0x03, 0x48, 0x38, 0xfb, 0x8d, 0x7a, 0xd2, 0x0d, 0x7c, 0x8c, 0x3a, 0x56,
	0xef, 0x60, 0x74, 0x90, 0x79, 0x86, 0x3e, 0x7a, 0x08, 0x65, 0x11, 0x24, 0xae, 0x47, 0xc2, 0x50,
	0xe1, 0x55, 0x83, 0x8b, 0x20, 0x39, 0x25, 0x61, 0x38, 0xf4, 0x51, 0x07, 0x2a, 0x0a, 0x97, 0x3c,
	0x8d, 0xaf, 0x55, 0x40, 0x53, 0x07, 0x80, 0x08, 0x92, 0xb1, 0x72, 0x0d, 0x7d, 0x84, 0x61, 0x8f,
	0xf8, 0x3e, 0xa7, 0x42, 0xe0, 0x1d, 0x0d, 0xe6, 0x26, 0xb2, 0x61, 0xff, 0x8a, 0x09, 0x19, 0x93,
	0x88, 0xe2, 0x23, 0x0d, 0xdd, 0xda, 0xe8, 0x19, 0x1c, 0x48, 0x4e, 0x62, 0x91, 0x30, 0x2e, 0x71,
	0xa3, 0x63, 0xf5, 0x6a, 0x83, 0x96, 0x93, 0x4d, 0xe9, 0x5c, 0x0c, 0xcf, 0xc7, 0x39, 0x38, 0x9a,
	0xc7, 0xa1, 0x36, 0xec, 0xc6, 0x69, 0x34, 0xa1, 0x1c, 0x97, 0x34, 0x5d, 0x66, 0xa1, 0x07, 0xb0,
	0xa7, 0x07, 0x90, 0x0c, 0xdf, 0x33, 0x80, 0x32, 0xc7, 0x4c, 0x75, 0x90, 0x0a, 0x75, 0x44, 0x11,
	0xc5, 0xf7, 0x4d, 0x07, 0xb9, 0xad, 0xb0, 0x84, 0x08, 0x71, 0xc3, 0xb8, 0x8f, 0x77, 0x0d, 0x96,
	0xdb, 0xe8, 0x04, 0x0e, 0x38, 0x63, 0x91, 0xab, 0x13, 0xf7, 0x0c, 0xa8, 0x1c, 0x3f, 0xa9, 0xc4,
	0xaf, 0xe1, 0x28, 0x99, 0x9f, 0xb3, 0x1b, 0xf8, 0x34, 0x96, 0x81, 0x9c, 0xe1, 0x7d, 0x1d, 0xd7,
	0x2c, 0x60, 0xc3, 0x0c, 0x42, 0x4f, 0xa1, 0x51, 0x4c, 0xd1, 0xb4, 0x35, 0x1d, 0x5e, 0x2f, 0xf8,
	0x57, 0xb1, 0x47, 0x54, 0x12, 0x9f, 0x48, 0x82, 0xeb, 0x4b, 0xec, 0x3f, 0x66, 0x10, 0xfa, 0x03,
	0xda, 0xc5, 0x14, 0x22, 0x25, 0x0f, 0x26, 0xa9, 0xa4, 0x02, 0x1f, 0x76, 0x4a, 0xbd, 0xf2, 0xe0,
	0xd4, 0xe1, 0x89, 0xe7, 0x6c, 0xb0, 0x2c, 0x4e, 0xc1, 0xf5, 0xe2, 0x96, 0xe5, 0x2c, 0x96, 0x7c,
	0x36, 0x6a, 0x25, 0xab, 0x30, 0x74, 0x04, 0xf7, 0x25, 0xbb, 0xa6, 0x31, 0x3e, 0xd0, 0xfd, 0x19,
	0x03, 0xb5, 0x60, 0xf7, 0x46, 0xb8, 0x29, 0x0f, 0x31, 0x18, 0xf7, 0x8d, 0xf8, 0x85, 0x87, 0x08,
	0xc1, 0x3d, 0x5f, 0x46, 0x97, 0xb8, 0xac, 0x9d, 0xfa, 0x1b, 0x3d, 0x86, 0x6a, 0x12, 0x92, 0x99,
	0xeb, 0x07, 0x24, 0x94, 0x2c, 0xa6, 0xb8, 0xd2, 0xb1, 0x7a, 0xfb, 0xa3, 0x8a, 0x72, 0xbe, 0xca,
	0x7c, 0xe8, 0x2d, 0xec, 0x5d, 0x51, 0xe2, 0x53, 0x2e, 0x70, 0x4b, 0x8f, 0xf4, 0xcd, 0xc6, 0x23,
	0xbd, 0x31, 0x79, 0x66, 0x88, 0x9c, 0x05, 0xa5, 0xd0, 0xca, 0x3e, 0x5d, 0xc9, 0x8a, 0x27, 0xd6,
	0xd6, 0xf4, 0x2f, 0xb6, 0xa5, 0x1f, 0xb3, 0xc5, 0xf3, 0x6a, 0x5e, 0x2d, 0x23, 0xaa, 0xec, 0xbc,
	0x96, 0xaa, 0x9c, 0x4f, 0x65, 0x6f, 0x59, 0x76, 0xce, 0x39, 0x66, 0x77, 0x26, 0x6c, 0x92, 0x65,
	0x04, 0xbd, 0x84, 0x7a, 0x10, 0x7b, 0x61, 0xea, 0xd3, 0xdb, 0x82, 0x27, 0xfa, 0xca, 0x1d, 0x17,
	0xaf, 0x9c, 0x89, 0x7e, 0x6b, 0x84, 0x67, 0x54, 0xcb, 0x32, 0x72, 0x8e, 0xef, 0xa1, 0x41, 0x63,
	0x32, 0x09, 0xa9, 0xef, 0x5e, 0x52, 0x22, 0x53, 0x4e, 0x05, 0x3e, 0xee, 0x94, 0x7a, 0xb5, 0x41,
	0xb3, 0x48, 0xf2, 0xda, 0x60, 0xa3, 0x7a, 0x16, 0x9c, 0xd9, 0xba, 0x07, 0x1e, 0xc4, 0xd3, 0x20,
	0x9e, 0xba, 0x32, 0x88, 0x28, 0x4b, 0x25, 0x7e, 0xd0, 0xb1, 0x7a, 0xe5, 0xc1, 0xb1, 0x63, 0xb4,
	0xd0, 0xc9, 0xb5, 0xd0, 0x79, 0x95, 0x69, 0xe1, 0xa8, 0x96, 0x65, 0x8c, 0x4d, 0x02, 0x3a, 0x83,
	0xc3, 0x88, 0xbc, 0x37, 0x62, 0x95, 0x0b, 0x26, 0xc6, 0xff, 0xc5, 0x52, 0x8f, 0xc8, 0x7b, 0xa5,
	0x66, 0xb9, 0xc3, 0x7e, 0x03, 0xf6, 0xfa, 0x45, 0x47, 0x0d, 0x28, 0x5d, 0xd3, 0x19, 0xb6, 0xf4,
	0x8e, 0xaa, 0x4f, 0xb5, 0xe3, 0xef, 0x48, 0x98, 0xd2, 0x4c, 0xdf, 0x8c, 0xf1, 0x7c, 0xe7, 0x3b,
	0xcb, 0x7e, 0x0e, 0x95, 0xe2, 0xe9, 0x6f, 0x95, 0xfb, 0x1a, 0xf0, 0xba, 0xe5, 0xd9, 0x96, 0x67,
	0xdd, 0x36, 0x6c, 0xc3, 0xd3, 0xfd, 0xcb, 0x82, 0x2f, 0x3e, 0xbd, 0x7a, 0x22, 0x61, 0xb1, 0xa0,
	0xe8, 0x09, 0xd4, 0xee, 0xea, 0x5f, 0xc6, 0x5f, 0xbd, 0xa3, 0x7c, 0x6b, 0x65, 0x72, 0x67, 0xbd,
	0x4c, 0x2e, 0x3c, 0x46, 0xa5, 0x85, 0xc7, 0xa8, 0xfb, 0xe7, 0x0e, 0x3c, 0xc9, 0x5b, 0xd4, 0x0f,
	0xc4, 0x25, 0xe5, 0xab, 0x5f, 0xbd, 0x05, 0x26, 0x6b, 0xf1, 0x59, 0x7b, 0x04, 0x65, 0x99, 0x11,
	0xa8, 0x57, 0xc3, 0xf4, 0x04, 0xb9, 0x6b, 0xcc, 0x96, 0x65, 0xa9, 0xb4, 0x42, 0x96, 0x7e, 0x9e,
	0xcb, 0xd2, 0x3d, 0x7d, 0x81, 0xbf, 0xbd, 0x73, 0x81, 0x3f, 0xd9, 0xe2, 0x6a, 0x61, 0xfa, 0x3f,
	0x1b, 0x35, 0xf8, 0xc7, 0x82, 0xf2, 0xc5, 0xf0, 0x3c, 0x2f, 0x8f, 0x6e, 0xe0, 0x68, 0xd5, 0x0f,
	0x89, 0x7a, 0x9b, 0xca, 0x8c, 0xfd, 0x74, 0x83, 0x48, 0xb3, 0x15, 0x5d, 0xf8, 0xf8, 0xc1, 0xda,
	0x6d, 0x58, 0x3f, 0x58, 0x5f, 0x59, 0x48, 0x40, 0x7b, 0xf5, 0xec, 0xe8, 0xcb, 0xcd, 0x0f, 0xc8,
	0x6e, 0x2f, 0x5d, 0xe9, 0x33, 0xf5, 0x0f, 0xaa, 0xdb, 0xfa, 0xf8, 0xc1, 0x3a, 0x6c, 0x58, 0x76,
	0x15, 0x15, 0x7f, 0xe3, 0x97, 0x9f, 0xff, 0xfa, 0x68, 0x1a, 0xc8, 0xab, 0x74, 0xe2, 0x78, 0x2c,
	0xea, 0x67, 0x92, 0x64, 0xfe, 0x60, 0x79, 0x2c, 0xec, 0xf3, 0xc4, 0x9b, 0xec, 0x6a, 0xeb, 0xd9,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x14, 0x7f, 0x97, 0xaf, 0x09, 0x00, 0x00,
}
