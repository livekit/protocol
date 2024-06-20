package livekit

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSIPTrunkAs(t *testing.T) {
	t.Run("inbound", func(t *testing.T) {
		in := &SIPInboundTrunkInfo{
			SipTrunkId:       "id",
			Name:             "name",
			Metadata:         "{}",
			Numbers:          []string{"+1"},
			AllowedAddresses: []string{"sip.com"},
			AllowedNumbers:   []string{"+2", "+3"},
			AuthUsername:     "user",
			AuthPassword:     "pass",
		}
		got := in.AsTrunkInfo().AsInbound()
		require.Equal(t, in, got)
	})
	t.Run("outbound", func(t *testing.T) {
		out := &SIPOutboundTrunkInfo{
			SipTrunkId:   "id",
			Name:         "name",
			Metadata:     "{}",
			Address:      "sip.com",
			Transport:    1,
			Numbers:      []string{"+1"},
			AuthUsername: "user",
			AuthPassword: "pass",
		}
		got := out.AsTrunkInfo().AsOutbound()
		require.Equal(t, out, got)
	})
}
