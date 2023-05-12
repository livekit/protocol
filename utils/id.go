package utils

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"

	"github.com/jxskiss/base62"
	"github.com/lithammer/shortuuid/v4"

	"github.com/livekit/protocol/livekit"
)

const (
	RoomPrefix         = "RM_"
	NodePrefix         = "ND_"
	ParticipantPrefix  = "PA_"
	TrackPrefix        = "TR_"
	APIKeyPrefix       = "API"
	EgressPrefix       = "EG_"
	IngressPrefix      = "IN_"
	RPCPrefix          = "RPC_"
	WHIPResourcePrefix = "WH_"
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

var ErrInvalidGuid = errors.New("invalid guid")

var b62Index = newB62Index()
var b62Chars = []byte(shortuuid.DefaultAlphabet)

func newB62Index() [256]byte {
	var index [256]byte
	for i := 0; i < len(shortuuid.DefaultAlphabet); i++ {
		index[shortuuid.DefaultAlphabet[i]] = byte(i)
	}
	return index
}

func guidPrefix[T livekit.Guid]() string {
	var id T
	switch any(id).(type) {
	case livekit.TrackID:
		return TrackPrefix
	case livekit.ParticipantID:
		return ParticipantPrefix
	case livekit.RoomID:
		return RoomPrefix
	case livekit.NodeID:
		return NodePrefix
	default:
		panic("unreachable")
	}
}

func MarshalGuid[T livekit.Guid](id T) livekit.GuidBlock {
	var b livekit.GuidBlock
	prefix := guidPrefix[T]()
	idb := []byte(id)[len(prefix):]
	for i := 0; i < 12; i += 4 {
		j := (i * 6) >> 3
		b[j] |= b62Index[idb[i]]<<2 | b62Index[idb[i+1]]>>4
		b[j+1] |= b62Index[idb[i+1]]<<4 | b62Index[idb[i+2]]>>2
		b[j+2] |= b62Index[idb[i+2]]<<6 | b62Index[idb[i+3]]
	}
	return b
}

func UnmarshalGuid[T livekit.Guid](b livekit.GuidBlock) string {
	prefix := guidPrefix[T]()
	id := make([]byte, len(prefix)+12)
	copy(id, []byte(prefix))
	idb := id[len(prefix):]
	for i := 0; i < 3; i++ {
		j := i * 3
		k := i * 4
		idb[k] = b62Chars[b[j]>>2]
		idb[k+1] = b62Chars[(b[j]&3)<<4|b[j+1]>>4]
		idb[k+2] = b62Chars[(b[j+1]&15)<<2|b[j+2]>>6]
		idb[k+3] = b62Chars[b[j+2]&63]
	}
	return string(id)
}
