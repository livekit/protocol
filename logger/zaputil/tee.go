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
	"slices"

	"go.uber.org/zap/zapcore"
)

// Tee duplicates every entry a derived logger writes. The enabler handed to
// newCore is the level that logger resolved, and is what the built core must
// gate on.
type Tee struct {
	newCore func(enab zapcore.LevelEnabler) zapcore.Core
	fields  []zapcore.Field
}

func NewTee(newCore func(enab zapcore.LevelEnabler) zapcore.Core) Tee {
	return Tee{newCore: newCore}
}

// Core returns nil for the zero Tee.
func (t Tee) Core(enab zapcore.LevelEnabler) zapcore.Core {
	if t.newCore == nil {
		return nil
	}
	core := t.newCore(enab)
	if len(t.fields) == 0 {
		return core
	}
	return core.With(t.fields)
}

func (t Tee) WithValues(kvs ...any) Tee {
	if t.newCore == nil {
		return t
	}
	// Clip so sibling loggers derived from this one cannot write into a shared
	// backing array.
	t.fields = append(slices.Clip(t.fields), valueFields(kvs...)...)
	return t
}
