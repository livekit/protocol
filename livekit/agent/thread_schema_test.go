// Copyright 2026 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package agent

import "testing"

func TestTaskStatesRoundTripWithoutCollapsingOutcomes(t *testing.T) {
	states := map[string]AgentTask_Status{
		"submitted":      AgentTask_SUBMITTED,
		"working":        AgentTask_WORKING,
		"input-required": AgentTask_INPUT_REQUIRED,
		"completed":      AgentTask_COMPLETED,
		"failed":         AgentTask_FAILED,
		"canceled":       AgentTask_CANCELED,
		"rejected":       AgentTask_REJECTED,
	}
	for state, expected := range states {
		status, ok := TaskStatus(state)
		if !ok || status != expected {
			t.Errorf("state %q mapped to %v, %v", state, status, ok)
		}
		got := TaskStates(expected)
		if len(got) != 1 || got[0] != state {
			t.Errorf("filter %v expands to %v", expected, got)
		}
	}
	if _, ok := TaskStatus("unknown"); ok {
		t.Fatal("unknown state accepted")
	}
	for _, status := range []AgentTask_Status{AgentTask_UNSPECIFIED, AgentTask_Status(999)} {
		if got := TaskStates(status); len(got) != 0 {
			t.Fatalf("invalid filter %v expands to %v", status, got)
		}
	}
}
