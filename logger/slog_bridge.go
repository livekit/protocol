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
	"context"
	"log/slog"
	"runtime"
	"strconv"
	"time"

	"go.uber.org/zap/zapcore"
)

// ---------------------------------------------------------------- Logger <-> *slog.Logger

// ToSlog returns a *slog.Logger that writes through l. FromSlog converts it
// back to l, so code can pass either type across API boundaries.
func ToSlog(l Logger) *slog.Logger {
	if l == nil {
		return slog.New(NewSlogDiscard())
	}
	return slog.New(ToSlogHandler(l))
}

// FromSlog returns a Logger that writes through l.
//
// A *slog.Logger obtained from ToSlog (or slog.Default after InitFromConfig)
// converts back to the Logger it came from, with every feature intact. Any
// other handler is adapted: names become a "logger" attribute, values are
// rendered as plain slog values (zap fields, marshalers and errors as the zap
// backend would encode them), deferred values work, and levels and sampling
// are left to the handler.
func FromSlog(l *slog.Logger) Logger {
	if l == nil {
		return GetDiscardLogger()
	}
	h := l.Handler()
	if lh, ok := h.(*loggerHandler); ok && lh.log != nil {
		return lh.log
	}
	return &slogAdapter{h: portableHandler{h}}
}

// WithName is Logger.WithName for a *slog.Logger.
func WithName(l *slog.Logger, name string) *slog.Logger {
	return ToSlog(FromSlog(l).WithName(name))
}

// WithComponent is Logger.WithComponent for a *slog.Logger.
func WithComponent(l *slog.Logger, component string) *slog.Logger {
	return ToSlog(FromSlog(l).WithComponent(component))
}

// WithDeferredValues is Logger.WithDeferredValues for a *slog.Logger.
func WithDeferredValues(l *slog.Logger) (*slog.Logger, DeferredFieldResolver) {
	d, r := FromSlog(l).WithDeferredValues()
	return ToSlog(d), r
}

// WithItemSampler is Logger.WithItemSampler for a *slog.Logger.
func WithItemSampler(l *slog.Logger) *slog.Logger {
	return ToSlog(FromSlog(l).WithItemSampler())
}

// WithoutSampler is Logger.WithoutSampler for a *slog.Logger.
func WithoutSampler(l *slog.Logger) *slog.Logger {
	return ToSlog(FromSlog(l).WithoutSampler())
}

// Err is the attribute Warnw and Errorw add for their error argument.
func Err(err error) slog.Attr {
	return slog.Any("error", err)
}

// Float32 keeps a float32's precision, which slog.Float64 and slog.Any widen
// (0.1 would be logged as 0.10000000149011612).
func Float32(key string, v float32) slog.Attr {
	return slog.Attr{Key: key, Value: slog.AnyValue(float32Value(v))}
}

// LogDepth logs through l like l.Log, reporting the caller depth frames above
// LogDepth's caller, as Logger.WithCallDepth does.
func LogDepth(ctx context.Context, l *slog.Logger, depth int, level slog.Level, msg string, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	if !l.Enabled(ctx, level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2+depth, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(args...)
	_ = l.Handler().Handle(ctx, r)
}

// loggerHandler is the slog.Handler view of a Logger. log is the equivalent
// Logger, kept up to date through WithAttrs, or nil once that is impossible.
type loggerHandler struct {
	log Logger
	h   slog.Handler
}

func newLoggerHandler(l *slogLogger) *loggerHandler {
	return &loggerHandler{log: l, h: l.handler()}
}

func (h *loggerHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return h.h.Enabled(ctx, l)
}

func (h *loggerHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.h.Handle(ctx, r)
}

func (h *loggerHandler) WithAttrs(as []slog.Attr) slog.Handler {
	if len(as) == 0 {
		return h
	}
	if l, ok := h.log.(*slogLogger); ok {
		return newLoggerHandler(l.withAttrs(as))
	}
	var log Logger
	if kvs, ok := attrsKVs(as); ok && h.log != nil {
		log = h.log.WithValues(kvs...)
	}
	return &loggerHandler{log: log, h: h.h.WithAttrs(as)}
}

func (h *loggerHandler) WithGroup(n string) slog.Handler {
	if n == "" {
		return h
	}
	if l, ok := h.log.(*slogLogger); ok {
		return newLoggerHandler(l.withGroup(n))
	}
	return &loggerHandler{h: h.h.WithGroup(n)}
}

// attrsKVs converts attrs to key/values for Logger.WithValues, failing for
// groups, which key/values cannot express.
func attrsKVs(as []slog.Attr) ([]any, bool) {
	kvs := make([]any, 0, 2*len(as))
	for _, a := range as {
		v := a.Value
		var x any
		switch v.Kind() {
		case slog.KindGroup:
			return nil, false
		case slog.KindLogValuer:
			if lv := unwrapLogValuer(v.LogValuer()); isMarshaler(lv) {
				x = lv
			} else if r := v.Resolve(); r.Kind() == slog.KindGroup {
				return nil, false
			} else {
				x = r.Any()
			}
		default:
			switch y := v.Any().(type) {
			case bareField:
				return nil, false
			case float32Value:
				x = float32(y)
			default:
				x = y
			}
		}
		kvs = append(kvs, a.Key, x)
	}
	return kvs, true
}

// ---------------------------------------------------------------- Logger over any slog.Handler

// slogAdapter implements Logger over a handler from outside this package.
type slogAdapter struct {
	h     slog.Handler
	name  string
	depth int
}

var _ Logger = (*slogAdapter)(nil)

func (l *slogAdapter) handler() slog.Handler {
	if l.name == "" {
		return l.h
	}
	return l.h.WithAttrs([]slog.Attr{slog.String("logger", l.name)})
}

func (l *slogAdapter) log(lvl slog.Level, msg string, err error, kvs []any) {
	ctx := context.Background()
	if !l.h.Enabled(ctx, lvl) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(3+l.depth, pcs[:])
	emit := func(lvl slog.Level, msg string, attrs ...slog.Attr) {
		r := slog.NewRecord(time.Now(), lvl, msg, pcs[0])
		if l.name != "" {
			r.AddAttrs(slog.String("logger", l.name))
		}
		r.AddAttrs(attrs...)
		_ = l.h.Handle(ctx, r)
	}
	if err != nil {
		kvs = append(kvs, "error", err)
	}
	attrs := sweeten(kvs, func(diag string, a slog.Attr) {
		if l.h.Enabled(ctx, slog.LevelError) {
			emit(slog.LevelError, diag, a)
		}
	})
	emit(lvl, msg, attrs...)
}

func (l *slogAdapter) Debugw(msg string, kvs ...any) { l.log(slog.LevelDebug, msg, nil, kvs) }
func (l *slogAdapter) Infow(msg string, kvs ...any)  { l.log(slog.LevelInfo, msg, nil, kvs) }
func (l *slogAdapter) Warnw(msg string, err error, kvs ...any) {
	l.log(slog.LevelWarn, msg, err, kvs)
}
func (l *slogAdapter) Errorw(msg string, err error, kvs ...any) {
	l.log(slog.LevelError, msg, err, kvs)
}

func (l *slogAdapter) dup() *slogAdapter {
	c := *l
	return &c
}

func (l *slogAdapter) WithValues(kvs ...any) Logger {
	c := l.dup()
	c.h = l.h.WithAttrs(valuesAttrs(kvs))
	return c
}

func (l *slogAdapter) WithUnlikelyValues(kvs ...any) UnlikelyLogger {
	return UnlikelyLogger{l, kvs}
}

func (l *slogAdapter) WithName(name string) Logger {
	c := l.dup()
	if c.name == "" {
		c.name = name
	} else {
		c.name += "." + name
	}
	return c
}

// WithComponent names the logger; levels are the handler's own.
func (l *slogAdapter) WithComponent(component string) Logger { return l.WithName(component) }

func (l *slogAdapter) WithCallDepth(depth int) Logger {
	c := l.dup()
	c.depth += depth
	return c
}

// WithItemSampler is a no-op: sampling is configured by Config, which an
// arbitrary handler does not have.
func (l *slogAdapter) WithItemSampler() Logger { return l }
func (l *slogAdapter) WithoutSampler() Logger  { return l }

func (l *slogAdapter) WithDeferredValues() (Logger, DeferredFieldResolver) {
	c := l.dup()
	d := &slogDeferrer{}
	c.h = &deferredHandler{l.h, d}
	return c, d
}

// valuesAttrs converts WithValues key/values: pairs with non-string keys, and
// a trailing key, are dropped, as the zap backend does.
func valuesAttrs(kvs []any) []slog.Attr {
	attrs := make([]slog.Attr, 0, len(kvs)/2)
	for i := 1; i < len(kvs); i += 2 {
		if k, ok := kvs[i-1].(string); ok {
			attrs = append(attrs, attrOf(k, kvs[i]))
		}
	}
	return attrs
}

// ---------------------------------------------------------------- portable values

// portableHandler renders this package's values (zap fields, marshalers,
// errors, ...) as plain slog values before they reach a handler from outside
// the package, using the same rules as the encoder. zap namespaces become
// groups.
type portableHandler struct{ h slog.Handler }

func (p portableHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return p.h.Enabled(ctx, l)
}

func (p portableHandler) Handle(ctx context.Context, r slog.Record) error {
	if !recordNeedsPortable(r) {
		return p.h.Handle(ctx, r)
	}
	attrs := make([]slog.Attr, 0, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})
	out := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	out.AddAttrs(portableAttrs(attrs)...)
	return p.h.Handle(ctx, out)
}

func (p portableHandler) WithAttrs(as []slog.Attr) slog.Handler {
	h := p.h
	start := 0
	for i, a := range as {
		if f, ok := a.Value.Any().(bareField); ok && f.Type == zapcore.NamespaceType {
			if i > start {
				h = h.WithAttrs(portableAttrs(as[start:i]))
			}
			h = h.WithGroup(f.Key)
			start = i + 1
		}
	}
	if start < len(as) {
		h = h.WithAttrs(portableAttrs(as[start:]))
	}
	return portableHandler{h}
}

func (p portableHandler) WithGroup(n string) slog.Handler {
	return portableHandler{p.h.WithGroup(n)}
}

func needsPortable(a slog.Attr) bool {
	switch a.Value.Kind() {
	case slog.KindAny, slog.KindLogValuer, slog.KindGroup:
		return true
	}
	return false
}

func recordNeedsPortable(r slog.Record) bool {
	needs := false
	r.Attrs(func(a slog.Attr) bool {
		needs = needsPortable(a)
		return !needs
	})
	return needs
}

func portableAttrs(as []slog.Attr) []slog.Attr {
	needs := false
	for _, a := range as {
		if needs = needsPortable(a); needs {
			break
		}
	}
	if !needs {
		return as
	}
	e := &attrEncoder{}
	for _, a := range as {
		addAttr(e, a)
	}
	return e.result()
}

// marshalerValue lets a marshaler render through LogValue for handlers that
// do not stream marshalers. This package's handlers stream it instead.
type marshalerValue struct{ m any }

func (v marshalerValue) LogValue() slog.Value { return marshalerSlogValue(v.m) }

// unwrapLogValuer returns the marshaler inside a marshalerValue.
func unwrapLogValuer(lv slog.LogValuer) any {
	if mv, ok := lv.(marshalerValue); ok {
		return mv.m
	}
	return lv
}

// marshalerSlogValue renders an object as a group and an array as []any.
func marshalerSlogValue(m any) slog.Value {
	if om, ok := asObjectMarshaler(m); ok {
		e := &attrEncoder{}
		_ = om.MarshalLogObject(e)
		return slog.GroupValue(e.result()...)
	}
	if am, ok := asArrayMarshaler(m); ok {
		a := &anyArrayEncoder{items: []any{}}
		_ = am.MarshalLogArray(a)
		return slog.AnyValue(a.items)
	}
	return slog.AnyValue(m)
}

// attrEncoder is an ObjectEncoder that builds slog attributes in order. It
// cannot record sensitive values, so marshalers redact them.
type attrEncoder struct {
	attrs []slog.Attr
	open  []attrNamespace
}

type attrNamespace struct {
	key    string
	parent []slog.Attr
}

var _ ObjectEncoder = (*attrEncoder)(nil)

func (e *attrEncoder) add(a slog.Attr) { e.attrs = append(e.attrs, a) }

func (e *attrEncoder) result() []slog.Attr {
	for i := len(e.open) - 1; i >= 0; i-- {
		ns := e.open[i]
		e.attrs = append(ns.parent, slog.Attr{Key: ns.key, Value: slog.GroupValue(e.attrs...)})
	}
	e.open = nil
	return e.attrs
}

func (e *attrEncoder) AddArray(k string, m ArrayMarshaler) error {
	a := &anyArrayEncoder{items: []any{}}
	err := m.MarshalLogArray(a)
	e.add(slog.Any(k, a.items))
	return err
}

func (e *attrEncoder) AddObject(k string, m ObjectMarshaler) error {
	sub := &attrEncoder{}
	err := m.MarshalLogObject(sub)
	e.add(slog.Attr{Key: k, Value: slog.GroupValue(sub.result()...)})
	return err
}

func (e *attrEncoder) AddBinary(k string, v []byte)     { e.add(slog.Any(k, v)) }
func (e *attrEncoder) AddByteString(k string, v []byte) { e.add(slog.String(k, string(v))) }
func (e *attrEncoder) AddBool(k string, v bool)         { e.add(slog.Bool(k, v)) }
func (e *attrEncoder) AddComplex128(k string, v complex128) {
	e.add(slog.String(k, formatComplex(v, 64)))
}
func (e *attrEncoder) AddComplex64(k string, v complex64) {
	e.add(slog.String(k, formatComplex(complex128(v), 32)))
}
func (e *attrEncoder) AddDuration(k string, v time.Duration) { e.add(slog.Duration(k, v)) }
func (e *attrEncoder) AddFloat64(k string, v float64)        { e.add(slog.Float64(k, v)) }
func (e *attrEncoder) AddFloat32(k string, v float32)        { e.add(Float32(k, v)) }
func (e *attrEncoder) AddInt(k string, v int)                { e.add(slog.Int(k, v)) }
func (e *attrEncoder) AddInt64(k string, v int64)            { e.add(slog.Int64(k, v)) }
func (e *attrEncoder) AddInt32(k string, v int32)            { e.add(slog.Int64(k, int64(v))) }
func (e *attrEncoder) AddInt16(k string, v int16)            { e.add(slog.Int64(k, int64(v))) }
func (e *attrEncoder) AddInt8(k string, v int8)              { e.add(slog.Int64(k, int64(v))) }
func (e *attrEncoder) AddString(k, v string)                 { e.add(slog.String(k, v)) }
func (e *attrEncoder) AddTime(k string, v time.Time)         { e.add(slog.Time(k, v)) }
func (e *attrEncoder) AddUint(k string, v uint)              { e.add(slog.Uint64(k, uint64(v))) }
func (e *attrEncoder) AddUint64(k string, v uint64)          { e.add(slog.Uint64(k, v)) }
func (e *attrEncoder) AddUint32(k string, v uint32)          { e.add(slog.Uint64(k, uint64(v))) }
func (e *attrEncoder) AddUint16(k string, v uint16)          { e.add(slog.Uint64(k, uint64(v))) }
func (e *attrEncoder) AddUint8(k string, v uint8)            { e.add(slog.Uint64(k, uint64(v))) }
func (e *attrEncoder) AddUintptr(k string, v uintptr)        { e.add(slog.Uint64(k, uint64(v))) }
func (e *attrEncoder) AddReflected(k string, v any) error    { e.add(slog.Any(k, v)); return nil }
func (e *attrEncoder) OpenNamespace(k string) {
	e.open = append(e.open, attrNamespace{key: k, parent: e.attrs})
	e.attrs = nil
}

// anyArrayEncoder is an ArrayEncoder that collects plain Go values; objects
// become maps.
type anyArrayEncoder struct{ items []any }

var _ ArrayEncoder = (*anyArrayEncoder)(nil)

func (a *anyArrayEncoder) add(v any) { a.items = append(a.items, v) }

func (a *anyArrayEncoder) AppendArray(m ArrayMarshaler) error {
	sub := &anyArrayEncoder{items: []any{}}
	err := m.MarshalLogArray(sub)
	a.add(sub.items)
	return err
}

func (a *anyArrayEncoder) AppendObject(m ObjectMarshaler) error {
	sub := &attrEncoder{}
	err := m.MarshalLogObject(sub)
	a.add(attrsMap(sub.result()))
	return err
}

func (a *anyArrayEncoder) AppendBool(v bool)              { a.add(v) }
func (a *anyArrayEncoder) AppendByteString(v []byte)      { a.add(string(v)) }
func (a *anyArrayEncoder) AppendComplex128(v complex128)  { a.add(formatComplex(v, 64)) }
func (a *anyArrayEncoder) AppendComplex64(v complex64)    { a.add(formatComplex(complex128(v), 32)) }
func (a *anyArrayEncoder) AppendFloat64(v float64)        { a.add(v) }
func (a *anyArrayEncoder) AppendFloat32(v float32)        { a.add(v) }
func (a *anyArrayEncoder) AppendInt(v int)                { a.add(v) }
func (a *anyArrayEncoder) AppendInt64(v int64)            { a.add(v) }
func (a *anyArrayEncoder) AppendInt32(v int32)            { a.add(v) }
func (a *anyArrayEncoder) AppendInt16(v int16)            { a.add(v) }
func (a *anyArrayEncoder) AppendInt8(v int8)              { a.add(v) }
func (a *anyArrayEncoder) AppendString(v string)          { a.add(v) }
func (a *anyArrayEncoder) AppendUint(v uint)              { a.add(v) }
func (a *anyArrayEncoder) AppendUint64(v uint64)          { a.add(v) }
func (a *anyArrayEncoder) AppendUint32(v uint32)          { a.add(v) }
func (a *anyArrayEncoder) AppendUint16(v uint16)          { a.add(v) }
func (a *anyArrayEncoder) AppendUint8(v uint8)            { a.add(v) }
func (a *anyArrayEncoder) AppendUintptr(v uintptr)        { a.add(v) }
func (a *anyArrayEncoder) AppendDuration(v time.Duration) { a.add(v) }
func (a *anyArrayEncoder) AppendTime(v time.Time)         { a.add(v) }
func (a *anyArrayEncoder) AppendReflected(v any) error    { a.add(v); return nil }

func attrsMap(as []slog.Attr) map[string]any {
	m := make(map[string]any, len(as))
	for _, a := range as {
		if v := a.Value.Resolve(); v.Kind() == slog.KindGroup {
			m[a.Key] = attrsMap(v.Group())
		} else {
			m[a.Key] = v.Any()
		}
	}
	return m
}

// formatComplex matches the encoder's complex formatting.
func formatComplex(v complex128, prec int) string {
	b := strconv.AppendFloat(nil, real(v), 'f', -1, prec)
	if imag(v) >= 0 {
		b = append(b, '+')
	}
	b = strconv.AppendFloat(b, imag(v), 'f', -1, prec)
	return string(append(b, 'i'))
}

// ---------------------------------------------------------------- tee to a slog.Handler

// slogTee duplicates entries to a slog.Handler (WithTeeHandler). The entry's
// logger name and stack trace are passed as "logger" and "stacktrace".
type slogTee struct{ h slog.Handler }

func newSlogTee(h slog.Handler) teeSink {
	if h == nil {
		return nil
	}
	return slogTee{portableHandler{h}}
}

func (t slogTee) write(ent *logEntry, r slog.Record) {
	ctx := context.Background()
	if !t.h.Enabled(ctx, ent.level) {
		return
	}
	pc := r.PC
	if pc == 0 && ent.caller.defined {
		pc = ent.caller.pc
	}
	out := slog.NewRecord(ent.time, ent.level, ent.msg, pc)
	if ent.name != "" {
		out.AddAttrs(slog.String("logger", ent.name))
	}
	r.Attrs(func(a slog.Attr) bool {
		out.AddAttrs(a)
		return true
	})
	if ent.stack != "" {
		out.AddAttrs(slog.String("stacktrace", ent.stack))
	}
	_ = t.h.Handle(ctx, out)
}

func (t slogTee) withValues(kvs []any) teeSink {
	return slogTee{t.h.WithAttrs(valuesAttrs(kvs))}
}

func (t slogTee) withAttrs(as []slog.Attr) teeSink {
	h := t.h
	start := 0
	for i, a := range as {
		if _, ok := a.Value.Any().(namespaceMarker); ok {
			if i > start {
				h = h.WithAttrs(as[start:i])
			}
			h = h.WithGroup(a.Key)
			start = i + 1
		}
	}
	if start < len(as) {
		h = h.WithAttrs(as[start:])
	}
	return slogTee{h}
}

// withGate keeps the handler: entries reach a tee only once the gate passed.
func (t slogTee) withGate(levelGate) teeSink { return t }

// multiTee fans out to several tees.
type multiTee []teeSink

func newMultiTee(tees ...teeSink) teeSink {
	var m multiTee
	for _, t := range tees {
		if t != nil {
			m = append(m, t)
		}
	}
	switch len(m) {
	case 0:
		return nil
	case 1:
		return m[0]
	}
	return m
}

func (m multiTee) write(ent *logEntry, r slog.Record) {
	for _, t := range m {
		t.write(ent, r)
	}
}

func (m multiTee) each(f func(teeSink) teeSink) teeSink {
	out := make(multiTee, len(m))
	for i, t := range m {
		out[i] = f(t)
	}
	return out
}

func (m multiTee) withValues(kvs []any) teeSink {
	return m.each(func(t teeSink) teeSink { return t.withValues(kvs) })
}

func (m multiTee) withAttrs(as []slog.Attr) teeSink {
	return m.each(func(t teeSink) teeSink { return t.withAttrs(as) })
}

func (m multiTee) withGate(g levelGate) teeSink {
	return m.each(func(t teeSink) teeSink { return t.withGate(g) })
}
