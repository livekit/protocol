package agent

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/auth"
)

func TestParseJobTokenClaims(t *testing.T) {
	t.Run("parses grants", func(t *testing.T) {
		grants, err := ParseJobTokenClaims(`{"identity":"avatar","kind":"agent","video":{"roomJoin":true,"canPublish":false},"attributes":{"a":"b"}}`)
		require.NoError(t, err)
		require.Equal(t, "avatar", grants.Identity)
		require.Equal(t, "agent", grants.Kind)
		require.True(t, grants.Video.RoomJoin)
		require.False(t, grants.Video.GetCanPublish())
		require.Equal(t, map[string]string{"a": "b"}, grants.Attributes)
	})

	t.Run("drops fields it does not know", func(t *testing.T) {
		grants, err := ParseJobTokenClaims(`{"identity":"a","foo":true,"exp":1,"video":{"roomJoin":true,"foo":true}}`)
		require.NoError(t, err)
		require.Equal(t, &auth.ClaimGrants{Identity: "a", Video: &auth.VideoGrant{RoomJoin: true}}, grants)

		// a whole grant the server does not know, and an unknown field inside
		// roomConfig, which has its own decoder
		grants, err = ParseJobTokenClaims(`{"identity":"a","newGrant":{"x":[1]},"roomConfig":{"name":"r","foo":true}}`)
		require.NoError(t, err)
		require.Equal(t, "a", grants.Identity)
		require.Equal(t, "r", grants.RoomConfig.Name)
	})

	refused := []struct {
		name   string
		claims string
	}{
		{"trailing data", `{"identity":"a"} {}`},
		{"not an object", `[]`},
		{"empty", ``},
	}
	for _, c := range refused {
		t.Run(c.name, func(t *testing.T) {
			_, err := ParseJobTokenClaims(c.claims)
			require.Error(t, err)
		})
	}

	neverIssued := []struct {
		name   string
		claims string
		err    string
	}{
		{"sha256", `{"sha256":"hash"}`, "claim not allowed: sha256"},
		{"sha256 in another case", `{"SHA256":"hash"}`, "claim not allowed: sha256"},
		{"video.agent", `{"video":{"roomJoin":true,"agent":true}}`, "claim not allowed: video.agent"},
		{"video.agent in another case", `{"Video":{"Agent":true}}`, "claim not allowed: video.agent"},
	}
	for _, c := range neverIssued {
		t.Run(c.name, func(t *testing.T) {
			_, err := ParseJobTokenClaims(c.claims)
			require.ErrorIs(t, err, ErrClaimNotAllowed)
			require.EqualError(t, err, c.err)
		})
	}
}

func TestBuildJobToken(t *testing.T) {
	grants := &auth.ClaimGrants{
		Identity: "avatar",
		Kind:     "agent",
		Video:    &auth.VideoGrant{RoomJoin: true, Room: "room"},
	}
	token, expiresAt, err := BuildJobToken("key", "secret", grants)
	require.NoError(t, err)
	require.WithinDuration(t, time.Now().Add(JobTokenValidFor), expiresAt, time.Second)

	v, err := auth.ParseAPIToken(token)
	require.NoError(t, err)
	require.Equal(t, "key", v.APIKey())
	claims, got, err := v.Verify("secret")
	require.NoError(t, err)
	require.Equal(t, grants, got)
	require.Equal(t, "avatar", claims.Subject)
	require.WithinDuration(t, expiresAt, claims.ExpiresAt.Time, time.Second)

	_, _, err = BuildJobToken("key", "", grants)
	require.ErrorIs(t, err, auth.ErrKeysMissing)

	_, _, err = BuildJobToken("key", "secret", nil)
	require.Error(t, err)
}

func TestBuildJobTokenRefusesClaimsNeverIssued(t *testing.T) {
	_, _, err := BuildJobToken("key", "secret", &auth.ClaimGrants{Sha256: "hash"})
	require.EqualError(t, err, "claim not allowed: sha256")

	_, _, err = BuildJobToken("key", "secret", &auth.ClaimGrants{Video: &auth.VideoGrant{Agent: true}})
	require.EqualError(t, err, "claim not allowed: video.agent")
}
