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

package logger

import "log/slog"

// ComponentLevels is a set of per-component levels derived from a logger's,
// for example a project's overrides on top of the service configuration. The
// derived levels can only widen the parent's: a level the parent enables stays
// enabled. It replaces building a zaputil.ComponentLeveler from
// ZapLogger.Leveler and applying it with WithComponentLeveler.
type ComponentLevels struct {
	native *slogLeveler
	zap    *ComponentLeveler
}

// NewComponentLevels derives levels from parent's. resolve is given a full
// dotted component path; returning false leaves that component to parent.
func NewComponentLevels(parent Logger, resolve func(component string) (slog.Level, bool)) *ComponentLevels {
	if l, ok := parent.(*slogLogger); ok {
		if root, ok := l.levels.(*slogLeveler); ok {
			return &ComponentLevels{native: root.derive(resolve)}
		}
	}
	if l, ok := parent.(ZapLogger); ok {
		return &ComponentLevels{zap: deriveZapLevels(l.Leveler(), resolve)}
	}
	return &ComponentLevels{}
}

// Refresh re-resolves every component, for after resolve's inputs changed.
// Loggers already using these levels see the result.
func (c *ComponentLevels) Refresh() {
	if c.native != nil {
		c.native.refresh()
	}
	if c.zap != nil {
		c.zap.Refresh()
	}
}

// WithComponentLevels returns l using c for its component levels. l must be
// the logger c was derived from, or derived from it; otherwise l is returned
// unchanged.
func WithComponentLevels(l Logger, c *ComponentLevels) Logger {
	if c == nil {
		return l
	}
	if sl, ok := l.(*slogLogger); ok && c.native != nil {
		return sl.withLevels(c.native)
	}
	if zl, ok := l.(ZapLogger); ok && c.zap != nil {
		return zl.WithComponentLeveler(c.zap)
	}
	return l
}
