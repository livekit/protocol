// Copyright 2026 LiveKit, Inc.
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

package datatrack_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/datatrack"
	"github.com/livekit/protocol/datatrack/datatracktest"
)

// Mutation-based fuzzing for packet deserialization.
//
// Run with:
//
//	go test -fuzz FuzzPacketUnmarshal ./datatrack
//
// A plain `go test` only exercises the seed corpus below.
func FuzzPacketUnmarshal(f *testing.F) {
	// Fixed vectors from packet_test.go: no extensions, participant SID
	// extension, and extension with padding.
	f.Add([]byte{
		0x18, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
		0xde, 0xad, 0xbe, 0xef, 0xff, 0xfe, 0xfd, 0xfc,
		0xfb, 0xfa,
	})
	f.Add([]byte{
		0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
		0xde, 0xad, 0xbe, 0xef, 0x00, 0x04, 0x01, 0x10,
		0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x72,
		0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
		0xff, 0xfe, 0xfd, 0xfc,
	})
	f.Add([]byte{
		0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
		0xde, 0xad, 0xbe, 0xef, 0x00, 0x03, 0x01, 0x0b,
		0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
		0x61, 0x6e, 0x74, 0x00, 0xff, 0xfe, 0xfd, 0xfc,
	})

	// Generated multi-packet frames with the participant SID extension.
	for _, raw := range datatracktest.GenerateRawDataPackets(1, 1, 1, 4, 600, 10*time.Millisecond) {
		f.Add(raw)
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		var p datatrack.Packet
		if err := p.Unmarshal(data); err != nil {
			// Malformed input must be rejected with an error, never a panic.
			return
		}

		// Round-trip stability: a packet that parsed must marshal, and the
		// re-parsed result must be structurally identical.
		buf, err := p.Marshal()
		require.NoError(t, err)

		var p2 datatrack.Packet
		require.NoError(t, p2.Unmarshal(buf))
		require.Equal(t, p, p2)
	})
}
