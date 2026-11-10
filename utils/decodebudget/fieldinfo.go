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
	"sync"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// messageInfo caches what the wire walk asks of a message descriptor, so the
// walk reads plain fields instead of calling protoreflect per occurrence.
type messageInfo struct {
	desc       protoreflect.MessageDescriptor
	layout     messageLayout
	extendable bool
	fields     []fieldInfo // by field index
	byNumber   []int32     // field number -> index+1, when numbers are dense
}

// fieldInfo is a field's descriptor with the facts and fixed costs the
// decoder models need, computed once by the descriptor-based cost functions.
type fieldInfo struct {
	protoreflect.FieldDescriptor
	child *messageInfo
	// oneof is set only for a real oneof member.
	oneof protoreflect.OneofDescriptor
	kind  protoreflect.Kind
	wire  protowire.Type

	elemSize   int
	allocCost  int // fieldCost(field, 0, msgCost)
	presence   int
	msgCost    int
	appendCost int
	checkCost  int
	valueCost  int
	mapSlot    int

	list, isMap, packable, payload, pointers  bool
	extension, generatedExt, mapEntry, mapKey bool
	mapPointers, mapValueMessage              bool
}

// messageInfos assumes one descriptor per name. A schema rebuilt at runtime
// under the same names is still estimated correctly, just not cached.
var messageInfos sync.Map // protoreflect.FullName -> *messageInfo

// infoOf returns desc's cached info. A descriptor that is not the one cached
// under its name gets an uncached info, so the walk never mixes the two.
func infoOf(desc protoreflect.MessageDescriptor) *messageInfo {
	if v, ok := messageInfos.Load(desc.FullName()); ok {
		if mi := v.(*messageInfo); mi.desc == desc {
			return mi
		}
		return buildInfo(desc, map[protoreflect.MessageDescriptor]*messageInfo{})
	}
	built := map[protoreflect.MessageDescriptor]*messageInfo{}
	mi := buildInfo(desc, built)
	for d, b := range built {
		messageInfos.LoadOrStore(d.FullName(), b)
	}
	return mi
}

func buildInfo(desc protoreflect.MessageDescriptor, built map[protoreflect.MessageDescriptor]*messageInfo) *messageInfo {
	if mi, ok := built[desc]; ok {
		return mi
	}
	if v, ok := messageInfos.Load(desc.FullName()); ok && v.(*messageInfo).desc == desc {
		return v.(*messageInfo)
	}
	fields := desc.Fields()
	mi := &messageInfo{
		desc:       desc,
		layout:     newLayout(desc),
		extendable: desc.ExtensionRanges().Len() > 0,
		fields:     make([]fieldInfo, fields.Len()),
	}
	built[desc] = mi
	maxNumber := 0
	for i := range mi.fields {
		field := fields.Get(i)
		var child *messageInfo
		if md := field.Message(); md != nil {
			child = buildInfo(md, built)
		}
		mi.fields[i] = newFieldInfo(field, child)
		maxNumber = max(maxNumber, int(field.Number()))
	}
	if maxNumber <= 2*len(mi.fields)+32 {
		mi.byNumber = make([]int32, maxNumber+1)
		for i := range mi.fields {
			mi.byNumber[mi.fields[i].Number()] = int32(i + 1)
		}
	}
	return mi
}

func (mi *messageInfo) field(number protowire.Number) *fieldInfo {
	if mi.byNumber != nil {
		if uint64(number) < uint64(len(mi.byNumber)) {
			if i := mi.byNumber[number]; i > 0 {
				return &mi.fields[i-1]
			}
		}
		return nil
	}
	if field := mi.desc.Fields().ByNumber(number); field != nil {
		return &mi.fields[field.Index()]
	}
	return nil
}

func newFieldInfo(field protoreflect.FieldDescriptor, child *messageInfo) fieldInfo {
	f := fieldInfo{
		FieldDescriptor: field,
		child:           child,
		kind:            field.Kind(),
		wire:            wireTypeOf(field.Kind()),
		elemSize:        goValueSize(field),
		presence:        presenceCost(field),
		appendCost:      reflectAppendCost(field),
		checkCost:       extensionCheckCost(field),
		list:            field.IsList(),
		isMap:           field.IsMap(),
		packable:        isPackable(field),
		pointers:        hasPointers(field),
		extension:       field.IsExtension(),
		generatedExt:    isGeneratedExtension(field),
		mapEntry:        field.ContainingMessage().IsMapEntry(),
	}
	f.payload = f.kind == protoreflect.StringKind || f.kind == protoreflect.BytesKind
	f.mapKey = f.mapEntry && field.Number() == 1
	if o := field.ContainingOneof(); o != nil && !o.IsSynthetic() {
		f.oneof = o
	}
	if child != nil {
		f.msgCost = messageCost(field, child.layout)
	}
	ownCost := f.msgCost
	if f.isMap {
		ownCost = 0
		if value := child.field(2).child; value != nil {
			ownCost = value.layout.size
		}
	}
	f.allocCost = fieldCost(field, 0, ownCost)
	f.valueCost = extensionValueCost(field, f.msgCost)
	if f.isMap {
		f.mapSlot, f.mapPointers = mapSlot(field)
		f.mapValueMessage = field.MapValue().Message() != nil
	}
	return f
}

func wireTypeOf(kind protoreflect.Kind) protowire.Type {
	switch kind {
	case protoreflect.MessageKind, protoreflect.StringKind, protoreflect.BytesKind:
		return protowire.BytesType
	case protoreflect.GroupKind:
		return protowire.StartGroupType
	case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		return protowire.Fixed32Type
	case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		return protowire.Fixed64Type
	default:
		return protowire.VarintType
	}
}

func (f *fieldInfo) matches(kind protowire.Type) bool {
	return kind == f.wire || (kind == protowire.BytesType && f.packable)
}
