package agent

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"go.uber.org/multierr"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
)

func BuildAgentToken(
	apiKey, secret, roomName, participantIdentity, participantName, participantMetadata string,
	participantAttributes map[string]string,
	permissions *livekit.ParticipantPermission,
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
		SetValidFor(1 * time.Hour).
		SetMetadata(participantMetadata).
		SetAttributes(participantAttributes)

	return at.ToJWT()
}

type WorkerTokenConfig struct {
	Secret  string        `yaml:"secret,omitempty"`
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

var DefaultWorkerTokenConfig = WorkerTokenConfig{
	Timeout: 60 * time.Minute,
}

type WorkerClaims struct {
	jwt.Claims
}

type WorkerTokenProvider struct {
	nodeID  livekit.NodeID
	keys    [][]byte
	timeout time.Duration
}

func NewWorkerTokenProvider(nodeID livekit.NodeID, config WorkerTokenConfig) *WorkerTokenProvider {
	keys := bytes.Split([]byte(config.Secret), []byte(","))
	for i := range keys {
		keys[i] = bytes.TrimSpace(keys[i])
	}

	return &WorkerTokenProvider{
		nodeID:  nodeID,
		keys:    keys,
		timeout: config.Timeout,
	}
}

func (t *WorkerTokenProvider) Encode(claims WorkerClaims) (string, error) {
	signer, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.HS256,
		Key:       t.keys[0],
	}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create signer: %w", err)
	}

	now := time.Now()

	claims.Issuer = string(t.nodeID)
	claims.Expiry = jwt.NewNumericDate(now.Add(t.timeout))
	claims.NotBefore = jwt.NewNumericDate(now)
	claims.IssuedAt = jwt.NewNumericDate(now)

	token, err := jwt.Signed(signer).Claims(claims).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("failed to create signed jwt: %w", err)
	}
	return token, nil
}

func (t *WorkerTokenProvider) Decode(token string) (*WorkerClaims, error) {
	tok, err := jwt.ParseSigned(token)
	if err != nil {
		return nil, err
	}

	claims := &WorkerClaims{}
	for _, k := range t.keys {
		if erri := tok.Claims(k, &claims); erri != nil {
			err = multierr.Append(err, erri)
			continue
		}
		if erri := claims.Validate(jwt.Expected{Time: time.Now()}); erri != nil {
			err = multierr.Append(err, erri)
			continue
		}
		return claims, nil
	}
	return nil, err
}
