package decodebudget

import (
	"bytes"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils/protojson"
)

func TestProtoFields(t *testing.T) {
	desc := protoTestFile(t).Messages().ByName("Node")
	field := desc.Fields().ByName("name")
	encoded := protowire.AppendString(protowire.AppendTag(nil, field.Number(), protowire.BytesType), "value")
	budget := Proto{MaxFields: 3, MaxDepth: 3}
	require.NoError(t, budget.Validate(bytes.Repeat(encoded, 3), desc))

	msg := dynamicpb.NewMessage(desc)
	msg.Set(field, protoreflect.ValueOfString("preserved"))
	before := proto.Clone(msg)
	require.ErrorIs(t, budget.Unmarshal(bytes.Repeat(encoded, 4), msg), ErrMessageComplexity)
	require.True(t, proto.Equal(before, msg))

	unknown := protowire.AppendBytes(protowire.AppendTag(nil, 9999, protowire.BytesType), []byte("discard"))
	require.NoError(t, budget.Unmarshal(unknown, msg))
	require.Empty(t, msg.GetUnknown())
	require.NoError(t, budget.Validate(bytes.Repeat(unknown, 3), desc))
	require.ErrorIs(t, budget.Validate(bytes.Repeat(unknown, 4), desc), ErrMessageComplexity)
}

func TestProtoPackedFields(t *testing.T) {
	desc := protoTestFile(t).Messages().ByName("Repeated")
	for _, tc := range []struct {
		name  string
		value []byte
	}{
		{name: "integers", value: protowire.AppendVarint(nil, 1)},
		{name: "states", value: protowire.AppendVarint(nil, 1)},
		{name: "fixed32", value: protowire.AppendFixed32(nil, 1)},
		{name: "fixed64", value: protowire.AppendFixed64(nil, 1)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			field := desc.Fields().ByName(protoreflect.Name(tc.name))
			require.False(t, field.IsPacked())
			packed := func(count int) []byte {
				return protowire.AppendBytes(protowire.AppendTag(nil, field.Number(), protowire.BytesType), bytes.Repeat(tc.value, count))
			}
			budget := Proto{MaxFields: 3, MaxDepth: 3}
			msg := dynamicpb.NewMessage(desc)
			// The decoder accepts packed wire values even with packed=false.
			require.NoError(t, proto.Unmarshal(packed(3), msg))
			require.Equal(t, 3, msg.Get(field).List().Len())
			before := proto.Clone(msg)
			require.ErrorIs(t, budget.Unmarshal(packed(3), msg), ErrMessageComplexity)
			require.True(t, proto.Equal(before, msg))
			require.NoError(t, budget.Unmarshal(packed(2), msg))
			require.Equal(t, 2, msg.Get(field).List().Len())
			kind := protowire.VarintType
			switch field.Kind() {
			case protoreflect.Fixed32Kind:
				kind = protowire.Fixed32Type
			case protoreflect.Fixed64Kind:
				kind = protowire.Fixed64Type
			}
			unpacked := append(protowire.AppendTag(nil, field.Number(), kind), tc.value...)
			require.NoError(t, budget.Unmarshal(bytes.Repeat(unpacked, 3), msg))
			require.Equal(t, 3, msg.Get(field).List().Len())
		})
	}
}

func TestProtoMaps(t *testing.T) {
	input := &livekit.ParticipantInfo{Attributes: map[string]string{"a": "one", "b": "two"}}
	data, err := proto.Marshal(input)
	require.NoError(t, err)
	msg := &livekit.ParticipantInfo{Identity: "preserved"}
	before := proto.Clone(msg)
	require.ErrorIs(t, (Proto{MaxFields: 5, MaxDepth: 2}).Unmarshal(data, msg), ErrMessageComplexity)
	require.True(t, proto.Equal(before, msg))
	require.NoError(t, (Proto{MaxFields: 6, MaxDepth: 2}).Unmarshal(data, msg))
	require.True(t, proto.Equal(input, msg))
	require.ErrorIs(t, (Proto{MaxFields: 6, MaxDepth: 1}).Unmarshal(data, msg), ErrMessageComplexity)
}

func TestProtoExtensions(t *testing.T) {
	file := protoTestFile(t)
	desc := file.Messages().ByName("Container")
	node := file.Messages().ByName("Node")
	extension := file.Extensions().ByName("items")
	child := node.Fields().ByName("children")
	emptyChild := protowire.AppendBytes(protowire.AppendTag(nil, child.Number(), protowire.BytesType), nil)
	encode := func(count int) []byte {
		return protowire.AppendBytes(protowire.AppendTag(nil, extension.Number(), protowire.BytesType), bytes.Repeat(emptyChild, count))
	}
	msg := dynamicpb.NewMessage(desc)
	require.NoError(t, proto.Unmarshal(encode(4), msg))
	xt, err := protoregistry.GlobalTypes.FindExtensionByNumber(desc.FullName(), extension.Number())
	require.NoError(t, err)
	require.Equal(t, 4, msg.Get(xt.TypeDescriptor()).List().Get(0).Message().Get(child).List().Len())
	before := proto.Clone(msg)
	budget := Proto{MaxFields: 4, MaxDepth: 4}
	require.ErrorIs(t, budget.Unmarshal(encode(4), msg), ErrMessageComplexity)
	require.True(t, proto.Equal(before, msg))
	require.NoError(t, budget.Unmarshal(encode(1), msg))
	require.Equal(t, 1, msg.Get(xt.TypeDescriptor()).List().Get(0).Message().Get(child).List().Len())

	packed := file.Extensions().ByName("numbers")
	require.False(t, packed.IsPacked())
	data := protowire.AppendBytes(protowire.AppendTag(nil, packed.Number(), protowire.BytesType), bytes.Repeat([]byte{1}, 4))
	require.NoError(t, proto.Unmarshal(data, dynamicpb.NewMessage(desc)))
	require.ErrorIs(t, budget.Unmarshal(data, msg), ErrMessageComplexity)
	unknown := protowire.AppendBytes(protowire.AppendTag(nil, 150, protowire.BytesType), []byte("discard"))
	require.NoError(t, budget.Unmarshal(unknown, msg))
	require.Empty(t, msg.GetUnknown())
}

func TestProtoDepth(t *testing.T) {
	desc := protoTestFile(t).Messages().ByName("Node")
	budget := Proto{MaxFields: 10, MaxDepth: 3}
	var data []byte
	for depth := 1; depth <= 4; depth++ {
		msg := dynamicpb.NewMessage(desc)
		if depth <= budget.MaxDepth {
			require.NoError(t, budget.Validate(data, desc))
			require.NoError(t, budget.Unmarshal(data, msg))
		} else {
			require.ErrorIs(t, budget.Validate(data, desc), ErrMessageComplexity)
			require.ErrorIs(t, budget.Unmarshal(data, msg), ErrMessageComplexity)
		}
		data = protowire.AppendBytes(protowire.AppendTag(nil, 1, protowire.BytesType), data)
	}
}

func TestProtoInvalidInput(t *testing.T) {
	desc := protoTestFile(t).Messages().ByName("Node")
	budget := Proto{MaxFields: 10, MaxDepth: 3}
	for _, data := range [][]byte{{0xff}, {0x12, 5, 'x'}, {0x08, 0x80}} {
		require.Error(t, budget.Validate(data, desc))
	}
	for _, kind := range []protowire.Type{protowire.StartGroupType, protowire.EndGroupType} {
		require.ErrorIs(t, budget.Validate(protowire.AppendTag(nil, 999, kind), desc), ErrMessageComplexity)
	}
	group := protowire.AppendTag(nil, 999, protowire.StartGroupType)
	group = protowire.AppendTag(group, 999, protowire.EndGroupType)
	require.NoError(t, proto.Unmarshal(group, dynamicpb.NewMessage(desc)))
	require.ErrorIs(t, budget.Validate(group, desc), ErrMessageComplexity)
	repeated := protoTestFile(t).Messages().ByName("Repeated")
	for _, data := range [][]byte{{0x0a, 1, 0x80}, {0x1a, 3, 1, 2, 3}, {0x22, 7, 1, 2, 3, 4, 5, 6, 7}} {
		require.Error(t, budget.Validate(data, repeated))
	}
	require.Error(t, budget.Validate(nil, nil))
	require.Error(t, budget.Unmarshal(nil, nil))
	var nilMessage *livekit.SignalRequest
	require.Error(t, budget.Unmarshal(nil, nilMessage))
	var nilDynamic *dynamicpb.Message
	require.Error(t, budget.Unmarshal(nil, nilDynamic))
	for _, invalid := range []Proto{{}, {MaxFields: 1}, {MaxDepth: 1}, {MaxFields: -1, MaxDepth: 1}} {
		require.Error(t, invalid.Validate(nil, desc))
		require.Error(t, invalid.Unmarshal(nil, dynamicpb.NewMessage(desc)))
	}
	// The bounded binary decoder retains required-field validation.
	require.Error(t, budget.Unmarshal(nil, &descriptorpb.UninterpretedOption_NamePart{}))
}

func TestJSON(t *testing.T) {
	budget := JSON{MaxTokens: 20, MaxDepth: 3}
	for _, data := range []string{`null`, `true`, `false`, `1e100000`, `"value"`, `{}`, `[]`, `{"name":"camera","layers":[{}]}`} {
		t.Run(data, func(t *testing.T) { require.NoError(t, budget.Validate([]byte(data))) })
	}
	for _, data := range []string{"", " \n\t", `{} {}`, `true false`, `null 0`, `{"a":`, `[1,]`, `]`} {
		t.Run("invalid "+data, func(t *testing.T) { require.Error(t, budget.Validate([]byte(data))) })
	}
	require.NoError(t, (JSON{MaxTokens: 3, MaxDepth: 1}).Validate([]byte(`[0]`)))
	require.ErrorIs(t, (JSON{MaxTokens: 2, MaxDepth: 1}).Validate([]byte(`[0]`)), ErrMessageComplexity)
	require.NoError(t, (JSON{MaxTokens: 4, MaxDepth: 1}).Validate([]byte(`{"a":0}`)))
	require.ErrorIs(t, (JSON{MaxTokens: 3, MaxDepth: 1}).Validate([]byte(`{"a":0}`)), ErrMessageComplexity)
	require.NoError(t, budget.Validate([]byte(`[[[]]]`)))
	require.ErrorIs(t, budget.Validate([]byte(`[[[[]]]]`)), ErrMessageComplexity)
	require.ErrorIs(t, budget.Validate([]byte("["+strings.Repeat("{},", 20)+"{}]")), ErrMessageComplexity)
	for _, invalid := range []JSON{{}, {MaxTokens: 1}, {MaxDepth: 1}, {MaxTokens: -1, MaxDepth: 1}} {
		require.Error(t, invalid.Validate([]byte(`null`)))
	}
}

var protoTestDescriptors = sync.OnceValues(func() (protoreflect.FileDescriptor, error) {
	optional := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	repeated := descriptorpb.FieldDescriptorProto_LABEL_REPEATED
	message := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING
	int32Kind := descriptorpb.FieldDescriptorProto_TYPE_INT32
	enum := descriptorpb.FieldDescriptorProto_TYPE_ENUM
	fixed32 := descriptorpb.FieldDescriptorProto_TYPE_FIXED32
	fixed64 := descriptorpb.FieldDescriptorProto_TYPE_FIXED64
	field := func(name string, number int32, label descriptorpb.FieldDescriptorProto_Label, kind descriptorpb.FieldDescriptorProto_Type, typeName string) *descriptorpb.FieldDescriptorProto {
		f := &descriptorpb.FieldDescriptorProto{Name: proto.String(name), Number: proto.Int32(number), Label: &label, Type: &kind}
		if typeName != "" {
			f.TypeName = proto.String(typeName)
		}
		return f
	}
	packed := func(name string, number int32, kind descriptorpb.FieldDescriptorProto_Type, typeName string) *descriptorpb.FieldDescriptorProto {
		f := field(name, number, repeated, kind, typeName)
		f.Options = &descriptorpb.FieldOptions{Packed: proto.Bool(false)}
		return f
	}
	items := field("items", 100, repeated, message, ".decode_budget_test.Node")
	items.Extendee = proto.String(".decode_budget_test.Container")
	numbers := packed("numbers", 101, int32Kind, "")
	numbers.Extendee = items.Extendee
	file, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name: proto.String("decode_budget_test.proto"), Package: proto.String("decode_budget_test"), Syntax: proto.String("proto2"),
		EnumType: []*descriptorpb.EnumDescriptorProto{{Name: proto.String("State"), Value: []*descriptorpb.EnumValueDescriptorProto{
			{Name: proto.String("ZERO"), Number: proto.Int32(0)}, {Name: proto.String("ONE"), Number: proto.Int32(1)},
		}}},
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: proto.String("Node"), Field: []*descriptorpb.FieldDescriptorProto{
				field("child", 1, optional, message, ".decode_budget_test.Node"), field("name", 2, optional, str, ""),
				field("children", 3, repeated, message, ".decode_budget_test.Node"),
			}},
			{Name: proto.String("Repeated"), Field: []*descriptorpb.FieldDescriptorProto{
				packed("integers", 1, int32Kind, ""), packed("states", 2, enum, ".decode_budget_test.State"),
				packed("fixed32", 3, fixed32, ""), packed("fixed64", 4, fixed64, ""),
			}},
			{Name: proto.String("Container"), ExtensionRange: []*descriptorpb.DescriptorProto_ExtensionRange{{Start: proto.Int32(100), End: proto.Int32(200)}}},
		},
		Extension: []*descriptorpb.FieldDescriptorProto{items, numbers},
	}, nil)
	if err != nil {
		return nil, err
	}
	for i := 0; i < file.Extensions().Len(); i++ {
		if err = protoregistry.GlobalTypes.RegisterExtension(dynamicpb.NewExtensionType(file.Extensions().Get(i))); err != nil {
			return nil, err
		}
	}
	return file, nil
})

func protoTestFile(t *testing.T) protoreflect.FileDescriptor {
	t.Helper()
	file, err := protoTestDescriptors()
	require.NoError(t, err)
	return file
}

func TestProtoDecodedBytes(t *testing.T) {
	input := &livekit.ParticipantInfo{Identity: strings.Repeat("x", 1000), Attributes: map[string]string{"a": "one"}}
	data, err := proto.Marshal(input)
	require.NoError(t, err)
	desc := input.ProtoReflect().Descriptor()

	unlimited := Proto{MaxFields: 100, MaxDepth: 4}
	estimate, err := unlimited.Estimate(data, desc)
	require.NoError(t, err)
	require.Greater(t, estimate, 1000)

	require.NoError(t, (Proto{MaxFields: 100, MaxDepth: 4, MaxDecodedBytes: estimate}).Validate(data, desc))
	require.ErrorIs(t, (Proto{MaxFields: 100, MaxDepth: 4, MaxDecodedBytes: estimate - 1}).Validate(data, desc), ErrMessageComplexity)
	require.Error(t, (Proto{MaxFields: 100, MaxDepth: 4, MaxDecodedBytes: -1}).Validate(data, desc))

	jsonData, err := protojson.Marshal(input)
	require.NoError(t, err)
	jsonEstimate, err := unlimited.EstimateJSON(jsonData, desc)
	require.NoError(t, err)
	require.ErrorIs(t, (Proto{MaxFields: 100, MaxDepth: 4, MaxDecodedBytes: jsonEstimate - 1}).ValidateJSON(jsonData, desc), ErrMessageComplexity)
}

func TestProtoJSONStructure(t *testing.T) {
	desc := (&livekit.ParticipantInfo{}).ProtoReflect().Descriptor()
	budget := Proto{MaxFields: 4, MaxDepth: 2}
	require.NoError(t, budget.ValidateJSON([]byte(`{"identity":"a","tracks":[{"sid":"TR"}]}`), desc))
	require.ErrorIs(t, budget.ValidateJSON([]byte(`{"tracks":[{"sid":"a"},{"sid":"b"},{"sid":"c"}]}`), desc), ErrMessageComplexity)
	require.ErrorIs(t, budget.ValidateJSON([]byte(`{"tracks":[{"layers":[{}]}]}`), desc), ErrMessageComplexity)
	require.ErrorIs(t, (Proto{MaxFields: 4, MaxDepth: 2}).ValidateJSON([]byte(`{"unknown":[[[[1]]]]}`), desc), ErrMessageComplexity)
	require.Error(t, budget.ValidateJSON([]byte(`{} {}`), desc))
	require.Error(t, budget.ValidateJSON([]byte(`{"identity":`), desc))
}

func TestAppendAllocBytes(t *testing.T) {
	require.Equal(t, 0, appendAllocBytes(0, 8))
	// 8, 16, 32 and 64 byte arrays for caps 1, 2, 4 and 8.
	require.Equal(t, 8+16+32+64, appendAllocBytes(5, 8))
}
