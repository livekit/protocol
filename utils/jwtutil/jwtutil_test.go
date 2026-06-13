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

package jwtutil_test

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils/jwtutil"
)

func TestHMACKeyID_StableSHA256First8(t *testing.T) {
	// Wire-compat with cloud-protocol/hostedagentsutil: id is base64(RawStdEncoding, sha256(key)[:8]).
	key := []byte("the-secret")
	want := func() string {
		h := sha256.Sum256(key)
		return base64.RawStdEncoding.EncodeToString(h[:8])
	}()
	require.Equal(t, want, jwtutil.HMACKeyID(key))
}

func TestNewHMACKeySet_ValidatesMethod(t *testing.T) {
	_, err := jwtutil.NewHMACKeySet(jwt.SigningMethodRS256, jwtutil.HMACKey{ID: "a", Key: []byte("k")})
	require.ErrorIs(t, err, jwtutil.ErrSigningMethod)
}

func TestNewHMACKeySet_RejectsEmpty(t *testing.T) {
	_, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256)
	require.ErrorIs(t, err, jwtutil.ErrEmptyKeySet)
}

func TestNewHMACKeySet_RejectsDuplicateID(t *testing.T) {
	_, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256,
		jwtutil.HMACKey{ID: "x", Key: []byte("k1")},
		jwtutil.HMACKey{ID: "x", Key: []byte("k2")},
	)
	require.ErrorIs(t, err, jwtutil.ErrDuplicateKeyID)
}

func TestSignVerifyRoundTrip(t *testing.T) {
	ks, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, jwtutil.HMACKey{ID: "a", Key: []byte("secret-a")})
	require.NoError(t, err)

	claims := jwt.RegisteredClaims{
		Issuer:    "tester",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
	}
	tokStr, err := ks.Sign(&claims)
	require.NoError(t, err)

	out := &jwt.RegisteredClaims{}
	parsed, err := ks.ParseWithClaims(tokStr, out)
	require.NoError(t, err)
	require.Equal(t, "a", parsed.Header["kid"])
	require.Equal(t, "tester", out.Issuer)
}

func TestRotation_NewSetAcceptsOldKID(t *testing.T) {
	keyA := jwtutil.HMACKey{ID: "a", Key: []byte("secret-a")}
	keyB := jwtutil.HMACKey{ID: "b", Key: []byte("secret-b")}

	signer, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, keyA, keyB)
	require.NoError(t, err)
	tokStr, err := signer.Sign(&jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
	})
	require.NoError(t, err)

	// Rotate: B is now the active signer but A is still accepted.
	verifier, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, keyB, keyA)
	require.NoError(t, err)
	out := &jwt.RegisteredClaims{}
	_, err = verifier.ParseWithClaims(tokStr, out)
	require.NoError(t, err)
}

func TestRotation_OldVerifierRejectsNewKID(t *testing.T) {
	keyA := jwtutil.HMACKey{ID: "a", Key: []byte("secret-a")}
	keyB := jwtutil.HMACKey{ID: "b", Key: []byte("secret-b")}

	// New signer uses key B; old verifier only knows A.
	signer, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, keyB)
	require.NoError(t, err)
	tokStr, err := signer.Sign(&jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
	})
	require.NoError(t, err)

	verifier, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, keyA)
	require.NoError(t, err)
	_, err = verifier.ParseWithClaims(tokStr, &jwt.RegisteredClaims{})
	require.Error(t, err)
	require.ErrorIs(t, err, jwtutil.ErrUnknownKeyID)
}

func TestVerify_RejectsTokenSignedWithWrongKeyBytes(t *testing.T) {
	// Two sets share a kid ("a") but the actual key bytes differ — verify must
	// reject the foreign signature instead of trusting the kid alone.
	signer, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, jwtutil.HMACKey{ID: "a", Key: []byte("attacker-key")})
	require.NoError(t, err)
	tokStr, err := signer.Sign(&jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
	})
	require.NoError(t, err)

	verifier, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, jwtutil.HMACKey{ID: "a", Key: []byte("real-key")})
	require.NoError(t, err)
	_, err = verifier.ParseWithClaims(tokStr, &jwt.RegisteredClaims{})
	require.Error(t, err)
	require.ErrorIs(t, err, jwt.ErrSignatureInvalid)
}

func TestVerify_EnforcesSigningMethod(t *testing.T) {
	// HS256 verifier rejects an HS512-signed token even when the kid + secret match.
	key := jwtutil.HMACKey{ID: "a", Key: []byte("secret")}
	tok512 := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
	})
	tok512.Header["kid"] = key.ID
	tokStr, err := tok512.SignedString(key.Key)
	require.NoError(t, err)

	verifier, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, key)
	require.NoError(t, err)
	_, err = verifier.ParseWithClaims(tokStr, &jwt.RegisteredClaims{})
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), "signing method"), "got %v", err)
}

func TestParseWithClaims_ExtraOptionsApplied(t *testing.T) {
	key := jwtutil.HMACKey{ID: "a", Key: []byte("secret")}
	ks, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, key)
	require.NoError(t, err)

	tokStr, err := ks.Sign(&jwt.RegisteredClaims{
		Issuer:    "expected-issuer",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
	})
	require.NoError(t, err)

	_, err = ks.ParseWithClaims(tokStr, &jwt.RegisteredClaims{}, jwt.WithIssuer("not-this-issuer"))
	require.Error(t, err)
	require.ErrorIs(t, err, jwt.ErrTokenInvalidIssuer)
}

func TestNewHMACKeySetFromBytes_DerivesIDs(t *testing.T) {
	ks, err := jwtutil.NewHMACKeySetFromBytes(jwt.SigningMethodHS256, []byte("k1"), []byte("k2"))
	require.NoError(t, err)

	tokStr, err := ks.Sign(&jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
	})
	require.NoError(t, err)

	// kid should match HMACKeyID("k1") (the active key)
	parsed, err := ks.ParseWithClaims(tokStr, &jwt.RegisteredClaims{})
	require.NoError(t, err)
	require.Equal(t, jwtutil.HMACKeyID([]byte("k1")), parsed.Header["kid"])
}
