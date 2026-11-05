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
			// Without a resolvable "@type" the decoder rejects a non-empty Any.
			if mt, err := protoregistry.GlobalTypes.FindMessageByURL(w.anyTypeURL()); err == nil {
				return true, w.any(mt.Descriptor(), depth)
			}
		}
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
// just been consumed, as protojson does, and returns "" if there is none.
func (w *jsonWalk) anyTypeURL() string {
	dec := json.NewDecoder(bytes.NewReader(w.data[w.dec.InputOffset()-1:]))
	if _, err := dec.Token(); err != nil {
		return ""
	}
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return ""
		}
		if key == "@type" {
			tok, _ := dec.Token()
			url, _ := tok.(string)
			return url
		}
		var member json.RawMessage
		if err := dec.Decode(&member); err != nil {
			return ""
		}
	}
	return ""
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
