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

package auth_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/auth"
)

func TestVerifier(t *testing.T) {
	apiKey := "APID3B67uxk4Nj2GKiRPibAZ9"
	secret := "YHC-CUhbQhGeVCaYgn1BNA++"
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDg5MzAzMDgsImlzcyI6IkFQSUQzQjY3dXhrNE5qMkdLaVJQaWJBWjkiLCJuYmYiOjE2MDg5MjY3MDgsInJvb21fam9pbiI6dHJ1ZSwicm9vbV9zaWQiOiJteWlkIiwic3ViIjoiQVBJRDNCNjd1eGs0TmoyR0tpUlBpYkFaOSJ9.cmHEBq0MLyRqphmVLM2cLXg5ao5Sro7am8yXhcYKcwE"
	t.Run("cannot decode with incorrect key", func(t *testing.T) {
		v, err := auth.ParseAPIToken(accessToken)
		require.NoError(t, err)

		require.Equal(t, apiKey, v.APIKey())
		_, _, err = v.Verify("")
		require.Error(t, err)

		_, _, err = v.Verify("anothersecret")
		require.Error(t, err)
	})

	t.Run("key has expired", func(t *testing.T) {
		v, err := auth.ParseAPIToken(accessToken)
		require.NoError(t, err)

		_, _, err = v.Verify(secret)
		require.Error(t, err)
	})

	t.Run("token without exp is rejected", func(t *testing.T) {
		// hand-rolled JWT with no exp claim (the Go SDK always sets one, so
		// build it directly to model a third-party minter forgetting exp)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": apiKey,
			"nbf": jwt.NewNumericDate(time.Now().Add(-time.Minute)),
			"video": map[string]interface{}{
				"roomCreate": true,
			},
		})
		authToken, err := token.SignedString([]byte(secret))
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(authToken)
		require.NoError(t, err)

		_, _, err = v.Verify(secret)
		require.Error(t, err)
	})

	t.Run("token issued in the future without nbf is rejected", func(t *testing.T) {
		// hand-rolled JWT with iat 2h in the future and no nbf claim. The Go
		// SDK always sets nbf (see AccessToken.ToJWT), so build this directly
		// to model a third-party minter that omits it: without nbf, and
		// without WithIssuedAt() on the parser, nothing stops a token from
		// verifying immediately no matter how far in the future it claims to
		// have been issued.
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": apiKey,
			"iat": jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			"exp": jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			"video": map[string]interface{}{
				"roomCreate": true,
			},
		})
		authToken, err := token.SignedString([]byte(secret))
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(authToken)
		require.NoError(t, err)

		_, _, err = v.Verify(secret)
		require.Error(t, err)
	})

	t.Run("unexpired token is verified", func(t *testing.T) {
		claim := auth.VideoGrant{RoomCreate: true}
		at := auth.NewAccessToken(apiKey, secret).
			SetVideoGrant(&claim).
			SetValidFor(time.Minute).
			SetIdentity("me")
		authToken, err := at.ToJWT()
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(authToken)
		require.NoError(t, err)
		require.Equal(t, apiKey, v.APIKey())
		require.Equal(t, "me", v.Identity())

		_, decoded, err := v.Verify(secret)
		require.NoError(t, err)
		require.Equal(t, &claim, decoded.Video)
	})

	t.Run("ensure metadata can be passed through", func(t *testing.T) {
		metadata := map[string]interface{}{
			"user":   "value",
			"number": float64(3),
		}
		md, _ := json.Marshal(metadata)
		attrs := map[string]string{"mykey": "myval", "secondkey": "secondval"}
		at := auth.NewAccessToken(apiKey, secret).
			SetVideoGrant(&auth.VideoGrant{
				RoomAdmin: true,
				Room:      "myroom",
			}).
			SetMetadata(string(md)).
			SetAttributes(attrs)

		authToken, err := at.ToJWT()
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(authToken)
		require.NoError(t, err)

		_, decoded, err := v.Verify(secret)
		require.NoError(t, err)

		require.EqualValues(t, string(md), decoded.Metadata)
		require.EqualValues(t, attrs, decoded.Attributes)
	})

	t.Run("unknown fields are ignored for forward compatibility", func(t *testing.T) {
		// Simulate a token issued by a newer client whose claims include fields
		// this server does not yet know about. The server should still accept the
		// token rather than failing with `unknown field`. This guards against
		// requiring server upgrades before client upgrades can roll out.
		claims := jwt.MapClaims{
			"iss": apiKey,
			"sub": "me",
			"nbf": jwt.NewNumericDate(time.Now()),
			"exp": jwt.NewNumericDate(time.Now().Add(time.Minute)),
			// unknown top-level claim grants field
			"someFutureGrant": map[string]any{"enabled": true},
			"video": map[string]any{
				"roomJoin": true,
				"room":     "myroom",
				// unknown field inside a known grant
				"someFutureVideoField": "future-value",
			},
			"roomConfig": map[string]any{
				"name": "myroom",
				// unknown field inside a protojson-decoded message
				"someFutureRoomConfigField": "future-value",
			},
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(token)
		require.NoError(t, err)

		_, decoded, err := v.Verify(secret)
		require.NoError(t, err)
		require.NotNil(t, decoded.Video)
		require.Equal(t, "myroom", decoded.Video.Room)
		require.True(t, decoded.Video.RoomJoin)
		require.NotNil(t, decoded.RoomConfig)
		require.Equal(t, "myroom", decoded.RoomConfig.Name)
	})

	t.Run("nil permissions are handled", func(t *testing.T) {
		grant := &auth.VideoGrant{
			Room:     "myroom",
			RoomJoin: true,
		}
		grant.SetCanPublishData(false)
		at := auth.NewAccessToken(apiKey, secret).
			SetVideoGrant(grant)
		token, err := at.ToJWT()
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(token)
		require.NoError(t, err)
		_, decoded, err := v.Verify(secret)
		require.NoError(t, err)

		require.Nil(t, decoded.Video.CanSubscribe)
		require.Nil(t, decoded.Video.CanPublish)
		require.False(t, *decoded.Video.CanPublishData)
	})
}

func TestVerifyWithKey(t *testing.T) {
	apiKey := "APID3B67uxk4Nj2GKiRPibAZ9"
	secret := "YHC-CUhbQhGeVCaYgn1BNA++"
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	v1Key := auth.VerificationKey{Version: 1, Secret: secret}
	v2Key := auth.VerificationKey{Version: 2, PublicKey: pub}

	grant := &auth.VideoGrant{RoomJoin: true, Room: "myroom"}
	hsToken, err := auth.NewAccessToken(apiKey, secret).SetVideoGrant(grant).SetIdentity("me").ToJWT()
	require.NoError(t, err)
	edToken, err := auth.NewAccessTokenWithPrivateKey(apiKey, priv).SetVideoGrant(grant).SetIdentity("me").ToJWT()
	require.NoError(t, err)

	parse := func(t *testing.T, token string) *auth.APIKeyTokenVerifier {
		v, err := auth.ParseAPIToken(token)
		require.NoError(t, err)
		require.Equal(t, apiKey, v.APIKey())
		require.Equal(t, "me", v.Identity())
		return v
	}

	t.Run("v1 key verifies HS256", func(t *testing.T) {
		_, decoded, err := parse(t, hsToken).VerifyWithKey(v1Key)
		require.NoError(t, err)
		require.Equal(t, grant, decoded.Video)
	})

	t.Run("v2 key verifies EdDSA", func(t *testing.T) {
		_, decoded, err := parse(t, edToken).VerifyWithKey(v2Key)
		require.NoError(t, err)
		require.Equal(t, grant, decoded.Video)
	})

	t.Run("v2 key refuses HS256", func(t *testing.T) {
		_, _, err := parse(t, hsToken).VerifyWithKey(v2Key)
		require.ErrorIs(t, err, auth.ErrKeyAlgorithmMismatch)
	})

	t.Run("v1 key refuses EdDSA", func(t *testing.T) {
		_, _, err := parse(t, edToken).VerifyWithKey(v1Key)
		require.ErrorIs(t, err, auth.ErrKeyAlgorithmMismatch)
	})

	t.Run("v2 key refuses HS256 signed with the public key as secret", func(t *testing.T) {
		// algorithm confusion: the attacker knows the public key and uses it as an HMAC secret
		token, err := auth.NewAccessToken(apiKey, string(pub)).SetVideoGrant(grant).SetIdentity("me").ToJWT()
		require.NoError(t, err)
		_, _, err = parse(t, token).VerifyWithKey(v2Key)
		require.ErrorIs(t, err, auth.ErrKeyAlgorithmMismatch)
	})

	t.Run("v2 key refuses EdDSA signed with another key", func(t *testing.T) {
		_, otherPriv, err := ed25519.GenerateKey(rand.Reader)
		require.NoError(t, err)
		token, err := auth.NewAccessTokenWithPrivateKey(apiKey, otherPriv).SetIdentity("me").ToJWT()
		require.NoError(t, err)
		_, _, err = parse(t, token).VerifyWithKey(v2Key)
		require.Error(t, err)
	})

	t.Run("v2 key refuses a token with no kid or a kid that is not the issuer", func(t *testing.T) {
		for _, kid := range []string{"", "APIother"} {
			tok := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
				"iss": apiKey,
				"sub": "me",
				"exp": jwt.NewNumericDate(time.Now().Add(time.Hour)),
			})
			if kid != "" {
				tok.Header["kid"] = kid
			}
			token, err := tok.SignedString(priv)
			require.NoError(t, err)
			_, _, err = parse(t, token).VerifyWithKey(v2Key)
			require.ErrorIs(t, err, auth.ErrKeyIDMismatch)
		}
	})

	t.Run("v2 key refuses a token that is not of the v2 type", func(t *testing.T) {
		// jwt sets typ to JWT by default
		tok := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
			"iss": apiKey,
			"sub": "me",
			"exp": jwt.NewNumericDate(time.Now().Add(time.Hour)),
		})
		tok.Header["kid"] = apiKey
		token, err := tok.SignedString(priv)
		require.NoError(t, err)
		_, _, err = parse(t, token).VerifyWithKey(v2Key)
		require.ErrorIs(t, err, auth.ErrTokenTypeMismatch)
	})

	t.Run("expired EdDSA token is refused", func(t *testing.T) {
		token, err := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
			"iss": apiKey,
			"sub": "me",
			"exp": jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		}).SignedString(priv)
		require.NoError(t, err)
		_, _, err = parse(t, token).VerifyWithKey(v2Key)
		require.ErrorIs(t, err, jwt.ErrTokenExpired)
	})

	t.Run("v1 key with empty secret is refused", func(t *testing.T) {
		_, _, err := parse(t, hsToken).VerifyWithKey(auth.VerificationKey{Version: 1})
		require.ErrorIs(t, err, auth.ErrKeysMissing)
	})

	t.Run("v2 key with empty or wrong-size public key is refused", func(t *testing.T) {
		_, _, err := parse(t, edToken).VerifyWithKey(auth.VerificationKey{Version: 2})
		require.ErrorIs(t, err, auth.ErrKeysMissing)
		_, _, err = parse(t, edToken).VerifyWithKey(auth.VerificationKey{Version: 2, PublicKey: pub[:16]})
		require.ErrorIs(t, err, auth.ErrInvalidPublicKey)
	})

	t.Run("unknown version is refused", func(t *testing.T) {
		for _, version := range []int{0, 3} {
			_, _, err := parse(t, hsToken).VerifyWithKey(auth.VerificationKey{Version: version, Secret: secret, PublicKey: pub})
			require.ErrorIs(t, err, auth.ErrUnsupportedKeyVersion)
		}
	})

	t.Run("Verify refuses EdDSA", func(t *testing.T) {
		v := parse(t, edToken)
		_, _, err := v.Verify(secret)
		require.ErrorIs(t, err, auth.ErrKeyAlgorithmMismatch)
		_, _, err = v.Verify(pub)
		require.ErrorIs(t, err, auth.ErrKeyAlgorithmMismatch)
	})
}
