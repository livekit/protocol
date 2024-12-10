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
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x72, 0xdb, 0x36,
	0x10, 0x1e, 0x5a, 0x89, 0x6c, 0xad, 0x7e, 0x0d, 0x59, 0x0a, 0x4c, 0x4f, 0x13, 0x55, 0x69, 0x66,
	0x94, 0x1e, 0xa8, 0x56, 0x99, 0xce, 0x74, 0x72, 0xe8, 0x34, 0x71, 0x9c, 0x89, 0x0e, 0x6d, 0x3c,
	0xb2, 0x7a, 0xe9, 0x85, 0x03, 0x91, 0xb0, 0xcc, 0x9a, 0x24, 0x58, 0x00, 0x8c, 0xa3, 0xbe, 0x41,
	0x0e, 0x7d, 0x87, 0x3e, 0x43, 0x1e, 0xa4, 0x2f, 0xd4, 0x4b, 0x07, 0x00, 0x69, 0xd3, 0xfa, 0x71,
	0xac, 0x1b, 0x77, 0xbf, 0xdd, 0x0f, 0xbb, 0xab, 0xc5, 0x07, 0x41, 0x9d, 0x27, 0xde, 0x50, 0x04,
	0x89, 0x93, 0x70, 0x26, 0x19, 0x2a, 0xf1, 0xc4, 0xb3, 0x8f, 0xe6, 0x8c, 0xcd, 0x43, 0x3a, 0xd4,
	0xae, 0x59, 0x7a, 0x3e, 0xa4, 0x51, 0x22, 0x17, 0x26, 0xc2, 0x7e, 0xbc, 0x0c, 0xfa, 0x29, 0x27,
	0x32, 0x60, 0x71, 0x86, 0xd7, 0x59, 0xa2, 0x2c, 0x91, 0x99, 0xfb, 0x61, 0xf0, 0x81, 0x5e, 0x06,
	0xd2, 0xbd, 0x3e, 0xa3, 0xff, 0x6f, 0x0d, 0x9e, 0x8e, 0x63, 0x49, 0x79, 0x4c, 0xc2, 0x63, 0x4e,
	0x89, 0xa4, 0x67, 0xe3, 0xd3, 0x53, 0xc2, 0x65, 0xe0, 0x05, 0x09, 0x89, 0xe5, 0x84, 0xfe, 0x99,
	0x52, 0x21, 0xd1, 0x57, 0x00, 0x09, 0x67, 0x7f, 0x50, 0x4f, 0xba, 0x81, 0x8f, 0x51, 0xcf, 0x1a,
	0x54, 0x26, 0x95, 0xcc, 0x33, 0xf6, 0xd1, 0x63, 0xa8, 0x8a, 0x20, 0x71, 0x3d, 0x12, 0x86, 0x0a,
	0xaf, 0x1b, 0x5c, 0x04, 0xc9, 0x31, 0x09, 0xc3, 0xb1, 0x8f, 0x7a, 0x50, 0x53, 0xb8, 0xe4, 0x69,
	0x7c, 0xa9, 0x02, 0xda, 0x3a, 0x00, 0x44, 0x90, 0x4c, 0x95, 0x6b, 0xec, 0x23, 0x0c, 0xbb, 0xc4,
	0xf7, 0x39, 0x15, 0x02, 0xef, 0x68, 0x30, 0x37, 0x91, 0x0d, 0x7b, 0x17, 0x4c, 0xc8, 0x98, 0x44,
	0x14, 0x1f, 0x68, 0xe8, 0xda, 0x46, 0x2f, 0xa0, 0x22, 0x39, 0x89, 0x45, 0xc2, 0xb8, 0xc4, 0xad,
	0x9e, 0x35, 0x68, 0x8c, 0x3a, 0x4e, 0xd6, 0xa5, 0x73, 0x36, 0x3e, 0x9d, 0xe6, 0xe0, 0xe4, 0x26,
	0x0e, 0x75, 0xa1, 0x1c, 0xa7, 0xd1, 0x8c, 0x72, 0x5c, 0xd2, 0x74, 0x99, 0x85, 0x1e, 0xc1, 0xae,
	0x6e, 0x40, 0x32, 0xfc, 0xc0, 0x00, 0xca, 0x9c, 0x32, 0x55, 0x41, 0x2a, 0xd4, 0x88, 0x22, 0x8a,
	0x1f, 0x9a, 0x0a, 0x72, 0x5b, 0x61, 0x09, 0x11, 0xe2, 0x8a, 0x71, 0x1f, 0x97, 0x0d, 0x96, 0xdb,
	0xe8, 0x08, 0x2a, 0x9c, 0xb1, 0xc8, 0xd5, 0x89, 0xbb, 0x06, 0x54, 0x8e, 0x5f, 0x55, 0xe2, 0xf7,
	0x70, 0x90, 0xdc, 0xcc, 0xd9, 0x0d, 0x7c, 0x1a, 0xcb, 0x40, 0x2e, 0xf0, 0x9e, 0x8e, 0x6b, 0x17,
	0xb0, 0x71, 0x06, 0xa1, 0xe7, 0xd0, 0x2a, 0xa6, 0x68, 0xda, 0x86, 0x0e, 0x6f, 0x16, 0xfc, 0xeb,
	0xd8, 0x23, 0x2a, 0x89, 0x4f, 0x24, 0xc1, 0xcd, 0x15, 0xf6, 0x5f, 0x32, 0x08, 0xfd, 0x05, 0xdd,
	0x62, 0x0a, 0x91, 0x92, 0x07, 0xb3, 0x54, 0x52, 0x81, 0xf7, 0x7b, 0xa5, 0x41, 0x75, 0x74, 0xec,
	0xf0, 0xc4, 0x73, 0xee, 0xb1, 0x2c, 0x4e, 0xc1, 0xf5, 0xea, 0x9a, 0xe5, 0x24, 0x96, 0x7c, 0x31,
	0xe9, 0x24, 0xeb, 0x30, 0x74, 0x00, 0x0f, 0x25, 0xbb, 0xa4, 0x31, 0xae, 0xe8, 0xfa, 0x8c, 0x81,
	0x3a, 0x50, 0xbe, 0x12, 0x6e, 0xca, 0x43, 0x0c, 0xc6, 0x7d, 0x25, 0x7e, 0xe3, 0x21, 0x42, 0xf0,
	0xc0, 0x97, 0xd1, 0x39, 0xae, 0x6a, 0xa7, 0xfe, 0x46, 0x4f, 0xa1, 0x9e, 0x84, 0x64, 0xe1, 0xfa,
	0x01, 0x09, 0x25, 0x8b, 0x29, 0xae, 0xf5, 0xac, 0xc1, 0xde, 0xa4, 0xa6, 0x9c, 0x6f, 0x32, 0x1f,
	0x7a, 0x0f, 0xbb, 0x17, 0x94, 0xf8, 0x94, 0x0b, 0xdc, 0xd1, 0x2d, 0xfd, 0x70, 0xef, 0x96, 0xde,
	0x99, 0x3c, 0xd3, 0x44, 0xce, 0x82, 0x52, 0xe8, 0x64, 0x9f, 0xae, 0x64, 0xc5, 0x89, 0x75, 0x35,
	0xfd, 0xab, 0x6d, 0xe9, 0xa7, 0x6c, 0x79, 0x5e, 0xed, 0x8b, 0x55, 0x44, 0x1d, 0x7b, 0x73, 0x96,
	0x3a, 0x39, 0xef, 0xca, 0xde, 0xf2, 0xd8, 0x1b, 0xce, 0x29, 0xbb, 0xd5, 0x61, 0x9b, 0xac, 0x22,
	0xe8, 0x35, 0x34, 0x83, 0xd8, 0x0b, 0x53, 0x9f, 0x5e, 0x1f, 0x78, 0xa4, 0xaf, 0xdc, 0x61, 0xf1,
	0xca, 0x99, 0xe8, 0xf7, 0x46, 0x78, 0x26, 0x8d, 0x2c, 0x23, 0xe7, 0xf8, 0x09, 0x5a, 0x34, 0x26,
	0xb3, 0x90, 0xfa, 0xee, 0x39, 0x25, 0x32, 0xe5, 0x54, 0xe0, 0xc3, 0x5e, 0x69, 0xd0, 0x18, 0xb5,
	0x8b, 0x24, 0x6f, 0x0d, 0x36, 0x69, 0x66, 0xc1, 0x99, 0xad, 0x6b, 0xe0, 0x41, 0x3c, 0x0f, 0xe2,
	0xb9, 0x2b, 0x83, 0x88, 0xb2, 0x54, 0xe2, 0x47, 0x3d, 0x6b, 0x50, 0x1d, 0x1d, 0x3a, 0x46, 0x0b,
	0x9d, 0x5c, 0x0b, 0x9d, 0x37, 0x99, 0x16, 0x4e, 0x1a, 0x59, 0xc6, 0xd4, 0x24, 0xa0, 0x13, 0xd8,
	0x8f, 0xc8, 0x47, 0x23, 0x56, 0xb9, 0x60, 0x62, 0xfc, 0x25, 0x96, 0x66, 0x44, 0x3e, 0x2a, 0x35,
	0xcb, 0x1d, 0xf6, 0x3b, 0xb0, 0x37, 0x2f, 0x3a, 0x6a, 0x41, 0xe9, 0x92, 0x2e, 0xb0, 0xa5, 0x77,
	0x54, 0x7d, 0xaa, 0x1d, 0xff, 0x40, 0xc2, 0x94, 0x66, 0xfa, 0x66, 0x8c, 0x97, 0x3b, 0x3f, 0x5a,
	0xf6, 0x4b, 0xa8, 0x15, 0xa7, 0xbf, 0x55, 0xee, 0x5b, 0xc0, 0x9b, 0x96, 0x67, 0x5b, 0x9e, 0x4d,
	0xdb, 0xb0, 0x0d, 0x4f, 0xff, 0x1f, 0x0b, 0xbe, 0xb9, 0x7b, 0xf5, 0x44, 0xc2, 0x62, 0x41, 0xd1,
	0x33, 0x68, 0xdc, 0xd6, 0xbf, 0x8c, 0xbf, 0x7e, 0x4b, 0xf9, 0x36, 0xca, 0xe4, 0xce, 0x66, 0x99,
	0x5c, 0x7a, 0x8c, 0x4a, 0x4b, 0x8f, 0x51, 0xff, 0x6f, 0x0b, 0x9e, 0xe5, 0x25, 0xea, 0x07, 0xe2,
	0x9c, 0xf2, 0xf5, 0xaf, 0xde, 0x12, 0x93, 0xb5, 0xfc, 0xac, 0x3d, 0x81, 0xaa, 0xcc, 0x08, 0xd4,
	0xab, 0x61, 0x6a, 0x82, 0xdc, 0x35, 0x65, 0xab, 0xb2, 0x54, 0x5a, 0x95, 0xa5, 0xd1, 0x7f, 0x16,
	0x54, 0xcf, 0xc6, 0xa7, 0x79, 0x49, 0xe8, 0x0a, 0x0e, 0xd6, 0x4d, 0x0e, 0x0d, 0xee, 0x7b, 0xaf,
	0xed, 0xe7, 0xf7, 0x88, 0x34, 0x3f, 0x43, 0x1f, 0x3e, 0x7f, 0xb2, 0xca, 0x2d, 0xeb, 0x67, 0xeb,
	0x3b, 0x0b, 0x09, 0xe8, 0xae, 0x9f, 0x07, 0xfa, 0xf6, 0x16, 0xe1, 0x9d, 0x43, 0xb3, 0xbb, 0x2b,
	0x77, 0xe8, 0x44, 0xfd, 0x65, 0xe9, 0x77, 0x3e, 0x7f, 0xb2, 0xf6, 0x5b, 0x96, 0x5d, 0x47, 0xc5,
	0xa1, 0xbe, 0xfe, 0xfa, 0xf7, 0x27, 0xf3, 0x40, 0x5e, 0xa4, 0x33, 0xc7, 0x63, 0xd1, 0x30, 0xd3,
	0x00, 0xf3, 0x8f, 0xc6, 0x63, 0xe1, 0x90, 0x27, 0xde, 0xac, 0xac, 0xad, 0x17, 0xff, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x33, 0x3c, 0x94, 0x37, 0x20, 0x09, 0x00, 0x00,
}
