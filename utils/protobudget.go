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

package utils

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

// ProtoBudget limits allocation-driving protobuf structure before decoding.
// Callers must separately limit the serialized byte size. Both limits must be
// positive; the zero value does not disable validation.
type ProtoBudget struct {
	// MaxFields counts every wire-field occurrence and each packed scalar value,
	// including occurrences of singular and unknown fields.
	MaxFields int
	// MaxDepth counts nested messages, including the root message at depth one.
	MaxDepth int
}

// Validate checks data against the message descriptor without constructing the
// decoded message. Known extensions use protoregistry.GlobalTypes, matching the
// default protobuf decoder. Legacy protobuf groups are deliberately unsupported.
func (b ProtoBudget) Validate(data []byte, desc protoreflect.MessageDescriptor) error {
	if b.MaxFields <= 0 || b.MaxDepth <= 0 {
		return errors.New("protobuf budget limits must be positive")
	}
	if desc == nil {
		return errors.New("protobuf descriptor must not be nil")
	}
	remaining := b.MaxFields
	return b.validateFields(data, desc, &remaining, 1)
}

func (b ProtoBudget) validateFields(data []byte, desc protoreflect.MessageDescriptor, remaining *int, depth int) error {
	if depth > b.MaxDepth {
		return ErrMessageComplexity
	}
	for len(data) > 0 {
		*remaining--
		if *remaining < 0 {
			return ErrMessageComplexity
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
		if kind == protowire.BytesType {
			value, size := protowire.ConsumeBytes(data)
			if size < 0 {
				return protowire.ParseError(size)
			}
			if field != nil && field.Message() != nil {
				if err := b.validateFields(value, field.Message(), remaining, depth+1); err != nil {
					return err
				}
			} else if isPackableProtoField(field) {
				// Decoders accept packed input even when the schema prefers unpacked.
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
					*remaining--
					if *remaining < 0 {
						return ErrMessageComplexity
					}
					value = value[consumed:]
				}
			}
			data = data[size:]
		} else {
			size := protowire.ConsumeFieldValue(number, kind, data)
			if size < 0 {
				return protowire.ParseError(size)
			}
			data = data[size:]
		}
	}
	return nil
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
func (b ProtoBudget) Unmarshal(data []byte, msg proto.Message) error {
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

// JSONBudget limits JSON structure before a downstream decoder constructs
// collections. Callers must separately limit the serialized byte size. Both
// limits must be positive; the zero value does not disable validation.
type JSONBudget struct {
	// MaxTokens includes scalar values, object keys, and opening/closing delimiters.
	MaxTokens int
	// MaxDepth counts nested arrays/objects, including a root container at depth one.
	MaxDepth int
}

// Validate accepts exactly one complete JSON value, including scalar roots.
func (b JSONBudget) Validate(data []byte) error {
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
