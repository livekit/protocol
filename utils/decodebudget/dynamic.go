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
// messages, which Profile.DynamicMessages prices.
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
	p := defaultProfile
	if b.Profile != nil {
		p = *b.Profile
	}
	p.DynamicMessages = true
	b.Profile = &p
	return b
}

func (p *Profile) rootSize(desc protoreflect.MessageDescriptor) int {
	if p.DynamicMessages {
		return dynamicMessageCost + dynamicResetCost
	}
	return protoStructSize(desc)
}

// swissMapGrowth is the allocation for inserting a map's nth entry, following
// internal/runtime/maps: one group of 8 slots, then a 16-slot table doubling
// at 7/8 load up to 1024 slots, then full tables splitting in two.
func swissMapGrowth(n, slot int) int {
	const groupSlots, maxCapacity, tableSize = 8, 1024, 32
	group := groupSlots + groupSlots*slot
	table := func(capacity int) int {
		size, _ := pointerAlloc(capacity / groupSlots * group)
		return size + tableSize
	}
	const full = maxCapacity * 7 / 8
	switch {
	case n == 1:
		size, _ := pointerAlloc(group)
		return size
	case n <= groupSlots:
		return 0
	case n == groupSlots+1:
		return table(2*groupSlots) + 8 // and a one-table directory
	case n <= full:
		for capacity := 2 * groupSlots; capacity < maxCapacity; capacity *= 2 {
			if n == capacity*7/8+1 {
				return table(2 * capacity)
			}
		}
		return 0
	}
	// Seeded hashes keep tables evenly filled, so the 2^k tables split
	// together as each nears full; charge every such wave at 7/8 of that, with
	// 64 bytes per split for the directory.
	for tables := 1; ; tables *= 2 {
		at := full + 1
		if tables > 1 {
			at = tables*full*7/8 + 1
		}
		switch {
		case n == at:
			return tables * (2*table(maxCapacity) + 64)
		case n < at:
			return 0
		}
	}
}

// pointerAlloc is allocSize for an object holding pointers, which carries an
// 8-byte malloc header between 512 bytes and 32 KiB, and the usable size.
func pointerAlloc(n int) (size, usable int) {
	const minHeader, maxSmall, header = 512, 32768, 8
	if n > minHeader && n <= maxSmall-header {
		size = allocSize(n + header)
		return size, size - header
	}
	size = allocSize(n)
	return size, size
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

// dynamicField is the known-map growth for the first occurrence of field in a
// dynamic message, plus the list, map or extension value Mutable creates then.
func (m *messageState) dynamicField(p *Profile, field protoreflect.FieldDescriptor) (*fieldState, int) {
	n := len(m.fields)
	fs := m.get(field)
	if len(m.fields) == n || field.ContainingMessage().IsMapEntry() {
		return fs, 0
	}
	cost := swissMapGrowth(len(m.fields), knownSlotSize)
	switch {
	case field.IsExtension():
		extensions := 0
		for i := range m.fields {
			if m.fields[i].field.IsExtension() {
				extensions++
			}
		}
		cost += swissMapGrowth(extensions, extSlotSize) + p.ExtensionCost
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

func (m *messageState) dynamicMessageOccurrence(p *Profile, field protoreflect.FieldDescriptor) (int, *messageState) {
	if field.IsList() || field.IsMap() {
		return m.dynamicOccurrence(p, field, 0), &messageState{}
	}
	fs, cost := m.dynamicField(p, field)
	m.setOneof(field, fs)
	if fs.child == nil {
		fs.child = &messageState{}
		// A map value message was charged with its entry.
		if !field.IsExtension() && !field.ContainingMessage().IsMapEntry() {
			cost += dynamicMessageCost
		}
	}
	return cost, fs.child
}

func (m *messageState) dynamicOccurrence(p *Profile, field protoreflect.FieldDescriptor, payloadLen int) int {
	fs, cost := m.dynamicField(p, field)
	if k := field.Kind(); k == protoreflect.StringKind || k == protoreflect.BytesKind {
		cost += allocSize(payloadLen)
	}
	switch {
	case field.IsMap():
		fs.len++
		cost += swissMapGrowth(fs.len, mapEntrySlotSize)
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

func (m *messageState) dynamicPacked(p *Profile, field protoreflect.FieldDescriptor, count int) int {
	fs, cost := m.dynamicField(p, field)
	mem, newCap := appendValues(fs.len, fs.cap, count)
	fs.len += count
	fs.cap = newCap
	return cost + mem + count*dynamicBoxCost(field, 0)
}

// dynamicMismatch is what a field with an unexpected wire type allocates: the
// reflective decoder creates a list or map before checking the wire type.
func (m *messageState) dynamicMismatch(p *Profile, field protoreflect.FieldDescriptor) int {
	if !p.DynamicMessages || field == nil || !(field.IsList() || field.IsMap()) {
		return 0
	}
	_, cost := m.dynamicField(p, field)
	return cost
}
