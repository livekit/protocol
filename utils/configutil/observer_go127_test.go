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

//go:build go1.27

package configutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestObserverDeriveMethod(t *testing.T) {
	conf := &deriveAppConfig{Sweeper: deriveSweeperConfig{Period: time.Second}}
	src := NewStaticObserver(conf)
	sweeper := src.Derive(func(c *deriveAppConfig) *deriveSweeperConfig { return &c.Sweeper })

	require.Same(t, &conf.Sweeper, sweeper.Load())

	var observed []time.Duration
	unsubscribe := sweeper.Observe(func(c *deriveSweeperConfig) {
		observed = append(observed, c.Period)
	})

	src.EmitConfigUpdate(&deriveAppConfig{Sweeper: deriveSweeperConfig{Period: 2 * time.Second}})
	require.Equal(t, 2*time.Second, sweeper.Load().Period)
	require.Equal(t, []time.Duration{2 * time.Second}, observed)

	unsubscribe()
	src.EmitConfigUpdate(&deriveAppConfig{Sweeper: deriveSweeperConfig{Period: 3 * time.Second}})
	require.Equal(t, []time.Duration{2 * time.Second}, observed, "unsubscribed callback still fired")
}

func TestObserverDeriveMethodChained(t *testing.T) {
	src := NewStaticObserver(&deriveAppConfig{
		Nested: deriveNestedConfig{Inner: deriveInnerConfig{Name: "a"}},
	})
	nested := src.Derive(func(c *deriveAppConfig) *deriveNestedConfig { return &c.Nested })
	inner := Derive(nested, func(c *deriveNestedConfig) *deriveInnerConfig { return &c.Inner })

	var seen []string
	stop := inner.ObserveCurrent(func(c *deriveInnerConfig) { seen = append(seen, c.Name) })
	src.EmitConfigUpdate(&deriveAppConfig{Nested: deriveNestedConfig{Inner: deriveInnerConfig{Name: "b"}}})
	stop()

	require.Equal(t, []string{"a", "b"}, seen)
}

func TestObserverDeriveMethodNilReceiver(t *testing.T) {
	var src *Observer[deriveAppConfig]
	derived := src.Derive(func(c *deriveAppConfig) *deriveSweeperConfig { return &c.Sweeper })

	// A typed-nil wrapper would pass a != nil check and panic on first use.
	require.True(t, derived == nil)
}
