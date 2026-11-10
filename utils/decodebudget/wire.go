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
	// extensions holds the infos of extension types seen in this walk, so
	// their states compare by pointer like other fields'.
	extensions []*fieldInfo
}

// decoderModel prices what one protobuf-go decoder allocates for each event of
// the wire walk, reading and updating the mirrored decoder state.
type decoderModel interface {
	root(info *messageInfo) int
	// message is the allocation for one occurrence of a message field, and the
	// state of a singular field, whose child holds what its contents decode
	// into once the walker creates it. Elements of lists and maps have none.
	message(m *messageState, f *fieldInfo) (int, *fieldState)
	// occurrence is the allocation for one occurrence of field with a string or
	// bytes payload of payloadLen.
	occurrence(m *messageState, f *fieldInfo, payloadLen int) int
	packed(m *messageState, f *fieldInfo, count int) int
	// mismatch is what the decoder allocates for a known field before it
	// rejects an unexpected wire type.
	mismatch(m *messageState, f *fieldInfo) int
	// unknown is what the decoder allocates for a field number the message
	// does not know before it keeps the field as unknown.
	unknown(m *messageState, info *messageInfo) int
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
	field *fieldInfo
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
	field *fieldInfo
}

func (m *messageState) find(f *fieldInfo) *fieldState {
	for i := range m.fields {
		if m.fields[i].field == f {
			return &m.fields[i]
		}
	}
	return nil
}

// add returns field's state, reporting whether it was just created.
func (m *messageState) add(f *fieldInfo) (*fieldState, bool) {
	if fs := m.find(f); fs != nil {
		return fs, false
	}
	m.fields = append(m.fields, fieldState{field: f})
	if f.extension {
		m.extensions++
	}
	return &m.fields[len(m.fields)-1], true
}

func (m *messageState) get(f *fieldInfo) *fieldState {
	fs, _ := m.add(f)
	return fs
}

// setOneof records field as its oneof's set member. Switching members discards
// the previous member's value, so a later return to field starts over.
func (m *messageState) setOneof(f *fieldInfo, fs *fieldState) {
	oneof := f.oneof
	if oneof == nil {
		return
	}
	for i := range m.oneofs {
		if m.oneofs[i].oneof == oneof {
			if m.oneofs[i].field != f {
				*fs = fieldState{field: f}
				m.oneofs[i].field = f
			}
			return
		}
	}
	m.oneofs = append(m.oneofs, oneofState{oneof: oneof, field: f})
}

func (m *messageState) isSet(f *fieldInfo) bool {
	for _, o := range m.oneofs {
		if o.oneof == f.oneof && o.field == f {
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
func (w *wireWalk) message(data []byte, info *messageInfo, depth int, state *messageState, model decoderModel) error {
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
		f := info.field(number)
		if f == nil && info.extendable && info.desc.ExtensionRanges().Has(number) {
			extension, err := protoregistry.GlobalTypes.FindExtensionByNumber(info.desc.FullName(), number)
			if err == nil {
				f = w.extension(extension.TypeDescriptor())
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
		if f == nil || !f.matches(kind) {
			cost := state.unknown(n + size)
			if f != nil {
				cost += model.mismatch(state, f)
			} else {
				cost += model.unknown(state, info)
			}
			if err := w.alloc(cost); err != nil {
				return err
			}
			continue
		}
		var err error
		switch {
		case kind == protowire.BytesType && f.child != nil:
			cost, fs := model.message(state, f)
			if err = w.alloc(cost); err == nil {
				err = w.child(value, f.child, depth+1, fs, valueModel(model, f))
			}
		case kind == protowire.BytesType && f.packable:
			err = w.packed(value, f, state, model)
		default:
			err = w.alloc(model.occurrence(state, f, len(value)))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *wireWalk) extension(xd protoreflect.ExtensionTypeDescriptor) *fieldInfo {
	for _, f := range w.extensions {
		if f.FieldDescriptor == protoreflect.FieldDescriptor(xd) {
			return f
		}
	}
	var child *messageInfo
	if md := xd.Message(); md != nil {
		child = infoOf(md)
	}
	f := newFieldInfo(xd, child)
	w.extensions = append(w.extensions, &f)
	return &f
}

func (w *wireWalk) child(data []byte, info *messageInfo, depth int, fs *fieldState, model decoderModel) error {
	if fs == nil {
		state := w.states.get()
		err := w.message(data, info, depth, state, model)
		w.states.put(state)
		return err
	}
	if fs.child == nil {
		fs.child = w.states.get()
	}
	return w.message(data, info, depth, fs.child, model)
}

// packed walks a packed scalar run; decoders accept one even when the schema
// prefers unpacked.
func (w *wireWalk) packed(value []byte, f *fieldInfo, state *messageState, model decoderModel) error {
	count := 0
	for len(value) > 0 {
		var consumed int
		switch f.wire {
		case protowire.Fixed64Type:
			_, consumed = protowire.ConsumeFixed64(value)
		case protowire.Fixed32Type:
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
	return w.alloc(model.packed(state, f, count))
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
