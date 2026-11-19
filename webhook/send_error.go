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
	"regexp"
	"strings"
)

// One text for a refused and an unresolvable destination, so the webhook's
// owner cannot tell an internal name that resolves from one that does not.
const noAllowedAddressText = "destination host has no allowed address"

// Texts of a refused or unresolvable destination. net exports neither of its own.
var noAllowedAddressCauses = []string{
	ErrBlockedAddress.Error(),
	"no such host",
	"no suitable address found",
}

var (
	ipv4Token = regexp.MustCompile(`\d{1,3}(?:\.\d{1,3}){3}`)
	ipv6Token = regexp.MustCompile(`(?i)[0-9a-f]*:[0-9a-f:.]*:[0-9a-f.]*(?:%[0-9a-z]+)?`)
)

// Rewrites a delivery's SendError for a host whose webhook owners are tenants.
// A refused or unresolvable destination always gives the same text. In any other
// error, each address in the private and reserved ranges that BlockInternal
// refuses becomes a placeholder. Ranges a caller configures are not consulted.
func RedactSendError(text string) string {
	// net/http keeps only the text of an error it times out, so errors.Is would miss a blocked dial
	for _, cause := range noAllowedAddressCauses {
		if strings.Contains(text, cause) {
			return noAllowedAddressText
		}
	}
	text = ipv6Token.ReplaceAllStringFunc(text, redactAddress)
	return ipv4Token.ReplaceAllStringFunc(text, redactAddress)
}

func redactAddress(token string) string {
	// the IPv6 pattern takes in punctuation after a bare address, as in
	// "fd00::: i/o timeout", where only one of the colons is the separator
	for end := len(token); end > 0; end-- {
		host, _, _ := strings.Cut(token[:end], "%")
		if ip := net.ParseIP(host); ip != nil {
			if isBlockedIP(ip) {
				return "internal-address" + token[end:]
			}
			return token
		}
		if c := token[end-1]; c != ':' && c != '.' {
			return token
		}
	}
	return token
}
