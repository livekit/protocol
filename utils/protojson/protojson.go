// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protojson

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type MarshalOptions = protojson.MarshalOptions
type UnmarshalOptions = protojson.UnmarshalOptions

// Format formats the message as a multiline string.
// This function is only intended for human consumption and ignores errors.
// Do not depend on the output being stable. Its output will change across
// different builds of your program, even when using the same version of the
// protobuf module.
func Format(m proto.Message) string {
	return MarshalOptions{
		Multiline:    true,
		AllowPartial: true,
	}.Format(m)
}

// Marshal writes the given [proto.Message] in JSON format using default options.
// Do not depend on the output being stable. Its output will change across
// different builds of your program, even when using the same version of the
// protobuf module.
func Marshal(m proto.Message) ([]byte, error) {
	return MarshalOptions{
		AllowPartial: true,
	}.Marshal(m)
}

// Unmarshal reads the given []byte into the given [proto.Message].
// The provided message must be mutable (e.g., a non-nil pointer to a message).
func Unmarshal(b []byte, m proto.Message) error {
	return UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}.Unmarshal(b, m)
}
