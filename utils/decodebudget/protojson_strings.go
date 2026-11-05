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
	i := indexRebuild(in)
	if i == len(in) {
		return allocSize(decodedLen)
	}
	cost, l, c := 0, i, i
	grow := func(add int) {
		mem, newCap := appendGrow(l, c, add, 1)
		cost, l, c = cost+mem, l+add, newCap
	}
	for in = in[i:]; len(in) > 0; {
		if in[0] == '\\' {
			n, size := jsonEscapeLen(in)
			grow(n)
			in = in[size:]
			continue
		}
		_, n := utf8.DecodeRune(in)
		n += indexRebuild(in[n:])
		grow(n)
		in = in[n:]
	}
	return cost + allocSize(decodedLen)
}

const replacementChar = string(utf8.RuneError)

// indexRebuild is where protojson's parseString stops aliasing its input.
func indexRebuild(in []byte) int {
	i := bytes.IndexByte(in, '\\')
	if i < 0 {
		i = len(in)
	}
	if j := bytes.Index(in[:i], []byte(replacementChar)); j >= 0 {
		return j
	}
	return i
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
