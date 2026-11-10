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

// ThreadSchemaVersion is the newest conversation schema a thread's database
// can hold.
const ThreadSchemaVersion = 1

// ThreadSchemaV1 creates version 1 of the conversation schema, one statement
// per ";" (comments hold none).
//
//go:embed thread_schema_v1.sql
var ThreadSchemaV1 string

// taskStatuses maps the values of tasks.status (A2A task states) to
// AgentTask.Status: a task not yet finished is running, a rejected one failed.
var taskStatuses = map[string]AgentTask_Status{
	"submitted":      AgentTask_RUNNING,
	"working":        AgentTask_RUNNING,
	"input-required": AgentTask_RUNNING,
	"completed":      AgentTask_COMPLETED,
	"failed":         AgentTask_FAILED,
	"rejected":       AgentTask_FAILED,
	"canceled":       AgentTask_CANCELED,
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

