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

package xch

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils/must"
)

var testInput = []string{
	"020458be-1192-4377-88c1-2f4b74086533",
	"037b537f-cced-4cfe-9818-e0c4b986f943",
	"03977371-3ef0-4663-b912-3b2512ecb280",
	"064def22-5215-434a-a9a2-67dd58185a4a",
	"0a40a17b-a212-4e1f-bba7-91d53dc97092",
	"100aa6f5-7d55-4f3f-ad39-08dd0490bf80",
	"11eb610a-8c98-414b-a7a6-c71c34b53ce1",
}

// state for uniq(value)
var expectedString = []string{
	"0001af7848d6",
	"0002af7848d6d4fcfc11",
	"0003af7848d6d4fcfc11d98a7548",
	"0004af7848d6d4fcfc11f5600dfdd98a7548",
	"0005af7848d6a6955a69d4fcfc11f5600dfdd98a7548",
	"0006af7848d6a6955a69d4fcfc11f5600dfdd98a7548decd1781",
	"0007af7848d6a6955a69d4fcfc11f5600dfdd98a7548ef255dfddecd1781",
}

// state for uniq(value, value)
var expectedStringString = []string{
	"0001bd5466a2",
	"00026c040cecbd5466a2",
	"00036c040cecdf97dde5bd5466a2",
	"0004cf59c1c16c040cecdf97dde5bd5466a2",
	"0005cf59c1c18f3133956c040cecdf97dde5bd5466a2",
	"0006cf59c1c18f3133956c040cecdf97dde5bd5466a21699a7d6",
	"0007cf59c1c14d0091f48f3133956c040cecdf97dde5bd5466a21699a7d6",
}

type testBuf [32]byte

func (t *testBuf) FormatInt(i int) string {
	b := strconv.AppendInt(t[:0], int64(i), 32)
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func TestUniqueHashSet(t *testing.T) {
	t.Run("encoded set matches ch generated value", func(t *testing.T) {
		for i := range testInput {
			hs := NewUniquesHashSet()
			for _, v := range testInput[:i+1] {
				hs.InsertString(v)
			}
			require.Equal(t, expectedString[i], hex.EncodeToString(must.Get(hs.MarshalBinary())))
		}
	})

	t.Run("encoded variadic set matches ch generated value", func(t *testing.T) {
		for i := range testInput {
			hs := NewUniquesHashSetVar()
			for _, v := range testInput[:i+1] {
				hs.InsertString(v, v)
			}
			require.Equal(t, expectedStringString[i], hex.EncodeToString(must.Get(hs.MarshalBinary())))
		}
	})

	t.Run("estimate is within 1%% of real cardinality", func(t *testing.T) {
		var b testBuf

		hash := NewUniquesHashSet()
		for i := 0; i <= 20; i++ {
			for j := range 1 << i {
				hash.InsertString(b.FormatInt(j))
			}
			require.InEpsilon(t, 1<<i, hash.Size(), 0.01, "expected %d got %d", 1<<i, hash.Size())
		}
	})

	t.Run("merged sets approximate union cardinality", func(t *testing.T) {
		var b testBuf

		// fill s0 with 0-750k
		s0 := NewUniquesHashSet()
		for i := 0; i < 750000; i++ {
			s0.InsertString(b.FormatInt(i))
		}

		// fill s1 with 250k-1M
		s1 := NewUniquesHashSet()
		for i := 250000; i < 1000000; i++ {
			s1.InsertString(b.FormatInt(i))
		}

		s2 := NewUniquesHashSet()
		s2.Merge(s0)
		s2.Merge(s1)

		require.InEpsilon(t, 1000000, s2.Size(), 0.01, "expected %d got %d", 1000000, s2.Size())
	})

	t.Run("marshal/unmarshal", func(t *testing.T) {
		var b testBuf

		s0 := NewUniquesHashSet()
		for i := range 100000 {
			s0.InsertString(b.FormatInt(i))
		}

		buf, err := s0.MarshalBinary()
		require.NoError(t, err)

		s1 := NewUniquesHashSet()
		err = s1.UnmarshalBinary(buf)
		require.NoError(t, err)

		require.Equal(t, s0.Size(), s1.Size())
	})
}

func BenchmarkHashSet(b *testing.B) {
	var temp [8]byte

	hash := NewUniquesHashSet()
	for i := range b.N {
		for j := range 1000 {
			binary.LittleEndian.PutUint64(temp[:], (uint64(i)<<8)|uint64(j))
			hash.InsertString(unsafe.String(&temp[0], 8))
		}
	}
}
