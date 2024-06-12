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
// 	protoc        v5.27.0
// source: rpc/sip.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
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

type InternalCreateSIPParticipantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SipCallId string `protobuf:"bytes,13,opt,name=sip_call_id,json=sipCallId,proto3" json:"sip_call_id,omitempty"`
	// IP that SIP INVITE is sent too
	Address   string               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Transport livekit.SIPTransport `protobuf:"varint,16,opt,name=transport,proto3,enum=livekit.SIPTransport" json:"transport,omitempty"`
	// Number used to make the call
	Number string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	// Number to call to
	CallTo                string            `protobuf:"bytes,4,opt,name=call_to,json=callTo,proto3" json:"call_to,omitempty"`
	Username              string            `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password              string            `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	RoomName              string            `protobuf:"bytes,7,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	ParticipantIdentity   string            `protobuf:"bytes,8,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
	ParticipantName       string            `protobuf:"bytes,14,opt,name=participant_name,json=participantName,proto3" json:"participant_name,omitempty"`
	ParticipantMetadata   string            `protobuf:"bytes,15,opt,name=participant_metadata,json=participantMetadata,proto3" json:"participant_metadata,omitempty"`
	ParticipantAttributes map[string]string `protobuf:"bytes,17,rep,name=participant_attributes,json=participantAttributes,proto3" json:"participant_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// optional token that should be used when creating LiveKit participant
	Token string `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	// optional websocket url that should be used when creating LiveKit participant
	WsUrl string `protobuf:"bytes,10,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
	// Optionally send following DTMF digits (extension codes) when making a call.
	// Character 'w' can be used to add a 0.5 sec delay.
	Dtmf string `protobuf:"bytes,11,opt,name=dtmf,proto3" json:"dtmf,omitempty"`
	// Optionally play ringtone in the room as an audible indicator for existing participants
	PlayRingtone bool `protobuf:"varint,12,opt,name=play_ringtone,json=playRingtone,proto3" json:"play_ringtone,omitempty"`
}

func (x *InternalCreateSIPParticipantRequest) Reset() {
	*x = InternalCreateSIPParticipantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalCreateSIPParticipantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalCreateSIPParticipantRequest) ProtoMessage() {}

func (x *InternalCreateSIPParticipantRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use InternalCreateSIPParticipantRequest.ProtoReflect.Descriptor instead.
func (*InternalCreateSIPParticipantRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sip_proto_rawDescGZIP(), []int{0}
}

func (x *InternalCreateSIPParticipantRequest) GetSipCallId() string {
	if x != nil {
		return x.SipCallId
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetTransport() livekit.SIPTransport {
	if x != nil {
		return x.Transport
	}
	return livekit.SIPTransport(0)
}

func (x *InternalCreateSIPParticipantRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetCallTo() string {
	if x != nil {
		return x.CallTo
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetParticipantIdentity() string {
	if x != nil {
		return x.ParticipantIdentity
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetParticipantName() string {
	if x != nil {
		return x.ParticipantName
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetParticipantMetadata() string {
	if x != nil {
		return x.ParticipantMetadata
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetParticipantAttributes() map[string]string {
	if x != nil {
		return x.ParticipantAttributes
	}
	return nil
}

func (x *InternalCreateSIPParticipantRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetDtmf() string {
	if x != nil {
		return x.Dtmf
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetPlayRingtone() bool {
	if x != nil {
		return x.PlayRingtone
	}
	return false
}

type InternalCreateSIPParticipantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId       string `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	ParticipantIdentity string `protobuf:"bytes,2,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
	SipCallId           string `protobuf:"bytes,3,opt,name=sip_call_id,json=sipCallId,proto3" json:"sip_call_id,omitempty"`
}

func (x *InternalCreateSIPParticipantResponse) Reset() {
	*x = InternalCreateSIPParticipantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalCreateSIPParticipantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalCreateSIPParticipantResponse) ProtoMessage() {}

func (x *InternalCreateSIPParticipantResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use InternalCreateSIPParticipantResponse.ProtoReflect.Descriptor instead.
func (*InternalCreateSIPParticipantResponse) Descriptor() ([]byte, []int) {
	return file_rpc_sip_proto_rawDescGZIP(), []int{1}
}

func (x *InternalCreateSIPParticipantResponse) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *InternalCreateSIPParticipantResponse) GetParticipantIdentity() string {
	if x != nil {
		return x.ParticipantIdentity
	}
	return ""
}

func (x *InternalCreateSIPParticipantResponse) GetSipCallId() string {
	if x != nil {
		return x.SipCallId
	}
	return ""
}

var File_rpc_sip_proto protoreflect.FileDescriptor

var file_rpc_sip_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x73, 0x69, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x05, 0x0a, 0x23, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0b, 0x73, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x49, 0x50, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x7a, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x11,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x77, 0x73, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x74, 0x6d, 0x66, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x74, 0x6d,
	0x66, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x74, 0x6f,
	0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x69,
	0x6e, 0x67, 0x74, 0x6f, 0x6e, 0x65, 0x1a, 0x48, 0x0a, 0x1a, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xa0, 0x01, 0x0a, 0x24, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x70, 0x43, 0x61, 0x6c,
	0x6c, 0x49, 0x64, 0x32, 0x84, 0x01, 0x0a, 0x0b, 0x53, 0x49, 0x50, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x12, 0x75, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x08, 0xb2, 0x89, 0x01, 0x04, 0x10, 0x01, 0x30, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_rpc_sip_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_sip_proto_goTypes = []interface{}{
	(*InternalCreateSIPParticipantRequest)(nil),  // 0: rpc.InternalCreateSIPParticipantRequest
	(*InternalCreateSIPParticipantResponse)(nil), // 1: rpc.InternalCreateSIPParticipantResponse
	nil,                       // 2: rpc.InternalCreateSIPParticipantRequest.ParticipantAttributesEntry
	(livekit.SIPTransport)(0), // 3: livekit.SIPTransport
}
var file_rpc_sip_proto_depIdxs = []int32{
	3, // 0: rpc.InternalCreateSIPParticipantRequest.transport:type_name -> livekit.SIPTransport
	2, // 1: rpc.InternalCreateSIPParticipantRequest.participant_attributes:type_name -> rpc.InternalCreateSIPParticipantRequest.ParticipantAttributesEntry
	0, // 2: rpc.SIPInternal.CreateSIPParticipant:input_type -> rpc.InternalCreateSIPParticipantRequest
	1, // 3: rpc.SIPInternal.CreateSIPParticipant:output_type -> rpc.InternalCreateSIPParticipantResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_sip_proto_init() }
func file_rpc_sip_proto_init() {
	if File_rpc_sip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalCreateSIPParticipantRequest); i {
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
			switch v := v.(*InternalCreateSIPParticipantResponse); i {
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
			NumMessages:   3,
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
