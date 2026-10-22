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

package datatrack

import (
	"bytes"
	"testing"

	"github.com/livekit/protocol/livekit"
	"github.com/stretchr/testify/require"
)

func TestPacket(t *testing.T) {
	t.Run("without extension", func(t *testing.T) {
		payload := make([]byte, 6)
		for i := range len(payload) {
			payload[i] = byte(255 - i)
		}
		packet := &Packet{
			Header: Header{
				Version:        0,
				IsStartOfFrame: true,
				IsFinalOfFrame: true,
				Handle:         3333,
				SequenceNumber: 6666,
				FrameNumber:    9999,
				Timestamp:      0xdeadbeef,
			},
			Payload: payload,
		}
		rawPacket, err := packet.Marshal()
		require.NoError(t, err)

		expectedRawPacket := []byte{
			0x18, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0xff, 0xfe, 0xfd, 0xfc,
			0xfb, 0xfa,
		}
		require.Equal(t, expectedRawPacket, rawPacket)

		var unmarshaled Packet
		err = unmarshaled.Unmarshal(rawPacket)
		require.NoError(t, err)
		require.Equal(t, packet, &unmarshaled)
	})

	t.Run("with extension", func(t *testing.T) {
		payload := make([]byte, 4)
		for i := range len(payload) {
			payload[i] = byte(255 - i)
		}
		packet := &Packet{
			Header: Header{
				Version:        0,
				IsStartOfFrame: true,
				IsFinalOfFrame: false,
				Handle:         3333,
				SequenceNumber: 6666,
				FrameNumber:    9999,
				Timestamp:      0xdeadbeef,
			},
			Payload: payload,
		}
		if extParticipantSid, err := NewExtensionParticipantSid("test_participant"); err == nil {
			if ext, err := extParticipantSid.Marshal(); err == nil {
				packet.AddExtension(ext)
			}
		}
		rawPacket, err := packet.Marshal()
		require.NoError(t, err)

		expectedRawPacket := []byte{
			0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0x00, 0x04, 0xff, 0x10,
			0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x72,
			0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
			0xff, 0xfe, 0xfd, 0xfc,
		}
		require.Equal(t, expectedRawPacket, rawPacket)

		var unmarshaled Packet
		err = unmarshaled.Unmarshal(rawPacket)
		require.NoError(t, err)
		require.Equal(t, packet, &unmarshaled)

		ext, err := unmarshaled.GetExtension(uint8(livekit.DataTrackExtensionID_DTEI_PARTICIPANT_SID))
		require.NoError(t, err)

		var extParticipantSid ExtensionParticipantSid
		require.NoError(t, extParticipantSid.Unmarshal(ext))
		require.Equal(t, livekit.ParticipantID("test_participant"), extParticipantSid.ParticipantID())
	})

	t.Run("with extension padding", func(t *testing.T) {
		payload := make([]byte, 4)
		for i := range len(payload) {
			payload[i] = byte(255 - i)
		}
		packet := &Packet{
			Header: Header{
				Version:        0,
				IsStartOfFrame: true,
				IsFinalOfFrame: false,
				Handle:         3333,
				SequenceNumber: 6666,
				FrameNumber:    9999,
				Timestamp:      0xdeadbeef,
			},
			Payload: payload,
		}
		if extParticipantSid, err := NewExtensionParticipantSid("participant"); err == nil {
			if ext, err := extParticipantSid.Marshal(); err == nil {
				packet.AddExtension(ext)
			}
		}
		rawPacket, err := packet.Marshal()
		require.NoError(t, err)

		expectedRawPacket := []byte{
			0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0x00, 0x03, 0xff, 0x0b,
			0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
			0x61, 0x6e, 0x74, 0x00, 0xff, 0xfe, 0xfd, 0xfc,
		}
		require.Equal(t, expectedRawPacket, rawPacket)

		var unmarshaled Packet
		err = unmarshaled.Unmarshal(rawPacket)
		require.NoError(t, err)
		require.Equal(t, packet, &unmarshaled)

		ext, err := unmarshaled.GetExtension(uint8(livekit.DataTrackExtensionID_DTEI_PARTICIPANT_SID))
		require.NoError(t, err)

		var extParticipantSid ExtensionParticipantSid
		require.NoError(t, extParticipantSid.Unmarshal(ext))
		require.Equal(t, livekit.ParticipantID("participant"), extParticipantSid.ParticipantID())
	})

	t.Run("replace extension", func(t *testing.T) {
		payload := make([]byte, 4)
		for i := range len(payload) {
			payload[i] = byte(255 - i)
		}
		packet := &Packet{
			Header: Header{
				Version:        0,
				IsStartOfFrame: true,
				IsFinalOfFrame: false,
				Handle:         3333,
				SequenceNumber: 6666,
				FrameNumber:    9999,
				Timestamp:      0xdeadbeef,
			},
			Payload: payload,
		}
		if extParticipantSid, err := NewExtensionParticipantSid("participant"); err == nil {
			if ext, err := extParticipantSid.Marshal(); err == nil {
				packet.AddExtension(ext)
			}
		}
		rawPacket, err := packet.Marshal()
		require.NoError(t, err)

		expectedRawPacket := []byte{
			0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0x00, 0x03, 0xff, 0x0b,
			0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
			0x61, 0x6e, 0x74, 0x00, 0xff, 0xfe, 0xfd, 0xfc,
		}
		require.Equal(t, expectedRawPacket, rawPacket)

		// replace existing extension ID and ensure that marshalled packet is updated
		if extParticipantSid, err := NewExtensionParticipantSid("test_participant"); err == nil {
			if ext, err := extParticipantSid.Marshal(); err == nil {
				packet.AddExtension(ext)
			}
		}
		rawPacket, err = packet.Marshal()
		require.NoError(t, err)

		expectedRawPacket = []byte{
			0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0x00, 0x04, 0xff, 0x10,
			0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x72,
			0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
			0xff, 0xfe, 0xfd, 0xfc,
		}
		require.Equal(t, expectedRawPacket, rawPacket)

		var unmarshaled Packet
		err = unmarshaled.Unmarshal(rawPacket)
		require.NoError(t, err)
		require.Equal(t, packet, &unmarshaled)

		ext, err := unmarshaled.GetExtension(uint8(livekit.DataTrackExtensionID_DTEI_PARTICIPANT_SID))
		require.NoError(t, err)

		var extParticipantSid ExtensionParticipantSid
		require.NoError(t, extParticipantSid.Unmarshal(ext))
		require.Equal(t, livekit.ParticipantID("test_participant"), extParticipantSid.ParticipantID())
	})

	t.Run("with client extensions", func(t *testing.T) {
		// mirrors the serialization test vector of the Rust client (livekit-datatrack)
		payload := bytes.Repeat([]byte{0xfa}, 1024)
		packet := &Packet{
			Header: Header{
				Version:        0,
				IsStartOfFrame: false,
				IsFinalOfFrame: true,
				Handle:         0x8811,
				SequenceNumber: 0x4422,
				FrameNumber:    0x4411,
				Timestamp:      0x44221188,
			},
			Payload: payload,
		}
		var iv [ExtensionE2EEIVLength]byte
		for i := range iv {
			iv[i] = 0x3c
		}
		ext, err := NewExtensionE2EE(0xfa, iv).Marshal()
		require.NoError(t, err)
		packet.AddExtension(ext)
		ext, err = NewExtensionUserTimestamp(0x4411221111118811).Marshal()
		require.NoError(t, err)
		packet.AddExtension(ext)

		rawPacket, err := packet.Marshal()
		require.NoError(t, err)
		require.Len(t, rawPacket, 1064)

		expectedHeader := []byte{
			0x0c, 0x00, 0x88, 0x11, 0x44, 0x22, 0x44, 0x11, // version 0, final, extensions; handle; sequence; frame
			0x44, 0x22, 0x11, 0x88, 0x00, 0x06, // timestamp; extension words
			0x01, 0x0d, 0xfa, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, // E2EE
			0x02, 0x08, 0x44, 0x11, 0x22, 0x11, 0x11, 0x11, 0x88, 0x11, // user timestamp
			0x00, // padding
		}
		require.Equal(t, expectedHeader, rawPacket[:len(expectedHeader)])
		require.Equal(t, payload, rawPacket[len(expectedHeader):])

		var unmarshaled Packet
		require.NoError(t, unmarshaled.Unmarshal(rawPacket))
		require.Equal(t, packet, &unmarshaled)

		ext, err = unmarshaled.GetExtension(ExtensionE2EEID)
		require.NoError(t, err)
		var extE2EE ExtensionE2EE
		require.NoError(t, extE2EE.Unmarshal(ext))
		require.Equal(t, uint8(0xfa), extE2EE.KeyIndex())
		require.Equal(t, iv, extE2EE.IV())

		ext, err = unmarshaled.GetExtension(ExtensionUserTimestampID)
		require.NoError(t, err)
		var extUserTimestamp ExtensionUserTimestamp
		require.NoError(t, extUserTimestamp.Unmarshal(ext))
		require.Equal(t, uint64(0x4411221111118811), extUserTimestamp.Timestamp())
	})

	t.Run("bad packet", func(t *testing.T) {
		var unmarshaled Packet
		// extensions size too small
		badPacket := []byte{
			0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0x00, 0x02, 0x01, 0x0b,
			0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
			0x61, 0x6e, 0x74, 0x00, 0xff, 0xfe, 0xfd, 0xfc,
		}
		err := unmarshaled.Unmarshal(badPacket)
		require.Error(t, err)

		// get an invalid extension id
		badPacket = []byte{
			0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0x00, 0x03, 0x02, 0x0b,
			0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
			0x61, 0x6e, 0x74, 0x00, 0xff, 0xfe, 0xfd, 0xfc,
		}
		err = unmarshaled.Unmarshal(badPacket)
		require.NoError(t, err)
		_, err = unmarshaled.GetExtension(uint8(livekit.DataTrackExtensionID_DTEI_PARTICIPANT_SID))
		require.Error(t, err)

		// extension payload size bigger than payload
		badPacket = []byte{
			0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0x00, 0x03, 0x01, 0x0d,
			0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
			0x61, 0x6e, 0x74, 0x00, 0xff, 0xfe, 0xfd, 0xfc,
		}
		err = unmarshaled.Unmarshal(badPacket)
		require.Error(t, err)

		// extension payload size smaller than payload
		badPacket = []byte{
			0x14, 0x00, 0x0d, 0x05, 0x1a, 0x0a, 0x27, 0x0f,
			0xde, 0xad, 0xbe, 0xef, 0x00, 0x03, 0x01, 0x07,
			0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
			0x61, 0x6e, 0x74, 0x00, 0xff, 0xfe, 0xfd, 0xfc,
		}
		err = unmarshaled.Unmarshal(badPacket)
		require.Error(t, err)
	})

	t.Run("oversized extension padding does not panic", func(t *testing.T) {
		var unmarshaled Packet
		// HasExtensions set, extensionsSize describes more bytes than present,
		// terminated by a 0x00 padding id -> hdrSize would exceed len(buf)
		badPacket := []byte{
			0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}
		err := unmarshaled.Unmarshal(badPacket)
		require.Error(t, err)
	})

	t.Run("extensions size wraparound does not panic", func(t *testing.T) {
		var unmarshaled Packet
		// 0xFFFF extensions-size field wraps (raw+1)*4 uint16 arithmetic to a huge
		// remainingSize; the 0x00 padding id must not push hdrSize past len(buf)
		badPacket := []byte{
			0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00,
		}
		err := unmarshaled.Unmarshal(badPacket)
		require.Error(t, err)
	})

	t.Run("truncated extensions size field does not panic", func(t *testing.T) {
		var unmarshaled Packet
		// HasExtensions set but buffer too short to hold the extensionsSize field
		badPacket := []byte{
			0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
		}
		err := unmarshaled.Unmarshal(badPacket)
		require.Error(t, err)
	})
}
