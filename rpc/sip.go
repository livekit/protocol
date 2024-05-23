package rpc

import "github.com/livekit/protocol/livekit"

// NewCreateSIPParticipantRequest fills InternalCreateSIPParticipantRequest from
// livekit.CreateSIPParticipantRequest and livekit.SIPTrunkInfo.
func NewCreateSIPParticipantRequest(
	callID, wsUrl, token string,
	req *livekit.CreateSIPParticipantRequest,
	trunk *livekit.SIPTrunkInfo,
) *InternalCreateSIPParticipantRequest {
	return &InternalCreateSIPParticipantRequest{
		SipCallId:           callID,
		Address:             trunk.OutboundAddress,
		Transport:           trunk.Transport,
		Number:              trunk.OutboundNumber,
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
