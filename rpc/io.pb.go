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
// 	protoc        v4.23.4
// source: rpc/io.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
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

type GetEgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EgressId string `protobuf:"bytes,1,opt,name=egress_id,json=egressId,proto3" json:"egress_id,omitempty"`
}

func (x *GetEgressRequest) Reset() {
	*x = GetEgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEgressRequest) ProtoMessage() {}

func (x *GetEgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEgressRequest.ProtoReflect.Descriptor instead.
func (*GetEgressRequest) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{0}
}

func (x *GetEgressRequest) GetEgressId() string {
	if x != nil {
		return x.EgressId
	}
	return ""
}

// Query an ingress info from an ingress ID or stream key
type GetIngressInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngressId string `protobuf:"bytes,1,opt,name=ingress_id,json=ingressId,proto3" json:"ingress_id,omitempty"`
	StreamKey string `protobuf:"bytes,2,opt,name=stream_key,json=streamKey,proto3" json:"stream_key,omitempty"`
}

func (x *GetIngressInfoRequest) Reset() {
	*x = GetIngressInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIngressInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngressInfoRequest) ProtoMessage() {}

func (x *GetIngressInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngressInfoRequest.ProtoReflect.Descriptor instead.
func (*GetIngressInfoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{1}
}

func (x *GetIngressInfoRequest) GetIngressId() string {
	if x != nil {
		return x.IngressId
	}
	return ""
}

func (x *GetIngressInfoRequest) GetStreamKey() string {
	if x != nil {
		return x.StreamKey
	}
	return ""
}

type GetIngressInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *livekit.IngressInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Token string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	WsUrl string               `protobuf:"bytes,3,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
}

func (x *GetIngressInfoResponse) Reset() {
	*x = GetIngressInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIngressInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngressInfoResponse) ProtoMessage() {}

func (x *GetIngressInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngressInfoResponse.ProtoReflect.Descriptor instead.
func (*GetIngressInfoResponse) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{2}
}

func (x *GetIngressInfoResponse) GetInfo() *livekit.IngressInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *GetIngressInfoResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetIngressInfoResponse) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

// Request to store an update to the ingress state ingress -> service
type UpdateIngressStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngressId string                `protobuf:"bytes,1,opt,name=ingress_id,json=ingressId,proto3" json:"ingress_id,omitempty"`
	State     *livekit.IngressState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *UpdateIngressStateRequest) Reset() {
	*x = UpdateIngressStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIngressStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIngressStateRequest) ProtoMessage() {}

func (x *UpdateIngressStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIngressStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateIngressStateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateIngressStateRequest) GetIngressId() string {
	if x != nil {
		return x.IngressId
	}
	return ""
}

func (x *UpdateIngressStateRequest) GetState() *livekit.IngressState {
	if x != nil {
		return x.State
	}
	return nil
}

type GetSIPTrunkAuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// What Number is calling
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// What Number was called
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// What is the IP address of the called number
	SrcAddress string `protobuf:"bytes,4,opt,name=src_address,json=srcAddress,proto3" json:"src_address,omitempty"`
}

func (x *GetSIPTrunkAuthenticationRequest) Reset() {
	*x = GetSIPTrunkAuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSIPTrunkAuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSIPTrunkAuthenticationRequest) ProtoMessage() {}

func (x *GetSIPTrunkAuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSIPTrunkAuthenticationRequest.ProtoReflect.Descriptor instead.
func (*GetSIPTrunkAuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{4}
}

func (x *GetSIPTrunkAuthenticationRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GetSIPTrunkAuthenticationRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *GetSIPTrunkAuthenticationRequest) GetSrcAddress() string {
	if x != nil {
		return x.SrcAddress
	}
	return ""
}

type GetSIPTrunkAuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expected username and password
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetSIPTrunkAuthenticationResponse) Reset() {
	*x = GetSIPTrunkAuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSIPTrunkAuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSIPTrunkAuthenticationResponse) ProtoMessage() {}

func (x *GetSIPTrunkAuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSIPTrunkAuthenticationResponse.ProtoReflect.Descriptor instead.
func (*GetSIPTrunkAuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{5}
}

func (x *GetSIPTrunkAuthenticationResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetSIPTrunkAuthenticationResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type EvaluateSIPDispatchRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SipParticipantId string `protobuf:"bytes,1,opt,name=sip_participant_id,json=sipParticipantId,proto3" json:"sip_participant_id,omitempty"`
	// What Number is calling
	CallingNumber string `protobuf:"bytes,2,opt,name=calling_number,json=callingNumber,proto3" json:"calling_number,omitempty"`
	// What Number was called
	CalledNumber string `protobuf:"bytes,3,opt,name=called_number,json=calledNumber,proto3" json:"called_number,omitempty"`
	// What is the IP address of the called number
	SrcAddress string `protobuf:"bytes,4,opt,name=src_address,json=srcAddress,proto3" json:"src_address,omitempty"`
	// What pin has been entered if any
	Pin string `protobuf:"bytes,5,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *EvaluateSIPDispatchRulesRequest) Reset() {
	*x = EvaluateSIPDispatchRulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateSIPDispatchRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateSIPDispatchRulesRequest) ProtoMessage() {}

func (x *EvaluateSIPDispatchRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateSIPDispatchRulesRequest.ProtoReflect.Descriptor instead.
func (*EvaluateSIPDispatchRulesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{6}
}

func (x *EvaluateSIPDispatchRulesRequest) GetSipParticipantId() string {
	if x != nil {
		return x.SipParticipantId
	}
	return ""
}

func (x *EvaluateSIPDispatchRulesRequest) GetCallingNumber() string {
	if x != nil {
		return x.CallingNumber
	}
	return ""
}

func (x *EvaluateSIPDispatchRulesRequest) GetCalledNumber() string {
	if x != nil {
		return x.CalledNumber
	}
	return ""
}

func (x *EvaluateSIPDispatchRulesRequest) GetSrcAddress() string {
	if x != nil {
		return x.SrcAddress
	}
	return ""
}

func (x *EvaluateSIPDispatchRulesRequest) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

type EvaluateSIPDispatchRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// non-empty string if SIPParticipant should be placed a room
	RoomName string `protobuf:"bytes,1,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	// What should be used for the LiveKit identity
	ParticipantIdentity string `protobuf:"bytes,2,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
	// Pin should be requested from SIPParticipant
	RequestPin bool `protobuf:"varint,3,opt,name=request_pin,json=requestPin,proto3" json:"request_pin,omitempty"`
}

func (x *EvaluateSIPDispatchRulesResponse) Reset() {
	*x = EvaluateSIPDispatchRulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateSIPDispatchRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateSIPDispatchRulesResponse) ProtoMessage() {}

func (x *EvaluateSIPDispatchRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateSIPDispatchRulesResponse.ProtoReflect.Descriptor instead.
func (*EvaluateSIPDispatchRulesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{7}
}

func (x *EvaluateSIPDispatchRulesResponse) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *EvaluateSIPDispatchRulesResponse) GetParticipantIdentity() string {
	if x != nil {
		return x.ParticipantIdentity
	}
	return ""
}

func (x *EvaluateSIPDispatchRulesResponse) GetRequestPin() bool {
	if x != nil {
		return x.RequestPin
	}
	return false
}

var File_rpc_io_proto protoreflect.FileDescriptor

var file_rpc_io_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x70, 0x63, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x55,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4b, 0x65, 0x79, 0x22, 0x6f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x15, 0x0a, 0x06, 0x77, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x77, 0x73, 0x55, 0x72, 0x6c, 0x22, 0x67, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x67, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x53, 0x49, 0x50, 0x54, 0x72, 0x75, 0x6e, 0x6b, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x72, 0x63, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x72,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5b, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53,
	0x49, 0x50, 0x54, 0x72, 0x75, 0x6e, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x1f, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x65, 0x53, 0x49, 0x50, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x69, 0x70,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x69, 0x70, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x72, 0x63, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x20, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x69, 0x6e, 0x32, 0xf0, 0x04, 0x0a,
	0x06, 0x49, 0x4f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x15,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6a, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x53, 0x49, 0x50, 0x54, 0x72, 0x75, 0x6e, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x49, 0x50, 0x54, 0x72, 0x75, 0x6e, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x49, 0x50, 0x54, 0x72, 0x75, 0x6e, 0x6b,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x18, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x65, 0x53, 0x49, 0x50, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x12, 0x24, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x53, 0x49, 0x50, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_io_proto_rawDescOnce sync.Once
	file_rpc_io_proto_rawDescData = file_rpc_io_proto_rawDesc
)

func file_rpc_io_proto_rawDescGZIP() []byte {
	file_rpc_io_proto_rawDescOnce.Do(func() {
		file_rpc_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_io_proto_rawDescData)
	})
	return file_rpc_io_proto_rawDescData
}

var file_rpc_io_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_io_proto_goTypes = []interface{}{
	(*GetEgressRequest)(nil),                  // 0: rpc.GetEgressRequest
	(*GetIngressInfoRequest)(nil),             // 1: rpc.GetIngressInfoRequest
	(*GetIngressInfoResponse)(nil),            // 2: rpc.GetIngressInfoResponse
	(*UpdateIngressStateRequest)(nil),         // 3: rpc.UpdateIngressStateRequest
	(*GetSIPTrunkAuthenticationRequest)(nil),  // 4: rpc.GetSIPTrunkAuthenticationRequest
	(*GetSIPTrunkAuthenticationResponse)(nil), // 5: rpc.GetSIPTrunkAuthenticationResponse
	(*EvaluateSIPDispatchRulesRequest)(nil),   // 6: rpc.EvaluateSIPDispatchRulesRequest
	(*EvaluateSIPDispatchRulesResponse)(nil),  // 7: rpc.EvaluateSIPDispatchRulesResponse
	(*livekit.IngressInfo)(nil),               // 8: livekit.IngressInfo
	(*livekit.IngressState)(nil),              // 9: livekit.IngressState
	(*livekit.EgressInfo)(nil),                // 10: livekit.EgressInfo
	(*livekit.ListEgressRequest)(nil),         // 11: livekit.ListEgressRequest
	(*emptypb.Empty)(nil),                     // 12: google.protobuf.Empty
	(*livekit.ListEgressResponse)(nil),        // 13: livekit.ListEgressResponse
}
var file_rpc_io_proto_depIdxs = []int32{
	8,  // 0: rpc.GetIngressInfoResponse.info:type_name -> livekit.IngressInfo
	9,  // 1: rpc.UpdateIngressStateRequest.state:type_name -> livekit.IngressState
	10, // 2: rpc.IOInfo.CreateEgress:input_type -> livekit.EgressInfo
	10, // 3: rpc.IOInfo.UpdateEgress:input_type -> livekit.EgressInfo
	0,  // 4: rpc.IOInfo.GetEgress:input_type -> rpc.GetEgressRequest
	11, // 5: rpc.IOInfo.ListEgress:input_type -> livekit.ListEgressRequest
	1,  // 6: rpc.IOInfo.GetIngressInfo:input_type -> rpc.GetIngressInfoRequest
	3,  // 7: rpc.IOInfo.UpdateIngressState:input_type -> rpc.UpdateIngressStateRequest
	4,  // 8: rpc.IOInfo.GetSIPTrunkAuthentication:input_type -> rpc.GetSIPTrunkAuthenticationRequest
	6,  // 9: rpc.IOInfo.EvaluateSIPDispatchRules:input_type -> rpc.EvaluateSIPDispatchRulesRequest
	12, // 10: rpc.IOInfo.CreateEgress:output_type -> google.protobuf.Empty
	12, // 11: rpc.IOInfo.UpdateEgress:output_type -> google.protobuf.Empty
	10, // 12: rpc.IOInfo.GetEgress:output_type -> livekit.EgressInfo
	13, // 13: rpc.IOInfo.ListEgress:output_type -> livekit.ListEgressResponse
	2,  // 14: rpc.IOInfo.GetIngressInfo:output_type -> rpc.GetIngressInfoResponse
	12, // 15: rpc.IOInfo.UpdateIngressState:output_type -> google.protobuf.Empty
	5,  // 16: rpc.IOInfo.GetSIPTrunkAuthentication:output_type -> rpc.GetSIPTrunkAuthenticationResponse
	7,  // 17: rpc.IOInfo.EvaluateSIPDispatchRules:output_type -> rpc.EvaluateSIPDispatchRulesResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_io_proto_init() }
func file_rpc_io_proto_init() {
	if File_rpc_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEgressRequest); i {
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
		file_rpc_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIngressInfoRequest); i {
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
		file_rpc_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIngressInfoResponse); i {
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
		file_rpc_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIngressStateRequest); i {
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
		file_rpc_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSIPTrunkAuthenticationRequest); i {
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
		file_rpc_io_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSIPTrunkAuthenticationResponse); i {
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
		file_rpc_io_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluateSIPDispatchRulesRequest); i {
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
		file_rpc_io_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluateSIPDispatchRulesResponse); i {
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
			RawDescriptor: file_rpc_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_io_proto_goTypes,
		DependencyIndexes: file_rpc_io_proto_depIdxs,
		MessageInfos:      file_rpc_io_proto_msgTypes,
	}.Build()
	File_rpc_io_proto = out.File
	file_rpc_io_proto_rawDesc = nil
	file_rpc_io_proto_goTypes = nil
	file_rpc_io_proto_depIdxs = nil
}
