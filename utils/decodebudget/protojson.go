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

	"google.golang.org/protobuf/reflect/protoreflect"
)

// protojson allocates per token and per object on top of the decoded message
// and the string copies charged separately, and its reflection-based list and
// map setters box each scalar element and map entry.
const (
	jsonTokenCost       = 8
	jsonListScalarCost  = 24
	jsonMapEntryBoxCost = 32
	jsonObjectCost      = 8
)

// ValidateJSON checks protojson input against the budget in protobuf terms: an
// object key or array element is a field, and nesting is message depth.
func (b Proto) ValidateJSON(data []byte, desc protoreflect.MessageDescriptor) error {
	_, err := b.EstimateJSON(data, desc)
	return err
}

// EstimateJSON validates protojson input like ValidateJSON and returns the
// estimated bytes protojson allocates decoding it. Malformed or mistyped input
// is left for the decoder to reject.
func (b Proto) EstimateJSON(data []byte, desc protoreflect.MessageDescriptor) (int, error) {
	t, err := b.newTally(desc)
	if err != nil {
		return 0, err
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	w := &jsonWalk{tally: t, dec: dec}

	tok, err := w.next()
	if err != nil {
		return w.total(), err
	}
	if err := w.alloc(protoStructSize(desc)); err != nil {
		return w.total(), err
	}
	if tok != json.Delim('{') {
		err := w.skip(tok, 1)
		return w.total(), err
	}
	if err := w.object(desc, 1); err != nil {
		return w.total(), err
	}
	if _, err := dec.Token(); err != io.EOF {
		return w.total(), errors.New("unexpected value after JSON document")
	}
	return w.total(), nil
}

// jsonWalk walks protojson input with encoding/json tokens.
type jsonWalk struct {
	*tally
	dec *json.Decoder
}

func (w *jsonWalk) next() (json.Token, error) {
	tok, err := w.dec.Token()
	if err == io.EOF {
		return nil, io.ErrUnexpectedEOF
	}
	if err != nil {
		return nil, err
	}
	cost := jsonTokenCost
	if str, ok := tok.(string); ok {
		// protojson's tokenizer copies every string, key or value.
		cost += allocSize(len(str))
	}
	return tok, w.alloc(cost)
}

// object walks the members of a message whose '{' has been consumed.
func (w *jsonWalk) object(desc protoreflect.MessageDescriptor, depth int) error {
	if depth > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
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
		tok, err = w.next()
		if err != nil {
			return err
		}
		if field == nil {
			// Unknown fields and extensions are discarded or rejected by the decoder.
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

func (w *jsonWalk) field(field protoreflect.FieldDescriptor, tok json.Token, depth int) error {
	switch {
	case tok == nil:
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
					cost += n * jsonListScalarCost
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
		if err := w.alloc(mapEntryCost(field) + jsonMapEntryBoxCost); err != nil {
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
	if tok == nil {
		return nil
	}
	if !field.IsMap() && field.ContainingMessage().IsMapEntry() {
		// Map values carry no slot or wrapper of their own.
		if md := field.Message(); md != nil {
			return w.message(md, tok, depth)
		}
		return w.alloc(jsonScalarCost(field, tok))
	}
	cost := fieldAllocCost(field, 0) + jsonScalarCost(field, tok) + presenceCost(field)
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
	if err := w.alloc(protoStructSize(md) + jsonObjectCost); err != nil {
		return err
	}
	if tok != json.Delim('{') || md.FullName().Parent() == "google.protobuf" {
		// Well-known types have their own JSON forms; walk them generically.
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

// jsonScalarCost is what decoding a scalar token allocates beyond the copy the
// tokenizer already made: base64 bytes are decoded into a new buffer.
func jsonScalarCost(field protoreflect.FieldDescriptor, tok json.Token) int {
	if str, ok := tok.(string); ok && field.Kind() == protoreflect.BytesKind {
		return allocSize(len(str) * 3 / 4)
	}
	return 0
}
