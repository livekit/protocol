package agent

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

func TestTaskStatesKeepBusinessOutcomes(t *testing.T) {
	for _, state := range []string{"submitted", "working", "input-required", "completed", "failed", "canceled", "rejected"} {
		actual, ok := TaskStatus(state)
		if !ok || fmt.Sprint(actual) != state {
			t.Errorf("task state %q projected as %q", state, fmt.Sprint(actual))
		}
	}
}

func TestTaskNamePresenceAndStatusSurviveJSON(t *testing.T) {
	for _, name := range []*string{nil, new(string)} {
		original := &AgentTask_Task{Status: "input-required", Name: name}
		data, err := protojson.Marshal(original)
		if err != nil {
			t.Fatal(err)
		}
		var restored AgentTask_Task
		if err := protojson.Unmarshal(data, &restored); err != nil {
			t.Fatal(err)
		}
		if restored.Status != original.Status || (restored.Name == nil) != (original.Name == nil) {
			t.Fatalf("JSON lost state or label presence: %s", data)
		}
	}
}
