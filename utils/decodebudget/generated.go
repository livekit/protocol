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
	"google.golang.org/protobuf/reflect/protoreflect"
)

// mapInsertCost is the allocation for inserting a new key into a generated
// map field holding n entries, excluding the map header.
func mapInsertCost(field protoreflect.FieldDescriptor, n int) int {
	slot, pointers := mapSlot(field)
	return mapGrowth(n, slot, pointers)
}

func mapSlot(field protoreflect.FieldDescriptor) (size int, pointers bool) {
	key, value := field.MapKey(), field.MapValue()
	return (goValueSize(key) + goValueSize(value) + 7) &^ 7, hasPointers(key) || hasPointers(value)
}

// mapBoxCost is the allocation for passing one map key or value of field's
// kind through reflect: a scalar is boxed by Interface and again by Convert.
func mapBoxCost(field protoreflect.FieldDescriptor) int {
	if field.Message() != nil {
		return 0
	}
	return 2 * allocSize(goValueSize(field))
}

// reflectAppendCost is appending one element to a generated list through
// protoreflect: reflect.Append moves the slice header to the heap, and boxes a
// scalar, which the tiny allocator leaves unrounded when it has no pointers.
func reflectAppendCost(field protoreflect.FieldDescriptor) int {
	const sliceHeader = 24
	if field.Message() != nil {
		return sliceHeader
	}
	return sliceHeader + goValueSize(field)
}

func oneofWrapperCost(field protoreflect.FieldDescriptor) int {
	return allocSize(goValueSize(field))
}

// presenceCost is the pointer a generated struct allocates the first time a
// scalar field with explicit presence, such as proto3 optional, is set.
func presenceCost(field protoreflect.FieldDescriptor) int {
	if !field.HasPresence() || field.IsList() || field.Message() != nil || field.IsExtension() ||
		(field.ContainingOneof() != nil && !field.ContainingOneof().IsSynthetic()) ||
		field.ContainingMessage().IsMapEntry() {
		return 0
	}
	return allocSize(goValueSize(field))
}

// fieldExtraCost is the allocation for one occurrence of field with a string or
// bytes payload of payloadLen, without the field's own message, which the
// walks charge themselves.
func fieldExtraCost(field protoreflect.FieldDescriptor, payloadLen int) int {
	return fieldCost(field, payloadLen, 0)
}

// fieldCost is fieldExtraCost plus msgCost for the field's own message, or
// for a map field, the value message the decoder allocates before reading
// each entry, whether or not the entry has one.
func fieldCost(field protoreflect.FieldDescriptor, payloadLen, msgCost int) int {
	cost := 0
	// A map value message is charged with its entry.
	if !field.ContainingMessage().IsMapEntry() {
		cost += msgCost
	}
	if k := field.Kind(); k == protoreflect.StringKind || k == protoreflect.BytesKind {
		cost += allocSize(payloadLen)
	}
	switch {
	case field.IsMap():
		// Map growth depends on earlier entries; the caller charges it.
		cost += mapBoxCost(field.MapKey()) + mapBoxCost(field.MapValue())
	case field.IsList():
		// Slice growth depends on earlier occurrences; the caller charges it.
	case field.ContainingOneof() != nil && !field.ContainingOneof().IsSynthetic():
		cost += oneofWrapperCost(field)
	}
	return cost
}

// generatedModel prices protobuf-go's fast-path decoder writing into
// generated Go structs.
type generatedModel struct{}

func (g generatedModel) root(info *messageInfo) int {
	return info.layout.size
}

func (g generatedModel) message(m *messageState, f *fieldInfo) (int, *fieldState) {
	if f.list || f.isMap {
		// Every element or map value is a new message.
		return g.occurrence(m, f, 0), nil
	}
	fs, isNew := m.add(f)
	cost := 0
	if f.extension {
		cost = extensionCost(m, f, isNew)
	}
	m.setOneof(f, fs)
	if fs.child != nil {
		return cost, fs
	}
	return cost + f.allocCost, fs
}

func (g generatedModel) occurrence(m *messageState, f *fieldInfo, payloadLen int) int {
	cost := f.allocCost
	if f.payload {
		cost += allocSize(payloadLen)
	}
	if f.extension {
		fs, isNew := m.add(f)
		cost += extensionCost(m, f, isNew)
		if f.list {
			cost += extensionListCost(fs, f, 1, payloadLen, isNew)
		}
		return cost
	}
	switch {
	case f.list:
		fs := m.get(f)
		grow := appendGrow
		if f.pointers {
			grow = appendGrowPointers
		}
		var mem int
		mem, fs.cap = grow(fs.len, fs.cap, 1, f.elemSize)
		fs.len++
		cost += mem
	case f.isMap:
		// Every entry is taken as a new key; duplicates are over-charged.
		fs := m.get(f)
		cost += g.allocateMap(fs)
		cost += mapGrowth(fs.entries, f.mapSlot, f.mapPointers)
		fs.entries++
	case f.presence > 0:
		if fs := m.get(f); !fs.present {
			fs.present = true
			cost += f.presence
		}
	case f.oneof != nil:
		m.setOneof(f, m.get(f))
	}
	return cost
}

// packed reallocates the field's slice at exactly its new length: each run is
// a fresh allocation, so many small runs cost quadratically.
func (g generatedModel) packed(m *messageState, f *fieldInfo, count int) int {
	if f.extension {
		// Even an empty run resolves the extension and sets its list.
		fs, isNew := m.add(f)
		return extensionCost(m, f, isNew) + extensionListCost(fs, f, count, 0, isNew)
	}
	if count == 0 {
		return 0
	}
	fs := m.get(f)
	fs.len += count
	fs.cap = fs.len
	return allocSize(fs.len * f.elemSize)
}

// mismatch is what the decoder allocates for a known field before it rejects
// the wire type: a nil map field's map, the wrapper of a oneof member other
// than the set one, which it then drops without changing the oneof, and an
// unset extension's value, each time.
func (g generatedModel) mismatch(m *messageState, f *fieldInfo) int {
	if f.extension {
		cost := extensionMapHeader(m)
		if !f.generatedExt {
			cost += uncachedExtensionCost
		}
		if m.find(f) == nil {
			cost += f.valueCost
		}
		return cost
	}
	if f.isMap {
		return g.allocateMap(m.get(f))
	}
	if f.oneof == nil || m.isSet(f) {
		return 0
	}
	return allocSize(f.elemSize)
}

func (g generatedModel) unknown(m *messageState, info *messageInfo) int {
	if !info.extendable {
		return 0
	}
	return extensionMapHeader(m)
}

func (g generatedModel) allocateMap(fs *fieldState) int {
	if fs.allocated {
		return 0
	}
	fs.allocated = true
	return allocSize(mapHeaderSize)
}
