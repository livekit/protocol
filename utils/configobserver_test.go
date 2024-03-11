package utils

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const testConfig0 = `foo: a`
const testConfig1 = `foo: b`

type TestConfig struct {
	Foo string `yaml:"foo"`
	Bar string `yaml:"bar"`
}

type testConfigBuilder struct{}

func (testConfigBuilder) New() (*TestConfig, error) {
	return &TestConfig{}, nil
}

func (testConfigBuilder) InitDefaults(c *TestConfig) {
	c.Bar = "c"
}

func TestConfigObserver(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "lk-test-*.yaml")
	t.Cleanup(func() {
		_ = f.Close()
	})
	require.NoError(t, err)
	_, err = f.WriteString(testConfig0)
	require.NoError(t, err)

	obs, conf, err := NewConfigObserver(f.Name(), testConfigBuilder{})
	require.NoError(t, err)

	require.Equal(t, "a", conf.Foo)
	require.Equal(t, "c", conf.Bar)

	done := make(chan struct{})
	obs.Observe(func(c *TestConfig) {
		require.Equal(t, "b", c.Foo)
		require.Equal(t, "c", c.Bar)
		close(done)
	})

	_, err = f.WriteAt([]byte(testConfig1), 0)
	require.NoError(t, err)

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		require.FailNow(t, "timed out waiting for config update")
	}
}
