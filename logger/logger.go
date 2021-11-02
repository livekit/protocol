package logger

import (
	"github.com/go-logr/logr"
)

var (
	defaultLogger = logr.Discard()
)

// Note: SetLogger already adds extra depth 1
func GetLogger() logr.Logger {
	return defaultLogger
}

// Note: only pass in logr.Logger with default depth
func SetLogger(l logr.Logger, name string) {
	defaultLogger = l.WithCallDepth(1).WithName(name)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	WithLogger(defaultLogger).Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	WithLogger(defaultLogger).Infow(msg, keysAndValues...)
}

func Warnw(msg string, err error, keysAndValues ...interface{}) {
	WithLogger(defaultLogger).Warnw(msg, err, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...interface{}) {
	WithLogger(defaultLogger).Errorw(msg, err, keysAndValues...)
}

type Logger logr.Logger

func WithLogger(logger logr.Logger) *Logger {
	l := Logger(logger)
	return &l
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	logger := defaultLogger
	if l != nil {
		logger = logr.Logger(*l)
	}

	logger.V(1).Info(msg, keysAndValues...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	logger := defaultLogger
	if l != nil {
		logger = logr.Logger(*l)
	}

	logger.Info(msg, keysAndValues...)
}

func (l *Logger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	logger := defaultLogger
	if l != nil {
		logger = logr.Logger(*l)
	}

	logger.Info(msg, keysAndValues...)
}

func (l *Logger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	logger := defaultLogger
	if l != nil {
		logger = logr.Logger(*l)
	}

	logger.Error(err, msg, keysAndValues...)
}
