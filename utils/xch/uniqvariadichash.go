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
	"unsafe"

	"github.com/go-faster/city"
)

// SEE: https://github.com/ClickHouse/ClickHouse/src/AggregateFunctions/UniqVariadicHash.h

type UniqVariadicHash struct {
	n   int
	sum uint64
}

func (h *UniqVariadicHash) WriteString(v string) (int, error) {
	return h.Write(unsafe.Slice(unsafe.StringData(v), len(v)))
}

func (h *UniqVariadicHash) Write(v []byte) (int, error) {
	ch := city.CH64(v)
	if h.n++; h.n == 1 {
		h.sum = ch
	} else {
		h.sum = hash128to64(ch, h.sum)
	}
	return len(v), nil
}

func (h UniqVariadicHash) Sum() uint64 {
	return h.sum
}

func hash128to64(low, high uint64) uint64 {
	const mul = uint64(0x9ddfea08eb382d69)
	a := (low ^ high) * mul
	a ^= a >> 47
	b := (high ^ a) * mul
	b ^= b >> 47
	b *= mul
	return b
}
