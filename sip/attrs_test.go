package sip

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestInviteTimeAttributes(t *testing.T) {
	at := time.UnixMilli(time.Now().UnixMilli())
	require.True(t, at.Equal(ParseInviteTime(InviteTimeAttributes(at))))

	require.True(t, ParseInviteTime(nil).IsZero())
	require.True(t, ParseInviteTime(map[string]string{livekit.AttrSIPInviteTime: "bad"}).IsZero())
}
