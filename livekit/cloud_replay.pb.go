// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: cloud_replay.proto

package livekit

import (
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

type LoadReplayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcRoomId    string `protobuf:"bytes,1,opt,name=src_room_id,json=srcRoomId,proto3" json:"src_room_id,omitempty"`
	DestRoomName string `protobuf:"bytes,2,opt,name=dest_room_name,json=destRoomName,proto3" json:"dest_room_name,omitempty"`
	StartingPts  int64  `protobuf:"varint,3,opt,name=starting_pts,json=startingPts,proto3" json:"starting_pts,omitempty"`
}

func (x *LoadReplayRequest) Reset() {
	*x = LoadReplayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_replay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadReplayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadReplayRequest) ProtoMessage() {}

func (x *LoadReplayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_replay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadReplayRequest.ProtoReflect.Descriptor instead.
func (*LoadReplayRequest) Descriptor() ([]byte, []int) {
	return file_cloud_replay_proto_rawDescGZIP(), []int{0}
}

func (x *LoadReplayRequest) GetSrcRoomId() string {
	if x != nil {
		return x.SrcRoomId
	}
	return ""
}

func (x *LoadReplayRequest) GetDestRoomName() string {
	if x != nil {
		return x.DestRoomName
	}
	return ""
}

func (x *LoadReplayRequest) GetStartingPts() int64 {
	if x != nil {
		return x.StartingPts
	}
	return 0
}

type LoadReplayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplayId string `protobuf:"bytes,1,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
}

func (x *LoadReplayResponse) Reset() {
	*x = LoadReplayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_replay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadReplayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadReplayResponse) ProtoMessage() {}

func (x *LoadReplayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_replay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadReplayResponse.ProtoReflect.Descriptor instead.
func (*LoadReplayResponse) Descriptor() ([]byte, []int) {
	return file_cloud_replay_proto_rawDescGZIP(), []int{1}
}

func (x *LoadReplayResponse) GetReplayId() string {
	if x != nil {
		return x.ReplayId
	}
	return ""
}

type RoomSeekRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplayId string `protobuf:"bytes,1,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
	Pts      int64  `protobuf:"varint,2,opt,name=pts,proto3" json:"pts,omitempty"`
}

func (x *RoomSeekRequest) Reset() {
	*x = RoomSeekRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_replay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomSeekRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomSeekRequest) ProtoMessage() {}

func (x *RoomSeekRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_replay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomSeekRequest.ProtoReflect.Descriptor instead.
func (*RoomSeekRequest) Descriptor() ([]byte, []int) {
	return file_cloud_replay_proto_rawDescGZIP(), []int{2}
}

func (x *RoomSeekRequest) GetReplayId() string {
	if x != nil {
		return x.ReplayId
	}
	return ""
}

func (x *RoomSeekRequest) GetPts() int64 {
	if x != nil {
		return x.Pts
	}
	return 0
}

type CloseReplayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplayId string `protobuf:"bytes,1,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
}

func (x *CloseReplayRequest) Reset() {
	*x = CloseReplayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_replay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseReplayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseReplayRequest) ProtoMessage() {}

func (x *CloseReplayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_replay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseReplayRequest.ProtoReflect.Descriptor instead.
func (*CloseReplayRequest) Descriptor() ([]byte, []int) {
	return file_cloud_replay_proto_rawDescGZIP(), []int{3}
}

func (x *CloseReplayRequest) GetReplayId() string {
	if x != nil {
		return x.ReplayId
	}
	return ""
}

var File_cloud_replay_proto protoreflect.FileDescriptor

var file_cloud_replay_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7c, 0x0a, 0x11, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x72, 0x63, 0x5f, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x72, 0x63,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x74, 0x73, 0x22,
	0x31, 0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x49, 0x64, 0x22, 0x40, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x65, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x70, 0x74, 0x73, 0x22, 0x31, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x32, 0xf0, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x79, 0x12, 0x53, 0x0a, 0x0a, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x65, 0x65, 0x6b, 0x46,
	0x6f, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x65, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x49, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x22,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x46, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0xaa, 0x02, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0xea, 0x02, 0x0e, 0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_replay_proto_rawDescOnce sync.Once
	file_cloud_replay_proto_rawDescData = file_cloud_replay_proto_rawDesc
)

func file_cloud_replay_proto_rawDescGZIP() []byte {
	file_cloud_replay_proto_rawDescOnce.Do(func() {
		file_cloud_replay_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_replay_proto_rawDescData)
	})
	return file_cloud_replay_proto_rawDescData
}

var file_cloud_replay_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_replay_proto_goTypes = []interface{}{
	(*LoadReplayRequest)(nil),  // 0: cloud_protocol.LoadReplayRequest
	(*LoadReplayResponse)(nil), // 1: cloud_protocol.LoadReplayResponse
	(*RoomSeekRequest)(nil),    // 2: cloud_protocol.RoomSeekRequest
	(*CloseReplayRequest)(nil), // 3: cloud_protocol.CloseReplayRequest
	(*emptypb.Empty)(nil),      // 4: google.protobuf.Empty
}
var file_cloud_replay_proto_depIdxs = []int32{
	0, // 0: cloud_protocol.Replay.LoadReplay:input_type -> cloud_protocol.LoadReplayRequest
	2, // 1: cloud_protocol.Replay.SeekForRoom:input_type -> cloud_protocol.RoomSeekRequest
	3, // 2: cloud_protocol.Replay.CloseReplay:input_type -> cloud_protocol.CloseReplayRequest
	1, // 3: cloud_protocol.Replay.LoadReplay:output_type -> cloud_protocol.LoadReplayResponse
	4, // 4: cloud_protocol.Replay.SeekForRoom:output_type -> google.protobuf.Empty
	4, // 5: cloud_protocol.Replay.CloseReplay:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_replay_proto_init() }
func file_cloud_replay_proto_init() {
	if File_cloud_replay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_replay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadReplayRequest); i {
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
		file_cloud_replay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadReplayResponse); i {
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
		file_cloud_replay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomSeekRequest); i {
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
		file_cloud_replay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseReplayRequest); i {
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
			RawDescriptor: file_cloud_replay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_replay_proto_goTypes,
		DependencyIndexes: file_cloud_replay_proto_depIdxs,
		MessageInfos:      file_cloud_replay_proto_msgTypes,
	}.Build()
	File_cloud_replay_proto = out.File
	file_cloud_replay_proto_rawDesc = nil
	file_cloud_replay_proto_goTypes = nil
	file_cloud_replay_proto_depIdxs = nil
}
