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

// wireWalk walks the protobuf binary encoding, pricing each occurrence with
// the model of the decoder that will write its message.
type wireWalk struct {
	*tally
	states statePool
}

// decoderModel prices what one protobuf-go decoder allocates for each event of
// the wire walk, reading and updating the mirrored decoder state.
type decoderModel interface {
	root(desc protoreflect.MessageDescriptor) int
	// message is the allocation for one occurrence of a message field, and the
	// state of a singular field, whose child holds what its contents decode
	// into once the walker creates it. Elements of lists and maps have none.
	message(m *messageState, field protoreflect.FieldDescriptor) (int, *fieldState)
	// occurrence is the allocation for one occurrence of field with a string or
	// bytes payload of payloadLen.
	occurrence(m *messageState, field protoreflect.FieldDescriptor, payloadLen int) int
	packed(m *messageState, field protoreflect.FieldDescriptor, count int) int
	// mismatch is what the decoder allocates for a known field before it
	// rejects an unexpected wire type.
	mismatch(m *messageState, field protoreflect.FieldDescriptor) int
	// unknown is what the decoder allocates for a field number desc does not
	// know before it keeps the field as unknown.
	unknown(m *messageState, desc protoreflect.MessageDescriptor) int
}

func newDecoderModel(dynamic bool) decoderModel {
	if dynamic {
		return dynamicModel{}
	}
	return generatedModel{}
}

// messageState mirrors the decoder state of one message whose cost depends on
// earlier occurrences.
type messageState struct {
	fields     []fieldState
	oneofs     []oneofState
	unknownLen int
	unknownCap int
	// extensions counts the extension fields in fields, and extensionMap
	// records that a generated message's extension map exists.
	extensions   int
	extensionMap bool
}

type fieldState struct {
	field protoreflect.FieldDescriptor
	// len and cap mirror a repeated field's slice.
	len, cap int
	// entries counts a map field's insertions, and allocated records that the
	// map itself exists.
	entries   int
	allocated bool
	// present records that a presence field's pointer exists.
	present bool
	// child is the state of a singular message field. The decoder merges each
	// later occurrence into the same message, so its slices keep growing.
	child *messageState
}

type oneofState struct {
	oneof protoreflect.OneofDescriptor
	field protoreflect.FieldDescriptor
}

func (m *messageState) find(field protoreflect.FieldDescriptor) *fieldState {
	for i := range m.fields {
		if m.fields[i].field == field {
			return &m.fields[i]
		}
	}
	return nil
}

// add returns field's state, reporting whether it was just created.
func (m *messageState) add(field protoreflect.FieldDescriptor) (*fieldState, bool) {
	if fs := m.find(field); fs != nil {
		return fs, false
	}
	m.fields = append(m.fields, fieldState{field: field})
	if field.IsExtension() {
		m.extensions++
	}
	return &m.fields[len(m.fields)-1], true
}

func (m *messageState) get(field protoreflect.FieldDescriptor) *fieldState {
	fs, _ := m.add(field)
	return fs
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

func (m *messageState) isSet(oneof protoreflect.OneofDescriptor, field protoreflect.FieldDescriptor) bool {
	for _, o := range m.oneofs {
		if o.oneof == oneof && o.field == field {
			return true
		}
	}
	return false
}

// statePool recycles the states of list and map elements, which nothing
// reads once the element is walked, to keep the walk's own allocation down.
type statePool struct {
	free []*messageState
}

func (p *statePool) get() *messageState {
	n := len(p.free)
	if n == 0 {
		return &messageState{}
	}
	m := p.free[n-1]
	p.free = p.free[:n-1]
	return m
}

func (p *statePool) put(m *messageState) {
	for i := range m.fields {
		if child := m.fields[i].child; child != nil {
			p.put(child)
		}
	}
	clear(m.fields)
	clear(m.oneofs)
	*m = messageState{fields: m.fields[:0], oneofs: m.oneofs[:0]}
	p.free = append(p.free, m)
}

// unknown is the allocation for retaining an n-byte field in the message's
// unknown-fields slice.
func (m *messageState) unknown(n int) int {
	mem, newCap := appendGrow(m.unknownLen, m.unknownCap, n, 1)
	m.unknownLen += n
	m.unknownCap = newCap
	return mem
}

// message walks the fields of one message into state, which persists across
// occurrences that the decoder merges into the same message.
func (w *wireWalk) message(data []byte, desc protoreflect.MessageDescriptor, depth int, state *messageState, model decoderModel) error {
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
			cost := state.unknown(n + size)
			if field != nil {
				cost += model.mismatch(state, field)
			} else {
				cost += model.unknown(state, desc)
			}
			if err := w.alloc(cost); err != nil {
				return err
			}
			continue
		}
		var err error
		switch {
		case kind == protowire.BytesType && field.Message() != nil:
			cost, fs := model.message(state, field)
			if err = w.alloc(cost); err == nil {
				err = w.child(value, field.Message(), depth+1, fs, valueModel(model, field))
			}
		case kind == protowire.BytesType && isPackable(field):
			err = w.packed(value, field, state, model)
		default:
			err = w.alloc(model.occurrence(state, field, len(value)))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *wireWalk) child(data []byte, desc protoreflect.MessageDescriptor, depth int, fs *fieldState, model decoderModel) error {
	if fs == nil {
		state := w.states.get()
		err := w.message(data, desc, depth, state, model)
		w.states.put(state)
		return err
	}
	if fs.child == nil {
		fs.child = w.states.get()
	}
	return w.message(data, desc, depth, fs.child, model)
}

// packed walks a packed scalar run; decoders accept one even when the schema
// prefers unpacked.
func (w *wireWalk) packed(value []byte, field protoreflect.FieldDescriptor, state *messageState, model decoderModel) error {
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
	return w.alloc(model.packed(state, field, count))
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
