// Code generated by protoc-gen-psrpc v0.5.0, DO NOT EDIT.
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
	UpdateSIPParticipant(ctx context.Context, topic string, req *InternalUpdateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalUpdateSIPParticipantResponse, error)

	SendSIPParticipantDTMF(ctx context.Context, req *InternalSendSIPParticipantDTMFRequest, opts ...psrpc.RequestOption) (*InternalSendSIPParticipantDTMFResponse, error)
}

// ================================
// SIPInternal ServerImpl Interface
// ================================

type SIPInternalServerImpl interface {
	UpdateSIPParticipant(context.Context, *InternalUpdateSIPParticipantRequest) (*InternalUpdateSIPParticipantResponse, error)
	UpdateSIPParticipantAffinity(context.Context, *InternalUpdateSIPParticipantRequest) float32

	SendSIPParticipantDTMF(context.Context, *InternalSendSIPParticipantDTMFRequest) (*InternalSendSIPParticipantDTMFResponse, error)
	SendSIPParticipantDTMFAffinity(context.Context, *InternalSendSIPParticipantDTMFRequest) float32
}

// ============================
// SIPInternal Server Interface
// ============================

type SIPInternalServer interface {
	RegisterUpdateSIPParticipantTopic(topic string) error
	DeregisterUpdateSIPParticipantTopic(topic string)

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

	sd.RegisterMethod("UpdateSIPParticipant", true, false, true, false)
	sd.RegisterMethod("SendSIPParticipantDTMF", true, false, true, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &sIPInternalClient{
		client: rpcClient,
	}, nil
}

func (c *sIPInternalClient) UpdateSIPParticipant(ctx context.Context, topic string, req *InternalUpdateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalUpdateSIPParticipantResponse, error) {
	return client.RequestSingle[*InternalUpdateSIPParticipantResponse](ctx, c.client, "UpdateSIPParticipant", []string{topic}, req, opts...)
}

func (c *sIPInternalClient) SendSIPParticipantDTMF(ctx context.Context, req *InternalSendSIPParticipantDTMFRequest, opts ...psrpc.RequestOption) (*InternalSendSIPParticipantDTMFResponse, error) {
	return client.RequestSingle[*InternalSendSIPParticipantDTMFResponse](ctx, c.client, "SendSIPParticipantDTMF", nil, req, opts...)
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

	sd.RegisterMethod("UpdateSIPParticipant", true, false, true, false)
	sd.RegisterMethod("SendSIPParticipantDTMF", true, false, true, false)
	var err error
	err = server.RegisterHandler(s, "SendSIPParticipantDTMF", nil, svc.SendSIPParticipantDTMF, svc.SendSIPParticipantDTMFAffinity)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &sIPInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *sIPInternalServer) RegisterUpdateSIPParticipantTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "UpdateSIPParticipant", []string{topic}, s.svc.UpdateSIPParticipant, s.svc.UpdateSIPParticipantAffinity)
}

func (s *sIPInternalServer) DeregisterUpdateSIPParticipantTopic(topic string) {
	s.rpc.DeregisterHandler("UpdateSIPParticipant", []string{topic})
}

func (s *sIPInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *sIPInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor3 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0xee, 0xd2, 0x40,
	0x14, 0xc5, 0x53, 0xfe, 0x58, 0xe0, 0x1a, 0x8c, 0x19, 0x09, 0x4e, 0xea, 0xc6, 0x54, 0x21, 0x88,
	0xa6, 0xf5, 0xe3, 0x05, 0x8c, 0x31, 0x26, 0x5d, 0x68, 0x08, 0xe0, 0xc6, 0x0d, 0x29, 0x33, 0x23,
	0x4e, 0x68, 0x67, 0xc6, 0x99, 0x29, 0xc4, 0xb8, 0x74, 0xc5, 0x53, 0xf8, 0x0e, 0x3c, 0xa1, 0xe9,
	0x17, 0xda, 0x84, 0x20, 0xff, 0x55, 0x73, 0xee, 0xef, 0xdc, 0xde, 0x7b, 0x4f, 0x5a, 0xe8, 0x6b,
	0x45, 0x42, 0xc3, 0x55, 0xa0, 0xb4, 0xb4, 0x12, 0xdd, 0x68, 0x45, 0xbc, 0xbe, 0x54, 0x96, 0x4b,
	0x61, 0xca, 0x9a, 0xff, 0xbb, 0x05, 0x4f, 0x22, 0x61, 0x99, 0x16, 0x71, 0xf2, 0x59, 0xd1, 0xd8,
	0xb2, 0x45, 0x34, 0x9b, 0xc5, 0xda, 0x72, 0xc2, 0x55, 0x2c, 0xec, 0x9c, 0x7d, 0xcf, 0x98, 0xb1,
	0x68, 0x04, 0xf7, 0xd4, 0xdf, 0xea, 0x8a, 0x53, 0xec, 0x3c, 0x76, 0x26, 0xbd, 0x79, 0xff, 0x9f,
	0x6a, 0x44, 0x11, 0x86, 0x4e, 0x4c, 0xa9, 0x66, 0xc6, 0xe0, 0x56, 0xc1, 0x6b, 0x89, 0x86, 0xe0,
	0x8a, 0x2c, 0x5d, 0x33, 0x8d, 0x6f, 0x0a, 0x50, 0x29, 0xf4, 0x10, 0x3a, 0x24, 0x4e, 0x92, 0x95,
	0x95, 0xb8, 0x5d, 0x82, 0x5c, 0x2e, 0x25, 0xf2, 0xa0, 0x9b, 0x99, 0x7c, 0xaf, 0x94, 0xe1, 0x3b,
	0x05, 0x39, 0xe9, 0x9c, 0xa9, 0xd8, 0x98, 0xbd, 0xd4, 0x14, 0xbb, 0x25, 0xab, 0x35, 0x7a, 0x04,
	0x3d, 0x2d, 0x65, 0xba, 0x2a, 0x1a, 0x3b, 0x25, 0xcc, 0x0b, 0x9f, 0xf2, 0xc6, 0x57, 0x30, 0x68,
	0x9e, 0xc1, 0x84, 0xe5, 0xf6, 0x07, 0xee, 0x16, 0xbe, 0x07, 0x8d, 0x63, 0x4a, 0xe4, 0x8f, 0xe1,
	0xe9, 0xe5, 0x80, 0x8c, 0x92, 0xc2, 0x30, 0xff, 0x2b, 0x8c, 0x6a, 0xdf, 0x82, 0x09, 0xda, 0x74,
	0xbd, 0x5f, 0x7e, 0xfc, 0x70, 0xcb, 0x28, 0x87, 0xe0, 0x52, 0xbe, 0xe1, 0xb6, 0x4e, 0xb2, 0x52,
	0xfe, 0x04, 0xc6, 0xff, 0x9b, 0x53, 0x6e, 0xf4, 0xfa, 0x57, 0x0b, 0xee, 0x2e, 0xa2, 0x59, 0xed,
	0x46, 0x7b, 0x18, 0x9c, 0xbb, 0x00, 0x4d, 0x02, 0xad, 0x48, 0x70, 0xc5, 0x57, 0xe0, 0x3d, 0xbb,
	0xc2, 0x59, 0xc5, 0x01, 0xc7, 0x83, 0xe3, 0xde, 0x77, 0xde, 0x3a, 0x2f, 0x1d, 0xf4, 0x13, 0x86,
	0xe7, 0x57, 0x45, 0xd3, 0xc6, 0x0b, 0x2f, 0xe6, 0xe6, 0x3d, 0xbf, 0xca, 0x5b, 0x8d, 0xef, 0x1e,
	0x0f, 0x4e, 0x3b, 0x1f, 0xfe, 0xee, 0xc5, 0x97, 0xe9, 0x86, 0xdb, 0x6f, 0xd9, 0x3a, 0x20, 0x32,
	0x0d, 0x13, 0xbe, 0x63, 0x5b, 0x6e, 0x4f, 0x4f, 0xb5, 0xdd, 0x84, 0x86, 0xe9, 0x1d, 0x27, 0x2c,
	0xd4, 0x8a, 0xac, 0xdd, 0xe2, 0xb7, 0x78, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x9b, 0x7b,
	0x2f, 0x3b, 0x03, 0x00, 0x00,
}
