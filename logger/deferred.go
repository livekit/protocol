package logger

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

type deferredWriteBuffer struct {
	ready  chan struct{}
	fields []zapcore.Field

	mu     sync.Mutex
	writes []*deferredWrite
}

func (b *deferredWriteBuffer) append(core zapcore.Core, ent zapcore.Entry, fields []zapcore.Field) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.writes = append(b.writes, &deferredWrite{core, ent, fields})
}

func (b *deferredWriteBuffer) flush() {
	b.mu.Lock()
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

func (b *deferredWriteBuffer) write(core zapcore.Core, ent zapcore.Entry, fields []zapcore.Field) error {
	select {
	case <-b.ready:
		return core.Write(ent, append(fields[0:len(fields):len(fields)], b.fields...))
	default:
		b.append(core, ent, fields)
		return nil
	}
}

type DeferredFieldResolver func(args ...any)

type deferredValueCore struct {
	zapcore.Core
	buf *deferredWriteBuffer
}

func newDeferredValueCore(core zapcore.Core) (zapcore.Core, DeferredFieldResolver) {
	buf := &deferredWriteBuffer{ready: make(chan struct{})}
	var resolveOnce sync.Once
	resolve := func(args ...any) {
		resolveOnce.Do(func() {
			fields := make([]zapcore.Field, 0, len(args))
			for i := 0; i < len(args); {
				switch arg := args[i].(type) {
				case zapcore.Field:
					fields = append(fields, arg)
					i++
				case string:
					if i < len(args)-1 {
						fields = append(fields, zap.Any(arg, arg[i+1]))
						i += 2
					}
				}
			}

			buf.fields = fields
			close(buf.ready)
			buf.flush()

		})
	}

	return &deferredValueCore{core, buf}, resolve
}

func (c *deferredValueCore) With(fields []zapcore.Field) zapcore.Core {
	return &deferredValueCore{
		Core: c.Core.With(fields),
		buf:  c.buf,
	}
}

func (c *deferredValueCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *deferredValueCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	return c.buf.write(c.Core, ent, fields)
}
