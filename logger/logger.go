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
	"strings"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/puzpuzpuz/xsync/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger/zaputil"
)

var (
	discardLogger        = logr.Discard()
	defaultLogger Logger = LogRLogger(discardLogger)
	pkgLogger     Logger = LogRLogger(discardLogger)
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

// SetLogger lets you use a custom logger. Pass in a logr.Logger with default depth
func SetLogger(l Logger, name string) {
	defaultLogger = l.WithCallDepth(1).WithName(name)
	// pkg wrapper needs to drop two levels of depth
	pkgLogger = l.WithCallDepth(2).WithName(name)
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

type DeferredFieldResolver = zaputil.DeferredFieldResolver

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, err error, keysAndValues ...interface{})
	Errorw(msg string, err error, keysAndValues ...interface{})
	WithValues(keysAndValues ...interface{}) Logger
	WithName(name string) Logger
	// WithComponent creates a new logger with name as "<name>.<component>", and uses a log level as specified
	WithComponent(component string) Logger
	WithCallDepth(depth int) Logger
	WithItemSampler() Logger
	// WithoutSampler returns the original logger without sampling
	WithoutSampler() Logger
	WithDeferredValues() (Logger, DeferredFieldResolver)
}

type sharedConfig struct {
	level           zap.AtomicLevel
	mu              sync.Mutex
	componentLevels map[string]zap.AtomicLevel
	config          *Config
}

func newSharedConfig(conf *Config) *sharedConfig {
	sc := &sharedConfig{
		level:           zap.NewAtomicLevelAt(ParseZapLevel(conf.Level)),
		config:          conf,
		componentLevels: make(map[string]zap.AtomicLevel),
	}
	conf.AddUpdateObserver(sc.onConfigUpdate)
	_ = sc.onConfigUpdate(conf)
	return sc
}

func (c *sharedConfig) onConfigUpdate(conf *Config) error {
	// update log levels
	c.level.SetLevel(ParseZapLevel(conf.Level))

	// we have to update alla existing component levels
	c.mu.Lock()
	c.config = conf
	for component, atomicLevel := range c.componentLevels {
		effectiveLevel := c.level.Level()
		parts := strings.Split(component, ".")
	confSearch:
		for len(parts) > 0 {
			search := strings.Join(parts, ".")
			if compLevel, ok := conf.ComponentLevels[search]; ok {
				effectiveLevel = ParseZapLevel(compLevel)
				break confSearch
			}
			parts = parts[:len(parts)-1]
		}
		atomicLevel.SetLevel(effectiveLevel)
	}
	c.mu.Unlock()
	return nil
}

// ensure we have an atomic level in the map representing the full component path
// this makes it possible to update the log level after the fact
func (c *sharedConfig) ComponentLevel(component string) zap.AtomicLevel {
	c.mu.Lock()
	defer c.mu.Unlock()
	if compLevel, ok := c.componentLevels[component]; ok {
		return compLevel
	}

	// search up the hierarchy to find the first level that is set
	atomicLevel := zap.NewAtomicLevelAt(c.level.Level())
	c.componentLevels[component] = atomicLevel
	parts := strings.Split(component, ".")
	for len(parts) > 0 {
		search := strings.Join(parts, ".")
		if compLevel, ok := c.config.ComponentLevels[search]; ok {
			atomicLevel.SetLevel(ParseZapLevel(compLevel))
			return atomicLevel
		}
		parts = parts[:len(parts)-1]
	}
	return atomicLevel
}

type zapConfig struct {
	conf          *Config
	sc            *sharedConfig
	writeEnablers *xsync.MapOf[string, *zaputil.WriteEnabler]
	tap           *zaputil.WriteEnabler
}

type ZapLoggerOption func(*zapConfig)

func WithTap(tap *zaputil.WriteEnabler) ZapLoggerOption {
	return func(zc *zapConfig) {
		zc.tap = tap
	}
}

type ZapLogger interface {
	Logger
	ToZap() *zap.SugaredLogger
}

type zapLogger[T zaputil.Encoder[T]] struct {
	zap *zap.SugaredLogger
	*zapConfig
	enc       T
	component string
	deferred  []*zaputil.Deferrer
	sampler   *zaputil.Sampler
}

func NewZapLogger(conf *Config, opts ...ZapLoggerOption) (ZapLogger, error) {
	zap := zap.New(nil).WithOptions(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)).Sugar()

	zc := &zapConfig{
		conf:          conf,
		sc:            newSharedConfig(conf),
		writeEnablers: xsync.NewMapOf[string, *zaputil.WriteEnabler](),
		tap:           zaputil.NewDiscardWriteEnabler(),
	}
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
	} else {
		return newZapLogger(zap, zc, zaputil.NewDevelopmentEncoder(), sampler), nil
	}
}

func newZapLogger[T zaputil.Encoder[T]](zap *zap.SugaredLogger, zc *zapConfig, enc T, sampler *zaputil.Sampler) ZapLogger {
	l := &zapLogger[T]{
		zap:       zap,
		zapConfig: zc,
		enc:       enc,
		sampler:   sampler,
	}
	l.zap = l.ToZap()
	return l
}

func (l *zapLogger[T]) ToZap() *zap.SugaredLogger {
	console, _ := l.writeEnablers.LoadOrCompute(l.component, func() *zaputil.WriteEnabler {
		return zaputil.NewWriteEnabler(os.Stderr, l.sc.ComponentLevel(l.component))
	})

	c := l.enc.Core(console, l.tap)
	for i := range l.deferred {
		c = zaputil.NewDeferredValueCore(c, l.deferred[i])
	}
	if l.sampler != nil {
		c = zaputil.NewSamplerCore(c, l.sampler)
	}

	return l.zap.WithOptions(zap.WrapCore(func(zapcore.Core) zapcore.Core { return c }))
}

func (l *zapLogger[T]) Debugw(msg string, keysAndValues ...interface{}) {
	l.zap.Debugw(msg, keysAndValues...)
}

func (l *zapLogger[T]) Infow(msg string, keysAndValues ...interface{}) {
	l.zap.Infow(msg, keysAndValues...)
}

func (l *zapLogger[T]) Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.zap.Warnw(msg, keysAndValues...)
}

func (l *zapLogger[T]) Errorw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.zap.Errorw(msg, keysAndValues...)
}

func (l *zapLogger[T]) WithValues(keysAndValues ...interface{}) Logger {
	dup := *l
	dup.enc = dup.enc.WithValues(keysAndValues...)
	dup.zap = dup.ToZap()
	return &dup
}

func (l *zapLogger[T]) WithName(name string) Logger {
	dup := *l
	dup.zap = dup.zap.Named(name)
	return &dup
}

func (l *zapLogger[T]) WithComponent(component string) Logger {
	dup := *l
	dup.zap = dup.zap.Named(component)
	if dup.component == "" {
		dup.component = component
	} else {
		dup.component = dup.component + "." + component
	}
	dup.zap = dup.ToZap()
	return &dup
}

func (l *zapLogger[T]) WithCallDepth(depth int) Logger {
	dup := *l
	dup.zap = dup.zap.WithOptions(zap.AddCallerSkip(depth))
	return &dup
}

func (l *zapLogger[T]) WithItemSampler() Logger {
	if l.conf.ItemSampleSeconds == 0 {
		return l
	}
	dup := *l
	dup.sampler = zaputil.NewSampler(
		time.Duration(l.conf.ItemSampleSeconds)*time.Second,
		l.conf.ItemSampleInitial,
		l.conf.ItemSampleInterval,
	)
	dup.zap = dup.ToZap()
	return &dup
}

func (l *zapLogger[T]) WithoutSampler() Logger {
	dup := *l
	dup.sampler = nil
	dup.zap = dup.ToZap()
	return &dup
}

func (l *zapLogger[T]) WithDeferredValues() (Logger, DeferredFieldResolver) {
	dup := *l
	def, resolve := zaputil.NewDeferrer()
	dup.deferred = append(dup.deferred[0:len(dup.deferred):len(dup.deferred)], def)
	dup.zap = dup.ToZap()
	return &dup, resolve
}

type LogRLogger logr.Logger

func (l LogRLogger) toLogr() logr.Logger {
	if logr.Logger(l).GetSink() == nil {
		return discardLogger
	}
	return logr.Logger(l)
}

func (l LogRLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.toLogr().V(1).Info(msg, keysAndValues...)
}

func (l LogRLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.toLogr().Info(msg, keysAndValues...)
}

func (l LogRLogger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.toLogr().Info(msg, keysAndValues...)
}

func (l LogRLogger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	l.toLogr().Error(err, msg, keysAndValues...)
}

func (l LogRLogger) WithValues(keysAndValues ...interface{}) Logger {
	return LogRLogger(l.toLogr().WithValues(keysAndValues...))
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
	return l, func(args ...any) {}
}
