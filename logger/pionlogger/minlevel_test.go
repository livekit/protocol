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
}
