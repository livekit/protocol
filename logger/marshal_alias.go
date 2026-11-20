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

//go:build !logger_own_marshalers

package logger

import "go.uber.org/zap/zapcore"

// See marshal.go.
type (
	ObjectEncoder       = zapcore.ObjectEncoder
	ArrayEncoder        = zapcore.ArrayEncoder
	ObjectMarshaler     = zapcore.ObjectMarshaler
	ArrayMarshaler      = zapcore.ArrayMarshaler
	ObjectMarshalerFunc = zapcore.ObjectMarshalerFunc
	ArrayMarshalerFunc  = zapcore.ArrayMarshalerFunc
)

// While the interfaces are aliases, converting between the flavours is the
// identity.

func fromZapObject(m zapcore.ObjectMarshaler) ObjectMarshaler    { return m }
func fromZapArray(m zapcore.ArrayMarshaler) ArrayMarshaler       { return m }
func toZapObject(m ObjectMarshaler) zapcore.ObjectMarshaler      { return m }
func toZapArray(m ArrayMarshaler) zapcore.ArrayMarshaler         { return m }
func toZapObjectEncoder(e ObjectEncoder) zapcore.ObjectEncoder   { return e }
func fromZapObjectEncoder(e zapcore.ObjectEncoder) ObjectEncoder { return e }
func zapValue(v any) any                                         { return v }
func zapValueChanged(v any) (any, bool)                          { return v, false }
func zapValues(kvs []any) []any                                  { return kvs }
