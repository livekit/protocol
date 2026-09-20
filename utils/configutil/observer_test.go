// Copyright 2024 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configutil

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
)

const testConfig0 = `foo: a`
const testConfig1 = `foo: b`
const testConfigInvalid = `foo: [unterminated`
const testConfig1Recover = `foo: b # recovered after invalid`

type TestConfig struct {
	Foo string `yaml:"foo"`
	Bar string `yaml:"bar"`
}

var (
	_ Defaulter[TestConfig] = testConfigBuilder{}
	_ Validator[TestConfig] = testConfigBuilder{}
)

type testConfigBuilder struct {
	initErr  error
	validate func(*TestConfig) error
}

func (testConfigBuilder) New() (*TestConfig, error) {
	return &TestConfig{}, nil
}

func (b testConfigBuilder) InitDefaults(c *TestConfig) error {
	if b.initErr != nil {
		return b.initErr
	}
	c.Bar = "c"
	return nil
}

func (b testConfigBuilder) Validate(c *TestConfig) error {
	if b.validate == nil {
		return nil
	}
	return b.validate(c)
}

type newOnlyBuilder struct{}

func (newOnlyBuilder) New() (*TestConfig, error) {
	return &TestConfig{Foo: "new"}, nil
}

func TestConfigObserver(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "lk-test-*.yaml")
	t.Cleanup(func() {
		_ = f.Close()
	})
	require.NoError(t, err)
	_, err = f.WriteString(testConfig0)
	require.NoError(t, err)

	obs, conf, err := NewObserver(f.Name(), testConfigBuilder{})
	require.NoError(t, err)
	t.Cleanup(obs.Close)

	require.Equal(t, "a", conf.Foo)
	require.Equal(t, "c", conf.Bar)

	atomicFoo := NewAtomicValue(obs, func(c *TestConfig) string {
		return c.Foo
	})

	require.Equal(t, "a", atomicFoo.Load())

	// the initial load publishes the config hash but does not count as a reload,
	// and the load state is populated as healthy (0) from the start
	require.Zero(t, counterVecValue(promConfigReloadTotal, f.Name(), "success"))
	require.Equal(t,
		float64(configHash([]byte(testConfig0))),
		gaugeVecValue(promConfigHash, f.Name()),
	)
	require.Zero(t, gaugeVecValue(promConfigLoadState, f.Name()))

	done := make(chan struct{})
	unsubscribe := obs.Observe(func(c *TestConfig) {
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

	require.Equal(t, "b", atomicFoo.Load())

	// the reload is counted, the hash gauge tracks the new config, and the load
	// state stays healthy (0)
	require.Equal(t, float64(1), counterVecValue(promConfigReloadTotal, f.Name(), "success"))
	require.Equal(t,
		float64(configHash([]byte(testConfig1))),
		gaugeVecValue(promConfigHash, f.Name()),
	)
	require.Zero(t, gaugeVecValue(promConfigLoadState, f.Name()))

	// stop the one-shot observer above; the remaining writes don't expect it
	unsubscribe()

	_, err = f.WriteAt([]byte(testConfigInvalid), 0)
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		return counterVecValue(promConfigReloadTotal, f.Name(), "failure") == 1
	}, time.Second, 5*time.Millisecond)
	require.Equal(t,
		float64(configHash([]byte(testConfig1))),
		gaugeVecValue(promConfigHash, f.Name()),
	)
	require.Equal(t, float64(1), gaugeVecValue(promConfigLoadState, f.Name()))

	// a subsequent valid load clears the failure state back to 0
	_, err = f.WriteAt([]byte(testConfig1Recover), 0)
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		return counterVecValue(promConfigReloadTotal, f.Name(), "success") == 2
	}, time.Second, 5*time.Millisecond)
	require.Zero(t, gaugeVecValue(promConfigLoadState, f.Name()))
}

const testConfigRejected = `foo: x`

var errRejected = errors.New("rejected")

func rejectFooX(c *TestConfig) error {
	if c.Foo == "x" {
		return errRejected
	}
	return nil
}

func TestConfigObserverValidation(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "lk-test-*.yaml")
	require.NoError(t, err)
	t.Cleanup(func() { _ = f.Close() })
	_, err = f.WriteString(testConfig0)
	require.NoError(t, err)

	obs, conf, err := NewObserver(f.Name(), testConfigBuilder{validate: rejectFooX})
	require.NoError(t, err)
	t.Cleanup(obs.Close)
	require.Equal(t, "a", conf.Foo)

	var emitted atomic.Int32
	obs.Observe(func(*TestConfig) { emitted.Inc() })

	_, err = f.WriteAt([]byte(testConfigRejected), 0)
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		return counterVecValue(promConfigReloadTotal, f.Name(), "failure") == 1
	}, time.Second, 5*time.Millisecond)

	// the rejected config is neither stored nor emitted, and the hash still
	// reflects the config in use
	require.Equal(t, "a", obs.Load().Foo)
	require.Zero(t, emitted.Load())
	require.Equal(t, float64(1), gaugeVecValue(promConfigLoadState, f.Name()))
	require.Equal(t,
		float64(configHash([]byte(testConfig0))),
		gaugeVecValue(promConfigHash, f.Name()),
	)

	_, err = f.WriteAt([]byte(testConfig1), 0)
	require.NoError(t, err)

	require.Eventually(t, func() bool { return emitted.Load() == 1 }, time.Second, 5*time.Millisecond)
	require.Equal(t, "b", obs.Load().Foo)
	require.Zero(t, gaugeVecValue(promConfigLoadState, f.Name()))
}

func TestNewObserverLoadErrors(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "lk-test-*.yaml")
	require.NoError(t, err)
	t.Cleanup(func() { _ = f.Close() })
	_, err = f.WriteString(testConfig0)
	require.NoError(t, err)

	errDefaults := errors.New("defaults")
	rejectAll := func(*TestConfig) error { return errRejected }

	for _, tc := range []struct {
		name    string
		builder testConfigBuilder
		want    error
	}{
		{"validate", testConfigBuilder{validate: rejectAll}, errRejected},
		{"init_defaults", testConfigBuilder{initErr: errDefaults}, errDefaults},
	} {
		t.Run(tc.name+"/file", func(t *testing.T) {
			obs, conf, err := NewObserver(f.Name(), tc.builder)
			require.ErrorIs(t, err, tc.want)
			require.Nil(t, obs)
			require.Nil(t, conf)
			require.Equal(t, float64(1), gaugeVecValue(promConfigLoadState, f.Name()))
		})
		t.Run(tc.name+"/nofile", func(t *testing.T) {
			obs, conf, err := NewObserver("", tc.builder)
			require.ErrorIs(t, err, tc.want)
			require.Nil(t, obs)
			require.Nil(t, conf)
		})
	}
}

func TestNewObserverBuilderOnly(t *testing.T) {
	obs, conf, err := NewObserver("", newOnlyBuilder{})
	require.NoError(t, err)
	t.Cleanup(obs.Close)
	require.Equal(t, &TestConfig{Foo: "new"}, conf)
	require.Same(t, conf, obs.Load())
}

func gaugeVecValue(g *prometheus.GaugeVec, labels ...string) float64 {
	m, err := g.GetMetricWithLabelValues(labels...)
	if err != nil {
		return 0
	}
	var dtoM dto.Metric
	if err := m.Write(&dtoM); err != nil {
		return 0
	}
	return dtoM.GetGauge().GetValue()
}

func counterVecValue(c *prometheus.CounterVec, labels ...string) float64 {
	m, err := c.GetMetricWithLabelValues(labels...)
	if err != nil {
		return 0
	}
	var dtoM dto.Metric
	if err := m.Write(&dtoM); err != nil {
		return 0
	}
	return dtoM.GetCounter().GetValue()
}
