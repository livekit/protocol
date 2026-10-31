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

// Package decodebudget bounds protobuf and protojson input before decoding,
// by field count, nesting depth and estimated decoded allocation. It depends
// only on the protobuf runtime so callers need not import utils.
package decodebudget

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// ErrMessageComplexity indicates that a message exceeds its structural budget.
var ErrMessageComplexity = errors.New("message exceeds structural complexity limit")

// Proto limits allocation-driving protobuf structure before decoding.
// Callers must separately limit the serialized byte size. MaxFields and
// MaxDepth must be positive; the zero value does not disable validation.
type Proto struct {
	// MaxFields counts every wire-field occurrence and each packed scalar value,
	// including occurrences of singular and unknown fields.
	MaxFields int
	// MaxDepth counts nested messages, including the root message at depth one.
	MaxDepth int
	// MaxDecodedBytes bounds the estimated bytes the decoder allocates: the Go
	// struct of each message, copied strings and bytes, oneof wrappers, and
	// slice and map growth. Zero leaves it unchecked.
	MaxDecodedBytes int
}

// Validate checks data against the message descriptor without constructing the
// decoded message. Known extensions use protoregistry.GlobalTypes, matching the
// default protobuf decoder. Legacy protobuf groups are deliberately unsupported.
func (b Proto) Validate(data []byte, desc protoreflect.MessageDescriptor) error {
	_, err := b.Estimate(data, desc)
	return err
}

// Estimate validates data like Validate and returns the estimated bytes the
// binary decoder allocates for it, including the root message.
func (b Proto) Estimate(data []byte, desc protoreflect.MessageDescriptor) (int, error) {
	w, err := b.newWalk(desc)
	if err != nil {
		return 0, err
	}
	if err := w.alloc(protoStructSize(desc)); err != nil {
		return w.total(), err
	}
	err = w.message(data, desc, 1)
	return w.total(), err
}

func (b Proto) newWalk(desc protoreflect.MessageDescriptor) (*protoWalk, error) {
	if b.MaxFields <= 0 || b.MaxDepth <= 0 || b.MaxDecodedBytes < 0 {
		return nil, errors.New("protobuf budget limits must be positive")
	}
	if desc == nil {
		return nil, errors.New("protobuf descriptor must not be nil")
	}
	return &protoWalk{budget: b}, nil
}

type protoWalk struct {
	budget Proto
	fields int
	bytes  int
}

func (w *protoWalk) count() error {
	w.fields++
	if w.fields > w.budget.MaxFields {
		return ErrMessageComplexity
	}
	return nil
}

func (w *protoWalk) total() int {
	return withMargin(w.bytes)
}

func (w *protoWalk) alloc(n int) error {
	w.bytes += n
	if w.budget.MaxDecodedBytes > 0 && w.total() > w.budget.MaxDecodedBytes {
		return ErrMessageComplexity
	}
	return nil
}

// repeatedCounts tallies unpacked occurrences of each repeated field in one
// message, so slice growth is charged once for the final length.
type repeatedCounts []struct {
	field protoreflect.FieldDescriptor
	n     int
}

func (r *repeatedCounts) add(field protoreflect.FieldDescriptor) {
	for i := range *r {
		if (*r)[i].field == field {
			(*r)[i].n++
			return
		}
	}
	*r = append(*r, struct {
		field protoreflect.FieldDescriptor
		n     int
	}{field, 1})
}

func (w *protoWalk) growth(r repeatedCounts) error {
	for _, c := range r {
		if err := w.alloc(appendAllocBytes(c.n, goValueSize(c.field))); err != nil {
			return err
		}
	}
	return nil
}

func (w *protoWalk) message(data []byte, desc protoreflect.MessageDescriptor, depth int) error {
	if depth > w.budget.MaxDepth {
		return ErrMessageComplexity
	}
	var repeated repeatedCounts
	for len(data) > 0 {
		if err := w.count(); err != nil {
			return err
		}
		number, kind, n := protowire.ConsumeTag(data)
		if n < 0 {
			return protowire.ParseError(n)
		}
		data = data[n:]
		if kind == protowire.StartGroupType || kind == protowire.EndGroupType {
			return ErrMessageComplexity
		}
		field := desc.Fields().ByNumber(number)
		if field == nil && desc.ExtensionRanges().Has(number) {
			extension, err := protoregistry.GlobalTypes.FindExtensionByNumber(desc.FullName(), number)
			if err == nil {
				field = extension.TypeDescriptor()
			} else if !errors.Is(err, protoregistry.NotFound) {
				return err
			}
		}
		if kind != protowire.BytesType {
			size := protowire.ConsumeFieldValue(number, kind, data)
			if size < 0 {
				return protowire.ParseError(size)
			}
			data = data[size:]
			cost := unknownFieldCost(n + size)
			if field != nil {
				cost = fieldAllocCost(field, 0)
				if field.IsList() {
					repeated.add(field)
				}
			}
			if err := w.alloc(cost); err != nil {
				return err
			}
			continue
		}

		value, size := protowire.ConsumeBytes(data)
		if size < 0 {
			return protowire.ParseError(size)
		}
		data = data[size:]
		switch {
		case field == nil:
			if err := w.alloc(unknownFieldCost(n + size)); err != nil {
				return err
			}
		case field.Message() != nil:
			if err := w.alloc(fieldAllocCost(field, 0)); err != nil {
				return err
			}
			if err := w.message(value, field.Message(), depth+1); err != nil {
				return err
			}
		case isPackableProtoField(field):
			if err := w.packed(value, field); err != nil {
				return err
			}
		default:
			if err := w.alloc(fieldAllocCost(field, len(value))); err != nil {
				return err
			}
		}
		if field != nil && field.IsList() && (!isPackableProtoField(field) || kind != protowire.BytesType) {
			repeated.add(field)
		}
	}
	return w.growth(repeated)
}

// packed walks a packed scalar run; decoders accept one even when the schema
// prefers unpacked.
func (w *protoWalk) packed(value []byte, field protoreflect.FieldDescriptor) error {
	count := 0
	for len(value) > 0 {
		var consumed int
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
			_, consumed = protowire.ConsumeFixed64(value)
		case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
			_, consumed = protowire.ConsumeFixed32(value)
		default:
			_, consumed = protowire.ConsumeVarint(value)
		}
		if consumed < 0 {
			return protowire.ParseError(consumed)
		}
		if err := w.count(); err != nil {
			return err
		}
		count++
		value = value[consumed:]
	}
	return w.alloc(packedCost(field, count))
}

func isPackableProtoField(field protoreflect.FieldDescriptor) bool {
	if field == nil || !field.IsList() {
		return false
	}
	switch field.Kind() {
	case protoreflect.BoolKind, protoreflect.EnumKind,
		protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
		protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind,
		protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		return true
	default:
		return false
	}
}

// Unmarshal validates the budget before decoding into a non-nil mutable message.
// It discards unknown fields, uses MaxDepth as the decoder recursion limit, and
// retains the binary decoder's required-field checks. Preflight rejection leaves
// the destination unchanged; other decode errors follow normal protobuf semantics.
func (b Proto) Unmarshal(data []byte, msg proto.Message) error {
	if msg == nil || (reflect.ValueOf(msg).Kind() == reflect.Pointer && reflect.ValueOf(msg).IsNil()) {
		return errors.New("protobuf message must be non-nil and mutable")
	}
	reflected := msg.ProtoReflect()
	if !reflected.IsValid() {
		return errors.New("protobuf message must be non-nil and mutable")
	}
	if err := b.Validate(data, reflected.Descriptor()); err != nil {
		return err
	}
	return (proto.UnmarshalOptions{DiscardUnknown: true, RecursionLimit: b.MaxDepth}).Unmarshal(data, msg)
}

// JSON limits JSON structure before a downstream decoder constructs
// collections. Callers must separately limit the serialized byte size. Both
// limits must be positive; the zero value does not disable validation.
type JSON struct {
	// MaxTokens includes scalar values, object keys, and opening/closing delimiters.
	MaxTokens int
	// MaxDepth counts nested arrays/objects, including a root container at depth one.
	MaxDepth int
}

// Validate accepts exactly one complete JSON value, including scalar roots.
func (b JSON) Validate(data []byte) error {
	if b.MaxTokens <= 0 || b.MaxDepth <= 0 {
		return errors.New("JSON budget limits must be positive")
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	depth, tokens := 0, 0
	seen, complete := false, false
	for {
		token, err := dec.Token()
		if err == io.EOF {
			if !seen || depth != 0 {
				return io.ErrUnexpectedEOF
			}
			return nil
		}
		if err != nil {
			return err
		}
		if complete {
			return errors.New("unexpected value after JSON document")
		}
		tokens++
		if tokens > b.MaxTokens {
			return ErrMessageComplexity
		}
		seen = true
		if d, ok := token.(json.Delim); ok {
			switch d {
			case '{', '[':
				depth++
			case '}', ']':
				depth--
			}
			if depth > b.MaxDepth {
				return ErrMessageComplexity
			}
		}
		if depth == 0 {
			complete = true
		}
	}
}
