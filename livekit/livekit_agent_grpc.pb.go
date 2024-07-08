// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: livekit_agent.proto

package livekit

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
	AgentService_CreateDispatch_FullMethodName = "/livekit.AgentService/CreateDispatch"
	AgentService_DeleteDispatch_FullMethodName = "/livekit.AgentService/DeleteDispatch"
	AgentService_ListDispatch_FullMethodName   = "/livekit.AgentService/ListDispatch"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	CreateDispatch(ctx context.Context, in *CreateAgentDispatchRequest, opts ...grpc.CallOption) (*AgentDispatch, error)
	DeleteDispatch(ctx context.Context, in *DeleteAgentDispatchRequest, opts ...grpc.CallOption) (*AgentDispatch, error)
	ListDispatch(ctx context.Context, in *ListAgentDispatchRequesst, opts ...grpc.CallOption) (*ListAgentDispatchResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) CreateDispatch(ctx context.Context, in *CreateAgentDispatchRequest, opts ...grpc.CallOption) (*AgentDispatch, error) {
	out := new(AgentDispatch)
	err := c.cc.Invoke(ctx, AgentService_CreateDispatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) DeleteDispatch(ctx context.Context, in *DeleteAgentDispatchRequest, opts ...grpc.CallOption) (*AgentDispatch, error) {
	out := new(AgentDispatch)
	err := c.cc.Invoke(ctx, AgentService_DeleteDispatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ListDispatch(ctx context.Context, in *ListAgentDispatchRequesst, opts ...grpc.CallOption) (*ListAgentDispatchResponse, error) {
	out := new(ListAgentDispatchResponse)
	err := c.cc.Invoke(ctx, AgentService_ListDispatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	CreateDispatch(context.Context, *CreateAgentDispatchRequest) (*AgentDispatch, error)
	DeleteDispatch(context.Context, *DeleteAgentDispatchRequest) (*AgentDispatch, error)
	ListDispatch(context.Context, *ListAgentDispatchRequesst) (*ListAgentDispatchResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) CreateDispatch(context.Context, *CreateAgentDispatchRequest) (*AgentDispatch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDispatch not implemented")
}
func (UnimplementedAgentServiceServer) DeleteDispatch(context.Context, *DeleteAgentDispatchRequest) (*AgentDispatch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDispatch not implemented")
}
func (UnimplementedAgentServiceServer) ListDispatch(context.Context, *ListAgentDispatchRequesst) (*ListAgentDispatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDispatch not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_CreateDispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentDispatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).CreateDispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_CreateDispatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).CreateDispatch(ctx, req.(*CreateAgentDispatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_DeleteDispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentDispatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).DeleteDispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_DeleteDispatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).DeleteDispatch(ctx, req.(*DeleteAgentDispatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ListDispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentDispatchRequesst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListDispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_ListDispatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListDispatch(ctx, req.(*ListAgentDispatchRequesst))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "livekit.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDispatch",
			Handler:    _AgentService_CreateDispatch_Handler,
		},
		{
			MethodName: "DeleteDispatch",
			Handler:    _AgentService_DeleteDispatch_Handler,
		},
		{
			MethodName: "ListDispatch",
			Handler:    _AgentService_ListDispatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "livekit_agent.proto",
}
