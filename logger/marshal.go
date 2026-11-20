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

package logger

// Custom types log themselves as structured values by implementing
// ObjectMarshaler or ArrayMarshaler:
//
//	func (t *T) MarshalLogObject(e logger.ObjectEncoder) error
//
// These interfaces are aliases of zap's today (marshal_alias.go). Building
// with -tags logger_own_marshalers replaces them with this package's own,
// identical interfaces (marshal_own.go); that becomes the default once no
// implementer names the zapcore types. Code that receives a marshaler
// therefore accepts both flavours, logger's first: while the aliases are in
// place the two checks see the same interface and the first one wins.

import "go.uber.org/zap/zapcore"

// asObjectMarshaler returns v as an ObjectMarshaler of either flavour.
func asObjectMarshaler(v any) (ObjectMarshaler, bool) {
	if m, ok := v.(ObjectMarshaler); ok {
		return m, true
	}
	if m, ok := v.(zapcore.ObjectMarshaler); ok {
		return fromZapObject(m), true
	}
	return nil, false
}

// asArrayMarshaler returns v as an ArrayMarshaler of either flavour.
func asArrayMarshaler(v any) (ArrayMarshaler, bool) {
	if m, ok := v.(ArrayMarshaler); ok {
		return m, true
	}
	if m, ok := v.(zapcore.ArrayMarshaler); ok {
		return fromZapArray(m), true
	}
	return nil, false
}

func isMarshaler(v any) bool {
	if _, ok := asObjectMarshaler(v); ok {
		return true
	}
	_, ok := asArrayMarshaler(v)
	return ok
}
