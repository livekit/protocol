package logger

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger/testutil"
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
		ws := &testutil.BufferedWriteSyncer{}
		l, err := NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel)))
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

	t.Run("component enabler for tapped logger returns lowest enabled level", func(t *testing.T) {
		tapLevel := zap.NewAtomicLevel()
		l, err := NewZapLogger(&Config{Level: "info"}, WithTap(zaputil.NewWriteEnabler(&testutil.BufferedWriteSyncer{}, tapLevel)))
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

type TestLogOutput struct {
	testutil.TestLogOutput
	Bar string
}

type auditLogOutput struct {
	testutil.TestLogOutput
	Action      string
	UserID      string
	ProjectID   string
	ServiceName string
	AuditLog    bool
	Error       string
	Room        string
}

func TestAudit(t *testing.T) {
	t.Cleanup(func() {
		defaultLogger = LogRLogger(discardLogger)
		pkgLogger = LogRLogger(discardLogger)
	})

	setup := func() *testutil.BufferedWriteSyncer {
		ws := &testutil.BufferedWriteSyncer{}
		l := must.Get(NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel))))
		SetLogger(l, "TEST")
		return ws
	}

	scope := AuditScope{UserID: "user-456", ProjectID: "proj-123"}

	t.Run("success", func(t *testing.T) {
		ws := setup()
		NewAuditEmitter(GetLogger(), "my-service").Scope(scope).Log("user signed in", "action", "login")

		var log auditLogOutput
		require.NoError(t, ws.Unmarshal(&log))
		require.Equal(t, "info", log.Level)
		require.Equal(t, "user signed in", log.Msg)
		require.Equal(t, "my-service", log.ServiceName)
		require.Equal(t, "proj-123", log.ProjectID)
		require.Equal(t, "user-456", log.UserID)
		require.True(t, log.AuditLog)
		require.Equal(t, "login", log.Action)
		require.Empty(t, log.Error)
		// caller must point at this test file, not logger.go
		require.Contains(t, log.Caller, "logger_test.go")
	})

	t.Run("error", func(t *testing.T) {
		ws := setup()
		NewAuditEmitter(GetLogger(), "my-service").Scope(scope).LogError("user sign-in failed", errors.New("boom"))

		var log auditLogOutput
		require.NoError(t, ws.Unmarshal(&log))
		require.Equal(t, "error", log.Level)
		require.Equal(t, "user sign-in failed", log.Msg)
		require.Equal(t, "my-service", log.ServiceName)
		require.Equal(t, "proj-123", log.ProjectID)
		require.Equal(t, "user-456", log.UserID)
		require.True(t, log.AuditLog)
		require.Equal(t, "boom", log.Error)
	})

	t.Run("direct log with varying scope", func(t *testing.T) {
		ws := setup()
		NewAuditEmitter(GetLogger(), "my-service").Log(scope, "invoice created")

		var log auditLogOutput
		require.NoError(t, ws.Unmarshal(&log))
		require.Equal(t, "my-service", log.ServiceName)
		require.Equal(t, "invoice created", log.Msg)
	})

	t.Run("inherits context bound on the base logger", func(t *testing.T) {
		ws := setup()
		l := GetLogger().WithValues("room", "room-1")
		NewAuditEmitter(l, "billing-service").Scope(scope).Log("invoice created")

		var log auditLogOutput
		require.NoError(t, ws.Unmarshal(&log))
		require.Equal(t, "billing-service", log.ServiceName)
		require.Equal(t, "invoice created", log.Msg)
		require.Equal(t, "room-1", log.Room) // came from the base logger
		require.Equal(t, "user-456", log.UserID)
	})

	t.Run("empty scope and service fall back to unknown", func(t *testing.T) {
		ws := setup()
		NewAuditEmitter(GetLogger(), "").Log(AuditScope{}, "no context")

		var log auditLogOutput
		require.NoError(t, ws.Unmarshal(&log))
		require.Equal(t, "unknown-service", log.ServiceName)
		require.Equal(t, "unknown", log.UserID)
		require.Equal(t, "unknown", log.ProjectID)
	})
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
			l := must.Get(NewZapLogger(&Config{}, WithTap(zaputil.NewWriteEnabler(ws, zapcore.DebugLevel))))

			testLogCaller(getLogFunc(l))

			var log TestLogOutput
			require.NoError(t, ws.Unmarshal(&log))
			require.True(t, strings.HasSuffix(caller, log.Caller), `caller mismatch expected suffix match on "%s" got "%s"`, caller, log.Caller)
		})
	}
}
