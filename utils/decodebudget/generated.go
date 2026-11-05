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
	key, value := field.MapKey(), field.MapValue()
	slot := (goValueSize(key) + goValueSize(value) + 7) &^ 7
	return mapGrowth(n, slot, hasPointers(key) || hasPointers(value))
}

// mapBoxCost is the allocation for passing one map key or value of field's
// kind through reflect: a scalar is boxed by Interface and again by Convert.
func mapBoxCost(field protoreflect.FieldDescriptor) int {
	if field.Message() != nil {
		return 0
	}
	return 2 * allocSize(goValueSize(field))
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

// extensionCost is a set extension's ExtensionField value, and
// extensionMapSize the map a message allocates for its first extension: a map
// header and one 8-slot group.
const (
	extensionCost    = 64
	extensionMapSize = 528
)

// extensionMapCost is the extension map a generated message allocates when
// field is its first extension; call it before m.get records field.
func extensionMapCost(m *messageState, field protoreflect.FieldDescriptor) int {
	if !field.IsExtension() || m.extensions > 0 {
		return 0
	}
	return extensionMapSize
}

// fieldAllocCost is the allocation for one occurrence of field with a string or
// bytes payload of payloadLen. Nested message contents are charged separately.
func fieldAllocCost(field protoreflect.FieldDescriptor, payloadLen int) int {
	cost := 0
	// A map value message is charged with its entry.
	if field.Message() != nil && !field.IsMap() && !field.ContainingMessage().IsMapEntry() {
		cost += protoStructSize(field.Message())
	}
	if k := field.Kind(); k == protoreflect.StringKind || k == protoreflect.BytesKind {
		cost += allocSize(payloadLen)
	}
	switch {
	case field.IsMap():
		// The decoder allocates a message value before reading the entry,
		// whether or not the entry has one. Map growth depends on earlier
		// entries; the caller charges it.
		cost += mapBoxCost(field.MapKey()) + mapBoxCost(field.MapValue())
		if value := field.MapValue().Message(); value != nil {
			cost += protoStructSize(value)
		}
	case field.IsList():
		// Slice growth depends on earlier occurrences; the caller charges it.
	case field.IsExtension():
		cost += extensionCost
	case field.ContainingOneof() != nil && !field.ContainingOneof().IsSynthetic():
		cost += oneofWrapperCost(field)
	}
	return cost
}

// generatedModel prices protobuf-go's fast-path decoder writing into
// generated Go structs.
type generatedModel struct{}

func (g generatedModel) root(desc protoreflect.MessageDescriptor) int {
	return protoStructSize(desc)
}

func (g generatedModel) message(m *messageState, field protoreflect.FieldDescriptor) (int, *fieldState) {
	if field.IsList() || field.IsMap() {
		// Every element or map value is a new message.
		return g.occurrence(m, field, 0), nil
	}
	cost := extensionMapCost(m, field)
	fs := m.get(field)
	m.setOneof(field, fs)
	if fs.child != nil {
		return cost, fs
	}
	return cost + fieldAllocCost(field, 0), fs
}

func (g generatedModel) occurrence(m *messageState, field protoreflect.FieldDescriptor, payloadLen int) int {
	cost := fieldAllocCost(field, payloadLen) + extensionMapCost(m, field)
	switch {
	case field.IsList():
		fs := m.get(field)
		var mem int
		mem, fs.cap = appendGrow(fs.len, fs.cap, 1, goValueSize(field))
		fs.len++
		cost += mem
	case field.IsMap():
		// Every entry is taken as a new key; duplicates are over-charged.
		fs := m.get(field)
		cost += g.allocateMap(fs)
		cost += mapInsertCost(field, fs.entries)
		fs.entries++
	case presenceCost(field) > 0:
		if fs := m.get(field); !fs.present {
			fs.present = true
			cost += presenceCost(field)
		}
	case field.ContainingOneof() != nil:
		m.setOneof(field, m.get(field))
	}
	return cost
}

// packed reallocates the field's slice at exactly its new length: each run is
// a fresh allocation, so many small runs cost quadratically.
func (g generatedModel) packed(m *messageState, field protoreflect.FieldDescriptor, count int) int {
	if count == 0 {
		return 0
	}
	fs := m.get(field)
	fs.len += count
	fs.cap = fs.len
	return allocSize(fs.len * goValueSize(field))
}

// mismatch is what the decoder allocates for a known field before it rejects
// the wire type: a nil map field's map, and the wrapper of a oneof member other
// than the set one, which it then drops without changing the oneof.
func (g generatedModel) mismatch(m *messageState, field protoreflect.FieldDescriptor) int {
	if field.IsMap() {
		return g.allocateMap(m.get(field))
	}
	oneof := field.ContainingOneof()
	if oneof == nil || oneof.IsSynthetic() || m.isSet(oneof, field) {
		return 0
	}
	return oneofWrapperCost(field)
}

func (g generatedModel) allocateMap(fs *fieldState) int {
	if fs.allocated {
		return 0
	}
	fs.allocated = true
	return allocSize(mapHeaderSize)
}
