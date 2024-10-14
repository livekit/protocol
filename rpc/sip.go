package rpc

import (
	"errors"
	"maps"
	"math/rand/v2"
	"strings"

	"github.com/livekit/protocol/livekit"
)

// NewCreateSIPParticipantRequest fills InternalCreateSIPParticipantRequest from
// livekit.CreateSIPParticipantRequest and livekit.SIPTrunkInfo.
func NewCreateSIPParticipantRequest(
	projectID, callID, host, wsUrl, token string,
	req *livekit.CreateSIPParticipantRequest,
	trunk *livekit.SIPOutboundTrunkInfo,
) (*InternalCreateSIPParticipantRequest, error) {
	if len(trunk.Numbers) == 0 {
		return nil, errors.New("no numbers on outbound trunk")
	}
	outboundNumber := trunk.Numbers[rand.IntN(len(trunk.Numbers))]
	// A sanity check for the number format for well-known providers.
	switch {
	case strings.HasSuffix(trunk.Address, "telnyx.com"):
		// Telnyx omits leading '+' by default.
		outboundNumber = strings.TrimPrefix(outboundNumber, "+")
	case strings.HasSuffix(trunk.Address, "twilio.com"):
		// Twilio requires leading '+'.
		if !strings.HasPrefix(outboundNumber, "+") {
			outboundNumber = "+" + outboundNumber
		}
	}
	attrs := maps.Clone(req.ParticipantAttributes)
	if attrs == nil {
		attrs = make(map[string]string)
	}
	attrs[livekit.AttrSIPCallID] = callID
	trunkID := req.SipTrunkId
	if trunkID == "" {
		trunkID = trunk.SipTrunkId
	}
	attrs[livekit.AttrSIPTrunkID] = trunkID
	if !req.HidePhoneNumber {
		attrs[livekit.AttrSIPPhoneNumber] = req.SipCallTo
		attrs[livekit.AttrSIPTrunkNumber] = outboundNumber
	}
	return &InternalCreateSIPParticipantRequest{
		ProjectId:             projectID,
		SipCallId:             callID,
		SipTrunkId:            trunkID,
		Address:               trunk.Address,
		Hostname:              host,
		Transport:             trunk.Transport,
		Number:                outboundNumber,
		Username:              trunk.AuthUsername,
		Password:              trunk.AuthPassword,
		CallTo:                req.SipCallTo,
		WsUrl:                 wsUrl,
		Token:                 token,
		RoomName:              req.RoomName,
		ParticipantIdentity:   req.ParticipantIdentity,
		ParticipantName:       req.ParticipantName,
		ParticipantMetadata:   req.ParticipantMetadata,
		ParticipantAttributes: attrs,
		Dtmf:                  req.Dtmf,
		PlayRingtone:          req.PlayRingtone,
		Headers:               trunk.Headers,
		HeadersToAttributes:   trunk.HeadersToAttributes,
		RingingTimeout:        req.RingingTimeout,
		MaxCallDuration:       req.MaxCallDuration,
	}, nil
}
