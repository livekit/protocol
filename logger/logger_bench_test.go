package logger

import (
	"io"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger/testutil"
	"github.com/livekit/protocol/logger/zaputil"
	"github.com/livekit/protocol/utils/must"
)

var benchSink Logger

type loggerBenchCase struct {
	label string
	conf  *Config
	opts  []ZapLoggerOption
}

func (c loggerBenchCase) logger(b *testing.B) ZapLogger {
	b.Helper()
	return must.Get(NewZapLogger(c.conf, c.opts...))
}

func derivationBenchCases() []loggerBenchCase {
	return []loggerBenchCase{
		{label: "console", conf: &Config{}},
		{label: "json", conf: &Config{JSON: true}},
		{label: "console tee", conf: &Config{}, opts: withDiscardTee()},
		{label: "json tee", conf: &Config{JSON: true}, opts: withDiscardTee()},
	}
}

func withDiscardTee() []ZapLoggerOption {
	return []ZapLoggerOption{WithTee(zaputil.NewTee(testutil.NewJSONCoreFactory(zapcore.AddSync(io.Discard))))}
}

func BenchmarkLoggerWithValues(b *testing.B) {
	for _, c := range derivationBenchCases() {
		b.Run(c.label, func(b *testing.B) {
			root := c.logger(b)

			b.Run("root", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					benchSink = root.WithValues("participant", "PA_1234")
				}
			})

			b.Run("derived", func(b *testing.B) {
				derived := root.WithValues("room", "RM_1234")
				b.ReportAllocs()
				for range b.N {
					benchSink = derived.WithValues("participant", "PA_1234")
				}
			})
		})
	}
}

func BenchmarkLoggerWithDeferredValues(b *testing.B) {
	for _, c := range derivationBenchCases() {
		b.Run(c.label, func(b *testing.B) {
			root := c.logger(b)

			b.Run("derive", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					benchSink, _ = root.WithDeferredValues()
				}
			})

			b.Run("derive and resolve", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					l, resolver := root.WithDeferredValues()
					resolver.Resolve("participant", "PA_1234")
					benchSink = l
				}
			})
		})
	}
}

func BenchmarkLoggerWrite(b *testing.B) {
	silenceStderr(b)

	cases := []loggerBenchCase{
		{label: "console", conf: &Config{Level: "debug"}},
		{label: "json", conf: &Config{JSON: true, Level: "debug"}},
		{label: "level disabled", conf: &Config{Level: "info"}},
		{label: "console tee", conf: &Config{Level: "debug"}, opts: withDiscardTee()},
		{label: "json tee", conf: &Config{JSON: true, Level: "debug"}, opts: withDiscardTee()},
	}

	for _, c := range cases {
		b.Run(c.label, func(b *testing.B) {
			root := c.logger(b)

			b.Run("plain", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					root.Debugw("bench", "participant", "PA_1234")
				}
			})

			b.Run("deferred", func(b *testing.B) {
				l, resolver := root.WithDeferredValues()
				resolver.Resolve("room", "RM_1234")
				b.ReportAllocs()
				for range b.N {
					l.Debugw("bench", "participant", "PA_1234")
				}
			})

			b.Run("deferred flush", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					l, resolver := root.WithDeferredValues()
					l.Debugw("bench", "participant", "PA_1234")
					resolver.Resolve("room", "RM_1234")
				}
			})
		})
	}
}
