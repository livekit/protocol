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
	"google.golang.org/protobuf/runtime/protoimpl"
)

// Extension allocations of protobuf-go's binary decoder: a generated message's
// extension map (a header and one group of 8 slots, then 56-byte slots), a
// repeated extension's list and the slice header reflect.Append allocates per
// element, the field info rebuilt per occurrence of an extension type it
// cannot cache, such as dynamicpb's, and dynamicpb's type check of such a
// type's values.
const (
	extensionMapSize          = 528
	extensionSlotSize         = 56
	extensionListSize         = 72
	extensionAppendCost       = 24
	uncachedExtensionCost     = 96 + 16
	dynamicExtensionCheckCost = 48
	dynamicExtensionListCheck = 16
)

func isGeneratedExtension(field protoreflect.FieldDescriptor) bool {
	xd, ok := field.(protoreflect.ExtensionTypeDescriptor)
	if !ok {
		return false
	}
	_, ok = xd.Type().(*protoimpl.ExtensionInfo)
	return ok
}

// extensionCost is what a generated message allocates for an occurrence of
// extension field beyond its value: map growth when field is new, so call it
// before m.get records field, and the field info of a type it cannot cache.
func extensionCost(m *messageState, field protoreflect.FieldDescriptor) int {
	cost := 0
	if !isGeneratedExtension(field) {
		cost = uncachedExtensionCost
	}
	switch {
	case m.has(field):
	case m.extensions == 0:
		cost += extensionMapSize
	default:
		cost += mapGrowth(m.extensions, extensionSlotSize, true)
	}
	return cost
}

// extensionCheckCost is a dynamicpb message type-checking a value of a
// dynamicpb extension type.
func extensionCheckCost(field protoreflect.FieldDescriptor) int {
	switch {
	case !field.IsExtension() || isGeneratedExtension(field):
		return 0
	case field.IsList():
		return dynamicExtensionListCheck
	}
	return dynamicExtensionCheckCost
}

// extensionListCost is appending count elements one at a time to a repeated
// extension, whose list layout follows its extension type, not the message.
func extensionListCost(fs *fieldState, field protoreflect.FieldDescriptor, count, payloadLen int) int {
	cost := 0
	if fs.len == 0 {
		cost = extensionListSize
	}
	if !isGeneratedExtension(field) {
		mem, newCap := appendValues(fs.len, fs.cap, count)
		fs.len += count
		fs.cap = newCap
		return cost + mem + count*dynamicBoxCost(field, payloadLen)
	}
	size := goValueSize(field)
	for fs.len += count; fs.cap < fs.len; {
		mem := allocSize(nextSliceCap(fs.cap+1, fs.cap) * size)
		cost += mem
		fs.cap = mem / size
	}
	box := 0
	if field.Message() == nil {
		// Boxed scalars come from the tiny allocator, unrounded.
		box = size
	}
	return cost + count*(extensionAppendCost+box)
}
