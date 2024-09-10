// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: livekit_metrics.proto

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

// index from [0; 512] are for predefined labels (`MetricLabel`)
type MetricLabel int32

const (
	MetricLabel_AGENTS_LLM_TTFT        MetricLabel = 0 // time to first token from LLM
	MetricLabel_AGENTS_STT_TTFT        MetricLabel = 1 // time to final transcription
	MetricLabel_AGENTS_TTS_TTFB        MetricLabel = 2 // time to first byte
	MetricLabel_METRIC_LABEL_MAX_VALUE MetricLabel = 512
)

// Enum value maps for MetricLabel.
var (
	MetricLabel_name = map[int32]string{
		0:   "AGENTS_LLM_TTFT",
		1:   "AGENTS_STT_TTFT",
		2:   "AGENTS_TTS_TTFB",
		512: "METRIC_LABEL_MAX_VALUE",
	}
	MetricLabel_value = map[string]int32{
		"AGENTS_LLM_TTFT":        0,
		"AGENTS_STT_TTFT":        1,
		"AGENTS_TTS_TTFB":        2,
		"METRIC_LABEL_MAX_VALUE": 512,
	}
)

func (x MetricLabel) Enum() *MetricLabel {
	p := new(MetricLabel)
	*p = x
	return p
}

func (x MetricLabel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricLabel) Descriptor() protoreflect.EnumDescriptor {
	return file_livekit_metrics_proto_enumTypes[0].Descriptor()
}

func (MetricLabel) Type() protoreflect.EnumType {
	return &file_livekit_metrics_proto_enumTypes[0]
}

func (x MetricLabel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricLabel.Descriptor instead.
func (MetricLabel) EnumDescriptor() ([]byte, []int) {
	return file_livekit_metrics_proto_rawDescGZIP(), []int{0}
}

type MetricsBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // time at which this batch is sent based on a monotonic clock
	// To avoid repeating string values, we store them in a separate list and reference them by index
	// This is useful for storing participant identities, track names, etc.
	// There is also a predefined list of labels that can be used to reference common metrics.
	// They have reserved indices from 0 to 512
	StrData    []string            `protobuf:"bytes,2,rep,name=str_data,json=strData,proto3" json:"str_data,omitempty"`
	TimeSeries []*TimeSeriesMetric `protobuf:"bytes,3,rep,name=time_series,json=timeSeries,proto3" json:"time_series,omitempty"`
	Events     []*EventMetric      `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *MetricsBatch) Reset() {
	*x = MetricsBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsBatch) ProtoMessage() {}

func (x *MetricsBatch) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsBatch.ProtoReflect.Descriptor instead.
func (*MetricsBatch) Descriptor() ([]byte, []int) {
	return file_livekit_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *MetricsBatch) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MetricsBatch) GetStrData() []string {
	if x != nil {
		return x.StrData
	}
	return nil
}

func (x *MetricsBatch) GetTimeSeries() []*TimeSeriesMetric {
	if x != nil {
		return x.TimeSeries
	}
	return nil
}

func (x *MetricsBatch) GetEvents() []*EventMetric {
	if x != nil {
		return x.Events
	}
	return nil
}

type TimeSeriesMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metric name e.g "speech_probablity". The string value is not directly stored in the message, but referenced by index
	// in the `str_data` field of `MetricsBatch`
	Label               uint32          `protobuf:"varint,1,opt,name=label,proto3" json:"label,omitempty"`
	ParticipantIdentity uint32          `protobuf:"varint,2,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
	TrackSid            uint32          `protobuf:"varint,3,opt,name=track_sid,json=trackSid,proto3" json:"track_sid,omitempty"`
	Samples             []*MetricSample `protobuf:"bytes,4,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *TimeSeriesMetric) Reset() {
	*x = TimeSeriesMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesMetric) ProtoMessage() {}

func (x *TimeSeriesMetric) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesMetric.ProtoReflect.Descriptor instead.
func (*TimeSeriesMetric) Descriptor() ([]byte, []int) {
	return file_livekit_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *TimeSeriesMetric) GetLabel() uint32 {
	if x != nil {
		return x.Label
	}
	return 0
}

func (x *TimeSeriesMetric) GetParticipantIdentity() uint32 {
	if x != nil {
		return x.ParticipantIdentity
	}
	return 0
}

func (x *TimeSeriesMetric) GetTrackSid() uint32 {
	if x != nil {
		return x.TrackSid
	}
	return 0
}

func (x *TimeSeriesMetric) GetSamples() []*MetricSample {
	if x != nil {
		return x.Samples
	}
	return nil
}

type MetricSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // time of metric based on a monotonic clock
	Value     float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MetricSample) Reset() {
	*x = MetricSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricSample) ProtoMessage() {}

func (x *MetricSample) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricSample.ProtoReflect.Descriptor instead.
func (*MetricSample) Descriptor() ([]byte, []int) {
	return file_livekit_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *MetricSample) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MetricSample) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type EventMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label               uint32 `protobuf:"varint,1,opt,name=label,proto3" json:"label,omitempty"`
	ParticipantIdentity uint32 `protobuf:"varint,2,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
	TrackSid            uint32 `protobuf:"varint,3,opt,name=track_sid,json=trackSid,proto3" json:"track_sid,omitempty"`                   // 0 =
	StartTimestamp      int64  `protobuf:"varint,4,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"` // start time of event based on a monotonic clock
	EndTimestamp        *int64 `protobuf:"varint,5,opt,name=end_timestamp,json=endTimestamp,proto3,oneof" json:"end_timestamp,omitempty"` // end time of event based on a monotonic clock, if needed
	Metadata            string `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *EventMetric) Reset() {
	*x = EventMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetric) ProtoMessage() {}

func (x *EventMetric) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetric.ProtoReflect.Descriptor instead.
func (*EventMetric) Descriptor() ([]byte, []int) {
	return file_livekit_metrics_proto_rawDescGZIP(), []int{3}
}

func (x *EventMetric) GetLabel() uint32 {
	if x != nil {
		return x.Label
	}
	return 0
}

func (x *EventMetric) GetParticipantIdentity() uint32 {
	if x != nil {
		return x.ParticipantIdentity
	}
	return 0
}

func (x *EventMetric) GetTrackSid() uint32 {
	if x != nil {
		return x.TrackSid
	}
	return 0
}

func (x *EventMetric) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *EventMetric) GetEndTimestamp() int64 {
	if x != nil && x.EndTimestamp != nil {
		return *x.EndTimestamp
	}
	return 0
}

func (x *EventMetric) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

var File_livekit_metrics_proto protoreflect.FileDescriptor

var file_livekit_metrics_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x22, 0xb1, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x10, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x64, 0x12,
	0x2f, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x22, 0x42, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xf4, 0x01, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x31, 0x0a, 0x14, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x69, 0x0a, 0x0b, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x53, 0x5f, 0x4c, 0x4c, 0x4d, 0x5f, 0x54, 0x54, 0x46, 0x54, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x53, 0x54, 0x54, 0x5f, 0x54, 0x54,
	0x46, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x54,
	0x54, 0x53, 0x5f, 0x54, 0x54, 0x46, 0x42, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x16, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x10, 0x80, 0x04, 0x42, 0x46, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0xaa, 0x02, 0x0d,
	0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xea, 0x02, 0x0e,
	0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_livekit_metrics_proto_rawDescOnce sync.Once
	file_livekit_metrics_proto_rawDescData = file_livekit_metrics_proto_rawDesc
)

func file_livekit_metrics_proto_rawDescGZIP() []byte {
	file_livekit_metrics_proto_rawDescOnce.Do(func() {
		file_livekit_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_livekit_metrics_proto_rawDescData)
	})
	return file_livekit_metrics_proto_rawDescData
}

var file_livekit_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_livekit_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_livekit_metrics_proto_goTypes = []interface{}{
	(MetricLabel)(0),         // 0: livekit.MetricLabel
	(*MetricsBatch)(nil),     // 1: livekit.MetricsBatch
	(*TimeSeriesMetric)(nil), // 2: livekit.TimeSeriesMetric
	(*MetricSample)(nil),     // 3: livekit.MetricSample
	(*EventMetric)(nil),      // 4: livekit.EventMetric
}
var file_livekit_metrics_proto_depIdxs = []int32{
	2, // 0: livekit.MetricsBatch.time_series:type_name -> livekit.TimeSeriesMetric
	4, // 1: livekit.MetricsBatch.events:type_name -> livekit.EventMetric
	3, // 2: livekit.TimeSeriesMetric.samples:type_name -> livekit.MetricSample
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_livekit_metrics_proto_init() }
func file_livekit_metrics_proto_init() {
	if File_livekit_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_livekit_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsBatch); i {
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
		file_livekit_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesMetric); i {
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
		file_livekit_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricSample); i {
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
		file_livekit_metrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetric); i {
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
	file_livekit_metrics_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_livekit_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_livekit_metrics_proto_goTypes,
		DependencyIndexes: file_livekit_metrics_proto_depIdxs,
		EnumInfos:         file_livekit_metrics_proto_enumTypes,
		MessageInfos:      file_livekit_metrics_proto_msgTypes,
	}.Build()
	File_livekit_metrics_proto = out.File
	file_livekit_metrics_proto_rawDesc = nil
	file_livekit_metrics_proto_goTypes = nil
	file_livekit_metrics_proto_depIdxs = nil
}
