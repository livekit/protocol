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
	res, err := NewCreateSIPParticipantRequest("call-id", "url", "token", r, tr)
	require.NoError(t, err)
	require.Equal(t, &InternalCreateSIPParticipantRequest{
		SipCallId:           "call-id",
		Address:             "sip.example.com",
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
	res, err = NewCreateSIPParticipantRequest("call-id", "url", "token", r, tr)
	require.NoError(t, err)
	require.Equal(t, &InternalCreateSIPParticipantRequest{
		SipCallId:           "call-id",
		Address:             "sip.example.com",
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
