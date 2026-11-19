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

import (
	"fmt"
	"log/slog"
	"reflect"
	"time"

	"go.uber.org/zap/zapcore"
)

// addAny encodes v under key exactly as zap.Any(key, v).AddTo(enc) would,
// without depending on zap's field machinery. The zapcore marshaler and
// encoder interfaces remain the contract for custom types.
func addAny(enc ObjectEncoder, key string, v any) {
	var err error
	switch x := v.(type) {
	case ObjectMarshaler:
		err = enc.AddObject(key, x)
	case ArrayMarshaler:
		err = enc.AddArray(key, x)
	case []zapcore.Field:
		err = enc.AddObject(key, fieldsObject(x))
	case bool:
		enc.AddBool(key, x)
	case *bool:
		err = addPtr(enc, key, x, enc.AddBool)
	case []bool:
		err = enc.AddArray(key, BoolSlice(x))
	case complex128:
		enc.AddComplex128(key, x)
	case *complex128:
		err = addPtr(enc, key, x, enc.AddComplex128)
	case []complex128:
		err = enc.AddArray(key, Complex128Slice(x))
	case complex64:
		enc.AddComplex64(key, x)
	case *complex64:
		err = addPtr(enc, key, x, enc.AddComplex64)
	case []complex64:
		err = enc.AddArray(key, Complex64Slice(x))
	case float64:
		enc.AddFloat64(key, x)
	case *float64:
		err = addPtr(enc, key, x, enc.AddFloat64)
	case []float64:
		err = enc.AddArray(key, Float64Slice(x))
	case float32:
		enc.AddFloat32(key, x)
	case *float32:
		err = addPtr(enc, key, x, enc.AddFloat32)
	case []float32:
		err = enc.AddArray(key, Float32Slice(x))
	case int:
		enc.AddInt(key, x)
	case *int:
		err = addPtr(enc, key, x, enc.AddInt)
	case []int:
		err = enc.AddArray(key, IntSlice(x))
	case int64:
		enc.AddInt64(key, x)
	case *int64:
		err = addPtr(enc, key, x, enc.AddInt64)
	case []int64:
		err = enc.AddArray(key, Int64Slice(x))
	case int32:
		enc.AddInt32(key, x)
	case *int32:
		err = addPtr(enc, key, x, enc.AddInt32)
	case []int32:
		err = enc.AddArray(key, Int32Slice(x))
	case int16:
		enc.AddInt16(key, x)
	case *int16:
		err = addPtr(enc, key, x, enc.AddInt16)
	case []int16:
		err = enc.AddArray(key, Int16Slice(x))
	case int8:
		enc.AddInt8(key, x)
	case *int8:
		err = addPtr(enc, key, x, enc.AddInt8)
	case []int8:
		err = enc.AddArray(key, Int8Slice(x))
	case string:
		enc.AddString(key, x)
	case *string:
		err = addPtr(enc, key, x, func(k, v string) { enc.AddString(k, v) })
	case []string:
		err = enc.AddArray(key, StringSlice(x))
	case uint:
		enc.AddUint(key, x)
	case *uint:
		err = addPtr(enc, key, x, enc.AddUint)
	case []uint:
		err = enc.AddArray(key, UintSlice(x))
	case uint64:
		enc.AddUint64(key, x)
	case *uint64:
		err = addPtr(enc, key, x, enc.AddUint64)
	case []uint64:
		err = enc.AddArray(key, Uint64Slice(x))
	case uint32:
		enc.AddUint32(key, x)
	case *uint32:
		err = addPtr(enc, key, x, enc.AddUint32)
	case []uint32:
		err = enc.AddArray(key, Uint32Slice(x))
	case uint16:
		enc.AddUint16(key, x)
	case *uint16:
		err = addPtr(enc, key, x, enc.AddUint16)
	case []uint16:
		err = enc.AddArray(key, Uint16Slice(x))
	case uint8:
		enc.AddUint8(key, x)
	case *uint8:
		err = addPtr(enc, key, x, enc.AddUint8)
	case []byte:
		enc.AddBinary(key, x)
	case uintptr:
		enc.AddUintptr(key, x)
	case *uintptr:
		err = addPtr(enc, key, x, enc.AddUintptr)
	case []uintptr:
		err = enc.AddArray(key, UintptrSlice(x))
	case time.Time:
		enc.AddTime(key, x)
	case *time.Time:
		err = addPtr(enc, key, x, enc.AddTime)
	case []time.Time:
		err = enc.AddArray(key, TimeSlice(x))
	case time.Duration:
		enc.AddDuration(key, x)
	case *time.Duration:
		err = addPtr(enc, key, x, enc.AddDuration)
	case []time.Duration:
		err = enc.AddArray(key, DurationSlice(x))
	case error:
		err = encodeError(key, x, enc)
	case []error:
		err = enc.AddArray(key, errArray(x))
	case fmt.Stringer:
		err = encodeStringer(key, x, enc)
	default:
		err = enc.AddReflected(key, x)
	}
	if err != nil {
		enc.AddString(key+"Error", err.Error())
	}
}

// addPtr is zap's *T fields: a nil pointer is encoded as JSON null.
func addPtr[T any](enc ObjectEncoder, key string, p *T, add func(string, T)) error {
	if p == nil {
		return enc.AddReflected(key, nil)
	}
	add(key, *p)
	return nil
}

type fieldsObject []zapcore.Field

func (fs fieldsObject) MarshalLogObject(enc ObjectEncoder) error {
	for _, f := range fs {
		f.AddTo(enc)
	}
	return nil
}

// encodeError is zapcore's encodeError: "key", plus "keyCauses" for error
// groups or "keyVerbose" for errors with a richer %+v.
func encodeError(key string, err error, enc ObjectEncoder) (retErr error) {
	defer func() {
		if r := recover(); r != nil {
			if v := reflect.ValueOf(err); v.Kind() == reflect.Pointer && v.IsNil() {
				enc.AddString(key, "<nil>")
				return
			}
			retErr = fmt.Errorf("PANIC=%v", r)
		}
	}()

	basic := err.Error()
	enc.AddString(key, basic)

	switch e := err.(type) {
	case interface{ Errors() []error }:
		return enc.AddArray(key+"Causes", errArray(e.Errors()))
	case fmt.Formatter:
		if verbose := fmt.Sprintf("%+v", e); verbose != basic {
			enc.AddString(key+"Verbose", verbose)
		}
	}
	return nil
}

func encodeStringer(key string, s fmt.Stringer, enc ObjectEncoder) (retErr error) {
	defer func() {
		if r := recover(); r != nil {
			if v := reflect.ValueOf(s); v.Kind() == reflect.Pointer && v.IsNil() {
				enc.AddString(key, "<nil>")
				return
			}
			retErr = fmt.Errorf("PANIC=%v", r)
		}
	}()
	enc.AddString(key, s.String())
	return nil
}

type errArray []error

func (errs errArray) MarshalLogArray(arr ArrayEncoder) error {
	for _, err := range errs {
		if err != nil {
			_ = arr.AppendObject(errElem{err})
		}
	}
	return nil
}

type errElem struct{ err error }

func (e errElem) MarshalLogObject(enc ObjectEncoder) error {
	if err := encodeError("error", e.err, enc); err != nil {
		enc.AddString("errorError", err.Error())
	}
	return nil
}

// ---------------------------------------------------------------- slog attrs

// float32Value keeps float32 width, which slog.AnyValue would widen to
// float64 and so format differently.
type float32Value float32

// bareField is a zapcore.Field passed in key position (zap.Inline, zap.String,
// ...). A Field passed as a value is reflected, as zap does.
type bareField zapcore.Field

// asBareField wraps a zapcore.Field passed in key position.
func asBareField(v any) (slog.Attr, bool) {
	f, ok := v.(zapcore.Field)
	if !ok {
		return slog.Attr{}, false
	}
	return slog.Attr{Key: f.Key, Value: slog.AnyValue(bareField(f))}, true
}

// attrOf builds the attr for a sugared key/value pair.
func attrOf(k string, v any) slog.Attr {
	if x, ok := v.(float32); ok {
		return slog.Attr{Key: k, Value: slog.AnyValue(float32Value(x))}
	}
	return slog.Any(k, v)
}

// addAttr encodes an attr as the zap backend would have encoded the original
// key/value pair.
func addAttr(enc ObjectEncoder, a slog.Attr) {
	v := a.Value
	switch v.Kind() {
	case slog.KindString:
		enc.AddString(a.Key, v.String())
	case slog.KindInt64:
		enc.AddInt64(a.Key, v.Int64())
	case slog.KindUint64:
		enc.AddUint64(a.Key, v.Uint64())
	case slog.KindFloat64:
		enc.AddFloat64(a.Key, v.Float64())
	case slog.KindBool:
		enc.AddBool(a.Key, v.Bool())
	case slog.KindDuration:
		enc.AddDuration(a.Key, v.Duration())
	case slog.KindTime:
		enc.AddTime(a.Key, v.Time())
	case slog.KindGroup:
		if err := enc.AddObject(a.Key, attrGroup(v.Group())); err != nil {
			enc.AddString(a.Key+"Error", err.Error())
		}
	case slog.KindLogValuer:
		if om, ok := v.LogValuer().(ObjectMarshaler); ok {
			addAny(enc, a.Key, om)
			return
		}
		addAttr(enc, slog.Attr{Key: a.Key, Value: v.Resolve()})
	default:
		switch x := v.Any().(type) {
		case bareField:
			zapcore.Field(x).AddTo(enc)
		case float32Value:
			enc.AddFloat32(a.Key, float32(x))
		default:
			addAny(enc, a.Key, x)
		}
	}
}

type attrGroup []slog.Attr

func (g attrGroup) MarshalLogObject(enc ObjectEncoder) error {
	for _, a := range g {
		addAttr(enc, a)
	}
	return nil
}
