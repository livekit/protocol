package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger/zaputil"
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

		var log struct {
			Level  string
			TS     float64
			Caller string
			Msg    string
			Bar    string
		}
		require.NoError(t, json.Unmarshal(ws.Bytes(), &log))

		require.Equal(t, "debug", log.Level)
		require.NotEqual(t, 0, log.TS)
		require.NotEqual(t, "", log.Caller)
		require.Equal(t, "foo", log.Msg)
		require.Equal(t, "baz", log.Bar)
	})
}

type testBufferedWriteSyncer struct {
	bytes.Buffer
}

func (t *testBufferedWriteSyncer) Sync() error { return nil }
