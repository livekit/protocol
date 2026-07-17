package livekit

import (
	errors "errors"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/livekit/protocol/utils/prototest"
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
	type validateable interface {
		Validate() error
	}
	type resultValidateable interface {
		ValidateResult() ValidationResult
	}
	type validateTestCase struct {
		name         string
		req          validateable
		exp          bool
		wantSoftErrs bool
	}

	cases := map[string][]validateTestCase{
		"SIPInboundTrunkInfo": {
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
				name: "inbound invalid header value",
				req: &SIPInboundTrunkInfo{
					Numbers: []string{"+1111"},
					HeadersToAttributes: map[string]string{
						"From": "<sip:u7@example.com", // missing closing bracket
					},
				},
				exp:          true,
				wantSoftErrs: true,
			},
		},
		"SIPOutboundTrunkInfo": {
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
			{
				name: "outbound invalid header value",
				req: &SIPOutboundTrunkInfo{
					Address: "sip.example.com",
					Numbers: []string{"+2222"},
					HeadersToAttributes: map[string]string{
						"From": "<sip:u7@example.com", // missing closing bracket
					},
				},
				exp:          true,
				wantSoftErrs: true,
			},
		},
		"SIPMediaConfig": {
			{
				name: "media_timeout_missing",
				req:  &SIPMediaConfig{},
				exp:  true,
			},
			{
				name: "media_timeout_ok",
				req: &SIPMediaConfig{
					MediaTimeout: durationpb.New(5 * time.Minute),
				},
				exp: true,
			},
			{
				name: "media_timeout_negative",
				req: &SIPMediaConfig{
					MediaTimeout: durationpb.New(-1 * time.Minute),
				},
				exp: false,
			},
			{
				name: "media_timeout_zero",
				req: &SIPMediaConfig{
					MediaTimeout: durationpb.New(0 * time.Minute),
				},
				exp: true,
			},
			{
				name: "media_timeout_over_max",
				req: &SIPMediaConfig{
					MediaTimeout: durationpb.New(20 * time.Minute),
				},
				exp: false,
			},
		},
		"CreateSIPDispatchRuleRequest": {
			{
				name: "dispatch_rule_validates_embedded_media_ok",
				req: &CreateSIPDispatchRuleRequest{
					DispatchRule: &SIPDispatchRuleInfo{
						Rule: &SIPDispatchRule{
							Rule: &SIPDispatchRule_DispatchRuleDirect{
								DispatchRuleDirect: &SIPDispatchRuleDirect{RoomName: "r"},
							},
						},
						// Sanity
					},
				},
				exp: true,
			},
			{
				name: "dispatch_rule_validates_embedded_media_timeout_over_max",
				req: &CreateSIPDispatchRuleRequest{
					DispatchRule: &SIPDispatchRuleInfo{
						Rule: &SIPDispatchRule{
							Rule: &SIPDispatchRule_DispatchRuleDirect{
								DispatchRuleDirect: &SIPDispatchRuleDirect{RoomName: "r"},
							},
						},
						// Just here to make sure it links to SIPMediaConfig.Validate()
						Media: &SIPMediaConfig{
							MediaTimeout: durationpb.New(20 * time.Minute),
						},
					},
				},
				exp: false,
			},
		},
		"UpdateSIPDispatchRuleRequest": {
			{
				name: "update_diff_validates_embedded_media_ok",
				req: &UpdateSIPDispatchRuleRequest{
					SipDispatchRuleId: "id",
					Action: &UpdateSIPDispatchRuleRequest_Update{
						Update: &SIPDispatchRuleUpdate{}, // Sanity
					},
				},
				exp: true,
			},
			{
				name: "update_diff_validates_embedded_media_timeout_over_max",
				req: &UpdateSIPDispatchRuleRequest{
					SipDispatchRuleId: "id",
					Action: &UpdateSIPDispatchRuleRequest_Update{
						Update: &SIPDispatchRuleUpdate{
							// Just here to make sure it links to SIPMediaConfig.Validate()
							Media: &SIPMediaConfig{
								MediaTimeout: durationpb.New(20 * time.Minute),
							},
						},
					},
				},
				exp: false,
			},
		},
		"CreateSIPParticipantRequest": {
			{
				name: "participant_validates_embedded_media_ok",
				req: &CreateSIPParticipantRequest{
					SipTrunkId: "trunk",
					SipCallTo:  "+1000",
					RoomName:   "room",
					// Sanity
				},
				exp: true,
			},
			{
				name: "participant_validates_embedded_media_timeout_over_max",
				req: &CreateSIPParticipantRequest{
					SipTrunkId: "trunk",
					SipCallTo:  "+1000",
					RoomName:   "room",
					// Just here to make sure it links to SIPMediaConfig.Validate()
					Media: &SIPMediaConfig{
						MediaTimeout: durationpb.New(20 * time.Minute),
					},
				},
				exp: false,
			},
			{
				name: "invalid_header_value_soft_failure",
				req: &CreateSIPParticipantRequest{
					SipTrunkId: "trunk",
					SipCallTo:  "+1000",
					RoomName:   "room",
					Headers: map[string]string{
						"From": "<sip:u7@example.com", // missing closing bracket
					},
					// Sanity
				},
				exp:          true,
				wantSoftErrs: true,
			},
		},
		"CreateSIPInboundTrunkRequest": {
			{
				name: "create_inbound_missing_trunk",
				req:  &CreateSIPInboundTrunkRequest{},
				exp:  false,
			},
			{
				name: "create_inbound_trunk_id_set",
				req: &CreateSIPInboundTrunkRequest{
					Trunk: &SIPInboundTrunkInfo{
						SipTrunkId: "id",
						Numbers:    []string{"+1111"},
					},
				},
				exp: false,
			},
			{
				name: "create_inbound_valid",
				req: &CreateSIPInboundTrunkRequest{
					Trunk: &SIPInboundTrunkInfo{
						Numbers: []string{"+1111"},
					},
				},
				exp: true,
			},
			{
				name: "create_inbound_valid_soft_failure",
				req: &CreateSIPInboundTrunkRequest{
					Trunk: &SIPInboundTrunkInfo{
						Numbers: []string{"+1111"},
						Headers: map[string]string{
							"From": "<sip:u7@example.com", // missing closing bracket
						},
					},
				},
				exp:          true,
				wantSoftErrs: true,
			},
		},
		"CreateSIPOutboundTrunkRequest": {
			{
				name: "create_outbound_missing_trunk",
				req:  &CreateSIPOutboundTrunkRequest{},
				exp:  false,
			},
			{
				name: "create_outbound_trunk_id_set",
				req: &CreateSIPOutboundTrunkRequest{
					Trunk: &SIPOutboundTrunkInfo{
						SipTrunkId: "id",
					},
				},
				exp: false,
			},
			{
				name: "create_outbound_valid",
				req: &CreateSIPOutboundTrunkRequest{
					Trunk: &SIPOutboundTrunkInfo{
						Address: "sip.example.com",
						Numbers: []string{"+2222"},
					},
				},
				exp: true,
			},
			{
				name: "create_outbound_valid_soft_failure",
				req: &CreateSIPOutboundTrunkRequest{
					Trunk: &SIPOutboundTrunkInfo{
						Address: "sip.example.com",
						Numbers: []string{"+2222"},
						Headers: map[string]string{
							"From": "<sip:u7@example.com", // missing closing bracket
						},
					},
				},
				exp:          true,
				wantSoftErrs: true,
			},
		},
		"UpdateSIPInboundTrunkRequest": {
			{
				name: "update_inbound_missing_id",
				req:  &UpdateSIPInboundTrunkRequest{},
				exp:  false,
			},
			{
				name: "update_inbound_missing_action",
				req: &UpdateSIPInboundTrunkRequest{
					SipTrunkId: "id",
				},
				exp: false,
			},
			{
				name: "update_inbound_update_ok",
				req: &UpdateSIPInboundTrunkRequest{
					SipTrunkId: "id",
					Action: &UpdateSIPInboundTrunkRequest_Update{
						Update: &SIPInboundTrunkUpdate{},
					},
				},
				exp: true,
			},
			{
				name: "update_inbound_replace_ok",
				req: &UpdateSIPInboundTrunkRequest{
					SipTrunkId: "id",
					Action: &UpdateSIPInboundTrunkRequest_Replace{
						Replace: &SIPInboundTrunkInfo{
							Numbers: []string{"+1111"},
						},
					},
				},
				exp: true,
			},
			{
				name: "update_inbound_replace_ok_soft_failrues",
				req: &UpdateSIPInboundTrunkRequest{
					SipTrunkId: "id",
					Action: &UpdateSIPInboundTrunkRequest_Replace{
						Replace: &SIPInboundTrunkInfo{
							Numbers: []string{"+1111"},
							Headers: map[string]string{
								"From": "<sip:u7@example.com", // missing closing bracket
							},
						},
					},
				},
				exp:          true,
				wantSoftErrs: true,
			},
			{
				name: "update_inbound_replace_id_mismatch",
				req: &UpdateSIPInboundTrunkRequest{
					SipTrunkId: "id",
					Action: &UpdateSIPInboundTrunkRequest_Replace{
						Replace: &SIPInboundTrunkInfo{
							SipTrunkId: "other",
							Numbers:    []string{"+1111"},
						},
					},
				},
				exp: false,
			},
		},
		"UpdateSIPOutboundTrunkRequest": {
			{
				name: "update_outbound_missing id",
				req:  &UpdateSIPOutboundTrunkRequest{},
				exp:  false,
			},
			{
				name: "update_outbound_missing_action",
				req: &UpdateSIPOutboundTrunkRequest{
					SipTrunkId: "id",
				},
				exp: false,
			},
			{
				name: "update_outbound_update_ok",
				req: &UpdateSIPOutboundTrunkRequest{
					SipTrunkId: "id",
					Action: &UpdateSIPOutboundTrunkRequest_Update{
						Update: &SIPOutboundTrunkUpdate{},
					},
				},
				exp: true,
			},
			{
				name: "update_outbound_replace_ok",
				req: &UpdateSIPOutboundTrunkRequest{
					SipTrunkId: "id",
					Action: &UpdateSIPOutboundTrunkRequest_Replace{
						Replace: &SIPOutboundTrunkInfo{
							Address: "sip.example.com",
							Numbers: []string{"+2222"},
						},
					},
				},
				exp: true,
			},
			{
				name: "update_outbound_replace_ok_soft_failures",
				req: &UpdateSIPOutboundTrunkRequest{
					SipTrunkId: "id",
					Action: &UpdateSIPOutboundTrunkRequest_Replace{
						Replace: &SIPOutboundTrunkInfo{
							Address: "sip.example.com",
							Numbers: []string{"+2222"},
							Headers: map[string]string{
								"From": "<sip:u7@example.com", // missing closing bracket
							},
						},
					},
				},
				exp:          true,
				wantSoftErrs: true,
			},
			{
				name: "update_outbound_replace_id_mismatch",
				req: &UpdateSIPOutboundTrunkRequest{
					SipTrunkId: "id",
					Action: &UpdateSIPOutboundTrunkRequest_Replace{
						Replace: &SIPOutboundTrunkInfo{
							SipTrunkId: "other",
							Address:    "sip.example.com",
							Numbers:    []string{"+2222"},
						},
					},
				},
				exp: false,
			},
		},
		"TransferSIPParticipantRequest": {
			{
				name: "transfer_valid",
				req: &TransferSIPParticipantRequest{
					RoomName:            "room1",
					ParticipantIdentity: "participant1",
					TransferTo:          "tel:+15105550100",
					Headers: map[string]string{
						"From": "<sip:u7@example.com",
					},
				},
				exp:          true,
				wantSoftErrs: true,
			},
			{
				name: "transfer_valid_soft_failures",
				req: &TransferSIPParticipantRequest{
					RoomName:            "room1",
					ParticipantIdentity: "participant1",
					TransferTo:          "tel:+15105550100",
				},
				exp: true,
			},
			{
				name: "transfer_missing_room",
				req:  &TransferSIPParticipantRequest{},
				exp:  false,
			},
			{
				name: "transfer_missing_participant",
				req: &TransferSIPParticipantRequest{
					RoomName: "room1",
				},
				exp: false,
			},
			{
				name: "transfer_invalid_uri",
				req: &TransferSIPParticipantRequest{
					RoomName:            "room1",
					ParticipantIdentity: "participant1",
					TransferTo:          "http://example.com",
				},
				exp: false,
			},
		},
	}

	resultValidatableTypes := make(map[string]bool)

	for name, class := range cases {
		t.Run(name, func(t *testing.T) {
			for _, c := range class {
				t.Run(c.name, func(t *testing.T) {
					err := c.req.Validate()
					require.Equal(t, c.exp, err == nil, "error: %v", err)

					if v, ok := c.req.(resultValidateable); ok {
						result := v.ValidateResult()
						require.Equal(t, c.exp, result.Error() == nil, "error: %v", err)
						resultValidatableTypes[fmt.Sprintf("%T", v)] = true

						hasSoftErrs := len(result.SoftErrors()) > 0
						if c.wantSoftErrs {
							require.Equal(t, c.wantSoftErrs, hasSoftErrs, "got zero soft validation errors; want at least one")
						} else {
							require.Equal(t, c.wantSoftErrs, hasSoftErrs, "got at least one soft validation error; want zero; %v", result.SoftErrors())
						}
					}
				})
			}
		})
	}

	wantResultValidateableTypes := map[string]bool{
		"*livekit.CreateSIPInboundTrunkRequest":  true,
		"*livekit.CreateSIPOutboundTrunkRequest": true,
		"*livekit.CreateSIPParticipantRequest":   true,
		"*livekit.SIPInboundTrunkInfo":           true,
		"*livekit.SIPOutboundTrunkInfo":          true,
		"*livekit.TransferSIPParticipantRequest": true,
		"*livekit.UpdateSIPInboundTrunkRequest":  true,
		"*livekit.UpdateSIPOutboundTrunkRequest": true,
	}
	require.Equal(t, wantResultValidateableTypes, resultValidatableTypes)
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

func TestTransferSIPParticipantRequestValidate(t *testing.T) {
	cases := []struct {
		name        string
		req         *TransferSIPParticipantRequest
		expectError bool
		expectedURI string // Expected TransferTo after validation
	}{
		{
			name: "valid sip URI without brackets",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "sip:+15105550100@sip.telnyx.com",
			},
			expectError: false,
			expectedURI: "<sip:+15105550100@sip.telnyx.com>",
		},
		{
			name: "valid sips URI without brackets",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "sips:+15105550100@sip.telnyx.com",
			},
			expectError: false,
			expectedURI: "<sips:+15105550100@sip.telnyx.com>",
		},
		{
			name: "valid tel URI without brackets",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "tel:+15105550100",
			},
			expectError: false,
			expectedURI: "tel:+15105550100",
		},
		{
			name: "valid sip URI with brackets",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "<sip:+15105550100@sip.telnyx.com>",
			},
			expectError: false,
			expectedURI: "<sip:+15105550100@sip.telnyx.com>",
		},
		{
			name: "valid sips URI with brackets",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "<sips:+15105550100@sip.telnyx.com>",
			},
			expectError: false,
			expectedURI: "<sips:+15105550100@sip.telnyx.com>",
		},
		{
			name: "valid tel URI with brackets",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "<tel:+15105550100>",
			},
			expectError: false,
			expectedURI: "tel:+15105550100",
		},
		{
			name: "invalid URI - http",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "http://example.com",
			},
			expectError: true,
		},
		{
			name: "invalid URI - mailto",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "mailto:test@example.com",
			},
			expectError: true,
		},
		{
			name: "invalid URI - plain text",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "just-a-phone-number",
			},
			expectError: true,
		},
		{
			name: "invalid URI - empty",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "",
			},
			expectError: true,
		},
		{
			name: "missing room name",
			req: &TransferSIPParticipantRequest{
				RoomName:            "",
				ParticipantIdentity: "participant1",
				TransferTo:          "sip:+15105550100@sip.telnyx.com",
			},
			expectError: true,
		},
		{
			name: "missing participant identity",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "",
				TransferTo:          "sip:+15105550100@sip.telnyx.com",
			},
			expectError: true,
		},
		{
			name: "invalid URI with brackets - http",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "<http://example.com>",
			},
			expectError: true,
		},
		{
			name: "complex sip URI with transport",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "sip:+15105550100@my-livekit-demo.pstn.twilio.com;transport=tcp",
			},
			expectError: false,
			expectedURI: "<sip:+15105550100@my-livekit-demo.pstn.twilio.com;transport=tcp>",
		},
		{
			name: "sip URI with user and parameters",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "sip:+15053128762@testdomain.com,123",
			},
			expectError: false,
			expectedURI: "<sip:+15053128762@testdomain.com,123>",
		},
		{
			name: "sip URI with multiple parameters",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "sip:user@example.com;transport=tcp;lr;ttl=60",
			},
			expectError: false,
			expectedURI: "<sip:user@example.com;transport=tcp;lr;ttl=60>",
		},
		{
			name: "sip URI with port",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "sip:+1234567890@example.com:5060",
			},
			expectError: false,
			expectedURI: "<sip:+1234567890@example.com:5060>",
		},
		{
			name: "tel URI with extension",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "tel:+1234567890;ext=123",
			},
			expectError: false,
			expectedURI: "tel:+1234567890;ext=123",
		},
		{
			name: "tel URI with parameters",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "tel:+1234567890;phone-context=example.com",
			},
			expectError: false,
			expectedURI: "tel:+1234567890;phone-context=example.com",
		},
		{
			name: "sip URI with headers",
			req: &TransferSIPParticipantRequest{
				RoomName:            "room1",
				ParticipantIdentity: "participant1",
				TransferTo:          "sip:user@example.com?subject=test&priority=urgent",
			},
			expectError: false,
			expectedURI: "<sip:user@example.com?subject=test&priority=urgent>",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.req.Validate()
			if c.expectError {
				require.Error(t, err, "Expected validation to fail")
			} else {
				require.NoError(t, err, "Expected validation to pass")
				require.Equal(t, c.expectedURI, c.req.TransferTo, "TransferTo should be RFC-compliant after validation")
			}
		})
	}
}

func TestGRPCStatus(t *testing.T) {
	e := &SIPStatus{Code: SIPStatusCode_SIP_STATUS_BUSY_HERE}
	st, ok := status.FromError(e)
	require.True(t, ok)
	require.Equal(t, codes.ResourceExhausted, st.Code())
	require.Equal(t, "sip status 486: BUSY_HERE", st.Message())
	details := st.Details()
	require.Len(t, details, 1)
	e2, ok := details[0].(*SIPStatus)
	require.True(t, ok)
	require.True(t, proto.Equal(e, e2))
}

func TestInboundTrunkUpdate(t *testing.T) {
	update := func(r *SIPInboundTrunkInfo, u *SIPInboundTrunkUpdate) *SIPInboundTrunkInfo {
		upd := &UpdateSIPInboundTrunkRequest{
			Action: &UpdateSIPInboundTrunkRequest_Update{
				Update: u,
			},
		}
		out, err := upd.Action.(UpdateSIPInboundTrunkRequestAction).Apply(r)
		require.NoError(t, err)
		require.True(t, r != out)
		return out
	}
	r := &SIPInboundTrunkInfo{
		Name:     "Test",
		Numbers:  []string{"T1", "T2"},
		Metadata: "test",
	}
	out := update(r, &SIPInboundTrunkUpdate{
		Name: new("Test2"),
		Numbers: &ListUpdate{
			Set: []string{"T3"},
		},
	})
	prototest.Equals(t, &SIPInboundTrunkInfo{
		Name:     "Test2",
		Numbers:  []string{"T3"},
		Metadata: "test",
		Media:    &SIPMediaConfig{},
	}, out)

	// Update of legacy MediaEncryption propagates to a new SIPMediaConfig.
	out = update(out, &SIPInboundTrunkUpdate{
		MediaEncryption: new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
	})
	prototest.Equals(t, &SIPInboundTrunkInfo{
		Name:            "Test2",
		Numbers:         []string{"T3"},
		Metadata:        "test",
		MediaEncryption: SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE,
		Media: &SIPMediaConfig{
			Encryption: new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
		},
	}, out)

	// Update of a new SIPMediaConfig downgrades to legacy MediaEncryption.
	out = update(out, &SIPInboundTrunkUpdate{
		Media: &SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		},
	})
	prototest.Equals(t, &SIPInboundTrunkInfo{
		Name:            "Test2",
		Numbers:         []string{"T3"},
		Metadata:        "test",
		MediaEncryption: SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW,
		Media: &SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		},
	}, out)

	r2 := proto.CloneOf(r)
	r2.Numbers = []string{"T4"}
	upd2 := &UpdateSIPInboundTrunkRequest{
		Action: &UpdateSIPInboundTrunkRequest_Replace{
			Replace: r2,
		},
	}

	out, err := upd2.Action.(UpdateSIPInboundTrunkRequestAction).Apply(r)
	require.NoError(t, err)
	require.True(t, r != out)
	require.True(t, r2 != out)
	prototest.Equals(t, r2, out)
}

func TestOutboundTrunkUpdate(t *testing.T) {
	update := func(r *SIPOutboundTrunkInfo, u *SIPOutboundTrunkUpdate) *SIPOutboundTrunkInfo {
		upd := &UpdateSIPOutboundTrunkRequest{
			Action: &UpdateSIPOutboundTrunkRequest_Update{
				Update: u,
			},
		}
		out, err := upd.Action.(UpdateSIPOutboundTrunkRequestAction).Apply(r)
		require.NoError(t, err)
		require.True(t, r != out)
		return out
	}
	r := &SIPOutboundTrunkInfo{
		Name:     "Test",
		Address:  "sip.example.com",
		Numbers:  []string{"T1", "T2"},
		Metadata: "test",
	}

	out := update(r, &SIPOutboundTrunkUpdate{
		Name:    new("Test2"),
		Address: new("sip2.example.com"),
		Numbers: &ListUpdate{
			Set: []string{"T3"},
		},
	})
	prototest.Equals(t, &SIPOutboundTrunkInfo{
		Name:     "Test2",
		Address:  "sip2.example.com",
		Numbers:  []string{"T3"},
		Metadata: "test",
		Media:    &SIPMediaConfig{},
	}, out)

	// Update of legacy MediaEncryption propagates to a new SIPMediaConfig.
	out = update(out, &SIPOutboundTrunkUpdate{
		MediaEncryption: new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
	})
	prototest.Equals(t, &SIPOutboundTrunkInfo{
		Name:            "Test2",
		Address:         "sip2.example.com",
		Numbers:         []string{"T3"},
		Metadata:        "test",
		MediaEncryption: SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE,
		Media: &SIPMediaConfig{
			Encryption: new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
		},
	}, out)

	// Update of a new SIPMediaConfig downgrades to legacy MediaEncryption.
	out = update(out, &SIPOutboundTrunkUpdate{
		Media: &SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		},
	})
	prototest.Equals(t, &SIPOutboundTrunkInfo{
		Name:            "Test2",
		Address:         "sip2.example.com",
		Numbers:         []string{"T3"},
		Metadata:        "test",
		MediaEncryption: SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW,
		Media: &SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		},
	}, out)

	r2 := proto.CloneOf(r)
	r2.Numbers = []string{"T4"}
	upd2 := &UpdateSIPOutboundTrunkRequest{
		Action: &UpdateSIPOutboundTrunkRequest_Replace{
			Replace: r2,
		},
	}

	out, err := upd2.Action.(UpdateSIPOutboundTrunkRequestAction).Apply(r)
	require.NoError(t, err)
	require.True(t, r != out)
	require.True(t, r2 != out)
	prototest.Equals(t, r2, out)
}

func TestDispatchRuleUpdate(t *testing.T) {
	update := func(r *SIPDispatchRuleInfo, u *SIPDispatchRuleUpdate) *SIPDispatchRuleInfo {
		upd := &UpdateSIPDispatchRuleRequest{
			Action: &UpdateSIPDispatchRuleRequest_Update{
				Update: u,
			},
		}
		out, err := upd.Action.(UpdateSIPDispatchRuleRequestAction).Apply(r)
		require.NoError(t, err)
		require.True(t, r != out)
		return out
	}
	r := &SIPDispatchRuleInfo{
		Name:     "Test",
		TrunkIds: []string{"T1", "T2"},
		Rule: &SIPDispatchRule{
			Rule: &SIPDispatchRule_DispatchRuleDirect{
				DispatchRuleDirect: &SIPDispatchRuleDirect{RoomName: "test"},
			},
		},
	}

	out := update(r, &SIPDispatchRuleUpdate{
		Name: new("Test2"),
		TrunkIds: &ListUpdate{
			Set: []string{"T3"},
		},
	})
	prototest.Equals(t, &SIPDispatchRuleInfo{
		Name:     "Test2",
		TrunkIds: []string{"T3"},
		Rule: &SIPDispatchRule{
			Rule: &SIPDispatchRule_DispatchRuleDirect{
				DispatchRuleDirect: &SIPDispatchRuleDirect{RoomName: "test"},
			},
		},
		Media: &SIPMediaConfig{},
	}, out)

	// Update of legacy MediaEncryption propagates to a new SIPMediaConfig.
	out = update(out, &SIPDispatchRuleUpdate{
		MediaEncryption: new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
	})
	prototest.Equals(t, &SIPDispatchRuleInfo{
		Name:     "Test2",
		TrunkIds: []string{"T3"},
		Rule: &SIPDispatchRule{
			Rule: &SIPDispatchRule_DispatchRuleDirect{
				DispatchRuleDirect: &SIPDispatchRuleDirect{RoomName: "test"},
			},
		},
		MediaEncryption: SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE,
		Media: &SIPMediaConfig{
			Encryption: new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
		},
	}, out)

	// Update of a new SIPMediaConfig downgrades to legacy MediaEncryption.
	out = update(out, &SIPDispatchRuleUpdate{
		Media: &SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		},
	})
	prototest.Equals(t, &SIPDispatchRuleInfo{
		Name:     "Test2",
		TrunkIds: []string{"T3"},
		Rule: &SIPDispatchRule{
			Rule: &SIPDispatchRule_DispatchRuleDirect{
				DispatchRuleDirect: &SIPDispatchRuleDirect{RoomName: "test"},
			},
		},
		MediaEncryption: SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW,
		Media: &SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		},
	}, out)

	r2 := proto.CloneOf(r)
	r2.TrunkIds = []string{"T4"}
	upd2 := &UpdateSIPDispatchRuleRequest{
		Action: &UpdateSIPDispatchRuleRequest_Replace{
			Replace: r2,
		},
	}

	out, err := upd2.Action.(UpdateSIPDispatchRuleRequestAction).Apply(r)
	require.NoError(t, err)
	require.True(t, r != out)
	require.True(t, r2 != out)
	require.True(t, proto.Equal(r2, out))
}

func TestDestinationValidation(t *testing.T) {
	tests := []struct {
		name        string
		destination *Destination
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil destination should pass",
			destination: nil,
			expectError: false,
		},
		{
			name:        "empty destination should pass",
			destination: &Destination{},
			expectError: false,
		},
		{
			name: "valid country only",
			destination: &Destination{
				Country: "US",
			},
			expectError: false,
		},
		{
			name: "valid country and city",
			destination: &Destination{
				Country: "US",
				City:    "New York",
			},
			expectError: false,
		},
		{
			name: "valid country and region",
			destination: &Destination{
				Country: "US",
				Region:  "us-east-1",
			},
			expectError: false,
		},
		{
			name: "city without country or region should fail",
			destination: &Destination{
				City: "New York",
			},
			expectError: true,
			errorMsg:    "if city is specified, country or region must also be specified",
		},
		{
			name: "invalid country code - 3 letters",
			destination: &Destination{
				Country: "USA",
			},
			expectError: true,
			errorMsg:    "country must be a valid ISO 3166-1 alpha-2 code",
		},
		{
			name: "invalid country code - 1 letter",
			destination: &Destination{
				Country: "U",
			},
			expectError: true,
			errorMsg:    "country must be a valid ISO 3166-1 alpha-2 code",
		},
		{
			name: "invalid country code - XX (should fail with IsCountry check)",
			destination: &Destination{
				Country: "XX",
			},
			expectError: true,
			errorMsg:    "country must be a valid ISO 3166-1 alpha-2 code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.destination.Validate()

			if tt.expectError {
				require.Error(t, err)
				if tt.errorMsg != "" {
					require.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidationResultCombine(t *testing.T) {
	type testCase struct {
		name     string
		source   ValidationResult
		other    ValidationResult
		expected ValidationResult
	}

	sourceErr := errors.New("source error")
	otherErr := errors.New("other error")

	testCases := []testCase{
		{
			name:     "empty source and other",
			source:   ValidationResult{},
			other:    ValidationResult{},
			expected: ValidationResult{},
		},
		{
			name:     "empty source",
			source:   ValidationResult{nil, nil},
			other:    ValidationResult{otherErr, nil},
			expected: ValidationResult{otherErr, nil},
		},
		{
			name:     "empty other",
			source:   ValidationResult{sourceErr, nil},
			other:    ValidationResult{},
			expected: ValidationResult{sourceErr, nil},
		},
		{
			name:     "both non-empty",
			source:   ValidationResult{sourceErr, nil},
			other:    ValidationResult{otherErr, nil},
			expected: ValidationResult{errors.Join(errors.Join(sourceErr, otherErr)), nil},
		},
		{
			name:   "soft errors",
			source: ValidationResult{nil, []error{errors.New("soft error 1"), errors.New("soft error 2")}},
			other:  ValidationResult{nil, []error{errors.New("soft error 3"), errors.New("soft error 4")}},
			expected: ValidationResult{nil, []error{
				errors.New("soft error 1"),
				errors.New("soft error 2"),
				errors.New("soft error 3"),
				errors.New("soft error 4"),
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.source.Combine(testCase.other)
			if testCase.source.Error() != nil {
				require.ErrorContains(t, result.Error(), testCase.source.Error().Error())
			}
			if testCase.other.Error() != nil {
				require.ErrorContains(t, result.Error(), testCase.other.Error().Error())
			}
			require.Equal(t, result.SoftErrors(), testCase.expected.SoftErrors())
		})
	}
}
