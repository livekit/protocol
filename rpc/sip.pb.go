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
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.0
// source: rpc/sip.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type InternalCreateSIPParticipantRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Used in Cloud only
	ProjectId  string `protobuf:"bytes,18,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	SipCallId  string `protobuf:"bytes,13,opt,name=sip_call_id,json=sipCallId,proto3" json:"sip_call_id,omitempty"`
	SipTrunkId string `protobuf:"bytes,19,opt,name=sip_trunk_id,json=sipTrunkId,proto3" json:"sip_trunk_id,omitempty"`
	// IP or hostname that SIP INVITE is sent too
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Hostname for the 'From' SIP address in INVITE
	Hostname  string               `protobuf:"bytes,20,opt,name=hostname,proto3" json:"hostname,omitempty"`
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
	ParticipantAttributes map[string]string `protobuf:"bytes,17,rep,name=participant_attributes,json=participantAttributes,proto3" json:"participant_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// optional token that should be used when creating LiveKit participant
	Token string `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	// optional websocket url that should be used when creating LiveKit participant
	WsUrl string `protobuf:"bytes,10,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
	// Optionally send following DTMF digits (extension codes) when making a call.
	// Character 'w' can be used to add a 0.5 sec delay.
	Dtmf string `protobuf:"bytes,11,opt,name=dtmf,proto3" json:"dtmf,omitempty"`
	// Optionally play dialtone in the room as an audible indicator for existing participants
	PlayDialtone        bool              `protobuf:"varint,12,opt,name=play_dialtone,json=playDialtone,proto3" json:"play_dialtone,omitempty"`
	Headers             map[string]string `protobuf:"bytes,21,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	HeadersToAttributes map[string]string `protobuf:"bytes,22,rep,name=headers_to_attributes,json=headersToAttributes,proto3" json:"headers_to_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Map LiveKit attributes to SIP X-* headers when sending BYE or REFER requests.
	// Keys are the names of attributes and values are the names of X-* headers they will be mapped to.
	AttributesToHeaders map[string]string `protobuf:"bytes,26,rep,name=attributes_to_headers,json=attributesToHeaders,proto3" json:"attributes_to_headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Map SIP headers from 200 OK to sip.h.* participant attributes automatically.
	//
	// When the names of required headers is known, using headers_to_attributes is strongly recommended.
	//
	// When mapping 200 OK headers to follow-up request headers with attributes_to_headers map,
	// lowercase header names should be used, for example: sip.h.x-custom-header.
	IncludeHeaders  livekit.SIPHeaderOptions `protobuf:"varint,27,opt,name=include_headers,json=includeHeaders,proto3,enum=livekit.SIPHeaderOptions" json:"include_headers,omitempty"`
	EnabledFeatures []livekit.SIPFeature     `protobuf:"varint,25,rep,packed,name=enabled_features,json=enabledFeatures,proto3,enum=livekit.SIPFeature" json:"enabled_features,omitempty"`
	// Max time for the callee to answer the call.
	RingingTimeout *durationpb.Duration `protobuf:"bytes,23,opt,name=ringing_timeout,json=ringingTimeout,proto3" json:"ringing_timeout,omitempty"`
	// Max call duration.
	MaxCallDuration *durationpb.Duration       `protobuf:"bytes,24,opt,name=max_call_duration,json=maxCallDuration,proto3" json:"max_call_duration,omitempty"`
	MediaEncryption livekit.SIPMediaEncryption `protobuf:"varint,28,opt,name=media_encryption,json=mediaEncryption,proto3,enum=livekit.SIPMediaEncryption" json:"media_encryption,omitempty"`
	// Wait for the answer for the call before returning.
	WaitUntilAnswered bool `protobuf:"varint,29,opt,name=wait_until_answered,json=waitUntilAnswered,proto3" json:"wait_until_answered,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *InternalCreateSIPParticipantRequest) Reset() {
	*x = InternalCreateSIPParticipantRequest{}
	mi := &file_rpc_sip_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalCreateSIPParticipantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalCreateSIPParticipantRequest) ProtoMessage() {}

func (x *InternalCreateSIPParticipantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sip_proto_msgTypes[0]
	if x != nil {
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

func (x *InternalCreateSIPParticipantRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetSipCallId() string {
	if x != nil {
		return x.SipCallId
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetSipTrunkId() string {
	if x != nil {
		return x.SipTrunkId
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InternalCreateSIPParticipantRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
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

func (x *InternalCreateSIPParticipantRequest) GetPlayDialtone() bool {
	if x != nil {
		return x.PlayDialtone
	}
	return false
}

func (x *InternalCreateSIPParticipantRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *InternalCreateSIPParticipantRequest) GetHeadersToAttributes() map[string]string {
	if x != nil {
		return x.HeadersToAttributes
	}
	return nil
}

func (x *InternalCreateSIPParticipantRequest) GetAttributesToHeaders() map[string]string {
	if x != nil {
		return x.AttributesToHeaders
	}
	return nil
}

func (x *InternalCreateSIPParticipantRequest) GetIncludeHeaders() livekit.SIPHeaderOptions {
	if x != nil {
		return x.IncludeHeaders
	}
	return livekit.SIPHeaderOptions(0)
}

func (x *InternalCreateSIPParticipantRequest) GetEnabledFeatures() []livekit.SIPFeature {
	if x != nil {
		return x.EnabledFeatures
	}
	return nil
}

func (x *InternalCreateSIPParticipantRequest) GetRingingTimeout() *durationpb.Duration {
	if x != nil {
		return x.RingingTimeout
	}
	return nil
}

func (x *InternalCreateSIPParticipantRequest) GetMaxCallDuration() *durationpb.Duration {
	if x != nil {
		return x.MaxCallDuration
	}
	return nil
}

func (x *InternalCreateSIPParticipantRequest) GetMediaEncryption() livekit.SIPMediaEncryption {
	if x != nil {
		return x.MediaEncryption
	}
	return livekit.SIPMediaEncryption(0)
}

func (x *InternalCreateSIPParticipantRequest) GetWaitUntilAnswered() bool {
	if x != nil {
		return x.WaitUntilAnswered
	}
	return false
}

type InternalCreateSIPParticipantResponse struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	ParticipantId       string                 `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	ParticipantIdentity string                 `protobuf:"bytes,2,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
	SipCallId           string                 `protobuf:"bytes,3,opt,name=sip_call_id,json=sipCallId,proto3" json:"sip_call_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *InternalCreateSIPParticipantResponse) Reset() {
	*x = InternalCreateSIPParticipantResponse{}
	mi := &file_rpc_sip_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalCreateSIPParticipantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalCreateSIPParticipantResponse) ProtoMessage() {}

func (x *InternalCreateSIPParticipantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sip_proto_msgTypes[1]
	if x != nil {
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

type InternalTransferSIPParticipantRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	SipCallId  string                 `protobuf:"bytes,1,opt,name=sip_call_id,json=sipCallId,proto3" json:"sip_call_id,omitempty"`
	TransferTo string                 `protobuf:"bytes,2,opt,name=transfer_to,json=transferTo,proto3" json:"transfer_to,omitempty"`
	// Optionally play dialtone to the SIP participant as an audible indicator of being transferred
	PlayDialtone bool `protobuf:"varint,3,opt,name=play_dialtone,json=playDialtone,proto3" json:"play_dialtone,omitempty"`
	// Add the following headers to the REFER SIP request.
	Headers       map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalTransferSIPParticipantRequest) Reset() {
	*x = InternalTransferSIPParticipantRequest{}
	mi := &file_rpc_sip_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalTransferSIPParticipantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalTransferSIPParticipantRequest) ProtoMessage() {}

func (x *InternalTransferSIPParticipantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sip_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalTransferSIPParticipantRequest.ProtoReflect.Descriptor instead.
func (*InternalTransferSIPParticipantRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sip_proto_rawDescGZIP(), []int{2}
}

func (x *InternalTransferSIPParticipantRequest) GetSipCallId() string {
	if x != nil {
		return x.SipCallId
	}
	return ""
}

func (x *InternalTransferSIPParticipantRequest) GetTransferTo() string {
	if x != nil {
		return x.TransferTo
	}
	return ""
}

func (x *InternalTransferSIPParticipantRequest) GetPlayDialtone() bool {
	if x != nil {
		return x.PlayDialtone
	}
	return false
}

func (x *InternalTransferSIPParticipantRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_rpc_sip_proto protoreflect.FileDescriptor

var file_rpc_sip_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x0d, 0x0a, 0x23, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x69,
	0x70, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x69,
	0x70, 0x5f, 0x74, 0x72, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x69, 0x70, 0x54, 0x72, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x53, 0x49, 0x50, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x7a,
	0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x15, 0x0a, 0x06, 0x77, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x77, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x74, 0x6d, 0x66, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x74, 0x6d, 0x66, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x74, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x69, 0x61, 0x6c, 0x74, 0x6f, 0x6e, 0x65,
	0x12, 0x4f, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x75, 0x0a, 0x15, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x41, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x54, 0x6f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x13, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x75, 0x0a, 0x15, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x42, 0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x53, 0x49, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x49, 0x50, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x72, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x45, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x63,
	0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6d,
	0x61, 0x78, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46,
	0x0a, 0x10, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x53, 0x49, 0x50, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x75,
	0x6e, 0x74, 0x69, 0x6c, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x18, 0x1d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x77, 0x61, 0x69, 0x74, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x1a, 0x48, 0x0a, 0x1a, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x18,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x18, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa0, 0x01, 0x0a,
	0x24, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x14,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x1e, 0x0a, 0x0b, 0x73, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x22,
	0x9c, 0x02, 0x0a, 0x25, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x69, 0x70,
	0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x74, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x69, 0x61, 0x6c, 0x74, 0x6f, 0x6e, 0x65, 0x12,
	0x51, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xf9,
	0x01, 0x0a, 0x0b, 0x53, 0x49, 0x50, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x75,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xb2, 0x89, 0x01,
	0x04, 0x10, 0x01, 0x30, 0x01, 0x12, 0x73, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12,
	0x2a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x49, 0x50, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x15, 0xb2, 0x89, 0x01, 0x11, 0x10, 0x01, 0x1a, 0x0d, 0x12, 0x0b, 0x73,
	0x69, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69,
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

var file_rpc_sip_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_sip_proto_goTypes = []any{
	(*InternalCreateSIPParticipantRequest)(nil),   // 0: rpc.InternalCreateSIPParticipantRequest
	(*InternalCreateSIPParticipantResponse)(nil),  // 1: rpc.InternalCreateSIPParticipantResponse
	(*InternalTransferSIPParticipantRequest)(nil), // 2: rpc.InternalTransferSIPParticipantRequest
	nil,                             // 3: rpc.InternalCreateSIPParticipantRequest.ParticipantAttributesEntry
	nil,                             // 4: rpc.InternalCreateSIPParticipantRequest.HeadersEntry
	nil,                             // 5: rpc.InternalCreateSIPParticipantRequest.HeadersToAttributesEntry
	nil,                             // 6: rpc.InternalCreateSIPParticipantRequest.AttributesToHeadersEntry
	nil,                             // 7: rpc.InternalTransferSIPParticipantRequest.HeadersEntry
	(livekit.SIPTransport)(0),       // 8: livekit.SIPTransport
	(livekit.SIPHeaderOptions)(0),   // 9: livekit.SIPHeaderOptions
	(livekit.SIPFeature)(0),         // 10: livekit.SIPFeature
	(*durationpb.Duration)(nil),     // 11: google.protobuf.Duration
	(livekit.SIPMediaEncryption)(0), // 12: livekit.SIPMediaEncryption
	(*emptypb.Empty)(nil),           // 13: google.protobuf.Empty
}
var file_rpc_sip_proto_depIdxs = []int32{
	8,  // 0: rpc.InternalCreateSIPParticipantRequest.transport:type_name -> livekit.SIPTransport
	3,  // 1: rpc.InternalCreateSIPParticipantRequest.participant_attributes:type_name -> rpc.InternalCreateSIPParticipantRequest.ParticipantAttributesEntry
	4,  // 2: rpc.InternalCreateSIPParticipantRequest.headers:type_name -> rpc.InternalCreateSIPParticipantRequest.HeadersEntry
	5,  // 3: rpc.InternalCreateSIPParticipantRequest.headers_to_attributes:type_name -> rpc.InternalCreateSIPParticipantRequest.HeadersToAttributesEntry
	6,  // 4: rpc.InternalCreateSIPParticipantRequest.attributes_to_headers:type_name -> rpc.InternalCreateSIPParticipantRequest.AttributesToHeadersEntry
	9,  // 5: rpc.InternalCreateSIPParticipantRequest.include_headers:type_name -> livekit.SIPHeaderOptions
	10, // 6: rpc.InternalCreateSIPParticipantRequest.enabled_features:type_name -> livekit.SIPFeature
	11, // 7: rpc.InternalCreateSIPParticipantRequest.ringing_timeout:type_name -> google.protobuf.Duration
	11, // 8: rpc.InternalCreateSIPParticipantRequest.max_call_duration:type_name -> google.protobuf.Duration
	12, // 9: rpc.InternalCreateSIPParticipantRequest.media_encryption:type_name -> livekit.SIPMediaEncryption
	7,  // 10: rpc.InternalTransferSIPParticipantRequest.headers:type_name -> rpc.InternalTransferSIPParticipantRequest.HeadersEntry
	0,  // 11: rpc.SIPInternal.CreateSIPParticipant:input_type -> rpc.InternalCreateSIPParticipantRequest
	2,  // 12: rpc.SIPInternal.TransferSIPParticipant:input_type -> rpc.InternalTransferSIPParticipantRequest
	1,  // 13: rpc.SIPInternal.CreateSIPParticipant:output_type -> rpc.InternalCreateSIPParticipantResponse
	13, // 14: rpc.SIPInternal.TransferSIPParticipant:output_type -> google.protobuf.Empty
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_rpc_sip_proto_init() }
func file_rpc_sip_proto_init() {
	if File_rpc_sip_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_sip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
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
