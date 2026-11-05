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

package sip

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmergencyNumber(t *testing.T) {
	cases := []struct {
		number string
		exp    bool
	}{
		{"911", true},
		{"112", true},
		{"999", true},
		{"000", true},
		{"9-1-1", true},
		{"(911)", true},
		{"+1911", false},
		{"+1 911", false},
		{"1911", false},
		{"411", false},
		{"9111", false},
		{"0911", false},
		{"+14155559110", false},
		{"+19115551212", false},
		{"alice", false},
		{"", false},
	}
	for _, c := range cases {
		t.Run(c.number, func(t *testing.T) {
			require.Equal(t, c.exp, IsEmergencyNumber(c.number))
		})
	}
}
