package rpc

import (
	"strings"

	"github.com/livekit/protocol/livekit"
)

// NewCreateSIPParticipantRequest fills InternalCreateSIPParticipantRequest from
// livekit.CreateSIPParticipantRequest and livekit.SIPTrunkInfo.
func NewCreateSIPParticipantRequest(
	callID, wsUrl, token string,
	req *livekit.CreateSIPParticipantRequest,
	trunk *livekit.SIPTrunkInfo,
) *InternalCreateSIPParticipantRequest {
	// A sanity check for the number format for well-known providers.
	outboundNumber := trunk.OutboundNumber
	switch {
	case strings.HasSuffix(trunk.OutboundAddress, "telnyx.com"):
		// Telnyx omits leading '+' by default.
		outboundNumber = strings.TrimPrefix(outboundNumber, "+")
	case strings.HasSuffix(trunk.OutboundAddress, "twilio.com"):
		// Twilio requires leading '+'.
		if !strings.HasPrefix(outboundNumber, "+") {
			outboundNumber = "+" + outboundNumber
		}
	}
	return &InternalCreateSIPParticipantRequest{
		SipCallId:           callID,
		Address:             trunk.OutboundAddress,
		Transport:           trunk.Transport,
		Number:              outboundNumber,
		Username:            trunk.OutboundUsername,
		Password:            trunk.OutboundPassword,
		CallTo:              req.SipCallTo,
		WsUrl:               wsUrl,
		Token:               token,
		RoomName:            req.RoomName,
		ParticipantIdentity: req.ParticipantIdentity,
		ParticipantName:     req.ParticipantName,
		ParticipantMetadata: req.ParticipantMetadata,
		Dtmf:                req.Dtmf,
		PlayRingtone:        req.PlayRingtone,
	}
}
