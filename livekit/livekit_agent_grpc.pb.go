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
// - protoc             v5.27.0
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
	AgentService_CreateAgentJobDefinition_FullMethodName = "/livekit.AgentService/CreateAgentJobDefinition"
	AgentService_DeleteAgentJobDefinition_FullMethodName = "/livekit.AgentService/DeleteAgentJobDefinition"
	AgentService_ListAgentJobDefinitions_FullMethodName  = "/livekit.AgentService/ListAgentJobDefinitions"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	CreateAgentJobDefinition(ctx context.Context, in *CreateAgentJobDefinitionRequest, opts ...grpc.CallOption) (*JobDefinition, error)
	DeleteAgentJobDefinition(ctx context.Context, in *DeleteAgentJobDefinitionRequest, opts ...grpc.CallOption) (*JobDefinition, error)
	ListAgentJobDefinitions(ctx context.Context, in *ListAgentJobDefinitionsRequesst, opts ...grpc.CallOption) (*ListAgentJobDefinitionsResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) CreateAgentJobDefinition(ctx context.Context, in *CreateAgentJobDefinitionRequest, opts ...grpc.CallOption) (*JobDefinition, error) {
	out := new(JobDefinition)
	err := c.cc.Invoke(ctx, AgentService_CreateAgentJobDefinition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) DeleteAgentJobDefinition(ctx context.Context, in *DeleteAgentJobDefinitionRequest, opts ...grpc.CallOption) (*JobDefinition, error) {
	out := new(JobDefinition)
	err := c.cc.Invoke(ctx, AgentService_DeleteAgentJobDefinition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ListAgentJobDefinitions(ctx context.Context, in *ListAgentJobDefinitionsRequesst, opts ...grpc.CallOption) (*ListAgentJobDefinitionsResponse, error) {
	out := new(ListAgentJobDefinitionsResponse)
	err := c.cc.Invoke(ctx, AgentService_ListAgentJobDefinitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	CreateAgentJobDefinition(context.Context, *CreateAgentJobDefinitionRequest) (*JobDefinition, error)
	DeleteAgentJobDefinition(context.Context, *DeleteAgentJobDefinitionRequest) (*JobDefinition, error)
	ListAgentJobDefinitions(context.Context, *ListAgentJobDefinitionsRequesst) (*ListAgentJobDefinitionsResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) CreateAgentJobDefinition(context.Context, *CreateAgentJobDefinitionRequest) (*JobDefinition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentJobDefinition not implemented")
}
func (UnimplementedAgentServiceServer) DeleteAgentJobDefinition(context.Context, *DeleteAgentJobDefinitionRequest) (*JobDefinition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgentJobDefinition not implemented")
}
func (UnimplementedAgentServiceServer) ListAgentJobDefinitions(context.Context, *ListAgentJobDefinitionsRequesst) (*ListAgentJobDefinitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgentJobDefinitions not implemented")
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

func _AgentService_CreateAgentJobDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentJobDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).CreateAgentJobDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_CreateAgentJobDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).CreateAgentJobDefinition(ctx, req.(*CreateAgentJobDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_DeleteAgentJobDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentJobDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).DeleteAgentJobDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_DeleteAgentJobDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).DeleteAgentJobDefinition(ctx, req.(*DeleteAgentJobDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ListAgentJobDefinitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentJobDefinitionsRequesst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListAgentJobDefinitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_ListAgentJobDefinitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListAgentJobDefinitions(ctx, req.(*ListAgentJobDefinitionsRequesst))
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
			MethodName: "CreateAgentJobDefinition",
			Handler:    _AgentService_CreateAgentJobDefinition_Handler,
		},
		{
			MethodName: "DeleteAgentJobDefinition",
			Handler:    _AgentService_DeleteAgentJobDefinition_Handler,
		},
		{
			MethodName: "ListAgentJobDefinitions",
			Handler:    _AgentService_ListAgentJobDefinitions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "livekit_agent.proto",
}
