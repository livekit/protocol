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
	"net/netip"
	"strings"
	"syscall"
)

// ErrWebhookTargetBlocked rejects a webhook URL that points at an address this
// server must not reach on a caller's behalf.
var ErrWebhookTargetBlocked = errors.New("webhook url must resolve to a public address")

// blockedWebhookPrefixes are reserved ranges netip's predicates do not cover.
var blockedWebhookPrefixes = []netip.Prefix{
	netip.MustParsePrefix("0.0.0.0/8"),     // "this network"; 0.0.0.0 reaches localhost
	netip.MustParsePrefix("100.64.0.0/10"), // carrier-grade NAT
	netip.MustParsePrefix("192.0.0.0/24"),  // IETF protocol assignments
	netip.MustParsePrefix("198.18.0.0/15"), // benchmarking
	netip.MustParsePrefix("240.0.0.0/4"),   // reserved
}

// checkWebhookAddr rejects loopback, private, link-local (including cloud
// metadata at 169.254.169.254), multicast, unspecified and other reserved
// addresses, in IPv4, IPv6 and IPv4-mapped IPv6 form.
func checkWebhookAddr(addr netip.Addr) error {
	addr = addr.Unmap()
	if !addr.IsValid() || !addr.IsGlobalUnicast() || addr.IsPrivate() {
		return ErrWebhookTargetBlocked
	}
	for _, p := range blockedWebhookPrefixes {
		if p.Contains(addr) {
			return ErrWebhookTargetBlocked
		}
	}
	return nil
}

// CheckWebhookHost rejects a URL host that is localhost or a blocked IP
// literal, so the mistake surfaces when the webhook is saved. Hostnames are
// only resolved, and checked, when a delivery dials them.
func CheckWebhookHost(host string) error {
	host = strings.TrimSuffix(strings.ToLower(host), ".")
	if host == "localhost" || strings.HasSuffix(host, ".localhost") {
		return ErrWebhookTargetBlocked
	}
	if addr, err := netip.ParseAddr(host); err == nil {
		return checkWebhookAddr(addr)
	}
	return nil
}

// GuardWebhookDial runs after DNS resolution, on the address actually
// dialed, so a hostname resolving (or rebinding) to a blocked address is
// refused before a connection opens.
func GuardWebhookDial(_, address string, _ syscall.RawConn) error {
	ap, err := netip.ParseAddrPort(address)
	if err != nil {
		return ErrWebhookTargetBlocked
	}
	return checkWebhookAddr(ap.Addr())
}
