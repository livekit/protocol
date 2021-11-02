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

type withLogger struct {
	logger logr.Logger
}

func WithLogger(logger logr.Logger) *withLogger {
	return &withLogger{logger: logger}
}

func (l *withLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.V(1).Info(msg, keysAndValues...)
}

func (l *withLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, keysAndValues...)
}

func (l *withLogger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.logger.Info(msg, keysAndValues...)
}

func (l *withLogger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	l.logger.Error(err, msg, keysAndValues...)
}
