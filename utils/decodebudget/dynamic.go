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

// isDynamicDestination reports whether msg takes protobuf-go's reflective
// decoder, as a dynamicpb message does.
func isDynamicDestination(msg protoreflect.Message) bool {
	methods := msg.ProtoMethods()
	return methods == nil || methods.Unmarshal == nil || methods.Flags&protoiface.SupportUnmarshalDiscardUnknown == 0
}

// dynamicBoxCost is the allocation for boxing one value of field into an
// interface, which dynamicpb does to type-check every value it stores.
func dynamicBoxCost(field protoreflect.FieldDescriptor, payloadLen int) int {
	return kindBoxCost(field.Kind(), payloadLen)
}

func kindBoxCost(kind protoreflect.Kind, payloadLen int) int {
	switch kind {
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

func (d dynamicModel) root(*messageInfo) int {
	return dynamicMessageCost + dynamicResetCost
}

// field is the known-map growth for the first occurrence of field in a
// dynamic message, plus the list, map or extension value Mutable creates then.
func (d dynamicModel) field(m *messageState, f *fieldInfo) (fs *fieldState, isNew bool, cost int) {
	fs, isNew = m.add(f)
	if !isNew || f.mapEntry {
		return fs, isNew, 0
	}
	cost = mapGrowth(len(m.fields)-1, knownSlotSize, true)
	switch {
	case f.extension:
		cost += mapGrowth(m.extensions-1, extSlotSize, true)
		if f.child != nil && !f.list {
			cost += f.msgCost
		}
	case f.list:
		cost += dynamicListCost
	case f.isMap:
		cost += dynamicMapCost
	}
	return fs, isNew, cost
}

func (d dynamicModel) message(m *messageState, f *fieldInfo) (int, *fieldState) {
	if f.list || f.isMap {
		return d.occurrence(m, f, 0), nil
	}
	fs, _, cost := d.field(m, f)
	cost += f.checkCost
	m.setOneof(f, fs)
	// A map value message was charged with its entry.
	if fs.child == nil && !f.extension && !f.mapEntry {
		cost += dynamicMessageCost
	}
	return cost, fs
}

func (d dynamicModel) occurrence(m *messageState, f *fieldInfo, payloadLen int) int {
	fs, isNew, cost := d.field(m, f)
	if f.payload {
		cost += allocSize(payloadLen)
	}
	switch {
	case f.isMap:
		cost += mapGrowth(fs.entries, mapEntrySlotSize, true)
		fs.entries++
		if f.mapValueMessage {
			cost += dynamicMessageCost
		}
	case f.list && f.extension:
		cost += extensionListCost(fs, f, 1, payloadLen, isNew) + f.checkCost
		if f.child != nil {
			cost += f.msgCost
		}
	case f.list:
		var mem int
		mem, fs.cap = appendValues(fs.len, fs.cap, 1)
		fs.len++
		cost += mem + kindBoxCost(f.kind, payloadLen)
		if f.child != nil {
			cost += dynamicMessageCost
		}
	default:
		cost += kindBoxCost(f.kind, payloadLen) + f.checkCost
		if f.mapKey {
			// The map boxes its key again to store it.
			cost += kindBoxCost(f.kind, payloadLen)
		}
		m.setOneof(f, fs)
	}
	return cost
}

func (d dynamicModel) packed(m *messageState, f *fieldInfo, count int) int {
	fs, isNew, cost := d.field(m, f)
	if f.extension {
		return cost + extensionListCost(fs, f, count, 0, isNew) + f.checkCost
	}
	mem, newCap := appendValues(fs.len, fs.cap, count)
	fs.len += count
	fs.cap = newCap
	return cost + mem + count*kindBoxCost(f.kind, 0)
}

// mismatch is what a field with an unexpected wire type allocates: the
// reflective decoder resolves an extension and creates a list or map before
// checking the wire type.
func (d dynamicModel) mismatch(m *messageState, f *fieldInfo) int {
	cost := 0
	if f.extension && !f.generatedExt {
		cost = dynamicExtensionResolveCost
	}
	if !f.list && !f.isMap {
		return cost
	}
	_, isNew, fieldCost := d.field(m, f)
	if isNew && f.extension {
		cost += extensionListSize
	}
	return cost + fieldCost
}

func (d dynamicModel) unknown(*messageState, *messageInfo) int {
	return 0
}
