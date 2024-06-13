package rpc

import (
	"errors"
	"math/rand/v2"
	"strings"

	"github.com/livekit/protocol/livekit"
)

// NewCreateSIPParticipantRequest fills InternalCreateSIPParticipantRequest from
// livekit.CreateSIPParticipantRequest and livekit.SIPTrunkInfo.
func NewCreateSIPParticipantRequest(
	callID, wsUrl, token string,
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
	return &InternalCreateSIPParticipantRequest{
		SipCallId:           callID,
		Address:             trunk.Address,
		Transport:           trunk.Transport,
		Number:              outboundNumber,
		Username:            trunk.AuthUsername,
		Password:            trunk.AuthPassword,
		CallTo:              req.SipCallTo,
		WsUrl:               wsUrl,
		Token:               token,
		RoomName:            req.RoomName,
		ParticipantIdentity: req.ParticipantIdentity,
		ParticipantName:     req.ParticipantName,
		ParticipantMetadata: req.ParticipantMetadata,
		Dtmf:                req.Dtmf,
		PlayRingtone:        req.PlayRingtone,
	}, nil
}
