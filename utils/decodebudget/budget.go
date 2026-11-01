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

// Package decodebudget bounds untrusted input before decoding. Proto checks
// protobuf and protojson input by field count, nesting depth and estimated
// decoded allocation; JSON checks plain JSON by token count and depth. It
// depends only on the protobuf runtime so callers need not import utils.
package decodebudget

import (
	"errors"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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
	t, err := b.newTally(desc)
	if err != nil {
		return 0, err
	}
	if err := t.alloc(protoStructSize(desc)); err != nil {
		return t.total(), err
	}
	err = (&wireWalk{t}).message(data, desc, 1)
	return t.total(), err
}

func (b Proto) newTally(desc protoreflect.MessageDescriptor) (*tally, error) {
	if b.MaxFields <= 0 || b.MaxDepth <= 0 || b.MaxDecodedBytes < 0 {
		return nil, errors.New("protobuf budget limits must be positive")
	}
	if desc == nil {
		return nil, errors.New("protobuf descriptor must not be nil")
	}
	return &tally{budget: b}, nil
}

// tally is the running field count and allocation estimate shared by the wire
// and protojson walks.
type tally struct {
	budget Proto
	fields int
	bytes  int
}

func (t *tally) count() error {
	t.fields++
	if t.fields > t.budget.MaxFields {
		return ErrMessageComplexity
	}
	return nil
}

func (t *tally) total() int {
	return withMargin(t.bytes)
}

func (t *tally) alloc(n int) error {
	t.bytes += n
	if t.budget.MaxDecodedBytes > 0 && t.total() > t.budget.MaxDecodedBytes {
		return ErrMessageComplexity
	}
	return nil
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
