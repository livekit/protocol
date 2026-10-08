// Copyright 2026 LiveKit, Inc.
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

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/logger/testutil"
	"github.com/livekit/protocol/logger/zaputil"
	"github.com/livekit/protocol/utils/must"
)

type errorLogger interface {
	Errorw(msg string, err error, keysAndValues ...any)
}

func logError(l errorLogger) {
	l.Errorw("msg", nil)
}

func TestOverrideLoggerChildren(t *testing.T) {
	f, logs := testutil.NewObserverCoreFactory()
	zl := must.Get(logger.NewZapLogger(&logger.Config{Level: "debug"}, logger.WithTee(zaputil.NewTee(f))))
	l := NewOverrideLogger(zl)
	deferred, resolver := l.WithDeferredValues()
	resolver.Resolve()

	cases := map[string]errorLogger{
		"NewOverrideLogger":  l,
		"WithValues":         l.WithValues("k", "v"),
		"WithUnlikelyValues": l.WithUnlikelyValues("k", "v"),
		"WithName":           l.WithName("n"),
		"WithComponent":      l.WithComponent("c"),
		"WithCallDepth":      l.WithCallDepth(0),
		"WithItemSampler":    l.WithItemSampler(),
		"WithoutSampler":     l.WithoutSampler(),
		"WithDeferredValues": deferred,
	}
	loggerType := reflect.TypeFor[logger.Logger]()
	unlikelyType := reflect.TypeFor[logger.UnlikelyLogger]()
	for i := range loggerType.NumMethod() {
		m := loggerType.Method(i)
		if out := m.Type.Out; m.Type.NumOut() > 0 && (out(0) == loggerType || out(0) == unlikelyType) {
			require.Contains(t, cases, m.Name, "wrap %s in OverrideLogger and add it here", m.Name)
		}
	}
	for label, c := range cases {
		t.Run(label, func(t *testing.T) {
			logError(c)
			entries := logs.TakeAll()
			require.Len(t, entries, 1)
			require.Equal(t, zapcore.WarnLevel, entries[0].Level)
			require.True(t, strings.HasSuffix(entries[0].Caller.Function, ".logError"), "caller %s", entries[0].Caller.Function)
		})
	}
}
