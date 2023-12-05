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
	"math"
	"regexp"
	"sort"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/rpc"
)

// DispatchRulePriority returns sorting priority for dispatch rules. Lower value means higher priority.
func DispatchRulePriority(info *livekit.SIPDispatchRuleInfo) int32 {
	// In all these cases, prefer pin-protected rules.
	// Thus, the order will be the following:
	// - 0: Direct or Pin (both pin-protected)
	// - 1: Individual (pin-protected)
	// - 100: Direct (open)
	// - 101: Individual (open)
	const (
		last = math.MaxInt32
	)
	// TODO: Maybe allow setting specific priorities for dispatch rules?
	switch rule := info.GetRule().GetRule().(type) {
	default:
		return last
	case *livekit.SIPDispatchRule_DispatchRuleDirect:
		if rule.DispatchRuleDirect.GetPin() != "" {
			return 0
		}
		return 100
	case *livekit.SIPDispatchRule_DispatchRuleIndividual:
		if rule.DispatchRuleIndividual.GetPin() != "" {
			return 1
		}
		return 101
	}
}

// SortDispatchRules predictably sorts dispatch rules by priority (first one is highest).
func SortDispatchRules(rules []*livekit.SIPDispatchRuleInfo) {
	sort.Slice(rules, func(i, j int) bool {
		p1, p2 := DispatchRulePriority(rules[i]), DispatchRulePriority(rules[j])
		if p1 < p2 {
			return true
		} else if p1 > p2 {
			return false
		}
		// For predictable sorting order.
		room1, _, _ := GetPinAndRoom(rules[i])
		room2, _, _ := GetPinAndRoom(rules[j])
		return room1 < room2
	})
}

// SelectDispatchRule takes a list of dispatch rules, and takes the decision which one should be selected.
// It returns an error if there are conflicting rules. Returns nil if no rules match.
func SelectDispatchRule(rules []*livekit.SIPDispatchRuleInfo, req *rpc.EvaluateSIPDispatchRulesRequest) (*livekit.SIPDispatchRuleInfo, error) {
	if len(rules) == 0 {
		return nil, nil
	}
	// Sorting will do the selection for us. We already filtered out irrelevant ones in matchSIPDispatchRule.
	SortDispatchRules(rules)
	byPin := make(map[string]*livekit.SIPDispatchRuleInfo)
	var (
		pinRule  *livekit.SIPDispatchRuleInfo
		openRule *livekit.SIPDispatchRuleInfo
	)
	openCnt := 0
	for _, r := range rules {
		_, pin, err := GetPinAndRoom(r)
		if err != nil {
			return nil, err
		}
		if pin == "" {
			openRule = r // last one
			openCnt++
		} else if r2 := byPin[pin]; r2 != nil {
			return nil, fmt.Errorf("Conflicting SIP Dispatch Rules: Same PIN for %q and %q",
				r.SipDispatchRuleId, r2.SipDispatchRuleId)
		} else {
			byPin[pin] = r
			// Pick the first one with a Pin. If Pin was provided in the request, we already filtered the right rules.
			// If not, this rule will just be used to send RequestPin=true flag.
			if pinRule == nil {
				pinRule = r
			}
		}
	}
	if req.GetPin() != "" {
		// If it's still nil that's fine. We will report "no rules matched" later.
		return pinRule, nil
	}
	if pinRule != nil {
		return pinRule, nil
	}
	if openCnt > 1 {
		return nil, fmt.Errorf("Conflicting SIP Dispatch Rules: Matched %d open rules for %q", openCnt, req.CallingNumber)
	}
	return openRule, nil
}

// GetPinAndRoom returns a room name/prefix and the pin for a dispatch rule. Just a convenience wrapper.
func GetPinAndRoom(info *livekit.SIPDispatchRuleInfo) (room, pin string, err error) {
	// TODO: Could probably add methods on SIPDispatchRuleInfo struct instead.
	switch rule := info.GetRule().GetRule().(type) {
	default:
		return "", "", fmt.Errorf("Unsupported SIP Dispatch Rule: %T", rule)
	case *livekit.SIPDispatchRule_DispatchRuleDirect:
		pin = rule.DispatchRuleDirect.GetPin()
		room = rule.DispatchRuleDirect.GetRoomName()
	case *livekit.SIPDispatchRule_DispatchRuleIndividual:
		pin = rule.DispatchRuleIndividual.GetPin()
		room = rule.DispatchRuleIndividual.GetRoomPrefix()
	}
	return room, pin, nil
}

// MatchTrunk finds a SIP Trunk definition matching the request.
// Returns nil if no rules matched or an error if there are conflicting definitions.
func MatchTrunk(trunks []*livekit.SIPTrunkInfo, calling, called string) (*livekit.SIPTrunkInfo, error) {
	var (
		selectedTrunk   *livekit.SIPTrunkInfo
		defaultTrunk    *livekit.SIPTrunkInfo
		defaultTrunkCnt int // to error in case there are multiple ones
	)
	for _, tr := range trunks {
		// Do not consider it if regexp doesn't match.
		matches := len(tr.InboundNumbersRegex) == 0
		for _, reStr := range tr.InboundNumbersRegex {
			// TODO: we should cache it
			re, err := regexp.Compile(reStr)
			if err != nil {
				logger.Errorw("cannot parse SIP trunk regexp", err, "trunkID", tr.SipTrunkId)
				continue
			}
			if re.MatchString(calling) {
				matches = true
				break
			}
		}
		if !matches {
			continue
		}
		if tr.OutboundNumber == "" {
			// Default/wildcard trunk.
			defaultTrunk = tr
			defaultTrunkCnt++
		} else if tr.OutboundNumber == called {
			// Trunk specific to the number.
			if selectedTrunk != nil {
				return nil, fmt.Errorf("Multiple SIP Trunks matched for %q", called)
			}
			selectedTrunk = tr
			// Keep searching! We want to know if there are any conflicting Trunk definitions.
		}
	}
	if selectedTrunk != nil {
		return selectedTrunk, nil
	}
	if defaultTrunkCnt > 1 {
		return nil, fmt.Errorf("Multiple default SIP Trunks matched for %q", called)
	}
	// Could still be nil here.
	return defaultTrunk, nil
}

// MatchDispatchRule finds the best dispatch rule matching the request parameters. Returns an error if no rule matched.
// Trunk parameter can be nil, in which case only wildcard dispatch rules will be effective (ones without Trunk IDs).
func MatchDispatchRule(trunk *livekit.SIPTrunkInfo, rules []*livekit.SIPDispatchRuleInfo, req *rpc.EvaluateSIPDispatchRulesRequest) (*livekit.SIPDispatchRuleInfo, error) {
	// Trunk can still be nil here in case none matched or were defined.
	// This is still fine, but only in case we'll match exactly one wildcard dispatch rule.
	if len(rules) == 0 {
		return nil, fmt.Errorf("No SIP Dispatch Rules defined")
	}
	// We split the matched dispatch rules into two sets: specific and default (aka wildcard).
	// First, attempt to match any of the specific rules, where we did match the Trunk ID.
	// If nothing matches there - fallback to default/wildcard rules, where no Trunk IDs were mentioned.
	var (
		specificRules []*livekit.SIPDispatchRuleInfo
		defaultRules  []*livekit.SIPDispatchRuleInfo
	)
	noPin := req.NoPin
	sentPin := req.GetPin()
	for _, info := range rules {
		_, rulePin, err := GetPinAndRoom(info)
		if err != nil {
			logger.Errorw("Invalid SIP Dispatch Rule", err, "dispatchRuleID", info.SipDispatchRuleId)
			continue
		}
		// Filter heavily on the Pin, so that only relevant rules remain.
		if noPin {
			if rulePin != "" {
				// Skip pin-protected rules if no pin mode requested.
				continue
			}
		} else if sentPin != "" {
			if rulePin == "" {
				// Pin already sent, skip non-pin-protected rules.
				continue
			}
			if sentPin != rulePin {
				// Pin doesn't match. Don't return an error here, just wait for other rule to match (or none at all).
				// Note that we will NOT match non-pin-protected rules, thus it will not fallback to open rules.
				continue
			}
		}
		if len(info.TrunkIds) == 0 {
			// Default/wildcard dispatch rule.
			defaultRules = append(defaultRules, info)
			continue
		}
		// Specific dispatch rules. Require a Trunk associated with the number.
		if trunk == nil {
			continue
		}
		matches := false
		for _, id := range info.TrunkIds {
			if id == trunk.SipTrunkId {
				matches = true
				break
			}
		}
		if !matches {
			continue
		}
		specificRules = append(specificRules, info)
	}
	best, err := SelectDispatchRule(specificRules, req)
	if err != nil {
		return nil, err
	} else if best != nil {
		return best, nil
	}
	best, err = SelectDispatchRule(defaultRules, req)
	if err != nil {
		return nil, err
	} else if best != nil {
		return best, nil
	}
	if trunk == nil {
		return nil, fmt.Errorf("No SIP Trunk or Dispatch Rules matched for %q", req.CalledNumber)
	}
	return nil, fmt.Errorf("No SIP Dispatch Rules matched for %q", req.CalledNumber)
}

// EvaluateDispatchRule checks a selected Dispatch Rule against the provided request.
func EvaluateDispatchRule(rule *livekit.SIPDispatchRuleInfo, req *rpc.EvaluateSIPDispatchRulesRequest) (*rpc.EvaluateSIPDispatchRulesResponse, error) {
	sentPin := req.GetPin()

	from := req.CallingNumber
	if rule.HidePhoneNumber {
		// TODO: Decide on the phone masking format.
		//       Maybe keep regional code, but mask all but 4 last digits?
		from = from[len(from)-4:]
	}
	fromName := "Phone " + from

	room, rulePin, err := GetPinAndRoom(rule)
	if err != nil {
		return nil, err
	}
	if rulePin != "" {
		if sentPin == "" {
			return &rpc.EvaluateSIPDispatchRulesResponse{
				RequestPin: true,
			}, nil
		}
		if rulePin != sentPin {
			// This should never happen in practice, because matchSIPDispatchRule should remove rules with the wrong pin.
			return nil, fmt.Errorf("Incorrect PIN for SIP room")
		}
	} else {
		// Pin was sent, but room doesn't require one. Assume user accidentally pressed phone button.
	}
	switch rule := rule.GetRule().GetRule().(type) {
	case *livekit.SIPDispatchRule_DispatchRuleIndividual:
		// TODO: Decide on the suffix. Do we need to escape specific characters?
		room = rule.DispatchRuleIndividual.GetRoomPrefix() + from
	}
	return &rpc.EvaluateSIPDispatchRulesResponse{
		RoomName:            room,
		ParticipantIdentity: fromName,
	}, nil
}
