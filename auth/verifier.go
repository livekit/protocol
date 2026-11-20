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
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// matches go-jose's previous DefaultLeeway for validation tolerance.
const tokenLeeway = time.Minute

// TokenTypeV2 is the typ header of a token signed with a version 2 key.
const TokenTypeV2 = "livekit-v2+jwt"

var (
	// Parsing runs before the key is known, so it allows both. The key picks one.
	allowedSigningMethods = []string{jwt.SigningMethodHS256.Alg(), jwt.SigningMethodEdDSA.Alg()}
	unverifiedParser      = jwt.NewParser(jwt.WithValidMethods(allowedSigningMethods))
)

// VerificationKey is what a key record gives the verifier.
//
// Do not use version 2 keys yet: no LiveKit server or SDK supports them.
type VerificationKey struct {
	Version   int               // 1: HS256 with Secret, 2: EdDSA with PublicKey
	Secret    string            // version 1
	PublicKey ed25519.PublicKey // version 2
}

type APIKeyTokenVerifier struct {
	raw      string
	identity string
	apiKey   string
	alg      string
	keyID    string // kid header
	typ      string // typ header
}

// ParseAPIToken parses an encoded JWT token without verifying its signature.
// The signature is verified later by Verify or VerifyWithKey, which require the key.
func ParseAPIToken(raw string) (*APIKeyTokenVerifier, error) {
	claims := &jwt.RegisteredClaims{}
	token, _, err := unverifiedParser.ParseUnverified(raw, claims)
	if err != nil {
		return nil, err
	}

	v := &APIKeyTokenVerifier{
		raw:      raw,
		apiKey:   claims.Issuer,
		identity: claims.Subject,
		alg:      token.Method.Alg(),
	}
	v.keyID, _ = token.Header["kid"].(string)
	v.typ, _ = token.Header["typ"].(string)
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

// Verify verifies a token of a version 1 key: HS256 with the API secret.
func (v *APIKeyTokenVerifier) Verify(key interface{}) (*jwt.RegisteredClaims, *ClaimGrants, error) {
	if key == nil || key == "" {
		return nil, nil, ErrKeysMissing
	}
	if s, ok := key.(string); ok {
		key = []byte(s)
	}
	return v.verify(jwt.SigningMethodHS256, key)
}

// VerifyWithKey picks the algorithm from the key, never from the token header:
// HS256 with the secret for version 1, EdDSA with the public key for version 2.
// A version 2 token must also name the API key in its kid header and have the
// type TokenTypeV2.
func (v *APIKeyTokenVerifier) VerifyWithKey(k VerificationKey) (*jwt.RegisteredClaims, *ClaimGrants, error) {
	switch k.Version {
	case 1:
		if k.Secret == "" {
			return nil, nil, ErrKeysMissing
		}
		return v.verify(jwt.SigningMethodHS256, []byte(k.Secret))
	case 2:
		if len(k.PublicKey) == 0 {
			return nil, nil, ErrKeysMissing
		}
		if len(k.PublicKey) != ed25519.PublicKeySize {
			return nil, nil, ErrInvalidPublicKey
		}
		claims, grants, err := v.verify(jwt.SigningMethodEdDSA, k.PublicKey)
		if err != nil {
			return nil, nil, err
		}
		if v.keyID != v.apiKey {
			return nil, nil, ErrKeyIDMismatch
		}
		if v.typ != TokenTypeV2 {
			return nil, nil, ErrTokenTypeMismatch
		}
		return claims, grants, nil
	default:
		return nil, nil, fmt.Errorf("%w: %d", ErrUnsupportedKeyVersion, k.Version)
	}
}

func (v *APIKeyTokenVerifier) verify(method jwt.SigningMethod, key interface{}) (*jwt.RegisteredClaims, *ClaimGrants, error) {
	if v.alg != method.Alg() {
		return nil, nil, fmt.Errorf("%w: key requires %s, token uses %s", ErrKeyAlgorithmMismatch, method.Alg(), v.alg)
	}

	claims := &tokenClaims{}
	_, err := jwt.ParseWithClaims(v.raw, claims, func(*jwt.Token) (interface{}, error) {
		return key, nil
	},
		jwt.WithValidMethods([]string{method.Alg()}),
		jwt.WithIssuer(v.apiKey),
		jwt.WithLeeway(tokenLeeway),
		// Without this, a token carrying no exp claim passes validation and
		// never expires. First-party SDKs always set exp, but any hand-rolled
		// or third-party minter that forgets it silently mints a permanent
		// credential.
		jwt.WithExpirationRequired(),
		// Without this, iat is never validated. First-party SDKs always set
		// nbf (see AccessToken.ToJWT), which the parser already validates by
		// default when present, but a token that omits nbf entirely verifies
		// immediately regardless of how far in the future it claims to have
		// been issued unless iat is checked too. (A token omitting both iat
		// and nbf is unaffected by this check either way.)
		jwt.WithIssuedAt(),
	)
	if err != nil {
		return nil, nil, err
	}

	claims.ClaimGrants.Identity = v.identity
	return &claims.RegisteredClaims, &claims.ClaimGrants, nil
}
