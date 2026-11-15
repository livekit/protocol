package agent

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestTaskNamePresenceAndStatusSurviveSerialization(t *testing.T) {
	for _, name := range []*string{nil, new(string)} {
		original := &AgentTask_Task{Status: AgentTask_INPUT_REQUIRED, Name: name}
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
				var restored AgentTask_Task
				if err := codec.unmarshal(data, &restored); err != nil {
					t.Fatal(err)
				}
				if restored.Status != original.Status || (restored.Name == nil) != (original.Name == nil) {
					t.Fatalf("serialization lost state or label presence: %v", &restored)
				}
			})
		}
	}
}
