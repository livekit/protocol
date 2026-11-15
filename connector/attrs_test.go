package connector

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestTimeAttributes(t *testing.T) {
	cases := []struct {
		key    string
		format func(time.Time) map[string]string
		parse  func(map[string]string) time.Time
	}{
		{livekit.AttrConnectorRequestTime, RequestTimeAttributes, ParseRequestTime},
		{livekit.AttrConnectorAnswerTime, AnswerTimeAttributes, ParseAnswerTime},
	}
	at := time.UnixMilli(time.Now().UnixMilli())
	for _, c := range cases {
		t.Run(c.key, func(t *testing.T) {
			attrs := c.format(at)
			require.Contains(t, attrs, c.key)
			require.True(t, at.Equal(c.parse(attrs)))

			require.True(t, c.parse(nil).IsZero())
			require.True(t, c.parse(map[string]string{c.key: "bad"}).IsZero())
		})
	}
}
