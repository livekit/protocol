package livekit

import (
	"testing"

	"github.com/stretchr/testify/require"
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

func TestSIPValidate(t *testing.T) {
	cases := []struct {
		name string
		req  interface {
			Validate() error
		}
		exp bool
	}{
		{
			name: "inbound empty",
			req:  &SIPInboundTrunkInfo{},
			exp:  false,
		},
		{
			name: "inbound numbers",
			req: &SIPInboundTrunkInfo{
				Numbers: []string{"+1111"},
			},
			exp: true,
		},
		{
			name: "inbound ips",
			req: &SIPInboundTrunkInfo{
				AllowedAddresses: []string{"1.1.1.1"},
			},
			exp: true,
		},
		{
			name: "inbound auth",
			req: &SIPInboundTrunkInfo{
				AuthUsername: "user",
				AuthPassword: "pass",
			},
			exp: true,
		},
		{
			name: "inbound x-header",
			req: &SIPInboundTrunkInfo{
				Numbers: []string{"+1111"},
				HeadersToAttributes: map[string]string{
					"X-Test": "test",
				},
			},
			exp: true,
		},
		{
			name: "inbound other header",
			req: &SIPInboundTrunkInfo{
				Numbers: []string{"+1111"},
				HeadersToAttributes: map[string]string{
					"From": "from",
				},
			},
			exp: false,
		},
		{
			name: "outbound empty",
			req:  &SIPOutboundTrunkInfo{},
			exp:  false,
		},
		{
			name: "outbound no numbers",
			req: &SIPOutboundTrunkInfo{
				Address: "sip.example.com",
			},
			exp: false,
		},
		{
			name: "outbound with numbers",
			req: &SIPOutboundTrunkInfo{
				Address: "sip.example.com",
				Numbers: []string{"+2222"},
			},
			exp: true,
		},
		{
			name: "outbound with user",
			req: &SIPOutboundTrunkInfo{
				Address: "user@sip.example.com",
				Numbers: []string{"+2222"},
			},
			exp: false,
		},
		{
			name: "outbound with transport",
			req: &SIPOutboundTrunkInfo{
				Address: "sip.example.com;transport=tcp",
				Numbers: []string{"+2222"},
			},
			exp: false,
		},
		{
			name: "outbound with schema",
			req: &SIPOutboundTrunkInfo{
				Address: "sip:example.com",
				Numbers: []string{"+2222"},
			},
			exp: false,
		},
		{
			name: "outbound with schema (tls)",
			req: &SIPOutboundTrunkInfo{
				Address: "sips:example.com",
				Numbers: []string{"+2222"},
			},
			exp: false,
		},
		{
			name: "outbound x-header",
			req: &SIPOutboundTrunkInfo{
				Address: "sip.example.com",
				Numbers: []string{"+2222"},
				HeadersToAttributes: map[string]string{
					"X-Test": "test",
				},
			},
			exp: true,
		},
		{
			name: "outbound other header",
			req: &SIPOutboundTrunkInfo{
				Address: "sip.example.com",
				Numbers: []string{"+2222"},
				HeadersToAttributes: map[string]string{
					"From": "from",
				},
			},
			exp: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.req.Validate()
			require.Equal(t, c.exp, err == nil, "error: %v", err)
		})
	}
}
