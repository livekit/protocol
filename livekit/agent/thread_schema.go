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

// TaskStatus validates a framework A2A state without collapsing its outcome.
func TaskStatus(state string) (string, bool) {
	switch state {
	case "submitted", "working", "input-required", "completed", "failed", "canceled", "rejected":
		return state, true
	default:
		return "", false
	}
}

// TaskStates is the exact state selected by a task filter.
func TaskStates(state string) []string {
	if state, ok := TaskStatus(state); ok {
		return []string{state}
	}
	return nil
}
