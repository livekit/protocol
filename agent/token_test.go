package agent

import (
	"testing"
	"time"

	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils/guid"
)

func TestToken(t *testing.T) {
	cases := []struct {
		label        string
		keys0, keys1 string
		success      bool
	}{
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
}
