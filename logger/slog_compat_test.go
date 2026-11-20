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

package logger_test

// These tests run the same calls through the zap backend and the slog backend
// (WithSlogBackend) and require byte-identical output, on the primary output
// and on the tee, in JSON and console format. Both loggers are called from the
// same source line, so callers and stack traces are comparable.

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/logger/medialogutils"
	"github.com/livekit/protocol/logger/pionlogger"
	"github.com/livekit/protocol/logger/testutil"
	"github.com/livekit/protocol/logger/zaputil"
)

var fixedNow = time.Date(2026, 11, 19, 12, 34, 56, 789_123_456, time.UTC)

func fixedClock() time.Time { return fixedNow }

// wrap adds one stack frame, for WithCallDepth(1).
type wrap struct{ l logger.Logger }

func (w wrap) Infow(m string, kv ...any)           { w.l.Infow(m, kv...) }
func (w wrap) Errorw(m string, e error, kv ...any) { w.l.Errorw(m, e, kv...) }

type both struct {
	conf *logger.Config
	z, s logger.Logger
}

// log runs fn against both loggers from the same source line.
func (b *both) log(fn func(l logger.Logger)) {
	for _, l := range []logger.Logger{b.z, b.s} {
		fn(l)
	}
}

func (b *both) derive(fn func(l logger.Logger) logger.Logger) {
	b.z, b.s = fn(b.z), fn(b.s)
}

// outputs holds, per backend, the primary output, the zap-core tee
// (WithTee) and the slog-handler tee (WithTeeHandler).
type outputs struct{ z, s, zt, st, zh, sh bytes.Buffer }

func newBoth(t *testing.T, conf *logger.Config, o *outputs) *both {
	t.Helper()
	newLogger := func(out, tee, teeHandler *bytes.Buffer, extra ...logger.ZapLoggerOption) logger.Logger {
		opts := append([]logger.ZapLoggerOption{
			logger.WithOutputForTest(out),
			logger.WithClockForTest(fixedClock),
			logger.WithTee(zaputil.NewTee(testutil.NewJSONCoreFactory(zapcore.AddSync(tee)))),
			logger.WithTeeHandler(slog.NewJSONHandler(teeHandler, &slog.HandlerOptions{Level: slog.LevelDebug})),
		}, extra...)
		l, err := logger.NewZapLogger(conf, opts...)
		require.NoError(t, err)
		return l
	}
	return &both{
		conf: conf,
		z:    newLogger(&o.z, &o.zt, &o.zh),
		s:    newLogger(&o.s, &o.st, &o.sh, logger.WithSlogBackend()),
	}
}

func participant() *livekit.ParticipantInfo {
	return &livekit.ParticipantInfo{
		Sid: "PA_abcdef", Identity: "user-1", Name: "Alice", State: livekit.ParticipantInfo_ACTIVE,
		JoinedAt: 1795000000, Version: 3, IsPublisher: true, Region: "us-east",
		Tracks: []*livekit.TrackInfo{
			{Sid: "TR_a", Type: livekit.TrackType_AUDIO, Name: "mic", Source: livekit.TrackSource_MICROPHONE},
			{Sid: "TR_v", Type: livekit.TrackType_VIDEO, Name: "cam", Width: 1280, Height: 720, Source: livekit.TrackSource_CAMERA},
		},
		Attributes: map[string]string{"k": "v"},
	}
}

type stringer struct{ v string }

func (s stringer) String() string { return "S(" + s.v + ")" }

type verboseErr struct{ msg string }

func (e verboseErr) Error() string { return e.msg }
func (e verboseErr) Format(f fmt.State, c rune) {
	if c == 'v' && f.Flag('+') {
		fmt.Fprintf(f, "%s\nstack:\n\tmain.go:1", e.msg)
		return
	}
	fmt.Fprint(f, e.msg)
}

type nilPtrErr struct{ msg string }

func (e *nilPtrErr) Error() string { return e.msg }

type obj struct{ a string }

func (o obj) MarshalLogObject(e logger.ObjectEncoder) error {
	e.AddString("a", o.a)
	e.OpenNamespace("ns")
	e.AddInt("n", 1)
	return nil
}

// zapObj implements the zapcore flavour of the marshaler interface, like
// types written before logger.ObjectEncoder existed.
type zapObj struct{ a string }

func (o zapObj) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddString("za", o.a)
	return e.AddObject("inner", logger.ZapObjectMarshaler(obj{o.a}))
}

func pkgLog() {
	logger.Infow("pkg info", "k", 1)
	logger.Errorw("pkg err", errors.New("e"))
}

type tcase struct {
	name string
	conf func(c *logger.Config)
	run  func(b *both)
}

func cases(t *testing.T) []tcase {
	p := participant()
	str := "pointee"
	var nilStr *string
	var nilErr *nilPtrErr
	return []tcase{
		{name: "scalars", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("scalars", "s", "str", "i", 42, "i8", int8(-3), "u", uint(7), "u64", uint64(1<<63), "f", 3.25,
					"f32", float32(0.1), "b", true, "d", 1500*time.Millisecond, "t", time.Date(2025, 1, 2, 3, 4, 5, 6, time.UTC), "nil", nil, "uptr", uintptr(9))
			})
		}},
		{name: "floats", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("floats", "nan", math.NaN(), "inf", math.Inf(1), "ninf", math.Inf(-1), "big", 1e21, "small", 1e-7, "negz", math.Copysign(0, -1))
			})
		}},
		{name: "strings", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("msg \"q\"\n\x01", "ctl", "a\x00b\x08c\x0cd\x1f", "q", "quote\" back\\ slash", "nl", "l1\nl2\r\t", "html", "<a&b>",
					"uni", "héllo 🌍", "bad", "\xff\xfe", "ls", " ", "k\"ey", 1)
			})
		}},
		{name: "bytes/complex", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("bc", "bytes", []byte("hi\x00"), "c128", complex(1, -2), "c64", complex64(complex(0.1, 0.2)))
			})
		}},
		{name: "slices/maps/structs", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("coll", "ss", []string{"a", "b"}, "is", []int{1, 2}, "bs", []bool{true}, "ds", []time.Duration{time.Second},
					"m", map[string]int{"b": 2, "a": 1}, "st", struct {
						A int
						B string
					}{1, "<x>"}, "unexp", obj{"x"}, "ch", make(chan int), "pid", livekit.ParticipantID("PA_1"))
			})
		}},
		{name: "pointers/stringer", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("ptr", "p", &str, "np", nilStr, "sg", stringer{"v"}, "sgp", &stringer{"w"})
			})
		}},
		{name: "errors", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("errs", "plain", errors.New("boom"), "wrapped", fmt.Errorf("ctx: %w", errors.New("inner")),
					"multi", multierr.Combine(errors.New("a"), verboseErr{"b"}), "verbose", verboseErr{"v"}, "nilptr", error(nilErr))
			})
		}},
		{name: "proto/marshalers", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("proto", "p", logger.Proto(p), "lim", logger.ProtoWithLimit(p, 10), "tracks", logger.ProtoSlice(p.Tracks),
					"ids", logger.StringSlice([]string{"x", "y"}), "obj", obj{"o"}, "zobj", zapObj{"z"}, "after", 1)
			})
		}},
		{name: "zap fields", run: func(b *both) {
			b.log(func(l logger.Logger) {
				l.Infow("fields", zap.String("zf", "v"), zap.Inline(zapObj{"in"}), zap.Skip(), "k", 1, zap.Namespace("ns"), "inner", 2)
			})
		}},
		{name: "sweeten quirks", run: func(b *both) {
			b.log(func(l logger.Logger) { l.Infow("keyless err", errors.New("x"), "k", 1) })
			b.log(func(l logger.Logger) { l.Infow("empty") })
		}},
		{name: "levels/errors", run: func(b *both) {
			b.log(func(l logger.Logger) { l.Debugw("dbg", "k", 1) })
			b.log(func(l logger.Logger) { l.Warnw("warn", errors.New("w"), "k", 1) })
			b.log(func(l logger.Logger) { l.Warnw("warn nil", nil) })
			b.log(func(l logger.Logger) { l.Errorw("error", verboseErr{"e"}, "k", 1) })
			b.log(func(l logger.Logger) { l.Errorw("error nil", nil) })
		}},
		{name: "WithValues", run: func(b *both) {
			b.derive(func(l logger.Logger) logger.Logger {
				return l.WithValues("room", "RM_1", 5, "dropped", "n", 3, "d", time.Second, "p", logger.Proto(p.Tracks[0]), zap.String("f", "x"), "tail")
			})
			b.log(func(l logger.Logger) { l.Infow("ctx", "k", "v") })
			b.derive(func(l logger.Logger) logger.Logger { return l.WithValues("more", true) })
			b.log(func(l logger.Logger) { l.Infow("ctx2") })
			b.log(func(l logger.Logger) { l.Infow("ctx ns", zap.Namespace("ns"), "k", 1) })
		}},
		{name: "names/components", run: func(b *both) {
			b.derive(func(l logger.Logger) logger.Logger { return l.WithName("lk") })
			b.log(func(l logger.Logger) { l.Infow("named") })
			b.log(func(l logger.Logger) { l.WithComponent("rtc").WithComponent("sub").Infow("component") })
			b.log(func(l logger.Logger) {
				q := l.WithComponent("quiet")
				q.Infow("suppressed")
				q.Warnw("quiet warn", nil)
			})
		}},
		{name: "call depth", run: func(b *both) {
			b.log(func(l logger.Logger) {
				w := wrap{l.WithCallDepth(1)}
				w.Infow("via wrapper")
				w.Errorw("via wrapper err", errors.New("e"))
			})
		}},
		{name: "deferred", run: func(b *both) {
			z, zd := b.z.WithDeferredValues()
			s, sd := b.s.WithDeferredValues()
			z, s = z.WithValues("ctx", 1), s.WithValues("ctx", 1)
			for _, l := range []logger.Logger{z, s} {
				l.Infow("buffered", "k", 1)
				l.Errorw("buffered err", errors.New("e"))
			}
			zd.Resolve("pID", "PA_1", 7, "x", zap.Int("zf", 2))
			sd.Resolve("pID", "PA_1", 7, "x", zap.Int("zf", 2))
			for _, l := range []logger.Logger{z, s} {
				l.Infow("direct", "k", 2)
			}
			zd.Resolve("pID", "PA_2", "room", "RM")
			sd.Resolve("pID", "PA_2", "room", "RM")
			for _, l := range []logger.Logger{z, s} {
				l.Infow("re-resolved")
			}
		}},
		{name: "global sampler", conf: func(c *logger.Config) {
			c.Sample, c.ItemSampleInitial, c.ItemSampleInterval = true, 3, 4
		}, run: func(b *both) {
			b.log(func(l logger.Logger) {
				for i := range 20 {
					l.Infow("hot", "i", i)
				}
				l.Infow("cold")
				l.WithValues("derived", true).Infow("hot", "i", 99) // shares counters
			})
		}},
		{name: "item sampler", conf: func(c *logger.Config) {
			c.ItemSampleSeconds, c.ItemSampleInitial, c.ItemSampleInterval = 1, 2, 3
		}, run: func(b *both) {
			b.log(func(l logger.Logger) {
				il := l.WithItemSampler()
				for i := range 10 {
					il.Infow("item", "i", i)
				}
				ul := il.WithoutSampler()
				for i := range 3 {
					ul.Infow("item", "i", i)
				}
			})
		}},
		{name: "ToZap", run: func(b *both) {
			b.derive(func(l logger.Logger) logger.Logger {
				return l.WithName("lk").WithComponent("rtc").WithValues("room", "RM_1")
			})
			b.log(func(l logger.Logger) {
				tz := l.(logger.ZapLogger).ToZap()
				tz.Infow("bridge", "k", 1, "d", time.Second)
				tz.Desugar().Info("desugared", zap.Int("n", 2), zap.Object("o", zapObj{"x"}))
				tz.Named("sub").With("w", 1).Warnw("named with", "e", errors.New("x"))
				tz.Errorw("bridge err", "error", errors.New("boom"))
				tz.Desugar().WithOptions(zap.WithCaller(false)).Info("nocaller")
				tz.Debugw("dbg")
				zap.NewStdLog(tz.Desugar()).Print("std log")
			})
		}},
		{name: "ToZap deferred+sampled", conf: func(c *logger.Config) {
			c.Sample, c.ItemSampleInitial, c.ItemSampleInterval = true, 2, 5
		}, run: func(b *both) {
			z, zd := b.z.WithDeferredValues()
			s, sd := b.s.WithDeferredValues()
			for _, l := range []logger.Logger{z, s} {
				tz := l.(logger.ZapLogger).ToZap()
				for i := range 8 {
					tz.Infow("bridged hot", "i", i)
				}
			}
			zd.Resolve("pID", "PA_1")
			sd.Resolve("pID", "PA_1")
		}},
		{name: "component leveler", run: func(b *both) {
			b.log(func(l logger.Logger) {
				zl := l.(logger.ZapLogger)
				widened := zl.WithComponentLeveler(zaputil.NewComponentLeveler(zl.Leveler(), logger.FixedComponentLevel(zapcore.DebugLevel)))
				widened.WithComponent("quiet").Infow("widened")
				l.WithComponent("quiet").Infow("still quiet")
				l.Infow("enabled?", "quietInfo", zl.ComponentLeveler().ComponentLevel("quiet").Enabled(zapcore.InfoLevel),
					"rootDebug", zl.ComponentLeveler().ComponentLevel("x").Enabled(zapcore.DebugLevel))
			})
		}},
		{name: "config hot reload", run: func(b *both) {
			qz, qs := b.z.WithComponent("quiet"), b.s.WithComponent("quiet")
			for _, l := range []logger.Logger{qz, qs} {
				l.Infow("before reload")
			}
			c := b.conf
			_ = c.Update(&logger.Config{JSON: c.JSON, Level: "info", ComponentLevels: map[string]string{"quiet": "debug"}})
			for _, l := range []logger.Logger{qz, qs, b.z, b.s} {
				l.Debugw("after reload")
			}
		}},
		{name: "package funcs", run: func(b *both) {
			for _, l := range []logger.Logger{b.z, b.s} {
				logger.SetLogger(l, "pkg")
				pkgLog()
			}
			logger.SetLogger(logger.GetDiscardLogger(), "")
		}},
		{name: "unlikely/override", run: func(b *both) {
			b.log(func(l logger.Logger) {
				u := l.WithUnlikelyValues("u", 1)
				u.Infow("unlikely", "k", 2)
				u.WithValues("u2", 3).Errorw("unlikely err", errors.New("e"))
				o := medialogutils.NewOverrideLogger(l)
				o.Errorw("overridden", errors.New("x"))
				o.WithValues("a", 1).Infow("override values")
			})
		}},
		{name: "slog round trip", run: func(b *both) {
			b.log(func(l logger.Logger) {
				s := logger.ToSlog(l).With("k", 1, "p", logger.Proto(p.Tracks[0]))
				logger.FromSlog(s).WithComponent("rtc").Infow("round trip", "x", 2)
				require.Equal(t, l, logger.FromSlog(logger.ToSlog(l)))
			})
		}},
		{name: "slog helpers", run: func(b *both) {
			b.log(func(l logger.Logger) {
				s := logger.WithComponent(logger.WithName(logger.ToSlog(l), "lk"), "rtc")
				logger.FromSlog(s).Infow("named via slog")
				d, r := logger.WithDeferredValues(s)
				logger.FromSlog(d).Infow("deferred via slog")
				r.Resolve("pID", "PA_1")
				logger.FromSlog(logger.WithoutSampler(logger.WithItemSampler(s))).Infow("sampler via slog")
			})
		}},
		{name: "component levels", run: func(b *both) {
			levels := map[string]slog.Level{"quiet": slog.LevelDebug}
			resolve := func(c string) (slog.Level, bool) {
				l, ok := levels[c]
				return l, ok
			}
			zc, sc := logger.NewComponentLevels(b.z, resolve), logger.NewComponentLevels(b.s, resolve)
			z, s := logger.WithComponentLevels(b.z, zc).WithComponent("quiet"), logger.WithComponentLevels(b.s, sc).WithComponent("quiet")
			for _, l := range []logger.Logger{z, s} {
				l.Infow("widened")
			}
			delete(levels, "quiet")
			zc.Refresh()
			sc.Refresh()
			for _, l := range []logger.Logger{z, s} {
				l.Infow("back to warn")
				l.Warnw("still warn", nil)
			}
		}},
		{name: "pion", run: func(b *both) {
			b.log(func(l logger.Logger) {
				pl := pionlogger.NewLoggerFactory(l).NewLogger("ice")
				pl.Infof("hello %d", 1)
				pl.Debug("dbg")
				pl.Warnf("warn %s", "x")
				pl.Errorf("err %s", "y")
				pl.Info("pingAllCandidates called with no candidate pairs") // filtered
				pl.Trace("trace")
			})
		}},
	}
}

func diffLines(t *testing.T, what, z, s string) {
	t.Helper()
	if z == s {
		return
	}
	zl, sl := strings.Split(z, "\n"), strings.Split(s, "\n")
	for i := range max(len(zl), len(sl)) {
		var a, b string
		if i < len(zl) {
			a = zl[i]
		}
		if i < len(sl) {
			b = sl[i]
		}
		if a != b {
			t.Errorf("%s line %d differs\n zap:  %q\n slog: %q", what, i, a, b)
		}
	}
}

func TestSlogBackendMatchesZap(t *testing.T) {
	for _, json := range []bool{true, false} {
		for _, c := range cases(t) {
			t.Run(fmt.Sprintf("json=%v/%s", json, c.name), func(t *testing.T) {
				conf := &logger.Config{JSON: json, Level: "debug", ComponentLevels: map[string]string{"quiet": "warn"}}
				if c.conf != nil {
					c.conf(conf)
				}
				var o outputs
				b := newBoth(t, conf, &o)
				c.run(b)
				require.NotZero(t, o.z.Len(), "no output")
				require.NotZero(t, o.zt.Len(), "no tee output")
				diffLines(t, "main", o.z.String(), o.s.String())
				diffLines(t, "tee", o.zt.String(), o.st.String())
				diffLines(t, "slog tee", o.zh.String(), o.sh.String())
			})
		}
	}
}

// The zap backend reports malformed key/values in separate entries, which it
// attributes to its own wrapper; the slog backend attributes them to the
// caller. Everything else matches.
func TestSlogBackendSweetenDiagnostics(t *testing.T) {
	var o outputs
	b := newBoth(t, &logger.Config{JSON: true, Level: "debug"}, &o)
	b.log(func(l logger.Logger) { l.Infow("two keyless", errors.New("x"), errors.New("y")) })
	b.log(func(l logger.Logger) { l.Infow("odd", "k", 1, "dangling") })
	b.log(func(l logger.Logger) { l.Infow("nonstring", 1, "v", "ok", 2) })
	strip := regexp.MustCompile(`"caller":"[^"]*"|"stacktrace":"(?:[^"\\]|\\.)*"`)
	require.Equal(t, 6, strings.Count(o.z.String(), "\n"))
	diffLines(t, "main", strip.ReplaceAllString(o.z.String(), ""), strip.ReplaceAllString(o.s.String(), ""))
	diffLines(t, "tee", strip.ReplaceAllString(o.zt.String(), ""), strip.ReplaceAllString(o.st.String(), ""))
}

func TestSlogBackendToSlogHandler(t *testing.T) {
	var o outputs
	b := newBoth(t, &logger.Config{JSON: true, Level: "debug"}, &o)
	s := b.s.WithName("lk").WithValues("room", "RM_1")
	slog.New(logger.ToSlogHandler(s)).With("k", 1).WithGroup("g").Info("via slog", "n", 2)
	line := o.s.String()
	require.Contains(t, line, `"logger":"lk"`)
	require.Contains(t, line, `"caller":"logger/slog_compat_test.go:`)
	require.Contains(t, line, `"msg":"via slog","room":"RM_1","k":1,"g":{"n":2}}`)
	require.Contains(t, o.st.String(), `"g":{"n":2}`)
}

func TestSlogBackendRejectsZapLogger(t *testing.T) {
	_, err := logger.FromZapLogger(zap.NewNop(), &logger.Config{}, logger.WithSlogBackend())
	require.Error(t, err)
}

// The comparison tests above would pass if both backends rendered values the
// same wrong way, so check the rendering itself, for both marshaler flavours.
func TestSlogBackendMarshalerOutput(t *testing.T) {
	var o outputs
	b := newBoth(t, &logger.Config{JSON: true, Level: "debug"}, &o)
	b.log(func(l logger.Logger) {
		l.Infow("m", "p", logger.Proto(participant()), "obj", obj{"o"}, "zobj", zapObj{"z"}, "ids", logger.StringSlice([]string{"x"}))
	})
	for name, out := range map[string]string{"zap": o.z.String(), "slog": o.s.String(), "zap tee": o.zt.String(), "slog tee": o.st.String()} {
		require.Contains(t, out, `"p":{"sid":"PA_abcdef","identity":"user-1","state":"ACTIVE","tracks":[{"sid":"TR_a","type":"AUDIO","name":"<redacted>"`, name)
		require.Contains(t, out, `"obj":{"a":"o","ns":{"n":1}}`, name)
		require.Contains(t, out, `"zobj":{"za":"z","inner":{"a":"z","ns":{"n":1}}}`, name)
		require.Contains(t, out, `"ids":["x"]`, name)
	}
	for name, out := range map[string]string{"zap": o.zh.String(), "slog": o.sh.String()} {
		require.Contains(t, out, `"p":{"sid":"PA_abcdef","identity":"user-1","state":"ACTIVE","tracks":[{`, name)
		require.Contains(t, out, `"obj":{"a":"o","ns":{"n":1}}`, name)
		require.Contains(t, out, `"zobj":{"za":"z","inner":{"a":"z","ns":{"n":1}}}`, name)
		require.Contains(t, out, `"ids":["x"]`, name)
	}
}

func TestFromSlogForeignHandler(t *testing.T) {
	var buf bytes.Buffer
	opts := &slog.HandlerOptions{Level: slog.LevelDebug, ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{}
		}
		return a
	}}
	s := slog.New(slog.NewJSONHandler(&buf, opts))
	l := logger.FromSlog(s).WithName("lk").WithValues("room", "RM_1")
	require.Equal(t, l, logger.FromSlog(logger.ToSlog(l)))

	l.Infow("values", "p", logger.Proto(participant()), "err", verboseErr{"v"}, "ss", []string{"a"},
		"f32", float32(0.1), "zobj", zapObj{"z"}, zap.Namespace("ns"), "inner", 1)
	d, r := l.WithDeferredValues()
	d.Infow("buffered", "k", 1)
	r.Resolve("pID", "PA_1")
	s.Info("direct slog", "p", logger.Proto(participant().Tracks[0]))

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	require.Len(t, lines, 3)
	// Objects inside arrays become maps for foreign handlers, so their keys are sorted.
	require.Contains(t, lines[0], `{"level":"INFO","msg":"values","room":"RM_1","logger":"lk","p":{"sid":"PA_abcdef","identity":"user-1","state":"ACTIVE","tracks":[{"backupCodecPolicy":"PREFER_REGRESSION"`)
	require.Contains(t, lines[0], `"joinedAt":1795000000,"name":"<redacted>"`)
	require.Contains(t, lines[0], `"err":"v","errVerbose":"v\nstack:\n\tmain.go:1","ss":["a"],"f32":0.1,"zobj":{"za":"z","inner":{"a":"z","ns":{"n":1}}},"ns":{"inner":1}}`)
	require.Equal(t, `{"level":"INFO","msg":"buffered","room":"RM_1","pID":"PA_1","logger":"lk","k":1}`, lines[1])
	require.Contains(t, lines[2], `"msg":"direct slog","p":{"sid":"TR_a","type":"AUDIO","name":"<redacted>"`)
}

func logDepthHelper(l *slog.Logger) {
	logger.LogDepth(context.Background(), l, 1, slog.LevelInfo, "depth", "k", 1)
}

func TestLogDepth(t *testing.T) {
	var o outputs
	b := newBoth(t, &logger.Config{JSON: true, Level: "debug"}, &o)
	_, _, line, _ := runtime.Caller(0)
	logDepthHelper(logger.ToSlog(b.s))
	require.Contains(t, o.s.String(), fmt.Sprintf(`"caller":"logger/slog_compat_test.go:%d","msg":"depth","k":1`, line+1))
}

func TestBackendOptions(t *testing.T) {
	_, err := logger.FromZapLogger(zap.NewNop(), &logger.Config{}, logger.WithSlogBackend(), logger.WithZapBackend())
	require.NoError(t, err)
	require.NotNil(t, logger.FromSlog(nil))
}
