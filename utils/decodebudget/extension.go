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
// extension map (its header, then one group of 8 slots, then 56-byte slots), a
// repeated extension's list, the field info rebuilt per occurrence of an extension type it
// cannot cache, such as dynamicpb's, and for such a type on a dynamicpb
// message, the descriptor resolving it allocates and the type check of its
// singular values.
const (
	extensionGroupSize          = 480
	extensionMapSize            = 48 + extensionGroupSize
	extensionSlotSize           = 56
	extensionListSize           = 72
	uncachedExtensionCost       = 96 + 16
	dynamicExtensionResolveCost = 16
	dynamicExtensionCheckCost   = 48
)

func isGeneratedExtension(field protoreflect.FieldDescriptor) bool {
	xd, ok := field.(protoreflect.ExtensionTypeDescriptor)
	if !ok {
		return false
	}
	_, ok = xd.Type().(*protoimpl.ExtensionInfo)
	return ok
}

// extensionMapHeader is the extension map a generated message creates for the
// first field number it has no field for, extension or not.
func extensionMapHeader(m *messageState) int {
	if m.extensionMap {
		return 0
	}
	m.extensionMap = true
	return allocSize(mapHeaderSize)
}

// extensionCost is what a generated message allocates for an occurrence of
// extension field beyond its value: the field info of a type it cannot cache,
// and map growth when the field is new.
func extensionCost(m *messageState, field protoreflect.FieldDescriptor, isNew bool) int {
	cost := extensionMapHeader(m)
	if !isGeneratedExtension(field) {
		cost += uncachedExtensionCost
	}
	switch {
	case !isNew:
	case m.extensions == 1:
		cost += extensionGroupSize
	default:
		cost += mapGrowth(m.extensions-1, extensionSlotSize, true)
	}
	return cost
}

// extensionValueCost is the value a decoder creates for an extension that is
// not yet set before it reads the wire: a message or a list.
func extensionValueCost(field protoreflect.FieldDescriptor) int {
	switch {
	case field.IsList():
		return extensionListSize
	case field.Message() == nil:
		return 0
	case isGeneratedExtension(field):
		return protoStructSize(field.Message())
	}
	return dynamicMessageCost
}

// extensionCheckCost is a dynamicpb message type-checking a value of a
// dynamicpb extension type.
func extensionCheckCost(field protoreflect.FieldDescriptor) int {
	switch {
	case !field.IsExtension() || isGeneratedExtension(field):
		return 0
	case field.IsList():
		return dynamicExtensionResolveCost
	}
	return dynamicExtensionCheckCost
}

// extensionListCost is appending count elements one at a time to a repeated
// extension, whose list layout follows its extension type, not the message.
// isNew charges the list itself.
func extensionListCost(fs *fieldState, field protoreflect.FieldDescriptor, count, payloadLen int, isNew bool) int {
	cost := 0
	if isNew {
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
	return cost + count*reflectAppendCost(field)
}
