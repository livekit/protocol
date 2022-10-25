package utils

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/jxskiss/base62"
	"github.com/lithammer/shortuuid/v3"
)

const (
	RoomPrefix        = "RM_"
	NodePrefix        = "ND_"
	ParticipantPrefix = "PA_"
	TrackPrefix       = "TR_"
	APIKeyPrefix      = "API"
	EgressPrefix      = "EG_"
	IngressPrefix     = "IN_"
	RPCPrefix         = "RPC_"
)

func NewGuid(prefix string) string {
	return prefix + shortuuid.New()[:12]
}

// Creates a hashed ID from a unique string
func HashedID(id string) string {
	h := sha1.New()
	h.Write([]byte(id))
	val := h.Sum(nil)

	return base62.EncodeToString(val)
}

func LocalNodeID() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", NodePrefix, HashedID(hostname)[:8]), nil
}
