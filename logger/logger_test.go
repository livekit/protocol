package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestLoggerComponent(t *testing.T) {
	t.Run("inheriting parent level", func(t *testing.T) {
		l, err := NewZapLogger(&Config{
			Level: "info",
			ComponentLevels: map[string]string{
				"mycomponent": "warn",
			},
		})
		require.NoError(t, err)

		sub := l.WithComponent("sub")
		require.True(t, sub.(*ZapLogger).isEnabled(zapcore.InfoLevel))
		require.False(t, sub.(*ZapLogger).isEnabled(zapcore.DebugLevel))

		compLogger := l.WithComponent("mycomponent").WithComponent("level2")
		require.True(t, compLogger.(*ZapLogger).isEnabled(zapcore.WarnLevel))
		require.False(t, compLogger.(*ZapLogger).isEnabled(zapcore.InfoLevel))
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

		sub := l.WithComponent("sub")
		sub2 := l.WithComponent("sub2")
		require.True(t, sub.(*ZapLogger).isEnabled(zapcore.DebugLevel))
		require.False(t, sub2.(*ZapLogger).isEnabled(zapcore.InfoLevel))
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

		sub := l.WithComponent("sub")
		sub2 := l.WithComponent("sub2.test")
		err = config.Update(&Config{
			Level: "debug",
			ComponentLevels: map[string]string{
				"sub": "info",
				// sub2 removed
			},
		})
		require.NoError(t, err)

		require.True(t, l.isEnabled(zapcore.DebugLevel))
		require.False(t, sub.(*ZapLogger).isEnabled(zapcore.DebugLevel))
		require.True(t, sub.(*ZapLogger).isEnabled(zapcore.InfoLevel))
		require.True(t, sub2.(*ZapLogger).isEnabled(zapcore.InfoLevel))
	})
}
