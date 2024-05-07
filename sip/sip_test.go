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
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/rpc"
)

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
		name: "not matching regexp",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber2, InboundNumbersRegex: []string{`^\d+$`}},
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
		name: "regexp",
		trunks: []*livekit.SIPTrunkInfo{
			{SipTrunkId: "aaa", OutboundNumber: sipNumber3},
			{SipTrunkId: "bbb", OutboundNumber: sipNumber2},
			{SipTrunkId: "ccc", OutboundNumber: sipNumber2, InboundNumbersRegex: []string{`^\d+$`}},
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
}

func TestSIPMatchTrunk(t *testing.T) {
	for _, c := range trunkCases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got, err := MatchTrunk(c.trunks, sipNumber1, sipNumber2)
			if c.expErr {
				require.Error(t, err)
				require.Nil(t, got)
				t.Log(err)
			} else {
				var exp *livekit.SIPTrunkInfo
				if c.exp >= 0 {
					exp = c.trunks[c.exp]
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
			err := ValidateTrunks(c.trunks)
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
					got, err := MatchDispatchRule(c.trunk, c.rules, newSIPReqDispatch(pin, c.noPin))
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
			err := ValidateDispatchRules(c.rules)
			if c.invalid {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
