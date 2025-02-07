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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.23.4
// source: livekit_agent_dispatch.proto

package livekit

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAgentDispatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AgentName     string                 `protobuf:"bytes,1,opt,name=agent_name,json=agentName,proto3" json:"agent_name,omitempty"`
	Room          string                 `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	Metadata      string                 `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAgentDispatchRequest) Reset() {
	*x = CreateAgentDispatchRequest{}
	mi := &file_livekit_agent_dispatch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAgentDispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgentDispatchRequest) ProtoMessage() {}

func (x *CreateAgentDispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_agent_dispatch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgentDispatchRequest.ProtoReflect.Descriptor instead.
func (*CreateAgentDispatchRequest) Descriptor() ([]byte, []int) {
	return file_livekit_agent_dispatch_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAgentDispatchRequest) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

func (x *CreateAgentDispatchRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *CreateAgentDispatchRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type RoomAgentDispatch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AgentName     string                 `protobuf:"bytes,1,opt,name=agent_name,json=agentName,proto3" json:"agent_name,omitempty"`
	Metadata      string                 `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomAgentDispatch) Reset() {
	*x = RoomAgentDispatch{}
	mi := &file_livekit_agent_dispatch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomAgentDispatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomAgentDispatch) ProtoMessage() {}

func (x *RoomAgentDispatch) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_agent_dispatch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomAgentDispatch.ProtoReflect.Descriptor instead.
func (*RoomAgentDispatch) Descriptor() ([]byte, []int) {
	return file_livekit_agent_dispatch_proto_rawDescGZIP(), []int{1}
}

func (x *RoomAgentDispatch) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

func (x *RoomAgentDispatch) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type DeleteAgentDispatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DispatchId    string                 `protobuf:"bytes,1,opt,name=dispatch_id,json=dispatchId,proto3" json:"dispatch_id,omitempty"`
	Room          string                 `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAgentDispatchRequest) Reset() {
	*x = DeleteAgentDispatchRequest{}
	mi := &file_livekit_agent_dispatch_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAgentDispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAgentDispatchRequest) ProtoMessage() {}

func (x *DeleteAgentDispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_agent_dispatch_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAgentDispatchRequest.ProtoReflect.Descriptor instead.
func (*DeleteAgentDispatchRequest) Descriptor() ([]byte, []int) {
	return file_livekit_agent_dispatch_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAgentDispatchRequest) GetDispatchId() string {
	if x != nil {
		return x.DispatchId
	}
	return ""
}

func (x *DeleteAgentDispatchRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

type ListAgentDispatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DispatchId    string                 `protobuf:"bytes,1,opt,name=dispatch_id,json=dispatchId,proto3" json:"dispatch_id,omitempty"` // if set, only the dispatch whose id is given will be returned
	Room          string                 `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`                               // name of the room to list agents for. Must be set.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAgentDispatchRequest) Reset() {
	*x = ListAgentDispatchRequest{}
	mi := &file_livekit_agent_dispatch_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAgentDispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentDispatchRequest) ProtoMessage() {}

func (x *ListAgentDispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_agent_dispatch_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentDispatchRequest.ProtoReflect.Descriptor instead.
func (*ListAgentDispatchRequest) Descriptor() ([]byte, []int) {
	return file_livekit_agent_dispatch_proto_rawDescGZIP(), []int{3}
}

func (x *ListAgentDispatchRequest) GetDispatchId() string {
	if x != nil {
		return x.DispatchId
	}
	return ""
}

func (x *ListAgentDispatchRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

type ListAgentDispatchResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	AgentDispatches []*AgentDispatch       `protobuf:"bytes,1,rep,name=agent_dispatches,json=agentDispatches,proto3" json:"agent_dispatches,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListAgentDispatchResponse) Reset() {
	*x = ListAgentDispatchResponse{}
	mi := &file_livekit_agent_dispatch_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAgentDispatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentDispatchResponse) ProtoMessage() {}

func (x *ListAgentDispatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_agent_dispatch_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentDispatchResponse.ProtoReflect.Descriptor instead.
func (*ListAgentDispatchResponse) Descriptor() ([]byte, []int) {
	return file_livekit_agent_dispatch_proto_rawDescGZIP(), []int{4}
}

func (x *ListAgentDispatchResponse) GetAgentDispatches() []*AgentDispatch {
	if x != nil {
		return x.AgentDispatches
	}
	return nil
}

type AgentDispatch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AgentName     string                 `protobuf:"bytes,2,opt,name=agent_name,json=agentName,proto3" json:"agent_name,omitempty"`
	Room          string                 `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
	Metadata      string                 `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	State         *AgentDispatchState    `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AgentDispatch) Reset() {
	*x = AgentDispatch{}
	mi := &file_livekit_agent_dispatch_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentDispatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentDispatch) ProtoMessage() {}

func (x *AgentDispatch) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_agent_dispatch_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentDispatch.ProtoReflect.Descriptor instead.
func (*AgentDispatch) Descriptor() ([]byte, []int) {
	return file_livekit_agent_dispatch_proto_rawDescGZIP(), []int{5}
}

func (x *AgentDispatch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgentDispatch) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

func (x *AgentDispatch) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *AgentDispatch) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *AgentDispatch) GetState() *AgentDispatchState {
	if x != nil {
		return x.State
	}
	return nil
}

type AgentDispatchState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// For dispatches of tyoe JT_ROOM, there will be at most 1 job.
	// For dispatches of type JT_PUBLISHER, there will be 1 per publisher.
	Jobs          []*Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	CreatedAt     int64  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt     int64  `protobuf:"varint,3,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AgentDispatchState) Reset() {
	*x = AgentDispatchState{}
	mi := &file_livekit_agent_dispatch_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentDispatchState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentDispatchState) ProtoMessage() {}

func (x *AgentDispatchState) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_agent_dispatch_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentDispatchState.ProtoReflect.Descriptor instead.
func (*AgentDispatchState) Descriptor() ([]byte, []int) {
	return file_livekit_agent_dispatch_proto_rawDescGZIP(), []int{6}
}

func (x *AgentDispatchState) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *AgentDispatchState) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AgentDispatchState) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_livekit_agent_dispatch_proto protoreflect.FileDescriptor

var file_livekit_agent_dispatch_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x1a, 0x13, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x1a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x11, 0x52, 0x6f, 0x6f,
	0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x1a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x4f, 0x0a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x5e, 0x0a,
	0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0xa1, 0x01,
	0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x74, 0x0a, 0x12, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x8b, 0x02, 0x0a, 0x14, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x23, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x4d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x23, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x55,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x21,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0xaa, 0x02, 0x0d, 0x4c,
	0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xea, 0x02, 0x0e, 0x4c,
	0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_livekit_agent_dispatch_proto_rawDescOnce sync.Once
	file_livekit_agent_dispatch_proto_rawDescData []byte
)

func file_livekit_agent_dispatch_proto_rawDescGZIP() []byte {
	file_livekit_agent_dispatch_proto_rawDescOnce.Do(func() {
		file_livekit_agent_dispatch_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_livekit_agent_dispatch_proto_rawDesc), len(file_livekit_agent_dispatch_proto_rawDesc)))
	})
	return file_livekit_agent_dispatch_proto_rawDescData
}

var file_livekit_agent_dispatch_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_livekit_agent_dispatch_proto_goTypes = []any{
	(*CreateAgentDispatchRequest)(nil), // 0: livekit.CreateAgentDispatchRequest
	(*RoomAgentDispatch)(nil),          // 1: livekit.RoomAgentDispatch
	(*DeleteAgentDispatchRequest)(nil), // 2: livekit.DeleteAgentDispatchRequest
	(*ListAgentDispatchRequest)(nil),   // 3: livekit.ListAgentDispatchRequest
	(*ListAgentDispatchResponse)(nil),  // 4: livekit.ListAgentDispatchResponse
	(*AgentDispatch)(nil),              // 5: livekit.AgentDispatch
	(*AgentDispatchState)(nil),         // 6: livekit.AgentDispatchState
	(*Job)(nil),                        // 7: livekit.Job
}
var file_livekit_agent_dispatch_proto_depIdxs = []int32{
	5, // 0: livekit.ListAgentDispatchResponse.agent_dispatches:type_name -> livekit.AgentDispatch
	6, // 1: livekit.AgentDispatch.state:type_name -> livekit.AgentDispatchState
	7, // 2: livekit.AgentDispatchState.jobs:type_name -> livekit.Job
	0, // 3: livekit.AgentDispatchService.CreateDispatch:input_type -> livekit.CreateAgentDispatchRequest
	2, // 4: livekit.AgentDispatchService.DeleteDispatch:input_type -> livekit.DeleteAgentDispatchRequest
	3, // 5: livekit.AgentDispatchService.ListDispatch:input_type -> livekit.ListAgentDispatchRequest
	5, // 6: livekit.AgentDispatchService.CreateDispatch:output_type -> livekit.AgentDispatch
	5, // 7: livekit.AgentDispatchService.DeleteDispatch:output_type -> livekit.AgentDispatch
	4, // 8: livekit.AgentDispatchService.ListDispatch:output_type -> livekit.ListAgentDispatchResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_livekit_agent_dispatch_proto_init() }
func file_livekit_agent_dispatch_proto_init() {
	if File_livekit_agent_dispatch_proto != nil {
		return
	}
	file_livekit_agent_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_livekit_agent_dispatch_proto_rawDesc), len(file_livekit_agent_dispatch_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_livekit_agent_dispatch_proto_goTypes,
		DependencyIndexes: file_livekit_agent_dispatch_proto_depIdxs,
		MessageInfos:      file_livekit_agent_dispatch_proto_msgTypes,
	}.Build()
	File_livekit_agent_dispatch_proto = out.File
	file_livekit_agent_dispatch_proto_goTypes = nil
	file_livekit_agent_dispatch_proto_depIdxs = nil
}
