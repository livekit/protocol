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

// Every AgentTask.Status must be reachable from a stored state, or a filter on
// it would silently match nothing.
func TestEveryTaskStatusHasStates(t *testing.T) {
	for value := range AgentTask_Status_name {
		if len(TaskStates(AgentTask_Status(value))) == 0 {
			t.Errorf("%s has no tasks.status value", AgentTask_Status(value))
		}
	}
}
