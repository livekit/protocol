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

package sip

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/netip"
	"strconv"
	"testing"

	"github.com/dennwc/iters"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/rpc"
)

func TestNormalizeNumber(t *testing.T) {
	cases := []struct {
		name string
		num  string
		exp  string
	}{
		{"empty", "", ""},
		{"number", "123", "+123"},
		{"plus", "+123", "+123"},
		{"user", "user", "user"},
		{"human", "(123) 456 7890", "+1234567890"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.exp, NormalizeNumber(c.num))
		})
	}
}

const (
	sipNumber1  = "1111 1111"
	sipNumber2  = "2222 2222"
	sipNumber3  = "3333 3333"
	sipTrunkID1 = "aaa"
	sipTrunkID2 = "bbb"
)

var trunkCases = []struct {
	name    string
	trunks  []*livekit.SIPTrunkInfo
	exp     int
	expErr  bool
	invalid bool
	from    string
	to      string
	src     string
	host    string
}{
	{
		name:   "empty",
		trunks: nil,
		exp:    -1, // no error; nil result
	},
	{
		name: "one wildcard",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa"},
		},
		exp: 0,
	},
	{
		name: "matching",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber2},
		},
		exp: 0,
	},
	{
		name: "matching inbound",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber2, InboundNumbers: []string{sipNumber1}},
		},
		exp: 0,
	},
	{
		name: "matching regexp",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber2, InboundNumbersRegex: []string{`^\d+ \d+$`}},
		},
		exp: 0,
	},
	{
		name: "not matching",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber3},
		},
		exp: -1,
	},
	{
		name: "not matching inbound",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber2, InboundNumbers: []string{sipNumber1 + "1"}},
		},
		exp: -1,
	},
	{
		name: "one match",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber3},
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2},
		},
		exp: 1,
	},
	{
		name: "many matches",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber3},
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2},
			{SipTrunkId: "ccc", OutboundNumber: sipNumber2},
		},
		expErr:  true,
		invalid: true,
	},
	{
		name: "many matches default",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber3},
			{SipTrunkId: "bbb"},
			{SipTrunkId: "ccc", OutboundNumber: sipNumber2},
			{SipTrunkId: "ddd"},
		},
		exp:     2,
		invalid: true, // it can successfully select "ccc", but the overall configuration is invalid
	},
	{
		name: "inbound",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber3},
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2},
			{SipTrunkId: "ccc", OutboundNumber: sipNumber2, InboundNumbers: []string{sipNumber1 + "1"}},
		},
		exp: 1,
	},
	{
		name: "multiple defaults",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber3},
			{SipTrunkId: "bbb"},
			{SipTrunkId: "ccc"},
		},
		expErr:  true,
		invalid: true,
	},
	{
		name: "inbound with ip exact",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2, InboundAddresses: []string{
				"10.10.10.10",
				"1.1.1.1",
			}},
		},
		exp: 0,
	},
	{
		name: "inbound with ip exact miss",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2, InboundAddresses: []string{
				"10.10.10.10",
			}},
		},
		exp: -1,
	},
	{
		name: "inbound with ip mask",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2, InboundAddresses: []string{
				"10.10.10.0/24",
				"1.1.1.0/24",
			}},
		},
		exp: 0,
	},
	{
		name: "inbound with ip mask miss",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2, InboundAddresses: []string{
				"10.10.10.0/24",
			}},
		},
		exp: -1,
	},
	{
		name: "inbound with host mask",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2, InboundAddresses: []string{
				"10.10.10.0/24",
				"sip.example.com",
			}},
		},
		exp: 0,
	},
	{
		name: "inbound with plus",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: "+" + sipNumber3},
			{SipTrunkId: "bbb", OutboundNumber: "+" + sipNumber2},
		},
		exp: 1,
	},
	{
		name: "inbound without plus",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber3},
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2},
		},
		from: "+" + sipNumber1,
		to:   "+" + sipNumber2,
		exp:  1,
	},
}

func toInboundTrunks(trunks []*livekit.SIPTrunkInfo) []*livekit.SIPInboundTrunkInfo {
	out := make([]*livekit.SIPInboundTrunkInfo, 0, len(trunks))
	for _, t := range trunks {
		out = append(out, t.AsInbound())
	}
	return out
}

func TestSIPMatchTrunk(t *testing.T) {
	for _, c := range trunkCases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			from, to, src, host := c.from, c.to, c.src, c.host
			if from == "" {
				from = sipNumber1
			}
			if to == "" {
				to = sipNumber2
			}
			if src == "" {
				src = "1.1.1.1"
			}
			if host == "" {
				host = "sip.example.com"
			}
			trunks := toInboundTrunks(c.trunks)
			call := &rpc.SIPCall{
				SipCallId: "test-call-id",
				SourceIp:  src,
				From: &livekit.SIPUri{
					User: from,
					Host: host,
				},
				To: &livekit.SIPUri{
					User: to,
				},
			}
			call.Address = call.To
			got, err := MatchTrunkIter(iters.Slice(trunks), call, WithTrunkConflict(func(t1, t2 *livekit.SIPInboundTrunkInfo, reason TrunkConflictReason) {
				t.Logf("conflict: %v\n%v\nvs\n%v", reason, t1, t2)
			}))
			if c.expErr {
				require.Error(t, err)
				require.Nil(t, got)
				t.Log(err)
			} else {
				var exp *livekit.SIPInboundTrunkInfo
				if c.exp >= 0 {
					exp = trunks[c.exp]
				}
				require.NoError(t, err)
				require.Equal(t, exp, got)
			}
		})
	}
}

func TestSIPValidateTrunks(t *testing.T) {
	for _, c := range trunkCases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			for i, r := range c.trunks {
				if r.SipTrunkId == "" {
					r.SipTrunkId = strconv.Itoa(i)
				}
			}
			err := ValidateTrunks(toInboundTrunks(c.trunks))
			if c.invalid {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func newSIPTrunkDispatch() *livekit.SIPTrunkInfo {
	return &livekit.SIPTrunkInfo{
		SipTrunkId:     sipTrunkID1,
		OutboundNumber: sipNumber2,
	}
}

func newSIPReqDispatch(pin string, noPin bool) *rpc.EvaluateSIPDispatchRulesRequest {
	return &rpc.EvaluateSIPDispatchRulesRequest{
		CallingNumber: sipNumber1,
		CalledNumber:  sipNumber2,
		Pin:           pin,
		//NoPin: noPin, // TODO
	}
}

func newDirectDispatch(room, pin string) *livekit.SIPDispatchRule {
	return &livekit.SIPDispatchRule{
		Rule: &livekit.SIPDispatchRule_DispatchRuleDirect{
			DispatchRuleDirect: &livekit.SIPDispatchRuleDirect{
				RoomName: room, Pin: pin,
			},
		},
	}
}

func newIndividualDispatch(roomPref, pin string) *livekit.SIPDispatchRule {
	return &livekit.SIPDispatchRule{
		Rule: &livekit.SIPDispatchRule_DispatchRuleIndividual{
			DispatchRuleIndividual: &livekit.SIPDispatchRuleIndividual{
				RoomPrefix: roomPref, Pin: pin,
			},
		},
	}
}

var dispatchCases = []struct {
	name    string
	trunk   *livekit.SIPTrunkInfo
	rules   []*livekit.SIPDispatchRuleInfo
	reqPin  string
	noPin   bool
	exp     int
	expErr  bool
	invalid bool
}{
	// These cases just validate that no rules produce an error.
	{
		name:   "empty",
		trunk:  nil,
		rules:  nil,
		expErr: true,
	},
	{
		name:   "only trunk",
		trunk:  newSIPTrunkDispatch(),
		rules:  nil,
		expErr: true,
	},
	// Default rules should work even if no trunk is defined.
	{
		name:  "one rule/no trunk",
		trunk: nil,
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip", "")},
		},
		exp: 0,
	},
	// Default rule should work with a trunk too.
	{
		name:  "one rule/default trunk",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip", "")},
		},
		exp: 0,
	},
	// Rule matching the trunk should be selected.
	{
		name:  "one rule/specific trunk",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: []string{sipTrunkID1, sipTrunkID2}, Rule: newDirectDispatch("sip", "")},
		},
		exp: 0,
	},
	// Rule NOT matching the trunk should NOT be selected.
	{
		name:  "one rule/wrong trunk",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: []string{"zzz"}, Rule: newDirectDispatch("sip", "")},
		},
		expErr: true,
	},
	// Direct rule with a pin should be selected, even if no pin is provided.
	{
		name:  "direct pin/correct",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: []string{sipTrunkID1}, Rule: newDirectDispatch("sip", "123")},
			{TrunkIds: []string{sipTrunkID2}, Rule: newDirectDispatch("sip", "456")},
		},
		reqPin: "123",
		exp:    0,
	},
	// Direct rule with a pin should reject wrong pin.
	{
		name:  "direct pin/wrong",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: []string{sipTrunkID1}, Rule: newDirectDispatch("sip", "123")},
			{TrunkIds: []string{sipTrunkID2}, Rule: newDirectDispatch("sip", "456")},
		},
		reqPin: "zzz",
		expErr: true,
	},
	// Multiple direct rules with the same pin should result in an error.
	{
		name:  "direct pin/conflict",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: []string{sipTrunkID1}, Rule: newDirectDispatch("sip1", "123")},
			{TrunkIds: []string{sipTrunkID1, sipTrunkID2}, Rule: newDirectDispatch("sip2", "123")},
		},
		reqPin:  "123",
		expErr:  true,
		invalid: true,
	},
	// Multiple direct rules with the same pin on different trunks are ok.
	{
		name:  "direct pin/no conflict on different trunk",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: []string{sipTrunkID1}, Rule: newDirectDispatch("sip1", "123")},
			{TrunkIds: []string{sipTrunkID2}, Rule: newDirectDispatch("sip2", "123")},
		},
		reqPin: "123",
		exp:    0,
	},
	// Specific direct rules should take priority over default direct rules.
	{
		name:  "direct pin/default and specific",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", "123")},
			{TrunkIds: []string{sipTrunkID1}, Rule: newDirectDispatch("sip2", "123")},
		},
		reqPin: "123",
		exp:    1,
	},
	// Specific direct rules should take priority over default direct rules. No pin.
	{
		name:  "direct/default and specific",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", "")},
			{TrunkIds: []string{sipTrunkID1}, Rule: newDirectDispatch("sip2", "")},
		},
		exp: 1,
	},
	// Specific direct rules should take priority over default direct rules. One with pin, other without.
	{
		name:  "direct/default and specific/mixed 1",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", "123")},
			{TrunkIds: []string{sipTrunkID1}, Rule: newDirectDispatch("sip2", "")},
		},
		exp: 1,
	},
	{
		name:  "direct/default and specific/mixed 2",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", "")},
			{TrunkIds: []string{sipTrunkID1}, Rule: newDirectDispatch("sip2", "123")},
		},
		exp: 1,
	},
	// Multiple default direct rules are not allowed.
	{
		name:  "direct/multiple defaults",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", "")},
			{TrunkIds: nil, Rule: newDirectDispatch("sip2", "")},
		},
		expErr:  true,
		invalid: true,
	},
	// Rules for specific numbers take priority.
	{
		name:  "direct/number specific",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", "")},
			{TrunkIds: nil, Rule: newDirectDispatch("sip2", ""), InboundNumbers: []string{sipNumber1}},
		},
		exp: 1,
	},
	{
		name:  "direct/number specific pin",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", "123")},
			{TrunkIds: nil, Rule: newDirectDispatch("sip2", "123"), InboundNumbers: []string{sipNumber1}},
		},
		exp: 1,
	},
	{
		name:  "direct/number specific conflict",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", ""), InboundNumbers: []string{sipNumber1}},
			{TrunkIds: nil, Rule: newDirectDispatch("sip2", ""), InboundNumbers: []string{sipNumber1, sipNumber2}},
		},
		expErr:  true,
		invalid: true,
	},
	// Check the "personal room" use case. Rule that accepts a number without a pin and requires pin for everyone else.
	{
		name:  "direct/open specific vs pin generic",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newDirectDispatch("sip1", "123")},
			{TrunkIds: nil, Rule: newDirectDispatch("sip2", ""), InboundNumbers: []string{sipNumber1}},
		},
		exp: 1,
	},
	// Cannot use both direct and individual rules with the same pin setup.
	{
		name:  "direct vs individual/private",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newIndividualDispatch("pref_", "123")},
			{TrunkIds: nil, Rule: newDirectDispatch("sip", "123")},
		},
		expErr:  true,
		invalid: true,
	},
	{
		name:  "direct vs individual/open",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newIndividualDispatch("pref_", "")},
			{TrunkIds: nil, Rule: newDirectDispatch("sip", "")},
		},
		expErr:  true,
		invalid: true,
	},
	// Direct rules take priority over individual rules.
	{
		name:  "direct vs individual/priority",
		trunk: newSIPTrunkDispatch(),
		rules: []*livekit.SIPDispatchRuleInfo{
			{TrunkIds: nil, Rule: newIndividualDispatch("pref_", "123")},
			{TrunkIds: nil, Rule: newDirectDispatch("sip", "456")},
		},
		reqPin: "456",
		exp:    1,
	},
}

func TestSIPMatchDispatchRule(t *testing.T) {
	for _, c := range dispatchCases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			pins := []string{c.reqPin}
			if !c.expErr && c.reqPin != "" {
				// Should match the same rule, even if no pin is set (so that it can be requested).
				pins = append(pins, "")
			}
			for i, r := range c.rules {
				if r.SipDispatchRuleId == "" {
					r.SipDispatchRuleId = fmt.Sprintf("rule_%d", i)
				}
			}
			for _, pin := range pins {
				pin := pin
				name := pin
				if name == "" {
					name = "no pin"
				}
				t.Run(name, func(t *testing.T) {
					got, err := MatchDispatchRuleIter(c.trunk.AsInbound(), iters.Slice(c.rules), newSIPReqDispatch(pin, c.noPin), WithDispatchRuleConflict(func(r1, r2 *livekit.SIPDispatchRuleInfo, reason DispatchRuleConflictReason) {
						t.Logf("conflict: %v\n%v\nvs\n%v", reason, r1, r2)
					}))
					if c.expErr {
						require.Error(t, err)
						require.Nil(t, got)
						t.Log(err)
					} else {
						var exp *livekit.SIPDispatchRuleInfo
						if c.exp >= 0 {
							exp = c.rules[c.exp]
						}
						require.NoError(t, err)
						require.Equal(t, exp, got)
					}
				})
			}
		})
	}
}

func TestSIPValidateDispatchRules(t *testing.T) {
	for _, c := range dispatchCases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			for i, r := range c.rules {
				if r.SipDispatchRuleId == "" {
					r.SipDispatchRuleId = strconv.Itoa(i)
				}
			}
			_, err := ValidateDispatchRulesIter(iters.Slice(c.rules), WithDispatchRuleConflict(func(r1, r2 *livekit.SIPDispatchRuleInfo, reason DispatchRuleConflictReason) {
				t.Logf("conflict: %v\n%v\nvs\n%v", reason, r1, r2)
			}))
			if c.invalid {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestEvaluateDispatchRule(t *testing.T) {
	d := &livekit.SIPDispatchRuleInfo{
		SipDispatchRuleId: "rule",
		Rule:              newDirectDispatch("room", ""),
		HidePhoneNumber:   false,
		InboundNumbers:    nil,
		Name:              "",
		Metadata:          "rule-meta",
		Attributes: map[string]string{
			"rule-attr": "1",
		},
	}
	r := &rpc.EvaluateSIPDispatchRulesRequest{
		SipCallId:     "call-id",
		CallingNumber: "+11112222",
		CallingHost:   "sip.example.com",
		CalledNumber:  "+3333",
		ExtraAttributes: map[string]string{
			"prov-attr": "1",
		},
	}
	tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}
	res, err := EvaluateDispatchRule("p_123", tr, d, r)
	require.NoError(t, err)
	require.Equal(t, &rpc.EvaluateSIPDispatchRulesResponse{
		ProjectId:           "p_123",
		Result:              rpc.SIPDispatchResult_ACCEPT,
		SipTrunkId:          "trunk",
		SipDispatchRuleId:   "rule",
		RoomName:            "room",
		ParticipantIdentity: "sip_+11112222",
		ParticipantName:     "Phone +11112222",
		ParticipantMetadata: "rule-meta",
		ParticipantAttributes: map[string]string{
			"rule-attr":                   "1",
			"prov-attr":                   "1",
			livekit.AttrSIPCallID:         "call-id",
			livekit.AttrSIPTrunkID:        "trunk",
			livekit.AttrSIPDispatchRuleID: "rule",
			livekit.AttrSIPPhoneNumber:    "+11112222",
			livekit.AttrSIPTrunkNumber:    "+3333",
			livekit.AttrSIPHostName:       "sip.example.com",
		},
	}, res)

	d.HidePhoneNumber = true
	res, err = EvaluateDispatchRule("p_123", tr, d, r)
	require.NoError(t, err)
	require.Equal(t, &rpc.EvaluateSIPDispatchRulesResponse{
		ProjectId:           "p_123",
		Result:              rpc.SIPDispatchResult_ACCEPT,
		SipTrunkId:          "trunk",
		SipDispatchRuleId:   "rule",
		RoomName:            "room",
		ParticipantIdentity: "sip_c15a31c71649a522",
		ParticipantName:     "Phone 2222",
		ParticipantMetadata: "rule-meta",
		ParticipantAttributes: map[string]string{
			"rule-attr":                   "1",
			"prov-attr":                   "1",
			livekit.AttrSIPCallID:         "call-id",
			livekit.AttrSIPTrunkID:        "trunk",
			livekit.AttrSIPDispatchRuleID: "rule",
		},
	}, res)
}

func TestEvaluateDispatchRuleDynamic(t *testing.T) {
	// Test dynamic dispatch rule with accept action
	t.Run("dynamic_accept", func(t *testing.T) {
		// Create a test server to mock the webhook
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "POST", r.Method)
			require.Equal(t, "application/json", r.Header.Get("Content-Type"))

			// Verify request body
			var reqBody map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&reqBody)
			require.NoError(t, err)

			require.Equal(t, "p_123", reqBody["project_id"])
			require.Equal(t, "trunk", reqBody["trunk_id"])
			require.Equal(t, "call-id", reqBody["call_id"])
			require.Equal(t, "+11112222", reqBody["from_user"])
			require.Equal(t, "sip.example.com", reqBody["from_host"])
			require.Equal(t, "+3333", reqBody["to_user"])
			require.Equal(t, "sip.example.com", reqBody["to_host"])

			// Return accept response
			response := map[string]interface{}{
				"dispatch_action": "accept",
				"room_name":       "dynamic-room",
				"pin":             "1234",
				"metadata":        "dynamic-meta",
				"attributes": map[string]string{
					"dynamic-attr": "value",
				},
				"media_encryption": "SRTP",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		d := &livekit.SIPDispatchRuleInfo{
			SipDispatchRuleId: "dynamic-rule",
			Rule: &livekit.SIPDispatchRule{
				Rule: &livekit.SIPDispatchRule_DispatchRuleDynamic{
					DispatchRuleDynamic: &livekit.SIPDispatchRuleDynamic{
						Url: server.URL,
					},
				},
			},
			Attributes: map[string]string{
				"rule-attr": "1",
			},
		}

		r := &rpc.EvaluateSIPDispatchRulesRequest{
			SipCallId:     "call-id",
			CallingNumber: "+11112222",
			CallingHost:   "sip.example.com",
			CalledNumber:  "+3333",
			CalledHost:    "sip.example.com",
			Pin:           "1234",
			ExtraAttributes: map[string]string{
				"prov-attr": "1",
			},
		}

		tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}

		res, err := EvaluateDispatchRule("p_123", tr, d, r)
		require.NoError(t, err)
		require.Equal(t, &rpc.EvaluateSIPDispatchRulesResponse{
			ProjectId:           "p_123",
			Result:              rpc.SIPDispatchResult_ACCEPT,
			SipTrunkId:          "trunk",
			SipDispatchRuleId:   "dynamic-rule",
			RoomName:            "dynamic-room",
			ParticipantIdentity: "sip_+11112222",
			ParticipantName:     "Phone +11112222",
			ParticipantMetadata: "dynamic-meta",
			ParticipantAttributes: map[string]string{
				"rule-attr":                   "1",
				"prov-attr":                   "1",
				"dynamic-attr":                "value",
				livekit.AttrSIPCallID:         "call-id",
				livekit.AttrSIPTrunkID:        "trunk",
				livekit.AttrSIPDispatchRuleID: "dynamic-rule",
				livekit.AttrSIPPhoneNumber:    "+11112222",
				livekit.AttrSIPTrunkNumber:    "+3333",
				livekit.AttrSIPHostName:       "sip.example.com",
			},
			MediaEncryption: livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE,
		}, res)
	})

	// Test dynamic dispatch rule with reject action
	t.Run("dynamic_reject", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := map[string]interface{}{
				"dispatch_action": "reject",
				"reject_reason":   "Call not allowed",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		d := &livekit.SIPDispatchRuleInfo{
			SipDispatchRuleId: "dynamic-rule",
			Rule: &livekit.SIPDispatchRule{
				Rule: &livekit.SIPDispatchRule_DispatchRuleDynamic{
					DispatchRuleDynamic: &livekit.SIPDispatchRuleDynamic{
						Url: server.URL,
					},
				},
			},
		}

		r := &rpc.EvaluateSIPDispatchRulesRequest{
			SipCallId:     "call-id",
			CallingNumber: "+11112222",
			CallingHost:   "sip.example.com",
			CalledNumber:  "+3333",
			CalledHost:    "sip.example.com",
		}

		tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}

		res, err := EvaluateDispatchRule("p_123", tr, d, r)
		require.NoError(t, err)
		require.Equal(t, &rpc.EvaluateSIPDispatchRulesResponse{
			ProjectId:           "p_123",
			SipTrunkId:          "trunk",
			SipDispatchRuleId:   "dynamic-rule",
			Result:              rpc.SIPDispatchResult_REJECT,
			ParticipantMetadata: "Call not allowed",
		}, res)
	})

	// Test dynamic dispatch rule with transfer action
	t.Run("dynamic_transfer", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := map[string]interface{}{
				"dispatch_action":          "transfer",
				"transfer_to":              "+55556666",
				"destination_name":         "Support Desk",
				"transfer_ringing_timeout": 30,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		d := &livekit.SIPDispatchRuleInfo{
			SipDispatchRuleId: "dynamic-rule",
			Rule: &livekit.SIPDispatchRule{
				Rule: &livekit.SIPDispatchRule_DispatchRuleDynamic{
					DispatchRuleDynamic: &livekit.SIPDispatchRuleDynamic{
						Url: server.URL,
					},
				},
			},
		}

		r := &rpc.EvaluateSIPDispatchRulesRequest{
			SipCallId:     "call-id",
			CallingNumber: "+11112222",
			CallingHost:   "sip.example.com",
			CalledNumber:  "+3333",
			CalledHost:    "sip.example.com",
		}

		tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}

		res, err := EvaluateDispatchRule("p_123", tr, d, r)
		require.NoError(t, err)
		require.Equal(t, &rpc.EvaluateSIPDispatchRulesResponse{
			ProjectId:           "p_123",
			SipTrunkId:          "trunk",
			SipDispatchRuleId:   "dynamic-rule",
			Result:              rpc.SIPDispatchResult_DROP,
			ParticipantMetadata: "+55556666",
		}, res)
	})

	// Test webhook HTTP error
	t.Run("webhook_http_error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		d := &livekit.SIPDispatchRuleInfo{
			SipDispatchRuleId: "dynamic-rule",
			Rule: &livekit.SIPDispatchRule{
				Rule: &livekit.SIPDispatchRule_DispatchRuleDynamic{
					DispatchRuleDynamic: &livekit.SIPDispatchRuleDynamic{
						Url: server.URL,
					},
				},
			},
		}

		r := &rpc.EvaluateSIPDispatchRulesRequest{
			SipCallId:     "call-id",
			CallingNumber: "+11112222",
			CallingHost:   "sip.example.com",
			CalledNumber:  "+3333",
			CalledHost:    "sip.example.com",
		}

		tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}

		_, err := EvaluateDispatchRule("p_123", tr, d, r)
		require.Error(t, err)
		require.Contains(t, err.Error(), "Webhook returned non-200 status: 500")
	})

	// Test webhook network error
	t.Run("webhook_network_error", func(t *testing.T) {
		d := &livekit.SIPDispatchRuleInfo{
			SipDispatchRuleId: "dynamic-rule",
			Rule: &livekit.SIPDispatchRule{
				Rule: &livekit.SIPDispatchRule_DispatchRuleDynamic{
					DispatchRuleDynamic: &livekit.SIPDispatchRuleDynamic{
						Url: "http://invalid-url-that-does-not-exist.com/webhook",
					},
				},
			},
		}

		r := &rpc.EvaluateSIPDispatchRulesRequest{
			SipCallId:     "call-id",
			CallingNumber: "+11112222",
			CallingHost:   "sip.example.com",
			CalledNumber:  "+3333",
			CalledHost:    "sip.example.com",
		}

		tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}

		_, err := EvaluateDispatchRule("p_123", tr, d, r)
		require.Error(t, err)
		require.Contains(t, err.Error(), "Failed to make HTTP request")
	})

	// Test invalid JSON response
	t.Run("invalid_json_response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("invalid json"))
		}))
		defer server.Close()

		d := &livekit.SIPDispatchRuleInfo{
			SipDispatchRuleId: "dynamic-rule",
			Rule: &livekit.SIPDispatchRule{
				Rule: &livekit.SIPDispatchRule_DispatchRuleDynamic{
					DispatchRuleDynamic: &livekit.SIPDispatchRuleDynamic{
						Url: server.URL,
					},
				},
			},
		}

		r := &rpc.EvaluateSIPDispatchRulesRequest{
			SipCallId:     "call-id",
			CallingNumber: "+11112222",
			CallingHost:   "sip.example.com",
			CalledNumber:  "+3333",
			CalledHost:    "sip.example.com",
		}

		tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}

		_, err := EvaluateDispatchRule("p_123", tr, d, r)
		require.Error(t, err)
		require.Contains(t, err.Error(), "Failed to decode response")
	})

	// Test unknown dispatch action
	t.Run("unknown_dispatch_action", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := map[string]interface{}{
				"dispatch_action": "unknown_action",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		d := &livekit.SIPDispatchRuleInfo{
			SipDispatchRuleId: "dynamic-rule",
			Rule: &livekit.SIPDispatchRule{
				Rule: &livekit.SIPDispatchRule_DispatchRuleDynamic{
					DispatchRuleDynamic: &livekit.SIPDispatchRuleDynamic{
						Url: server.URL,
					},
				},
			},
		}

		r := &rpc.EvaluateSIPDispatchRulesRequest{
			SipCallId:     "call-id",
			CallingNumber: "+11112222",
			CallingHost:   "sip.example.com",
			CalledNumber:  "+3333",
			CalledHost:    "sip.example.com",
		}

		tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}

		// Should continue with normal dispatch rule evaluation
		res, err := EvaluateDispatchRule("p_123", tr, d, r)
		require.NoError(t, err)
		require.Equal(t, rpc.SIPDispatchResult_ACCEPT, res.Result)
	})

	// Test dynamic dispatch with PIN validation
	t.Run("dynamic_with_pin_validation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := map[string]interface{}{
				"dispatch_action": "accept",
				"room_name":       "pin-room",
				"pin":             "5678",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		d := &livekit.SIPDispatchRuleInfo{
			SipDispatchRuleId: "dynamic-rule",
			Rule: &livekit.SIPDispatchRule{
				Rule: &livekit.SIPDispatchRule_DispatchRuleDynamic{
					DispatchRuleDynamic: &livekit.SIPDispatchRuleDynamic{
						Url: server.URL,
					},
				},
			},
		}

		// Test without PIN - should request PIN
		r := &rpc.EvaluateSIPDispatchRulesRequest{
			SipCallId:     "call-id",
			CallingNumber: "+11112222",
			CallingHost:   "sip.example.com",
			CalledNumber:  "+3333",
			CalledHost:    "sip.example.com",
		}

		tr := &livekit.SIPInboundTrunkInfo{SipTrunkId: "trunk"}

		res, err := EvaluateDispatchRule("p_123", tr, d, r)
		require.NoError(t, err)
		require.Equal(t, rpc.SIPDispatchResult_REQUEST_PIN, res.Result)
		require.True(t, res.RequestPin)

		// Test with correct PIN
		r.Pin = "5678"
		res, err = EvaluateDispatchRule("p_123", tr, d, r)
		require.NoError(t, err)
		require.Equal(t, rpc.SIPDispatchResult_ACCEPT, res.Result)
		require.Equal(t, "pin-room", res.RoomName)

		// Test with incorrect PIN - should return permission denied error
		r.Pin = "9999"
		_, err = EvaluateDispatchRule("p_123", tr, d, r)
		require.Error(t, err)
		require.Contains(t, err.Error(), "Incorrect PIN for SIP room")
	})
}

func TestMatchIP(t *testing.T) {
	cases := []struct {
		addr  string
		mask  string
		valid bool
		exp   bool
	}{
		{addr: "192.168.0.10", mask: "192.168.0.10", valid: true, exp: true},
		{addr: "192.168.0.10", mask: "192.168.0.11", valid: true, exp: false},
		{addr: "192.168.0.10", mask: "192.168.0.0/24", valid: true, exp: true},
		{addr: "192.168.0.10", mask: "192.168.0.10/0", valid: true, exp: true},
		{addr: "192.168.0.10", mask: "192.170.0.0/24", valid: true, exp: false},
	}
	for _, c := range cases {
		t.Run(c.mask, func(t *testing.T) {
			ip, err := netip.ParseAddr(c.addr)
			require.NoError(t, err)
			got := isValidMask(c.mask)
			require.Equal(t, c.valid, got)
			got = matchAddrMask(ip, c.mask)
			require.Equal(t, c.exp, got)
		})
	}
}

func TestMatchMasks(t *testing.T) {
	cases := []struct {
		name  string
		addr  string
		host  string
		masks []string
		exp   bool
	}{
		{
			name:  "no masks",
			addr:  "192.168.0.10",
			masks: nil,
			exp:   true,
		},
		{
			name: "single ip",
			addr: "192.168.0.10",
			masks: []string{
				"192.168.0.10",
			},
			exp: true,
		},
		{
			name: "wrong ip",
			addr: "192.168.0.10",
			masks: []string{
				"192.168.0.11",
			},
			exp: false,
		},
		{
			name: "ip mask",
			addr: "192.168.0.10",
			masks: []string{
				"192.168.0.0/24",
			},
			exp: true,
		},
		{
			name: "wrong mask",
			addr: "192.168.0.10",
			masks: []string{
				"192.168.1.0/24",
			},
			exp: false,
		},
		{
			name: "hostname",
			addr: "192.168.0.10",
			host: "sip.example.com",
			masks: []string{
				"sip.example.com",
			},
			exp: true,
		},
		{
			name: "invalid hostname",
			addr: "192.168.0.10",
			host: "sip.example.com",
			masks: []string{
				"some.domain",
			},
			exp: false,
		},
		{
			name: "invalid and valid range",
			addr: "192.168.0.10",
			masks: []string{
				"some.domain,192.168.0.10/24",
				"192.168.0.0/24",
			},
			exp: true,
		},
		{
			name: "invalid and wrong range",
			addr: "192.168.0.10",
			masks: []string{
				"some.domain",
				"192.168.1.0/24",
			},
			exp: false,
		},
		{
			name: "domain name",
			addr: "192.168.0.10",
			host: "sip.example.com",
			masks: []string{
				"some.domain",
				"192.168.1.0/24",
				"sip.example.com",
			},
			exp: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := matchAddrMasks(c.addr, c.host, c.masks)
			require.Equal(t, c.exp, got)
		})
	}
}

func TestMatchTrunkDetailed(t *testing.T) {
	for _, c := range []struct {
		name            string
		trunks          []*livekit.SIPInboundTrunkInfo
		expMatchType    TrunkMatchType
		expTrunkID      string
		expDefaultCount int
		expErr          bool
		from            string
		to              string
		src             string
		host            string
	}{
		{
			name:         "empty",
			trunks:       nil,
			expMatchType: TrunkMatchEmpty,
			expTrunkID:   "",
			expErr:       false,
		},
		{
			name: "one wildcard",
			trunks: []*livekit.SIPInboundTrunkInfo{
				{SipTrunkId: "aaa"},
			},
			expMatchType:    TrunkMatchDefault,
			expTrunkID:      "aaa",
			expDefaultCount: 1,
			expErr:          false,
		},
		{
			name: "specific match",
			trunks: []*livekit.SIPInboundTrunkInfo{
				{SipTrunkId: "aaa", Numbers: []string{sipNumber2}},
			},
			expMatchType:    TrunkMatchSpecific,
			expTrunkID:      "aaa",
			expDefaultCount: 0,
			expErr:          false,
		},
		{
			name: "no match with trunks",
			trunks: []*livekit.SIPInboundTrunkInfo{
				{SipTrunkId: "aaa", Numbers: []string{sipNumber3}},
			},
			expMatchType:    TrunkMatchNone,
			expTrunkID:      "",
			expDefaultCount: 0,
			expErr:          false,
		},
		{
			name: "multiple defaults",
			trunks: []*livekit.SIPInboundTrunkInfo{
				{SipTrunkId: "aaa"},
				{SipTrunkId: "bbb"},
			},
			expMatchType:    TrunkMatchDefault,
			expTrunkID:      "aaa",
			expDefaultCount: 2,
			expErr:          true,
		},
		{
			name: "specific over default",
			trunks: []*livekit.SIPInboundTrunkInfo{
				{SipTrunkId: "aaa"},
				{SipTrunkId: "bbb", Numbers: []string{sipNumber2}},
			},
			expMatchType:    TrunkMatchSpecific,
			expTrunkID:      "bbb",
			expDefaultCount: 1,
			expErr:          false,
		},
		{
			name: "multiple specific",
			trunks: []*livekit.SIPInboundTrunkInfo{
				{SipTrunkId: "aaa", Numbers: []string{sipNumber2}},
				{SipTrunkId: "bbb", Numbers: []string{sipNumber2}},
			},
			expMatchType:    TrunkMatchSpecific,
			expTrunkID:      "aaa",
			expDefaultCount: 0,
			expErr:          true,
		},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			from, to, src, host := c.from, c.to, c.src, c.host
			if from == "" {
				from = sipNumber1
			}
			if to == "" {
				to = sipNumber2
			}
			if src == "" {
				src = "1.1.1.1"
			}
			if host == "" {
				host = "sip.example.com"
			}
			call := &rpc.SIPCall{
				SipCallId: "test-call-id",
				SourceIp:  src,
				From: &livekit.SIPUri{
					User: from,
					Host: host,
				},
				To: &livekit.SIPUri{
					User: to,
				},
			}
			call.Address = call.To

			var conflicts []string
			result, err := MatchTrunkDetailed(iters.Slice(c.trunks), call, WithTrunkConflict(func(t1, t2 *livekit.SIPInboundTrunkInfo, reason TrunkConflictReason) {
				conflicts = append(conflicts, fmt.Sprintf("%v: %v vs %v", reason, t1.SipTrunkId, t2.SipTrunkId))
			}))

			if c.expErr {
				require.Error(t, err)
				require.NotEmpty(t, conflicts, "expected conflicts but got none")
			} else {
				require.NoError(t, err)
				require.Empty(t, conflicts, "unexpected conflicts: %v", conflicts)

				if c.expTrunkID == "" {
					require.Nil(t, result.Trunk)
				} else {
					require.NotNil(t, result.Trunk)
					require.Equal(t, c.expTrunkID, result.Trunk.SipTrunkId)
				}

				require.Equal(t, c.expMatchType, result.MatchType)
				require.Equal(t, c.expDefaultCount, result.DefaultTrunkCount)
			}
		})
	}
}
