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
	defaultLogger.V(1).Info(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	defaultLogger.Info(msg, keysAndValues...)
}

func Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	defaultLogger.Info(msg, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...interface{}) {
	defaultLogger.Error(err, msg, keysAndValues...)
}
