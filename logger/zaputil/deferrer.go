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
	"sync"

	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type deferredWrite struct {
	core   zapcore.Core
	ent    zapcore.Entry
	fields []zapcore.Field
}

type Deferrer struct {
	mu     sync.Mutex
	fields atomic.Pointer[[]zapcore.Field]
	writes []*deferredWrite
}

func (b *Deferrer) buffer(core zapcore.Core, ent zapcore.Entry, fields []zapcore.Field) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.fields.Load() != nil {
		return false
	}
	b.writes = append(b.writes, &deferredWrite{core, ent, fields})
	return true
}

func (b *Deferrer) flush() {
	b.mu.Lock()
	writes := b.writes
	b.writes = nil
	b.mu.Unlock()

	fields := slices.Clone(*b.fields.Load())
	n := len(fields)

	for _, w := range writes {
		fields = append(fields[:n], w.fields...)
		w.core.Write(w.ent, fields)
	}
}

func (b *Deferrer) write(core zapcore.Core, ent zapcore.Entry, fields []zapcore.Field) error {
	for {
		if dfs := b.fields.Load(); dfs != nil {
			return core.Write(ent, slices.Concat(fields, *dfs))
		}
		if b.buffer(core, ent, fields) {
			return nil
		}
	}
}

type DeferredFieldResolver func(args ...any)

func NewDeferrer() (*Deferrer, DeferredFieldResolver) {
	buf := &Deferrer{}

	resolve := func(args ...any) {
		fields := make([]zapcore.Field, 0, len(args))
		for i := 0; i < len(args); i++ {
			switch arg := args[i].(type) {
			case zapcore.Field:
				fields = append(fields, arg)
			case string:
				if i < len(args)-1 {
					fields = append(fields, zap.Any(arg, args[i+1]))
					i++
				}
			}
		}

		if buf.fields.Swap(&fields) == nil {
			buf.flush()
		}
	}
	return buf, resolve
}

type deferredValueCore struct {
	zapcore.Core
	def *Deferrer
}

func NewDeferredValueCore(core zapcore.Core, def *Deferrer) zapcore.Core {
	return &deferredValueCore{core, def}
}

func (c *deferredValueCore) With(fields []zapcore.Field) zapcore.Core {
	return &deferredValueCore{
		Core: c.Core.With(fields),
		def:  c.def,
	}
}

func (c *deferredValueCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *deferredValueCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	return c.def.write(c.Core, ent, fields)
}
