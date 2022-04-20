package logger

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger = logr.Discard()
	pkgLogger     = Logger(logr.Discard())
)

// InitFromConfig initializes a Zap-based logger
func InitFromConfig(conf Config, name string) {
	lvl := ParseZapLevel(conf.Level)
	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(lvl),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	if conf.Sample {
		zapConfig.Sampling = &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		}
	}
	if conf.JSON {
		zapConfig.Encoding = "json"
		zapConfig.EncoderConfig = zap.NewProductionEncoderConfig()
	}
	l, _ := zapConfig.Build()
	zapLogger := zapr.NewLogger(l)
	SetLogger(zapLogger, name)
}

// GetLogger returns the logger that was set with SetLogger with an extra depth of 1
func GetLogger() logr.Logger {
	return defaultLogger
}

// GetDefaultLogger returns the logger that was set but with LiveKit wrappers
func GetDefaultLogger() Logger {
	return Logger(defaultLogger)
}

// SetLogger lets you use a custom logger. Pass in a logr.Logger with default depth
func SetLogger(l logr.Logger, name string) {
	defaultLogger = l.WithCallDepth(1).WithName(name)
	// pkg wrapper needs to drop two levels of depth
	pkgLogger = Logger(l.WithCallDepth(2).WithName(name))
}

func Debugw(msg string, keysAndValues ...interface{}) {
	pkgLogger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	pkgLogger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, err error, keysAndValues ...interface{}) {
	pkgLogger.Warnw(msg, err, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...interface{}) {
	pkgLogger.Errorw(msg, err, keysAndValues...)
}

func ParseZapLevel(level string) zapcore.Level {
	lvl := zapcore.InfoLevel
	if level != "" {
		_ = lvl.UnmarshalText([]byte(level))
	}
	return lvl
}

type Logger logr.Logger

func (l Logger) toLogr() logr.Logger {
	if logr.Logger(l).GetSink() == nil {
		return defaultLogger
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
