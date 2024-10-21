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
	"sort"
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
	// - 1: Caller, aka Individual (pin-protected)
	// - 2: Callee (pin-protected)
	// - 100: Direct (open)
	// - 101: Caller, aka Individual (open)
	// - 102: Callee (open)
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
	case *livekit.SIPDispatchRule_DispatchRuleCallee:
		if rule.DispatchRuleCallee.GetPin() != "" {
			priority = 2
		} else {
			priority = 102
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
				key := ruleKey{Pin: pin, Trunk: trunk, Number: normalizeNumber(number)}
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
	case *livekit.SIPDispatchRule_DispatchRuleCallee:
		pin = rule.DispatchRuleCallee.GetPin()
		room = rule.DispatchRuleCallee.GetRoomPrefix()
	}
	return room, pin, nil
}

func printNumbers(numbers []string) string {
	if len(numbers) == 0 {
		return "<any>"
	}
	return fmt.Sprintf("%q", numbers)
}

func normalizeNumber(num string) string {
	if num == "" {
		return ""
	}
	// TODO: Always keep "number" as-is if it's not E.164.
	//       This will only matter for native SIP clients which have '+' in the username.
	if !strings.HasPrefix(num, `+`) {
		num = "+" + num
	}
	return num
}

func validateTrunkInbound(byInbound map[string]*livekit.SIPInboundTrunkInfo, t *livekit.SIPInboundTrunkInfo) error {
	if len(t.AllowedNumbers) == 0 {
		if t2 := byInbound[""]; t2 != nil {
			return fmt.Errorf("Conflicting inbound SIP Trunks: %q and %q, using the same number(s) %s without AllowedNumbers set",
				printID(t.SipTrunkId), printID(t2.SipTrunkId), printNumbers(t.Numbers))
		}
		byInbound[""] = t
	} else {
		for _, num := range t.AllowedNumbers {
			inboundKey := normalizeNumber(num)
			t2 := byInbound[inboundKey]
			if t2 != nil {
				return fmt.Errorf("Conflicting inbound SIP Trunks: %q and %q, using the same number(s) %s and AllowedNumber %q",
					printID(t.SipTrunkId), printID(t2.SipTrunkId), printNumbers(t.Numbers), num)
			}
			byInbound[inboundKey] = t
		}
	}
	return nil
}

// ValidateTrunks checks a set of trunks for conflicts.
func ValidateTrunks(trunks []*livekit.SIPInboundTrunkInfo) error {
	if len(trunks) == 0 {
		return nil
	}
	byOutboundAndInbound := make(map[string]map[string]*livekit.SIPInboundTrunkInfo)
	for _, t := range trunks {
		if len(t.Numbers) == 0 {
			byInbound := byOutboundAndInbound[""]
			if byInbound == nil {
				byInbound = make(map[string]*livekit.SIPInboundTrunkInfo)
				byOutboundAndInbound[""] = byInbound
			}
			if err := validateTrunkInbound(byInbound, t); err != nil {
				return err
			}
		} else {
			for _, num := range t.Numbers {
				byInbound := byOutboundAndInbound[num]
				if byInbound == nil {
					byInbound = make(map[string]*livekit.SIPInboundTrunkInfo)
					byOutboundAndInbound[num] = byInbound
				}
				if err := validateTrunkInbound(byInbound, t); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func matchAddrMask(ip netip.Addr, mask string) bool {
	if !strings.Contains(mask, "/") {
		expIP, err := netip.ParseAddr(mask)
		if err != nil {
			return false
		}
		return ip == expIP
	}
	pref, err := netip.ParsePrefix(mask)
	if err != nil {
		return false
	}
	return pref.Contains(ip)
}

func matchAddrMasks(addr netip.Addr, masks []string) bool {
	if !addr.IsValid() || len(masks) == 0 {
		return true
	}
	for _, mask := range masks {
		if matchAddrMask(addr, mask) {
			return true
		}
	}
	return false
}

func matchNumbers(num string, allowed []string) bool {
	if len(allowed) == 0 {
		return true
	}
	num = normalizeNumber(num)
	for _, allow := range allowed {
		if num == normalizeNumber(allow) {
			return true
		}
	}
	return false
}

// MatchTrunk finds a SIP Trunk definition matching the request.
// Returns nil if no rules matched or an error if there are conflicting definitions.
func MatchTrunk(trunks []*livekit.SIPInboundTrunkInfo, srcIP netip.Addr, calling, called string) (*livekit.SIPInboundTrunkInfo, error) {
	var (
		selectedTrunk   *livekit.SIPInboundTrunkInfo
		defaultTrunk    *livekit.SIPInboundTrunkInfo
		defaultTrunkCnt int // to error in case there are multiple ones
	)
	calledNorm := normalizeNumber(called)
	for _, tr := range trunks {
		// Do not consider it if number doesn't match.
		if !matchNumbers(calling, tr.AllowedNumbers) {
			continue
		}
		if !matchAddrMasks(srcIP, tr.AllowedAddresses) {
			continue
		}
		if len(tr.Numbers) == 0 {
			// Default/wildcard trunk.
			defaultTrunk = tr
			defaultTrunkCnt++
		} else {
			for _, num := range tr.Numbers {
				if normalizeNumber(num) == calledNorm {
					// Trunk specific to the number.
					if selectedTrunk != nil {
						return nil, fmt.Errorf("Multiple SIP Trunks matched for %q", called)
					}
					selectedTrunk = tr
					// Keep searching! We want to know if there are any conflicting Trunk definitions.
				}
			}
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
func MatchDispatchRule(trunk *livekit.SIPInboundTrunkInfo, rules []*livekit.SIPDispatchRuleInfo, req *rpc.EvaluateSIPDispatchRulesRequest) (*livekit.SIPDispatchRuleInfo, error) {
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
func EvaluateDispatchRule(projectID string, trunk *livekit.SIPInboundTrunkInfo, rule *livekit.SIPDispatchRuleInfo, req *rpc.EvaluateSIPDispatchRulesRequest) (*rpc.EvaluateSIPDispatchRulesResponse, error) {
	sentPin := req.GetPin()

	trunkID := req.SipTrunkId
	if trunk != nil {
		trunkID = trunk.SipTrunkId
	}
	attrs := maps.Clone(rule.Attributes)
	if attrs == nil {
		attrs = make(map[string]string)
	}
	for k, v := range req.ExtraAttributes {
		attrs[k] = v
	}
	attrs[livekit.AttrSIPCallID] = req.SipCallId
	attrs[livekit.AttrSIPTrunkID] = trunkID

	to := req.CalledNumber
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
		attrs[livekit.AttrSIPPhoneNumber] = req.CallingNumber
		attrs[livekit.AttrSIPTrunkNumber] = req.CalledNumber
	}

	room, rulePin, err := GetPinAndRoom(rule)
	if err != nil {
		return nil, err
	}
	if rulePin != "" {
		if sentPin == "" {
			return &rpc.EvaluateSIPDispatchRulesResponse{
				ProjectId:         projectID,
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
		// TODO: Remove "_" if the prefix is empty for consistency with Callee dispatch rule.
		// TODO: Do we need to escape specific characters in the number?
		// TODO: Include actual SIP call ID in the room name?
		room = fmt.Sprintf("%s_%s_%s", rule.DispatchRuleIndividual.GetRoomPrefix(), from, guid.New(""))
	case *livekit.SIPDispatchRule_DispatchRuleCallee:
		room = to
		if pref := rule.DispatchRuleCallee.GetRoomPrefix(); pref != "" {
			room = pref + "_" + to
		}
		if rule.DispatchRuleCallee.Randomize {
			room += "_" + guid.New("")
		}
	}
	attrs[livekit.AttrSIPDispatchRuleID] = rule.SipDispatchRuleId
	resp := &rpc.EvaluateSIPDispatchRulesResponse{
		ProjectId:             projectID,
		SipTrunkId:            trunkID,
		SipDispatchRuleId:     rule.SipDispatchRuleId,
		Result:                rpc.SIPDispatchResult_ACCEPT,
		RoomName:              room,
		ParticipantIdentity:   fromID,
		ParticipantName:       fromName,
		ParticipantMetadata:   rule.Metadata,
		ParticipantAttributes: attrs,
	}
	if trunk != nil {
		resp.Headers = trunk.Headers
		resp.HeadersToAttributes = trunk.HeadersToAttributes
		resp.RingingTimeout = trunk.RingingTimeout
		resp.MaxCallDuration = trunk.MaxCallDuration
		if trunk.KrispEnabled {
			resp.EnabledFeatures = append(resp.EnabledFeatures, rpc.SIPFeature_KRISP_ENABLED)
		}
	}
	return resp, nil
}
