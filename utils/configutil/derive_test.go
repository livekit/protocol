// Copyright 2026 LiveKit, Inc.
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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// *Observer is the observable Derive exists to narrow, so hold it to the
// interface here rather than discovering a mismatch at a call site.
var _ Observable[deriveAppConfig] = (*Observer[deriveAppConfig])(nil)

type deriveAppConfig struct {
	Sweeper deriveSweeperConfig
	Nested  deriveNestedConfig
}

type deriveSweeperConfig struct {
	Period time.Duration
}

type deriveNestedConfig struct {
	Inner deriveInnerConfig
}

type deriveInnerConfig struct {
	Name string
}

func TestDerive(t *testing.T) {
	src := NewStaticObserver(&deriveAppConfig{
		Sweeper: deriveSweeperConfig{Period: time.Second},
	})
	sweeper := Derive(src, func(c *deriveAppConfig) *deriveSweeperConfig { return &c.Sweeper })

	require.Equal(t, time.Second, sweeper.Load().Period)

	var observed []time.Duration
	unsubscribe := sweeper.Observe(func(c *deriveSweeperConfig) {
		observed = append(observed, c.Period)
	})

	src.EmitConfigUpdate(&deriveAppConfig{Sweeper: deriveSweeperConfig{Period: 2 * time.Second}})
	require.Equal(t, 2*time.Second, sweeper.Load().Period)
	require.Equal(t, []time.Duration{2 * time.Second}, observed)

	unsubscribe()
	src.EmitConfigUpdate(&deriveAppConfig{Sweeper: deriveSweeperConfig{Period: 3 * time.Second}})
	require.Equal(t, 3*time.Second, sweeper.Load().Period)
	require.Equal(t, []time.Duration{2 * time.Second}, observed, "unsubscribed callback still fired")
}

func TestDeriveLoadIdentity(t *testing.T) {
	conf := &deriveAppConfig{Sweeper: deriveSweeperConfig{Period: time.Second}}
	sweeper := Derive(NewStaticObserver(conf), func(c *deriveAppConfig) *deriveSweeperConfig { return &c.Sweeper })

	require.Same(t, &conf.Sweeper, sweeper.Load())
	require.Same(t, sweeper.Load(), sweeper.Load())
}

func TestDeriveChained(t *testing.T) {
	src := NewStaticObserver(&deriveAppConfig{
		Nested: deriveNestedConfig{Inner: deriveInnerConfig{Name: "a"}},
	})
	nested := Derive(src, func(c *deriveAppConfig) *deriveNestedConfig { return &c.Nested })
	inner := Derive(nested, func(c *deriveNestedConfig) *deriveInnerConfig { return &c.Inner })

	require.Equal(t, "a", inner.Load().Name)

	done := make(chan string, 1)
	inner.Observe(func(c *deriveInnerConfig) { done <- c.Name })

	src.EmitConfigUpdate(&deriveAppConfig{Nested: deriveNestedConfig{Inner: deriveInnerConfig{Name: "b"}}})
	require.Equal(t, "b", <-done)
	require.Equal(t, "b", inner.Load().Name)
}

func TestDeriveFeedsAtomic(t *testing.T) {
	src := NewStaticObserver(&deriveAppConfig{
		Sweeper: deriveSweeperConfig{Period: time.Second},
	})
	sweeper := Derive(src, func(c *deriveAppConfig) *deriveSweeperConfig { return &c.Sweeper })
	period := NewAtomicDuration(sweeper, func(c *deriveSweeperConfig) time.Duration { return c.Period })

	require.Equal(t, time.Second, period.Load())

	src.EmitConfigUpdate(&deriveAppConfig{Sweeper: deriveSweeperConfig{Period: 2 * time.Second}})
	require.Equal(t, 2*time.Second, period.Load())
}
