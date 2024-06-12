package rpc

import (
	"maps"

	"github.com/livekit/protocol/livekit"
)

// NewCreateSIPParticipantRequest fills InternalCreateSIPParticipantRequest from
// livekit.CreateSIPParticipantRequest and livekit.SIPTrunkInfo.
func NewCreateSIPParticipantRequest(
	callID, wsUrl, token string,
	req *livekit.CreateSIPParticipantRequest,
	trunk *livekit.SIPTrunkInfo,
) *InternalCreateSIPParticipantRequest {
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
		attrs[livekit.AttrSIPTrunkNumber] = trunk.OutboundNumber
	}
	return &InternalCreateSIPParticipantRequest{
		SipCallId:             callID,
		Address:               trunk.OutboundAddress,
		Transport:             trunk.Transport,
		Number:                trunk.OutboundNumber,
		Username:              trunk.OutboundUsername,
		Password:              trunk.OutboundPassword,
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
	}
}
