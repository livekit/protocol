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
	ready  bool
	fields []zapcore.Field
	writes []*deferredWrite
}

func (b *Deferrer) buffer(core zapcore.Core, ent zapcore.Entry, fields []zapcore.Field) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.ready {
		return false
	}
	b.writes = append(b.writes, &deferredWrite{core, ent, fields})
	return true
}

func (b *Deferrer) flush() {
	b.mu.Lock()
	b.ready = true
	writes := b.writes
	b.writes = nil
	b.mu.Unlock()

	var fields []zapcore.Field
	for _, w := range writes {
		fields = append(fields[:0], b.fields...)
		fields = append(fields, w.fields...)
		w.core.Write(w.ent, fields)
	}
}

func (b *Deferrer) write(core zapcore.Core, ent zapcore.Entry, fields []zapcore.Field) error {
	if !b.buffer(core, ent, fields) {
		return core.Write(ent, append(fields[0:len(fields):len(fields)], b.fields...))
	}
	return nil
}

type DeferredFieldResolver func(args ...any)

func NewDeferrer() (*Deferrer, DeferredFieldResolver) {
	buf := &Deferrer{}
	var resolveOnce sync.Once
	resolve := func(args ...any) {
		resolveOnce.Do(func() {
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

			buf.fields = fields
			buf.flush()
		})
	}
	return buf, resolve
}

type deferredValueCore struct {
	zapcore.Core
	def *Deferrer
}

func NewDeferredValueCore(core zapcore.Core, def *Deferrer) zapcore.Core {
	def.mu.Lock()
	defer def.mu.Unlock()
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
