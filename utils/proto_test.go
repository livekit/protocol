package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/livekit/logger"
)

func TestProtoRedact(t *testing.T) {
	msg := &livekit.ParticipantInfo{
		Identity: "testIdentity",
		Name:     "testName",
		Metadata: "testMetadata",
		Attributes: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		Tracks: []*livekit.TrackInfo{
			{
				Sid:  "TR_AUDIO",
				Name: "Microphone",
			},
			{
				Sid:  "TR_VIDEO",
				Name: "Camera",
			},
		},
	}
	cloned := CloneProtoRedacted(msg)
	expectedCloned := &livekit.ParticipantInfo{
		Identity: "testIdentity",
		Tracks: []*livekit.TrackInfo{
			{
				Sid: "TR_AUDIO",
			},
			{
				Sid: "TR_VIDEO",
			},
		},
	}
	require.True(t, proto.Equal(expectedCloned, cloned))
}

func TestCloneProtoRedactedLivekitMaps(t *testing.T) {
	cases := []struct {
		name string
		msg  proto.Message
		want proto.Message
	}{
		{
			name: "top-level map<string,string>",
			msg:  &livekit.CreateRoomRequest{Name: "room", Tags: map[string]string{"k": "v"}},
			want: &livekit.CreateRoomRequest{Name: "room", Tags: map[string]string{"k": "v"}},
		},
		{
			name: "map<string,string> in oneof message",
			msg: &livekit.DataPacket{Value: &livekit.DataPacket_StreamHeader{StreamHeader: &livekit.DataStream_Header{
				StreamId:   "stream",
				Attributes: map[string]string{"k": "v"},
			}}},
			want: &livekit.DataPacket{Value: &livekit.DataPacket_StreamHeader{StreamHeader: &livekit.DataStream_Header{
				StreamId:   "stream",
				Attributes: map[string]string{"k": "v"},
			}}},
		},
		{
			name: "map<string,string> in signal offer",
			msg: &livekit.SignalRequest{Message: &livekit.SignalRequest_Offer{Offer: &livekit.SessionDescription{
				Type:         "offer",
				Sdp:          "v=0",
				MidToTrackId: map[string]string{"0": "TR_VCxyz"},
			}}},
			want: &livekit.SignalRequest{Message: &livekit.SignalRequest_Offer{Offer: &livekit.SessionDescription{
				Type:         "offer",
				Sdp:          "v=0",
				MidToTrackId: map[string]string{"0": "TR_VCxyz"},
			}}},
		},
		{
			name: "map<int32,uint32>",
			msg:  &livekit.RTPStats{Packets: 3, GapHistogram: map[int32]uint32{1: 2}},
			want: &livekit.RTPStats{Packets: 3, GapHistogram: map[int32]uint32{1: 2}},
		},
		{
			name: "unannotated map beside annotated fields",
			msg: &livekit.CreateSIPInboundTrunkRequest{Trunk: &livekit.SIPInboundTrunkInfo{
				Name:                "support line",
				Numbers:             []string{"+15105550100"},
				AuthUsername:        "user",
				AuthPassword:        "pass",
				Headers:             map[string]string{"X-Secret": "s"},
				HeadersToAttributes: map[string]string{"X-Customer-Id": "customer.id"},
			}},
			want: &livekit.CreateSIPInboundTrunkRequest{Trunk: &livekit.SIPInboundTrunkInfo{
				Name:                "support line",
				Numbers:             []string{"+15105550100"},
				HeadersToAttributes: map[string]string{"X-Customer-Id": "customer.id"},
			}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			before := proto.Clone(c.msg)
			got := CloneProtoRedacted(c.msg)
			require.True(t, proto.Equal(c.want, got), "got %v", got)
			require.True(t, proto.Equal(before, c.msg), "input mutated")
		})
	}
}

var scalarKinds = []descriptorpb.FieldDescriptorProto_Type{
	descriptorpb.FieldDescriptorProto_TYPE_BOOL,
	descriptorpb.FieldDescriptorProto_TYPE_ENUM,
	descriptorpb.FieldDescriptorProto_TYPE_INT32,
	descriptorpb.FieldDescriptorProto_TYPE_SINT32,
	descriptorpb.FieldDescriptorProto_TYPE_UINT32,
	descriptorpb.FieldDescriptorProto_TYPE_INT64,
	descriptorpb.FieldDescriptorProto_TYPE_SINT64,
	descriptorpb.FieldDescriptorProto_TYPE_UINT64,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED32,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED32,
	descriptorpb.FieldDescriptorProto_TYPE_FLOAT,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED64,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED64,
	descriptorpb.FieldDescriptorProto_TYPE_DOUBLE,
	descriptorpb.FieldDescriptorProto_TYPE_STRING,
	descriptorpb.FieldDescriptorProto_TYPE_BYTES,
}

func sensitive(s logger.Sensitivity) *descriptorpb.FieldOptions {
	opts := &descriptorpb.FieldOptions{}
	proto.SetExtension(opts, logger.E_Sensitivity, s)
	return opts
}

func mapField(
	msg *descriptorpb.DescriptorProto,
	name string,
	value *descriptorpb.FieldDescriptorProto,
	opts *descriptorpb.FieldOptions,
) {
	var entry strings.Builder
	for _, part := range strings.Split(name, "_") {
		entry.WriteString(strings.ToUpper(part[:1]) + part[1:])
	}
	entry.WriteString("Entry")
	key := &descriptorpb.FieldDescriptorProto{
		Name:   proto.String("key"),
		Number: proto.Int32(1),
		Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
	}
	value.Name, value.Number = proto.String("value"), proto.Int32(2)
	value.Label = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
	msg.NestedType = append(msg.NestedType, &descriptorpb.DescriptorProto{
		Name:    proto.String(entry.String()),
		Field:   []*descriptorpb.FieldDescriptorProto{key, value},
		Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
	})
	msg.Field = append(msg.Field, &descriptorpb.FieldDescriptorProto{
		Name:     proto.String(name),
		Number:   proto.Int32(int32(len(msg.Field) + 1)),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String("." + msg.GetName() + "." + entry.String()),
		Options:  opts,
	})
}

func plainField(
	msg *descriptorpb.DescriptorProto,
	name string,
	label descriptorpb.FieldDescriptorProto_Label,
	typ descriptorpb.FieldDescriptorProto_Type,
	typeName string,
	opts *descriptorpb.FieldOptions,
) {
	f := &descriptorpb.FieldDescriptorProto{
		Name:    proto.String(name),
		Number:  proto.Int32(int32(len(msg.Field) + 1)),
		Label:   label.Enum(),
		Type:    typ.Enum(),
		Options: opts,
	}
	if typeName != "" {
		f.TypeName = proto.String(typeName)
	}
	msg.Field = append(msg.Field, f)
}

func scalarFieldProto(k descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	f := &descriptorpb.FieldDescriptorProto{Type: k.Enum()}
	if k == descriptorpb.FieldDescriptorProto_TYPE_ENUM {
		f.TypeName = proto.String(".Color")
	}
	return f
}

// buildRedactTestFile synthesizes a schema with every map value kind, since livekit protos have no enum- or
// bytes-valued maps.
func buildRedactTestFile(t *testing.T) (root, leaf protoreflect.MessageDescriptor) {
	optional := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	repeated := descriptorpb.FieldDescriptorProto_LABEL_REPEATED
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING
	msgType := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	pii := sensitive(logger.Sensitivity_SENSITIVITY_PII)
	secret := sensitive(logger.Sensitivity_SENSITIVITY_SECRET)

	leafMsg := &descriptorpb.DescriptorProto{Name: proto.String("Leaf")}
	plainField(leafMsg, "id", optional, str, "", nil)
	plainField(leafMsg, "secret", optional, str, "", secret)
	mapField(leafMsg, "pii_map", scalarFieldProto(str), pii)
	mapField(leafMsg, "plain_map", scalarFieldProto(str), nil)

	rootMsg := &descriptorpb.DescriptorProto{Name: proto.String("Root")}
	for _, k := range scalarKinds {
		mapField(rootMsg, "m_"+strings.ToLower(strings.TrimPrefix(k.String(), "TYPE_")), scalarFieldProto(k), nil)
	}
	mapField(rootMsg, "pii_enum_map", scalarFieldProto(descriptorpb.FieldDescriptorProto_TYPE_ENUM), pii)
	mapField(rootMsg, "secret_bytes_map", scalarFieldProto(descriptorpb.FieldDescriptorProto_TYPE_BYTES), secret)
	mapField(rootMsg, "m_leaf", &descriptorpb.FieldDescriptorProto{Type: msgType.Enum(), TypeName: proto.String(".Leaf")}, nil)
	mapField(rootMsg, "pii_leaf_map", &descriptorpb.FieldDescriptorProto{Type: msgType.Enum(), TypeName: proto.String(".Leaf")}, pii)
	plainField(rootMsg, "l_leaf", repeated, msgType, ".Leaf", nil)
	plainField(rootMsg, "l_str", repeated, str, "", nil)
	plainField(rootMsg, "pii_l_str", repeated, str, "", pii)
	plainField(rootMsg, "leaf", optional, msgType, ".Leaf", nil)
	plainField(rootMsg, "secret", optional, str, "", secret)

	fdp := &descriptorpb.FileDescriptorProto{
		Name:   proto.String("utils_redact_test.proto"),
		Syntax: proto.String("proto3"),
		EnumType: []*descriptorpb.EnumDescriptorProto{{
			Name: proto.String("Color"),
			Value: []*descriptorpb.EnumValueDescriptorProto{
				{Name: proto.String("COLOR_UNSPECIFIED"), Number: proto.Int32(0)},
				{Name: proto.String("COLOR_RED"), Number: proto.Int32(1)},
			},
		}},
		MessageType: []*descriptorpb.DescriptorProto{leafMsg, rootMsg},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	require.NoError(t, err)
	return fd.Messages().ByName("Root"), fd.Messages().ByName("Leaf")
}

func sampleValue(k protoreflect.Kind) protoreflect.Value {
	switch k {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(true)
	case protoreflect.EnumKind:
		return protoreflect.ValueOfEnum(1)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(-7)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return protoreflect.ValueOfUint32(7)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(-7)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return protoreflect.ValueOfUint64(7)
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(1.5)
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(1.5)
	case protoreflect.StringKind:
		return protoreflect.ValueOfString("v")
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes([]byte("v"))
	}
	panic("unexpected kind " + k.String())
}

func setMapEntry(m protoreflect.Message, name string, key string, v protoreflect.Value) {
	fd := m.Descriptor().Fields().ByName(protoreflect.Name(name))
	m.Mutable(fd).Map().Set(protoreflect.ValueOfString(key).MapKey(), v)
}

func setField(m protoreflect.Message, name string, v protoreflect.Value) {
	m.Set(m.Descriptor().Fields().ByName(protoreflect.Name(name)), v)
}

func newLeaf(desc protoreflect.MessageDescriptor, redacted bool) protoreflect.Message {
	leaf := dynamicpb.NewMessage(desc)
	setField(leaf, "id", protoreflect.ValueOfString("id"))
	setMapEntry(leaf, "plain_map", "k", protoreflect.ValueOfString("v"))
	if !redacted {
		setField(leaf, "secret", protoreflect.ValueOfString("s"))
		setMapEntry(leaf, "pii_map", "k", protoreflect.ValueOfString("v"))
	}
	return leaf
}

func TestCloneProtoRedactedShapes(t *testing.T) {
	rootDesc, leafDesc := buildRedactTestFile(t)

	type tc struct {
		name string
		msg  proto.Message
		want proto.Message
	}
	var cases []tc

	for _, k := range scalarKinds {
		name := "m_" + strings.ToLower(strings.TrimPrefix(k.String(), "TYPE_"))
		fd := rootDesc.Fields().ByName(protoreflect.Name(name))
		msg := dynamicpb.NewMessage(rootDesc)
		setMapEntry(msg, name, "k", sampleValue(fd.MapValue().Kind()))
		cases = append(cases, tc{name: "map<string," + fd.MapValue().Kind().String() + ">", msg: msg, want: proto.Clone(msg)})
	}

	containers := func(redacted bool) *dynamicpb.Message {
		msg := dynamicpb.NewMessage(rootDesc)
		setMapEntry(msg, "m_leaf", "k", protoreflect.ValueOfMessage(newLeaf(leafDesc, redacted)))
		lst := msg.Mutable(rootDesc.Fields().ByName("l_leaf")).List()
		lst.Append(protoreflect.ValueOfMessage(newLeaf(leafDesc, redacted)))
		lst.Append(protoreflect.ValueOfMessage(newLeaf(leafDesc, redacted)))
		msg.Mutable(rootDesc.Fields().ByName("l_str")).List().Append(protoreflect.ValueOfString("v"))
		setField(msg, "leaf", protoreflect.ValueOfMessage(newLeaf(leafDesc, redacted)))
		setMapEntry(msg, "m_string", "k", protoreflect.ValueOfString("v"))
		setMapEntry(msg, "m_enum", "k", protoreflect.ValueOfEnum(1))
		if !redacted {
			setField(msg, "secret", protoreflect.ValueOfString("s"))
			setMapEntry(msg, "pii_enum_map", "k", protoreflect.ValueOfEnum(1))
			setMapEntry(msg, "secret_bytes_map", "k", protoreflect.ValueOfBytes([]byte("s")))
			setMapEntry(msg, "pii_leaf_map", "k", protoreflect.ValueOfMessage(newLeaf(leafDesc, false)))
			msg.Mutable(rootDesc.Fields().ByName("pii_l_str")).List().Append(protoreflect.ValueOfString("v"))
		}
		return msg
	}
	cases = append(cases, tc{name: "annotated fields in every container", msg: containers(false), want: containers(true)})

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			before := proto.Clone(c.msg)
			got := CloneProtoRedacted(c.msg)
			require.True(t, proto.Equal(c.want, got), "got %v", got)
			require.True(t, proto.Equal(before, c.msg), "input mutated")
		})
	}
}
