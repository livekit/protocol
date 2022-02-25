package auth_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gopkg.in/square/go-jose.v2/json"

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
		_, err = v.Verify("")
		require.Error(t, err)

		_, err = v.Verify("anothersecret")
		require.Error(t, err)
	})

	t.Run("key has expired", func(t *testing.T) {
		v, err := auth.ParseAPIToken(accessToken)
		require.NoError(t, err)

		_, err = v.Verify(secret)
		require.Error(t, err)
	})

	t.Run("unexpired token is verified", func(t *testing.T) {
		claim := auth.VideoGrant{RoomCreate: true}
		at := auth.NewAccessToken(apiKey, secret).
			AddGrant(&claim).
			SetValidFor(time.Minute).
			SetIdentity("me")
		authToken, err := at.ToJWT()
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(authToken)
		require.NoError(t, err)
		require.Equal(t, apiKey, v.APIKey())
		require.Equal(t, "me", v.Identity())

		decoded, err := v.Verify(secret)
		require.NoError(t, err)
		require.Equal(t, &claim, decoded.Video)
	})

	t.Run("ensure metadata can be passed through", func(t *testing.T) {
		metadata := map[string]interface{}{
			"user":   "value",
			"number": float64(3),
		}
		md, _ := json.Marshal(metadata)
		at := auth.NewAccessToken(apiKey, secret).
			AddGrant(&auth.VideoGrant{
				RoomAdmin: true,
				Room:      "myroom",
			}).
			SetMetadata(string(md))

		authToken, err := at.ToJWT()
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(authToken)
		require.NoError(t, err)

		decoded, err := v.Verify(secret)
		require.NoError(t, err)

		require.EqualValues(t, string(md), decoded.Metadata)
	})

	t.Run("nil permissions are handled", func(t *testing.T) {
		grant := &auth.VideoGrant{
			Room:     "myroom",
			RoomJoin: true,
		}
		grant.SetCanPublishData(false)
		at := auth.NewAccessToken(apiKey, secret).
			AddGrant(grant)
		token, err := at.ToJWT()
		require.NoError(t, err)

		v, err := auth.ParseAPIToken(token)
		require.NoError(t, err)
		decoded, err := v.Verify(secret)
		require.NoError(t, err)

		require.Nil(t, decoded.Video.CanSubscribe)
		require.Nil(t, decoded.Video.CanPublish)
		require.False(t, *decoded.Video.CanPublishData)
	})
}
