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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.4
// source: rpc/sip.proto

package rpc

import (
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InternalUpdateSIPParticipantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId string `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	// IP that SIP INVITE is sent too
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Number used to make the call
	Number string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	// Number to call to
	CallTo              string `protobuf:"bytes,4,opt,name=call_to,json=callTo,proto3" json:"call_to,omitempty"`
	Username            string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password            string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	RoomName            string `protobuf:"bytes,7,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	ParticipantIdentity string `protobuf:"bytes,8,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
}

func (x *InternalUpdateSIPParticipantRequest) Reset() {
	*x = InternalUpdateSIPParticipantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalUpdateSIPParticipantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalUpdateSIPParticipantRequest) ProtoMessage() {}

func (x *InternalUpdateSIPParticipantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalUpdateSIPParticipantRequest.ProtoReflect.Descriptor instead.
func (*InternalUpdateSIPParticipantRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sip_proto_rawDescGZIP(), []int{0}
}

func (x *InternalUpdateSIPParticipantRequest) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *InternalUpdateSIPParticipantRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InternalUpdateSIPParticipantRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *InternalUpdateSIPParticipantRequest) GetCallTo() string {
	if x != nil {
		return x.CallTo
	}
	return ""
}

func (x *InternalUpdateSIPParticipantRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *InternalUpdateSIPParticipantRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *InternalUpdateSIPParticipantRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *InternalUpdateSIPParticipantRequest) GetParticipantIdentity() string {
	if x != nil {
		return x.ParticipantIdentity
	}
	return ""
}

type InternalUpdateSIPParticipantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InternalUpdateSIPParticipantResponse) Reset() {
	*x = InternalUpdateSIPParticipantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalUpdateSIPParticipantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalUpdateSIPParticipantResponse) ProtoMessage() {}

func (x *InternalUpdateSIPParticipantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalUpdateSIPParticipantResponse.ProtoReflect.Descriptor instead.
func (*InternalUpdateSIPParticipantResponse) Descriptor() ([]byte, []int) {
	return file_rpc_sip_proto_rawDescGZIP(), []int{1}
}

type InternalSendSIPParticipantDTMFRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId string `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	Digits        string `protobuf:"bytes,2,opt,name=digits,proto3" json:"digits,omitempty"`
}

func (x *InternalSendSIPParticipantDTMFRequest) Reset() {
	*x = InternalSendSIPParticipantDTMFRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalSendSIPParticipantDTMFRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalSendSIPParticipantDTMFRequest) ProtoMessage() {}

func (x *InternalSendSIPParticipantDTMFRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalSendSIPParticipantDTMFRequest.ProtoReflect.Descriptor instead.
func (*InternalSendSIPParticipantDTMFRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sip_proto_rawDescGZIP(), []int{2}
}

func (x *InternalSendSIPParticipantDTMFRequest) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *InternalSendSIPParticipantDTMFRequest) GetDigits() string {
	if x != nil {
		return x.Digits
	}
	return ""
}

type InternalSendSIPParticipantDTMFResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InternalSendSIPParticipantDTMFResponse) Reset() {
	*x = InternalSendSIPParticipantDTMFResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalSendSIPParticipantDTMFResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalSendSIPParticipantDTMFResponse) ProtoMessage() {}

func (x *InternalSendSIPParticipantDTMFResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalSendSIPParticipantDTMFResponse.ProtoReflect.Descriptor instead.
func (*InternalSendSIPParticipantDTMFResponse) Descriptor() ([]byte, []int) {
	return file_rpc_sip_proto_rawDescGZIP(), []int{3}
}

var File_rpc_sip_proto protoreflect.FileDescriptor

var file_rpc_sip_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x23, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x24, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x0a,
	0x25, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x49, 0x50,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x44, 0x54, 0x4d, 0x46, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x73, 0x22, 0x28, 0x0a, 0x26, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x44, 0x54, 0x4d, 0x46, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xff, 0x01, 0x0a, 0x0b, 0x53, 0x49, 0x50, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12,
	0x75, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xb2, 0x89,
	0x01, 0x04, 0x10, 0x01, 0x30, 0x01, 0x12, 0x79, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x49,
	0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x44, 0x54, 0x4d, 0x46,
	0x12, 0x2a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x44, 0x54, 0x4d, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x44, 0x54, 0x4d,
	0x46, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0xb2, 0x89, 0x01, 0x02, 0x30,
	0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_sip_proto_rawDescOnce sync.Once
	file_rpc_sip_proto_rawDescData = file_rpc_sip_proto_rawDesc
)

func file_rpc_sip_proto_rawDescGZIP() []byte {
	file_rpc_sip_proto_rawDescOnce.Do(func() {
		file_rpc_sip_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_sip_proto_rawDescData)
	})
	return file_rpc_sip_proto_rawDescData
}

var file_rpc_sip_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_sip_proto_goTypes = []interface{}{
	(*InternalUpdateSIPParticipantRequest)(nil),    // 0: rpc.InternalUpdateSIPParticipantRequest
	(*InternalUpdateSIPParticipantResponse)(nil),   // 1: rpc.InternalUpdateSIPParticipantResponse
	(*InternalSendSIPParticipantDTMFRequest)(nil),  // 2: rpc.InternalSendSIPParticipantDTMFRequest
	(*InternalSendSIPParticipantDTMFResponse)(nil), // 3: rpc.InternalSendSIPParticipantDTMFResponse
}
var file_rpc_sip_proto_depIdxs = []int32{
	0, // 0: rpc.SIPInternal.UpdateSIPParticipant:input_type -> rpc.InternalUpdateSIPParticipantRequest
	2, // 1: rpc.SIPInternal.SendSIPParticipantDTMF:input_type -> rpc.InternalSendSIPParticipantDTMFRequest
	1, // 2: rpc.SIPInternal.UpdateSIPParticipant:output_type -> rpc.InternalUpdateSIPParticipantResponse
	3, // 3: rpc.SIPInternal.SendSIPParticipantDTMF:output_type -> rpc.InternalSendSIPParticipantDTMFResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_sip_proto_init() }
func file_rpc_sip_proto_init() {
	if File_rpc_sip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalUpdateSIPParticipantRequest); i {
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
		file_rpc_sip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalUpdateSIPParticipantResponse); i {
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
		file_rpc_sip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalSendSIPParticipantDTMFRequest); i {
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
		file_rpc_sip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalSendSIPParticipantDTMFResponse); i {
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
			RawDescriptor: file_rpc_sip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_sip_proto_goTypes,
		DependencyIndexes: file_rpc_sip_proto_depIdxs,
		MessageInfos:      file_rpc_sip_proto_msgTypes,
	}.Build()
	File_rpc_sip_proto = out.File
	file_rpc_sip_proto_rawDesc = nil
	file_rpc_sip_proto_goTypes = nil
	file_rpc_sip_proto_depIdxs = nil
}
