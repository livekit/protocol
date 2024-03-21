// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: infra/link.proto

package infra

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Link_WatchLocalLinks_FullMethodName   = "/rpc.Link/WatchLocalLinks"
	Link_SimulateLinkState_FullMethodName = "/rpc.Link/SimulateLinkState"
)

// LinkClient is the client API for Link service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkClient interface {
	WatchLocalLinks(ctx context.Context, in *WatchLocalLinksRequest, opts ...grpc.CallOption) (Link_WatchLocalLinksClient, error)
	SimulateLinkState(ctx context.Context, in *SimulateLinkStateRequest, opts ...grpc.CallOption) (*SimulateLinkStateResponse, error)
}

type linkClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkClient(cc grpc.ClientConnInterface) LinkClient {
	return &linkClient{cc}
}

func (c *linkClient) WatchLocalLinks(ctx context.Context, in *WatchLocalLinksRequest, opts ...grpc.CallOption) (Link_WatchLocalLinksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Link_ServiceDesc.Streams[0], Link_WatchLocalLinks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &linkWatchLocalLinksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Link_WatchLocalLinksClient interface {
	Recv() (*WatchLocalLinksResponse, error)
	grpc.ClientStream
}

type linkWatchLocalLinksClient struct {
	grpc.ClientStream
}

func (x *linkWatchLocalLinksClient) Recv() (*WatchLocalLinksResponse, error) {
	m := new(WatchLocalLinksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *linkClient) SimulateLinkState(ctx context.Context, in *SimulateLinkStateRequest, opts ...grpc.CallOption) (*SimulateLinkStateResponse, error) {
	out := new(SimulateLinkStateResponse)
	err := c.cc.Invoke(ctx, Link_SimulateLinkState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServer is the server API for Link service.
// All implementations must embed UnimplementedLinkServer
// for forward compatibility
type LinkServer interface {
	WatchLocalLinks(*WatchLocalLinksRequest, Link_WatchLocalLinksServer) error
	SimulateLinkState(context.Context, *SimulateLinkStateRequest) (*SimulateLinkStateResponse, error)
	mustEmbedUnimplementedLinkServer()
}

// UnimplementedLinkServer must be embedded to have forward compatible implementations.
type UnimplementedLinkServer struct {
}

func (UnimplementedLinkServer) WatchLocalLinks(*WatchLocalLinksRequest, Link_WatchLocalLinksServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchLocalLinks not implemented")
}
func (UnimplementedLinkServer) SimulateLinkState(context.Context, *SimulateLinkStateRequest) (*SimulateLinkStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimulateLinkState not implemented")
}
func (UnimplementedLinkServer) mustEmbedUnimplementedLinkServer() {}

// UnsafeLinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkServer will
// result in compilation errors.
type UnsafeLinkServer interface {
	mustEmbedUnimplementedLinkServer()
}

func RegisterLinkServer(s grpc.ServiceRegistrar, srv LinkServer) {
	s.RegisterService(&Link_ServiceDesc, srv)
}

func _Link_WatchLocalLinks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchLocalLinksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LinkServer).WatchLocalLinks(m, &linkWatchLocalLinksServer{stream})
}

type Link_WatchLocalLinksServer interface {
	Send(*WatchLocalLinksResponse) error
	grpc.ServerStream
}

type linkWatchLocalLinksServer struct {
	grpc.ServerStream
}

func (x *linkWatchLocalLinksServer) Send(m *WatchLocalLinksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Link_SimulateLinkState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateLinkStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).SimulateLinkState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Link_SimulateLinkState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).SimulateLinkState(ctx, req.(*SimulateLinkStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Link_ServiceDesc is the grpc.ServiceDesc for Link service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Link_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Link",
	HandlerType: (*LinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SimulateLinkState",
			Handler:    _Link_SimulateLinkState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchLocalLinks",
			Handler:       _Link_WatchLocalLinks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "infra/link.proto",
}
