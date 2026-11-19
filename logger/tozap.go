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

// This file holds everything that exposes zap types: ToZap, and the adapters
// that let the slog backend serve the zap-typed parts of ZapLogger (levelers
// and tees from zaputil). Nothing else in the slog backend imports zap.

import (
	"context"
	"log/slog"
	"slices"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger/zaputil"
)

// ToZap returns a *zap.SugaredLogger that writes like this logger.
//
// Deprecated: ToZap exists for libraries that require a *zap.Logger. Use
// ToSlogHandler with slog.New where a *slog.Logger is accepted instead.
func (l *zapLogger) ToZap() *zap.SugaredLogger {
	return l.zap.WithOptions(zap.AddCallerSkip(-1))
}

// ToZap returns a *zap.SugaredLogger whose entries go through this logger's
// handler chain, so levels, name, context, tee, sampling and deferred values
// all apply.
//
// Deprecated: ToZap exists for libraries that require a *zap.Logger. Use
// ToSlogHandler with slog.New where a *slog.Logger is accepted instead.
func (l *slogLogger) ToZap() *zap.SugaredLogger {
	z := zap.New(&bridgeCore{l.h, l.sampler},
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
		zap.WithClock(clockFunc(l.now)),
		zap.AddCallerSkip(l.depth),
	)
	if n := l.h.loggerName(); n != "" {
		z = z.Named(n)
	}
	return z.Sugar()
}

type clockFunc func() time.Time

func (f clockFunc) Now() time.Time                         { return f() }
func (f clockFunc) NewTicker(d time.Duration) *time.Ticker { return time.NewTicker(d) }

// bridgeCore lets a *zap.Logger write through a slog handler chain.
type bridgeCore struct {
	h       chainHandler
	sampler *lineSampler
}

func (c *bridgeCore) Enabled(l zapcore.Level) bool {
	return c.h.Enabled(context.Background(), slog.Level(l)*4)
}

func (c *bridgeCore) With(fs []zapcore.Field) zapcore.Core {
	attrs := make([]slog.Attr, len(fs))
	for i, f := range fs {
		attrs[i], _ = asBareField(f)
	}
	return &bridgeCore{c.h.WithAttrs(attrs).(chainHandler), c.sampler}
}

func (c *bridgeCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if !c.Enabled(ent.Level) {
		return ce
	}
	if c.sampler != nil && !c.sampler.allow(slog.Level(ent.Level)*4, ent.Message, ent.Time) {
		return ce
	}
	return ce.AddCore(ent, c)
}

func (c *bridgeCore) Write(ent zapcore.Entry, fs []zapcore.Field) error {
	r := slog.NewRecord(ent.Time, slog.Level(ent.Level)*4, ent.Message, 0)
	for _, f := range fs {
		a, _ := asBareField(f)
		r.AddAttrs(a)
	}
	ctx := context.WithValue(context.Background(), entryMetaKey{}, &entryMeta{
		caller: callerInfo{
			defined:  ent.Caller.Defined,
			pc:       ent.Caller.PC,
			file:     ent.Caller.File,
			line:     ent.Caller.Line,
			function: ent.Caller.Function,
		},
		stack: ent.Stack,
		name:  ent.LoggerName,
	})
	return c.h.Handle(ctx, r)
}

func (c *bridgeCore) Sync() error { return nil }

// ---------------------------------------------------------------- levelers

// zapCompat holds the zaputil leveler handed out by Leveler, created on first
// use so loggers that never ask for it do not maintain it.
type zapCompat struct {
	once sync.Once
	root *zaputil.ComponentLeveler
}

// zapLevelSource adapts a zaputil.ComponentLeveler given to WithComponentLeveler.
type zapLevelSource struct{ lv *zaputil.ComponentLeveler }

func (s zapLevelSource) gate(component string) levelGate {
	return zapGate{s.lv.WriteEnabler(component)}
}

type zapGate struct{ we *zaputil.WriteEnabler }

func (g zapGate) enabled(l slog.Level) bool        { return g.we.Enabled(zapcore.Level(l / 4)) }
func (g zapGate) Write(p []byte) (int, error)      { return g.we.Write(p) }
func (g gateEnabler) Enabled(l zapcore.Level) bool { return g.g.enabled(slog.Level(l) * 4) }

// gateEnabler exposes a levelGate as a zapcore.LevelEnabler.
type gateEnabler struct{ g levelGate }

func (l *slogLogger) Leveler() *ComponentLeveler {
	if s, ok := l.levels.(zapLevelSource); ok {
		return s.lv
	}
	l.compat.once.Do(func() {
		root := zaputil.NewRootComponentLeveler(zapcore.AddSync(l.out), l.conf)
		l.conf.AddUpdateObserver(func(*Config) error {
			root.Refresh()
			return nil
		})
		l.compat.root = root
	})
	return l.compat.root
}

func (l *slogLogger) WithComponentLeveler(lv *ComponentLeveler) Logger {
	if lv == nil {
		return l
	}
	c := l.dup()
	c.levels = zapLevelSource{lv}
	c.h = l.h.withGate(c.levels.gate(l.comp))
	return c
}

type slogComponentLeveler struct{ l *slogLogger }

func (c slogComponentLeveler) ComponentLevel(component string) zapcore.LevelEnabler {
	if c.l.comp != "" {
		component = c.l.comp + "." + component
	}
	return gateEnabler{c.l.levels.gate(component)}
}

func (l *slogLogger) ComponentLeveler() ZapComponentLeveler {
	return slogComponentLeveler{l}
}

// ---------------------------------------------------------------- tee

// zapTee adapts a zaputil.Tee (a zapcore.Core factory) to teeSink.
type zapTee struct {
	t     zaputil.Tee
	gate  levelGate
	extra []zapcore.Field
	core  zapcore.Core
}

// newZapTee returns nil for the zero Tee.
func newZapTee(t zaputil.Tee, g levelGate) teeSink {
	zt := &zapTee{t: t, gate: g}
	if zt.build(); zt.core == nil {
		return nil
	}
	return zt
}

func (t *zapTee) build() {
	t.core = t.t.Core(gateEnabler{t.gate})
	if t.core != nil && len(t.extra) > 0 {
		t.core = t.core.With(t.extra)
	}
}

func (t *zapTee) clone() *zapTee {
	c := *t
	c.extra = slices.Clip(t.extra)
	return &c
}

func (t *zapTee) write(ent *logEntry, r slog.Record) {
	ze := zapcore.Entry{
		Level:      zapcore.Level(ent.level / 4),
		Time:       ent.time,
		LoggerName: ent.name,
		Message:    ent.msg,
		Caller: zapcore.EntryCaller{
			Defined:  ent.caller.defined,
			PC:       ent.caller.pc,
			File:     ent.caller.file,
			Line:     ent.caller.line,
			Function: ent.caller.function,
		},
		Stack: ent.stack,
	}
	if ce := t.core.Check(ze, nil); ce != nil {
		fields := make([]zapcore.Field, 0, r.NumAttrs())
		r.Attrs(func(a slog.Attr) bool {
			fields = append(fields, attrField(a))
			return true
		})
		ce.Write(fields...)
	}
}

func (t *zapTee) withValues(kvs []any) teeSink {
	c := t.clone()
	c.t = t.t.WithValues(kvs...)
	c.build()
	return c
}

func (t *zapTee) withAttrs(as []slog.Attr) teeSink {
	c := t.clone()
	for _, a := range as {
		if _, ok := a.Value.Any().(namespaceMarker); ok {
			c.extra = append(c.extra, zap.Namespace(a.Key))
		} else {
			c.extra = append(c.extra, attrField(a))
		}
	}
	c.build()
	return c
}

func (t *zapTee) withGate(g levelGate) teeSink {
	c := t.clone()
	c.gate = g
	c.build()
	return c
}

// attrField converts an attr to the zap field the sugared logger would have
// produced, so tee cores see the same field types as with the zap backend.
func attrField(a slog.Attr) zapcore.Field {
	v := a.Value
	switch v.Kind() {
	case slog.KindString:
		return zap.String(a.Key, v.String())
	case slog.KindInt64:
		return zap.Int64(a.Key, v.Int64())
	case slog.KindUint64:
		return zap.Uint64(a.Key, v.Uint64())
	case slog.KindFloat64:
		return zap.Float64(a.Key, v.Float64())
	case slog.KindBool:
		return zap.Bool(a.Key, v.Bool())
	case slog.KindDuration:
		return zap.Duration(a.Key, v.Duration())
	case slog.KindTime:
		return zap.Time(a.Key, v.Time())
	case slog.KindGroup:
		return zap.Object(a.Key, attrGroup(v.Group()))
	case slog.KindLogValuer:
		if om, ok := v.LogValuer().(ObjectMarshaler); ok {
			return zap.Object(a.Key, om)
		}
		return attrField(slog.Attr{Key: a.Key, Value: v.Resolve()})
	}
	switch x := v.Any().(type) {
	case bareField:
		return zapcore.Field(x)
	case float32Value:
		return zap.Float32(a.Key, float32(x))
	default:
		return zap.Any(a.Key, x)
	}
}
