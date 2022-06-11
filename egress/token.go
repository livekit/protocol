package egress

import (
	"time"

	"github.com/livekit/protocol/auth"
)

func BuildEgressToken(egressID, apiKey, secret, roomName string) (string, error) {
	f := false
	t := true
	grant := &auth.VideoGrant{
		RoomJoin:       true,
		Room:           roomName,
		CanSubscribe:   &t,
		CanPublish:     &f,
		CanPublishData: &f,
		Hidden:         true,
		Recorder:       true,
	}

	at := auth.NewAccessToken(apiKey, secret).
		AddGrant(grant).
		SetIdentity(egressID).
		SetValidFor(24 * time.Hour)

	return at.ToJWT()
}
