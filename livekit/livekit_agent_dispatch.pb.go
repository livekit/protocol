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
// 	protoc-gen-go v1.36.6
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

const file_livekit_agent_dispatch_proto_rawDesc = "" +
	"\n" +
	"\x1clivekit_agent_dispatch.proto\x12\alivekit\x1a\x13livekit_agent.proto\"k\n" +
	"\x1aCreateAgentDispatchRequest\x12\x1d\n" +
	"\n" +
	"agent_name\x18\x01 \x01(\tR\tagentName\x12\x12\n" +
	"\x04room\x18\x02 \x01(\tR\x04room\x12\x1a\n" +
	"\bmetadata\x18\x03 \x01(\tR\bmetadata\"N\n" +
	"\x11RoomAgentDispatch\x12\x1d\n" +
	"\n" +
	"agent_name\x18\x01 \x01(\tR\tagentName\x12\x1a\n" +
	"\bmetadata\x18\x02 \x01(\tR\bmetadata\"Q\n" +
	"\x1aDeleteAgentDispatchRequest\x12\x1f\n" +
	"\vdispatch_id\x18\x01 \x01(\tR\n" +
	"dispatchId\x12\x12\n" +
	"\x04room\x18\x02 \x01(\tR\x04room\"O\n" +
	"\x18ListAgentDispatchRequest\x12\x1f\n" +
	"\vdispatch_id\x18\x01 \x01(\tR\n" +
	"dispatchId\x12\x12\n" +
	"\x04room\x18\x02 \x01(\tR\x04room\"^\n" +
	"\x19ListAgentDispatchResponse\x12A\n" +
	"\x10agent_dispatches\x18\x01 \x03(\v2\x16.livekit.AgentDispatchR\x0fagentDispatches\"\xa1\x01\n" +
	"\rAgentDispatch\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"agent_name\x18\x02 \x01(\tR\tagentName\x12\x12\n" +
	"\x04room\x18\x03 \x01(\tR\x04room\x12\x1a\n" +
	"\bmetadata\x18\x04 \x01(\tR\bmetadata\x121\n" +
	"\x05state\x18\x05 \x01(\v2\x1b.livekit.AgentDispatchStateR\x05state\"t\n" +
	"\x12AgentDispatchState\x12 \n" +
	"\x04jobs\x18\x01 \x03(\v2\f.livekit.JobR\x04jobs\x12\x1d\n" +
	"\n" +
	"created_at\x18\x02 \x01(\x03R\tcreatedAt\x12\x1d\n" +
	"\n" +
	"deleted_at\x18\x03 \x01(\x03R\tdeletedAt2\x8b\x02\n" +
	"\x14AgentDispatchService\x12M\n" +
	"\x0eCreateDispatch\x12#.livekit.CreateAgentDispatchRequest\x1a\x16.livekit.AgentDispatch\x12M\n" +
	"\x0eDeleteDispatch\x12#.livekit.DeleteAgentDispatchRequest\x1a\x16.livekit.AgentDispatch\x12U\n" +
	"\fListDispatch\x12!.livekit.ListAgentDispatchRequest\x1a\".livekit.ListAgentDispatchResponseBFZ#github.com/livekit/protocol/livekit\xaa\x02\rLiveKit.Proto\xea\x02\x0eLiveKit::Protob\x06proto3"

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
