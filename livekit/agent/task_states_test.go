package agent

import (
	"fmt"
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
