package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/livekit/protocol/logger/testutil"
	"github.com/livekit/protocol/logger/zaputil"
	"github.com/livekit/protocol/utils/must"
)

func zapLoggerCore(l Logger) zapcore.Core {
	return l.(ZapLogger).ToZap().Desugar().Core()
}

func jsonTee(ws zapcore.WriteSyncer) ZapLoggerOption {
	return WithTee(zaputil.NewTee(testutil.NewJSONCoreFactory(ws)))
}

func observerTee() (ZapLoggerOption, *observer.ObservedLogs) {
	f, logs := testutil.NewObserverCoreFactory()
	return WithTee(zaputil.NewTee(f)), logs
}

func TestLoggerComponent(t *testing.T) {
	t.Run("inheriting parent level", func(t *testing.T) {
		l, err := NewZapLogger(&Config{
			Level: "info",
			ComponentLevels: map[string]string{
				"mycomponent": "warn",
			},
		})
		require.NoError(t, err)

		sub := zapLoggerCore(l.WithComponent("sub"))
		require.True(t, sub.Enabled(zapcore.InfoLevel))
		require.False(t, sub.Enabled(zapcore.DebugLevel))

		compLogger := zapLoggerCore(l.WithComponent("mycomponent").WithComponent("level2"))
		require.True(t, compLogger.Enabled(zapcore.WarnLevel))
		require.False(t, compLogger.Enabled(zapcore.InfoLevel))
	})

	t.Run("obeys component override", func(t *testing.T) {
		l, err := NewZapLogger(&Config{
			Level: "info",
			ComponentLevels: map[string]string{
				"sub":  "debug",
				"sub2": "error",
			},
		})
		require.NoError(t, err)

		sub := zapLoggerCore(l.WithComponent("sub"))
		sub2 := zapLoggerCore(l.WithComponent("sub2"))
		require.True(t, sub.Enabled(zapcore.DebugLevel))
		require.False(t, sub2.Enabled(zapcore.InfoLevel))
	})

	t.Run("updates dynamically", func(t *testing.T) {
		config := &Config{
			Level: "info",
			ComponentLevels: map[string]string{
				"sub":  "debug",
				"sub2": "error",
			},
		}
		l, err := NewZapLogger(config)
		require.NoError(t, err)

		sub := zapLoggerCore(l.WithComponent("sub"))
		sub2 := zapLoggerCore(l.WithComponent("sub2.test"))
		err = config.Update(&Config{
			Level: "debug",
			ComponentLevels: map[string]string{
				"sub": "info",
				// sub2 removed
			},
		})
		require.NoError(t, err)

		require.True(t, zapLoggerCore(l).Enabled(zapcore.DebugLevel))
		require.False(t, sub.Enabled(zapcore.DebugLevel))
		require.True(t, sub.Enabled(zapcore.InfoLevel))
		require.True(t, sub2.Enabled(zapcore.InfoLevel))
	})

	t.Run("log output matches expected values", func(t *testing.T) {
		silenceStderr(t)
		ws := &testutil.BufferedWriteSyncer{}
		l, err := NewZapLogger(&Config{Level: "debug"}, jsonTee(ws))
		require.NoError(t, err)
		l.Debugw("foo", "bar", "baz")

		var log TestLogOutput
		require.NoError(t, ws.Unmarshal(&log))

		require.Equal(t, "debug", log.Level)
		require.NotEqual(t, 0, log.TS)
		require.NotEqual(t, "", log.Caller)
		require.Equal(t, "foo", log.Msg)
		require.Equal(t, "baz", log.Bar)
	})

	t.Run("component enabler ignores the tee", func(t *testing.T) {
		tee, _ := observerTee()
		l, err := NewZapLogger(&Config{Level: "info"}, tee)
		require.NoError(t, err)

		lvl := l.ComponentLeveler().ComponentLevel("foo")

		require.False(t, lvl.Enabled(zapcore.DebugLevel))
		require.True(t, lvl.Enabled(zapcore.InfoLevel))
	})
}

func TestLoggerTee(t *testing.T) {
	t.Run("receives WithValues fields per derived logger", func(t *testing.T) {
		silenceStderr(t)
		tee, logs := observerTee()
		root := must.Get(NewZapLogger(&Config{Level: "debug"}, tee))

		room := root.WithValues("room", "RM_1")
		room.WithValues("participant", "PA_1").Debugw("participant")
		room.WithValues("track", "TR_1").Debugw("track")

		entries := logs.All()
		require.Len(t, entries, 2)
		require.Equal(t, map[string]any{"room": "RM_1", "participant": "PA_1"}, entries[0].ContextMap())
		require.Equal(t, map[string]any{"room": "RM_1", "track": "TR_1"}, entries[1].ContextMap())
	})

	t.Run("drops entries below the configured level", func(t *testing.T) {
		tee, logs := observerTee()
		l := must.Get(NewZapLogger(&Config{Level: "info"}, tee))

		l.Debugw("debug")

		require.Empty(t, logs.All())
	})

	t.Run("follows component levels", func(t *testing.T) {
		silenceStderr(t)
		tee, logs := observerTee()
		l := must.Get(NewZapLogger(&Config{
			Level: "info",
			ComponentLevels: map[string]string{
				"x":   "debug",
				"x.y": "info",
			},
		}, tee))

		x := l.WithComponent("x")
		x.Debugw("x")
		x.WithComponent("y").Debugw("xy")

		entries := logs.All()
		require.Len(t, entries, 1)
		require.Equal(t, "x", entries[0].Message)
	})

	t.Run("follows the min level floor", func(t *testing.T) {
		silenceStderr(t)
		tee, logs := observerTee()
		l := must.Get(NewZapLogger(&Config{Level: "info"}, tee))

		l.Debugw("dropped")
		l.WithMinLevel(zapcore.DebugLevel).Debugw("kept")

		entries := logs.All()
		require.Len(t, entries, 1)
		require.Equal(t, "kept", entries[0].Message)
	})

	t.Run("receives resolved deferred values", func(t *testing.T) {
		silenceStderr(t)
		tee, logs := observerTee()
		root := must.Get(NewZapLogger(&Config{Level: "debug"}, tee))

		l, resolver := root.WithDeferredValues()
		l.Debugw("deferred")
		require.Empty(t, logs.All())

		resolver.Resolve("participant", "PA_1")

		entries := logs.All()
		require.Len(t, entries, 1)
		require.Equal(t, map[string]any{"participant": "PA_1"}, entries[0].ContextMap())
	})

	t.Run("deferred entries honor the resolved level", func(t *testing.T) {
		silenceStderr(t)
		tee, logs := observerTee()
		root := must.Get(NewZapLogger(&Config{Level: "warn"}, tee))

		l, resolver := root.WithDeferredValues()
		l.Debugw("below")
		l.Warnw("at", nil)
		resolver.Resolve("participant", "PA_1")

		entries := logs.All()
		require.Len(t, entries, 1)
		require.Equal(t, "at", entries[0].Message)
	})

	t.Run("deferred entries run the tee's check-time logic", func(t *testing.T) {
		silenceStderr(t)
		core, logs := observer.New(zapcore.DebugLevel)
		var hooked int
		tee := zaputil.NewTee(func(enab zapcore.LevelEnabler) zapcore.Core {
			return zapcore.RegisterHooks(testutil.Leveled(core, enab), func(zapcore.Entry) error {
				hooked++
				return nil
			})
		})
		root := must.Get(NewZapLogger(&Config{Level: "debug"}, WithTee(tee)))

		l, resolver := root.WithDeferredValues()
		l.Debugw("deferred")
		resolver.Resolve("participant", "PA_1")

		require.Len(t, logs.All(), 1)
		require.Equal(t, 1, hooked)
	})

	t.Run("agrees with the console on malformed value lists", func(t *testing.T) {
		readStderr := captureStderr(t)
		tee, logs := observerTee()
		l := must.Get(NewZapLogger(&Config{JSON: true, Level: "debug"}, tee))

		l.WithValues("room", "RM_1", 42, "dropped", "track", "TR_1", "dangling").Debugw("test")

		var console map[string]any
		require.NoError(t, json.Unmarshal([]byte(readStderr()), &console))
		for _, k := range []string{"level", "ts", "caller", "msg"} {
			delete(console, k)
		}

		require.Equal(t, map[string]any{"room": "RM_1", "track": "TR_1"}, console)
		require.Equal(t, console, logs.All()[0].ContextMap())
	})
}

// The core captures os.Stderr when it is built, so these must run before the logger is created.

func silenceStderr(tb testing.TB) {
	tb.Helper()
	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	require.NoError(tb, err)
	prev := os.Stderr
	os.Stderr = f
	tb.Cleanup(func() {
		os.Stderr = prev
		f.Close()
	})
}

func captureStderr(tb testing.TB) func() string {
	tb.Helper()
	f, err := os.CreateTemp(tb.TempDir(), "stderr")
	require.NoError(tb, err)
	prev := os.Stderr
	os.Stderr = f
	tb.Cleanup(func() {
		os.Stderr = prev
		f.Close()
	})
	return func() string {
		b, err := os.ReadFile(f.Name())
		require.NoError(tb, err)
		return string(b)
	}
}

type TestLogOutput struct {
	testutil.TestLogOutput
	Bar string
}

type logFunc func(string, ...any)

func testLogCaller(f logFunc) {
	f("test")
}

func TestLoggerCallDepth(t *testing.T) {
	t.Cleanup(func() {
		defaultLogger = LogRLogger(discardLogger)
		pkgLogger = LogRLogger(discardLogger)
	})

	var caller string
	testLogCaller(func(string, ...any) {
		_, file, line, _ := runtime.Caller(1)
		caller = fmt.Sprintf("%s:%d", file, line)
	})

	cases := map[string]func(l Logger) logFunc{
		"NewZapLogger": func(l Logger) logFunc {
			return l.Debugw
		},
		"package logger": func(l Logger) logFunc {
			SetLogger(l, "TEST")
			return Debugw
		},
		"GetLogger": func(l Logger) logFunc {
			SetLogger(l, "TEST")
			return GetLogger().Debugw
		},
		"ToZap": func(l Logger) logFunc {
			return l.(ZapLogger).ToZap().Debugw
		},
		"WithUnlikelyValues": func(l Logger) logFunc {
			return l.WithUnlikelyValues().Debugw
		},
	}
	for label, getLogFunc := range cases {
		t.Run(label, func(t *testing.T) {
			ws := &testutil.BufferedWriteSyncer{}
			l := must.Get(NewZapLogger(&Config{Level: "debug"}, jsonTee(ws)))

			testLogCaller(getLogFunc(l))

			var log TestLogOutput
			require.NoError(t, ws.Unmarshal(&log))
			require.True(t, strings.HasSuffix(caller, log.Caller), `caller mismatch expected suffix match on "%s" got "%s"`, caller, log.Caller)
		})
	}
}
