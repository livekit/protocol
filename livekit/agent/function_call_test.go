package agent

import (
	"reflect"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestFunctionCallMetadataSurvivesSerialization(t *testing.T) {
	original := &FunctionCall{
		CallId: "call-1",
		Extra: map[string]string{
			"lk.approval":          `"granted"`,
			"lk.approval_identity": `"{\"amount\":42}"`,
			"lk.delegation.state":  `{"pending_task_id":"task-1"}`,
			"application_receipt":  `"confirmed-appointment"`,
		},
	}
	for format, codec := range map[string]struct {
		marshal   func(proto.Message) ([]byte, error)
		unmarshal func([]byte, proto.Message) error
	}{
		"json":     {protojson.Marshal, protojson.Unmarshal},
		"protobuf": {proto.Marshal, proto.Unmarshal},
	} {
		t.Run(format, func(t *testing.T) {
			data, err := codec.marshal(original)
			if err != nil {
				t.Fatal(err)
			}
			var restored FunctionCall
			if err := codec.unmarshal(data, &restored); err != nil {
				t.Fatal(err)
			}
			if restored.CallId != original.CallId || !reflect.DeepEqual(restored.Extra, original.Extra) {
				t.Fatalf("serialization lost call metadata: %v", &restored)
			}
		})
	}
}

func TestFunctionCallMetadataDoesNotReuseLegacyFieldSix(t *testing.T) {
	var call FunctionCall
	if number := call.ProtoReflect().Descriptor().Fields().ByName("extra").Number(); number != 7 {
		t.Fatalf("extra uses field %d; field 6 belongs to incompatible legacy data", number)
	}
	// A legacy length-delimited string in field 6 must remain an unknown field,
	// rather than being decoded as a metadata-map entry.
	legacy := []byte{0x32, 0x04, 't', 'u', 'r', 'n'}
	if err := proto.Unmarshal(legacy, &call); err != nil {
		t.Fatal(err)
	}
	if len(call.Extra) != 0 {
		t.Fatalf("legacy field 6 populated metadata: %v", call.Extra)
	}
}
