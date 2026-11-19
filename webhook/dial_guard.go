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
	"net/netip"
	"slices"
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

func BlockInternal(_, address string, _ syscall.RawConn) error {
	return checkDialAddress(address, true, nil, nil)
}

// Refuses addresses in blocked, and private ones when blockPrivate is set,
// unless they are also in allowed.
func NewDialControl(blockPrivate bool, blocked, allowed []string) (func(network, address string, c syscall.RawConn) error, error) {
	blockedPrefixes, err := parseCIDRs("blocked_cidrs", blocked)
	if err != nil {
		return nil, err
	}
	allowedPrefixes, err := parseCIDRs("allowed_cidrs", allowed)
	if err != nil {
		return nil, err
	}
	return func(_, address string, _ syscall.RawConn) error {
		return checkDialAddress(address, blockPrivate, blockedPrefixes, allowedPrefixes)
	}, nil
}

func checkDialAddress(address string, blockPrivate bool, blocked, allowed []netip.Prefix) error {
	ap, err := netip.ParseAddrPort(address)
	if err != nil {
		return err
	}
	addr := ap.Addr().WithZone("").Unmap()
	if prefixesContain(allowed, addr) {
		return nil
	}
	// an allowed entry is an exception, so only refusal widens to the embedded IPv4 address
	if ap.Addr().Zone() != "" ||
		(blockPrivate && isBlockedIP(addr.AsSlice())) ||
		prefixesContain(blocked, addr) ||
		prefixesContain(blocked, embeddedIPv4(addr)) {
		return fmt.Errorf("%w: %s", ErrBlockedAddress, ap.Addr())
	}
	return nil
}

func parseCIDRs(key string, cidrs []string) ([]netip.Prefix, error) {
	prefixes := make([]netip.Prefix, 0, len(cidrs))
	for _, c := range cidrs {
		p, err := netip.ParsePrefix(c)
		if err != nil {
			return nil, fmt.Errorf("invalid webhook %s entry %q: %w", key, c, err)
		}
		// a dialed address is unmapped before matching, so a mapped prefix must be too
		if p.Addr().Is4In6() {
			if p.Bits() < 96 {
				return nil, fmt.Errorf("invalid webhook %s entry %q: an IPv4-mapped prefix needs at least 96 bits", key, c)
			}
			p = netip.PrefixFrom(p.Addr().Unmap(), p.Bits()-96)
		}
		prefixes = append(prefixes, p)
	}
	return prefixes, nil
}

func prefixesContain(prefixes []netip.Prefix, addr netip.Addr) bool {
	return slices.ContainsFunc(prefixes, func(p netip.Prefix) bool { return p.Contains(addr) })
}

var (
	nat64Prefix    = netip.MustParsePrefix("64:ff9b::/96")
	v4CompatPrefix = netip.MustParsePrefix("::/96")
	sixToFour      = netip.MustParsePrefix("2002::/16")
)

// Returns the IPv4 address a NAT64, v4-compatible or 6to4 address carries, or
// the zero Addr, which no prefix contains.
func embeddedIPv4(addr netip.Addr) netip.Addr {
	b := addr.As16()
	switch {
	case nat64Prefix.Contains(addr), v4CompatPrefix.Contains(addr):
		return netip.AddrFrom4([4]byte(b[12:]))
	case sixToFour.Contains(addr):
		return netip.AddrFrom4([4]byte(b[2:6]))
	default:
		return netip.Addr{}
	}
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
