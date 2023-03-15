// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: livekit_rpc_internal.proto

package livekit

import (
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

// Deprecated: Do not use.
type StartEgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request metadata
	EgressId  string `protobuf:"bytes,1,opt,name=egress_id,json=egressId,proto3" json:"egress_id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SenderId  string `protobuf:"bytes,10,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	SentAt    int64  `protobuf:"varint,4,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	// request
	//
	// Types that are assignable to Request:
	//
	//	*StartEgressRequest_RoomComposite
	//	*StartEgressRequest_TrackComposite
	//	*StartEgressRequest_Track
	//	*StartEgressRequest_Web
	Request isStartEgressRequest_Request `protobuf_oneof:"request"`
	// connection info
	RoomId string `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Token  string `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	WsUrl  string `protobuf:"bytes,9,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
}

func (x *StartEgressRequest) Reset() {
	*x = StartEgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartEgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartEgressRequest) ProtoMessage() {}

func (x *StartEgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartEgressRequest.ProtoReflect.Descriptor instead.
func (*StartEgressRequest) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{0}
}

func (x *StartEgressRequest) GetEgressId() string {
	if x != nil {
		return x.EgressId
	}
	return ""
}

func (x *StartEgressRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *StartEgressRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *StartEgressRequest) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

func (m *StartEgressRequest) GetRequest() isStartEgressRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *StartEgressRequest) GetRoomComposite() *RoomCompositeEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_RoomComposite); ok {
		return x.RoomComposite
	}
	return nil
}

func (x *StartEgressRequest) GetTrackComposite() *TrackCompositeEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_TrackComposite); ok {
		return x.TrackComposite
	}
	return nil
}

func (x *StartEgressRequest) GetTrack() *TrackEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_Track); ok {
		return x.Track
	}
	return nil
}

func (x *StartEgressRequest) GetWeb() *WebEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_Web); ok {
		return x.Web
	}
	return nil
}

func (x *StartEgressRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *StartEgressRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StartEgressRequest) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

type isStartEgressRequest_Request interface {
	isStartEgressRequest_Request()
}

type StartEgressRequest_RoomComposite struct {
	RoomComposite *RoomCompositeEgressRequest `protobuf:"bytes,5,opt,name=room_composite,json=roomComposite,proto3,oneof"`
}

type StartEgressRequest_TrackComposite struct {
	TrackComposite *TrackCompositeEgressRequest `protobuf:"bytes,6,opt,name=track_composite,json=trackComposite,proto3,oneof"`
}

type StartEgressRequest_Track struct {
	Track *TrackEgressRequest `protobuf:"bytes,7,opt,name=track,proto3,oneof"`
}

type StartEgressRequest_Web struct {
	Web *WebEgressRequest `protobuf:"bytes,11,opt,name=web,proto3,oneof"`
}

func (*StartEgressRequest_RoomComposite) isStartEgressRequest_Request() {}

func (*StartEgressRequest_TrackComposite) isStartEgressRequest_Request() {}

func (*StartEgressRequest_Track) isStartEgressRequest_Request() {}

func (*StartEgressRequest_Web) isStartEgressRequest_Request() {}

// Deprecated: Do not use.
type EgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request metadata
	EgressId  string `protobuf:"bytes,1,opt,name=egress_id,json=egressId,proto3" json:"egress_id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SenderId  string `protobuf:"bytes,5,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// request
	//
	// Types that are assignable to Request:
	//
	//	*EgressRequest_UpdateStream
	//	*EgressRequest_Stop
	Request isEgressRequest_Request `protobuf_oneof:"request"`
}

func (x *EgressRequest) Reset() {
	*x = EgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressRequest) ProtoMessage() {}

func (x *EgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressRequest.ProtoReflect.Descriptor instead.
func (*EgressRequest) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{1}
}

func (x *EgressRequest) GetEgressId() string {
	if x != nil {
		return x.EgressId
	}
	return ""
}

func (x *EgressRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *EgressRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (m *EgressRequest) GetRequest() isEgressRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *EgressRequest) GetUpdateStream() *UpdateStreamRequest {
	if x, ok := x.GetRequest().(*EgressRequest_UpdateStream); ok {
		return x.UpdateStream
	}
	return nil
}

func (x *EgressRequest) GetStop() *StopEgressRequest {
	if x, ok := x.GetRequest().(*EgressRequest_Stop); ok {
		return x.Stop
	}
	return nil
}

type isEgressRequest_Request interface {
	isEgressRequest_Request()
}

type EgressRequest_UpdateStream struct {
	UpdateStream *UpdateStreamRequest `protobuf:"bytes,3,opt,name=update_stream,json=updateStream,proto3,oneof"`
}

type EgressRequest_Stop struct {
	Stop *StopEgressRequest `protobuf:"bytes,4,opt,name=stop,proto3,oneof"`
}

func (*EgressRequest_UpdateStream) isEgressRequest_Request() {}

func (*EgressRequest_Stop) isEgressRequest_Request() {}

// Deprecated: Do not use.
type EgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *EgressInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Error     string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	RequestId string      `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *EgressResponse) Reset() {
	*x = EgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressResponse) ProtoMessage() {}

func (x *EgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressResponse.ProtoReflect.Descriptor instead.
func (*EgressResponse) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{2}
}

func (x *EgressResponse) GetInfo() *EgressInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *EgressResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *EgressResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_livekit_rpc_internal_proto protoreflect.FileDescriptor

var file_livekit_rpc_internal_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x03, 0x0a, 0x12,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x12, 0x4c, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x2d, 0x0a, 0x03, 0x77, 0x65, 0x62, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x57, 0x65, 0x62, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x03, 0x77, 0x65, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x77, 0x73, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x73, 0x55, 0x72, 0x6c, 0x3a, 0x02, 0x18,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xee, 0x01, 0x0a,
	0x0d, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x04,
	0x73, 0x74, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x3a, 0x02,
	0x18, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x72, 0x0a,
	0x0e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x3a, 0x02, 0x18,
	0x01, 0x42, 0x46, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0xaa, 0x02, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x4b,
	0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xea, 0x02, 0x0e, 0x4c, 0x69, 0x76, 0x65, 0x4b,
	0x69, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_livekit_rpc_internal_proto_rawDescOnce sync.Once
	file_livekit_rpc_internal_proto_rawDescData = file_livekit_rpc_internal_proto_rawDesc
)

func file_livekit_rpc_internal_proto_rawDescGZIP() []byte {
	file_livekit_rpc_internal_proto_rawDescOnce.Do(func() {
		file_livekit_rpc_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_livekit_rpc_internal_proto_rawDescData)
	})
	return file_livekit_rpc_internal_proto_rawDescData
}

var file_livekit_rpc_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_livekit_rpc_internal_proto_goTypes = []interface{}{
	(*StartEgressRequest)(nil),          // 0: livekit.StartEgressRequest
	(*EgressRequest)(nil),               // 1: livekit.EgressRequest
	(*EgressResponse)(nil),              // 2: livekit.EgressResponse
	(*RoomCompositeEgressRequest)(nil),  // 3: livekit.RoomCompositeEgressRequest
	(*TrackCompositeEgressRequest)(nil), // 4: livekit.TrackCompositeEgressRequest
	(*TrackEgressRequest)(nil),          // 5: livekit.TrackEgressRequest
	(*WebEgressRequest)(nil),            // 6: livekit.WebEgressRequest
	(*UpdateStreamRequest)(nil),         // 7: livekit.UpdateStreamRequest
	(*StopEgressRequest)(nil),           // 8: livekit.StopEgressRequest
	(*EgressInfo)(nil),                  // 9: livekit.EgressInfo
}
var file_livekit_rpc_internal_proto_depIdxs = []int32{
	3, // 0: livekit.StartEgressRequest.room_composite:type_name -> livekit.RoomCompositeEgressRequest
	4, // 1: livekit.StartEgressRequest.track_composite:type_name -> livekit.TrackCompositeEgressRequest
	5, // 2: livekit.StartEgressRequest.track:type_name -> livekit.TrackEgressRequest
	6, // 3: livekit.StartEgressRequest.web:type_name -> livekit.WebEgressRequest
	7, // 4: livekit.EgressRequest.update_stream:type_name -> livekit.UpdateStreamRequest
	8, // 5: livekit.EgressRequest.stop:type_name -> livekit.StopEgressRequest
	9, // 6: livekit.EgressResponse.info:type_name -> livekit.EgressInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_livekit_rpc_internal_proto_init() }
func file_livekit_rpc_internal_proto_init() {
	if File_livekit_rpc_internal_proto != nil {
		return
	}
	file_livekit_egress_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_livekit_rpc_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartEgressRequest); i {
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
		file_livekit_rpc_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressRequest); i {
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
		file_livekit_rpc_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressResponse); i {
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
	file_livekit_rpc_internal_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StartEgressRequest_RoomComposite)(nil),
		(*StartEgressRequest_TrackComposite)(nil),
		(*StartEgressRequest_Track)(nil),
		(*StartEgressRequest_Web)(nil),
	}
	file_livekit_rpc_internal_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EgressRequest_UpdateStream)(nil),
		(*EgressRequest_Stop)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_livekit_rpc_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_livekit_rpc_internal_proto_goTypes,
		DependencyIndexes: file_livekit_rpc_internal_proto_depIdxs,
		MessageInfos:      file_livekit_rpc_internal_proto_msgTypes,
	}.Build()
	File_livekit_rpc_internal_proto = out.File
	file_livekit_rpc_internal_proto_rawDesc = nil
	file_livekit_rpc_internal_proto_goTypes = nil
	file_livekit_rpc_internal_proto_depIdxs = nil
}
