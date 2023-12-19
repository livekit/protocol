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
	UpdateSIPParticipant(ctx context.Context, req *InternalUpdateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalUpdateSIPParticipantResponse, error)

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

func (c *sIPInternalClient) UpdateSIPParticipant(ctx context.Context, req *InternalUpdateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalUpdateSIPParticipantResponse, error) {
	return client.RequestSingle[*InternalUpdateSIPParticipantResponse](ctx, c.client, "UpdateSIPParticipant", nil, req, opts...)
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
	var err error
	err = server.RegisterHandler(s, "UpdateSIPParticipant", nil, svc.UpdateSIPParticipant, svc.UpdateSIPParticipantAffinity)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("SendSIPParticipantDTMF", true, false, true, false)
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

func (s *sIPInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *sIPInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor3 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x8e, 0xda, 0x30,
	0x14, 0x85, 0x15, 0xa0, 0x01, 0x6e, 0x45, 0x17, 0x2e, 0xa2, 0x56, 0xba, 0xa9, 0xd2, 0x82, 0x28,
	0xad, 0x92, 0xfe, 0xbc, 0x40, 0x55, 0x55, 0x95, 0xb2, 0x68, 0x85, 0x80, 0x6e, 0xba, 0x41, 0xc1,
	0x76, 0xa9, 0x45, 0x62, 0xbb, 0xb6, 0x43, 0x35, 0x9a, 0xd5, 0x2c, 0xe7, 0x29, 0xe6, 0x1d, 0xe6,
	0x09, 0x47, 0x49, 0x08, 0x33, 0x19, 0x21, 0x86, 0x59, 0x45, 0xe7, 0x7c, 0xc7, 0xf1, 0xbd, 0x47,
	0x09, 0xf4, 0xb4, 0x22, 0xa1, 0xe1, 0x2a, 0x50, 0x5a, 0x5a, 0x89, 0x9a, 0x5a, 0x11, 0xaf, 0x27,
	0x95, 0xe5, 0x52, 0x98, 0xd2, 0xf3, 0xaf, 0x1a, 0xf0, 0x3a, 0x12, 0x96, 0x69, 0x11, 0x27, 0xbf,
	0x14, 0x8d, 0x2d, 0x9b, 0x47, 0xd3, 0x69, 0xac, 0x2d, 0x27, 0x5c, 0xc5, 0xc2, 0xce, 0xd8, 0xbf,
	0x8c, 0x19, 0x8b, 0x86, 0xf0, 0x4c, 0xdd, 0xba, 0x4b, 0x4e, 0xb1, 0xf3, 0xca, 0x19, 0x77, 0x67,
	0xbd, 0x3b, 0x6e, 0x44, 0x11, 0x86, 0x76, 0x4c, 0xa9, 0x66, 0xc6, 0xe0, 0x46, 0xc1, 0x2b, 0x89,
	0x06, 0xe0, 0x8a, 0x2c, 0x5d, 0x31, 0x8d, 0x9b, 0x05, 0xd8, 0x29, 0xf4, 0x02, 0xda, 0x24, 0x4e,
	0x92, 0xa5, 0x95, 0xb8, 0x55, 0x82, 0x5c, 0x2e, 0x24, 0xf2, 0xa0, 0x93, 0x99, 0x7c, 0xae, 0x94,
	0xe1, 0x27, 0x05, 0xd9, 0xeb, 0x9c, 0xa9, 0xd8, 0x98, 0xff, 0x52, 0x53, 0xec, 0x96, 0xac, 0xd2,
	0xe8, 0x25, 0x74, 0xb5, 0x94, 0xe9, 0xb2, 0x38, 0xd8, 0x2e, 0x61, 0x6e, 0xfc, 0xcc, 0x0f, 0x7e,
	0x84, 0x7e, 0x7d, 0x0d, 0x26, 0x2c, 0xb7, 0x67, 0xb8, 0x53, 0xe4, 0x9e, 0xd7, 0x96, 0x29, 0x91,
	0x3f, 0x82, 0x37, 0xc7, 0x0b, 0x32, 0x4a, 0x0a, 0xc3, 0xfc, 0x3f, 0x30, 0xac, 0x72, 0x73, 0x26,
	0x68, 0x3d, 0xf5, 0x6d, 0xf1, 0xe3, 0xfb, 0x23, 0xab, 0x1c, 0x80, 0x4b, 0xf9, 0x9a, 0xdb, 0xaa,
	0xc9, 0x9d, 0xf2, 0xc7, 0x30, 0x7a, 0xe8, 0x9e, 0x72, 0xa2, 0x4f, 0x17, 0x0d, 0x78, 0x3a, 0x8f,
	0xa6, 0x55, 0x1a, 0x65, 0xd0, 0x3f, 0xb4, 0x01, 0x1a, 0x07, 0x5a, 0x91, 0xe0, 0x84, 0xaf, 0xc0,
	0x7b, 0x7b, 0x42, 0x72, 0x57, 0x47, 0xe7, 0xfa, 0xd2, 0x69, 0x7d, 0x71, 0x3e, 0x38, 0xe8, 0x1c,
	0x06, 0x87, 0x07, 0x45, 0x93, 0xda, 0xeb, 0x8e, 0xb6, 0xe6, 0xbd, 0x3b, 0x29, 0x7b, 0xff, 0xf2,
	0xaf, 0xef, 0x7f, 0x4f, 0xd6, 0xdc, 0xfe, 0xcd, 0x56, 0x01, 0x91, 0x69, 0x98, 0xf0, 0x2d, 0xdb,
	0x70, 0xbb, 0x7f, 0xaa, 0xcd, 0x3a, 0x34, 0x4c, 0x6f, 0x39, 0x61, 0xa1, 0x56, 0x64, 0xe5, 0x16,
	0x3f, 0xc5, 0xe7, 0x9b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x70, 0xb3, 0x81, 0x39, 0x03, 0x00,
	0x00,
}
