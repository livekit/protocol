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
	"google.golang.org/protobuf/runtime/protoiface"
)

// Allocations of protobuf-go's reflective decoder writing into dynamicpb
// messages.
const (
	// The Message struct, its known and ext map headers, and the closure
	// checkInitialized passes to its Range.
	dynamicMessageCost = 64 + 2*48 + 32
	// Unmarshal resets the root message, replacing both maps.
	dynamicResetCost = 2 * 48
	dynamicListCost  = 48
	dynamicMapCost   = 24 + 48
	// Lists hold protoreflect.Values.
	dynamicValueSize = 24
	// Swiss map slots: FieldNumber to Value, any to Value, FieldNumber to
	// FieldDescriptor.
	knownSlotSize    = 32
	mapEntrySlotSize = 40
	extSlotSize      = 24
)

// forDestination prices the dynamic layout when msg takes protobuf-go's
// reflective decoder, as a dynamicpb message does.
func (b Proto) forDestination(msg protoreflect.Message) Proto {
	methods := msg.ProtoMethods()
	if methods != nil && methods.Unmarshal != nil && methods.Flags&protoiface.SupportUnmarshalDiscardUnknown != 0 {
		return b
	}
	b.dynamic = true
	return b
}

// dynamicBoxCost is the allocation for boxing one value of field into an
// interface, which dynamicpb does to type-check every value it stores.
func dynamicBoxCost(field protoreflect.FieldDescriptor, payloadLen int) int {
	switch field.Kind() {
	case protoreflect.BoolKind, protoreflect.MessageKind, protoreflect.GroupKind:
		return 0
	case protoreflect.StringKind:
		if payloadLen == 0 {
			return 0
		}
		return 16
	case protoreflect.BytesKind:
		return 24
	default:
		return 8
	}
}

// appendValues is the allocation for appending add Values one at a time to a
// slice of length l and capacity c, and the resulting capacity.
func appendValues(l, c, add int) (mem, newCap int) {
	for l+add > c {
		size, usable := pointerAlloc(nextSliceCap(c+1, c) * dynamicValueSize)
		mem += size
		c = usable / dynamicValueSize
	}
	return mem, c
}

// dynamicModel prices protobuf-go's reflective decoder writing into dynamicpb
// messages.
type dynamicModel struct{}

func (d dynamicModel) root(protoreflect.MessageDescriptor) int {
	return dynamicMessageCost + dynamicResetCost
}

// field is the known-map growth for the first occurrence of field in a
// dynamic message, plus the list, map or extension value Mutable creates then.
func (d dynamicModel) field(m *messageState, field protoreflect.FieldDescriptor) (*fieldState, int) {
	n := len(m.fields)
	fs := m.get(field)
	if len(m.fields) == n || field.ContainingMessage().IsMapEntry() {
		return fs, 0
	}
	cost := mapGrowth(n, knownSlotSize, true)
	switch {
	case field.IsExtension():
		extensions := 0
		for i := range m.fields[:n] {
			if m.fields[i].field.IsExtension() {
				extensions++
			}
		}
		cost += mapGrowth(extensions, extSlotSize, true) + extensionCost
		if md := field.Message(); md != nil && !field.IsList() {
			// Extension types are generated, so their messages are too.
			cost += max(dynamicMessageCost, protoStructSize(md))
		}
	case field.IsList():
		cost += dynamicListCost
	case field.IsMap():
		cost += dynamicMapCost
	}
	return fs, cost
}

func (d dynamicModel) message(m *messageState, field protoreflect.FieldDescriptor) (int, *fieldState) {
	if field.IsList() || field.IsMap() {
		return d.occurrence(m, field, 0), nil
	}
	fs, cost := d.field(m, field)
	m.setOneof(field, fs)
	// A map value message was charged with its entry.
	if fs.child == nil && !field.IsExtension() && !field.ContainingMessage().IsMapEntry() {
		cost += dynamicMessageCost
	}
	return cost, fs
}

func (d dynamicModel) occurrence(m *messageState, field protoreflect.FieldDescriptor, payloadLen int) int {
	fs, cost := d.field(m, field)
	if k := field.Kind(); k == protoreflect.StringKind || k == protoreflect.BytesKind {
		cost += allocSize(payloadLen)
	}
	switch {
	case field.IsMap():
		cost += mapGrowth(fs.entries, mapEntrySlotSize, true)
		fs.entries++
		if field.MapValue().Message() != nil {
			cost += dynamicMessageCost
		}
	case field.IsList():
		var mem int
		mem, fs.cap = appendValues(fs.len, fs.cap, 1)
		fs.len++
		cost += mem + dynamicBoxCost(field, payloadLen)
		if field.Message() != nil {
			cost += dynamicMessageCost
		}
	default:
		cost += dynamicBoxCost(field, payloadLen)
		if field.ContainingMessage().IsMapEntry() && field.Number() == 1 {
			// The map boxes its key again to store it.
			cost += dynamicBoxCost(field, payloadLen)
		}
		m.setOneof(field, fs)
	}
	return cost
}

func (d dynamicModel) packed(m *messageState, field protoreflect.FieldDescriptor, count int) int {
	fs, cost := d.field(m, field)
	mem, newCap := appendValues(fs.len, fs.cap, count)
	fs.len += count
	fs.cap = newCap
	return cost + mem + count*dynamicBoxCost(field, 0)
}

// mismatch is what a field with an unexpected wire type allocates: the
// reflective decoder creates a list or map before checking the wire type.
func (d dynamicModel) mismatch(m *messageState, field protoreflect.FieldDescriptor) int {
	if !field.IsList() && !field.IsMap() {
		return 0
	}
	_, cost := d.field(m, field)
	return cost
}
