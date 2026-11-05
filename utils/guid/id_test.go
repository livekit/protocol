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

package guid

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestMarshalUnmarshal(t *testing.T) {
	id0 := livekit.TrackID(New(TrackPrefix))
	b0 := Marshal(id0)
	id1 := Unmarshal[livekit.TrackID](b0)
	b1 := Marshal(id1)
	require.EqualValues(t, id0, id1)
	require.EqualValues(t, b0, b1)
}

func TestMarshalAppend(t *testing.T) {
	id0 := livekit.RoomID(New(RoomPrefix))
	id1 := livekit.RoomID(New(RoomPrefix))
	var b []byte
	b = MarshalAppend(b, id0)
	b = MarshalAppend(b, id1)
	require.Equal(t, id0, Unmarshal[livekit.RoomID](livekit.GuidBlock(b[0:])))
	require.Equal(t, id1, Unmarshal[livekit.RoomID](livekit.GuidBlock(b[9:])))
}

func TestMarshalInvalidID(t *testing.T) {
	for _, id := range []livekit.TrackID{
		"",
		TrackPrefix,
		TrackPrefix + "abc",
		TrackPrefix + "01IOl01IOl01",
		"TR_ABCDEFGHIJ",
	} {
		require.Equal(t, invalidBlock(), Marshal(id), id)
		require.Equal(t, []byte{0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, MarshalAppend([]byte{0}, id), id)
		require.Equal(t, livekit.TrackID(""), Unmarshal[livekit.TrackID](Marshal(id)), id)
	}
}

func TestUnmarshalInvalidBlock(t *testing.T) {
	for _, b := range []livekit.GuidBlock{
		{0xff},
		{0x03, 0x90},
		{0, 0, 57},
		{0, 0, 0, 0, 0, 0, 0xe4},
		{0, 0, 0, 0, 0, 0, 0, 0, 0x39},
		invalidBlock(),
	} {
		require.Equal(t, livekit.TrackID(""), Unmarshal[livekit.TrackID](b), b)
	}
	require.Equal(t, livekit.TrackID(TrackPrefix+"22222222222z"), Unmarshal[livekit.TrackID](livekit.GuidBlock{0, 0, 0, 0, 0, 0, 0, 0, 0x38}))
	require.Equal(t, livekit.TrackID(TrackPrefix+"222222222222"), Unmarshal[livekit.TrackID](livekit.GuidBlock{}))
}

func FuzzUnmarshal(f *testing.F) {
	f.Add(make([]byte, 9))
	f.Add(bytes.Repeat([]byte{0xff}, 9))
	f.Add([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0x38})
	f.Fuzz(func(t *testing.T, data []byte) {
		var b livekit.GuidBlock
		copy(b[:], data)
		id := Unmarshal[livekit.ParticipantID](b)
		if id != "" {
			require.Equal(t, b, Marshal(id))
		}
	})
}

func FuzzMarshal(f *testing.F) {
	f.Add("")
	f.Add("RM_abc")
	f.Add(New(RoomPrefix))
	f.Fuzz(func(t *testing.T, s string) {
		b := Marshal(livekit.RoomID(s))
		if b == invalidBlock() {
			return
		}
		require.Equal(t, livekit.RoomID(RoomPrefix+s[len(s)-Size:]), Unmarshal[livekit.RoomID](b))
	})
}

func TestHash(t *testing.T) {
	id := Hash(AgentPrefix, []byte("test"))
	require.Equal(t, "A_SFo4igEG5Dg5", id)
}

func BenchmarkNew(b *testing.B) {
	b.Run("new", func(b *testing.B) {
		var guid string
		for i := 0; i < b.N; i++ {
			guid = New(TrackPrefix)
		}
		_ = guid
	})
}

func TestIsValidID(t *testing.T) {
	require.True(t, IsValidID("A_SFo4igEG5Dg5"))
	require.True(t, IsValidID("NM_OJOHANNESBURG1A_K6SMQw2ZCZyB"))
	require.False(t, IsValidID("A_A_A_SFo4igEG5Dg5"))
	require.False(t, IsValidID("_A_SFo4igEG5Dg5"))
	require.False(t, IsValidID("_SFo4igEG5Dg5"))
	require.False(t, IsValidID("SFo4igEG5Dg5"))
}
