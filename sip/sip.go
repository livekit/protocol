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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"maps"
	"math"
	"net/netip"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/rpc"
	"github.com/livekit/protocol/utils"
	"github.com/livekit/protocol/utils/guid"
)

func NewCallID() string {
	return guid.New(utils.SIPCallPrefix)
}

type ErrNoDispatchMatched struct {
	NoRules      bool
	NoTrunks     bool
	CalledNumber string
}

func (e *ErrNoDispatchMatched) Error() string {
	if e.NoRules {
		return "No SIP Dispatch Rules defined"
	}
	if e.NoTrunks {
		return fmt.Sprintf("No SIP Trunk or Dispatch Rules matched for %q", e.CalledNumber)
	}
	return fmt.Sprintf("No SIP Dispatch Rules matched for %q", e.CalledNumber)
}

// DispatchRulePriority returns sorting priority for dispatch rules. Lower value means higher priority.
func DispatchRulePriority(info *livekit.SIPDispatchRuleInfo) int32 {
	// In all these cases, prefer pin-protected rules and rules for specific calling number.
	// Thus, the order will be the following:
	// - 0: Direct or Pin (both pin-protected)
	// - 1: Individual (pin-protected)
	// - 100: Direct (open)
	// - 101: Individual (open)
	// Also, add 1K penalty for not specifying the calling number.
	const (
		last = math.MaxInt32
	)
	// TODO: Maybe allow setting specific priorities for dispatch rules?
	priority := int32(0)
	switch rule := info.GetRule().GetRule().(type) {
	default:
		return last
	case *livekit.SIPDispatchRule_DispatchRuleDirect:
		if rule.DispatchRuleDirect.GetPin() != "" {
			priority = 0
		} else {
			priority = 100
		}
	case *livekit.SIPDispatchRule_DispatchRuleIndividual:
		if rule.DispatchRuleIndividual.GetPin() != "" {
			priority = 1
		} else {
			priority = 101
		}
	}
	if len(info.InboundNumbers) == 0 {
		priority += 1000
	}
	return priority
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

func printID(s string) string {
	if s == "" {
		return "<new>"
	}
	return s
}

// ValidateDispatchRules checks a set of dispatch rules for conflicts.
func ValidateDispatchRules(rules []*livekit.SIPDispatchRuleInfo) error {
	if len(rules) == 0 {
		return nil
	}
	type ruleKey struct {
		Pin    string
		Trunk  string
		Number string
	}
	byRuleKey := make(map[ruleKey]*livekit.SIPDispatchRuleInfo)
	for _, r := range rules {
		_, pin, err := GetPinAndRoom(r)
		if err != nil {
			return err
		}
		trunks := r.TrunkIds
		if len(trunks) == 0 {
			// This rule matches all trunks, but collides only with other default ones (specific rules take priority).
			trunks = []string{""}
		}
		numbers := r.InboundNumbers
		if len(numbers) == 0 {
			// This rule matches all numbers, but collides only with other default ones (specific rules take priority).
			numbers = []string{""}
		}
		for _, trunk := range trunks {
			for _, number := range numbers {
				key := ruleKey{Pin: pin, Trunk: trunk, Number: number}
				r2 := byRuleKey[key]
				if r2 != nil {
					return fmt.Errorf("Conflicting SIP Dispatch Rules: same Trunk+Number+PIN combination for for %q and %q",
						printID(r.SipDispatchRuleId), printID(r2.SipDispatchRuleId))
				}
				byRuleKey[key] = r
			}
		}
	}
	return nil
}

// SelectDispatchRule takes a list of dispatch rules, and takes the decision which one should be selected.
// It returns an error if there are conflicting rules. Returns nil if no rules match.
func SelectDispatchRule(rules []*livekit.SIPDispatchRuleInfo, req *rpc.EvaluateSIPDispatchRulesRequest) (*livekit.SIPDispatchRuleInfo, error) {
	if len(rules) == 0 {
		// Nil is fine here. We will report "no rules matched" later.
		return nil, nil
	}
	if err := ValidateDispatchRules(rules); err != nil {
		return nil, err
	}
	// Sorting will do the selection for us. We already filtered out irrelevant ones in MatchDispatchRule and above.
	SortDispatchRules(rules)
	return rules[0], nil
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

func printNumber(s string) string {
	if s == "" {
		return "<any>"
	}
	return strconv.Quote(s)
}

// ValidateTrunks checks a set of trunks for conflicts.
func ValidateTrunks(trunks []*livekit.SIPTrunkInfo) error {
	if len(trunks) == 0 {
		return nil
	}
	byOutboundAndInbound := make(map[string]map[string]*livekit.SIPTrunkInfo)
	for _, t := range trunks {
		if len(t.InboundNumbersRegex) != 0 {
			continue // can't effectively validate these
		}
		byInbound := byOutboundAndInbound[t.OutboundNumber]
		if byInbound == nil {
			byInbound = make(map[string]*livekit.SIPTrunkInfo)
			byOutboundAndInbound[t.OutboundNumber] = byInbound
		}
		if len(t.InboundNumbers) == 0 {
			if t2 := byInbound[""]; t2 != nil {
				return fmt.Errorf("Conflicting SIP Trunks: %q and %q, using the same OutboundNumber %s without InboundNumbers set",
					printID(t.SipTrunkId), printID(t2.SipTrunkId), printNumber(t.OutboundNumber))
			}
			byInbound[""] = t
		} else {
			for _, num := range t.InboundNumbers {
				t2 := byInbound[num]
				if t2 != nil {
					return fmt.Errorf("Conflicting SIP Trunks: %q and %q, using the same OutboundNumber %s and InboundNumber %q",
						printID(t.SipTrunkId), printID(t2.SipTrunkId), printNumber(t.OutboundNumber), num)
				}
				byInbound[num] = t
			}
		}
	}
	return nil
}

func matchAddrs(addr string, mask string) bool {
	if !strings.Contains(mask, "/") {
		return addr == mask
	}
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return false
	}
	pref, err := netip.ParsePrefix(mask)
	if err != nil {
		return false
	}
	return pref.Contains(ip)
}

func matchAddr(addr string, masks []string) bool {
	if addr == "" {
		return true
	}
	for _, mask := range masks {
		if !matchAddrs(addr, mask) {
			return false
		}
	}
	return true
}

// MatchTrunk finds a SIP Trunk definition matching the request.
// Returns nil if no rules matched or an error if there are conflicting definitions.
func MatchTrunk(trunks []*livekit.SIPTrunkInfo, srcIP, calling, called string) (*livekit.SIPTrunkInfo, error) {
	var (
		selectedTrunk   *livekit.SIPTrunkInfo
		defaultTrunk    *livekit.SIPTrunkInfo
		defaultTrunkCnt int // to error in case there are multiple ones
	)
	for _, tr := range trunks {
		// Do not consider it if number doesn't match.
		if len(tr.InboundNumbers) != 0 && !slices.Contains(tr.InboundNumbers, calling) {
			continue
		}
		if !matchAddr(srcIP, tr.InboundAddresses) {
			continue
		}
		// Deprecated, but we still check it for backward compatibility.
		matchesRe := len(tr.InboundNumbersRegex) == 0
		for _, reStr := range tr.InboundNumbersRegex {
			// TODO: we should cache it
			re, err := regexp.Compile(reStr)
			if err != nil {
				logger.Errorw("cannot parse SIP trunk regexp", err, "trunkID", tr.SipTrunkId)
				continue
			}
			if re.MatchString(calling) {
				matchesRe = true
				break
			}
		}
		if !matchesRe {
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
		return nil, &ErrNoDispatchMatched{NoRules: true, NoTrunks: trunk == nil, CalledNumber: req.CalledNumber}
	}
	// We split the matched dispatch rules into two sets in relation to Trunks: specific and default (aka wildcard).
	// First, attempt to match any of the specific rules, where we did match the Trunk ID.
	// If nothing matches there - fallback to default/wildcard rules, where no Trunk IDs were mentioned.
	var (
		specificRules []*livekit.SIPDispatchRuleInfo
		defaultRules  []*livekit.SIPDispatchRuleInfo
	)
	noPin := req.NoPin
	sentPin := req.GetPin()
	for _, info := range rules {
		if len(info.InboundNumbers) != 0 && !slices.Contains(info.InboundNumbers, req.CallingNumber) {
			continue
		}
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
		if !slices.Contains(info.TrunkIds, trunk.SipTrunkId) {
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
	return nil, &ErrNoDispatchMatched{NoRules: false, NoTrunks: trunk == nil, CalledNumber: req.CalledNumber}
}

// EvaluateDispatchRule checks a selected Dispatch Rule against the provided request.
func EvaluateDispatchRule(trunkID string, rule *livekit.SIPDispatchRuleInfo, req *rpc.EvaluateSIPDispatchRulesRequest) (*rpc.EvaluateSIPDispatchRulesResponse, error) {
	sentPin := req.GetPin()

	attrs := maps.Clone(rule.Attributes)
	if attrs == nil {
		attrs = make(map[string]string)
	}
	for k, v := range req.ExtraAttributes {
		attrs[k] = v
	}
	attrs[livekit.AttrSIPCallID] = req.SipCallId
	attrs[livekit.AttrSIPTrunkID] = trunkID

	from := req.CallingNumber
	fromName := "Phone " + req.CallingNumber
	fromID := "sip_" + req.CallingNumber
	if rule.HidePhoneNumber {
		// Mask the phone number, hash identity. Omit number in attrs.
		h := sha256.Sum256([]byte(req.CallingNumber))
		fromID = "sip_" + hex.EncodeToString(h[:8])
		// TODO: Maybe keep regional code, but mask all but 4 last digits?
		n := 4
		if len(from) <= 4 {
			n = 1
		}
		from = from[len(from)-n:]
		fromName = "Phone " + from
	} else {
		attrs[livekit.AttrSIPFromNumber] = req.CallingNumber
		attrs[livekit.AttrSIPToNumber] = req.CalledNumber
	}

	room, rulePin, err := GetPinAndRoom(rule)
	if err != nil {
		return nil, err
	}
	if rulePin != "" {
		if sentPin == "" {
			return &rpc.EvaluateSIPDispatchRulesResponse{
				SipTrunkId:        trunkID,
				SipDispatchRuleId: rule.SipDispatchRuleId,
				Result:            rpc.SIPDispatchResult_REQUEST_PIN,
				RequestPin:        true,
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
		// TODO: Do we need to escape specific characters in the number?
		// TODO: Include actual SIP call ID in the room name?
		room = fmt.Sprintf("%s_%s_%s", rule.DispatchRuleIndividual.GetRoomPrefix(), from, guid.New(""))
	}
	attrs[livekit.AttrSIPDispatchRuleID] = rule.SipDispatchRuleId
	return &rpc.EvaluateSIPDispatchRulesResponse{
		SipTrunkId:            trunkID,
		SipDispatchRuleId:     rule.SipDispatchRuleId,
		Result:                rpc.SIPDispatchResult_ACCEPT,
		RoomName:              room,
		ParticipantIdentity:   fromID,
		ParticipantName:       fromName,
		ParticipantMetadata:   rule.Metadata,
		ParticipantAttributes: attrs,
	}, nil
}
