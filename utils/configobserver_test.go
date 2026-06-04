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

package utils

import (
	"os"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils/configutil"
)

const testConfig0 = `foo: a`
const testConfig1 = `foo: b`
const testConfigInvalid = `foo: [unterminated`

type TestConfig struct {
	Foo string `yaml:"foo"`
	Bar string `yaml:"bar"`
}

type testConfigBuilder struct{}

func (testConfigBuilder) New() (*TestConfig, error) {
	return &TestConfig{}, nil
}

func (testConfigBuilder) InitDefaults(c *TestConfig) error {
	c.Bar = "c"
	return nil
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
	t.Cleanup(obs.Close)

	require.Equal(t, "a", conf.Foo)
	require.Equal(t, "c", conf.Bar)

	atomicFoo := configutil.NewAtomicValue(obs, func(c *TestConfig) string {
		return c.Foo
	})

	require.Equal(t, "a", atomicFoo.Load())

	// the initial load publishes the config hash but does not count as a reload
	require.Zero(t, counterVecValue(promConfigReloadTotal, f.Name(), "success"))
	require.Equal(t,
		float64(configHash([]byte(testConfig0))),
		gaugeVecValue(promConfigHash, f.Name()),
	)

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

	require.Equal(t, "b", atomicFoo.Load())

	// the reload is counted and the hash gauge tracks the new config
	require.Equal(t, float64(1), counterVecValue(promConfigReloadTotal, f.Name(), "success"))
	require.Equal(t,
		float64(configHash([]byte(testConfig1))),
		gaugeVecValue(promConfigHash, f.Name()),
	)

	_, err = f.WriteAt([]byte(testConfigInvalid), 0)
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		return counterVecValue(promConfigReloadTotal, f.Name(), "failure") == 1
	}, time.Second, 5*time.Millisecond)
	require.Equal(t,
		float64(configHash([]byte(testConfig1))),
		gaugeVecValue(promConfigHash, f.Name()),
	)
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
