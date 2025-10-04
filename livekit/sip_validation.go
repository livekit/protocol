// Copyright 2023 LiveKit, Inc.
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

package livekit

import (
	"errors"
	fmt "fmt"
	"regexp"
	"strings"
)

// RFC 3261 compliant validation functions for SIP headers and messages

// RFC 3261 Section 25.1 - Header field names
// token = 1*(alphanum / "-" / "." / "!" / "%" / "*" / "_" / "+" / "`" / "'" / "~")
// Specifically lowercase since we're converting to lowercase for case-insensitive comparison
var reHeaderName = regexp.MustCompile(`^[a-z0-9\-\.!%*_+` + "`" + `'~]+$`)

// RFC 3261 Section 25.1 - Header field values (basic validation)
// More specific validation is done per header type
var reHeaderValueBasic = regexp.MustCompile(`^[\x20-\x7E]*$`)

// RFC 3261 Section 19.1 - SIP URI validation
var reSIPURI = regexp.MustCompile(`^(sip|sips):([^@]+@)?([^;]+)(;.*)?$`)

// Required headers for SIP requests per RFC 3261 Section 8.1.1
var RequiredRequestHeaders = map[string]bool{
	"via":          true,
	"from":         true,
	"to":           true,
	"call-id":      true,
	"cseq":         true,
	"max-forwards": true,
}

// Required headers for SIP responses per RFC 3261 Section 8.2.1
var RequiredResponseHeaders = map[string]bool{
	"via":     true,
	"from":    true,
	"to":      true,
	"call-id": true,
	"cseq":    true,
}

// Crucial headers that can't be overridden by the user, and their shorthands
var FrobiddenSipHeaderNames = map[string]bool{
	"accept":           true,
	"accept-encoding":  true,
	"accept-language":  true,
	"allow":            true,
	"allow-events":     true, // rfc3903
	"call-id":          true,
	"contact":          true,
	"content-encoding": true,
	"content-length":   true,
	"content-type":     true,
	"cseq":             true,
	"event":            true, // rfc3903
	"expires":          true,
	"from":             true, // We might allow this in the future, but for now we're printing
	"max-forwards":     true,
	"record-route":     true,
	"refer-to":         true, // rfc3515
	"referred-by":      true, // rfc3892
	"reply-to":         true,
	"route":            true,
	"supported":        true,
	"to":               true, // We might allow this in the future, but for now we're printing
	"via":              true,

	// Single-letter shorthands, a.k.a compact form
	"b": true, // Referred-By; rfc3892
	"c": true, // Content-Type
	"e": true, // Content-Encoding
	"f": true, // From
	"i": true, // Call-ID
	"k": true, // Supported
	"l": true, // Content-Length
	"m": true, // Contact
	"o": true, // Event; rfc3903
	"r": true, // Refer-To; rfc3515
	"t": true, // To
	"u": true, // Allow-Events; rfc3903
	"v": true, // Via
}

// Headers that must comply with name-addr specification per RFC 3261 Section 20.10
// name-addr = [display-name] <addr-spec>
// addr-spec = SIP-URI / SIPS-URI / absoluteURI
var nameAddrHeaders = map[string]bool{
	"from":                true,
	"to":                  true,
	"contact":             true,
	"route":               true,
	"record-route":        true,
	"reply-to":            true,
	"p-asserted-identity": true, // RFC 3325 Section 9.1
}

// ValidateHeaderName validates a SIP header name per RFC 3261 Section 25.1
func ValidateHeaderName(name string) error {
	if name == "" {
		return errors.New("header name cannot be empty")
	}

	if len(name) > 255 {
		return errors.New("header name too long (max 255 characters)")
	}

	lowerName := strings.ToLower(name)
	if !reHeaderName.MatchString(lowerName) {
		return errors.New("header name contains invalid characters")
	}

	// Convert to lowercase for case-insensitive comparison
	if forbidden, exists := FrobiddenSipHeaderNames[lowerName]; exists && forbidden {
		return fmt.Errorf("header name %s not supported", name)
	}

	return nil
}

// ValidateHeaderValue validates a SIP header value per RFC 3261 Section 25.1
func ValidateHeaderValue(name, value string) error {
	if value == "" {
		return errors.New("header value cannot be empty")
	}

	if len(value) > 1024 {
		return errors.New("header value too long (max 1024 characters)")
	}

	// Basic character validation - printable ASCII
	if !reHeaderValueBasic.MatchString(value) {
		return errors.New("header value contains invalid characters")
	}

	// Convert to lowercase for case-insensitive comparison
	lowerName := strings.ToLower(name)
	if _, exists := nameAddrHeaders[lowerName]; exists && false {
		// TODO: Disabled since all supported headers are forbidden, re-enable when we allow some
		return validateNameAddrHeader(value)
	}

	return nil
}

// findAngleBrackets efficiently finds angle brackets in a single scan
// Returns: start, end positions (-1 = missing), and error status
func findAngleBrackets(value string) (int, int, error) {
	start := -1
	end := -1

	for i, r := range value {
		switch r {
		case '<':
			if start != -1 {
				return -1, -1, errors.New("multiple opening brackets")
			}
			start = i
		case '>':
			if end != -1 {
				return -1, -1, errors.New("multiple closing brackets")
			}
			end = i
		}
	}

	// Check for mismatched brackets
	if (start == -1) != (end == -1) {
		return -1, -1, errors.New("mismatched angle brackets")
	}

	// Check that < comes before >
	if start > end {
		return -1, -1, errors.New("malformed angle brackets")
	}

	return start, end, nil
}

// validateNameAddrHeader validates headers that use name-addr format per RFC 3261 Section 20.10
func validateNameAddrHeader(value string) error {
	// RFC 3261 Section 20.10 - name-addr format
	// name-addr = [display-name] <addr-spec>
	// addr-spec = SIP-URI / SIPS-URI / absoluteURI

	start, end, err := findAngleBrackets(value)
	if err != nil {
		return err
	}
	if start >= 0 || end >= 0 {
		displayName := strings.TrimSpace(value[:start])
		uri := value[start+1 : end]
		if displayName != "" {
			if err := validateDisplayName(displayName); err != nil {
				return err
			}
		}
		// Keep in mind, this ignores header parameters, while validating URI parameters
		return validateURI(uri)
	}

	// This is a bare URI, and should comply with addr-spec, no special characters
	if strings.ContainsAny(value, ";,? ") {
		return errors.New("bare URI with special characters")
	}

	return validateURI(value)
}

// validateDisplayName validates a display name in name-addr format
func validateDisplayName(displayName string) error {
	if displayName == "" {
		return nil
	}

	// Check if display name is quoted
	if strings.HasPrefix(displayName, "\"") && strings.HasSuffix(displayName, "\"") {
		// Quoted display name - basic validation
		quoted := displayName[1 : len(displayName)-1]
		// Check for proper escaping
		if strings.Contains(quoted, "\"") && !strings.Contains(quoted, "\\\"") {
			return errors.New("unquoted display name contains unescaped quotes")
		}
		return nil
	}

	// Unquoted display name - must not contain special characters
	for _, r := range displayName {
		if r == '<' || r == '>' || r == '"' || r == '\\' || r == '&' {
			return errors.New("unquoted display name contains special characters")
		}
	}

	return nil
}

// validateURI validates URIs that can appear in name-addr format
func validateURI(uri string) error {
	if uri == "" {
		return errors.New("URI cannot be empty")
	}

	// Check for SIP/SIPS scheme
	if strings.HasPrefix(uri, "sip:") || strings.HasPrefix(uri, "sips:") {
		return validateSIPURI(uri)
	}

	// Check for TEL scheme
	if strings.HasPrefix(uri, "tel:") {
		return validateTELURI(uri)
	}

	// For now, only support SIP/SIPS and TEL URIs
	// RFC 3261 allows other absolute URIs, but we'll be restrictive
	return errors.New("URI scheme must match one of sip, sips, or tel")
}

// validateSIPURI validates a SIP or SIPS URI per RFC 3261 Section 19.1
func validateSIPURI(uri string) error {
	if uri == "" {
		return errors.New("SIP URI cannot be empty")
	}

	// Check for SIP or SIPS scheme
	if !strings.HasPrefix(uri, "sip:") && !strings.HasPrefix(uri, "sips:") {
		return errors.New("SIP URI scheme must match sip or sips")
	}

	// Basic format validation
	if !reSIPURI.MatchString(uri) {
		return errors.New("SIP URI format is invalid")
	}

	// Validate URI parameters if present
	if strings.Contains(uri, ";") {
		return validateURIParameters(uri)
	}

	return nil
}

// validateURIParameters validates URI parameters
func validateURIParameters(uri string) error {
	// Split URI into base and parameters
	parts := strings.Split(uri, ";")
	if len(parts) < 2 {
		return errors.New("invalid URI parameters")
	}

	// Validate each parameter
	for _, param := range parts[1:] {
		param = strings.TrimSpace(param)
		if param == "" {
			return errors.New("empty URI parameter")
		}

		// Check for valid parameter format: name=value or name
		if strings.Contains(param, "=") {
			paramParts := strings.SplitN(param, "=", 2)
			if len(paramParts) != 2 {
				return errors.New("invalid URI parameter format")
			}
			name := strings.TrimSpace(paramParts[0])
			value := strings.TrimSpace(paramParts[1])

			if name == "" {
				return errors.New("URI parameter name cannot be empty")
			}
			if value == "" {
				return errors.New("URI parameter value cannot be empty")
			}

		} else {
			// Parameter without value - just validate name
			if strings.TrimSpace(param) == "" {
				return errors.New("URI parameter name cannot be empty")
			}
		}

		// Check for invalid characters in parameter
		if strings.Contains(param, " ") {
			return errors.New("URI parameter contains spaces")
		}
	}

	return nil
}

// validateTELURI validates a TEL URI per RFC 3966
func validateTELURI(uri string) error {
	if uri == "" {
		return errors.New("TEL URI cannot be empty")
	}

	if !strings.HasPrefix(uri, "tel:") {
		return errors.New("TEL URI scheme must match tel")
	}

	// Basic validation - TEL URIs are more complex, this is simplified
	if len(uri) < 5 { // "tel:" + at least one character
		return errors.New("TEL URI format is invalid")
	}

	return nil
}

// EscapeHeaderValue escapes special characters in header values per RFC 3261
func EscapeHeaderValue(value string) string {
	// Escape special characters that need to be quoted
	var result strings.Builder
	for _, r := range value {
		switch r {
		case ' ', '\t', '\r', '\n', '"', '\\':
			// These characters need to be escaped or quoted
			result.WriteString("\\")
			result.WriteRune(r)
		default:
			result.WriteRune(r)
		}
	}
	return result.String()
}

// UnescapeHeaderValue removes escaping from header values
func UnescapeHeaderValue(value string) string {
	var result strings.Builder
	escaped := false

	for _, r := range value {
		if escaped {
			result.WriteRune(r)
			escaped = false
		} else if r == '\\' {
			escaped = true
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
