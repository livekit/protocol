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

func TestTaskFiltersKeepNonterminalAndRejectedStatesSeparate(t *testing.T) {
	for _, state := range []string{"submitted", "working", "input-required", "rejected", "failed"} {
		got := TaskStates(state)
		if len(got) != 1 || got[0] != state {
			t.Errorf("filter %q expands to %v", state, got)
		}
	}
	if got := TaskStates("unknown"); len(got) != 0 {
		t.Fatalf("unknown state accepted: %v", got)
	}
}
