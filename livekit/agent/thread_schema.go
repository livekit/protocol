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

import _ "embed"

// ThreadSchemaVersion is the conversation schema a thread's database holds.
// Version 2 stores A2A v1: task states by their v1 names, messages and tasks
// as v1 protobuf JSON, and tasks.updated_at in milliseconds.
const ThreadSchemaVersion = 2

// ThreadSchema creates the conversation schema, one statement per ";"
// (comments hold none).
//
//go:embed thread_schema.sql
var ThreadSchema string

// taskStatuses maps the values of tasks.status (A2A v1 task state names) to
// AgentTask.Status: a task not yet finished is running, a rejected one failed.
var taskStatuses = map[string]AgentTask_Status{
	"TASK_STATE_SUBMITTED":      AgentTask_RUNNING,
	"TASK_STATE_WORKING":        AgentTask_RUNNING,
	"TASK_STATE_INPUT_REQUIRED": AgentTask_RUNNING,
	"TASK_STATE_AUTH_REQUIRED":  AgentTask_RUNNING,
	"TASK_STATE_COMPLETED":      AgentTask_COMPLETED,
	"TASK_STATE_FAILED":         AgentTask_FAILED,
	"TASK_STATE_REJECTED":       AgentTask_FAILED,
	"TASK_STATE_CANCELED":       AgentTask_CANCELED,
}

// TaskStatus is the AgentTask.Status of a tasks.status value.
func TaskStatus(state string) (AgentTask_Status, bool) {
	status, ok := taskStatuses[state]
	return status, ok
}

// TaskStates are the tasks.status values of an AgentTask.Status.
func TaskStates(status AgentTask_Status) []string {
	var states []string
	for state, s := range taskStatuses {
		if s == status {
			states = append(states, state)
		}
	}
	return states
}

