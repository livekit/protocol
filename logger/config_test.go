package logger

import (
	"io"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"

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

func TestConfigYAMLFields(t *testing.T) {
	conf, snap := reflect.TypeFor[Config](), reflect.TypeFor[configYAML]()

	var i int
	for _, f := range reflect.VisibleFields(conf) {
		if !f.IsExported() {
			continue
		}
		require.Less(t, i, snap.NumField(), "configYAML is missing %s", f.Name)

		s := snap.Field(i)
		require.Equal(t, f.Name, s.Name)
		require.Equal(t, f.Type, s.Type)
		require.Equal(t, f.Tag.Get("yaml"), s.Tag.Get("yaml"))
		i++
	}
	require.Equal(t, snap.NumField(), i, "configYAML has fields Config does not")
}

func TestConfigMarshalYAML(t *testing.T) {
	conf := &Config{
		JSON:            true,
		Level:           "debug",
		Sample:          true,
		ComponentLevels: map[string]string{"rtc.room": "debug"},
		SampleInitial:   5,
	}

	b, err := yaml.Marshal(conf)
	require.NoError(t, err)
	require.Equal(t, `json: true
level: debug
sample: true
component_levels:
    rtc.room: debug
sample_initial: 5
`, string(b))

	var out Config
	require.NoError(t, yaml.Unmarshal(b, &out))
	require.Equal(t, conf.ComponentLevels, out.ComponentLevels)
	require.Equal(t, conf.Level, out.Level)
	require.Equal(t, conf.SampleInitial, out.SampleInitial)

	// The encoder walks the snapshot after MarshalYAML returns, so a component level
	// dropped in the meantime must not change what was marshaled.
	delete(conf.ComponentLevels, "rtc.room")
	require.Contains(t, string(b), "rtc.room: debug")

	b, err = yaml.Marshal(&Config{})
	require.NoError(t, err)
	require.Equal(t, "{}\n", string(b))
}
