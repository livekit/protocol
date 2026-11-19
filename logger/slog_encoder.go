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

// The encoder below is ported from go.uber.org/zap/zapcore's jsonEncoder and
// consoleEncoder so that the slog backend writes byte-identical output:
//
// Copyright (c) 2016-2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package logger

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// encFlavor selects between the two formats the zap backend produces.
type encFlavor uint8

const (
	// flavorJSON matches zap.NewProductionEncoderConfig with the JSON encoder.
	flavorJSON encFlavor = iota
	// flavorConsole matches zap.NewDevelopmentEncoderConfig with the console
	// encoder, whose structured context is spaced JSON.
	flavorConsole
)

const iso8601Layout = "2006-01-02T15:04:05.000Z0700"

// jsonEnc implements ObjectEncoder and ArrayEncoder over a
// byte slice, so existing MarshalLogObject/MarshalLogArray implementations
// stream into it unchanged.
type jsonEnc struct {
	flavor     encFlavor
	buf        []byte
	openNS     int
	reflectBuf bytes.Buffer
	reflectEnc *json.Encoder
}

var _ ObjectEncoder = (*jsonEnc)(nil)
var _ ArrayEncoder = (*jsonEnc)(nil)

var jsonEncPool = sync.Pool{New: func() any { return &jsonEnc{buf: make([]byte, 0, 1024)} }}

func getJSONEnc(f encFlavor) *jsonEnc {
	e := jsonEncPool.Get().(*jsonEnc)
	e.flavor, e.buf, e.openNS = f, e.buf[:0], 0
	return e
}

func putJSONEnc(e *jsonEnc) {
	if cap(e.buf) <= 64<<10 {
		jsonEncPool.Put(e)
	}
}

func (e *jsonEnc) spaced() bool { return e.flavor == flavorConsole }

func (e *jsonEnc) addElementSeparator() {
	if n := len(e.buf); n > 0 {
		switch e.buf[n-1] {
		case '{', '[', ':', ',', ' ':
		default:
			e.buf = append(e.buf, ',')
			if e.spaced() {
				e.buf = append(e.buf, ' ')
			}
		}
	}
}

func (e *jsonEnc) addKey(k string) {
	e.addElementSeparator()
	e.buf = append(e.buf, '"')
	e.buf = appendJSONSafe(e.buf, k)
	e.buf = append(e.buf, '"', ':')
	if e.spaced() {
		e.buf = append(e.buf, ' ')
	}
}

func (e *jsonEnc) closeOpenNamespaces() {
	for ; e.openNS > 0; e.openNS-- {
		e.buf = append(e.buf, '}')
	}
}

func (e *jsonEnc) AddArray(k string, m ArrayMarshaler) error {
	e.addKey(k)
	return e.AppendArray(m)
}

func (e *jsonEnc) AddObject(k string, m ObjectMarshaler) error {
	e.addKey(k)
	return e.AppendObject(m)
}

func (e *jsonEnc) AddBinary(k string, v []byte) {
	e.AddString(k, base64.StdEncoding.EncodeToString(v))
}

func (e *jsonEnc) AddByteString(k string, v []byte)     { e.addKey(k); e.AppendByteString(v) }
func (e *jsonEnc) AddBool(k string, v bool)             { e.addKey(k); e.AppendBool(v) }
func (e *jsonEnc) AddComplex128(k string, v complex128) { e.addKey(k); e.AppendComplex128(v) }
func (e *jsonEnc) AddComplex64(k string, v complex64)   { e.addKey(k); e.AppendComplex64(v) }
func (e *jsonEnc) AddDuration(k string, v time.Duration) {
	e.addKey(k)
	e.AppendDuration(v)
}
func (e *jsonEnc) AddFloat64(k string, v float64) { e.addKey(k); e.AppendFloat64(v) }
func (e *jsonEnc) AddFloat32(k string, v float32) { e.addKey(k); e.AppendFloat32(v) }
func (e *jsonEnc) AddInt(k string, v int)         { e.AddInt64(k, int64(v)) }
func (e *jsonEnc) AddInt64(k string, v int64)     { e.addKey(k); e.AppendInt64(v) }
func (e *jsonEnc) AddInt32(k string, v int32)     { e.AddInt64(k, int64(v)) }
func (e *jsonEnc) AddInt16(k string, v int16)     { e.AddInt64(k, int64(v)) }
func (e *jsonEnc) AddInt8(k string, v int8)       { e.AddInt64(k, int64(v)) }
func (e *jsonEnc) AddString(k, v string)          { e.addKey(k); e.AppendString(v) }
func (e *jsonEnc) AddTime(k string, v time.Time)  { e.addKey(k); e.AppendTime(v) }
func (e *jsonEnc) AddUint(k string, v uint)       { e.AddUint64(k, uint64(v)) }
func (e *jsonEnc) AddUint64(k string, v uint64)   { e.addKey(k); e.AppendUint64(v) }
func (e *jsonEnc) AddUint32(k string, v uint32)   { e.AddUint64(k, uint64(v)) }
func (e *jsonEnc) AddUint16(k string, v uint16)   { e.AddUint64(k, uint64(v)) }
func (e *jsonEnc) AddUint8(k string, v uint8)     { e.AddUint64(k, uint64(v)) }
func (e *jsonEnc) AddUintptr(k string, v uintptr) { e.AddUint64(k, uint64(v)) }
func (e *jsonEnc) AppendInt(v int)                { e.AppendInt64(int64(v)) }
func (e *jsonEnc) AppendInt32(v int32)            { e.AppendInt64(int64(v)) }
func (e *jsonEnc) AppendInt16(v int16)            { e.AppendInt64(int64(v)) }
func (e *jsonEnc) AppendInt8(v int8)              { e.AppendInt64(int64(v)) }
func (e *jsonEnc) AppendUint(v uint)              { e.AppendUint64(uint64(v)) }
func (e *jsonEnc) AppendUint32(v uint32)          { e.AppendUint64(uint64(v)) }
func (e *jsonEnc) AppendUint16(v uint16)          { e.AppendUint64(uint64(v)) }
func (e *jsonEnc) AppendUint8(v uint8)            { e.AppendUint64(uint64(v)) }
func (e *jsonEnc) AppendUintptr(v uintptr)        { e.AppendUint64(uint64(v)) }
func (e *jsonEnc) AppendComplex128(v complex128)  { e.appendComplex(v, 64) }
func (e *jsonEnc) AppendComplex64(v complex64)    { e.appendComplex(complex128(v), 32) }
func (e *jsonEnc) AppendFloat64(v float64)        { e.appendFloat(v, 64) }
func (e *jsonEnc) AppendFloat32(v float32)        { e.appendFloat(float64(v), 32) }
func (e *jsonEnc) AppendBool(v bool)              { e.addElementSeparator(); e.buf = strconv.AppendBool(e.buf, v) }
func (e *jsonEnc) AppendInt64(v int64) {
	e.addElementSeparator()
	e.buf = strconv.AppendInt(e.buf, v, 10)
}
func (e *jsonEnc) AppendUint64(v uint64) {
	e.addElementSeparator()
	e.buf = strconv.AppendUint(e.buf, v, 10)
}
func (e *jsonEnc) OpenNamespace(k string) { e.addKey(k); e.buf = append(e.buf, '{'); e.openNS++ }
func (e *jsonEnc) AddReflected(k string, v any) error {
	b, err := e.encodeReflected(v)
	if err != nil {
		return err
	}
	e.addKey(k)
	e.buf = append(e.buf, b...)
	return nil
}

func (e *jsonEnc) encodeReflected(v any) ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}
	if e.reflectEnc == nil {
		e.reflectEnc = json.NewEncoder(&e.reflectBuf)
		e.reflectEnc.SetEscapeHTML(false)
	}
	e.reflectBuf.Reset()
	if err := e.reflectEnc.Encode(v); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(e.reflectBuf.Bytes(), []byte("\n")), nil
}

func (e *jsonEnc) AppendArray(m ArrayMarshaler) error {
	e.addElementSeparator()
	e.buf = append(e.buf, '[')
	err := m.MarshalLogArray(e)
	e.buf = append(e.buf, ']')
	return err
}

func (e *jsonEnc) AppendObject(m ObjectMarshaler) error {
	// Close only the namespaces opened by this object, as zap does.
	old := e.openNS
	e.openNS = 0
	e.addElementSeparator()
	e.buf = append(e.buf, '{')
	err := m.MarshalLogObject(e)
	e.buf = append(e.buf, '}')
	e.closeOpenNamespaces()
	e.openNS = old
	return err
}

func (e *jsonEnc) AppendByteString(v []byte) {
	e.addElementSeparator()
	e.buf = append(e.buf, '"')
	e.buf = appendJSONSafe(e.buf, v)
	e.buf = append(e.buf, '"')
}

func (e *jsonEnc) appendComplex(v complex128, prec int) {
	e.addElementSeparator()
	r, i := real(v), imag(v)
	e.buf = append(e.buf, '"')
	e.buf = strconv.AppendFloat(e.buf, r, 'f', -1, prec)
	if i >= 0 {
		e.buf = append(e.buf, '+')
	}
	e.buf = strconv.AppendFloat(e.buf, i, 'f', -1, prec)
	e.buf = append(e.buf, 'i', '"')
}

func (e *jsonEnc) AppendDuration(v time.Duration) {
	if e.flavor == flavorJSON {
		e.AppendFloat64(float64(v) / float64(time.Second)) // SecondsDurationEncoder
	} else {
		e.AppendString(v.String()) // StringDurationEncoder
	}
}

func (e *jsonEnc) appendFloat(v float64, bits int) {
	e.addElementSeparator()
	switch {
	case math.IsNaN(v):
		e.buf = append(e.buf, `"NaN"`...)
	case math.IsInf(v, 1):
		e.buf = append(e.buf, `"+Inf"`...)
	case math.IsInf(v, -1):
		e.buf = append(e.buf, `"-Inf"`...)
	default:
		e.buf = strconv.AppendFloat(e.buf, v, 'f', -1, bits)
	}
}

func (e *jsonEnc) AppendString(v string) {
	e.addElementSeparator()
	e.buf = append(e.buf, '"')
	e.buf = appendJSONSafe(e.buf, v)
	e.buf = append(e.buf, '"')
}

func (e *jsonEnc) AppendTimeLayout(t time.Time, layout string) {
	e.addElementSeparator()
	e.buf = append(e.buf, '"')
	e.buf = t.AppendFormat(e.buf, layout)
	e.buf = append(e.buf, '"')
}

func (e *jsonEnc) AppendTime(t time.Time) {
	if e.flavor == flavorJSON {
		e.AppendFloat64(float64(t.UnixNano()) / float64(time.Second)) // EpochTimeEncoder
	} else {
		e.AppendTimeLayout(t, iso8601Layout) // ISO8601TimeEncoder
	}
}

func (e *jsonEnc) AppendReflected(v any) error {
	b, err := e.encodeReflected(v)
	if err != nil {
		return err
	}
	e.addElementSeparator()
	e.buf = append(e.buf, b...)
	return nil
}

const lowerHex = "0123456789abcdef"

// appendJSONSafe is zapcore's safeAppendStringLike.
func appendJSONSafe[S []byte | string](buf []byte, s S) []byte {
	last := 0
	for i := 0; i < len(s); {
		c := s[i]
		if c >= utf8.RuneSelf {
			var r rune
			var size int
			switch x := any(s).(type) {
			case string:
				r, size = utf8.DecodeRuneInString(x[i:])
			case []byte:
				r, size = utf8.DecodeRune(x[i:])
			}
			if r != utf8.RuneError || size != 1 {
				i += size
				continue
			}
			buf = append(buf, s[last:i]...)
			buf = append(buf, "\\ufffd"...)
			i++
			last = i
			continue
		}
		if c >= 0x20 && c != '\\' && c != '"' {
			i++
			continue
		}
		buf = append(buf, s[last:i]...)
		switch c {
		case '\\', '"':
			buf = append(buf, '\\', c)
		case '\n':
			buf = append(buf, '\\', 'n')
		case '\r':
			buf = append(buf, '\\', 'r')
		case '\t':
			buf = append(buf, '\\', 't')
		default:
			buf = append(buf, '\\', 'u', '0', '0', lowerHex[c>>4], lowerHex[c&0xF])
		}
		i++
		last = i
	}
	return append(buf, s[last:]...)
}

// ---------------------------------------------------------------- levels

// The zap backend's levels, as slog levels (zap level × 4).
const (
	levelDPanic = slog.Level(12)
	levelPanic  = slog.Level(16)
	levelFatal  = slog.Level(20)
)

var (
	levelLower   = [...]string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"}
	levelCapital = [...]string{"DEBUG", "INFO", "WARN", "ERROR", "DPANIC", "PANIC", "FATAL"}
)

// normLevel maps an arbitrary slog level onto the nearest zap level at or
// below it, as zapslog does.
func normLevel(l slog.Level) slog.Level {
	switch {
	case l >= slog.LevelDebug && l <= levelFatal && l%4 == 0:
		return l
	case l >= slog.LevelError:
		return slog.LevelError
	case l >= slog.LevelWarn:
		return slog.LevelWarn
	case l >= slog.LevelInfo:
		return slog.LevelInfo
	}
	return slog.LevelDebug
}

func levelIndex(l slog.Level) int { return int(normLevel(l)-slog.LevelDebug) / 4 }

// ---------------------------------------------------------------- entries

type callerInfo struct {
	defined  bool
	pc       uintptr
	file     string
	line     int
	function string
}

// appendTrimmedPath is zapcore.EntryCaller.TrimmedPath without allocating.
func appendTrimmedPath(b []byte, c callerInfo) []byte {
	file := c.file
	if i := strings.LastIndexByte(file, '/'); i >= 0 {
		if j := strings.LastIndexByte(file[:i], '/'); j >= 0 {
			file = file[j+1:]
		}
	}
	b = append(b, file...)
	b = append(b, ':')
	return strconv.AppendInt(b, int64(c.line), 10)
}

type logEntry struct {
	level  slog.Level // normalised
	time   time.Time
	name   string
	caller callerInfo
	msg    string
	stack  string
}

// encodeEntry writes one line. ctx is the pre-encoded context (with ctxNS
// namespaces open) and fields adds the entry's own fields.
func encodeEntry(line *jsonEnc, ent *logEntry, ctx []byte, ctxNS int, fields func(e *jsonEnc)) {
	idx := levelIndex(ent.level)
	if line.flavor == flavorJSON {
		line.buf = append(line.buf, '{')
		line.AddString("level", levelLower[idx])
		line.AddTime("ts", ent.time)
		if ent.name != "" {
			line.AddString("logger", ent.name)
		}
		if ent.caller.defined {
			line.addKey("caller")
			line.buf = append(line.buf, '"')
			line.buf = appendTrimmedPath(line.buf, ent.caller)
			line.buf = append(line.buf, '"')
		}
		line.AddString("msg", ent.msg)
		if len(ctx) > 0 {
			line.addElementSeparator()
			line.buf = append(line.buf, ctx...)
		}
		line.openNS = ctxNS
		fields(line)
		line.closeOpenNamespaces()
		if ent.stack != "" {
			line.AddString("stacktrace", ent.stack)
		}
		line.buf = append(line.buf, '}', '\n')
		return
	}

	b := ent.time.AppendFormat(line.buf, iso8601Layout)
	b = append(append(b, '\t'), levelCapital[idx]...)
	if ent.name != "" {
		b = append(append(b, '\t'), ent.name...)
	}
	if ent.caller.defined {
		b = appendTrimmedPath(append(b, '\t'), ent.caller)
	}
	b = append(append(b, '\t'), ent.msg...)
	c := getJSONEnc(flavorConsole)
	c.buf = append(c.buf, ctx...)
	c.openNS = ctxNS
	fields(c)
	c.closeOpenNamespaces()
	if len(c.buf) > 0 {
		b = append(b, '\t', '{')
		b = append(b, c.buf...)
		b = append(b, '}')
	}
	putJSONEnc(c)
	if ent.stack != "" {
		b = append(append(b, '\n'), ent.stack...)
	}
	line.buf = append(b, '\n')
}
