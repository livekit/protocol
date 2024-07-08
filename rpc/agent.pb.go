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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckEnabledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckEnabledRequest) Reset() {
	*x = CheckEnabledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEnabledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnabledRequest) ProtoMessage() {}

func (x *CheckEnabledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomEnabled      bool `protobuf:"varint,1,opt,name=room_enabled,json=roomEnabled,proto3" json:"room_enabled,omitempty"`
	PublisherEnabled bool `protobuf:"varint,2,opt,name=publisher_enabled,json=publisherEnabled,proto3" json:"publisher_enabled,omitempty"`
	// Deprecated: Marked as deprecated in rpc/agent.proto.
	Namespaces []string `protobuf:"bytes,3,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	AgentNames []string `protobuf:"bytes,4,rep,name=agent_names,json=agentNames,proto3" json:"agent_names,omitempty"`
}

func (x *CheckEnabledResponse) Reset() {
	*x = CheckEnabledResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEnabledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnabledResponse) ProtoMessage() {}

func (x *CheckEnabledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

var File_rpc_agent_proto protoreflect.FileDescriptor

var file_rpc_agent_proto_rawDesc = []byte{
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
	0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x96, 0x02, 0x0a,
	0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x4b,
	0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x06, 0xb2, 0x89, 0x01, 0x02, 0x28, 0x01, 0x12, 0x53, 0x0a, 0x0a, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1f, 0xb2, 0x89, 0x01, 0x1b, 0x10, 0x01, 0x1a, 0x15, 0x12, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x30, 0x01,
	0x12, 0x63, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0xb2, 0x89, 0x01, 0x1b, 0x08, 0x01, 0x10, 0x01, 0x1a, 0x13,
	0x12, 0x11, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x28, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_agent_proto_rawDescOnce sync.Once
	file_rpc_agent_proto_rawDescData = file_rpc_agent_proto_rawDesc
)

func file_rpc_agent_proto_rawDescGZIP() []byte {
	file_rpc_agent_proto_rawDescOnce.Do(func() {
		file_rpc_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_agent_proto_rawDescData)
	})
	return file_rpc_agent_proto_rawDescData
}

var file_rpc_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_agent_proto_goTypes = []interface{}{
	(*CheckEnabledRequest)(nil),  // 0: rpc.CheckEnabledRequest
	(*CheckEnabledResponse)(nil), // 1: rpc.CheckEnabledResponse
	(*livekit.Job)(nil),          // 2: livekit.Job
	(*emptypb.Empty)(nil),        // 3: google.protobuf.Empty
}
var file_rpc_agent_proto_depIdxs = []int32{
	0, // 0: rpc.AgentInternal.CheckEnabled:input_type -> rpc.CheckEnabledRequest
	2, // 1: rpc.AgentInternal.JobRequest:input_type -> livekit.Job
	3, // 2: rpc.AgentInternal.WorkerRegistered:input_type -> google.protobuf.Empty
	1, // 3: rpc.AgentInternal.CheckEnabled:output_type -> rpc.CheckEnabledResponse
	3, // 4: rpc.AgentInternal.JobRequest:output_type -> google.protobuf.Empty
	3, // 5: rpc.AgentInternal.WorkerRegistered:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_agent_proto_init() }
func file_rpc_agent_proto_init() {
	if File_rpc_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEnabledRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEnabledResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_agent_proto_goTypes,
		DependencyIndexes: file_rpc_agent_proto_depIdxs,
		MessageInfos:      file_rpc_agent_proto_msgTypes,
	}.Build()
	File_rpc_agent_proto = out.File
	file_rpc_agent_proto_rawDesc = nil
	file_rpc_agent_proto_goTypes = nil
	file_rpc_agent_proto_depIdxs = nil
}
