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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.23.4
// source: rpc/agent.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type JobTerminateReason int32

const (
	JobTerminateReason_TERMINATION_REQUESTED JobTerminateReason = 0
	JobTerminateReason_AGENT_LEFT_ROOM       JobTerminateReason = 1
)

// Enum value maps for JobTerminateReason.
var (
	JobTerminateReason_name = map[int32]string{
		0: "TERMINATION_REQUESTED",
		1: "AGENT_LEFT_ROOM",
	}
	JobTerminateReason_value = map[string]int32{
		"TERMINATION_REQUESTED": 0,
		"AGENT_LEFT_ROOM":       1,
	}
)

func (x JobTerminateReason) Enum() *JobTerminateReason {
	p := new(JobTerminateReason)
	*p = x
	return p
}

func (x JobTerminateReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobTerminateReason) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_agent_proto_enumTypes[0].Descriptor()
}

func (JobTerminateReason) Type() protoreflect.EnumType {
	return &file_rpc_agent_proto_enumTypes[0]
}

func (x JobTerminateReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobTerminateReason.Descriptor instead.
func (JobTerminateReason) EnumDescriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{0}
}

type CheckEnabledRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckEnabledRequest) Reset() {
	*x = CheckEnabledRequest{}
	mi := &file_rpc_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckEnabledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnabledRequest) ProtoMessage() {}

func (x *CheckEnabledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEnabledRequest.ProtoReflect.Descriptor instead.
func (*CheckEnabledRequest) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{0}
}

type CheckEnabledResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RoomEnabled      bool                   `protobuf:"varint,1,opt,name=room_enabled,json=roomEnabled,proto3" json:"room_enabled,omitempty"`
	PublisherEnabled bool                   `protobuf:"varint,2,opt,name=publisher_enabled,json=publisherEnabled,proto3" json:"publisher_enabled,omitempty"`
	// Deprecated: Marked as deprecated in rpc/agent.proto.
	Namespaces    []string `protobuf:"bytes,3,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	AgentNames    []string `protobuf:"bytes,4,rep,name=agent_names,json=agentNames,proto3" json:"agent_names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckEnabledResponse) Reset() {
	*x = CheckEnabledResponse{}
	mi := &file_rpc_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckEnabledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnabledResponse) ProtoMessage() {}

func (x *CheckEnabledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEnabledResponse.ProtoReflect.Descriptor instead.
func (*CheckEnabledResponse) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{1}
}

func (x *CheckEnabledResponse) GetRoomEnabled() bool {
	if x != nil {
		return x.RoomEnabled
	}
	return false
}

func (x *CheckEnabledResponse) GetPublisherEnabled() bool {
	if x != nil {
		return x.PublisherEnabled
	}
	return false
}

// Deprecated: Marked as deprecated in rpc/agent.proto.
func (x *CheckEnabledResponse) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *CheckEnabledResponse) GetAgentNames() []string {
	if x != nil {
		return x.AgentNames
	}
	return nil
}

type JobRequestResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         *livekit.JobState      `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobRequestResponse) Reset() {
	*x = JobRequestResponse{}
	mi := &file_rpc_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequestResponse) ProtoMessage() {}

func (x *JobRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequestResponse.ProtoReflect.Descriptor instead.
func (*JobRequestResponse) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{2}
}

func (x *JobRequestResponse) GetState() *livekit.JobState {
	if x != nil {
		return x.State
	}
	return nil
}

type JobTerminateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Reason        JobTerminateReason     `protobuf:"varint,2,opt,name=reason,proto3,enum=rpc.JobTerminateReason" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTerminateRequest) Reset() {
	*x = JobTerminateRequest{}
	mi := &file_rpc_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTerminateRequest) ProtoMessage() {}

func (x *JobTerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTerminateRequest.ProtoReflect.Descriptor instead.
func (*JobTerminateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{3}
}

func (x *JobTerminateRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobTerminateRequest) GetReason() JobTerminateReason {
	if x != nil {
		return x.Reason
	}
	return JobTerminateReason_TERMINATION_REQUESTED
}

type JobTerminateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         *livekit.JobState      `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTerminateResponse) Reset() {
	*x = JobTerminateResponse{}
	mi := &file_rpc_agent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTerminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTerminateResponse) ProtoMessage() {}

func (x *JobTerminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTerminateResponse.ProtoReflect.Descriptor instead.
func (*JobTerminateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{4}
}

func (x *JobTerminateResponse) GetState() *livekit.JobState {
	if x != nil {
		return x.State
	}
	return nil
}

var File_rpc_agent_proto protoreflect.FileDescriptor

var file_rpc_agent_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xab,
	0x01, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72,
	0x6f, 0x6f, 0x6d, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x12,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x4a,
	0x6f, 0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x4a, 0x6f, 0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x14, 0x4a, 0x6f,
	0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x44, 0x0a, 0x12, 0x4a,
	0x6f, 0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10,
	0x01, 0x32, 0xee, 0x02, 0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0xb2, 0x89, 0x01, 0x02, 0x28, 0x01,
	0x12, 0x54, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x1a, 0x17, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0xb2, 0x89, 0x01, 0x1b, 0x10, 0x01, 0x1a, 0x15, 0x12,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x08, 0x6a, 0x6f, 0x62, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x62,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xb2, 0x89, 0x01,
	0x0c, 0x10, 0x01, 0x1a, 0x08, 0x12, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x12, 0x63, 0x0a,
	0x10, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1f, 0xb2, 0x89, 0x01, 0x1b, 0x08, 0x01, 0x10, 0x01, 0x1a, 0x13, 0x12, 0x11, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x28, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_rpc_agent_proto_rawDescOnce sync.Once
	file_rpc_agent_proto_rawDescData []byte
)

func file_rpc_agent_proto_rawDescGZIP() []byte {
	file_rpc_agent_proto_rawDescOnce.Do(func() {
		file_rpc_agent_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_agent_proto_rawDesc), len(file_rpc_agent_proto_rawDesc)))
	})
	return file_rpc_agent_proto_rawDescData
}

var file_rpc_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_agent_proto_goTypes = []any{
	(JobTerminateReason)(0),      // 0: rpc.JobTerminateReason
	(*CheckEnabledRequest)(nil),  // 1: rpc.CheckEnabledRequest
	(*CheckEnabledResponse)(nil), // 2: rpc.CheckEnabledResponse
	(*JobRequestResponse)(nil),   // 3: rpc.JobRequestResponse
	(*JobTerminateRequest)(nil),  // 4: rpc.JobTerminateRequest
	(*JobTerminateResponse)(nil), // 5: rpc.JobTerminateResponse
	(*livekit.JobState)(nil),     // 6: livekit.JobState
	(*livekit.Job)(nil),          // 7: livekit.Job
	(*emptypb.Empty)(nil),        // 8: google.protobuf.Empty
}
var file_rpc_agent_proto_depIdxs = []int32{
	6, // 0: rpc.JobRequestResponse.state:type_name -> livekit.JobState
	0, // 1: rpc.JobTerminateRequest.reason:type_name -> rpc.JobTerminateReason
	6, // 2: rpc.JobTerminateResponse.state:type_name -> livekit.JobState
	1, // 3: rpc.AgentInternal.CheckEnabled:input_type -> rpc.CheckEnabledRequest
	7, // 4: rpc.AgentInternal.JobRequest:input_type -> livekit.Job
	4, // 5: rpc.AgentInternal.JobTerminate:input_type -> rpc.JobTerminateRequest
	8, // 6: rpc.AgentInternal.WorkerRegistered:input_type -> google.protobuf.Empty
	2, // 7: rpc.AgentInternal.CheckEnabled:output_type -> rpc.CheckEnabledResponse
	3, // 8: rpc.AgentInternal.JobRequest:output_type -> rpc.JobRequestResponse
	5, // 9: rpc.AgentInternal.JobTerminate:output_type -> rpc.JobTerminateResponse
	8, // 10: rpc.AgentInternal.WorkerRegistered:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_agent_proto_init() }
func file_rpc_agent_proto_init() {
	if File_rpc_agent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_agent_proto_rawDesc), len(file_rpc_agent_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_agent_proto_goTypes,
		DependencyIndexes: file_rpc_agent_proto_depIdxs,
		EnumInfos:         file_rpc_agent_proto_enumTypes,
		MessageInfos:      file_rpc_agent_proto_msgTypes,
	}.Build()
	File_rpc_agent_proto = out.File
	file_rpc_agent_proto_goTypes = nil
	file_rpc_agent_proto_depIdxs = nil
}
