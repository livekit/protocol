// Copyright 2023 LiveKit, Inc.
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

package pionlogger

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger"
)

// componentMinLeveler resolves exact component paths. Real implementations
// search up the component hierarchy; that search is the caller's concern.
type componentMinLeveler map[string]zapcore.LevelEnabler

func (m componentMinLeveler) ComponentMinLevel(component string) zapcore.LevelEnabler {
	return m[component]
}

func debugLevel() zapcore.LevelEnabler { return zap.NewAtomicLevelAt(zapcore.DebugLevel) }

// capture swaps os.Stderr for the duration of f. Loggers must be built inside
// f: makeZap binds os.Stderr when the logger is created, not when it emits.
func capture(t *testing.T, f func(base logger.ZapLogger)) string {
	t.Helper()

	r, w, err := os.Pipe()
	require.NoError(t, err)

	orig := os.Stderr
	os.Stderr = w

	func() {
		defer func() { os.Stderr = orig }()

		// staging/production shape: everything at info, pion quieted to warn
		base, err := logger.NewZapLogger(&logger.Config{
			Level:           "info",
			ComponentLevels: map[string]string{"transport.pion": "warn"},
		})
		require.NoError(t, err)
		f(base)
	}()

	require.NoError(t, w.Close())
	out, err := io.ReadAll(r)
	require.NoError(t, err)
	return string(out)
}

// pionDebug logs a debug line from pion scope through the same path
// pkg/rtc/transport.go uses: a participant logger, WithComponent("transport"),
// handed to the factory.
func pionDebug(l logger.Logger, scope, msg string) {
	NewLoggerFactory(l.WithComponent("transport")).NewLogger(scope).Debug(msg)
}

func TestMinLevelReachesPionComponents(t *testing.T) {
	t.Run("no override leaves pion at its configured level", func(t *testing.T) {
		out := capture(t, func(base logger.ZapLogger) {
			pionDebug(base, "ice", "ice-debug")
		})
		require.NotContains(t, out, "ice-debug")
	})

	// The behavior existing `project_levels: {<id>: debug}` config relies on.
	// A logger-wide floor must not drag pion along with it.
	t.Run("min level alone does not open pion", func(t *testing.T) {
		out := capture(t, func(base logger.ZapLogger) {
			proj := base.WithMinLevel(debugLevel())
			proj.WithComponent("transport").Debugw("transport-debug")
			pionDebug(proj, "ice", "ice-debug")
		})
		require.Contains(t, out, "transport-debug", "min level should still apply to non-gated components")
		require.NotContains(t, out, "ice-debug")
	})

	t.Run("component min leveler opens only the listed component", func(t *testing.T) {
		out := capture(t, func(base logger.ZapLogger) {
			proj := base.WithComponentMinLeveler(componentMinLeveler{"transport.pion.ice": debugLevel()})
			pionDebug(proj, "ice", "ice-debug")
			pionDebug(proj, "sctp", "sctp-debug")
		})
		require.Contains(t, out, "ice-debug")
		require.NotContains(t, out, "sctp-debug", "an unlisted pion component must stay at its configured level")
	})

	// A project at `level: debug` that also lists one pion component gets that
	// component and no other: the logger-wide floor still must not leak into
	// the gated ones.
	t.Run("min level and component min leveler together", func(t *testing.T) {
		out := capture(t, func(base logger.ZapLogger) {
			proj := base.WithMinLevel(debugLevel()).(logger.ZapLogger).
				WithComponentMinLeveler(componentMinLeveler{"transport.pion.ice": debugLevel()})
			pionDebug(proj, "ice", "ice-debug")
			pionDebug(proj, "sctp", "sctp-debug")
		})
		require.Contains(t, out, "ice-debug")
		require.NotContains(t, out, "sctp-debug")
	})

	// Enablers are memoized on state shared by every logger built from one
	// config, so an override must never be cached under a bare component key.
	t.Run("override does not leak to sibling loggers", func(t *testing.T) {
		out := capture(t, func(base logger.ZapLogger) {
			withOverride := base.WithComponentMinLeveler(componentMinLeveler{"transport.pion.ice": debugLevel()})
			pionDebug(withOverride, "ice", "overridden-ice-debug")

			// same component path, different logger, no override
			pionDebug(base, "ice", "other-ice-debug")
		})
		require.Contains(t, out, "overridden-ice-debug")
		require.NotContains(t, out, "other-ice-debug")
	})

	t.Run("override does not leak to loggers built before it", func(t *testing.T) {
		out := capture(t, func(base logger.ZapLogger) {
			pionDebug(base, "ice", "first-ice-debug")

			withOverride := base.WithComponentMinLeveler(componentMinLeveler{"transport.pion.ice": debugLevel()})
			pionDebug(withOverride, "ice", "overridden-ice-debug")
			pionDebug(base, "ice", "last-ice-debug")
		})
		require.NotContains(t, out, "first-ice-debug")
		require.Contains(t, out, "overridden-ice-debug")
		require.NotContains(t, out, "last-ice-debug")
	})

	// pion asks the factory for a logger per object, not per peer connection -
	// once per RTPReceiver, per data channel, per stream - so the same
	// component is resolved over and over for one participant.
	t.Run("resolving a component repeatedly does not allocate", func(t *testing.T) {
		base, err := logger.NewZapLogger(&logger.Config{
			Level:           "info",
			ComponentLevels: map[string]string{"transport.pion": "warn"},
		})
		require.NoError(t, err)

		proj := base.WithComponentMinLeveler(componentMinLeveler{"transport.pion.ice": debugLevel()})
		leveler := proj.WithComponent("transport").(logger.ZapLogger).ComponentLeveler()
		require.True(t, leveler.ComponentLevel("pion.ice").Enabled(zapcore.DebugLevel))
		require.False(t, leveler.ComponentLevel("pion.sctp").Enabled(zapcore.DebugLevel))

		// sctp has no override and so takes the levels cached for every logger
		// at once; ice must cost the same, both of them only the component path
		// this builds to look itself up by
		overridden := testing.AllocsPerRun(100, func() { _ = leveler.ComponentLevel("pion.ice") })
		plain := testing.AllocsPerRun(100, func() { _ = leveler.ComponentLevel("pion.sctp") })
		require.Equal(t, plain, overridden)
	})

	// A memoized enabler holds the atomic levels themselves, so it keeps
	// tracking config after it is cached.
	t.Run("cached enablers follow a config reload", func(t *testing.T) {
		conf := &logger.Config{
			Level:           "info",
			ComponentLevels: map[string]string{"transport.pion": "warn"},
		}

		r, w, err := os.Pipe()
		require.NoError(t, err)

		orig := os.Stderr
		os.Stderr = w

		func() {
			defer func() { os.Stderr = orig }()

			base, err := logger.NewZapLogger(conf)
			require.NoError(t, err)

			level := zap.NewAtomicLevelAt(zapcore.WarnLevel)
			proj := base.WithComponentMinLeveler(componentMinLeveler{"transport.pion.ice": level})

			// caches the enabler at warn
			pionDebug(proj, "ice", "before-reload-debug")

			level.SetLevel(zapcore.DebugLevel)
			pionDebug(proj, "ice", "after-reload-debug")
		}()

		require.NoError(t, w.Close())
		out, err := io.ReadAll(r)
		require.NoError(t, err)

		require.NotContains(t, string(out), "before-reload-debug")
		require.Contains(t, string(out), "after-reload-debug", "a cached enabler must hold the level, not its value")
	})
}
