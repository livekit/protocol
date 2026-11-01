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
	"errors"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// wireWalk walks the protobuf binary encoding.
type wireWalk struct {
	*tally
}

// messageState mirrors the decoder state of one message whose cost depends on
// earlier occurrences: each repeated field's slice, whether a presence field's
// pointer exists, which member of each oneof is set, and the unknown-fields
// slice.
type messageState struct {
	fields     []fieldState
	oneofs     []oneofState
	unknownLen int
	unknownCap int
}

type fieldState struct {
	field protoreflect.FieldDescriptor
	len   int
	cap   int
	// child is the state of a singular message field. The decoder merges each
	// later occurrence into the same message, so its slices keep growing.
	child *messageState
}

type oneofState struct {
	oneof protoreflect.OneofDescriptor
	field protoreflect.FieldDescriptor
}

func (m *messageState) get(field protoreflect.FieldDescriptor) *fieldState {
	for i := range m.fields {
		if m.fields[i].field == field {
			return &m.fields[i]
		}
	}
	m.fields = append(m.fields, fieldState{field: field})
	return &m.fields[len(m.fields)-1]
}

// message walks the fields of one message into state, which persists across
// occurrences that the decoder merges into the same message.
func (w *wireWalk) message(data []byte, desc protoreflect.MessageDescriptor, depth int, state *messageState) error {
	if depth > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
	for len(data) > 0 {
		if err := w.count(); err != nil {
			return err
		}
		number, kind, n := protowire.ConsumeTag(data)
		if n < 0 {
			return protowire.ParseError(n)
		}
		data = data[n:]
		if kind == protowire.StartGroupType || kind == protowire.EndGroupType {
			return ErrMessageComplexity
		}
		field := desc.Fields().ByNumber(number)
		if field == nil && desc.ExtensionRanges().Has(number) {
			extension, err := protoregistry.GlobalTypes.FindExtensionByNumber(desc.FullName(), number)
			if err == nil {
				field = extension.TypeDescriptor()
			} else if !errors.Is(err, protoregistry.NotFound) {
				return err
			}
		}

		var value []byte
		var size int
		if kind == protowire.BytesType {
			value, size = protowire.ConsumeBytes(data)
		} else {
			size = protowire.ConsumeFieldValue(number, kind, data)
		}
		if size < 0 {
			return protowire.ParseError(size)
		}
		data = data[size:]

		// The decoder keeps a field with an unexpected wire type as unknown.
		if field == nil || !wireTypeMatches(field, kind) {
			if err := w.alloc(state.unknown(n + size)); err != nil {
				return err
			}
			continue
		}
		var err error
		switch {
		case kind == protowire.BytesType && field.Message() != nil:
			cost, child := state.messageOccurrence(field)
			if err = w.alloc(cost); err == nil {
				err = w.message(value, field.Message(), depth+1, child)
			}
		case kind == protowire.BytesType && isPackable(field):
			err = w.packed(value, field, state)
		default:
			err = w.alloc(state.occurrence(field, len(value)))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// messageOccurrence is the allocation for one occurrence of a message field,
// and the state its contents decode into.
func (m *messageState) messageOccurrence(field protoreflect.FieldDescriptor) (int, *messageState) {
	if field.IsList() || field.IsMap() {
		// Every element or map value is a new message.
		return m.occurrence(field, 0), &messageState{}
	}
	fs := m.get(field)
	m.setOneof(field, fs)
	if fs.child != nil {
		return 0, fs.child
	}
	fs.child = &messageState{}
	return fieldAllocCost(field, 0), fs.child
}

// setOneof records field as its oneof's set member. Switching members discards
// the previous member's value, so a later return to field starts over.
func (m *messageState) setOneof(field protoreflect.FieldDescriptor, fs *fieldState) {
	oneof := field.ContainingOneof()
	if oneof == nil || oneof.IsSynthetic() {
		return
	}
	for i := range m.oneofs {
		if m.oneofs[i].oneof == oneof {
			if m.oneofs[i].field != field {
				*fs = fieldState{field: field}
				m.oneofs[i].field = field
			}
			return
		}
	}
	m.oneofs = append(m.oneofs, oneofState{oneof: oneof, field: field})
}

// occurrence is the allocation for one occurrence of field with a string or
// bytes payload of payloadLen. Nested message contents are charged separately.
func (m *messageState) occurrence(field protoreflect.FieldDescriptor, payloadLen int) int {
	cost := fieldAllocCost(field, payloadLen)
	switch {
	case field.IsList():
		fs := m.get(field)
		var mem int
		mem, fs.cap = appendGrow(fs.len, fs.cap, 1, goValueSize(field))
		fs.len++
		cost += mem
	case presenceCost(field) > 0:
		if fs := m.get(field); fs.len == 0 {
			fs.len = 1
			cost += presenceCost(field)
		}
	case field.ContainingOneof() != nil:
		m.setOneof(field, m.get(field))
	}
	return cost
}

// unknown is the allocation for retaining an n-byte field in the message's
// unknown-fields slice.
func (m *messageState) unknown(n int) int {
	mem, newCap := appendGrow(m.unknownLen, m.unknownCap, n, 1)
	m.unknownLen += n
	m.unknownCap = newCap
	return mem
}

// packed walks a packed scalar run; decoders accept one even when the schema
// prefers unpacked. Each run reallocates the field's slice at exactly its new
// length, so many small runs cost quadratically.
func (w *wireWalk) packed(value []byte, field protoreflect.FieldDescriptor, state *messageState) error {
	count := 0
	for len(value) > 0 {
		var consumed int
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
			_, consumed = protowire.ConsumeFixed64(value)
		case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
			_, consumed = protowire.ConsumeFixed32(value)
		default:
			_, consumed = protowire.ConsumeVarint(value)
		}
		if consumed < 0 {
			return protowire.ParseError(consumed)
		}
		if err := w.count(); err != nil {
			return err
		}
		count++
		value = value[consumed:]
	}
	if count == 0 {
		return nil
	}
	fs := state.get(field)
	fs.len += count
	fs.cap = fs.len
	return w.alloc(allocSize(fs.len * goValueSize(field)))
}

func wireTypeMatches(field protoreflect.FieldDescriptor, kind protowire.Type) bool {
	if kind == protowire.BytesType && isPackable(field) {
		return true
	}
	switch field.Kind() {
	case protoreflect.MessageKind, protoreflect.StringKind, protoreflect.BytesKind:
		return kind == protowire.BytesType
	case protoreflect.GroupKind:
		return kind == protowire.StartGroupType
	case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		return kind == protowire.Fixed32Type
	case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		return kind == protowire.Fixed64Type
	default:
		return kind == protowire.VarintType
	}
}

func isPackable(field protoreflect.FieldDescriptor) bool {
	if field == nil || !field.IsList() {
		return false
	}
	switch field.Kind() {
	case protoreflect.BoolKind, protoreflect.EnumKind,
		protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
		protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind,
		protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		return true
	default:
		return false
	}
}
