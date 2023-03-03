// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: livekit_analytics.proto

package livekit

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AnalyticsRecorderServiceClient is the client API for AnalyticsRecorderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsRecorderServiceClient interface {
	IngestStats(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestStatsClient, error)
	IngestEvents(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestEventsClient, error)
}

type analyticsRecorderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsRecorderServiceClient(cc grpc.ClientConnInterface) AnalyticsRecorderServiceClient {
	return &analyticsRecorderServiceClient{cc}
}

func (c *analyticsRecorderServiceClient) IngestStats(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AnalyticsRecorderService_ServiceDesc.Streams[0], "/livekit.AnalyticsRecorderService/IngestStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &analyticsRecorderServiceIngestStatsClient{stream}
	return x, nil
}

type AnalyticsRecorderService_IngestStatsClient interface {
	Send(*AnalyticsStats) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type analyticsRecorderServiceIngestStatsClient struct {
	grpc.ClientStream
}

func (x *analyticsRecorderServiceIngestStatsClient) Send(m *AnalyticsStats) error {
	return x.ClientStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestStatsClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *analyticsRecorderServiceClient) IngestEvents(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AnalyticsRecorderService_ServiceDesc.Streams[1], "/livekit.AnalyticsRecorderService/IngestEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &analyticsRecorderServiceIngestEventsClient{stream}
	return x, nil
}

type AnalyticsRecorderService_IngestEventsClient interface {
	Send(*AnalyticsEvents) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type analyticsRecorderServiceIngestEventsClient struct {
	grpc.ClientStream
}

func (x *analyticsRecorderServiceIngestEventsClient) Send(m *AnalyticsEvents) error {
	return x.ClientStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestEventsClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnalyticsRecorderServiceServer is the server API for AnalyticsRecorderService service.
// All implementations must embed UnimplementedAnalyticsRecorderServiceServer
// for forward compatibility
type AnalyticsRecorderServiceServer interface {
	IngestStats(AnalyticsRecorderService_IngestStatsServer) error
	IngestEvents(AnalyticsRecorderService_IngestEventsServer) error
	mustEmbedUnimplementedAnalyticsRecorderServiceServer()
}

// UnimplementedAnalyticsRecorderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsRecorderServiceServer struct {
}

func (UnimplementedAnalyticsRecorderServiceServer) IngestStats(AnalyticsRecorderService_IngestStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method IngestStats not implemented")
}
func (UnimplementedAnalyticsRecorderServiceServer) IngestEvents(AnalyticsRecorderService_IngestEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method IngestEvents not implemented")
}
func (UnimplementedAnalyticsRecorderServiceServer) mustEmbedUnimplementedAnalyticsRecorderServiceServer() {
}

// UnsafeAnalyticsRecorderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsRecorderServiceServer will
// result in compilation errors.
type UnsafeAnalyticsRecorderServiceServer interface {
	mustEmbedUnimplementedAnalyticsRecorderServiceServer()
}

func RegisterAnalyticsRecorderServiceServer(s grpc.ServiceRegistrar, srv AnalyticsRecorderServiceServer) {
	s.RegisterService(&AnalyticsRecorderService_ServiceDesc, srv)
}

func _AnalyticsRecorderService_IngestStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnalyticsRecorderServiceServer).IngestStats(&analyticsRecorderServiceIngestStatsServer{stream})
}

type AnalyticsRecorderService_IngestStatsServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*AnalyticsStats, error)
	grpc.ServerStream
}

type analyticsRecorderServiceIngestStatsServer struct {
	grpc.ServerStream
}

func (x *analyticsRecorderServiceIngestStatsServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestStatsServer) Recv() (*AnalyticsStats, error) {
	m := new(AnalyticsStats)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AnalyticsRecorderService_IngestEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnalyticsRecorderServiceServer).IngestEvents(&analyticsRecorderServiceIngestEventsServer{stream})
}

type AnalyticsRecorderService_IngestEventsServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*AnalyticsEvents, error)
	grpc.ServerStream
}

type analyticsRecorderServiceIngestEventsServer struct {
	grpc.ServerStream
}

func (x *analyticsRecorderServiceIngestEventsServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestEventsServer) Recv() (*AnalyticsEvents, error) {
	m := new(AnalyticsEvents)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnalyticsRecorderService_ServiceDesc is the grpc.ServiceDesc for AnalyticsRecorderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsRecorderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "livekit.AnalyticsRecorderService",
	HandlerType: (*AnalyticsRecorderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "IngestStats",
			Handler:       _AnalyticsRecorderService_IngestStats_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "IngestEvents",
			Handler:       _AnalyticsRecorderService_IngestEvents_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "livekit_analytics.proto",
}
