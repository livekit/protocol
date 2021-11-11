// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: analytics.proto

package livekit

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type StreamType int32

const (
	StreamType_UPSTREAM   StreamType = 0
	StreamType_DOWNSTREAM StreamType = 1
)

// Enum value maps for StreamType.
var (
	StreamType_name = map[int32]string{
		0: "UPSTREAM",
		1: "DOWNSTREAM",
	}
	StreamType_value = map[string]int32{
		"UPSTREAM":   0,
		"DOWNSTREAM": 1,
	}
)

func (x StreamType) Enum() *StreamType {
	p := new(StreamType)
	*p = x
	return p
}

func (x StreamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamType) Descriptor() protoreflect.EnumDescriptor {
	return file_analytics_proto_enumTypes[0].Descriptor()
}

func (StreamType) Type() protoreflect.EnumType {
	return &file_analytics_proto_enumTypes[0]
}

func (x StreamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamType.Descriptor instead.
func (StreamType) EnumDescriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{0}
}

type AnalyticsEventType int32

const (
	AnalyticsEventType_ROOM_CREATED       AnalyticsEventType = 0
	AnalyticsEventType_ROOM_ENDED         AnalyticsEventType = 1
	AnalyticsEventType_PARTICIPANT_JOINED AnalyticsEventType = 2
	AnalyticsEventType_PARTICIPANT_LEFT   AnalyticsEventType = 3
	AnalyticsEventType_TRACK_PUBLISHED    AnalyticsEventType = 4
	AnalyticsEventType_TRACK_UNPUBLISHED  AnalyticsEventType = 5
	AnalyticsEventType_TRACK_SUBSCRIBED   AnalyticsEventType = 6
	AnalyticsEventType_TRACK_UNSUBSCRIBED AnalyticsEventType = 7
)

// Enum value maps for AnalyticsEventType.
var (
	AnalyticsEventType_name = map[int32]string{
		0: "ROOM_CREATED",
		1: "ROOM_ENDED",
		2: "PARTICIPANT_JOINED",
		3: "PARTICIPANT_LEFT",
		4: "TRACK_PUBLISHED",
		5: "TRACK_UNPUBLISHED",
		6: "TRACK_SUBSCRIBED",
		7: "TRACK_UNSUBSCRIBED",
	}
	AnalyticsEventType_value = map[string]int32{
		"ROOM_CREATED":       0,
		"ROOM_ENDED":         1,
		"PARTICIPANT_JOINED": 2,
		"PARTICIPANT_LEFT":   3,
		"TRACK_PUBLISHED":    4,
		"TRACK_UNPUBLISHED":  5,
		"TRACK_SUBSCRIBED":   6,
		"TRACK_UNSUBSCRIBED": 7,
	}
)

func (x AnalyticsEventType) Enum() *AnalyticsEventType {
	p := new(AnalyticsEventType)
	*p = x
	return p
}

func (x AnalyticsEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnalyticsEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_analytics_proto_enumTypes[1].Descriptor()
}

func (AnalyticsEventType) Type() protoreflect.EnumType {
	return &file_analytics_proto_enumTypes[1]
}

func (x AnalyticsEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnalyticsEventType.Descriptor instead.
func (AnalyticsEventType) EnumDescriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{1}
}

type AnalyticsStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind          StreamType             `protobuf:"varint,1,opt,name=kind,proto3,enum=livekit.StreamType" json:"kind,omitempty"`
	TimeStamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Node          string                 `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	Sid           *string                `protobuf:"bytes,4,opt,name=sid,proto3,oneof" json:"sid,omitempty"`
	ProjectId     *string                `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3,oneof" json:"project_id,omitempty"`
	ParticipantId *string                `protobuf:"bytes,6,opt,name=participant_id,json=participantId,proto3,oneof" json:"participant_id,omitempty"`
	RoomName      *string                `protobuf:"bytes,7,opt,name=room_name,json=roomName,proto3,oneof" json:"room_name,omitempty"`
	Jitter        *float64               `protobuf:"fixed64,8,opt,name=jitter,proto3,oneof" json:"jitter,omitempty"`
	PacketLost    *uint64                `protobuf:"varint,9,opt,name=packet_lost,json=packetLost,proto3,oneof" json:"packet_lost,omitempty"`
	RrTime        *uint64                `protobuf:"varint,10,opt,name=rr_time,json=rrTime,proto3,oneof" json:"rr_time,omitempty"`
	BytesSent     *uint64                `protobuf:"varint,11,opt,name=bytes_sent,json=bytesSent,proto3,oneof" json:"bytes_sent,omitempty"`
	BytesReceived *uint64                `protobuf:"varint,12,opt,name=bytes_received,json=bytesReceived,proto3,oneof" json:"bytes_received,omitempty"`
	AuthToken     *string                `protobuf:"bytes,13,opt,name=auth_token,json=authToken,proto3,oneof" json:"auth_token,omitempty"`
}

func (x *AnalyticsStat) Reset() {
	*x = AnalyticsStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsStat) ProtoMessage() {}

func (x *AnalyticsStat) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsStat.ProtoReflect.Descriptor instead.
func (*AnalyticsStat) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *AnalyticsStat) GetKind() StreamType {
	if x != nil {
		return x.Kind
	}
	return StreamType_UPSTREAM
}

func (x *AnalyticsStat) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *AnalyticsStat) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *AnalyticsStat) GetSid() string {
	if x != nil && x.Sid != nil {
		return *x.Sid
	}
	return ""
}

func (x *AnalyticsStat) GetProjectId() string {
	if x != nil && x.ProjectId != nil {
		return *x.ProjectId
	}
	return ""
}

func (x *AnalyticsStat) GetParticipantId() string {
	if x != nil && x.ParticipantId != nil {
		return *x.ParticipantId
	}
	return ""
}

func (x *AnalyticsStat) GetRoomName() string {
	if x != nil && x.RoomName != nil {
		return *x.RoomName
	}
	return ""
}

func (x *AnalyticsStat) GetJitter() float64 {
	if x != nil && x.Jitter != nil {
		return *x.Jitter
	}
	return 0
}

func (x *AnalyticsStat) GetPacketLost() uint64 {
	if x != nil && x.PacketLost != nil {
		return *x.PacketLost
	}
	return 0
}

func (x *AnalyticsStat) GetRrTime() uint64 {
	if x != nil && x.RrTime != nil {
		return *x.RrTime
	}
	return 0
}

func (x *AnalyticsStat) GetBytesSent() uint64 {
	if x != nil && x.BytesSent != nil {
		return *x.BytesSent
	}
	return 0
}

func (x *AnalyticsStat) GetBytesReceived() uint64 {
	if x != nil && x.BytesReceived != nil {
		return *x.BytesReceived
	}
	return 0
}

func (x *AnalyticsStat) GetAuthToken() string {
	if x != nil && x.AuthToken != nil {
		return *x.AuthToken
	}
	return ""
}

type AnalyticsStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*AnalyticsStat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *AnalyticsStats) Reset() {
	*x = AnalyticsStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsStats) ProtoMessage() {}

func (x *AnalyticsStats) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsStats.ProtoReflect.Descriptor instead.
func (*AnalyticsStats) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *AnalyticsStats) GetStats() []*AnalyticsStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type AnalyticsEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          AnalyticsEventType     `protobuf:"varint,1,opt,name=type,proto3,enum=livekit.AnalyticsEventType" json:"type,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RoomSid       string                 `protobuf:"bytes,3,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	Room          *Room                  `protobuf:"bytes,4,opt,name=room,proto3" json:"room,omitempty"`
	ParticipantId string                 `protobuf:"bytes,5,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	Participant   *ParticipantInfo       `protobuf:"bytes,6,opt,name=participant,proto3" json:"participant,omitempty"`
	Track         *TrackInfo             `protobuf:"bytes,7,opt,name=track,proto3" json:"track,omitempty"`
}

func (x *AnalyticsEvent) Reset() {
	*x = AnalyticsEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsEvent) ProtoMessage() {}

func (x *AnalyticsEvent) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsEvent.ProtoReflect.Descriptor instead.
func (*AnalyticsEvent) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *AnalyticsEvent) GetType() AnalyticsEventType {
	if x != nil {
		return x.Type
	}
	return AnalyticsEventType_ROOM_CREATED
}

func (x *AnalyticsEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AnalyticsEvent) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *AnalyticsEvent) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *AnalyticsEvent) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *AnalyticsEvent) GetParticipant() *ParticipantInfo {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *AnalyticsEvent) GetTrack() *TrackInfo {
	if x != nil {
		return x.Track
	}
	return nil
}

type AnalyticsEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*AnalyticsEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *AnalyticsEvents) Reset() {
	*x = AnalyticsEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsEvents) ProtoMessage() {}

func (x *AnalyticsEvents) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsEvents.ProtoReflect.Descriptor instead.
func (*AnalyticsEvents) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *AnalyticsEvents) GetEvents() []*AnalyticsEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_analytics_proto protoreflect.FileDescriptor

var file_analytics_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5,
	0x04, 0x0a, 0x0d, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x12, 0x27, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x73, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0d, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x04, 0x52, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x73,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x72, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x04, 0x48, 0x06, 0x52, 0x06, 0x72, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x48, 0x07, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x48, 0x08,
	0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x69, 0x64, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x73, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x73, 0x65, 0x6e, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x0e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x74, 0x61, 0x74, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0xc6, 0x02, 0x0a, 0x0e, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x22,
	0x42, 0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2a, 0x2a, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x44, 0x4f, 0x57, 0x4e, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x10, 0x01, 0x2a,
	0xbe, 0x01, 0x0a, 0x12, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4f, 0x4d,
	0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x41, 0x52, 0x54,
	0x49, 0x43, 0x49, 0x50, 0x41, 0x4e, 0x54, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x50, 0x41, 0x52, 0x54, 0x49, 0x43, 0x49, 0x50, 0x41, 0x4e, 0x54, 0x5f,
	0x4c, 0x45, 0x46, 0x54, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f,
	0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x54,
	0x52, 0x41, 0x43, 0x4b, 0x5f, 0x55, 0x4e, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44,
	0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x55, 0x42, 0x53,
	0x43, 0x52, 0x49, 0x42, 0x45, 0x44, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x52, 0x41, 0x43,
	0x4b, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x44, 0x10, 0x07,
	0x32, 0xa4, 0x01, 0x0a, 0x18, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a,
	0x0b, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x44, 0x0a, 0x0c, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analytics_proto_rawDescOnce sync.Once
	file_analytics_proto_rawDescData = file_analytics_proto_rawDesc
)

func file_analytics_proto_rawDescGZIP() []byte {
	file_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_proto_rawDescData)
	})
	return file_analytics_proto_rawDescData
}

var file_analytics_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_analytics_proto_goTypes = []interface{}{
	(StreamType)(0),               // 0: livekit.StreamType
	(AnalyticsEventType)(0),       // 1: livekit.AnalyticsEventType
	(*AnalyticsStat)(nil),         // 2: livekit.AnalyticsStat
	(*AnalyticsStats)(nil),        // 3: livekit.AnalyticsStats
	(*AnalyticsEvent)(nil),        // 4: livekit.AnalyticsEvent
	(*AnalyticsEvents)(nil),       // 5: livekit.AnalyticsEvents
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*Room)(nil),                  // 7: livekit.Room
	(*ParticipantInfo)(nil),       // 8: livekit.ParticipantInfo
	(*TrackInfo)(nil),             // 9: livekit.TrackInfo
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_analytics_proto_depIdxs = []int32{
	0,  // 0: livekit.AnalyticsStat.kind:type_name -> livekit.StreamType
	6,  // 1: livekit.AnalyticsStat.time_stamp:type_name -> google.protobuf.Timestamp
	2,  // 2: livekit.AnalyticsStats.stats:type_name -> livekit.AnalyticsStat
	1,  // 3: livekit.AnalyticsEvent.type:type_name -> livekit.AnalyticsEventType
	6,  // 4: livekit.AnalyticsEvent.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 5: livekit.AnalyticsEvent.room:type_name -> livekit.Room
	8,  // 6: livekit.AnalyticsEvent.participant:type_name -> livekit.ParticipantInfo
	9,  // 7: livekit.AnalyticsEvent.track:type_name -> livekit.TrackInfo
	4,  // 8: livekit.AnalyticsEvents.events:type_name -> livekit.AnalyticsEvent
	3,  // 9: livekit.AnalyticsRecorderService.IngestStats:input_type -> livekit.AnalyticsStats
	5,  // 10: livekit.AnalyticsRecorderService.IngestEvents:input_type -> livekit.AnalyticsEvents
	10, // 11: livekit.AnalyticsRecorderService.IngestStats:output_type -> google.protobuf.Empty
	10, // 12: livekit.AnalyticsRecorderService.IngestEvents:output_type -> google.protobuf.Empty
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_analytics_proto_init() }
func file_analytics_proto_init() {
	if File_analytics_proto != nil {
		return
	}
	file_livekit_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsStat); i {
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
		file_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsStats); i {
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
		file_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsEvent); i {
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
		file_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsEvents); i {
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
	file_analytics_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_analytics_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_proto_depIdxs,
		EnumInfos:         file_analytics_proto_enumTypes,
		MessageInfos:      file_analytics_proto_msgTypes,
	}.Build()
	File_analytics_proto = out.File
	file_analytics_proto_rawDesc = nil
	file_analytics_proto_goTypes = nil
	file_analytics_proto_depIdxs = nil
}
