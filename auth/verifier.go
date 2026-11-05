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

package auth

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// matches go-jose's previous DefaultLeeway for validation tolerance.
const tokenLeeway = time.Minute

var (
	allowedSigningMethods = []string{jwt.SigningMethodHS256.Alg()}
	unverifiedParser      = jwt.NewParser(jwt.WithValidMethods(allowedSigningMethods))
)

type APIKeyTokenVerifier struct {
	raw      string
	identity string
	apiKey   string
}

// ParseAPIToken parses an encoded JWT token without verifying its signature.
// The signature is verified later by Verify, which requires the API secret.
func ParseAPIToken(raw string) (*APIKeyTokenVerifier, error) {
	// fixed structs skip every other member, so unsigned input is never built into maps
	var header struct {
		Alg string `json:"alg"`
	}
	var claims struct {
		Issuer  string `json:"iss"`
		Subject string `json:"sub"`
		ID      string `json:"jti"`
	}
	if strings.Count(raw, ".") != 2 {
		return nil, fmt.Errorf("%w: token contains an invalid number of segments", jwt.ErrTokenMalformed)
	}
	h, rest, _ := strings.Cut(raw, ".")
	c, _, _ := strings.Cut(rest, ".")
	if err := decodeSegment(h, &header); err != nil {
		return nil, err
	}
	if jwt.GetSigningMethod(header.Alg) == nil {
		return nil, fmt.Errorf("%w: signing method (alg) is unavailable", jwt.ErrTokenUnverifiable)
	}
	if err := decodeSegment(c, &claims); err != nil {
		return nil, err
	}

	v := &APIKeyTokenVerifier{
		raw:      raw,
		apiKey:   claims.Issuer,
		identity: claims.Subject,
	}
	if v.identity == "" {
		v.identity = claims.ID
	}
	return v, nil
}

func decodeSegment(seg string, v any) error {
	b, err := unverifiedParser.DecodeSegment(seg)
	if err == nil {
		err = json.Unmarshal(b, v)
	}
	if err != nil {
		return fmt.Errorf("%w: %w", jwt.ErrTokenMalformed, err)
	}
	return nil
}

// APIKey returns the API key this token was signed with
func (v *APIKeyTokenVerifier) APIKey() string {
	return v.apiKey
}

func (v *APIKeyTokenVerifier) Identity() string {
	return v.identity
}

func (v *APIKeyTokenVerifier) Verify(key interface{}) (*jwt.RegisteredClaims, *ClaimGrants, error) {
	// HS256 is the only allowed method, and anyone can sign with an empty HMAC key
	var secret []byte
	switch k := key.(type) {
	case string:
		secret = []byte(k)
	case []byte:
		secret = k
	}
	if len(secret) == 0 {
		return nil, nil, ErrKeysMissing
	}

	// check the MAC first: golang-jwt decodes the header and grants before it does
	i := strings.LastIndexByte(v.raw, '.')
	if i < 0 {
		return nil, nil, jwt.ErrTokenMalformed
	}
	sig, err := unverifiedParser.DecodeSegment(v.raw[i+1:])
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", jwt.ErrTokenMalformed, err)
	}
	if err := jwt.SigningMethodHS256.Verify(v.raw[:i], sig, secret); err != nil {
		return nil, nil, fmt.Errorf("%w: %w", jwt.ErrTokenSignatureInvalid, err)
	}

	claims := &tokenClaims{}
	_, err = jwt.ParseWithClaims(v.raw, claims, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	},
		jwt.WithValidMethods(allowedSigningMethods),
		jwt.WithIssuer(v.apiKey),
		jwt.WithLeeway(tokenLeeway),
		// iat is optional, but a token must not be used before it was issued
		jwt.WithIssuedAt(),
		// Without this, a token carrying no exp claim passes validation and
		// never expires. First-party SDKs always set exp, but any hand-rolled
		// or third-party minter that forgets it silently mints a permanent
		// credential.
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, nil, err
	}

	claims.ClaimGrants.Identity = v.identity
	return &claims.RegisteredClaims, &claims.ClaimGrants, nil
}
