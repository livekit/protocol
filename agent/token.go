package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
)

func BuildAgentToken(
	apiKey, secret, roomName, participantIdentity, participantName, participantMetadata string,
	participantAttributes map[string]string,
	permissions *livekit.ParticipantPermission,
	kindDetails ...livekit.ParticipantInfo_KindDetail,
) (string, error) {
	grant := &auth.VideoGrant{
		RoomJoin:             true,
		Agent:                true,
		Room:                 roomName,
		CanSubscribe:         &permissions.CanSubscribe,
		CanPublish:           &permissions.CanPublish,
		CanPublishData:       &permissions.CanPublishData,
		Hidden:               permissions.Hidden,
		CanUpdateOwnMetadata: &permissions.CanUpdateMetadata,
		CanSubscribeMetrics:  &permissions.CanSubscribeMetrics,
	}

	at := auth.NewAccessToken(apiKey, secret).
		SetVideoGrant(grant).
		SetIdentity(participantIdentity).
		SetName(participantName).
		SetKind(livekit.ParticipantInfo_AGENT).
		SetKindDetail(kindDetails...).
		SetValidFor(1 * time.Hour).
		SetMetadata(participantMetadata).
		SetAttributes(participantAttributes)

	return at.ToJWT()
}

// JobTokenValidFor is the lifetime of a token a job asks for with
// CreateJobTokenRequest.
const JobTokenValidFor = 10 * time.Minute

var ErrClaimNotAllowed = errors.New("claim not allowed")

// ParseJobTokenClaims reads the claims of CreateJobTokenRequest into typed
// grants. A field the grants do not know is dropped, so a client that sends a
// newer claim still gets a token, without that claim. Anything after the JSON
// object, or a claim a job token never carries, is an error.
func ParseJobTokenClaims(claims string) (*auth.ClaimGrants, error) {
	dec := json.NewDecoder(strings.NewReader(claims))
	grants := &auth.ClaimGrants{}
	if err := dec.Decode(grants); err != nil {
		return nil, err
	}
	if _, err := dec.Token(); err != io.EOF {
		return nil, errors.New("unexpected data after claims")
	}
	if err := checkJobTokenGrants(grants); err != nil {
		return nil, err
	}
	return grants, nil
}

// checkJobTokenGrants refuses the claims a job token never carries: sha256,
// which a webhook receiver accepts as the hash of a body it was sent, and
// video.agent, which registers an agent worker.
func checkJobTokenGrants(grants *auth.ClaimGrants) error {
	if grants.Sha256 != "" {
		return fmt.Errorf("%w: sha256", ErrClaimNotAllowed)
	}
	if grants.Video != nil && grants.Video.Agent {
		return fmt.Errorf("%w: video.agent", ErrClaimNotAllowed)
	}
	return nil
}

// BuildJobToken signs a token that carries grants and is valid for
// JobTokenValidFor. The token holds only the grants given; the issuer and the
// times are set here. expiresAt is at most a second before the token's exp.
// Grants with a claim a job token never carries are refused here too, whatever
// produced them.
func BuildJobToken(apiKey, secret string, grants *auth.ClaimGrants) (token string, expiresAt time.Time, err error) {
	if grants == nil {
		return "", time.Time{}, errors.New("no grants")
	}
	if err := checkJobTokenGrants(grants); err != nil {
		return "", time.Time{}, err
	}
	at := auth.NewAccessToken(apiKey, secret).SetValidFor(JobTokenValidFor)
	*at.GetGrants() = *grants.Clone()

	expiresAt = time.Now().Add(JobTokenValidFor)
	token, err = at.ToJWT()
	if err != nil {
		return "", time.Time{}, err
	}
	return token, expiresAt, nil
}
