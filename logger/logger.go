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

package logger

import (
	"log/slog"
	"os"
	"slices"
	"sync/atomic"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger/zaputil"
)

var (
	discardLogger      = logr.Discard()
	discardLoggerIface = LogRLogger(discardLogger)
	defaultLogger      = Logger(discardLoggerIface)
	pkgLogger          = Logger(discardLoggerIface)
)

// InitFromConfig initializes a Zap-based logger
func InitFromConfig(conf *Config, name string) {
	l, err := NewZapLogger(conf)
	if err == nil {
		SetLogger(l, name)
		slog.SetDefault(slog.New(ToSlogHandler(l)))
	}
}

// GetLogger returns the logger that was set with SetLogger with an extra depth of 1
func GetLogger() Logger {
	return defaultLogger
}

func GetDiscardLogger() Logger {
	return discardLoggerIface
}

// SetLogger lets you use a custom logger. Pass in a logr.Logger with default depth
func SetLogger(l Logger, name string) {
	defaultLogger = l.WithName(name)
	// pkg wrapper needs to drop two levels of depth
	pkgLogger = l.WithCallDepth(1).WithName(name)
}

func Debugw(msg string, keysAndValues ...any) {
	pkgLogger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...any) {
	pkgLogger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, err error, keysAndValues ...any) {
	pkgLogger.Warnw(msg, err, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...any) {
	pkgLogger.Errorw(msg, err, keysAndValues...)
}

func ParseZapLevel(level string) zapcore.Level {
	lvl := zapcore.InfoLevel
	if level != "" {
		_ = lvl.UnmarshalText([]byte(level))
	}
	return lvl
}

type (
	ComponentLeveler       = zaputil.ComponentLeveler
	ComponentLevelResolver = zaputil.ComponentLevelResolver
	DeferredFieldResolver  = zaputil.DeferredFieldResolver
	FixedComponentLevel    = zaputil.FixedComponentLevel
)

type Logger interface {
	Debugw(msg string, keysAndValues ...any)
	Infow(msg string, keysAndValues ...any)
	Warnw(msg string, err error, keysAndValues ...any)
	Errorw(msg string, err error, keysAndValues ...any)
	WithValues(keysAndValues ...any) Logger
	WithUnlikelyValues(keysAndValues ...any) UnlikelyLogger
	WithName(name string) Logger
	// WithComponent creates a new logger with name as "<name>.<component>", and uses a log level as specified
	WithComponent(component string) Logger
	WithCallDepth(depth int) Logger
	WithItemSampler() Logger
	// WithoutSampler returns the original logger without sampling
	WithoutSampler() Logger
	WithDeferredValues() (Logger, DeferredFieldResolver)
}

type UnlikelyLogger struct {
	logger        Logger
	keysAndValues []any
}

func NewUnlikelyLogger(logger Logger, keysAndValues ...any) UnlikelyLogger {
	return UnlikelyLogger{logger, keysAndValues}
}

func (l UnlikelyLogger) makeLogger() Logger {
	return l.logger.WithCallDepth(1)
}

func (l UnlikelyLogger) Debugw(msg string, keysAndValues ...any) {
	l.makeLogger().Debugw(msg, slices.Concat(l.keysAndValues, keysAndValues)...)
}

func (l UnlikelyLogger) Infow(msg string, keysAndValues ...any) {
	l.makeLogger().Infow(msg, slices.Concat(l.keysAndValues, keysAndValues)...)
}

func (l UnlikelyLogger) Warnw(msg string, err error, keysAndValues ...any) {
	l.makeLogger().Warnw(msg, err, slices.Concat(l.keysAndValues, keysAndValues)...)
}

func (l UnlikelyLogger) Errorw(msg string, err error, keysAndValues ...any) {
	l.makeLogger().Errorw(msg, err, slices.Concat(l.keysAndValues, keysAndValues)...)
}

func (l UnlikelyLogger) WithValues(keysAndValues ...any) UnlikelyLogger {
	return UnlikelyLogger{l.logger, slices.Concat(l.keysAndValues, keysAndValues)}
}

type zapConfig struct {
	conf *Config
	root *ComponentLeveler
	tee  zaputil.Tee
}

type ZapLoggerOption func(*zapConfig)

// The tee does its own encoding, but shares the console's level: it cannot
// widen what the logger emits.
func WithTee(tee zaputil.Tee) ZapLoggerOption {
	return func(zc *zapConfig) {
		zc.tee = tee
	}
}

type ZapComponentLeveler interface {
	ComponentLevel(component string) zapcore.LevelEnabler
}

type ZapLogger interface {
	Logger
	ToZap() *zap.SugaredLogger
	// ComponentLeveler names components relative to this logger, unlike Leveler, which
	// takes paths from the root.
	ComponentLeveler() ZapComponentLeveler
	Leveler() *ComponentLeveler
	// lv supplies the write syncer as well as the level, so it must be derived from this
	// logger's Leveler or output goes wherever its root points.
	WithComponentLeveler(lv *ComponentLeveler) Logger
}

type zapLogger struct {
	zap *zap.SugaredLogger
	*zapConfig
	enc       zaputil.Encoder
	component string
	deferred  []*zaputil.Deferrer
	sampler   *zaputil.Sampler
	leveler   *ComponentLeveler
	tee       zaputil.Tee
}

func FromZapLogger(log *zap.Logger, conf *Config, opts ...ZapLoggerOption) (ZapLogger, error) {
	if log == nil {
		log = zap.New(nil).WithOptions(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	}
	zap := log.WithOptions(zap.AddCallerSkip(1)).Sugar()

	zc := &zapConfig{
		conf: conf,
		root: zaputil.NewRootComponentLeveler(os.Stderr, conf),
	}
	conf.AddUpdateObserver(func(*Config) error {
		zc.root.Refresh()
		return nil
	})
	for _, opt := range opts {
		opt(zc)
	}

	var sampler *zaputil.Sampler
	if conf.Sample {
		var initial = 20
		var interval = 100
		if conf.ItemSampleInitial != 0 {
			initial = conf.ItemSampleInitial
		}
		if conf.ItemSampleInterval != 0 {
			interval = conf.ItemSampleInterval
		}
		sampler = zaputil.NewSampler(time.Second, initial, interval)
	}

	if conf.JSON {
		return newZapLogger(zap, zc, zaputil.NewProductionEncoder(), sampler), nil
	}
	return newZapLogger(zap, zc, zaputil.NewDevelopmentEncoder(), sampler), nil
}

func NewZapLogger(conf *Config, opts ...ZapLoggerOption) (ZapLogger, error) {
	return FromZapLogger(nil, conf, opts...)
}

func newZapLogger(zap *zap.SugaredLogger, zc *zapConfig, enc zaputil.Encoder, sampler *zaputil.Sampler) ZapLogger {
	l := &zapLogger{
		zap:       zap,
		zapConfig: zc,
		enc:       enc,
		sampler:   sampler,
		leveler:   zc.root,
		tee:       zc.tee,
	}
	l.zap = l.makeZap()
	return l
}

func (l *zapLogger) makeZap() *zap.SugaredLogger {
	console := l.leveler.WriteEnabler(l.component)

	c := l.enc.Core(console)
	if tee := l.tee.Core(console); tee != nil {
		c = zapcore.NewTee(c, tee)
	}
	for i := range l.deferred {
		c = zaputil.NewDeferredValueCore(c, l.deferred[i])
	}
	if l.sampler != nil {
		c = zaputil.NewSamplerCore(c, l.sampler)
	}

	return l.zap.WithOptions(zap.WrapCore(func(zapcore.Core) zapcore.Core { return c }))
}

func (l *zapLogger) ToZap() *zap.SugaredLogger {
	return l.zap.WithOptions(zap.AddCallerSkip(-1))
}

type zapLoggerComponentLeveler struct {
	zl *zapLogger
}

func (l zapLoggerComponentLeveler) ComponentLevel(component string) zapcore.LevelEnabler {
	if l.zl.component != "" {
		component = l.zl.component + "." + component
	}

	return l.zl.leveler.ComponentLevel(component)
}

func (l *zapLogger) ComponentLeveler() ZapComponentLeveler {
	return zapLoggerComponentLeveler{l}
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...any) {
	l.zap.Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Leveler() *ComponentLeveler {
	return l.leveler
}

func (l *zapLogger) WithComponentLeveler(lv *ComponentLeveler) Logger {
	if lv == nil {
		return l
	}
	dup := *l
	dup.leveler = lv
	dup.zap = dup.makeZap()
	return &dup
}

func (l *zapLogger) Infow(msg string, keysAndValues ...any) {
	l.zap.Infow(msg, keysAndValues...)
}

func (l *zapLogger) Warnw(msg string, err error, keysAndValues ...any) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.zap.Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Errorw(msg string, err error, keysAndValues ...any) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.zap.Errorw(msg, keysAndValues...)
}

func (l *zapLogger) WithValues(keysAndValues ...any) Logger {
	dup := *l
	dup.enc = dup.enc.WithValues(keysAndValues...)
	dup.tee = dup.tee.WithValues(keysAndValues...)
	dup.zap = dup.makeZap()
	return &dup
}

func (l *zapLogger) WithUnlikelyValues(keysAndValues ...any) UnlikelyLogger {
	return UnlikelyLogger{l, keysAndValues}
}

func (l *zapLogger) WithName(name string) Logger {
	dup := *l
	dup.zap = dup.zap.Named(name)
	return &dup
}

func (l *zapLogger) WithComponent(component string) Logger {
	dup := *l
	dup.zap = dup.zap.Named(component)
	if dup.component == "" {
		dup.component = component
	} else {
		dup.component = dup.component + "." + component
	}
	dup.zap = dup.makeZap()
	return &dup
}

func (l *zapLogger) WithCallDepth(depth int) Logger {
	dup := *l
	dup.zap = dup.zap.WithOptions(zap.AddCallerSkip(depth))
	return &dup
}

func (l *zapLogger) WithItemSampler() Logger {
	if l.conf.ItemSampleSeconds == 0 {
		return l
	}
	dup := *l
	dup.sampler = zaputil.NewSampler(
		time.Duration(l.conf.ItemSampleSeconds)*time.Second,
		l.conf.ItemSampleInitial,
		l.conf.ItemSampleInterval,
	)
	dup.zap = dup.makeZap()
	return &dup
}

func (l *zapLogger) WithoutSampler() Logger {
	dup := *l
	dup.sampler = nil
	dup.zap = dup.makeZap()
	return &dup
}

func (l *zapLogger) WithDeferredValues() (Logger, DeferredFieldResolver) {
	dup := *l
	def := &zaputil.Deferrer{}
	dup.deferred = append(dup.deferred[0:len(dup.deferred):len(dup.deferred)], def)
	dup.zap = dup.makeZap()
	return &dup, def
}

type LogRLogger logr.Logger

func (l LogRLogger) toLogr() logr.Logger {
	if logr.Logger(l).GetSink() == nil {
		return discardLogger
	}
	return logr.Logger(l)
}

func (l LogRLogger) Debugw(msg string, keysAndValues ...any) {
	l.toLogr().V(1).Info(msg, keysAndValues...)
}

func (l LogRLogger) Infow(msg string, keysAndValues ...any) {
	l.toLogr().Info(msg, keysAndValues...)
}

func (l LogRLogger) Warnw(msg string, err error, keysAndValues ...any) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.toLogr().Info(msg, keysAndValues...)
}

func (l LogRLogger) Errorw(msg string, err error, keysAndValues ...any) {
	l.toLogr().Error(err, msg, keysAndValues...)
}

func (l LogRLogger) WithValues(keysAndValues ...any) Logger {
	return LogRLogger(l.toLogr().WithValues(keysAndValues...))
}

func (l LogRLogger) WithUnlikelyValues(keysAndValues ...any) UnlikelyLogger {
	return UnlikelyLogger{l, keysAndValues}
}

func (l LogRLogger) WithName(name string) Logger {
	return LogRLogger(l.toLogr().WithName(name))
}

func (l LogRLogger) WithComponent(component string) Logger {
	return LogRLogger(l.toLogr().WithName(component))
}

func (l LogRLogger) WithCallDepth(depth int) Logger {
	return LogRLogger(l.toLogr().WithCallDepth(depth))
}

func (l LogRLogger) WithItemSampler() Logger {
	// logr does not support sampling
	return l
}

func (l LogRLogger) WithoutSampler() Logger {
	return l
}

func (l LogRLogger) WithDeferredValues() (Logger, DeferredFieldResolver) {
	return l, zaputil.NoOpDeferrer{}
}

type TestLogger interface {
	Logf(format string, args ...any)
	Log(args ...any)
	Cleanup(f func())
}

func NewTestLogger(t TestLogger) Logger {
	return NewTestLoggerLevel(t, 0)
}

func NewTestLoggerLevel(t TestLogger, lvl int) Logger {
	var closed atomic.Bool
	t.Cleanup(func() {
		closed.Store(true)
	})
	return LogRLogger(funcr.New(func(prefix, args string) {
		if closed.Load() {
			return
		}
		if prefix != "" {
			t.Logf("%s: %s\n", prefix, args)
		} else {
			t.Log(args)
		}
	}, funcr.Options{Verbosity: lvl}))
}
