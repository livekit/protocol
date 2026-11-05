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

package decodebudget

import (
	"bytes"
	"strconv"
	"unicode/utf16"
	"unicode/utf8"
)

// jsonStringCost is what protojson's tokenizer allocates for the string token
// in raw, which may start with separators: one copy of the decoded string, and
// for a string it has to rebuild, the buffers its appends grow through.
func jsonStringCost(raw []byte, decodedLen int) int {
	in := raw[bytes.IndexByte(raw, '"')+1 : len(raw)-1]
	r := rebuildScan{in: in, slash: -1, fffd: -1}
	pos := r.next(0)
	if pos == len(in) {
		return allocSize(decodedLen)
	}
	cost, l, c := 0, pos, pos
	grow := func(add int) {
		mem, newCap := appendGrow(l, c, add, 1)
		cost, l, c = cost+mem, l+add, newCap
	}
	for pos < len(in) {
		if in[pos] == '\\' {
			n, size := jsonEscapeLen(in[pos:])
			grow(n)
			pos += size
			continue
		}
		_, n := utf8.DecodeRune(in[pos:])
		end := r.next(pos + n)
		grow(end - pos)
		pos = end
	}
	return cost + allocSize(decodedLen)
}

const replacementChar = string(utf8.RuneError)

// rebuildScan finds where protojson's parseString stops aliasing its input: a
// backslash or U+FFFD. It remembers the next of each so that a string of
// either is scanned once, not once per occurrence.
type rebuildScan struct {
	in          []byte
	slash, fffd int
}

func (r *rebuildScan) next(from int) int {
	if r.slash < from {
		r.slash = indexFrom(r.in, from, func(b []byte) int { return bytes.IndexByte(b, '\\') })
	}
	if r.fffd < from {
		r.fffd = indexFrom(r.in, from, func(b []byte) int { return bytes.Index(b, []byte(replacementChar)) })
	}
	return min(r.slash, r.fffd)
}

func indexFrom(in []byte, from int, index func([]byte) int) int {
	if i := index(in[from:]); i >= 0 {
		return from + i
	}
	return len(in)
}

// jsonEscapeLen returns the decoded length and raw size of the escape at the
// start of in.
func jsonEscapeLen(in []byte) (n, size int) {
	if len(in) < 6 || in[1] != 'u' {
		return 1, min(2, len(in))
	}
	r := jsonHex4(in[2:6])
	if utf16.IsSurrogate(r) && len(in) >= 12 && in[6] == '\\' && in[7] == 'u' {
		return utf8.RuneLen(utf16.DecodeRune(r, jsonHex4(in[8:12]))), 12
	}
	return max(utf8.RuneLen(r), 1), 6
}

func jsonHex4(b []byte) rune {
	v, _ := strconv.ParseUint(string(b), 16, 16)
	return rune(v)
}
