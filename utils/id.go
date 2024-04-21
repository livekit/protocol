// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jxskiss/base62"
	"github.com/lithammer/shortuuid/v4"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

const GuidSize = 12

const (
	RoomPrefix            = "RM_"
	NodePrefix            = "ND_"
	ParticipantPrefix     = "PA_"
	TrackPrefix           = "TR_"
	APIKeyPrefix          = "API"
	EgressPrefix          = "EG_"
	IngressPrefix         = "IN_"
	SIPTrunkPrefix        = "ST_"
	SIPDispatchRulePrefix = "SDR_"
	RPCPrefix             = "RPC_"
	WHIPResourcePrefix    = "WH_"
	RTMPResourcePrefix    = "RT_"
	URLResourcePrefix     = "UR_"
	AgentWorkerPrefix     = "AW_"
	AgentJobPrefix        = "AJ_"
)

var defaultGuidGenerator = newGuidGenerator(4096, GuidSize+10)

func NewGuid(prefix string) string {
	return defaultGuidGenerator.NewGuid(prefix)
}

// HashedID creates a hashed ID from a unique string
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

var b57Index = newB57Index()
var b57Chars = []byte(shortuuid.DefaultAlphabet)

func newB57Index() [256]byte {
	var index [256]byte
	for i := 0; i < len(b57Chars); i++ {
		index[b57Chars[i]] = byte(i)
	}
	return index
}

type guidGenerator struct {
	pool sync.Pool
	mu   sync.Mutex
	buf  []byte
	pos  int
}

func newGuidGenerator(bufSize, scratchSize int) *guidGenerator {
	return &guidGenerator{
		pool: sync.Pool{
			New: func() any {
				b := make([]byte, scratchSize)
				return &b
			},
		},
		buf: make([]byte, bufSize),
		pos: bufSize,
	}
}

func (g *guidGenerator) refillBuf() {
	for {
		_, err := rand.Read(g.buf)
		if err == nil {
			return
		}

		logger.Errorw("unable to refill guid buffer", err)
		time.Sleep(time.Millisecond)
	}
}

func (g *guidGenerator) uint32() uint32 {
	if g.pos == len(g.buf) {
		g.refillBuf()
		g.pos = 0
	}

	n := binary.BigEndian.Uint32(g.buf[g.pos:])
	g.pos += 4

	return n
}

func (g *guidGenerator) readIDChars(b []byte) {
	g.mu.Lock()
	defer g.mu.Unlock()

	var n int
	for {
		r := g.uint32()
		for i := 0; i < 5; i++ {
			if int(r&0x3f) < len(b57Chars) {
				b[n] = b57Chars[r&0x3f]
				n++
				if n == len(b) {
					return
				}
			}
			r >>= 6
		}
	}
}

func (g *guidGenerator) NewGuid(prefix string) string {
	b := g.pool.Get().(*[]byte)
	defer g.pool.Put(b)

	*b = append((*b)[:0], make([]byte, len(prefix)+GuidSize)...)
	copy(*b, prefix)
	g.readIDChars((*b)[len(prefix):])
	return string(*b)
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
	default:
		panic("unreachable")
	}
}

func MarshalGuid[T livekit.Guid](id T) livekit.GuidBlock {
	var b livekit.GuidBlock
	idb := []byte(id)[len(id)-GuidSize:]
	for i := 0; i < 3; i++ {
		j := i * 3
		k := i * 4
		b[j] = b57Index[idb[k]]<<2 | b57Index[idb[k+1]]>>4
		b[j+1] = b57Index[idb[k+1]]<<4 | b57Index[idb[k+2]]>>2
		b[j+2] = b57Index[idb[k+2]]<<6 | b57Index[idb[k+3]]
	}
	return b
}

func UnmarshalGuid[T livekit.Guid](b livekit.GuidBlock) T {
	prefix := guidPrefix[T]()
	id := make([]byte, len(prefix)+GuidSize)
	copy(id, []byte(prefix))
	idb := id[len(prefix):]
	for i := 0; i < 3; i++ {
		j := i * 3
		k := i * 4
		idb[k] = b57Chars[b[j]>>2]
		idb[k+1] = b57Chars[(b[j]&3)<<4|b[j+1]>>4]
		idb[k+2] = b57Chars[(b[j+1]&15)<<2|b[j+2]>>6]
		idb[k+3] = b57Chars[b[j+2]&63]
	}
	return T(id)
}
