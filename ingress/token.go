package ingress

import (
	"time"

	"github.com/livekit/protocol/auth"
)

func BuildIngressToken(ingressID, apiKey, secret, roomName string) (string, error) {
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
		SetIdentity(ingressID).
		SetValidFor(24 * time.Hour)

	return at.ToJWT()
}
