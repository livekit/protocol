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
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-logr/logr"
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
	WithTap(we *zaputil.WriteEnabler) Logger
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

type ZapLogger struct {
	zap           *zap.SugaredLogger
	conf          *Config
	sc            *sharedConfig
	console, json *zaputil.WriteEnabler
	enc           zaputil.Encoder
	name          string
	component     string
	deferred      []*zaputil.Deferrer
	sampler       *zaputil.Sampler
	callerSkip    int
}

func NewZapLogger(conf *Config) (*ZapLogger, error) {
	sc := newSharedConfig(conf)

	var enc zaputil.Encoder
	if conf.JSON {
		enc = zaputil.NewProductionEncoder()
	} else {
		enc = zaputil.NewDevelopmentEncoder()
	}

	l := &ZapLogger{
		conf:    conf,
		sc:      sc,
		console: zaputil.NewWriteEnabler(os.Stderr, sc.level),
		json:    zaputil.NewDiscardWriteEnabler(),
		enc:     enc,
	}

	if conf.Sample {
		var initial = 20
		var interval = 100
		if conf.ItemSampleInitial != 0 {
			initial = conf.ItemSampleInitial
		}
		if conf.ItemSampleInterval != 0 {
			interval = conf.ItemSampleInterval
		}
		l.sampler = zaputil.NewSampler(time.Second, initial, interval)
	}

	l.zap = l.ToZap()
	return l, nil
}

func (l *ZapLogger) ToZap() *zap.SugaredLogger {
	c := l.enc.Core(l.console, l.json)
	for i := range l.deferred {
		c = zaputil.NewDeferredValueCore(c, l.deferred[i])
	}
	if l.sampler != nil {
		c = zaputil.NewSamplerCore(c, l.sampler)
	}

	zl := zap.New(c, zap.AddCaller(), zap.AddCallerSkip(l.callerSkip))

	if l.name == "" || l.component == "" {
		zl = zl.Named(l.name + l.component)
	} else {
		zl = zl.Named(l.name + "." + l.component)
	}

	return zl.Sugar()
}

func (l *ZapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.zap.Debugw(msg, keysAndValues...)
}

func (l *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.zap.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.zap.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.zap.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) WithValues(keysAndValues ...interface{}) Logger {
	dup := *l
	dup.enc = dup.enc.WithValues(keysAndValues...)
	dup.zap = dup.ToZap()
	return &dup
}

func (l *ZapLogger) WithName(name string) Logger {
	dup := *l
	if dup.name == "" {
		dup.name = name
	} else {
		dup.name = dup.name + "." + name
	}
	dup.zap = dup.ToZap()
	return &dup
}

func (l *ZapLogger) WithComponent(component string) Logger {
	dup := *l
	if dup.component == "" {
		dup.component = component
	} else {
		dup.component = dup.component + "." + component
	}
	dup.console = zaputil.NewWriteEnabler(os.Stderr, l.sc.ComponentLevel(dup.component))
	dup.zap = dup.ToZap()
	return &dup
}

func (l *ZapLogger) WithCallDepth(depth int) Logger {
	dup := *l
	dup.callerSkip = depth
	dup.zap = dup.ToZap()
	return &dup
}

func (l *ZapLogger) WithItemSampler() Logger {
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

func (l *ZapLogger) WithoutSampler() Logger {
	dup := *l
	dup.sampler = nil
	dup.zap = dup.ToZap()
	return &dup
}

func (l *ZapLogger) WithDeferredValues() (Logger, DeferredFieldResolver) {
	dup := *l
	def, resolve := zaputil.NewDeferrer()
	dup.deferred = append(dup.deferred[0:len(dup.deferred):len(dup.deferred)], def)
	dup.zap = dup.ToZap()
	return &dup, resolve
}

func (l *ZapLogger) WithTap(we *zaputil.WriteEnabler) Logger {
	dup := *l
	dup.json = we
	dup.zap = dup.ToZap()
	return &dup
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

func (l LogRLogger) WithTap(we *zaputil.WriteEnabler) Logger {
	return l
}
