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
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

// The Go memory layout of decoded values: what a descriptor's generated type
// and field kinds occupy, independent of which decoder fills them.

// messageLayout records whether desc has a generated Go type, which a message
// built at runtime lacks so that only dynamicpb can decode it, and the size of
// one decoded message.
type messageLayout struct {
	generated bool
	size      int
}

var messageLayouts sync.Map // protoreflect.FullName -> messageLayout

func layoutOf(desc protoreflect.MessageDescriptor) messageLayout {
	if v, ok := messageLayouts.Load(desc.FullName()); ok {
		return v.(messageLayout)
	}
	// A dynamicpb message keeps its fields in a map.
	l := messageLayout{size: 64 + 32*desc.Fields().Len()}
	if mt, err := protoregistry.GlobalTypes.FindMessageByName(desc.FullName()); err == nil {
		if t := reflect.TypeOf(mt.Zero().Interface()); t != reflect.TypeFor[*dynamicpb.Message]() && t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Struct {
			l.generated = true
			l.size = int(t.Elem().Size())
		}
	}
	l.size = allocSize(l.size)
	messageLayouts.Store(desc.FullName(), l)
	return l
}

// protoStructSize is the allocation for one message of desc: its generated Go
// struct, or without one, a rough dynamicpb message.
func protoStructSize(desc protoreflect.MessageDescriptor) int {
	return layoutOf(desc).size
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

func hasPointers(field protoreflect.FieldDescriptor) bool {
	switch field.Kind() {
	case protoreflect.StringKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
		return true
	}
	return false
}
