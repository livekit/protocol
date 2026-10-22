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

package zaputil

import (
	"sync"

	"github.com/puzpuzpuz/xsync/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ComponentLevelResolver is given a full dotted component path. Returning false defers
// to the leveler's parent.
type ComponentLevelResolver interface {
	ResolveComponentLevel(component string) (zapcore.Level, bool)
}

type FixedComponentLevel zapcore.Level

func (l FixedComponentLevel) ResolveComponentLevel(string) (zapcore.Level, bool) {
	return zapcore.Level(l), true
}

// level is held apart from enab because only the former is settable once handed out.
type componentEntry struct {
	level zap.AtomicLevel
	enab  zapcore.LevelEnabler
	write *WriteEnabler
}

// One ComponentLeveler per configuration source.
type ComponentLeveler struct {
	parent   *ComponentLeveler
	resolver ComponentLevelResolver
	ws       zapcore.WriteSyncer

	mu      sync.Mutex
	entries *xsync.Map[string, *componentEntry]
}

func NewRootComponentLeveler(ws zapcore.WriteSyncer, r ComponentLevelResolver) *ComponentLeveler {
	return &ComponentLeveler{
		resolver: r,
		ws:       ws,
		entries:  xsync.NewMap[string, *componentEntry](),
	}
}

// NewComponentLeveler derives a leveler that can only widen parent: a level parent
// already enables stays enabled whatever r resolves.
func NewComponentLeveler(parent *ComponentLeveler, r ComponentLevelResolver) *ComponentLeveler {
	return &ComponentLeveler{
		parent:   parent,
		resolver: r,
		ws:       parent.ws,
		entries:  xsync.NewMap[string, *componentEntry](),
	}
}

func (l *ComponentLeveler) ComponentLevel(component string) zapcore.LevelEnabler {
	return l.entry(component).enab
}

// WriteEnabler is memoized per component so that rebuilding a core does not allocate.
func (l *ComponentLeveler) WriteEnabler(component string) *WriteEnabler {
	return l.entry(component).write
}

// Refresh re-resolves in place, so enablers already handed out see the new levels.
func (l *ComponentLeveler) Refresh() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries.Range(func(component string, e *componentEntry) bool {
		e.level.SetLevel(l.resolve(component))
		return true
	})
}

func (l *ComponentLeveler) entry(component string) *componentEntry {
	if e, ok := l.entries.Load(component); ok {
		return e
	}

	// Serialized against Refresh, otherwise an entry resolved here could miss a
	// concurrent configuration change and stay stale forever.
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.entries.Load(component); ok {
		return e
	}

	e := &componentEntry{level: zap.NewAtomicLevelAt(l.resolve(component))}
	e.enab = e.level
	if l.parent != nil {
		e.enab = OrLevelEnabler{l.parent.ComponentLevel(component), e.level}
	}
	e.write = NewWriteEnabler(l.ws, e.enab)
	l.entries.Store(component, e)
	return e
}

// InvalidLevel enables nothing, leaving the Or against the parent decisive.
func (l *ComponentLeveler) resolve(component string) zapcore.Level {
	if lvl, ok := l.resolver.ResolveComponentLevel(component); ok {
		return lvl
	}
	return zapcore.InvalidLevel
}
