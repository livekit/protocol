package agent

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"

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
	Keys    string        `yaml:"secret,omitempty"`
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

var DefaultWorkerTokenConfig = WorkerTokenConfig{
	Timeout: 60 * time.Minute,
}

type WorkerClaimsT[T any] struct {
	jwt.Claims
	Metadata T `json:"metadata,inline"`
}

type WorkerTokenProviderT[T any] struct {
	nodeID  livekit.NodeID
	keySet  jose.JSONWebKeySet
	timeout time.Duration
}

func NewWorkerTokenProviderT[T any](nodeID livekit.NodeID, config WorkerTokenConfig) *WorkerTokenProviderT[T] {
	var keySet jose.JSONWebKeySet
	keys := bytes.Split([]byte(config.Keys), []byte(","))
	for i := range keys {
		key := bytes.TrimSpace(keys[i])
		id := sha256.Sum256(key)
		keySet.Keys = append(keySet.Keys, jose.JSONWebKey{
			Key:       key,
			KeyID:     base64.RawStdEncoding.EncodeToString(id[:8]),
			Algorithm: string(jose.HS256),
			Use:       "sig",
		})
	}

	return &WorkerTokenProviderT[T]{
		nodeID:  nodeID,
		keySet:  keySet,
		timeout: config.Timeout,
	}
}

func (t *WorkerTokenProviderT[T]) Encode(claims WorkerClaimsT[T]) (string, error) {
	opts := &jose.SignerOptions{}
	opts.WithType("JWT")
	opts.WithHeader("kid", t.keySet.Keys[0].KeyID)

	signer, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.HS256,
		Key:       t.keySet.Keys[0],
	}, opts)
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

func (t *WorkerTokenProviderT[T]) Decode(token string) (*WorkerClaimsT[T], error) {
	tok, err := jwt.ParseSigned(token)
	if err != nil {
		return nil, err
	}

	claims := &WorkerClaimsT[T]{}
	if err := tok.Claims(t.keySet, &claims); err != nil {
		return nil, err
	}
	if err := claims.Validate(jwt.Expected{Time: time.Now()}); err != nil {
		return nil, err
	}
	return claims, nil
}

type (
	empty               struct{}
	WorkerClaims        = WorkerClaimsT[empty]
	WorkerTokenProvider = WorkerTokenProviderT[empty]
)

func NewWorkerTokenProvider(nodeID livekit.NodeID, config WorkerTokenConfig) *WorkerTokenProvider {
	return NewWorkerTokenProviderT[empty](nodeID, config)
}
