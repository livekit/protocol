package configutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils/events"
)

type testConfig struct {
	events.ObserverList[*testConfig]
}

func (c *testConfig) Observe(cb func(*testConfig)) func() {
	return c.On(cb)
}

func (c *testConfig) ObserveCurrent(cb func(*testConfig)) func() {
	stop := c.On(cb)
	cb(c)
	return stop
}

func (c *testConfig) Load() *testConfig {
	return c
}

func TestInterfaceAtomicValue(t *testing.T) {
	config := &testConfig{}
	_ = NewAtomicValue(config, func(*testConfig) any {
		return 1
	})
	config.Emit(&testConfig{})
}

func TestAtomicAfterEmit(t *testing.T) {
	obs := NewStaticObserver(&TestConfig{Foo: "a"})
	obs.EmitConfigUpdate(&TestConfig{Foo: "b"})

	foo := NewAtomicString(obs, func(c *TestConfig) string { return c.Foo })
	require.Equal(t, "b", foo.Load())

	obs.EmitConfigUpdate(&TestConfig{Foo: "c"})
	require.Equal(t, "c", foo.Load())
}
