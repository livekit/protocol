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
	"reflect"
	"slices"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

// The cost model overestimates on purpose: a budget that admits a payload must
// not let its decode allocate more than the estimate. Its tunable constants
// live in Profile.

// Go runtime size classes; an allocation is rounded up to the next one.
var goSizeClasses = []int{
	8, 16, 24, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256,
	288, 320, 352, 384, 416, 448, 480, 512, 576, 640, 704, 768, 896, 1024, 1152, 1280,
	1408, 1536, 1792, 2048, 2304, 2688, 3072, 3200, 3456, 4096, 4864, 5376, 6144, 6528,
	6784, 6912, 8192, 9472, 9728, 10240, 10880, 12288, 13568, 14336, 16384, 18432,
	19072, 20480, 21760, 24576, 27264, 28672, 32768,
}

func allocSize(n int) int {
	if n <= 0 {
		return 0
	}
	if i, _ := slices.BinarySearch(goSizeClasses, n); i < len(goSizeClasses) {
		return goSizeClasses[i]
	}
	const page = 8192
	return (n + page - 1) / page * page
}

// tmpCopySize is the allocation for a temporary copy of n bytes; one that does
// not escape and fits runtime.tmpBuf stays on the stack.
func tmpCopySize(n int) int {
	if n <= 32 {
		return 0
	}
	return allocSize(n)
}

// appendGrow is the allocation for appending add elements of elem bytes to a
// slice of length l and capacity c, and the resulting capacity, following
// runtime.growslice.
func appendGrow(l, c, add, elem int) (mem, newCap int) {
	newLen := l + add
	if newLen <= c {
		return 0, c
	}
	newCap = c + c
	switch {
	case newLen > newCap:
		newCap = newLen
	case c >= 256:
		newCap = c
		for newCap < newLen {
			newCap += (newCap + 3*256) >> 2
		}
	}
	mem = allocSize(newCap * elem)
	return mem, mem / elem
}

// appendAllocBytes is the total size of the backing arrays append allocates
// while growing a nil slice to n elements one at a time.
func appendAllocBytes(n, elem int) int {
	total, c := 0, 0
	for l := 0; l < n; l++ {
		mem, newCap := appendGrow(l, c, 1, elem)
		total += mem
		c = newCap
	}
	return total
}

var protoStructSizes sync.Map // protoreflect.FullName -> int

// protoStructSize is the allocation for one message of desc's generated Go type.
func protoStructSize(desc protoreflect.MessageDescriptor) int {
	if v, ok := protoStructSizes.Load(desc.FullName()); ok {
		return v.(int)
	}
	// A dynamicpb message keeps its fields in a map.
	size := 64 + 32*desc.Fields().Len()
	if mt, err := protoregistry.GlobalTypes.FindMessageByName(desc.FullName()); err == nil {
		if t := reflect.TypeOf(mt.Zero().Interface()); t != reflect.TypeFor[*dynamicpb.Message]() && t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Struct {
			size = int(t.Elem().Size())
		}
	}
	size = allocSize(size)
	protoStructSizes.Store(desc.FullName(), size)
	return size
}

// goValueSize is the size of one value of field's kind as a slice element, map
// slot or oneof wrapper field.
func goValueSize(field protoreflect.FieldDescriptor) int {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return 1
	case protoreflect.EnumKind, protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		return 4
	case protoreflect.StringKind:
		return 16
	case protoreflect.BytesKind:
		return 24
	default:
		// 64-bit scalars and message pointers.
		return 8
	}
}

// Go's map layout (internal/runtime/maps): a header, then one group of 8
// slots, then a table that doubles at 7/8 load up to 1024 slots, after which
// each full table splits into two new ones.
const (
	mapHeaderSize    = 48
	mapTableSize     = 32
	mapGroupSlots    = 8
	mapMaxTableSlots = 1024
)

// mapBoxCost is the allocation for passing one map key or value of field's
// kind through reflect: a scalar is boxed by Interface and again by Convert.
func mapBoxCost(field protoreflect.FieldDescriptor) int {
	if field.Message() != nil {
		return 0
	}
	return 2 * allocSize(goValueSize(field))
}

// mapInsertCost is the allocation for inserting a new key into a map of
// field's type that holds n entries, excluding the map header.
func mapInsertCost(field protoreflect.FieldDescriptor, n int) int {
	slot := (goValueSize(field.MapKey()) + goValueSize(field.MapValue()) + 7) &^ 7
	group := 8 + mapGroupSlots*slot
	table := func(slots int) int {
		return allocSize(mapTableSize) + allocSize(slots/mapGroupSlots*group)
	}
	switch {
	case n == 0:
		return allocSize(group)
	case n < mapGroupSlots:
		return 0
	case n == mapGroupSlots:
		// The first table comes with a one-entry directory.
		return table(2*mapGroupSlots) + allocSize(8)
	case n < mapMaxTableSlots*7/8:
		for c := 2 * mapGroupSlots; c < mapMaxTableSlots; c *= 2 {
			if n == c*7/8 {
				return table(2 * c)
			}
		}
		return 0
	}
	return mapSplitCost(n, table(mapMaxTableSlots))
}

// mapSplitCost charges the round of splits that takes a map from t to 2t
// full-size tables, plus the directory that doubles for it. Tables fill
// unevenly, so a round with several tables is spread over the inserts from
// 3/4 to 7/8 average load.
func mapSplitCost(n, table int) int {
	t := 1
	for n >= 2*t*mapMaxTableSlots*3/4 {
		t *= 2
	}
	full := t * mapMaxTableSlots * 7 / 8
	begin := full
	if t > 1 {
		begin = t * mapMaxTableSlots * 3 / 4
	}
	if n < begin || n > full {
		return 0
	}
	round := 2*t*table + allocSize(2*t*8)
	span := full - begin + 1
	return round*(n-begin+1)/span - round*(n-begin)/span
}

func oneofWrapperCost(field protoreflect.FieldDescriptor) int {
	switch field.Kind() {
	case protoreflect.StringKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
		return allocSize(goValueSize(field))
	default:
		return allocSize(goValueSize(field))
	}
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

// fieldAllocCost is the allocation for one occurrence of field with a string or
// bytes payload of payloadLen. Nested message contents are charged separately.
func (p *Profile) fieldAllocCost(field protoreflect.FieldDescriptor, payloadLen int) int {
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
		cost += p.ExtensionCost
	case field.ContainingOneof() != nil && !field.ContainingOneof().IsSynthetic():
		cost += oneofWrapperCost(field)
	}
	return cost
}
