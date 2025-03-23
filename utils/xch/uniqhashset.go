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
	"bytes"
	"fmt"
	"hash/crc32"
	"math"
	"math/bits"
	"unsafe"

	"github.com/ClickHouse/ch-go/proto"
	"github.com/go-faster/city"
)

// SEE: https://github.com/ClickHouse/ClickHouse/src/AggregateFunctions/UniquesHashSet.h
// SEE: https://github.com/ClickHouse/ClickHouse/src/AggregateFunctions/AggregateFunctionUniq.h

const (
	initialSizeDegree     = 4
	maxSizeDegree         = 17
	maxSize               = 1 << (maxSizeDegree - 1)
	bitsForSkip           = 32 - maxSizeDegree
	initialBufferCapacity = 1 << initialSizeDegree
)

type UniquesHashSet struct {
	uniquesHashSetImpl
}

func NewUniquesHashSet() *UniquesHashSet {
	return &UniquesHashSet{makeUniquesHashSet()}
}

func (s *UniquesHashSet) InsertString(v string) {
	b := unsafe.Slice(unsafe.StringData(v), len(v))
	s.uniquesHashSetImpl.insert(intHash64(city.CH64(b)))
}

func (s *UniquesHashSet) Merge(rhs *UniquesHashSet) {
	s.uniquesHashSetImpl.merge(rhs.uniquesHashSetImpl)
}

type UniquesHashSetVar struct {
	uniquesHashSetImpl
}

func NewUniquesHashSetVar() *UniquesHashSetVar {
	return &UniquesHashSetVar{makeUniquesHashSet()}
}

func (s *UniquesHashSetVar) InsertString(v ...string) {
	var h UniqVariadicHash
	for i := range v {
		h.WriteString(v[i])
	}
	s.uniquesHashSetImpl.insert(h.Sum())
}

func (s *UniquesHashSetVar) Merge(rhs *UniquesHashSetVar) {
	s.uniquesHashSetImpl.merge(rhs.uniquesHashSetImpl)
}

type uniquesHashSetImpl struct {
	buf        []uint32
	sizeDegree uint8
	skipDegree uint8
	size       uint32
	hasZero    bool
}

func makeUniquesHashSet() uniquesHashSetImpl {
	return uniquesHashSetImpl{
		buf:        make([]uint32, initialBufferCapacity),
		sizeDegree: initialSizeDegree,
	}
}

func (s *uniquesHashSetImpl) insert(hash uint64) {
	if !s.good(uint32(hash)) {
		return
	}
	s.insertImpl(uint32(hash))
	s.shrinkIfNeeded()
}

func (s *uniquesHashSetImpl) insertImpl(hash uint32) {
	if hash == 0 {
		if !s.hasZero {
			s.hasZero = true
			s.size++
		}
		return
	}

	pos := s.place(hash)
	for s.buf[pos] != 0 && s.buf[pos] != hash {
		pos = (pos + 1) & s.mask()
	}
	if s.buf[pos] == hash {
		return
	}
	s.buf[pos] = hash
	s.size++
}

func (s *uniquesHashSetImpl) good(hash uint32) bool {
	return hash == ((hash >> s.skipDegree) << s.skipDegree)
}

func (s *uniquesHashSetImpl) place(hash uint32) int {
	return int((hash >> bitsForSkip) & uint32(s.mask()))
}

func (s *uniquesHashSetImpl) mask() int {
	return (1 << s.sizeDegree) - 1
}

func (s *uniquesHashSetImpl) shrinkIfNeeded() {
	if s.size <= uint32(1<<(s.sizeDegree-1)) {
		return
	}
	if s.size > maxSize {
		for s.size > maxSize {
			s.skipDegree++
			s.rehash()
		}
	} else {
		s.resize(s.sizeDegree + 1)
	}
}

func (s *uniquesHashSetImpl) resize(newDegree uint8) {
	oldBuf := s.buf

	if newDegree == 0 {
		newDegree++
	}

	s.buf = append(s.buf, make([]uint32, 1<<newDegree-len(oldBuf))...)
	s.sizeDegree = newDegree

	for i := 0; i < len(oldBuf) || s.buf[i] != 0; i++ {
		hash := s.buf[i]
		if hash == 0 {
			continue
		}

		place := s.place(hash)
		if place == i {
			continue
		}

		for s.buf[place] != 0 && s.buf[place] != hash {
			place = (place + 1) & s.mask()
		}

		if s.buf[place] == hash {
			continue
		}

		s.buf[place] = hash
		s.buf[i] = 0
	}
}

func (s *uniquesHashSetImpl) reinsert(hash uint32) {
	pos := s.place(hash)
	for s.buf[pos] != 0 {
		pos = (pos + 1) & s.mask()
	}
	s.buf[pos] = hash
}

func (s *uniquesHashSetImpl) rehash() {
	for i, hash := range s.buf {
		if hash == 0 {
			continue
		}

		if !s.good(hash) {
			s.buf[i] = 0
			s.size--
		} else if i != s.place(hash) {
			s.buf[i] = 0
			s.reinsert(hash)
		}
	}

	for i := 0; i < len(s.buf) && s.buf[i] != 0; i++ {
		if i != s.place(s.buf[i]) {
			hash := s.buf[i]
			s.buf[i] = 0
			s.reinsert(hash)
		}
	}
}

func (s *uniquesHashSetImpl) Reset() {
	clear(s.buf)
	s.buf = s.buf[:initialBufferCapacity]
	s.size = initialBufferCapacity
}

func (s *uniquesHashSetImpl) Size() uint64 {
	if s.skipDegree == 0 {
		return uint64(s.size)
	}
	res := uint64(s.size) << s.skipDegree
	res += uint64(crc32.ChecksumIEEE([]byte{byte(s.size)})) & ((1 << s.skipDegree) - 1)
	p32 := float64(uint64(1) << 32)
	return uint64(math.Round(p32 * (math.Log(p32) - math.Log(p32-float64(res)))))
}

func (s *uniquesHashSetImpl) merge(rhs uniquesHashSetImpl) {
	if rhs.skipDegree > s.skipDegree {
		s.skipDegree = rhs.skipDegree
		s.rehash()
	}

	if rhs.hasZero && !s.hasZero {
		s.hasZero = true
		s.size++
		s.shrinkIfNeeded()
	}

	for _, hash := range rhs.buf {
		if hash != 0 && s.good(hash) {
			s.insertImpl(hash)
			s.shrinkIfNeeded()
		}
	}
}

func (s *uniquesHashSetImpl) Encode(b *proto.Buffer) {
	b.PutUInt8(s.skipDegree)
	b.PutUVarInt(uint64(s.size))

	if s.hasZero {
		b.PutUInt32(0)
	}

	for _, val := range s.buf {
		if val != 0 {
			b.PutUInt32(val)
		}
	}
}

func (s *uniquesHashSetImpl) Decode(r *proto.Reader) error {
	sd, err := r.UInt8()
	if err != nil {
		return err
	}
	s.skipDegree = sd

	sz, err := r.UVarInt()
	if err != nil {
		return err
	}

	if sz > maxSize {
		return fmt.Errorf("set too large to read")
	}
	s.size = uint32(sz)

	newDegree := uint8(initialSizeDegree)
	if s.size > 1 {
		newDegree = uint8(bits.Len32(s.size-1)) + 1
	}
	s.buf = make([]uint32, 1<<newDegree)
	s.sizeDegree = newDegree
	s.hasZero = false

	for i := uint32(0); i < s.size; i++ {
		x, err := r.UInt32()
		if err != nil {
			return err
		}
		if x == 0 {
			s.hasZero = true
		} else {
			s.reinsert(x)
		}
	}

	return nil
}

func (s *uniquesHashSetImpl) MarshalBinary() ([]byte, error) {
	b := &proto.Buffer{}
	s.Encode(b)
	return b.Buf, nil
}

func (s *uniquesHashSetImpl) UnmarshalBinary(b []byte) error {
	return s.Decode(proto.NewReader(bytes.NewReader(b)))
}
