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

//go:build logger_own_marshalers

package logger

import (
	"slices"
	"time"

	"go.uber.org/zap/zapcore"
)

// See marshal.go. These mirror zapcore's interfaces method for method, with
// this package's types in the signatures.

type ObjectMarshaler interface {
	MarshalLogObject(ObjectEncoder) error
}

type ObjectMarshalerFunc func(ObjectEncoder) error

func (f ObjectMarshalerFunc) MarshalLogObject(enc ObjectEncoder) error { return f(enc) }

type ArrayMarshaler interface {
	MarshalLogArray(ArrayEncoder) error
}

type ArrayMarshalerFunc func(ArrayEncoder) error

func (f ArrayMarshalerFunc) MarshalLogArray(enc ArrayEncoder) error { return f(enc) }

type ObjectEncoder interface {
	AddArray(key string, marshaler ArrayMarshaler) error
	AddObject(key string, marshaler ObjectMarshaler) error

	AddBinary(key string, value []byte)
	AddByteString(key string, value []byte)
	AddBool(key string, value bool)
	AddComplex128(key string, value complex128)
	AddComplex64(key string, value complex64)
	AddDuration(key string, value time.Duration)
	AddFloat64(key string, value float64)
	AddFloat32(key string, value float32)
	AddInt(key string, value int)
	AddInt64(key string, value int64)
	AddInt32(key string, value int32)
	AddInt16(key string, value int16)
	AddInt8(key string, value int8)
	AddString(key, value string)
	AddTime(key string, value time.Time)
	AddUint(key string, value uint)
	AddUint64(key string, value uint64)
	AddUint32(key string, value uint32)
	AddUint16(key string, value uint16)
	AddUint8(key string, value uint8)
	AddUintptr(key string, value uintptr)

	AddReflected(key string, value any) error
	OpenNamespace(key string)
}

type PrimitiveArrayEncoder interface {
	AppendBool(bool)
	AppendByteString([]byte)
	AppendComplex128(complex128)
	AppendComplex64(complex64)
	AppendFloat64(float64)
	AppendFloat32(float32)
	AppendInt(int)
	AppendInt64(int64)
	AppendInt32(int32)
	AppendInt16(int16)
	AppendInt8(int8)
	AppendString(string)
	AppendUint(uint)
	AppendUint64(uint64)
	AppendUint32(uint32)
	AppendUint16(uint16)
	AppendUint8(uint8)
	AppendUintptr(uintptr)
}

type ArrayEncoder interface {
	PrimitiveArrayEncoder

	AppendDuration(time.Duration)
	AppendTime(time.Time)

	AppendArray(ArrayMarshaler) error
	AppendObject(ObjectMarshaler) error

	AppendReflected(value any) error
}

// ---------------------------------------------------------------- marshalers

// zapObject runs a zapcore marshaler against a logger encoder.
type zapObject struct{ m zapcore.ObjectMarshaler }

func (z zapObject) MarshalLogObject(e ObjectEncoder) error {
	return z.m.MarshalLogObject(toZapObjectEncoder(e))
}

type zapArray struct{ m zapcore.ArrayMarshaler }

func (z zapArray) MarshalLogArray(e ArrayEncoder) error {
	return z.m.MarshalLogArray(toZapArrayEncoder(e))
}

// logObject runs a logger marshaler against a zapcore encoder.
type logObject struct{ m ObjectMarshaler }

func (l logObject) MarshalLogObject(e zapcore.ObjectEncoder) error {
	return l.m.MarshalLogObject(fromZapObjectEncoder(e))
}

type logArray struct{ m ArrayMarshaler }

func (l logArray) MarshalLogArray(e zapcore.ArrayEncoder) error {
	return l.m.MarshalLogArray(fromZapArrayEncoder(e))
}

func fromZapObject(m zapcore.ObjectMarshaler) ObjectMarshaler {
	if l, ok := m.(logObject); ok {
		return l.m
	}
	return zapObject{m}
}

func fromZapArray(m zapcore.ArrayMarshaler) ArrayMarshaler {
	if l, ok := m.(logArray); ok {
		return l.m
	}
	return zapArray{m}
}

func toZapObject(m ObjectMarshaler) zapcore.ObjectMarshaler {
	if z, ok := m.(zapObject); ok {
		return z.m
	}
	return logObject{m}
}

func toZapArray(m ArrayMarshaler) zapcore.ArrayMarshaler {
	if z, ok := m.(zapArray); ok {
		return z.m
	}
	return logArray{m}
}

// zapValueChanged makes v loggable by zap: a type that implements only this
// package's marshalers is wrapped in the zapcore flavour.
func zapValueChanged(v any) (any, bool) {
	switch v.(type) {
	case zapcore.ObjectMarshaler, zapcore.ArrayMarshaler:
		return v, false
	}
	if m, ok := v.(ObjectMarshaler); ok {
		return logObject{m}, true
	}
	if m, ok := v.(ArrayMarshaler); ok {
		return logArray{m}, true
	}
	return v, false
}

func zapValue(v any) any {
	z, _ := zapValueChanged(v)
	return z
}

// zapValues applies zapValue to sugared key/values, copying only if needed.
func zapValues(kvs []any) []any {
	var out []any
	for i, v := range kvs {
		if z, ok := zapValueChanged(v); ok {
			if out == nil {
				out = slices.Clone(kvs)
			}
			out[i] = z
		}
	}
	if out == nil {
		return kvs
	}
	return out
}

// ---------------------------------------------------------------- encoders

type zapSensitiveObjectEncoder interface {
	SensitiveObjectEncoder(Sensitivity) zapcore.ObjectEncoder
}

type zapSensitiveArrayEncoder interface {
	SensitiveArrayEncoder(Sensitivity) zapcore.ArrayEncoder
}

func toZapObjectEncoder(e ObjectEncoder) zapcore.ObjectEncoder {
	if l, ok := e.(*logObjEnc); ok {
		return l.e
	}
	return &zapObjEnc{e}
}

func fromZapObjectEncoder(e zapcore.ObjectEncoder) ObjectEncoder {
	if z, ok := e.(*zapObjEnc); ok {
		return z.e
	}
	return &logObjEnc{e}
}

func toZapArrayEncoder(e ArrayEncoder) zapcore.ArrayEncoder {
	if l, ok := e.(*logArrEnc); ok {
		return l.e
	}
	return &zapArrEnc{e}
}

func fromZapArrayEncoder(e zapcore.ArrayEncoder) ArrayEncoder {
	if z, ok := e.(*zapArrEnc); ok {
		return z.e
	}
	return &logArrEnc{e}
}

// zapObjEnc presents a logger ObjectEncoder as a zapcore one.
type zapObjEnc struct{ e ObjectEncoder }

func (z *zapObjEnc) AddArray(k string, m zapcore.ArrayMarshaler) error {
	return z.e.AddArray(k, fromZapArray(m))
}
func (z *zapObjEnc) AddObject(k string, m zapcore.ObjectMarshaler) error {
	return z.e.AddObject(k, fromZapObject(m))
}
func (z *zapObjEnc) AddBinary(k string, v []byte)          { z.e.AddBinary(k, v) }
func (z *zapObjEnc) AddByteString(k string, v []byte)      { z.e.AddByteString(k, v) }
func (z *zapObjEnc) AddBool(k string, v bool)              { z.e.AddBool(k, v) }
func (z *zapObjEnc) AddComplex128(k string, v complex128)  { z.e.AddComplex128(k, v) }
func (z *zapObjEnc) AddComplex64(k string, v complex64)    { z.e.AddComplex64(k, v) }
func (z *zapObjEnc) AddDuration(k string, v time.Duration) { z.e.AddDuration(k, v) }
func (z *zapObjEnc) AddFloat64(k string, v float64)        { z.e.AddFloat64(k, v) }
func (z *zapObjEnc) AddFloat32(k string, v float32)        { z.e.AddFloat32(k, v) }
func (z *zapObjEnc) AddInt(k string, v int)                { z.e.AddInt(k, v) }
func (z *zapObjEnc) AddInt64(k string, v int64)            { z.e.AddInt64(k, v) }
func (z *zapObjEnc) AddInt32(k string, v int32)            { z.e.AddInt32(k, v) }
func (z *zapObjEnc) AddInt16(k string, v int16)            { z.e.AddInt16(k, v) }
func (z *zapObjEnc) AddInt8(k string, v int8)              { z.e.AddInt8(k, v) }
func (z *zapObjEnc) AddString(k, v string)                 { z.e.AddString(k, v) }
func (z *zapObjEnc) AddTime(k string, v time.Time)         { z.e.AddTime(k, v) }
func (z *zapObjEnc) AddUint(k string, v uint)              { z.e.AddUint(k, v) }
func (z *zapObjEnc) AddUint64(k string, v uint64)          { z.e.AddUint64(k, v) }
func (z *zapObjEnc) AddUint32(k string, v uint32)          { z.e.AddUint32(k, v) }
func (z *zapObjEnc) AddUint16(k string, v uint16)          { z.e.AddUint16(k, v) }
func (z *zapObjEnc) AddUint8(k string, v uint8)            { z.e.AddUint8(k, v) }
func (z *zapObjEnc) AddUintptr(k string, v uintptr)        { z.e.AddUintptr(k, v) }
func (z *zapObjEnc) AddReflected(k string, v any) error    { return z.e.AddReflected(k, v) }
func (z *zapObjEnc) OpenNamespace(k string)                { z.e.OpenNamespace(k) }

func (z *zapObjEnc) SensitiveObjectEncoder(s Sensitivity) zapcore.ObjectEncoder {
	if e := ObjectEncoderFor(z.e, s); e != nil {
		return toZapObjectEncoder(e)
	}
	return nil
}

// logObjEnc presents a zapcore ObjectEncoder as a logger one, forwarding the
// sensitivity capability when the zapcore encoder has it.
type logObjEnc struct{ e zapcore.ObjectEncoder }

func (l *logObjEnc) AddArray(k string, m ArrayMarshaler) error {
	return l.e.AddArray(k, toZapArray(m))
}
func (l *logObjEnc) AddObject(k string, m ObjectMarshaler) error {
	return l.e.AddObject(k, toZapObject(m))
}
func (l *logObjEnc) AddBinary(k string, v []byte)          { l.e.AddBinary(k, v) }
func (l *logObjEnc) AddByteString(k string, v []byte)      { l.e.AddByteString(k, v) }
func (l *logObjEnc) AddBool(k string, v bool)              { l.e.AddBool(k, v) }
func (l *logObjEnc) AddComplex128(k string, v complex128)  { l.e.AddComplex128(k, v) }
func (l *logObjEnc) AddComplex64(k string, v complex64)    { l.e.AddComplex64(k, v) }
func (l *logObjEnc) AddDuration(k string, v time.Duration) { l.e.AddDuration(k, v) }
func (l *logObjEnc) AddFloat64(k string, v float64)        { l.e.AddFloat64(k, v) }
func (l *logObjEnc) AddFloat32(k string, v float32)        { l.e.AddFloat32(k, v) }
func (l *logObjEnc) AddInt(k string, v int)                { l.e.AddInt(k, v) }
func (l *logObjEnc) AddInt64(k string, v int64)            { l.e.AddInt64(k, v) }
func (l *logObjEnc) AddInt32(k string, v int32)            { l.e.AddInt32(k, v) }
func (l *logObjEnc) AddInt16(k string, v int16)            { l.e.AddInt16(k, v) }
func (l *logObjEnc) AddInt8(k string, v int8)              { l.e.AddInt8(k, v) }
func (l *logObjEnc) AddString(k, v string)                 { l.e.AddString(k, v) }
func (l *logObjEnc) AddTime(k string, v time.Time)         { l.e.AddTime(k, v) }
func (l *logObjEnc) AddUint(k string, v uint)              { l.e.AddUint(k, v) }
func (l *logObjEnc) AddUint64(k string, v uint64)          { l.e.AddUint64(k, v) }
func (l *logObjEnc) AddUint32(k string, v uint32)          { l.e.AddUint32(k, v) }
func (l *logObjEnc) AddUint16(k string, v uint16)          { l.e.AddUint16(k, v) }
func (l *logObjEnc) AddUint8(k string, v uint8)            { l.e.AddUint8(k, v) }
func (l *logObjEnc) AddUintptr(k string, v uintptr)        { l.e.AddUintptr(k, v) }
func (l *logObjEnc) AddReflected(k string, v any) error    { return l.e.AddReflected(k, v) }
func (l *logObjEnc) OpenNamespace(k string)                { l.e.OpenNamespace(k) }

func (l *logObjEnc) SensitiveObjectEncoder(s Sensitivity) ObjectEncoder {
	if z, ok := l.e.(zapSensitiveObjectEncoder); ok {
		if e := z.SensitiveObjectEncoder(s); e != nil {
			return fromZapObjectEncoder(e)
		}
	}
	return nil
}

// zapArrEnc presents a logger ArrayEncoder as a zapcore one.
type zapArrEnc struct{ e ArrayEncoder }

func (z *zapArrEnc) AppendArray(m zapcore.ArrayMarshaler) error {
	return z.e.AppendArray(fromZapArray(m))
}
func (z *zapArrEnc) AppendObject(m zapcore.ObjectMarshaler) error {
	return z.e.AppendObject(fromZapObject(m))
}
func (z *zapArrEnc) AppendBool(v bool)              { z.e.AppendBool(v) }
func (z *zapArrEnc) AppendByteString(v []byte)      { z.e.AppendByteString(v) }
func (z *zapArrEnc) AppendComplex128(v complex128)  { z.e.AppendComplex128(v) }
func (z *zapArrEnc) AppendComplex64(v complex64)    { z.e.AppendComplex64(v) }
func (z *zapArrEnc) AppendFloat64(v float64)        { z.e.AppendFloat64(v) }
func (z *zapArrEnc) AppendFloat32(v float32)        { z.e.AppendFloat32(v) }
func (z *zapArrEnc) AppendInt(v int)                { z.e.AppendInt(v) }
func (z *zapArrEnc) AppendInt64(v int64)            { z.e.AppendInt64(v) }
func (z *zapArrEnc) AppendInt32(v int32)            { z.e.AppendInt32(v) }
func (z *zapArrEnc) AppendInt16(v int16)            { z.e.AppendInt16(v) }
func (z *zapArrEnc) AppendInt8(v int8)              { z.e.AppendInt8(v) }
func (z *zapArrEnc) AppendString(v string)          { z.e.AppendString(v) }
func (z *zapArrEnc) AppendUint(v uint)              { z.e.AppendUint(v) }
func (z *zapArrEnc) AppendUint64(v uint64)          { z.e.AppendUint64(v) }
func (z *zapArrEnc) AppendUint32(v uint32)          { z.e.AppendUint32(v) }
func (z *zapArrEnc) AppendUint16(v uint16)          { z.e.AppendUint16(v) }
func (z *zapArrEnc) AppendUint8(v uint8)            { z.e.AppendUint8(v) }
func (z *zapArrEnc) AppendUintptr(v uintptr)        { z.e.AppendUintptr(v) }
func (z *zapArrEnc) AppendDuration(v time.Duration) { z.e.AppendDuration(v) }
func (z *zapArrEnc) AppendTime(v time.Time)         { z.e.AppendTime(v) }
func (z *zapArrEnc) AppendReflected(v any) error    { return z.e.AppendReflected(v) }

func (z *zapArrEnc) SensitiveArrayEncoder(s Sensitivity) zapcore.ArrayEncoder {
	if e := ArrayEncoderFor(z.e, s); e != nil {
		return toZapArrayEncoder(e)
	}
	return nil
}

// logArrEnc presents a zapcore ArrayEncoder as a logger one.
type logArrEnc struct{ e zapcore.ArrayEncoder }

func (l *logArrEnc) AppendArray(m ArrayMarshaler) error   { return l.e.AppendArray(toZapArray(m)) }
func (l *logArrEnc) AppendObject(m ObjectMarshaler) error { return l.e.AppendObject(toZapObject(m)) }
func (l *logArrEnc) AppendBool(v bool)                    { l.e.AppendBool(v) }
func (l *logArrEnc) AppendByteString(v []byte)            { l.e.AppendByteString(v) }
func (l *logArrEnc) AppendComplex128(v complex128)        { l.e.AppendComplex128(v) }
func (l *logArrEnc) AppendComplex64(v complex64)          { l.e.AppendComplex64(v) }
func (l *logArrEnc) AppendFloat64(v float64)              { l.e.AppendFloat64(v) }
func (l *logArrEnc) AppendFloat32(v float32)              { l.e.AppendFloat32(v) }
func (l *logArrEnc) AppendInt(v int)                      { l.e.AppendInt(v) }
func (l *logArrEnc) AppendInt64(v int64)                  { l.e.AppendInt64(v) }
func (l *logArrEnc) AppendInt32(v int32)                  { l.e.AppendInt32(v) }
func (l *logArrEnc) AppendInt16(v int16)                  { l.e.AppendInt16(v) }
func (l *logArrEnc) AppendInt8(v int8)                    { l.e.AppendInt8(v) }
func (l *logArrEnc) AppendString(v string)                { l.e.AppendString(v) }
func (l *logArrEnc) AppendUint(v uint)                    { l.e.AppendUint(v) }
func (l *logArrEnc) AppendUint64(v uint64)                { l.e.AppendUint64(v) }
func (l *logArrEnc) AppendUint32(v uint32)                { l.e.AppendUint32(v) }
func (l *logArrEnc) AppendUint16(v uint16)                { l.e.AppendUint16(v) }
func (l *logArrEnc) AppendUint8(v uint8)                  { l.e.AppendUint8(v) }
func (l *logArrEnc) AppendUintptr(v uintptr)              { l.e.AppendUintptr(v) }
func (l *logArrEnc) AppendDuration(v time.Duration)       { l.e.AppendDuration(v) }
func (l *logArrEnc) AppendTime(v time.Time)               { l.e.AppendTime(v) }
func (l *logArrEnc) AppendReflected(v any) error          { return l.e.AppendReflected(v) }

func (l *logArrEnc) SensitiveArrayEncoder(s Sensitivity) ArrayEncoder {
	if z, ok := l.e.(zapSensitiveArrayEncoder); ok {
		if e := z.SensitiveArrayEncoder(s); e != nil {
			return fromZapArrayEncoder(e)
		}
	}
	return nil
}
