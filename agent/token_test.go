package agent

import (
	"testing"
	"time"

	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils/guid"
)

type testMetadata struct {
	Foo string
}

func TestToken(t *testing.T) {
	t.Run("validation", func(t *testing.T) {
		cases := []struct {
			label        string
			keys0, keys1 string
			success      bool
		}{
			{"empty key", "", "", true},
			{"one key", "foo", "foo", true},
			{"one of key set", "foo", "foo,bar", true},
			{"invalid key", "foo", "bar", false},
		}

		for _, c := range cases {
			t.Run(c.label, func(t *testing.T) {
				wt0 := NewWorkerTokenProvider(
					livekit.NodeID(guid.New(guid.NodePrefix)),
					WorkerTokenConfig{
						Keys:    c.keys0,
						Timeout: time.Hour,
					},
				)

				workerID := guid.New(guid.AgentWorkerPrefix)
				token, err := wt0.Encode(WorkerClaims{
					Claims: jwt.Claims{
						Subject: workerID,
					},
				})
				require.NoError(t, err)

				wt1 := NewWorkerTokenProvider(
					livekit.NodeID(guid.New(guid.NodePrefix)),
					WorkerTokenConfig{
						Keys:    c.keys1,
						Timeout: time.Hour,
					},
				)

				claims, err := wt1.Decode(token)
				if c.success {
					require.NoError(t, err)
					require.Equal(t, workerID, claims.Subject)
				} else {
					require.Error(t, err)
				}
			})
		}
	})

	t.Run("meta", func(t *testing.T) {
		wt := NewWorkerTokenProviderT[testMetadata](
			livekit.NodeID(guid.New(guid.NodePrefix)),
			WorkerTokenConfig{
				Keys:    "test",
				Timeout: time.Hour,
			},
		)

		token, err := wt.Encode(WorkerClaimsT[testMetadata]{
			Claims: jwt.Claims{
				Subject: guid.New(guid.AgentWorkerPrefix),
			},
			Metadata: testMetadata{
				Foo: "foo",
			},
		})
		require.NoError(t, err)

		claims, err := wt.Decode(token)
		require.NoError(t, err)
		require.Equal(t, "foo", claims.Metadata.Foo)
	})
}
