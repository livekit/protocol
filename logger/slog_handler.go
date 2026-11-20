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

import (
	"context"
	"io"
	"log/slog"
	"math"
	"runtime"
	"slices"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/puzpuzpuz/xsync/v4"
	"github.com/zeebo/xxh3"
)

// ---------------------------------------------------------------- levels

// levelGate is a component's resolved level together with its output.
type levelGate interface {
	io.Writer
	enabled(l slog.Level) bool
}

// levelSource resolves components (dotted paths) to gates.
type levelSource interface {
	gate(component string) levelGate
}

const levelDisabled = math.MaxInt64

// slogLeveler resolves component levels, from a Config at the root or from
// a resolver for derived levelers, which can only widen their parent. Gates
// are memoized, and refresh updates them in place so loggers already derived
// see new levels.
type slogLeveler struct {
	conf    *Config
	resolve func(component string) (slog.Level, bool)
	parent  *slogLeveler
	w       io.Writer
	mu      sync.Mutex
	entries *xsync.Map[string, *slogLevelEntry]
}

type slogLevelEntry struct {
	level  atomic.Int64
	parent *slogLevelEntry
	w      io.Writer
}

func (e *slogLevelEntry) enabled(l slog.Level) bool {
	return int64(l) >= e.level.Load() || (e.parent != nil && e.parent.enabled(l))
}

func (e *slogLevelEntry) Write(p []byte) (int, error) { return e.w.Write(p) }

func newSlogLeveler(w io.Writer, conf *Config) *slogLeveler {
	return &slogLeveler{conf: conf, w: w, entries: xsync.NewMap[string, *slogLevelEntry]()}
}

// derive returns a leveler whose components are enabled at the level resolve
// returns, or at l's level if that is lower; resolve returning false leaves
// the component to l.
func (l *slogLeveler) derive(resolve func(component string) (slog.Level, bool)) *slogLeveler {
	return &slogLeveler{resolve: resolve, parent: l, w: l.w, entries: xsync.NewMap[string, *slogLevelEntry]()}
}

func (l *slogLeveler) level(component string) int64 {
	if l.resolve != nil {
		if lvl, ok := l.resolve(component); ok {
			return int64(lvl)
		}
		return levelDisabled
	}
	lvl, ok := l.conf.ResolveComponentLevel(component)
	if !ok {
		return levelDisabled
	}
	return int64(lvl) * 4 // zap level -> slog level
}

func (l *slogLeveler) gate(component string) levelGate {
	if e, ok := l.entries.Load(component); ok {
		return e
	}
	// Serialized against refresh so a concurrent config change is not missed.
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.entries.Load(component); ok {
		return e
	}
	e := &slogLevelEntry{w: l.w}
	if l.parent != nil {
		e.parent = l.parent.gate(component).(*slogLevelEntry)
	}
	e.level.Store(l.level(component))
	l.entries.Store(component, e)
	return e
}

func (l *slogLeveler) refresh() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries.Range(func(component string, e *slogLevelEntry) bool {
		e.level.Store(l.level(component))
		return true
	})
}

// ---------------------------------------------------------------- tee

// teeSink duplicates entries to a second destination, gated by the same
// resolved level as the primary output.
type teeSink interface {
	write(ent *logEntry, r slog.Record)
	withValues(kvs []any) teeSink
	withAttrs(as []slog.Attr) teeSink
	withGate(g levelGate) teeSink
}

// ---------------------------------------------------------------- caller / stack

var callerCache sync.Map // pc -> callerInfo

func callerAt(pc uintptr) callerInfo {
	if c, ok := callerCache.Load(pc); ok {
		return c.(callerInfo)
	}
	f, _ := runtime.CallersFrames([]uintptr{pc}).Next()
	c := callerInfo{defined: f.PC != 0, pc: f.PC, file: f.File, line: f.Line, function: f.Function}
	callerCache.Store(pc, c)
	return c
}

var pcsPool = sync.Pool{New: func() any { s := make([]uintptr, 64); return &s }}

// stackFrom formats the stack from the frame at pc outwards, dropping the
// final runtime frame, exactly like zap's stacktrace formatter.
func stackFrom(pc uintptr) string {
	pp := pcsPool.Get().(*[]uintptr)
	defer pcsPool.Put(pp)
	pcs := *pp
	n := runtime.Callers(1, pcs)
	for n == len(pcs) {
		pcs = make([]uintptr, len(pcs)*2)
		n = runtime.Callers(1, pcs)
		*pp = pcs
	}
	i := slices.Index(pcs[:n], pc)
	if i < 0 {
		return ""
	}
	e := getJSONEnc(flavorJSON)
	defer putJSONEnc(e)
	b := e.buf
	frames := runtime.CallersFrames(pcs[i:n])
	for first := true; ; first = false {
		f, more := frames.Next()
		if !more && !first {
			break
		}
		if !first {
			b = append(b, '\n')
		}
		b = append(b, f.Function...)
		b = append(b, '\n', '\t')
		b = append(b, f.File...)
		b = append(b, ':')
		b = strconv.AppendInt(b, int64(f.Line), 10)
		if !more {
			break
		}
	}
	e.buf = b
	return string(b)
}

// entryMeta carries caller, stack and name for records whose call site has
// already been resolved (deferred writes, the zap bridge).
type entryMeta struct {
	caller callerInfo
	stack  string
	name   string
}

type entryMetaKey struct{}

// ---------------------------------------------------------------- handler chain

// chainHandler is the handler contract the slog logger needs beyond
// slog.Handler.
type chainHandler interface {
	slog.Handler
	// withValues adds context with the zap backend's WithValues semantics:
	// pairs with non-string keys, and a trailing key, are silently dropped.
	withValues(kvs []any) chainHandler
	withName(n string) chainHandler
	withGate(g levelGate) chainHandler
	loggerName() string
}

// coreHandler encodes and writes entries, then duplicates them to the tee.
type coreHandler struct {
	flavor encFlavor
	gate   levelGate
	name   string
	ctx    []byte
	ctxNS  int
	tee    teeSink
}

var _ chainHandler = (*coreHandler)(nil)

func newCoreHandler(json bool, g levelGate, tee teeSink) *coreHandler {
	f := flavorConsole
	if json {
		f = flavorJSON
	}
	return &coreHandler{flavor: f, gate: g, tee: tee}
}

func (h *coreHandler) Enabled(_ context.Context, l slog.Level) bool {
	return h.gate.enabled(normLevel(l))
}

func (h *coreHandler) Handle(ctx context.Context, r slog.Record) error {
	ent := logEntry{level: normLevel(r.Level), time: r.Time, name: h.name, msg: r.Message}
	if m, ok := ctx.Value(entryMetaKey{}).(*entryMeta); ok {
		ent.caller, ent.stack, ent.name = m.caller, m.stack, m.name
	} else if r.PC != 0 {
		ent.caller = callerAt(r.PC)
		if ent.level >= slog.LevelError {
			ent.stack = stackFrom(r.PC)
		}
	}

	line := getJSONEnc(h.flavor)
	encodeEntry(line, &ent, h.ctx, h.ctxNS, func(e *jsonEnc) {
		r.Attrs(func(a slog.Attr) bool {
			addAttr(e, a)
			return true
		})
	})
	_, err := h.gate.Write(line.buf)
	putJSONEnc(line)

	if h.tee != nil {
		h.tee.write(&ent, r)
	}
	return err
}

func (h *coreHandler) clone() *coreHandler {
	c := *h
	c.ctx = slices.Clip(h.ctx)
	return &c
}

func (h *coreHandler) encodeCtx(add func(e *jsonEnc)) ([]byte, int) {
	e := getJSONEnc(h.flavor)
	e.buf = append(e.buf, h.ctx...)
	add(e)
	b, ns := slices.Clone(e.buf), h.ctxNS+e.openNS
	putJSONEnc(e)
	return b, ns
}

func (h *coreHandler) WithAttrs(as []slog.Attr) slog.Handler {
	if len(as) == 0 {
		return h
	}
	c := h.clone()
	c.ctx, c.ctxNS = h.encodeCtx(func(e *jsonEnc) {
		for _, a := range as {
			addAttr(e, a)
		}
	})
	if c.tee != nil {
		c.tee = c.tee.withAttrs(as)
	}
	return c
}

func (h *coreHandler) WithGroup(n string) slog.Handler {
	if n == "" {
		return h
	}
	c := h.clone()
	c.ctx, c.ctxNS = h.encodeCtx(func(e *jsonEnc) { e.OpenNamespace(n) })
	if c.tee != nil {
		c.tee = c.tee.withAttrs([]slog.Attr{{Key: n, Value: slog.AnyValue(namespaceMarker{})}})
	}
	return c
}

// namespaceMarker tells a tee to open a namespace, the zap equivalent of
// slog.Logger.WithGroup.
type namespaceMarker struct{}

func (h *coreHandler) withValues(kvs []any) chainHandler {
	c := h.clone()
	c.ctx, c.ctxNS = h.encodeCtx(func(e *jsonEnc) {
		for i := 1; i < len(kvs); i += 2 {
			if k, ok := kvs[i-1].(string); ok {
				addAny(e, k, kvs[i])
			}
		}
	})
	if c.tee != nil {
		c.tee = c.tee.withValues(kvs)
	}
	return c
}

func (h *coreHandler) withName(n string) chainHandler {
	c := h.clone()
	if c.name == "" {
		c.name = n
	} else {
		c.name += "." + n
	}
	return c
}

func (h *coreHandler) withGate(g levelGate) chainHandler {
	c := h.clone()
	c.gate = g
	if c.tee != nil {
		c.tee = c.tee.withGate(g)
	}
	return c
}

func (h *coreHandler) loggerName() string { return h.name }

// ---------------------------------------------------------------- deferred values

type deferredWrite struct {
	ctx context.Context
	h   slog.Handler
	r   slog.Record
}

// deferredHandler buffers records until the deferrer is resolved. Its inner
// handler is a chainHandler for slogLogger, or any handler for slogAdapter;
// the chainHandler methods are only used in the former case.

// slogDeferrer implements DeferredFieldResolver with zaputil.Deferrer's
// semantics, including its two field orderings.
type slogDeferrer struct {
	mu     sync.Mutex
	attrs  []slog.Attr
	set    bool
	writes []deferredWrite
}

var _ DeferredFieldResolver = (*slogDeferrer)(nil)

func (d *slogDeferrer) Resolve(args ...any) {
	// Fields pass through, string keys take the next argument, anything else
	// is skipped.
	var attrs []slog.Attr
	for i := 0; i < len(args); i++ {
		switch a := args[i].(type) {
		case string:
			if i < len(args)-1 {
				attrs = append(attrs, attrOf(a, args[i+1]))
				i++
			}
		default:
			if f, ok := asBareField(a); ok {
				attrs = append(attrs, f)
			}
		}
	}

	d.mu.Lock()
	if d.set {
		// Later resolutions override keys and keep the rest.
		for _, p := range d.attrs {
			if !slices.ContainsFunc(attrs, func(a slog.Attr) bool { return a.Key == p.Key }) {
				attrs = append(attrs, p)
			}
		}
		d.attrs = attrs
		d.mu.Unlock()
		return
	}
	d.attrs, d.set = attrs, true
	writes := d.writes
	d.writes = nil
	d.mu.Unlock()

	for _, w := range writes {
		if !w.h.Enabled(w.ctx, w.r.Level) {
			continue
		}
		// Buffered entries get the deferred fields before their own.
		r := slog.NewRecord(w.r.Time, w.r.Level, w.r.Message, w.r.PC)
		r.AddAttrs(attrs...)
		w.r.Attrs(func(a slog.Attr) bool {
			r.AddAttrs(a)
			return true
		})
		_ = w.h.Handle(w.ctx, r)
	}
}

func (d *slogDeferrer) Reset() {
	d.mu.Lock()
	d.attrs, d.set = nil, false
	d.mu.Unlock()
}

type deferredHandler struct {
	inner slog.Handler
	d     *slogDeferrer
}

var _ chainHandler = (*deferredHandler)(nil)

func (h *deferredHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return h.inner.Enabled(ctx, l)
}

func (h *deferredHandler) Handle(ctx context.Context, r slog.Record) error {
	h.d.mu.Lock()
	if !h.d.set {
		ctx = h.capture(ctx, r)
		h.d.writes = append(h.d.writes, deferredWrite{context.WithoutCancel(ctx), h.inner, r.Clone()})
		h.d.mu.Unlock()
		return nil
	}
	attrs := h.d.attrs
	h.d.mu.Unlock()
	// Direct writes get the deferred fields after their own.
	r.AddAttrs(attrs...)
	return h.inner.Handle(ctx, r)
}

// capture resolves caller and stack while the call is still on the stack.
func (h *deferredHandler) capture(ctx context.Context, r slog.Record) context.Context {
	ch, isChain := h.inner.(chainHandler)
	if _, ok := ctx.Value(entryMetaKey{}).(*entryMeta); ok || r.PC == 0 || !isChain {
		return ctx
	}
	m := &entryMeta{caller: callerAt(r.PC), name: ch.loggerName()}
	if normLevel(r.Level) >= slog.LevelError {
		m.stack = stackFrom(r.PC)
	}
	return context.WithValue(ctx, entryMetaKey{}, m)
}

func (h *deferredHandler) WithAttrs(as []slog.Attr) slog.Handler {
	return &deferredHandler{h.inner.WithAttrs(as), h.d}
}

func (h *deferredHandler) WithGroup(n string) slog.Handler {
	return &deferredHandler{h.inner.WithGroup(n), h.d}
}

func (h *deferredHandler) chain() chainHandler { return h.inner.(chainHandler) }

func (h *deferredHandler) withValues(kvs []any) chainHandler {
	return &deferredHandler{h.chain().withValues(kvs), h.d}
}

func (h *deferredHandler) withName(n string) chainHandler {
	return &deferredHandler{h.chain().withName(n), h.d}
}

func (h *deferredHandler) withGate(g levelGate) chainHandler {
	return &deferredHandler{h.chain().withGate(g), h.d}
}

func (h *deferredHandler) loggerName() string { return h.chain().loggerName() }

// ---------------------------------------------------------------- sampling

type sampleCounter struct {
	resetAt atomic.Int64
	n       atomic.Uint64
}

// lineSampler is zaputil.Sampler keyed on (level, message): the first N per
// tick are kept, then every Mth.
type lineSampler struct {
	counts            [len(levelLower)][4096]sampleCounter
	tick              time.Duration
	first, thereafter uint64
}

func newLineSampler(tick time.Duration, first, thereafter int) *lineSampler {
	return &lineSampler{tick: tick, first: uint64(first), thereafter: uint64(thereafter)}
}

func (s *lineSampler) allow(l slog.Level, msg string, t time.Time) bool {
	c := &s.counts[levelIndex(l)][xxh3.HashString(msg)%4096]
	tn := t.UnixNano()
	var n uint64
	if resetAt := c.resetAt.Load(); resetAt > tn {
		n = c.n.Add(1)
	} else {
		c.n.Store(1)
		n = 1
		if !c.resetAt.CompareAndSwap(resetAt, tn+s.tick.Nanoseconds()) {
			// Raced with another reset, which also set the count to 1.
			n = c.n.Add(1)
		}
	}
	return n <= s.first || (s.thereafter != 0 && (n-s.first)%s.thereafter == 0)
}

// samplingHandler samples records from slog.Logger callers, which bypass the
// LiveKit Logger methods.
type samplingHandler struct {
	slog.Handler
	s *lineSampler
}

func (h samplingHandler) Handle(ctx context.Context, r slog.Record) error {
	if !h.s.allow(r.Level, r.Message, r.Time) {
		return nil
	}
	return h.Handler.Handle(ctx, r)
}

func (h samplingHandler) WithAttrs(as []slog.Attr) slog.Handler {
	return samplingHandler{h.Handler.WithAttrs(as), h.s}
}

func (h samplingHandler) WithGroup(n string) slog.Handler {
	return samplingHandler{h.Handler.WithGroup(n), h.s}
}
