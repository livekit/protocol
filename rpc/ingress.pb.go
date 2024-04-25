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
// 	protoc        v4.23.4
// source: rpc/ingress.proto

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

type ListActiveIngressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListActiveIngressRequest) Reset() {
	*x = ListActiveIngressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ingress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActiveIngressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveIngressRequest) ProtoMessage() {}

func (x *ListActiveIngressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActiveIngressRequest.ProtoReflect.Descriptor instead.
func (*ListActiveIngressRequest) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_proto_rawDescGZIP(), []int{0}
}

type ListActiveIngressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngressIds []string `protobuf:"bytes,1,rep,name=ingress_ids,json=ingressIds,proto3" json:"ingress_ids,omitempty"`
}

func (x *ListActiveIngressResponse) Reset() {
	*x = ListActiveIngressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ingress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActiveIngressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveIngressResponse) ProtoMessage() {}

func (x *ListActiveIngressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActiveIngressResponse.ProtoReflect.Descriptor instead.
func (*ListActiveIngressResponse) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_proto_rawDescGZIP(), []int{1}
}

func (x *ListActiveIngressResponse) GetIngressIds() []string {
	if x != nil {
		return x.IngressIds
	}
	return nil
}

type DeleteWHIPResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	StreamKey  string `protobuf:"bytes,2,opt,name=stream_key,json=streamKey,proto3" json:"stream_key,omitempty"`
}

func (x *DeleteWHIPResourceRequest) Reset() {
	*x = DeleteWHIPResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ingress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWHIPResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWHIPResourceRequest) ProtoMessage() {}

func (x *DeleteWHIPResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWHIPResourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteWHIPResourceRequest) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteWHIPResourceRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *DeleteWHIPResourceRequest) GetStreamKey() string {
	if x != nil {
		return x.StreamKey
	}
	return ""
}

type StartIngressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info          *livekit.IngressInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Token         string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	WsUrl         string               `protobuf:"bytes,3,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
	LoggingFields map[string]string    `protobuf:"bytes,4,rep,name=logging_fields,json=loggingFields,proto3" json:"logging_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StartIngressRequest) Reset() {
	*x = StartIngressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ingress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartIngressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartIngressRequest) ProtoMessage() {}

func (x *StartIngressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartIngressRequest.ProtoReflect.Descriptor instead.
func (*StartIngressRequest) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_proto_rawDescGZIP(), []int{3}
}

func (x *StartIngressRequest) GetInfo() *livekit.IngressInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *StartIngressRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StartIngressRequest) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

func (x *StartIngressRequest) GetLoggingFields() map[string]string {
	if x != nil {
		return x.LoggingFields
	}
	return nil
}

var File_rpc_ingress_proto protoreflect.FileDescriptor

var file_rpc_ingress_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x73, 0x22, 0x5b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57,
	0x48, 0x49, 0x50, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4b,
	0x65, 0x79, 0x22, 0x82, 0x02, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x77, 0x73,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x73, 0x55, 0x72,
	0x6c, 0x12, 0x52, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xb7, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x46, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x06, 0xb2, 0x89, 0x01,
	0x02, 0x30, 0x01, 0x12, 0x5c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xb2, 0x89, 0x01, 0x04, 0x10, 0x01, 0x28,
	0x01, 0x32, 0x84, 0x02, 0x0a, 0x0e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x06, 0xb2, 0x89, 0x01,
	0x02, 0x10, 0x01, 0x12, 0x4d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x06, 0xb2, 0x89, 0x01, 0x02,
	0x10, 0x01, 0x12, 0x54, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x48, 0x49, 0x50,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x48, 0x49, 0x50, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x06, 0xb2, 0x89, 0x01, 0x02, 0x10, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_ingress_proto_rawDescOnce sync.Once
	file_rpc_ingress_proto_rawDescData = file_rpc_ingress_proto_rawDesc
)

func file_rpc_ingress_proto_rawDescGZIP() []byte {
	file_rpc_ingress_proto_rawDescOnce.Do(func() {
		file_rpc_ingress_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_ingress_proto_rawDescData)
	})
	return file_rpc_ingress_proto_rawDescData
}

var file_rpc_ingress_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_ingress_proto_goTypes = []interface{}{
	(*ListActiveIngressRequest)(nil),     // 0: rpc.ListActiveIngressRequest
	(*ListActiveIngressResponse)(nil),    // 1: rpc.ListActiveIngressResponse
	(*DeleteWHIPResourceRequest)(nil),    // 2: rpc.DeleteWHIPResourceRequest
	(*StartIngressRequest)(nil),          // 3: rpc.StartIngressRequest
	nil,                                  // 4: rpc.StartIngressRequest.LoggingFieldsEntry
	(*livekit.IngressInfo)(nil),          // 5: livekit.IngressInfo
	(*livekit.UpdateIngressRequest)(nil), // 6: livekit.UpdateIngressRequest
	(*livekit.DeleteIngressRequest)(nil), // 7: livekit.DeleteIngressRequest
	(*livekit.IngressState)(nil),         // 8: livekit.IngressState
	(*emptypb.Empty)(nil),                // 9: google.protobuf.Empty
}
var file_rpc_ingress_proto_depIdxs = []int32{
	5, // 0: rpc.StartIngressRequest.info:type_name -> livekit.IngressInfo
	4, // 1: rpc.StartIngressRequest.logging_fields:type_name -> rpc.StartIngressRequest.LoggingFieldsEntry
	3, // 2: rpc.IngressInternal.StartIngress:input_type -> rpc.StartIngressRequest
	0, // 3: rpc.IngressInternal.ListActiveIngress:input_type -> rpc.ListActiveIngressRequest
	6, // 4: rpc.IngressHandler.UpdateIngress:input_type -> livekit.UpdateIngressRequest
	7, // 5: rpc.IngressHandler.DeleteIngress:input_type -> livekit.DeleteIngressRequest
	2, // 6: rpc.IngressHandler.DeleteWHIPResource:input_type -> rpc.DeleteWHIPResourceRequest
	5, // 7: rpc.IngressInternal.StartIngress:output_type -> livekit.IngressInfo
	1, // 8: rpc.IngressInternal.ListActiveIngress:output_type -> rpc.ListActiveIngressResponse
	8, // 9: rpc.IngressHandler.UpdateIngress:output_type -> livekit.IngressState
	8, // 10: rpc.IngressHandler.DeleteIngress:output_type -> livekit.IngressState
	9, // 11: rpc.IngressHandler.DeleteWHIPResource:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_ingress_proto_init() }
func file_rpc_ingress_proto_init() {
	if File_rpc_ingress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_ingress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActiveIngressRequest); i {
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
		file_rpc_ingress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActiveIngressResponse); i {
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
		file_rpc_ingress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWHIPResourceRequest); i {
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
		file_rpc_ingress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartIngressRequest); i {
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
			RawDescriptor: file_rpc_ingress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_rpc_ingress_proto_goTypes,
		DependencyIndexes: file_rpc_ingress_proto_depIdxs,
		MessageInfos:      file_rpc_ingress_proto_msgTypes,
	}.Build()
	File_rpc_ingress_proto = out.File
	file_rpc_ingress_proto_rawDesc = nil
	file_rpc_ingress_proto_goTypes = nil
	file_rpc_ingress_proto_depIdxs = nil
}
