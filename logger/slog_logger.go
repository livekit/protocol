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
	"io"
	"log/slog"
	"runtime"
	"time"
)

// slogLogger is the log/slog-based implementation of ZapLogger, selected
// with WithSlogBackend. Its output is byte-identical to zapLogger's (see
// slog_compat_test.go); the zap-typed parts of the API are in tozap.go.
//
//	slogLogger --(key/values, sampling, call site)--> chainHandler
//	chainHandler: [deferredHandler] -> coreHandler (encode, write, tee)
type slogLogger struct {
	*slogShared
	h       chainHandler
	levels  levelSource
	comp    string
	depth   int
	sampler *lineSampler
}

var _ ZapLogger = (*slogLogger)(nil)

// slogShared is common to every logger derived from one root.
type slogShared struct {
	conf   *Config
	out    io.Writer
	now    func() time.Time
	compat zapCompat
}

func newSlogLogger(zc *zapConfig) *slogLogger {
	conf := zc.conf
	now := zc.now
	if now == nil {
		now = time.Now
	}
	root := newSlogLeveler(zc.out, conf)
	conf.AddUpdateObserver(func(*Config) error {
		root.refresh()
		return nil
	})
	gate := root.gate("")
	l := &slogLogger{
		slogShared: &slogShared{conf: conf, out: zc.out, now: now},
		h:          newCoreHandler(conf.JSON, gate, newMultiTee(newZapTee(zc.tee, gate), newSlogTee(zc.teeHandler))),
		levels:     root,
	}
	if conf.Sample {
		// Matches FromZapLogger, including its use of the ItemSample* settings.
		initial, interval := 20, 100
		if conf.ItemSampleInitial != 0 {
			initial = conf.ItemSampleInitial
		}
		if conf.ItemSampleInterval != 0 {
			interval = conf.ItemSampleInterval
		}
		l.sampler = newLineSampler(time.Second, initial, interval)
	}
	return l
}

// handler returns the slog.Handler behind this logger, for slog.New.
func (l *slogLogger) handler() slog.Handler {
	if l.sampler != nil {
		return samplingHandler{l.h, l.sampler}
	}
	return l.h
}

func (l *slogLogger) log(lvl slog.Level, msg string, err error, kvs []any) {
	ctx := context.Background()
	if !l.h.Enabled(ctx, lvl) {
		return
	}
	t := l.now()
	if l.sampler != nil && !l.sampler.allow(lvl, msg, t) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(3+l.depth, pcs[:])
	if err != nil {
		kvs = append(kvs, "error", err)
	}
	attrs := sweeten(kvs, func(diag string, a slog.Attr) {
		// The zap backend logs these through the same sampled core.
		t := l.now()
		if !l.h.Enabled(ctx, slog.LevelError) || (l.sampler != nil && !l.sampler.allow(slog.LevelError, diag, t)) {
			return
		}
		r := slog.NewRecord(t, slog.LevelError, diag, pcs[0])
		r.AddAttrs(a)
		_ = l.h.Handle(ctx, r)
	})
	r := slog.NewRecord(t, lvl, msg, pcs[0])
	r.AddAttrs(attrs...)
	_ = l.h.Handle(ctx, r)
}

func (l *slogLogger) Debugw(msg string, kvs ...any) { l.log(slog.LevelDebug, msg, nil, kvs) }
func (l *slogLogger) Infow(msg string, kvs ...any)  { l.log(slog.LevelInfo, msg, nil, kvs) }
func (l *slogLogger) Warnw(msg string, err error, kvs ...any) {
	l.log(slog.LevelWarn, msg, err, kvs)
}
func (l *slogLogger) Errorw(msg string, err error, kvs ...any) {
	l.log(slog.LevelError, msg, err, kvs)
}

func (l *slogLogger) dup() *slogLogger {
	c := *l
	return &c
}

func (l *slogLogger) WithValues(kvs ...any) Logger {
	c := l.dup()
	c.h = l.h.withValues(kvs)
	return c
}

func (l *slogLogger) WithUnlikelyValues(kvs ...any) UnlikelyLogger {
	return UnlikelyLogger{l, kvs}
}

func (l *slogLogger) WithName(name string) Logger {
	c := l.dup()
	c.h = l.h.withName(name)
	return c
}

func (l *slogLogger) WithComponent(component string) Logger {
	c := l.dup()
	if c.comp == "" {
		c.comp = component
	} else {
		c.comp += "." + component
	}
	c.h = l.h.withName(component).withGate(l.levels.gate(c.comp))
	return c
}

func (l *slogLogger) WithCallDepth(depth int) Logger {
	c := l.dup()
	c.depth += depth
	return c
}

func (l *slogLogger) WithItemSampler() Logger {
	if l.conf.ItemSampleSeconds == 0 {
		return l
	}
	c := l.dup()
	c.sampler = newLineSampler(
		time.Duration(l.conf.ItemSampleSeconds)*time.Second,
		l.conf.ItemSampleInitial,
		l.conf.ItemSampleInterval,
	)
	return c
}

func (l *slogLogger) WithoutSampler() Logger {
	c := l.dup()
	c.sampler = nil
	return c
}

// withAttrs and withGroup apply slog.Handler semantics, for the slog view of
// the logger.
func (l *slogLogger) withAttrs(as []slog.Attr) *slogLogger {
	c := l.dup()
	c.h = l.h.WithAttrs(as).(chainHandler)
	return c
}

func (l *slogLogger) withGroup(name string) *slogLogger {
	c := l.dup()
	c.h = l.h.WithGroup(name).(chainHandler)
	return c
}

func (l *slogLogger) withLevels(src levelSource) *slogLogger {
	c := l.dup()
	c.levels = src
	c.h = l.h.withGate(src.gate(l.comp))
	return c
}

func (l *slogLogger) WithDeferredValues() (Logger, DeferredFieldResolver) {
	c := l.dup()
	d := &slogDeferrer{}
	c.h = &deferredHandler{l.h, d}
	return c, d
}

// ---------------------------------------------------------------- key/values

type invalidPair struct {
	position   int
	key, value any
}

func (p invalidPair) MarshalLogObject(enc ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	addAny(enc, "key", p.key)
	addAny(enc, "value", p.value)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc ArrayEncoder) error {
	for _, p := range ps {
		_ = enc.AppendObject(p)
	}
	return nil
}

// sweeten converts sugared key/values with zap's SugaredLogger rules: zap
// fields pass through, the first bare error becomes "error", and malformed
// input is reported through diag as separate error entries.
func sweeten(args []any, diag func(msg string, a slog.Attr)) []slog.Attr {
	attrs := make([]slog.Attr, 0, len(args))
	var invalid invalidPairs
	seenError := false
	for i := 0; i < len(args); {
		if f, ok := asBareField(args[i]); ok {
			attrs = append(attrs, f)
			i++
			continue
		}
		if err, ok := args[i].(error); ok {
			if !seenError {
				seenError = true
				attrs = append(attrs, slog.Any("error", err))
			} else {
				diag("Multiple errors without a key.", slog.Any("error", err))
			}
			i++
			continue
		}
		if i == len(args)-1 {
			diag("Ignored key without a value.", slog.Any("ignored", args[i]))
			break
		}
		if k, ok := args[i].(string); ok {
			attrs = append(attrs, attrOf(k, args[i+1]))
		} else {
			invalid = append(invalid, invalidPair{i, args[i], args[i+1]})
		}
		i += 2
	}
	if len(invalid) > 0 {
		diag("Ignored key-value pairs with non-string keys.", slog.Any("invalid", ArrayMarshaler(invalid)))
	}
	return attrs
}
