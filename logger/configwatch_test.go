package logger

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	require.NoError(t, os.WriteFile(path, []byte(body), 0o600))
}

func TestApplyConfigFile(t *testing.T) {
	t.Run("a level change reaches a live logger", func(t *testing.T) {
		conf := &Config{Level: "info"}
		l, err := NewZapLogger(conf)
		require.NoError(t, err)
		core := zapLoggerCore(l)
		require.False(t, core.Enabled(zapcore.DebugLevel))

		path := filepath.Join(t.TempDir(), "logging.yaml")
		writeFile(t, path, "level: debug\n")

		applied, ok := applyConfigFile(conf, conf.snapshot(), path, nil)
		require.True(t, ok)
		require.NotEmpty(t, applied)
		require.True(t, zapLoggerCore(l).Enabled(zapcore.DebugLevel),
			"the atomic level behind the existing logger must move, not just Config.Level")
	})

	t.Run("keys absent from the file keep their current values", func(t *testing.T) {
		// Update assigns every field, so applying a partial file over a zero Config would wipe
		// these. component_levels is where livekit-server lands pion_level.
		conf := &Config{
			Level:           "info",
			JSON:            true,
			Sample:          true,
			SampleInitial:   7,
			ComponentLevels: map[string]string{"pion": "error"},
		}
		_, err := NewZapLogger(conf)
		require.NoError(t, err)

		path := filepath.Join(t.TempDir(), "logging.yaml")
		writeFile(t, path, "level: warn\n")

		_, ok := applyConfigFile(conf, conf.snapshot(), path, nil)
		require.True(t, ok)
		require.Equal(t, "warn", conf.Level)
		require.True(t, conf.JSON)
		require.True(t, conf.Sample)
		require.Equal(t, 7, conf.SampleInitial)
		require.Equal(t, map[string]string{"pion": "error"}, conf.ComponentLevels)
	})

	t.Run("a component level in the file merges with the existing ones", func(t *testing.T) {
		conf := &Config{Level: "info", ComponentLevels: map[string]string{"pion": "error"}}
		l, err := NewZapLogger(conf)
		require.NoError(t, err)

		path := filepath.Join(t.TempDir(), "logging.yaml")
		writeFile(t, path, "component_levels:\n  psrpc: debug\n")

		_, ok := applyConfigFile(conf, conf.snapshot(), path, nil)
		require.True(t, ok)
		require.Equal(t, "error", conf.ComponentLevels["pion"])
		require.Equal(t, "debug", conf.ComponentLevels["psrpc"])
		require.True(t, zapLoggerCore(l.WithComponent("psrpc")).Enabled(zapcore.DebugLevel))
	})

	t.Run("unchanged bytes are not reapplied", func(t *testing.T) {
		conf := &Config{Level: "info"}
		path := filepath.Join(t.TempDir(), "logging.yaml")
		writeFile(t, path, "level: debug\n")

		applied, ok := applyConfigFile(conf, conf.snapshot(), path, nil)
		require.True(t, ok)
		_, ok = applyConfigFile(conf, conf.snapshot(), path, applied)
		require.False(t, ok)
	})

	t.Run("malformed yaml keeps the last good config", func(t *testing.T) {
		conf := &Config{Level: "info"}
		l, err := NewZapLogger(conf)
		require.NoError(t, err)

		path := filepath.Join(t.TempDir(), "logging.yaml")
		writeFile(t, path, "level: [not, a, string\n")

		_, ok := applyConfigFile(conf, conf.snapshot(), path, nil)
		require.False(t, ok)
		require.Equal(t, "info", conf.Level)
		require.False(t, zapLoggerCore(l).Enabled(zapcore.DebugLevel))
	})

	t.Run("a missing file is tolerated", func(t *testing.T) {
		conf := &Config{Level: "info"}
		_, ok := applyConfigFile(conf, conf.snapshot(), filepath.Join(t.TempDir(), "absent.yaml"), nil)
		require.False(t, ok)
		require.Equal(t, "info", conf.Level)
	})
}

func TestWatchConfigFile(t *testing.T) {
	conf := &Config{Level: "info"}
	l, err := NewZapLogger(conf)
	require.NoError(t, err)

	path := filepath.Join(t.TempDir(), "logging.yaml")
	writeFile(t, path, "level: info\n")

	stop := WatchConfigFile(conf, path, 5*time.Millisecond)
	t.Cleanup(stop)

	writeFile(t, path, "level: debug\n")
	require.Eventually(t, func() bool {
		return zapLoggerCore(l).Enabled(zapcore.DebugLevel)
	}, 2*time.Second, 5*time.Millisecond, "watcher should pick up the rewritten file")

	stop()
	writeFile(t, path, "level: error\n")
	time.Sleep(50 * time.Millisecond)
	require.True(t, zapLoggerCore(l).Enabled(zapcore.DebugLevel), "stop must end the polling")
}

// Applying config while component levels are being resolved: sharedConfig.ComponentLevel reads
// under its own mutex while Update writes the Config under a different one, so it must be reading
// a copy it owns. Meaningful under -race.
func TestApplyConfigFileWhileResolvingComponents(t *testing.T) {
	conf := &Config{Level: "info", ComponentLevels: map[string]string{"pion": "error"}}
	l, err := NewZapLogger(conf)
	require.NoError(t, err)

	dir := t.TempDir()
	path := filepath.Join(dir, "logging.yaml")

	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 500; i++ {
			_ = zapLoggerCore(l.WithComponent("psrpc").WithComponent("Egress"))
		}
	}()

	var last []byte
	for i, level := range []string{"debug", "warn", "info", "error"} {
		writeFile(t, path, "level: "+level+"\n")
		applied, ok := applyConfigFile(conf, conf.snapshot(), path, last)
		require.True(t, ok, "iteration %d", i)
		last = applied
	}
	<-done
	require.Equal(t, "error", conf.Level)
	require.Equal(t, "error", conf.ComponentLevels["pion"])
}

// The reset path the chart documents: emptying the file must put the startup levels back, not
// leave the last override in force. Applying each file over a baseline rather than over the
// config currently in force is what makes this hold.
func TestEmptyFileRestoresStartupConfig(t *testing.T) {
	conf := &Config{Level: "info", ComponentLevels: map[string]string{"pion": "error"}}
	l, err := NewZapLogger(conf)
	require.NoError(t, err)
	baseline := conf.snapshot()

	path := filepath.Join(t.TempDir(), "logging.yaml")
	writeFile(t, path, "level: debug\ncomponent_levels:\n  psrpc: debug\n")
	applied, ok := applyConfigFile(conf, baseline, path, nil)
	require.True(t, ok)
	require.Equal(t, "debug", conf.Level)
	require.True(t, zapLoggerCore(l).Enabled(zapcore.DebugLevel))

	writeFile(t, path, "{}\n")
	_, ok = applyConfigFile(conf, baseline, path, applied)
	require.True(t, ok)
	require.Equal(t, "info", conf.Level)
	require.Equal(t, map[string]string{"pion": "error"}, conf.ComponentLevels,
		"a component the file no longer names must stop applying")
	require.False(t, zapLoggerCore(l).Enabled(zapcore.DebugLevel))
}
