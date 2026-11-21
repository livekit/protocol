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

package netguard

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"net/netip"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBlockedIP(t *testing.T) {
	cases := []struct {
		ip      string
		blocked bool
	}{
		// internal / reserved ranges the SSRF guard must reject
		{"127.0.0.1", true},                // IPv4 loopback
		{"10.0.0.1", true},                 // RFC1918
		{"172.16.0.1", true},               // RFC1918
		{"192.168.1.1", true},              // RFC1918
		{"169.254.169.254", true},          // link-local including cloud metadata
		{"100.64.0.1", true},               // RFC6598 CGNAT
		{"::1", true},                      // IPv6 loopback
		{"fc00::1", true},                  // IPv6 ULA
		{"fe80::1", true},                  // IPv6 link-local
		{"64:ff9b::10.0.0.1", true},        // well-known NAT64 carrying RFC1918
		{"64:ff9b::169.254.169.254", true}, // well-known NAT64 carrying cloud metadata
		{"64:ff9b::127.0.0.1", true},       // well-known NAT64 carrying loopback
		{"64:ff9b::8.8.8.8", false},        // well-known NAT64 carrying a public address
		{"64:ff9b:1::1", true},             // local-use NAT64
		{"64:ff9b:1::808:808", true},       // local-use NAT64, whatever it carries
		{"0.1.2.3", true},                  // 0.0.0.0/8 "this network"
		{"240.0.0.1", true},                // 240.0.0.0/4 reserved
		{"198.19.0.1", true},               // 198.18.0.0/15 benchmarking
		{"192.0.0.1", true},                // 192.0.0.0/24 IETF protocol assignments
		{"192.88.99.1", true},              // 192.88.99.0/24 6to4 relay anycast
		{"::c0a8:101", true},               // ::/96 v4-compatible (embeds 192.168.1.1)
		{"2002:c0a8:101::1", true},         // 6to4 (embeds 192.168.1.1)
		{"2001::1", true},                  // 2001::/32 Teredo (embeds IPv4)
		{"2001:1ff::1", true},              // 2001::/23 IETF protocol assignments, outside Teredo
		{"fec0::1", true},                  // fec0::/10 deprecated IPv6 site-local
		{"192.0.2.1", true},                // 192.0.2.0/24 TEST-NET-1
		{"198.51.100.1", true},             // 198.51.100.0/24 TEST-NET-2
		{"203.0.113.1", true},              // 203.0.113.0/24 TEST-NET-3
		{"100::1", true},                   // 100::/64 discard-only
		{"2001:db8::1", true},              // 2001:db8::/32 documentation
		{"3fff::1", true},                  // 3fff::/20 documentation
		{"5f00::1", true},                  // 5f00::/16 SRv6 SIDs
		{"2001:200::1", false},             // just past 2001::/23 (WIDE, public)
		{"::ffff:169.254.169.254", true},   // IPv4-mapped cloud metadata
		{"::ffff:0:a9fe:a9fe", true},       // ::ffff:0:0:0/96 IPv4-translated (embeds 169.254.169.254)
		{"::ffff:8.8.8.8", false},          // IPv4-mapped public address
		// public addresses that must remain reachable
		{"8.8.8.8", false},
		{"1.1.1.1", false},
		{"2606:4700:4700::1111", false},
	}
	for _, c := range cases {
		ip := net.ParseIP(c.ip)
		require.NotNil(t, ip, "parse %s", c.ip)
		require.Equal(t, c.blocked, IsBlockedIP(ip), "IsBlockedIP(%s)", c.ip)
	}

	for _, ip := range []net.IP{nil, {}, {127, 0, 0}, make(net.IP, 5)} {
		require.True(t, IsBlockedIP(ip), "IsBlockedIP(%#v)", ip)
	}
	require.False(t, IsBlockedIP(net.IPv4(8, 8, 8, 8).To4()), "4-byte form of a public address")
}

func TestCheckHost(t *testing.T) {
	for _, host := range []string{"example.com", "93.184.215.14", "2606:4700::1111", "[2606:4700::1111]",
		"example.com:443", "[2606:4700::1111]:443",
	} {
		require.NoError(t, CheckHost(host), host)
	}
	for _, host := range []string{
		"localhost", "LOCALHOST.", "api.localhost",
		"127.0.0.1", "10.0.0.1", "169.254.169.254",
		"::1", "::ffff:10.0.0.1", "0.0.0.0",
		"[::1]", "[::ffff:10.0.0.1]",
		"fe80::1%eth0", "2606:4700::1111%eth0", "[2606:4700::1111%eth0]",
		"127.0.0.1:80", "[::1]:80", "localhost:8080", "[fe80::1%eth0]:443",
	} {
		require.ErrorIs(t, CheckHost(host), ErrBlockedAddress, host)
	}
}

func TestNewDialControl(t *testing.T) {
	cases := []struct {
		blockPrivate     bool
		blocked, allowed []string
		want             map[string]bool
	}{
		{
			blockPrivate: true,
			blocked:      []string{"8.8.8.0/24"},
			allowed:      []string{"10.1.0.0/16"},
			want: map[string]bool{
				"8.8.8.8:443":             true,
				"[::ffff:8.8.8.8]:443":    true,
				"1.1.1.1:443":             false,
				"10.1.2.3:443":            false,
				"[::ffff:10.1.2.3]:80":    false,
				"10.2.0.1:443":            true,
				"[64:ff9b::a01:203]:80":   true,
				"[64:ff9b::101:101]:443":  false,
				"[::10.1.2.3]:80":         true,
				"[2002:a01:203::]:80":     true,
				"[64:ff9b::a02:1]:80":     true,
				"[fe80::1%eth0]:443":      true,
				"[2607:f8b0::1%eth0]:443": true,
			},
		},
		{
			blocked: []string{"10.0.0.0/8"},
			want: map[string]bool{
				"10.0.0.1:80":               true,
				"[64:ff9b::a00:1]:80":       true,
				"[::10.0.0.1]:80":           true,
				"[2002:a00:1::]:80":         true,
				"[64:ff9b::808:808]:80":     false,
				"8.8.8.8:80":                false,
				"[fe80::1%eth0]:443":        true,
				"[2607:f8b0::1%eth0]:443":   true,
				"[::ffff:8.8.8.8%eth0]:443": true,
			},
		},
		{
			blockPrivate: true,
			allowed:      []string{"0.0.0.0/0"},
			want: map[string]bool{
				"10.0.0.1:80":        false,
				"127.0.0.1:80":       false,
				"169.254.169.254:80": false,
				"[::1]:80":           true,
				"[::]:80":            true,
			},
		},
		{
			blockPrivate: true,
			allowed:      []string{"fe80::/10"},
			want: map[string]bool{
				"[fe80::1%eth0]:80":      false,
				"[2607:f8b0::1%eth0]:80": true,
			},
		},
		{
			blocked: []string{"::ffff:8.8.8.0/120"},
			want: map[string]bool{
				"8.8.8.8:443":          true,
				"[::ffff:8.8.8.8]:443": true,
				"8.8.9.1:443":          false,
			},
		},
		{
			blockPrivate: true,
			allowed:      []string{"::ffff:10.1.0.0/112"},
			want: map[string]bool{
				"10.1.2.3:443":          false,
				"[::ffff:10.1.2.3]:443": false,
				"10.2.0.1:443":          true,
			},
		},
	}
	for _, c := range cases {
		control := NewDialControl(c.blockPrivate, mustParseCIDRs(t, c.blocked), mustParseCIDRs(t, c.allowed))
		for address, blocked := range c.want {
			checkDial(t, control, address, blocked)
		}
	}
}

func TestParseCIDR(t *testing.T) {
	p, err := ParseCIDR("::ffff:10.1.0.0/112")
	require.NoError(t, err)
	require.Equal(t, netip.MustParsePrefix("10.1.0.0/16"), p)

	for _, cidr := range []string{"10.0.0.0/33", "not-a-cidr", "::ffff:0:0/80"} {
		_, err := ParseCIDR(cidr)
		require.Error(t, err, cidr)
	}
}

func TestReportOnly(t *testing.T) {
	type reported struct {
		address string
		err     error
	}
	var got []reported
	control := ReportOnly(BlockInternal, func(address string, err error) {
		got = append(got, reported{address, err})
	})

	require.NoError(t, control("tcp", "169.254.169.254:80", nil))
	require.Len(t, got, 1)
	require.Equal(t, "169.254.169.254:80", got[0].address)
	require.ErrorIs(t, got[0].err, ErrBlockedAddress)

	require.NoError(t, control("tcp", "8.8.8.8:443", nil))
	require.Error(t, control("tcp", "not-an-address", nil))
	require.Len(t, got, 1, "only refused addresses are reported")
}

func TestReportOnlyContext(t *testing.T) {
	type destKey struct{}
	type reported struct {
		dest, address string
		err           error
	}
	var got []reported
	control := ReportOnlyContext(BlockInternal, func(ctx context.Context, address string, err error) {
		dest, _ := ctx.Value(destKey{}).(string)
		got = append(got, reported{dest, address, err})
	})

	ctx := context.WithValue(context.Background(), destKey{}, "a")
	require.NoError(t, control(ctx, "tcp", "169.254.169.254:80", nil))
	require.NoError(t, control(ctx, "tcp", "8.8.8.8:443", nil))
	require.Error(t, control(ctx, "tcp", "not-an-address", nil))
	require.Len(t, got, 1, "only refused addresses are reported")
	require.Equal(t, "a", got[0].dest)
	require.Equal(t, "169.254.169.254:80", got[0].address)
	require.ErrorIs(t, got[0].err, ErrBlockedAddress)

	// http.Transport passes the request's context values to the dial.
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer srv.Close()
	dialer := &net.Dialer{ControlContext: control}
	client := &http.Client{Transport: &http.Transport{DialContext: dialer.DialContext}}
	defer client.CloseIdleConnections()

	got = nil
	req, err := http.NewRequestWithContext(context.WithValue(ctx, destKey{}, "b"), http.MethodGet, srv.URL, nil)
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err, "a report-only refusal must not fail the request")
	require.NoError(t, resp.Body.Close())
	require.Len(t, got, 1)
	require.Equal(t, "b", got[0].dest)
	require.Equal(t, srv.Listener.Addr().String(), got[0].address)
}

func TestBlockInternal(t *testing.T) {
	checkDial(t, BlockInternal, "[fe80::1%eth0]:443", true)
	checkDial(t, BlockInternal, "[2607:f8b0::1%eth0]:443", true)
	checkDial(t, BlockInternal, "127.0.0.1:80", true)
	checkDial(t, BlockInternal, "169.254.169.254:80", true)
	checkDial(t, BlockInternal, "8.8.8.8:443", false)
	checkDial(t, BlockInternal, "[64:ff9b::a9fe:a9fe]:80", true)
	checkDial(t, BlockInternal, "[64:ff9b::808:808]:443", false)
}

func mustParseCIDRs(t *testing.T, cidrs []string) []netip.Prefix {
	t.Helper()
	prefixes := make([]netip.Prefix, 0, len(cidrs))
	for _, c := range cidrs {
		p, err := ParseCIDR(c)
		require.NoError(t, err, c)
		prefixes = append(prefixes, p)
	}
	return prefixes
}

func checkDial(t *testing.T, control func(string, string, syscall.RawConn) error, address string, blocked bool) {
	t.Helper()
	err := control("tcp", address, nil)
	if blocked {
		require.ErrorIs(t, err, ErrBlockedAddress, address)
	} else {
		require.NoError(t, err, address)
	}
}
