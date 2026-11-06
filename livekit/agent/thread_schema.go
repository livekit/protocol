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

// TaskStatus converts a stored framework state to its typed API status.
func TaskStatus(state string) (AgentTask_Status, bool) {
	switch state {
	case "submitted":
		return AgentTask_SUBMITTED, true
	case "working":
		return AgentTask_WORKING, true
	case "input-required":
		return AgentTask_INPUT_REQUIRED, true
	case "completed":
		return AgentTask_COMPLETED, true
	case "failed":
		return AgentTask_FAILED, true
	case "canceled":
		return AgentTask_CANCELED, true
	case "rejected":
		return AgentTask_REJECTED, true
	default:
		return AgentTask_UNSPECIFIED, false
	}
}

// TaskStates returns the exact stored framework state selected by a filter.
func TaskStates(status AgentTask_Status) []string {
	switch status {
	case AgentTask_SUBMITTED:
		return []string{"submitted"}
	case AgentTask_WORKING:
		return []string{"working"}
	case AgentTask_INPUT_REQUIRED:
		return []string{"input-required"}
	case AgentTask_COMPLETED:
		return []string{"completed"}
	case AgentTask_FAILED:
		return []string{"failed"}
	case AgentTask_CANCELED:
		return []string{"canceled"}
	case AgentTask_REJECTED:
		return []string{"rejected"}
	default:
		return nil
	}
}
