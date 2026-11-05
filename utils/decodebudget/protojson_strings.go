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
	"encoding/json"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"

	"google.golang.org/protobuf/reflect/protoreflect"
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
// backslash or U+FFFD. It remembers the next backslash, and that no U+FFFD
// comes before fffd, so that each byte is searched once rather than once per
// occurrence.
type rebuildScan struct {
	in          []byte
	slash, fffd int
}

func (r *rebuildScan) next(from int) int {
	if r.slash < from {
		r.slash = len(r.in)
		if i := bytes.IndexByte(r.in[from:], '\\'); i >= 0 {
			r.slash = from + i
		}
	}
	if r.fffd < from {
		r.fffd = r.slash
		if i := bytes.Index(r.in[from:r.slash], []byte(replacementChar)); i >= 0 {
			r.fffd = from + i
		}
	}
	return min(r.slash, r.fffd)
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

// jsonScalarCost is what decoding a scalar token allocates beyond the copy the
// tokenizer already made: base64 bytes are decoded into a new buffer, a number
// in a string is tokenized again from a copy and parsed from another, and an
// integer's digits are normalized into a new string.
func jsonScalarCost(field protoreflect.FieldDescriptor, tok json.Token) int {
	cost := 0
	var lit string
	switch v := tok.(type) {
	case json.Number:
		lit = string(v)
	case string:
		switch field.Kind() {
		case protoreflect.BytesKind:
			return allocSize(len(v) * 3 / 4)
		case protoreflect.StringKind, protoreflect.BoolKind, protoreflect.EnumKind,
			protoreflect.MessageKind, protoreflect.GroupKind:
			return 0
		}
		lit, cost = v, 2*tmpCopySize(len(v))
	default:
		return 0
	}
	switch field.Kind() {
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return cost
	}
	return cost + jsonIntCost(lit)
}

// jsonIntCost is protojson normalizing an integer literal to its digits: the
// zeros an exponent adds are appended one at a time to a copy of the integer
// part, and the digits become a new string. Literals it rejects cost nothing.
func jsonIntCost(lit string) int {
	neg := strings.HasPrefix(lit, "-")
	intp, frac, exp, ok := numberParts(strings.TrimPrefix(lit, "-"))
	if !ok || (intp == "" && frac == "") {
		return 0
	}
	digits := len(intp) + exp
	if exp < 0 {
		if frac != "" || digits < 0 {
			return 0
		}
	} else if len(frac) > exp || digits > 20 {
		return 0
	}
	cost := intStringCost(digits, neg)
	if exp <= 0 {
		return cost
	}
	l, c := len(intp), len(intp)
	grow := func(add int) {
		mem, newCap := appendGrow(l, c, add, 1)
		cost, l, c = cost+mem, l+add, newCap
	}
	if frac != "" {
		grow(len(frac))
	}
	for range exp - len(frac) {
		grow(1)
	}
	return cost
}

// intStringCost is the string normalizeToIntString returns; Go keeps one-byte
// strings static, and builds a negative one from a stack copy.
func intStringCost(digits int, neg bool) int {
	switch {
	case neg:
		return allocSize(digits + 1)
	case digits <= 1:
		return 0
	}
	return allocSize(digits)
}

// numberParts splits an unsigned JSON number into its integer and fraction
// digits and its exponent.
func numberParts(lit string) (intp, frac string, exp int, ok bool) {
	end := strings.IndexAny(lit, ".eE")
	if end < 0 {
		return lit, "", 0, true
	}
	intp, lit = lit[:end], lit[end:]
	if strings.HasPrefix(lit, ".") {
		end = strings.IndexAny(lit, "eE")
		if end < 0 {
			end = len(lit)
		}
		frac, lit = lit[1:end], lit[end:]
	}
	if lit == "" {
		return intp, frac, 0, true
	}
	exp, err := strconv.Atoi(strings.TrimPrefix(lit[1:], "+"))
	return intp, frac, exp, err == nil
}
