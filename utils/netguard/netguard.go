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

// Package netguard refuses connections to internal and reserved addresses, so
// a destination a caller does not trust cannot reach the host's own network.
//
// A dial guard sees the address being connected to, which is the proxy's when
// the client uses one, so a guarded http.Transport should set Proxy to nil.
package netguard

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"slices"
	"strings"
	"syscall"
)

// ErrBlockedAddress marks a destination the guard refuses. It threads through a
// transport's url.Error, so callers can classify it as permanent.
var ErrBlockedAddress = errors.New("blocked destination address")

var errShortMappedPrefix = errors.New("an IPv4-mapped prefix needs at least 96 bits")

var blockedPrefixes = func() []netip.Prefix {
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
		"64:ff9b:1::/48",  // NAT64 local-use prefix (RFC 8215)
		"::/96",           // IPv4-compatible IPv6 (RFC 4291, deprecated; embeds a v4 address)
		"::ffff:0:0:0/96", // IPv4-translated (RFC 2765; embeds a v4 address)
		"2002::/16",       // 6to4 (RFC 3056; embeds a v4 address)
		"2001::/23",       // IETF protocol assignments (RFC 2928), includes Teredo 2001::/32 (RFC 4380; embeds a v4 address)
		"fec0::/10",       // IPv6 site-local (deprecated, RFC 3879)
		"100::/64",        // discard-only (RFC 6666)
		"2001:db8::/32",   // IPv6 documentation (RFC 3849)
		"3fff::/20",       // IPv6 documentation (RFC 9637)
		"5f00::/16",       // SRv6 SIDs (RFC 9602)
	}
	prefixes := make([]netip.Prefix, 0, len(cidrs))
	for _, c := range cidrs {
		prefixes = append(prefixes, netip.MustParsePrefix(c))
	}
	return prefixes
}()

// CheckHost rejects a URL host that is localhost, a blocked IP literal, or an
// IP literal with a zone. A port, and brackets around an IPv6 literal, are
// accepted. It only catches configuration mistakes early: host names, and IPv4
// spellings netip.ParseAddr rejects such as 127.1, pass here and are checked by
// a dial guard once resolved.
func CheckHost(host string) error {
	if h, _, err := net.SplitHostPort(host); err == nil {
		host = h
	}
	host = strings.TrimSuffix(strings.ToLower(host), ".")
	if host == "localhost" || strings.HasSuffix(host, ".localhost") {
		return ErrBlockedAddress
	}
	host = strings.TrimSuffix(strings.TrimPrefix(host, "["), "]")
	if addr, err := netip.ParseAddr(host); err == nil {
		if addr.Zone() != "" || IsBlockedIP(addr.AsSlice()) {
			return ErrBlockedAddress
		}
	}
	return nil
}

// BlockInternal is a net.Dialer Control function that refuses every address
// IsBlockedIP reports, and any zoned address.
func BlockInternal(_, address string, _ syscall.RawConn) error {
	return checkDialAddress(address, true, nil, nil)
}

// ReportOnly wraps a net.Dialer Control function so that an address it refuses
// with ErrBlockedAddress is passed to report and the connection goes ahead.
// Other errors from control are returned unchanged.
func ReportOnly(
	control func(network, address string, c syscall.RawConn) error,
	report func(address string, err error),
) func(network, address string, c syscall.RawConn) error {
	return func(network, address string, c syscall.RawConn) error {
		err := control(network, address, c)
		if errors.Is(err, ErrBlockedAddress) {
			report(address, err)
			return nil
		}
		return err
	}
}

// NewDialControl returns a net.Dialer Control function that refuses addresses
// in blocked, and private ones when blockPrivate is set, unless they are also
// in allowed. An allowed address is never refused, so a broad allowed range
// such as 0.0.0.0/0 also lets loopback and the cloud metadata address through.
// Prefixes should come from ParseCIDR.
func NewDialControl(blockPrivate bool, blocked, allowed []netip.Prefix) func(network, address string, c syscall.RawConn) error {
	return func(_, address string, _ syscall.RawConn) error {
		return checkDialAddress(address, blockPrivate, blocked, allowed)
	}
}

// ParseCIDR parses a range for NewDialControl. A dialed address is unmapped
// before matching, so an IPv4-mapped range is converted to its IPv4 range.
func ParseCIDR(cidr string) (netip.Prefix, error) {
	p, err := netip.ParsePrefix(cidr)
	if err != nil {
		return netip.Prefix{}, err
	}
	if p.Addr().Is4In6() {
		if p.Bits() < 96 {
			return netip.Prefix{}, errShortMappedPrefix
		}
		p = netip.PrefixFrom(p.Addr().Unmap(), p.Bits()-96)
	}
	return p, nil
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
		(blockPrivate && isBlocked(addr)) ||
		prefixesContain(blocked, addr) ||
		prefixesContain(blocked, embeddedIPv4(addr)) {
		return fmt.Errorf("%w: %s", ErrBlockedAddress, ap.Addr())
	}
	return nil
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

// IsBlockedIP reports whether ip falls in a range a guarded client must never
// dial: loopback, RFC1918 private and IPv6 ULA (fc00::/7), link-local unicast
// (including the 169.254.169.254 cloud-metadata IP and fe80::/10), link-local
// multicast, the unspecified address, any multicast, and the reserved ranges
// listed in blockedPrefixes. An address in the NAT64 well-known prefix is
// judged by the IPv4 address it carries. A nil or malformed ip is reported as
// blocked.
func IsBlockedIP(ip net.IP) bool {
	addr, ok := netip.AddrFromSlice(ip)
	if !ok {
		return true
	}
	return isBlocked(addr.Unmap())
}

// addr must be unmapped.
func isBlocked(addr netip.Addr) bool {
	if nat64Prefix.Contains(addr) {
		return isBlocked(embeddedIPv4(addr))
	}
	return addr.IsLoopback() ||
		addr.IsPrivate() ||
		addr.IsLinkLocalUnicast() ||
		addr.IsLinkLocalMulticast() ||
		addr.IsUnspecified() ||
		addr.IsMulticast() ||
		prefixesContain(blockedPrefixes, addr)
}
