// Copyright 2025 LiveKit, Inc.
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

package medialogutils

import "github.com/livekit/protocol/logger"

// OverrideLogger converts errors to warnings
type OverrideLogger struct {
	logger.Logger
}

func NewOverrideLogger(l logger.Logger) *OverrideLogger {
	if l == nil {
		l = logger.GetLogger()
	}

	return &OverrideLogger{
		Logger: l.WithCallDepth(1),
	}
}

func (l *OverrideLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.Logger.Debugw(msg, keysAndValues...)
}

func (l *OverrideLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.Logger.Infow(msg, keysAndValues...)
}

func (l *OverrideLogger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	l.Logger.Warnw(msg, err, keysAndValues...)
}

func (l *OverrideLogger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	l.Logger.Warnw(msg, err, keysAndValues...)
}

func (l *OverrideLogger) WithValues(keysAndValues ...interface{}) logger.Logger {
	return &OverrideLogger{Logger: l.Logger.WithValues(keysAndValues...)}
}

func (l *OverrideLogger) WithUnlikelyValues(keysAndValues ...interface{}) logger.UnlikelyLogger {
	return logger.NewUnlikelyLogger(l, keysAndValues...)
}

func (l *OverrideLogger) WithName(name string) logger.Logger {
	return &OverrideLogger{Logger: l.Logger.WithName(name)}
}

func (l *OverrideLogger) WithComponent(component string) logger.Logger {
	return &OverrideLogger{Logger: l.Logger.WithComponent(component)}
}

func (l *OverrideLogger) WithCallDepth(depth int) logger.Logger {
	return &OverrideLogger{Logger: l.Logger.WithCallDepth(depth)}
}

func (l *OverrideLogger) WithItemSampler() logger.Logger {
	return &OverrideLogger{Logger: l.Logger.WithItemSampler()}
}

func (l *OverrideLogger) WithoutSampler() logger.Logger {
	return &OverrideLogger{Logger: l.Logger.WithoutSampler()}
}

func (l *OverrideLogger) WithDeferredValues() (logger.Logger, logger.DeferredFieldResolver) {
	child, resolver := l.Logger.WithDeferredValues()
	return &OverrideLogger{Logger: child}, resolver
}
