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
	"strings"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils"
)

func TestAccessToken(t *testing.T) {
	t.Parallel()

	t.Run("keys must be set", func(t *testing.T) {
		token := NewAccessToken("", "")
		_, err := token.ToJWT()
		require.Equal(t, ErrKeysMissing, err)
	})

	t.Run("generates a decode-able key", func(t *testing.T) {
		apiKey, secret := apiKeypair()
		videoGrant := &VideoGrant{RoomJoin: true, Room: "myroom"}
		at := NewAccessToken(apiKey, secret).
			AddGrant(videoGrant).
			SetValidFor(time.Minute * 5).
			SetIdentity("user")
		value, err := at.ToJWT()
		//fmt.Println(raw)
		require.NoError(t, err)

		require.Len(t, strings.Split(value, "."), 3)

		// ensure it's a valid JWT
		token, err := jwt.ParseSigned(value)
		require.NoError(t, err)

		decodedGrant := ClaimGrants{}
		err = token.UnsafeClaimsWithoutVerification(&decodedGrant)
		require.NoError(t, err)

		require.EqualValues(t, videoGrant, decodedGrant.Video)
	})

	t.Run("default validity should be more than a minute", func(t *testing.T) {
		apiKey, secret := apiKeypair()
		videoGrant := &VideoGrant{RoomJoin: true, Room: "myroom"}
		at := NewAccessToken(apiKey, secret).
			AddGrant(videoGrant)
		value, err := at.ToJWT()
		require.NoError(t, err)
		token, err := jwt.ParseSigned(value)
		require.NoError(t, err)

		claim := jwt.Claims{}
		decodedGrant := ClaimGrants{}
		err = token.UnsafeClaimsWithoutVerification(&claim, &decodedGrant)
		require.NoError(t, err)
		require.EqualValues(t, videoGrant, decodedGrant.Video)

		// default validity
		require.True(t, claim.Expiry.Time().Sub(claim.IssuedAt.Time()) > time.Minute)
	})
}

func apiKeypair() (string, string) {
	return utils.NewGuid(utils.APIKeyPrefix), utils.RandomSecret()
}
