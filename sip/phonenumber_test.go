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

const (
	usNumber = "+14155550123"
	ukNumber = "+442071234567"
	auNumber = "+61212345678"
)

func TestIsEmergencyNumber(t *testing.T) {
	cases := []struct {
		number string
		from   string
		exp    bool
	}{
		{"911", usNumber, true},
		{"112", usNumber, true},
		{"+1911", usNumber, true},
		{"+1 911", usNumber, true},
		{"+1-911", usNumber, true},
		{"+1 (911)", usNumber, true},
		{"1911", usNumber, false},
		{"999", usNumber, false},
		{"933", usNumber, false},
		{"9111", usNumber, false},
		{"0911", usNumber, false},
		{"+14155559110", usNumber, false},
		{"+19115551212", usNumber, false},
		{"alice", usNumber, false},
		{"", usNumber, false},
		{"999", ukNumber, true},
		{"112", ukNumber, true},
		{"911", ukNumber, false},
		{"000", auNumber, true},
		{"+61000", usNumber, true},
		{"+44999", usNumber, true},
		{"+33112", usNumber, true},
		{"+1911", ukNumber, true},
		{"+1911", "", true},
		{"911", "", false},
	}
	for _, c := range cases {
		t.Run(c.number+"/"+c.from, func(t *testing.T) {
			require.Equal(t, c.exp, IsEmergencyNumber(c.number, c.from))
		})
	}
}
