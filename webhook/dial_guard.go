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
	"errors"
	"fmt"
	"net"
	"strings"
	"syscall"
)

// Marks a destination that the SSRF guard rejects. It threads through the transport's url.Error.
// Callers can then classify it as permanent.
var ErrBlockedAddress = errors.New("blocked destination address")

var blockedNets = func() []*net.IPNet {
	cidrs := []string{
		"100.64.0.0/10",   // RFC 6598 CGNAT (carrier-grade NAT)
		"0.0.0.0/8",       // "this network" (RFC 1122)
		"240.0.0.0/4",     // reserved for future use (RFC 1112)
		"198.18.0.0/15",   // benchmarking (RFC 2544)
		"192.0.0.0/24",    // IETF protocol assignments (RFC 6890)
		"192.88.99.0/24",  // 6to4 relay anycast (RFC 3068, deprecated)
		"192.0.2.0/24",    // TEST-NET-1 documentation (RFC 5737)
		"198.51.100.0/24", // TEST-NET-2 documentation (RFC 5737)
		"203.0.113.0/24",  // TEST-NET-3 documentation (RFC 5737)
		"64:ff9b::/96",    // NAT64 well-known prefix (RFC 6052)
		"64:ff9b:1::/48",  // NAT64 local-use prefix (RFC 8215)
		"::/96",           // IPv4-compatible IPv6 (RFC 4291, deprecated; embeds a v4 address)
		"2002::/16",       // 6to4 (RFC 3056; embeds a v4 address)
		"2001::/23",       // IETF protocol assignments (RFC 2928), includes Teredo 2001::/32 (RFC 4380; embeds a v4 address)
		"fec0::/10",       // IPv6 site-local (deprecated, RFC 3879)
		"100::/64",        // discard-only (RFC 6666)
		"2001:db8::/32",   // IPv6 documentation (RFC 3849)
		"3fff::/20",       // IPv6 documentation (RFC 9637)
		"5f00::/16",       // SRv6 SIDs (RFC 9602)
	}
	nets := make([]*net.IPNet, 0, len(cidrs))
	for _, c := range cidrs {
		_, n, _ := net.ParseCIDR(c)
		nets = append(nets, n)
	}
	return nets
}()

// CheckWebhookHost rejects a URL host that is localhost or a blocked IP
// literal, so the mistake surfaces when the webhook is saved. Hostnames are
// only resolved, and checked, when a delivery dials them.
func CheckWebhookHost(host string) error {
	host = strings.TrimSuffix(strings.ToLower(host), ".")
	if host == "localhost" || strings.HasSuffix(host, ".localhost") {
		return ErrBlockedAddress
	}
	if ip := net.ParseIP(host); ip != nil && isBlockedIP(ip) {
		return ErrBlockedAddress
	}
	return nil
}

func BlockInternal(network, address string, _ syscall.RawConn) error {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return err
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return fmt.Errorf("cannot resolve %q", address)
	}
	if isBlockedIP(ip) {
		return fmt.Errorf("%w: %s", ErrBlockedAddress, ip)
	}
	return nil
}

// Reports whether IP falls in a range the service must never dial:
// loopback, RFC1918 private and IPv6 ULA (fc00::/7), link-local unicast
// (including the 169.254.169.254 cloud-metadata IP and fe80::/10), link-local
// multicast, the unspecified address, any multicast, and the blockedNets ranges.
func isBlockedIP(ip net.IP) bool {
	if ip.IsLoopback() ||
		ip.IsPrivate() ||
		ip.IsLinkLocalUnicast() ||
		ip.IsLinkLocalMulticast() ||
		ip.IsUnspecified() ||
		ip.IsMulticast() {
		return true
	}
	for _, n := range blockedNets {
		if n.Contains(ip) {
			return true
		}
	}
	return false
}
