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
	"testing"
	"time"

	jose "github.com/go-jose/go-jose/v4"
	josejwt "github.com/go-jose/go-jose/v4/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils/jwtutil"
)

// HS256 in go-jose v4 enforces RFC 7518's 32-byte HMAC key minimum.
var interopKey = []byte("interop-secret-padded-to-32-bytes!!")

func TestGoJoseInterop_GoJoseSignsJwtutilVerifies(t *testing.T) {
	signerOpts := (&jose.SignerOptions{}).WithType("JWT").WithHeader("kid", "a")
	joseSigner, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: interopKey}, signerOpts)
	require.NoError(t, err)

	exp := time.Now().Add(time.Minute).Truncate(time.Second)
	tokStr, err := josejwt.Signed(joseSigner).Claims(josejwt.Claims{
		Issuer: "go-jose",
		Expiry: josejwt.NewNumericDate(exp),
	}).Serialize()
	require.NoError(t, err)

	ks, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, jwtutil.HMACKey{ID: "a", Key: interopKey})
	require.NoError(t, err)
	out := &jwt.RegisteredClaims{}
	parsed, err := ks.ParseWithClaims(tokStr, out)
	require.NoError(t, err)
	require.Equal(t, "a", parsed.Header["kid"])
	require.Equal(t, "go-jose", out.Issuer)
	require.True(t, out.ExpiresAt.Equal(exp))
}

func TestGoJoseInterop_JwtutilSignsGoJoseVerifies(t *testing.T) {
	ks, err := jwtutil.NewHMACKeySet(jwt.SigningMethodHS256, jwtutil.HMACKey{ID: "a", Key: interopKey})
	require.NoError(t, err)

	exp := time.Now().Add(time.Minute).Truncate(time.Second)
	tokStr, err := ks.Sign(&jwt.RegisteredClaims{
		Issuer:    "jwtutil",
		ExpiresAt: jwt.NewNumericDate(exp),
	})
	require.NoError(t, err)

	parsed, err := josejwt.ParseSigned(tokStr, []jose.SignatureAlgorithm{jose.HS256})
	require.NoError(t, err)
	require.Len(t, parsed.Headers, 1)
	require.Equal(t, "a", parsed.Headers[0].KeyID)

	out := josejwt.Claims{}
	require.NoError(t, parsed.Claims(interopKey, &out))
	require.Equal(t, "jwtutil", out.Issuer)
	require.True(t, out.Expiry.Time().Equal(exp))
}
