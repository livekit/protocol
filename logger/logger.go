package logger

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger = logr.Discard()
)

// Note: already with extra depth 1
func GetLogger() logr.Logger {
	if defaultLogger == logr.Discard() {
		InitDevelopment("default", "")
	}
	return defaultLogger
}

// Note: only pass in logr.Logger with default depth
func SetLogger(l logr.Logger, name string) {
	defaultLogger = l.WithCallDepth(1).WithName(name)
}

func InitDevelopment(name, logLevel string) {
	initLogger(zap.NewDevelopmentConfig(), name, logLevel)
}

func InitProduction(name, logLevel string) {
	initLogger(zap.NewProductionConfig(), name, logLevel)
}

// valid levels: debug, info, warn, error, fatal, panic
func initLogger(config zap.Config, name, level string) {
	if level != "" {
		lvl := zapcore.Level(0)
		if err := lvl.UnmarshalText([]byte(level)); err == nil {
			config.Level = zap.NewAtomicLevelAt(lvl)
		}
	}

	logger, _ := config.Build()
	SetLogger(zapr.NewLogger(logger), name)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	defaultLogger.V(2).Info(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	defaultLogger.V(1).Info(msg, keysAndValues...)
}

func Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append([]interface{}{"error", err}, keysAndValues...)
	}
	defaultLogger.V(1).Error(err, msg, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...interface{}) {
	defaultLogger.Error(err, msg, keysAndValues...)
}
