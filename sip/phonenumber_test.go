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
		{"1911", true},
		{"+1911", true},
		{"+1 911", true},
		{"+1-911", true},
		{"(911)", true},
		{"", false},
		{"9", false},
		{"9111", false},
		{"0911", false},
		{"+14155559110", false},
		{"+19115551212", false},
		{"112", false},
		{"alice", false},
		{"911abc", false},
	}
	for _, c := range cases {
		t.Run(c.number, func(t *testing.T) {
			require.Equal(t, c.exp, IsEmergencyNumber(c.number))
		})
	}
}

func TestURIUser(t *testing.T) {
	cases := []struct {
		uri string
		exp string
	}{
		{"<sip:911@carrier.example.com>", "911"},
		{"sip:911@carrier.example.com", "911"},
		{"sips:+1911@carrier.example.com:5061;transport=tls", "+1911"},
		{"<sip:alice@example.com;transport=tcp>", "alice"},
		{"\"Bob\" <sip:bob@example.com>", "bob"},
		{"tel:911", "911"},
		{"tel:+1-911;phone-context=example.com", "+1-911"},
		{"<tel:911>", "911"},
		{"sip:example.com", ""},
		{"", ""},
		{"911", ""},
	}
	for _, c := range cases {
		t.Run(c.uri, func(t *testing.T) {
			require.Equal(t, c.exp, URIUser(c.uri))
		})
	}
}
