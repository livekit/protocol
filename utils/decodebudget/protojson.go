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
	"errors"
	"io"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"

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

// ValidateProtoJSON checks protojson input against the budget in protobuf terms: an
// object key or array element is a field, and nesting is message depth.
func (b Proto) ValidateProtoJSON(data []byte, desc protoreflect.MessageDescriptor) error {
	_, err := b.EstimateProtoJSON(data, desc)
	return err
}

// EstimateProtoJSON validates protojson input like ValidateProtoJSON and returns the
// estimated bytes protojson allocates decoding it. Malformed or mistyped input
// is left for the decoder to reject.
func (b Proto) EstimateProtoJSON(data []byte, desc protoreflect.MessageDescriptor) (int, error) {
	t, err := b.newTally(desc)
	if err != nil {
		return 0, err
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	w := &jsonWalk{tally: t, dec: dec, data: data}

	tok, err := w.next()
	if err != nil {
		return w.total(), err
	}
	if err := w.message(desc, tok, 0); err != nil {
		return w.total(), err
	}
	if _, err := dec.Token(); err != io.EOF {
		return w.total(), errors.New("unexpected value after JSON document")
	}
	return w.total(), nil
}

// jsonMapReflectSize is the wrapper protojson allocates for every map field it
// sees, even an empty one.
const jsonMapReflectSize = 64

// jsonWalk walks protojson input with encoding/json tokens.
type jsonWalk struct {
	*tally
	dec   *json.Decoder
	data  []byte
	inAny bool
}

func (w *jsonWalk) next() (json.Token, error) {
	start := w.dec.InputOffset()
	tok, err := w.dec.Token()
	if err == io.EOF {
		return nil, io.ErrUnexpectedEOF
	}
	if err != nil {
		return nil, err
	}
	cost := w.profile.JSONTokenCost
	switch v := tok.(type) {
	case string:
		cost += jsonStringCost(w.data[start:w.dec.InputOffset()], len(v))
	case json.Number:
		// Token.Float parses from a copy of the literal.
		cost += tmpCopySize(len(v))
	}
	return tok, w.alloc(cost)
}

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

// object walks the members of a message whose '{' has been consumed.
func (w *jsonWalk) object(desc protoreflect.MessageDescriptor, depth int) error {
	if depth > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
	extended := false
	for {
		tok, err := w.next()
		if err != nil {
			return err
		}
		if tok == json.Delim('}') {
			return nil
		}
		if err := w.count(); err != nil {
			return err
		}
		key, _ := tok.(string)
		field := desc.Fields().ByJSONName(key)
		if field == nil {
			field = desc.Fields().ByTextName(key)
		}
		if field == nil {
			if field, err = jsonExtension(desc, key); err != nil {
				return err
			}
			if field != nil && !extended {
				extended = true
				if err := w.alloc(w.profile.ExtensionMapCost); err != nil {
					return err
				}
			}
		}
		tok, err = w.next()
		if err != nil {
			return err
		}
		if field == nil {
			// Unknown fields are discarded or rejected by the decoder.
			if err := w.skip(tok, depth); err != nil {
				return err
			}
			continue
		}
		if err := w.field(field, tok, depth); err != nil {
			return err
		}
	}
}

// jsonExtension resolves a "[full.name]" key the way protojson's default
// resolver does, returning nil when the decoder would discard or reject it.
func jsonExtension(desc protoreflect.MessageDescriptor, key string) (protoreflect.FieldDescriptor, error) {
	if len(key) < 2 || !strings.HasPrefix(key, "[") || !strings.HasSuffix(key, "]") {
		return nil, nil
	}
	extension, err := protoregistry.GlobalTypes.FindExtensionByName(protoreflect.FullName(key[1 : len(key)-1]))
	if errors.Is(err, protoregistry.NotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	field := extension.TypeDescriptor()
	if field.ContainingMessage().FullName() != desc.FullName() || !desc.ExtensionRanges().Has(field.Number()) {
		return nil, nil
	}
	return field, nil
}

func (w *jsonWalk) field(field protoreflect.FieldDescriptor, tok json.Token, depth int) error {
	switch {
	case tok == nil && !isKnownValue(field):
		return nil
	case field.IsMap():
		if tok != json.Delim('{') {
			return w.skip(tok, depth)
		}
		return w.mapEntries(field, depth)
	case field.IsList():
		if tok != json.Delim('[') {
			return w.skip(tok, depth)
		}
		for n := 0; ; n++ {
			tok, err := w.next()
			if err != nil {
				return err
			}
			if tok == json.Delim(']') {
				cost := appendAllocBytes(n, goValueSize(field))
				if field.Message() == nil {
					cost += n * w.profile.JSONListScalarCost
				}
				return w.alloc(cost)
			}
			if err := w.count(); err != nil {
				return err
			}
			if err := w.value(field, tok, depth); err != nil {
				return err
			}
		}
	default:
		return w.value(field, tok, depth)
	}
}

func (w *jsonWalk) mapEntries(field protoreflect.FieldDescriptor, depth int) error {
	if depth+1 > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
	if err := w.alloc(allocSize(mapHeaderSize) + jsonMapReflectSize); err != nil {
		return err
	}
	for n := 0; ; n++ {
		tok, err := w.next()
		if err != nil {
			return err
		}
		if tok == json.Delim('}') {
			return nil
		}
		if err := w.count(); err != nil {
			return err
		}
		// The duplicate-key check boxes the key once more than the set does.
		cost := mapInsertCost(field, n) + 2*mapBoxCost(field.MapKey()) + mapBoxCost(field.MapValue())
		if w.inAny {
			cost += jsonAnyMapEntryCost
		}
		if err := w.alloc(cost); err != nil {
			return err
		}
		tok, err = w.next()
		if err != nil {
			return err
		}
		if err := w.value(field.MapValue(), tok, depth+1); err != nil {
			return err
		}
	}
}

// value walks one occurrence of field. Slots, wrappers and map entries were
// charged by the caller only for maps; everything else is charged here.
func (w *jsonWalk) value(field protoreflect.FieldDescriptor, tok json.Token, depth int) error {
	if tok == nil && !isKnownValue(field) {
		return nil
	}
	if !field.IsMap() && field.ContainingMessage().IsMapEntry() {
		// Map values carry no slot or wrapper of their own.
		if md := field.Message(); md != nil {
			return w.message(md, tok, depth)
		}
		return w.alloc(jsonScalarCost(field, tok))
	}
	cost := w.profile.fieldAllocCost(field, 0) + jsonScalarCost(field, tok) + presenceCost(field)
	if field.Message() != nil {
		cost -= protoStructSize(field.Message())
	}
	if err := w.alloc(cost); err != nil {
		return err
	}
	if md := field.Message(); md != nil {
		return w.message(md, tok, depth)
	}
	if _, ok := tok.(json.Delim); ok {
		return w.skip(tok, depth)
	}
	return nil
}

func (w *jsonWalk) message(md protoreflect.MessageDescriptor, tok json.Token, depth int) error {
	if err := w.alloc(protoStructSize(md) + w.profile.JSONObjectCost); err != nil {
		return err
	}
	switch md.FullName() {
	case "google.protobuf.Struct", "google.protobuf.ListValue":
		return w.knownContainer(md, tok, depth+1)
	case "google.protobuf.Value":
		return w.knownValue(md, tok, depth+1)
	case "google.protobuf.Any":
		if tok == json.Delim('{') {
			if err := w.alloc(jsonAnyCost); err != nil {
				return err
			}
			// Without a resolvable "@type" the decoder rejects a non-empty Any.
			if mt, err := protoregistry.GlobalTypes.FindMessageByURL(w.anyTypeURL()); err == nil {
				return w.any(mt.Descriptor(), depth+1)
			}
		}
	case "google.protobuf.FieldMask":
		if str, ok := tok.(string); ok {
			if err := w.alloc(fieldMaskCost(str)); err != nil {
				return err
			}
		}
	}
	if tok != json.Delim('{') || hasCustomJSON(md) {
		return w.skip(tok, depth+1)
	}
	return w.object(md, depth+1)
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
	start := w.dec.InputOffset()
	inAny := w.inAny
	w.inAny = true
	var err error
	if hasCustomJSON(md) {
		err = w.anyValue(md, depth)
	} else if err = w.alloc(protoStructSize(md)); err == nil {
		err = w.object(md, depth)
	}
	w.inAny = inAny
	if err != nil {
		return err
	}
	return w.alloc(jsonAnyCopyFactor * allocSize(int(w.dec.InputOffset()-start)))
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

// skip walks a value without a descriptor, counting every token as a field.
func (w *jsonWalk) skip(tok json.Token, depth int) error {
	d, ok := tok.(json.Delim)
	if !ok || (d != '{' && d != '[') {
		return w.count()
	}
	for open := 1; open > 0; {
		if depth+open > w.budget.MaxDepth {
			return ErrMessageComplexity
		}
		tok, err := w.next()
		if err != nil {
			return err
		}
		if err := w.count(); err != nil {
			return err
		}
		switch tok {
		case json.Delim('{'), json.Delim('['):
			open++
		case json.Delim('}'), json.Delim(']'):
			open--
		}
	}
	return nil
}

// jsonScalarCost is what decoding a scalar token allocates beyond the copy the
// tokenizer already made: base64 bytes are decoded into a new buffer, and a
// number in a string is tokenized again from a copy and parsed from another.
func jsonScalarCost(field protoreflect.FieldDescriptor, tok json.Token) int {
	str, ok := tok.(string)
	if !ok {
		return 0
	}
	switch field.Kind() {
	case protoreflect.BytesKind:
		return allocSize(len(str) * 3 / 4)
	case protoreflect.StringKind, protoreflect.BoolKind, protoreflect.EnumKind,
		protoreflect.MessageKind, protoreflect.GroupKind:
		return 0
	}
	return 2 * tmpCopySize(len(str))
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
