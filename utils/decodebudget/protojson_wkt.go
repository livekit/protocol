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

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Allocations protojson makes for well-known types beyond the decoded messages,
// measured against protobuf-go v1.36.12.
const (
	// A Value boxes its scalar for the reflective oneof setter. A ListValue's
	// list is created even when empty; mapEntries charges a Struct's map.
	jsonKnownValueCost  = 24
	jsonKnownStructCost = 80
	jsonKnownListCost   = 24
	// A FieldMask path is split out, boxed and appended to the paths list.
	jsonFieldMaskPathCost = 80
	// An Any's object is tokenized a second time to find "@type", and the
	// embedded message is marshalled deterministically, sorting map keys as
	// reflect.Values. Its wire form can be several times its JSON: "0," in a
	// ListValue is eleven bytes.
	jsonAnyCost         = 32
	jsonAnyCopyFactor   = 6
	jsonAnyMapEntryCost = 80
	// A Timestamp whose offset is not a whole hour gets a new time.Location.
	jsonFixedZoneCost = 160
)

// wellKnown walks md if protojson gives it a special JSON form, reporting
// whether it did. Any and FieldMask are only charged here; their tokens are
// then walked by the caller when they have no resolvable form.
func (w *jsonWalk) wellKnown(md protoreflect.MessageDescriptor, tok json.Token, depth int) (bool, error) {
	switch md.FullName() {
	case "google.protobuf.Struct", "google.protobuf.ListValue":
		return true, w.knownContainer(md, tok, depth)
	case "google.protobuf.Value":
		return true, w.knownValue(md, tok, depth)
	case "google.protobuf.Any":
		if tok == json.Delim('{') {
			if err := w.alloc(jsonAnyCost); err != nil {
				return true, err
			}
			url := w.anyTypeURL()
			if url == "" {
				// Discarding unknown fields, the decoder reads an Any without
				// "@type" once looking for it and again to skip it, copying its
				// strings each time.
				start := w.dec.InputOffset()
				if err := w.skip(tok, depth); err != nil {
					return true, err
				}
				return true, w.alloc(jsonStringsCost(w.data[start:w.dec.InputOffset()]))
			}
			// Without a resolvable "@type" the decoder rejects a non-empty Any.
			if mt, err := protoregistry.GlobalTypes.FindMessageByURL(url); err == nil {
				return true, w.any(mt.Descriptor(), depth)
			}
		}
	case "google.protobuf.Timestamp", "google.protobuf.Duration":
		cost := 0
		if w.dynamic {
			// dynamicpb boxes the seconds and nanos it stores.
			cost = 2 * dynamicBoxCost(md.Fields().Get(0), 0)
		}
		if str, ok := tok.(string); ok && md.Name() == "Timestamp" && !wholeHourOffset(str) {
			cost += jsonFixedZoneCost
		}
		return false, w.alloc(cost)
	case "google.protobuf.BoolValue", "google.protobuf.Int32Value", "google.protobuf.Int64Value",
		"google.protobuf.UInt32Value", "google.protobuf.UInt64Value", "google.protobuf.FloatValue",
		"google.protobuf.DoubleValue", "google.protobuf.StringValue", "google.protobuf.BytesValue":
		value := md.Fields().Get(0)
		return false, w.alloc(jsonScalarCost(value, tok) + w.boxCost(value, tok))
	case "google.protobuf.FieldMask":
		if str, ok := tok.(string); ok {
			if err := w.alloc(fieldMaskCost(str)); err != nil {
				return true, err
			}
		}
	}
	return false, nil
}

// knownContainer walks a Struct or ListValue, whose JSON forms are its single
// map<string, Value> or repeated Value field.
func (w *jsonWalk) knownContainer(md protoreflect.MessageDescriptor, tok json.Token, depth int) error {
	if depth > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
	cost := jsonKnownListCost
	if md.FullName() == "google.protobuf.Struct" {
		cost = jsonKnownStructCost
	}
	if err := w.alloc(cost); err != nil {
		return err
	}
	return w.field(md.Fields().Get(0), tok, depth)
}

// knownValue walks a Value; every JSON value, null included, sets a member of
// its kind oneof.
func (w *jsonWalk) knownValue(md protoreflect.MessageDescriptor, tok json.Token, depth int) error {
	if depth > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
	member := md.Fields().ByName(knownValueMember(tok))
	cost := oneofWrapperCost(member) + jsonKnownValueCost
	if member.Message() == nil {
		cost += allocSize(goValueSize(member))
	}
	if err := w.alloc(cost); err != nil {
		return err
	}
	if member.Message() != nil {
		return w.message(member.Message(), tok, depth)
	}
	return nil
}

// any walks the members of an Any whose '{' has been consumed as a message of
// the embedded type md, which protojson then marshals into the Any's value.
func (w *jsonWalk) any(md protoreflect.MessageDescriptor, depth int) error {
	if w.dynamic {
		// The Any sets its type URL and value.
		if err := w.alloc(mapGrowth(0, knownSlotSize, true)); err != nil {
			return err
		}
	}
	start, entries := w.dec.InputOffset(), w.anyEntries
	var err error
	if hasCustomJSON(md) {
		err = w.anyValue(md, depth)
	} else if err = w.alloc(protoStructSize(md)); err == nil {
		err = w.object(md, depth)
	}
	if err != nil {
		return err
	}
	// Marshalling the Any's value sorts the map entries it holds directly;
	// a nested Any's entries were sorted into its own value already.
	sorted := w.anyEntries - entries
	w.anyEntries = entries
	return w.alloc(jsonAnyCopyFactor*allocSize(int(w.dec.InputOffset()-start)) + sorted*jsonAnyMapEntryCost)
}

// anyValue walks the members of an Any embedding a well-known type whose JSON
// form is its "value" member.
func (w *jsonWalk) anyValue(md protoreflect.MessageDescriptor, depth int) error {
	if depth > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
	for {
		key, err := w.next()
		if err != nil {
			return err
		}
		if key == json.Delim('}') {
			return nil
		}
		if err := w.count(); err != nil {
			return err
		}
		tok, err := w.next()
		if err != nil {
			return err
		}
		if key == "value" {
			err = w.message(md, tok, depth)
		} else {
			err = w.skip(tok, depth)
		}
		if err != nil {
			return err
		}
	}
}

// anyTypeURL looks ahead for the "@type" member of the object whose '{' has
// just been consumed, as protojson does, and returns "" if there is none. It
// scans the input in place: a nested Any repeats the scan over its subtree,
// so copying members would cost the estimator O(depth × size).
func (w *jsonWalk) anyTypeURL() string {
	s := jsonScanner{data: w.data, i: int(w.dec.InputOffset())}
	if s.consume('}') {
		return ""
	}
	for {
		key, ok := s.str()
		if !ok || !s.consume(':') {
			return ""
		}
		if jsonUnquote(key) == "@type" {
			url, _ := s.str()
			return jsonUnquote(url)
		}
		if !s.skip() || !s.consume(',') {
			return ""
		}
	}
}

// jsonScanner steps through JSON in place without validating it fully.
type jsonScanner struct {
	data []byte
	i    int
}

func (s *jsonScanner) space() {
	for s.i < len(s.data) && strings.IndexByte(" \t\n\r", s.data[s.i]) >= 0 {
		s.i++
	}
}

func (s *jsonScanner) consume(c byte) bool {
	s.space()
	if s.i < len(s.data) && s.data[s.i] == c {
		s.i++
		return true
	}
	return false
}

// str consumes a string and returns it with its quotes.
func (s *jsonScanner) str() ([]byte, bool) {
	s.space()
	if s.i >= len(s.data) || s.data[s.i] != '"' {
		return nil, false
	}
	for j := s.i + 1; j < len(s.data); j++ {
		switch s.data[j] {
		case '\\':
			j++
		case '"':
			quoted := s.data[s.i : j+1]
			s.i = j + 1
			return quoted, true
		}
	}
	return nil, false
}

// skip consumes one value.
func (s *jsonScanner) skip() bool {
	for depth := 0; ; {
		s.space()
		if s.i >= len(s.data) {
			return false
		}
		switch c := s.data[s.i]; c {
		case '"':
			if _, ok := s.str(); !ok {
				return false
			}
		case '{', '[':
			depth++
			s.i++
		case '}', ']', ',', ':':
			if depth == 0 {
				return false
			}
			if c == '}' || c == ']' {
				depth--
			}
			s.i++
		default:
			for s.i < len(s.data) && strings.IndexByte(" \t\n\r,:{}[]\"", s.data[s.i]) < 0 {
				s.i++
			}
		}
		if depth == 0 {
			return true
		}
	}
}

// jsonStringsCost is what tokenizing the strings in raw copies.
func jsonStringsCost(raw []byte) int {
	cost := 0
	for s := (jsonScanner{data: raw}); s.i < len(raw); {
		if raw[s.i] != '"' {
			s.i++
			continue
		}
		quoted, ok := s.str()
		if !ok {
			break
		}
		cost += jsonStringCost(quoted, len(jsonUnquote(quoted)))
	}
	return cost
}

// jsonUnquote decodes a quoted string from jsonScanner.str, or returns "" if
// it is malformed. Only strings with escapes allocate.
func jsonUnquote(quoted []byte) string {
	if len(quoted) < 2 {
		return ""
	}
	if bytes.IndexByte(quoted, '\\') < 0 {
		return string(quoted[1 : len(quoted)-1])
	}
	var s string
	if json.Unmarshal(quoted, &s) != nil {
		return ""
	}
	return s
}

func isKnownValue(field protoreflect.FieldDescriptor) bool {
	md := field.Message()
	return md != nil && md.FullName() == "google.protobuf.Value"
}

// knownValueMember is the member of google.protobuf.Value's kind oneof that
// protojson sets for tok.
func knownValueMember(tok json.Token) protoreflect.Name {
	switch tok.(type) {
	case bool:
		return "bool_value"
	case json.Number:
		return "number_value"
	case string:
		return "string_value"
	}
	switch tok {
	case json.Delim('{'):
		return "struct_value"
	case json.Delim('['):
		return "list_value"
	}
	return "null_value"
}

// hasCustomJSON reports whether protojson decodes md from a special JSON form
// rather than an object of its fields.
func hasCustomJSON(md protoreflect.MessageDescriptor) bool {
	if md.FullName().Parent() != "google.protobuf" {
		return false
	}
	switch md.Name() {
	case "Any", "Timestamp", "Duration", "Struct", "ListValue", "Value", "FieldMask", "Empty",
		"BoolValue", "Int32Value", "Int64Value", "UInt32Value", "UInt64Value",
		"FloatValue", "DoubleValue", "StringValue", "BytesValue":
		return true
	}
	return false
}

// wholeHourOffset reports whether an RFC 3339 timestamp's zone is one that
// time.FixedZone caches: Z, or a whole hour from -12 to +14.
func wholeHourOffset(ts string) bool {
	if strings.HasSuffix(ts, "Z") || strings.HasSuffix(ts, "z") || len(ts) < 6 {
		return true
	}
	zone := ts[len(ts)-6:]
	hours, err := strconv.Atoi(zone[1:3])
	if err != nil || zone[3] != ':' {
		return true
	}
	switch zone[0] {
	case '+':
		return zone[4:] == "00" && hours <= 14
	case '-':
		return zone[4:] == "00" && hours <= 12
	}
	return true
}

// fieldMaskCost is what protojson allocates splitting a FieldMask's JSON
// string into snake_case paths and appending them to its list.
func fieldMaskCost(str string) int {
	n := strings.Count(str, ",") + 1
	cost := appendAllocBytes(n, 16) + n*jsonFieldMaskPathCost
	for path := range strings.SplitSeq(str, ",") {
		// Up to two bytes per byte, appended to a nil slice and then copied.
		cost += appendAllocBytes(2*len(path), 1) + allocSize(2*len(path))
	}
	return cost
}
