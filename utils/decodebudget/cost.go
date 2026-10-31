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
// not let its decode allocate more than the estimate.
const (
	// Map tables grow by doubling and are at most 7/8 full.
	mapGrowthFactor  = 4
	mapEntryOverhead = 16
	extensionCost    = 64
	// estimateMarginPercent covers allocations the model does not itemize; the
	// calibration tests measured them at under 3%.
	estimateMarginPercent = 5
)

func withMargin(n int) int {
	return n + n*estimateMarginPercent/100
}

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

// appendAllocBytes is the total size of the backing arrays append allocates
// while growing a nil slice to n elements, following runtime.growslice.
func appendAllocBytes(n, elem int) int {
	total, capacity := 0, 0
	for capacity < n {
		newLen, newCap := capacity+1, capacity+capacity
		switch {
		case newLen > newCap:
			newCap = newLen
		case capacity >= 256:
			newCap = capacity
			for newCap < newLen {
				newCap += (newCap + 3*256) >> 2
			}
		}
		mem := allocSize(newCap * elem)
		total += mem
		capacity = mem / elem
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

// packedCost is exact: the decoder sizes a packed run's slice once.
func packedCost(field protoreflect.FieldDescriptor, count int) int {
	return allocSize(count * goValueSize(field))
}

func mapEntryCost(field protoreflect.FieldDescriptor) int {
	return mapGrowthFactor*(goValueSize(field.MapKey())+goValueSize(field.MapValue())) + mapEntryOverhead
}

func oneofWrapperCost(field protoreflect.FieldDescriptor) int {
	switch field.Kind() {
	case protoreflect.StringKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
		return allocSize(goValueSize(field))
	default:
		return allocSize(goValueSize(field))
	}
}

// unknownFieldCost covers retaining n raw bytes in a growing unknown-fields slice.
func unknownFieldCost(n int) int {
	return 2 * n
}

// fieldAllocCost is the allocation for one occurrence of field with a string or
// bytes payload of payloadLen. Nested message contents are charged separately.
func fieldAllocCost(field protoreflect.FieldDescriptor, payloadLen int) int {
	cost := 0
	if field.Message() != nil && !field.IsMap() {
		cost += protoStructSize(field.Message())
	}
	if k := field.Kind(); k == protoreflect.StringKind || k == protoreflect.BytesKind {
		cost += allocSize(payloadLen)
	}
	switch {
	case field.IsMap():
		cost += mapEntryCost(field)
	case field.IsList():
		// Slice growth is charged per message by protoWalk.
	case field.IsExtension():
		cost += extensionCost
	case field.ContainingOneof() != nil && !field.ContainingOneof().IsSynthetic():
		cost += oneofWrapperCost(field)
	}
	return cost
}
