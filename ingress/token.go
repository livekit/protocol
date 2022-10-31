package ingress

import (
	"time"

	"github.com/abdulhaseeb08/protocol/auth"
)

func BuildIngressToken(apiKey, secret, roomName, participantIdentity, participantName string) (string, error) {
	f := false
	t := true
	grant := &auth.VideoGrant{
		RoomJoin:     true,
		Room:         roomName,
		CanSubscribe: &f,
		CanPublish:   &t,
	}

	at := auth.NewAccessToken(apiKey, secret).
		AddGrant(grant).
		SetIdentity(participantIdentity).
		SetName(participantName).
		SetValidFor(24 * time.Hour)

	return at.ToJWT()
}
