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
	claims := &jwt.RegisteredClaims{}
	if _, _, err := unverifiedParser.ParseUnverified(raw, claims); err != nil {
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

// APIKey returns the API key this token was signed with
func (v *APIKeyTokenVerifier) APIKey() string {
	return v.apiKey
}

func (v *APIKeyTokenVerifier) Identity() string {
	return v.identity
}

func (v *APIKeyTokenVerifier) Verify(key interface{}) (*jwt.RegisteredClaims, *ClaimGrants, error) {
	if key == nil || key == "" {
		return nil, nil, ErrKeysMissing
	}
	if s, ok := key.(string); ok {
		key = []byte(s)
	}

	claims := &tokenClaims{}
	_, err := jwt.ParseWithClaims(v.raw, claims, func(*jwt.Token) (interface{}, error) {
		return key, nil
	},
		jwt.WithValidMethods(allowedSigningMethods),
		jwt.WithIssuer(v.apiKey),
		jwt.WithLeeway(tokenLeeway),
	)
	if err != nil {
		return nil, nil, err
	}

	claims.ClaimGrants.Identity = v.identity
	return &claims.RegisteredClaims, &claims.ClaimGrants, nil
}
