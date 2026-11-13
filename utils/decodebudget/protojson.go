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
	"math"
	"strconv"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
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
	return b.estimateProtoJSON(data, desc, false)
}

// estimateProtoJSON is EstimateProtoJSON decoding into a dynamicpb message when
// dynamic is set, as it must when desc has no generated type.
func (b Proto) estimateProtoJSON(data []byte, desc protoreflect.MessageDescriptor, dynamic bool) (int, error) {
	t, err := b.newTally(desc)
	if err != nil {
		return 0, err
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	w := &jsonWalk{tally: t, dec: dec, data: data, dynamic: dynamic || !layoutOf(desc).generated}
	if w.dynamic {
		if err := w.alloc(dynamicResetCost); err != nil {
			return w.total(), err
		}
	}

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

// protojson allocates per token and per object on top of the decoded message,
// its reflection-based list setter boxes each scalar element, it wraps every
// map field it sees, even an empty one, and it sets each extension value.
const (
	jsonExtensionCost  = 64
	jsonTokenCost      = 8
	jsonObjectCost     = 8
	jsonListScalarCost = 24
	jsonMapReflectSize = 64
	// A dynamicpb message without the closure binary decoding allocates.
	jsonDynamicMessageCost = 64 + 2*48
)

// jsonWalk walks protojson input with encoding/json tokens.
type jsonWalk struct {
	*tally
	dec  *json.Decoder
	data []byte
	// anyEntries counts map entries walked since the innermost Any began.
	anyEntries int
	// dynamic prices the message being walked as a dynamicpb message.
	dynamic bool
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
	cost := jsonTokenCost
	switch v := tok.(type) {
	case string:
		cost += jsonStringCost(w.data[start:w.dec.InputOffset()], len(v))
	case json.Number:
		// Token.Float parses from a copy of the literal.
		cost += tmpCopySize(len(v))
	}
	return tok, w.alloc(cost)
}

// object walks the members of a message whose '{' has been consumed.
func (w *jsonWalk) object(desc protoreflect.MessageDescriptor, depth int) error {
	if depth > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
	var set setFields
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
		field, err := jsonField(desc, key)
		if err != nil {
			return err
		}
		if tok, err = w.next(); err != nil {
			return err
		}
		if field == nil {
			// Unknown fields are discarded or rejected by the decoder.
			err = w.skip(tok, depth)
		} else if err = w.alloc(w.setCost(&set, field)); err == nil {
			err = w.field(field, tok, depth)
		}
		if err != nil {
			return err
		}
	}
}

// setFields counts the fields and extensions an object has set; its member
// names are unique, so each sets a new one.
type setFields struct {
	fields, extensions int
}

// setCost is the map growth of setting field in an object.
func (w *jsonWalk) setCost(set *setFields, field protoreflect.FieldDescriptor) int {
	cost := 0
	if field.IsExtension() {
		cost = w.extensionMapCost(set.extensions)
		set.extensions++
	}
	if w.dynamic {
		cost += mapGrowth(set.fields, knownSlotSize, true)
		set.fields++
	}
	return cost
}

// jsonField resolves an object key the way protojson does.
func jsonField(desc protoreflect.MessageDescriptor, key string) (protoreflect.FieldDescriptor, error) {
	if field := desc.Fields().ByJSONName(key); field != nil {
		return field, nil
	}
	if field := desc.Fields().ByTextName(key); field != nil {
		return field, nil
	}
	return jsonExtension(desc, key)
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
		if w.dynamic {
			if err := w.alloc(dynamicListCost); err != nil {
				return err
			}
		}
		for n := 0; ; n++ {
			tok, err := w.next()
			if err != nil {
				return err
			}
			if tok == json.Delim(']') {
				return w.alloc(w.listCost(field, n))
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
	cost := allocSize(mapHeaderSize) + jsonMapReflectSize
	if w.dynamic {
		cost = dynamicMapCost
	}
	if err := w.alloc(cost); err != nil {
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
		w.anyEntries++
		if err := w.alloc(w.mapEntryCost(field, n, tok)); err != nil {
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
		return w.alloc(jsonScalarCost(field, tok) + w.boxCost(field, tok))
	}
	cost := jsonScalarCost(field, tok) + w.boxCost(field, tok)
	if !w.dynamic {
		cost += fieldExtraCost(field, 0) + presenceCost(field)
	}
	if field.IsExtension() {
		cost += jsonExtensionCost
		if !isGeneratedExtension(field) {
			cost += dynamicExtensionResolveCost + dynamicExtensionCheckCost
		}
	}
	if err := w.alloc(cost); err != nil {
		return err
	}
	if md := field.Message(); md != nil {
		if !field.IsExtension() {
			return w.message(md, tok, depth)
		}
		// An extension's message follows its type.
		sub := *w
		sub.dynamic = !isGeneratedExtension(field)
		err := sub.message(md, tok, depth)
		w.anyEntries = sub.anyEntries
		return err
	}
	if _, ok := tok.(json.Delim); ok {
		return w.skip(tok, depth)
	}
	return nil
}

func (w *jsonWalk) message(md protoreflect.MessageDescriptor, tok json.Token, depth int) error {
	size := protoStructSize(md)
	if w.dynamic {
		size = jsonDynamicMessageCost
		if hasCustomJSON(md) && md.Fields().Len() > 0 && md.FullName() != "google.protobuf.Any" {
			// Its JSON form sets fields without walking an object.
			size += mapGrowth(0, knownSlotSize, true)
		}
	}
	if err := w.alloc(size + jsonObjectCost); err != nil {
		return err
	}
	if done, err := w.wellKnown(md, tok, depth+1); done || err != nil {
		return err
	}
	if tok != json.Delim('{') || hasCustomJSON(md) {
		return w.skip(tok, depth+1)
	}
	return w.object(md, depth+1)
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

// boxCost is a dynamicpb message boxing a value it stores. Go boxes a value
// whose bits are under 256 without allocating.
func (w *jsonWalk) boxCost(field protoreflect.FieldDescriptor, tok json.Token) int {
	if !w.dynamic {
		return 0
	}
	if n, ok := tok.(json.Number); ok {
		if kind := field.Kind(); kind == protoreflect.FloatKind || kind == protoreflect.DoubleKind {
			// A float is boxed by its bits, so only zero is free.
			if f, err := n.Float64(); err == nil && math.Float64bits(f) == 0 {
				return 0
			}
		} else if u, err := strconv.ParseUint(string(n), 10, 64); err == nil && u < 256 {
			return 0
		}
	}
	str, _ := tok.(string)
	return dynamicBoxCost(field, len(str))
}

// listCost is the growth of a list of n elements.
func (w *jsonWalk) listCost(field protoreflect.FieldDescriptor, n int) int {
	if w.dynamic {
		mem, _ := appendValues(0, 0, n)
		return mem
	}
	cost := appendAllocBytes(n, goValueSize(field))
	if field.Message() == nil {
		cost += n * jsonListScalarCost
	}
	return cost
}

// mapEntryCost is inserting a new key into a map field holding n entries. The
// duplicate-key check boxes the key once more than the set does.
func (w *jsonWalk) mapEntryCost(field protoreflect.FieldDescriptor, n int, key json.Token) int {
	if !w.dynamic {
		return mapInsertCost(field, n) + 2*mapBoxCost(field.MapKey()) + mapBoxCost(field.MapValue())
	}
	str, _ := key.(string)
	return mapGrowth(n, mapEntrySlotSize, true) + 3*dynamicBoxCost(field.MapKey(), len(str))
}

// extensionMapCost is the extension map growth for a message's next
// extension, having set n.
func (w *jsonWalk) extensionMapCost(n int) int {
	switch {
	case w.dynamic:
		return mapGrowth(n, extSlotSize, true)
	case n == 0:
		return extensionMapSize
	}
	return mapGrowth(n, extensionSlotSize, true)
}
