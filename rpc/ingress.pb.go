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
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListActiveIngressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListActiveIngressRequest) Reset() {
	*x = ListActiveIngressRequest{}
	mi := &file_rpc_ingress_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListActiveIngressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveIngressRequest) ProtoMessage() {}

func (x *ListActiveIngressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[0]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in rpc/ingress.proto.
	IngressIds      []string          `protobuf:"bytes,1,rep,name=ingress_ids,json=ingressIds,proto3" json:"ingress_ids,omitempty"`
	IngressSessions []*IngressSession `protobuf:"bytes,2,rep,name=ingress_sessions,json=ingressSessions,proto3" json:"ingress_sessions,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListActiveIngressResponse) Reset() {
	*x = ListActiveIngressResponse{}
	mi := &file_rpc_ingress_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListActiveIngressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveIngressResponse) ProtoMessage() {}

func (x *ListActiveIngressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[1]
	if x != nil {
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

// Deprecated: Marked as deprecated in rpc/ingress.proto.
func (x *ListActiveIngressResponse) GetIngressIds() []string {
	if x != nil {
		return x.IngressIds
	}
	return nil
}

func (x *ListActiveIngressResponse) GetIngressSessions() []*IngressSession {
	if x != nil {
		return x.IngressSessions
	}
	return nil
}

type DeleteWHIPResourceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceId    string                 `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	StreamKey     string                 `protobuf:"bytes,2,opt,name=stream_key,json=streamKey,proto3" json:"stream_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWHIPResourceRequest) Reset() {
	*x = DeleteWHIPResourceRequest{}
	mi := &file_rpc_ingress_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWHIPResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWHIPResourceRequest) ProtoMessage() {}

func (x *DeleteWHIPResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[2]
	if x != nil {
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

type ICERestartWHIPResourceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceId    string                 `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	StreamKey     string                 `protobuf:"bytes,2,opt,name=stream_key,json=streamKey,proto3" json:"stream_key,omitempty"`
	UserFragment  string                 `protobuf:"bytes,3,opt,name=user_fragment,json=userFragment,proto3" json:"user_fragment,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Candidates    []string               `protobuf:"bytes,5,rep,name=candidates,proto3" json:"candidates,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ICERestartWHIPResourceRequest) Reset() {
	*x = ICERestartWHIPResourceRequest{}
	mi := &file_rpc_ingress_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ICERestartWHIPResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ICERestartWHIPResourceRequest) ProtoMessage() {}

func (x *ICERestartWHIPResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ICERestartWHIPResourceRequest.ProtoReflect.Descriptor instead.
func (*ICERestartWHIPResourceRequest) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_proto_rawDescGZIP(), []int{3}
}

func (x *ICERestartWHIPResourceRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ICERestartWHIPResourceRequest) GetStreamKey() string {
	if x != nil {
		return x.StreamKey
	}
	return ""
}

func (x *ICERestartWHIPResourceRequest) GetUserFragment() string {
	if x != nil {
		return x.UserFragment
	}
	return ""
}

func (x *ICERestartWHIPResourceRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ICERestartWHIPResourceRequest) GetCandidates() []string {
	if x != nil {
		return x.Candidates
	}
	return nil
}

type ICERestartWHIPResourceResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TrickleIceSdpfrag string                 `protobuf:"bytes,1,opt,name=trickle_ice_sdpfrag,json=trickleIceSdpfrag,proto3" json:"trickle_ice_sdpfrag,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ICERestartWHIPResourceResponse) Reset() {
	*x = ICERestartWHIPResourceResponse{}
	mi := &file_rpc_ingress_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ICERestartWHIPResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ICERestartWHIPResourceResponse) ProtoMessage() {}

func (x *ICERestartWHIPResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ICERestartWHIPResourceResponse.ProtoReflect.Descriptor instead.
func (*ICERestartWHIPResourceResponse) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_proto_rawDescGZIP(), []int{4}
}

func (x *ICERestartWHIPResourceResponse) GetTrickleIceSdpfrag() string {
	if x != nil {
		return x.TrickleIceSdpfrag
	}
	return ""
}

type StartIngressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          *livekit.IngressInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	WsUrl         string                 `protobuf:"bytes,3,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
	LoggingFields map[string]string      `protobuf:"bytes,4,rep,name=logging_fields,json=loggingFields,proto3" json:"logging_fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartIngressRequest) Reset() {
	*x = StartIngressRequest{}
	mi := &file_rpc_ingress_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartIngressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartIngressRequest) ProtoMessage() {}

func (x *StartIngressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[5]
	if x != nil {
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
	return file_rpc_ingress_proto_rawDescGZIP(), []int{5}
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

type IngressSession struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IngressId     string                 `protobuf:"bytes,1,opt,name=ingress_id,json=ingressId,proto3" json:"ingress_id,omitempty"`
	ResourceId    string                 `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IngressSession) Reset() {
	*x = IngressSession{}
	mi := &file_rpc_ingress_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IngressSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressSession) ProtoMessage() {}

func (x *IngressSession) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressSession.ProtoReflect.Descriptor instead.
func (*IngressSession) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_proto_rawDescGZIP(), []int{6}
}

func (x *IngressSession) GetIngressId() string {
	if x != nil {
		return x.IngressId
	}
	return ""
}

func (x *IngressSession) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type KillIngressSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Session       *IngressSession        `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KillIngressSessionRequest) Reset() {
	*x = KillIngressSessionRequest{}
	mi := &file_rpc_ingress_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KillIngressSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillIngressSessionRequest) ProtoMessage() {}

func (x *KillIngressSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillIngressSessionRequest.ProtoReflect.Descriptor instead.
func (*KillIngressSessionRequest) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_proto_rawDescGZIP(), []int{7}
}

func (x *KillIngressSessionRequest) GetSession() *IngressSession {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_rpc_ingress_proto protoreflect.FileDescriptor

const file_rpc_ingress_proto_rawDesc = "" +
	"\n" +
	"\x11rpc/ingress.proto\x12\x03rpc\x1a\roptions.proto\x1a\x15livekit_ingress.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x1a\n" +
	"\x18ListActiveIngressRequest\"\x80\x01\n" +
	"\x19ListActiveIngressResponse\x12#\n" +
	"\vingress_ids\x18\x01 \x03(\tB\x02\x18\x01R\n" +
	"ingressIds\x12>\n" +
	"\x10ingress_sessions\x18\x02 \x03(\v2\x13.rpc.IngressSessionR\x0fingressSessions\"[\n" +
	"\x19DeleteWHIPResourceRequest\x12\x1f\n" +
	"\vresource_id\x18\x01 \x01(\tR\n" +
	"resourceId\x12\x1d\n" +
	"\n" +
	"stream_key\x18\x02 \x01(\tR\tstreamKey\"\xc0\x01\n" +
	"\x1dICERestartWHIPResourceRequest\x12\x1f\n" +
	"\vresource_id\x18\x01 \x01(\tR\n" +
	"resourceId\x12\x1d\n" +
	"\n" +
	"stream_key\x18\x02 \x01(\tR\tstreamKey\x12#\n" +
	"\ruser_fragment\x18\x03 \x01(\tR\fuserFragment\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\x12\x1e\n" +
	"\n" +
	"candidates\x18\x05 \x03(\tR\n" +
	"candidates\"P\n" +
	"\x1eICERestartWHIPResourceResponse\x12.\n" +
	"\x13trickle_ice_sdpfrag\x18\x01 \x01(\tR\x11trickleIceSdpfrag\"\x82\x02\n" +
	"\x13StartIngressRequest\x12(\n" +
	"\x04info\x18\x01 \x01(\v2\x14.livekit.IngressInfoR\x04info\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\x12\x15\n" +
	"\x06ws_url\x18\x03 \x01(\tR\x05wsUrl\x12R\n" +
	"\x0elogging_fields\x18\x04 \x03(\v2+.rpc.StartIngressRequest.LoggingFieldsEntryR\rloggingFields\x1a@\n" +
	"\x12LoggingFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"P\n" +
	"\x0eIngressSession\x12\x1d\n" +
	"\n" +
	"ingress_id\x18\x01 \x01(\tR\tingressId\x12\x1f\n" +
	"\vresource_id\x18\x02 \x01(\tR\n" +
	"resourceId\"J\n" +
	"\x19KillIngressSessionRequest\x12-\n" +
	"\asession\x18\x01 \x01(\v2\x13.rpc.IngressSessionR\asession2\xa8\x02\n" +
	"\x0fIngressInternal\x12F\n" +
	"\fStartIngress\x12\x18.rpc.StartIngressRequest\x1a\x14.livekit.IngressInfo\"\x06\xb2\x89\x01\x020\x01\x12\\\n" +
	"\x11ListActiveIngress\x12\x1d.rpc.ListActiveIngressRequest\x1a\x1e.rpc.ListActiveIngressResponse\"\b\xb2\x89\x01\x04\x10\x01(\x01\x12o\n" +
	"\x12KillIngressSession\x12\x1e.rpc.KillIngressSessionRequest\x1a\x16.google.protobuf.Empty\"!\xb2\x89\x01\x1d\x10\x01\x1a\x19\x12\n" +
	"ingress_id\x12\vresource_id2\xef\x02\n" +
	"\x0eIngressHandler\x12M\n" +
	"\rUpdateIngress\x12\x1d.livekit.UpdateIngressRequest\x1a\x15.livekit.IngressState\"\x06\xb2\x89\x01\x02\x10\x01\x12M\n" +
	"\rDeleteIngress\x12\x1d.livekit.DeleteIngressRequest\x1a\x15.livekit.IngressState\"\x06\xb2\x89\x01\x02\x10\x01\x12T\n" +
	"\x12DeleteWHIPResource\x12\x1e.rpc.DeleteWHIPResourceRequest\x1a\x16.google.protobuf.Empty\"\x06\xb2\x89\x01\x02\x10\x01\x12i\n" +
	"\x16ICERestartWHIPResource\x12\".rpc.ICERestartWHIPResourceRequest\x1a#.rpc.ICERestartWHIPResourceResponse\"\x06\xb2\x89\x01\x02\x10\x01B!Z\x1fgithub.com/livekit/protocol/rpcb\x06proto3"

var (
	file_rpc_ingress_proto_rawDescOnce sync.Once
	file_rpc_ingress_proto_rawDescData []byte
)

func file_rpc_ingress_proto_rawDescGZIP() []byte {
	file_rpc_ingress_proto_rawDescOnce.Do(func() {
		file_rpc_ingress_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_ingress_proto_rawDesc), len(file_rpc_ingress_proto_rawDesc)))
	})
	return file_rpc_ingress_proto_rawDescData
}

var file_rpc_ingress_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpc_ingress_proto_goTypes = []any{
	(*ListActiveIngressRequest)(nil),       // 0: rpc.ListActiveIngressRequest
	(*ListActiveIngressResponse)(nil),      // 1: rpc.ListActiveIngressResponse
	(*DeleteWHIPResourceRequest)(nil),      // 2: rpc.DeleteWHIPResourceRequest
	(*ICERestartWHIPResourceRequest)(nil),  // 3: rpc.ICERestartWHIPResourceRequest
	(*ICERestartWHIPResourceResponse)(nil), // 4: rpc.ICERestartWHIPResourceResponse
	(*StartIngressRequest)(nil),            // 5: rpc.StartIngressRequest
	(*IngressSession)(nil),                 // 6: rpc.IngressSession
	(*KillIngressSessionRequest)(nil),      // 7: rpc.KillIngressSessionRequest
	nil,                                    // 8: rpc.StartIngressRequest.LoggingFieldsEntry
	(*livekit.IngressInfo)(nil),            // 9: livekit.IngressInfo
	(*livekit.UpdateIngressRequest)(nil),   // 10: livekit.UpdateIngressRequest
	(*livekit.DeleteIngressRequest)(nil),   // 11: livekit.DeleteIngressRequest
	(*emptypb.Empty)(nil),                  // 12: google.protobuf.Empty
	(*livekit.IngressState)(nil),           // 13: livekit.IngressState
}
var file_rpc_ingress_proto_depIdxs = []int32{
	6,  // 0: rpc.ListActiveIngressResponse.ingress_sessions:type_name -> rpc.IngressSession
	9,  // 1: rpc.StartIngressRequest.info:type_name -> livekit.IngressInfo
	8,  // 2: rpc.StartIngressRequest.logging_fields:type_name -> rpc.StartIngressRequest.LoggingFieldsEntry
	6,  // 3: rpc.KillIngressSessionRequest.session:type_name -> rpc.IngressSession
	5,  // 4: rpc.IngressInternal.StartIngress:input_type -> rpc.StartIngressRequest
	0,  // 5: rpc.IngressInternal.ListActiveIngress:input_type -> rpc.ListActiveIngressRequest
	7,  // 6: rpc.IngressInternal.KillIngressSession:input_type -> rpc.KillIngressSessionRequest
	10, // 7: rpc.IngressHandler.UpdateIngress:input_type -> livekit.UpdateIngressRequest
	11, // 8: rpc.IngressHandler.DeleteIngress:input_type -> livekit.DeleteIngressRequest
	2,  // 9: rpc.IngressHandler.DeleteWHIPResource:input_type -> rpc.DeleteWHIPResourceRequest
	3,  // 10: rpc.IngressHandler.ICERestartWHIPResource:input_type -> rpc.ICERestartWHIPResourceRequest
	9,  // 11: rpc.IngressInternal.StartIngress:output_type -> livekit.IngressInfo
	1,  // 12: rpc.IngressInternal.ListActiveIngress:output_type -> rpc.ListActiveIngressResponse
	12, // 13: rpc.IngressInternal.KillIngressSession:output_type -> google.protobuf.Empty
	13, // 14: rpc.IngressHandler.UpdateIngress:output_type -> livekit.IngressState
	13, // 15: rpc.IngressHandler.DeleteIngress:output_type -> livekit.IngressState
	12, // 16: rpc.IngressHandler.DeleteWHIPResource:output_type -> google.protobuf.Empty
	4,  // 17: rpc.IngressHandler.ICERestartWHIPResource:output_type -> rpc.ICERestartWHIPResourceResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_ingress_proto_init() }
func file_rpc_ingress_proto_init() {
	if File_rpc_ingress_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_ingress_proto_rawDesc), len(file_rpc_ingress_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_rpc_ingress_proto_goTypes,
		DependencyIndexes: file_rpc_ingress_proto_depIdxs,
		MessageInfos:      file_rpc_ingress_proto_msgTypes,
	}.Build()
	File_rpc_ingress_proto = out.File
	file_rpc_ingress_proto_goTypes = nil
	file_rpc_ingress_proto_depIdxs = nil
}
