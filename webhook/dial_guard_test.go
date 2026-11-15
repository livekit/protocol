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

package webhook

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBlockedIP(t *testing.T) {
	cases := []struct {
		ip      string
		blocked bool
	}{
		// internal / reserved ranges the SSRF guard must reject
		{"127.0.0.1", true},        // IPv4 loopback
		{"10.0.0.1", true},         // RFC1918
		{"172.16.0.1", true},       // RFC1918
		{"192.168.1.1", true},      // RFC1918
		{"169.254.169.254", true},  // link-local including cloud metadata
		{"100.64.0.1", true},       // RFC6598 CGNAT
		{"::1", true},              // IPv6 loopback
		{"fc00::1", true},          // IPv6 ULA
		{"fe80::1", true},          // IPv6 link-local
		{"64:ff9b::1.2.3.4", true}, // well-known NAT64
		{"64:ff9b:1::1", true},     // local-use NAT64
		{"0.1.2.3", true},          // 0.0.0.0/8 "this network"
		{"240.0.0.1", true},        // 240.0.0.0/4 reserved
		{"198.19.0.1", true},       // 198.18.0.0/15 benchmarking
		{"192.0.0.1", true},        // 192.0.0.0/24 IETF protocol assignments
		{"192.88.99.1", true},      // 192.88.99.0/24 6to4 relay anycast
		{"::c0a8:101", true},       // ::/96 v4-compatible (embeds 192.168.1.1)
		{"2002:c0a8:101::1", true}, // 6to4 (embeds 192.168.1.1)
		{"2001::1", true},          // 2001::/32 Teredo (embeds IPv4)
		{"2001:1ff::1", true},      // 2001::/23 IETF protocol assignments, outside Teredo
		{"fec0::1", true},          // fec0::/10 deprecated IPv6 site-local
		{"192.0.2.1", true},        // 192.0.2.0/24 TEST-NET-1
		{"198.51.100.1", true},     // 198.51.100.0/24 TEST-NET-2
		{"203.0.113.1", true},      // 203.0.113.0/24 TEST-NET-3
		{"100::1", true},           // 100::/64 discard-only
		{"2001:db8::1", true},      // 2001:db8::/32 documentation
		{"3fff::1", true},          // 3fff::/20 documentation
		{"5f00::1", true},          // 5f00::/16 SRv6 SIDs
		{"2001:200::1", false},     // just past 2001::/23 (WIDE, public)
		// public addresses that must remain reachable
		{"8.8.8.8", false},
		{"1.1.1.1", false},
		{"2606:4700:4700::1111", false},
	}
	for _, c := range cases {
		ip := net.ParseIP(c.ip)
		require.NotNil(t, ip, "parse %s", c.ip)
		require.Equal(t, c.blocked, isBlockedIP(ip), "isBlockedIP(%s)", c.ip)
	}
}

func TestCheckWebhookHost(t *testing.T) {
	for _, host := range []string{"example.com", "93.184.215.14", "2606:4700::1111"} {
		require.NoError(t, CheckWebhookHost(host), host)
	}
	for _, host := range []string{
		"localhost", "LOCALHOST.", "api.localhost",
		"127.0.0.1", "10.0.0.1", "169.254.169.254",
		"::1", "::ffff:10.0.0.1", "0.0.0.0",
	} {
		require.ErrorIs(t, CheckWebhookHost(host), ErrBlockedAddress, host)
	}
}
