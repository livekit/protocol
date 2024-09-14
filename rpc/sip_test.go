package rpc

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestNewCreateSIPParticipantRequest(t *testing.T) {
	r := &livekit.CreateSIPParticipantRequest{
		SipCallTo:           "+3333",
		RoomName:            "room",
		ParticipantIdentity: "",
		ParticipantName:     "",
		ParticipantMetadata: "meta",
		ParticipantAttributes: map[string]string{
			"extra": "1",
		},
		Dtmf:         "1234#",
		PlayRingtone: true,
	}
	tr := &livekit.SIPOutboundTrunkInfo{
		SipTrunkId:   "trunk",
		Address:      "sip.example.com",
		Numbers:      []string{"+1111"},
		AuthUsername: "user",
		AuthPassword: "pass",
	}
	res, err := NewCreateSIPParticipantRequest("p_123", "call-id", "xyz.sip.livekit.cloud", "url", "token", r, tr)
	require.NoError(t, err)
	require.Equal(t, &InternalCreateSIPParticipantRequest{
		ProjectId:           "p_123",
		SipCallId:           "call-id",
		SipTrunkId:          "trunk",
		Address:             "sip.example.com",
		Hostname:            "xyz.sip.livekit.cloud",
		Number:              "+1111",
		CallTo:              "+3333",
		Username:            "user",
		Password:            "pass",
		RoomName:            "room",
		ParticipantMetadata: "meta",
		Token:               "token",
		WsUrl:               "url",
		Dtmf:                "1234#",
		PlayRingtone:        true,
		ParticipantAttributes: map[string]string{
			"extra":                    "1",
			livekit.AttrSIPCallID:      "call-id",
			livekit.AttrSIPTrunkID:     "trunk",
			livekit.AttrSIPTrunkNumber: "+1111",
			livekit.AttrSIPPhoneNumber: "+3333",
		},
	}, res)

	r.HidePhoneNumber = true
	res, err = NewCreateSIPParticipantRequest("p_123", "call-id", "xyz.sip.livekit.cloud", "url", "token", r, tr)
	require.NoError(t, err)
	require.Equal(t, &InternalCreateSIPParticipantRequest{
		ProjectId:           "p_123",
		SipCallId:           "call-id",
		SipTrunkId:          "trunk",
		Address:             "sip.example.com",
		Hostname:            "xyz.sip.livekit.cloud",
		Number:              "+1111",
		CallTo:              "+3333",
		Username:            "user",
		Password:            "pass",
		RoomName:            "room",
		ParticipantMetadata: "meta",
		Token:               "token",
		WsUrl:               "url",
		Dtmf:                "1234#",
		PlayRingtone:        true,
		ParticipantAttributes: map[string]string{
			"extra":                "1",
			livekit.AttrSIPCallID:  "call-id",
			livekit.AttrSIPTrunkID: "trunk",
		},
	}, res)
}
