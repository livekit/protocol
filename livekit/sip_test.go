package livekit

import (
	"slices"
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
			exp: true,
		},
		{
			name: "inbound invalid header",
			req: &SIPInboundTrunkInfo{
				Numbers: []string{"+1111"},
				HeadersToAttributes: map[string]string{
					"From ": "from",
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
			exp: true,
		},
		{
			name: "outbound invalid header",
			req: &SIPOutboundTrunkInfo{
				Address: "sip.example.com",
				Numbers: []string{"+2222"},
				HeadersToAttributes: map[string]string{
					"From ": "from",
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

func TestSIPInboundTrunkFilter(t *testing.T) {
	list := []*SIPInboundTrunkInfo{
		0: {SipTrunkId: "A"},
		1: {SipTrunkId: "B", Numbers: []string{"+111", "+222"}},
		2: {SipTrunkId: "C", Numbers: []string{"+333"}},
	}
	cases := []struct {
		name string
		req  *ListSIPInboundTrunkRequest
		exp  []*SIPInboundTrunkInfo
	}{
		{
			name: "all",
			req:  &ListSIPInboundTrunkRequest{},
			exp:  list,
		},
		{
			name: "id",
			req: &ListSIPInboundTrunkRequest{
				TrunkIds: []string{"B"},
			},
			exp: []*SIPInboundTrunkInfo{list[1]},
		},
		{
			name: "ids",
			req: &ListSIPInboundTrunkRequest{
				TrunkIds: []string{"B", "C"},
			},
			exp: []*SIPInboundTrunkInfo{list[1], list[2]},
		},
		{
			name: "ids order",
			req: &ListSIPInboundTrunkRequest{
				TrunkIds: []string{"C", "missing", "B"},
			},
			exp: []*SIPInboundTrunkInfo{list[2], nil, list[1]},
		},
		{
			name: "numbers",
			req: &ListSIPInboundTrunkRequest{
				Numbers: []string{"+222"},
			},
			exp: []*SIPInboundTrunkInfo{list[0], list[1]},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := slices.Clone(list)
			got := c.req.FilterSlice(list)
			require.Equal(t, c.exp, got)
		})
	}
}

func TestSIPOutboundTrunkFilter(t *testing.T) {
	list := []*SIPOutboundTrunkInfo{
		0: {SipTrunkId: "A"},
		1: {SipTrunkId: "B", Numbers: []string{"+111", "+222"}},
		2: {SipTrunkId: "C", Numbers: []string{"+333"}},
	}
	cases := []struct {
		name string
		req  *ListSIPOutboundTrunkRequest
		exp  []*SIPOutboundTrunkInfo
	}{
		{
			name: "all",
			req:  &ListSIPOutboundTrunkRequest{},
			exp:  list,
		},
		{
			name: "id",
			req: &ListSIPOutboundTrunkRequest{
				TrunkIds: []string{"B"},
			},
			exp: []*SIPOutboundTrunkInfo{list[1]},
		},
		{
			name: "ids",
			req: &ListSIPOutboundTrunkRequest{
				TrunkIds: []string{"B", "C"},
			},
			exp: []*SIPOutboundTrunkInfo{list[1], list[2]},
		},
		{
			name: "ids order",
			req: &ListSIPOutboundTrunkRequest{
				TrunkIds: []string{"C", "missing", "B"},
			},
			exp: []*SIPOutboundTrunkInfo{list[2], nil, list[1]},
		},
		{
			name: "numbers",
			req: &ListSIPOutboundTrunkRequest{
				Numbers: []string{"+222"},
			},
			exp: []*SIPOutboundTrunkInfo{list[0], list[1]},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := slices.Clone(list)
			got := c.req.FilterSlice(list)
			require.Equal(t, c.exp, got)
		})
	}
}

func TestSIPDispatchRuleFilter(t *testing.T) {
	list := []*SIPDispatchRuleInfo{
		0: {SipDispatchRuleId: "A"},
		1: {SipDispatchRuleId: "B", TrunkIds: []string{"T1", "T2"}},
		2: {SipDispatchRuleId: "C", TrunkIds: []string{"T3"}},
	}
	cases := []struct {
		name string
		req  *ListSIPDispatchRuleRequest
		exp  []*SIPDispatchRuleInfo
	}{
		{
			name: "all",
			req:  &ListSIPDispatchRuleRequest{},
			exp:  list,
		},
		{
			name: "id",
			req: &ListSIPDispatchRuleRequest{
				DispatchRuleIds: []string{"B"},
			},
			exp: []*SIPDispatchRuleInfo{list[1]},
		},
		{
			name: "ids",
			req: &ListSIPDispatchRuleRequest{
				DispatchRuleIds: []string{"B", "C"},
			},
			exp: []*SIPDispatchRuleInfo{list[1], list[2]},
		},
		{
			name: "ids order",
			req: &ListSIPDispatchRuleRequest{
				DispatchRuleIds: []string{"C", "missing", "B"},
			},
			exp: []*SIPDispatchRuleInfo{list[2], nil, list[1]},
		},
		{
			name: "numbers",
			req: &ListSIPDispatchRuleRequest{
				TrunkIds: []string{"T2"},
			},
			exp: []*SIPDispatchRuleInfo{list[0], list[1]},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := slices.Clone(list)
			got := c.req.FilterSlice(list)
			require.Equal(t, c.exp, got)
		})
	}
}
