package webhook

import "errors"

var (
	ErrNoAuthHeader    = errors.New("authorization header could not be found")
	ErrSecretNotFound  = errors.New("API secret could not be found")
	ErrInvalidChecksum = errors.New("could not verify authenticity of message")
)

const authHeader = "Authorization"

const (
	EventRoomStarted       = "room_started"
	EventRoomFinished      = "room_finished"
	EventParticipantJoined = "participant_joined"
	EventParticipantLeft   = "participant_left"
	EventTrackPublished    = "track_published"
	EventTrackUnpublished  = "track_unpublished"
	EventEgressStarted     = "egress_started"
	EventEgressUpdated     = "egress_updated"
	EventEgressEnded       = "egress_ended"
	EventIngressStarted    = "ingress_started"
	EventIngressEnded      = "ingress_ended"
)
