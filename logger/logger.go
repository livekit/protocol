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
	Logger(defaultLogger).Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	Logger(defaultLogger).Infow(msg, keysAndValues...)
}

func Warnw(msg string, err error, keysAndValues ...interface{}) {
	Logger(defaultLogger).Warnw(msg, err, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...interface{}) {
	Logger(defaultLogger).Errorw(msg, err, keysAndValues...)
}

type Logger logr.Logger

func (l Logger) toLogr() logr.Logger {
	if logr.Logger(l).GetSink() == nil {
		return logr.Logger(defaultLogger)
	}
	return logr.Logger(l)
}

func (l Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.toLogr().V(1).Info(msg, keysAndValues...)
}

func (l Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.toLogr().Info(msg, keysAndValues...)
}

func (l Logger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.toLogr().Info(msg, keysAndValues...)
}

func (l Logger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	l.toLogr().Error(err, msg, keysAndValues...)
}
