package logger

import (
	"context"
	"log/slog"

	"github.com/go-logr/logr"
	"go.uber.org/zap/exp/zapslog"
)

// NewSlogDiscard creates a slog.Handler that discards all logs.
func NewSlogDiscard() slog.Handler {
	return slogDiscard{}
}

// ToSlogHandler converts Logger to slog.Handler.
func ToSlogHandler(log Logger) slog.Handler {
	switch log := log.(type) {
	case ZapLogger:
		zlog := log.ToZap().Desugar()
		return zapslog.NewHandler(zlog.Core(), &zapslog.HandlerOptions{AddSource: true})
	case LogRLogger:
		return logr.ToSlogHandler(log.toLogr())
	}
	return slogHandler{log, ""}
}

type slogDiscard struct{}

func (_ slogDiscard) Enabled(ctx context.Context, level slog.Level) bool {
	return false
}

func (_ slogDiscard) Handle(ctx context.Context, record slog.Record) error {
	return nil
}

func (l slogDiscard) WithAttrs(attrs []slog.Attr) slog.Handler {
	return l
}

func (l slogDiscard) WithGroup(name string) slog.Handler {
	return l
}

type slogHandler struct {
	log   Logger
	group string
}

func (l slogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true // so such method on Logger
}

func (l slogHandler) getGroup() string {
	group := l.group
	if group != "" {
		group = group + "."
	}
	return group
}

func (l slogHandler) Handle(ctx context.Context, r slog.Record) error {
	keysAndValues := make([]any, 0, r.NumAttrs()*2)
	group := l.getGroup()
	r.Attrs(func(attr slog.Attr) bool {
		keysAndValues = append(keysAndValues, group+attr.Key, attr.Value.Resolve().Any())
		return true
	})
	switch r.Level {
	case slog.LevelDebug:
		l.log.Debugw(r.Message, keysAndValues...)
	default:
		fallthrough
	case slog.LevelInfo:
		l.log.Infow(r.Message, keysAndValues...)
	case slog.LevelWarn:
		l.log.Warnw(r.Message, nil, keysAndValues...)
	case slog.LevelError:
		l.log.Errorw(r.Message, nil, keysAndValues...)
	}
	return nil
}

func (l slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	log := l.log
	keysAndValues := make([]any, 0, len(attrs)*2)
	group := l.getGroup()
	for _, attr := range attrs {
		keysAndValues = append(keysAndValues, group+attr.Key, attr.Value.Resolve().Any())
	}
	log = log.WithValues(keysAndValues...)
	return slogHandler{log, l.group}
}

func (l slogHandler) WithGroup(name string) slog.Handler {
	group := name
	if l.group != "" {
		group = l.group + "." + name
	}
	return slogHandler{l.log, group}
}
