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
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: infra/link.proto

package infra

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WatchLocalLinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchLocalLinksRequest) Reset() {
	*x = WatchLocalLinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchLocalLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchLocalLinksRequest) ProtoMessage() {}

func (x *WatchLocalLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchLocalLinksRequest.ProtoReflect.Descriptor instead.
func (*WatchLocalLinksRequest) Descriptor() ([]byte, []int) {
	return file_infra_link_proto_rawDescGZIP(), []int{0}
}

type WatchLocalLinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalRegion  string                 `protobuf:"bytes,1,opt,name=local_region,json=localRegion,proto3" json:"local_region,omitempty"`
	RemoteRegion string                 `protobuf:"bytes,2,opt,name=remote_region,json=remoteRegion,proto3" json:"remote_region,omitempty"`
	Rtt          int64                  `protobuf:"varint,3,opt,name=rtt,proto3" json:"rtt,omitempty"`
	Jitter       int64                  `protobuf:"varint,4,opt,name=jitter,proto3" json:"jitter,omitempty"`
	PacketLoss   float64                `protobuf:"fixed64,5,opt,name=packet_loss,json=packetLoss,proto3" json:"packet_loss,omitempty"`
	Disabled     bool                   `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	LastRead     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_read,json=lastRead,proto3" json:"last_read,omitempty"`
}

func (x *WatchLocalLinksResponse) Reset() {
	*x = WatchLocalLinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchLocalLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchLocalLinksResponse) ProtoMessage() {}

func (x *WatchLocalLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchLocalLinksResponse.ProtoReflect.Descriptor instead.
func (*WatchLocalLinksResponse) Descriptor() ([]byte, []int) {
	return file_infra_link_proto_rawDescGZIP(), []int{1}
}

func (x *WatchLocalLinksResponse) GetLocalRegion() string {
	if x != nil {
		return x.LocalRegion
	}
	return ""
}

func (x *WatchLocalLinksResponse) GetRemoteRegion() string {
	if x != nil {
		return x.RemoteRegion
	}
	return ""
}

func (x *WatchLocalLinksResponse) GetRtt() int64 {
	if x != nil {
		return x.Rtt
	}
	return 0
}

func (x *WatchLocalLinksResponse) GetJitter() int64 {
	if x != nil {
		return x.Jitter
	}
	return 0
}

func (x *WatchLocalLinksResponse) GetPacketLoss() float64 {
	if x != nil {
		return x.PacketLoss
	}
	return 0
}

func (x *WatchLocalLinksResponse) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *WatchLocalLinksResponse) GetLastRead() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRead
	}
	return nil
}

type SimulateLinkStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalRegion  string   `protobuf:"bytes,1,opt,name=local_region,json=localRegion,proto3" json:"local_region,omitempty"`
	RemoteRegion string   `protobuf:"bytes,2,opt,name=remote_region,json=remoteRegion,proto3" json:"remote_region,omitempty"`
	Rtt          *int64   `protobuf:"varint,3,opt,name=rtt,proto3,oneof" json:"rtt,omitempty"`
	Jitter       *int64   `protobuf:"varint,4,opt,name=jitter,proto3,oneof" json:"jitter,omitempty"`
	PacketLoss   *float64 `protobuf:"fixed64,5,opt,name=packet_loss,json=packetLoss,proto3,oneof" json:"packet_loss,omitempty"`
	Disabled     *bool    `protobuf:"varint,6,opt,name=disabled,proto3,oneof" json:"disabled,omitempty"`
	Timeout      int64    `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *SimulateLinkStateRequest) Reset() {
	*x = SimulateLinkStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_link_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateLinkStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateLinkStateRequest) ProtoMessage() {}

func (x *SimulateLinkStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_link_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateLinkStateRequest.ProtoReflect.Descriptor instead.
func (*SimulateLinkStateRequest) Descriptor() ([]byte, []int) {
	return file_infra_link_proto_rawDescGZIP(), []int{2}
}

func (x *SimulateLinkStateRequest) GetLocalRegion() string {
	if x != nil {
		return x.LocalRegion
	}
	return ""
}

func (x *SimulateLinkStateRequest) GetRemoteRegion() string {
	if x != nil {
		return x.RemoteRegion
	}
	return ""
}

func (x *SimulateLinkStateRequest) GetRtt() int64 {
	if x != nil && x.Rtt != nil {
		return *x.Rtt
	}
	return 0
}

func (x *SimulateLinkStateRequest) GetJitter() int64 {
	if x != nil && x.Jitter != nil {
		return *x.Jitter
	}
	return 0
}

func (x *SimulateLinkStateRequest) GetPacketLoss() float64 {
	if x != nil && x.PacketLoss != nil {
		return *x.PacketLoss
	}
	return 0
}

func (x *SimulateLinkStateRequest) GetDisabled() bool {
	if x != nil && x.Disabled != nil {
		return *x.Disabled
	}
	return false
}

func (x *SimulateLinkStateRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type SimulateLinkStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SimulateLinkStateResponse) Reset() {
	*x = SimulateLinkStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_link_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateLinkStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateLinkStateResponse) ProtoMessage() {}

func (x *SimulateLinkStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_link_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateLinkStateResponse.ProtoReflect.Descriptor instead.
func (*SimulateLinkStateResponse) Descriptor() ([]byte, []int) {
	return file_infra_link_proto_rawDescGZIP(), []int{3}
}

var File_infra_link_proto protoreflect.FileDescriptor

var file_infra_link_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x81, 0x02, 0x0a, 0x17, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x74, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x74, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x37, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x22, 0xa7, 0x02, 0x0a, 0x18, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x03, 0x72,
	0x74, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x03, 0x72, 0x74, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x01, 0x52, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f,
	0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x74, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6a, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c,
	0x6f, 0x73, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x22, 0x1b, 0x0a, 0x19, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xaa, 0x01,
	0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x4e, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_link_proto_rawDescOnce sync.Once
	file_infra_link_proto_rawDescData = file_infra_link_proto_rawDesc
)

func file_infra_link_proto_rawDescGZIP() []byte {
	file_infra_link_proto_rawDescOnce.Do(func() {
		file_infra_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_link_proto_rawDescData)
	})
	return file_infra_link_proto_rawDescData
}

var file_infra_link_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infra_link_proto_goTypes = []interface{}{
	(*WatchLocalLinksRequest)(nil),    // 0: rpc.WatchLocalLinksRequest
	(*WatchLocalLinksResponse)(nil),   // 1: rpc.WatchLocalLinksResponse
	(*SimulateLinkStateRequest)(nil),  // 2: rpc.SimulateLinkStateRequest
	(*SimulateLinkStateResponse)(nil), // 3: rpc.SimulateLinkStateResponse
	(*timestamppb.Timestamp)(nil),     // 4: google.protobuf.Timestamp
}
var file_infra_link_proto_depIdxs = []int32{
	4, // 0: rpc.WatchLocalLinksResponse.last_read:type_name -> google.protobuf.Timestamp
	0, // 1: rpc.Link.WatchLocalLinks:input_type -> rpc.WatchLocalLinksRequest
	2, // 2: rpc.Link.SimulateLinkState:input_type -> rpc.SimulateLinkStateRequest
	1, // 3: rpc.Link.WatchLocalLinks:output_type -> rpc.WatchLocalLinksResponse
	3, // 4: rpc.Link.SimulateLinkState:output_type -> rpc.SimulateLinkStateResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_link_proto_init() }
func file_infra_link_proto_init() {
	if File_infra_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchLocalLinksRequest); i {
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
		file_infra_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchLocalLinksResponse); i {
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
		file_infra_link_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateLinkStateRequest); i {
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
		file_infra_link_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateLinkStateResponse); i {
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
	file_infra_link_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_link_proto_goTypes,
		DependencyIndexes: file_infra_link_proto_depIdxs,
		MessageInfos:      file_infra_link_proto_msgTypes,
	}.Build()
	File_infra_link_proto = out.File
	file_infra_link_proto_rawDesc = nil
	file_infra_link_proto_goTypes = nil
	file_infra_link_proto_depIdxs = nil
}
