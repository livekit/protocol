package logger

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger/zaputil"
)

func TestConfigResolveComponentLevel(t *testing.T) {
	conf := &Config{Level: "info", ComponentLevels: map[string]string{"rtc.room": "debug"}}
	root := zaputil.NewRootComponentLeveler(zapcore.AddSync(io.Discard), conf)

	require.True(t, root.ComponentLevel("rtc.room").Enabled(zapcore.DebugLevel))
	require.True(t, root.ComponentLevel("rtc.room.track").Enabled(zapcore.DebugLevel))
	require.False(t, root.ComponentLevel("rtc").Enabled(zapcore.DebugLevel))
	require.True(t, root.ComponentLevel("rtc").Enabled(zapcore.InfoLevel))

	lvl, ok := (&Config{}).ResolveComponentLevel("anything")
	require.True(t, ok)
	require.Equal(t, zapcore.InfoLevel, lvl)
}
