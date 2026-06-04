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

// Package jwtutil offers shared helpers for signing and verifying JWTs with
// statically-configured multi-key sets — i.e. kid-based key rotation. It
// replaces the convenience that go-jose's JSONWebKeySet provided when picking
// the right key out of a set during verification, which golang-jwt/jwt/v5
// leaves to the caller.
package jwtutil

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrEmptyKeySet     = errors.New("jwtutil: key set is empty")
	ErrDuplicateKeyID  = errors.New("jwtutil: duplicate key ID")
	ErrUnknownKeyID    = errors.New("jwtutil: unknown signing key")
	ErrSigningMethod   = errors.New("jwtutil: unsupported signing method")
)

// HMACKey is one entry in an HMACKeySet.
type HMACKey struct {
	// ID is the value placed in (and matched against) the JWT's `kid` header.
	ID string
	// Key is the HMAC shared secret.
	Key []byte
}

// HMACKeyID derives a stable, opaque identifier for an HMAC key as the first
// 8 bytes of its SHA-256, base64(RawStdEncoding)-encoded. Use when keys are
// configured as opaque secrets without explicit IDs.
func HMACKeyID(key []byte) string {
	h := sha256.Sum256(key)
	return base64.RawStdEncoding.EncodeToString(h[:8])
}

// HMACKeySet is an ordered set of HMAC keys for signing and verifying JWTs
// with kid-based key rotation. The first key is the active signer; all keys
// in the set are accepted during verification, allowing graceful rotation
// without invalidating in-flight tokens.
type HMACKeySet struct {
	method jwt.SigningMethod
	keys   []HMACKey
	byID   map[string][]byte
}

// NewHMACKeySet builds an HMACKeySet for the given signing method. method
// must be one of jwt.SigningMethodHS256 / HS384 / HS512. The first key in
// keys is the active signer.
func NewHMACKeySet(method jwt.SigningMethod, keys ...HMACKey) (*HMACKeySet, error) {
	switch method {
	case jwt.SigningMethodHS256, jwt.SigningMethodHS384, jwt.SigningMethodHS512:
	default:
		return nil, fmt.Errorf("%w: %q", ErrSigningMethod, method.Alg())
	}
	if len(keys) == 0 {
		return nil, ErrEmptyKeySet
	}
	ks := &HMACKeySet{
		method: method,
		keys:   make([]HMACKey, len(keys)),
		byID:   make(map[string][]byte, len(keys)),
	}
	for i, k := range keys {
		if _, exists := ks.byID[k.ID]; exists {
			return nil, fmt.Errorf("%w: %q", ErrDuplicateKeyID, k.ID)
		}
		ks.keys[i] = k
		ks.byID[k.ID] = k.Key
	}
	return ks, nil
}

// NewHMACKeySetFromBytes is a convenience that derives each key's ID via
// HMACKeyID. Use when keys are configured as opaque secrets and the caller
// doesn't manage key IDs explicitly.
func NewHMACKeySetFromBytes(method jwt.SigningMethod, keys ...[]byte) (*HMACKeySet, error) {
	hmacKeys := make([]HMACKey, len(keys))
	for i, k := range keys {
		hmacKeys[i] = HMACKey{ID: HMACKeyID(k), Key: k}
	}
	return NewHMACKeySet(method, hmacKeys...)
}

// SigningMethod returns the signing method enforced by this key set.
func (ks *HMACKeySet) SigningMethod() jwt.SigningMethod {
	return ks.method
}

// Sign returns a compact-serialized JWT signed with the active (first) key.
// The active key's ID is set as the `kid` header.
func (ks *HMACKeySet) Sign(claims jwt.Claims) (string, error) {
	active := ks.keys[0]
	tok := jwt.NewWithClaims(ks.method, claims)
	tok.Header["kid"] = active.ID
	return tok.SignedString(active.Key)
}

// ParseWithClaims parses and verifies the token, selecting the signing key
// from the set by matching the token's `kid` header. The set's signing
// method is enforced via jwt.WithValidMethods. Additional parser options
// (WithIssuer, WithLeeway, WithExpirationRequired, …) may be passed.
func (ks *HMACKeySet) ParseWithClaims(token string, claims jwt.Claims, opts ...jwt.ParserOption) (*jwt.Token, error) {
	allOpts := make([]jwt.ParserOption, 0, len(opts)+1)
	allOpts = append(allOpts, jwt.WithValidMethods([]string{ks.method.Alg()}))
	allOpts = append(allOpts, opts...)
	return jwt.ParseWithClaims(token, claims, ks.keyFunc, allOpts...)
}

func (ks *HMACKeySet) keyFunc(tok *jwt.Token) (any, error) {
	kid, _ := tok.Header["kid"].(string)
	key, ok := ks.byID[kid]
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownKeyID, kid)
	}
	return key, nil
}
