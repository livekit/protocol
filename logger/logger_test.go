package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger/zaputil"
	"github.com/livekit/protocol/utils/must"
)

func zapLoggerCore(l Logger) zapcore.Core {
	return l.(ZapLogger).ToZap().Desugar().Core()
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
		ws := &testBufferedWriteSyncer{}
		l, err := NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel)))
		require.NoError(t, err)
		l.Debugw("foo", "bar", "baz")

		log, err := unmarshalTestLogOutput(ws.Bytes())
		require.NoError(t, err)

		require.Equal(t, "debug", log.Level)
		require.NotEqual(t, 0, log.TS)
		require.NotEqual(t, "", log.Caller)
		require.Equal(t, "foo", log.Msg)
		require.Equal(t, "baz", log.Bar)
	})

	t.Run("component enabler for tapped logger returns lowest enabled level", func(t *testing.T) {
		tapLevel := zap.NewAtomicLevel()
		l, err := NewZapLogger(&Config{Level: "info"}, WithTap(zaputil.NewWriteEnabler(&testBufferedWriteSyncer{}, tapLevel)))
		require.NoError(t, err)

		lvl := l.ComponentLeveler().ComponentLevel("foo")

		// check config level
		require.False(t, lvl.Enabled(zapcore.DebugLevel))
		require.True(t, lvl.Enabled(zapcore.InfoLevel))

		// check tap level
		tapLevel.SetLevel(zapcore.DebugLevel)
		require.True(t, lvl.Enabled(zapcore.DebugLevel))
	})
}

type testLogOutput struct {
	Level  string
	TS     float64
	Caller string
	Msg    string
	Bar    string
}

func unmarshalTestLogOutput(b []byte) (*testLogOutput, error) {
	log := &testLogOutput{}
	return log, json.Unmarshal(b, &log)
}

type testBufferedWriteSyncer struct {
	bytes.Buffer
}

func (t *testBufferedWriteSyncer) Sync() error { return nil }

func testLogCaller(logFunc func(msg string, keysAndValues ...any)) {
	logFunc("test")
}

func getTestLogCallerCaller() string {
	var caller string
	testLogCaller(func(string, ...any) {
		_, file, line, _ := runtime.Caller(1)
		caller = fmt.Sprintf("%s:%d", file, line)
	})
	return caller
}

func TestLoggerCallDepth(t *testing.T) {
	caller := getTestLogCallerCaller()

	t.Run("NewZapLogger", func(t *testing.T) {
		ws := &testBufferedWriteSyncer{}
		l := must.Get(NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel))))

		testLogCaller(l.Debugw)
		log := must.Get(unmarshalTestLogOutput(ws.Bytes()))

		require.True(t, strings.HasSuffix(caller, log.Caller), `caller mismatch expected suffix match on "%s" got "%s"`, caller, log.Caller)
	})

	t.Run("package logger", func(t *testing.T) {
		ws := &testBufferedWriteSyncer{}
		l := must.Get(NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel))))
		SetLogger(l, "TEST")

		testLogCaller(Debugw)
		log := must.Get(unmarshalTestLogOutput(ws.Bytes()))

		require.True(t, strings.HasSuffix(caller, log.Caller), `caller mismatch expected suffix match on "%s" got "%s"`, caller, log.Caller)
	})

	t.Run("GetLogger", func(t *testing.T) {
		ws := &testBufferedWriteSyncer{}
		l := must.Get(NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel))))
		SetLogger(l, "TEST")

		testLogCaller(GetLogger().Debugw)
		log := must.Get(unmarshalTestLogOutput(ws.Bytes()))

		require.True(t, strings.HasSuffix(caller, log.Caller), `caller mismatch expected suffix match on "%s" got "%s"`, caller, log.Caller)
	})

	t.Run("ToZap", func(t *testing.T) {
		ws := &testBufferedWriteSyncer{}
		l := must.Get(NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel))))

		testLogCaller(l.ToZap().Debugw)
		log := must.Get(unmarshalTestLogOutput(ws.Bytes()))

		require.True(t, strings.HasSuffix(caller, log.Caller), `caller mismatch expected suffix match on "%s" got "%s"`, caller, log.Caller)
	})

	t.Run("WithUnlikelyValues", func(t *testing.T) {
		ws := &testBufferedWriteSyncer{}
		l := must.Get(NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel))))

		testLogCaller(l.WithUnlikelyValues().Debugw)
		log := must.Get(unmarshalTestLogOutput(ws.Bytes()))

		require.True(t, strings.HasSuffix(caller, log.Caller), `caller mismatch expected suffix match on "%s" got "%s"`, caller, log.Caller)
	})
}
