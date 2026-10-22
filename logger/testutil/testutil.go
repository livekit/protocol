package testutil

import (
	"bytes"
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type TestLogOutput struct {
	Level  string
	TS     float64
	Caller string
	Msg    string
}

type BufferedWriteSyncer struct {
	bytes.Buffer
}

func (t *BufferedWriteSyncer) Unmarshal(v any) error {
	return json.Unmarshal(t.Bytes(), &v)
}

func (t *BufferedWriteSyncer) Sync() error { return nil }

func NewJSONCore(ws zapcore.WriteSyncer, enab zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), ws, enab)
}

// zaputil's own tests import this package, so it must not import zaputil;
// callers build the Tee themselves.
type CoreFactory = func(enab zapcore.LevelEnabler) zapcore.Core

func NewJSONCoreFactory(ws zapcore.WriteSyncer) CoreFactory {
	return func(enab zapcore.LevelEnabler) zapcore.Core {
		return NewJSONCore(ws, enab)
	}
}

// NewObserverCoreFactory records every entry the core is handed.
func NewObserverCoreFactory() (CoreFactory, *observer.ObservedLogs) {
	core, logs := observer.New(zapcore.DebugLevel)
	return func(enab zapcore.LevelEnabler) zapcore.Core {
		return Leveled(core, enab)
	}, logs
}

// Leveled overrides a fixed-level core's level decisions with enab.
func Leveled(core zapcore.Core, enab zapcore.LevelEnabler) zapcore.Core {
	return leveledCore{core, enab}
}

type leveledCore struct {
	zapcore.Core
	enab zapcore.LevelEnabler
}

func (c leveledCore) Enabled(lvl zapcore.Level) bool {
	return c.enab.Enabled(lvl)
}

func (c leveledCore) With(fields []zapcore.Field) zapcore.Core {
	return leveledCore{c.Core.With(fields), c.enab}
}

func (c leveledCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if !c.enab.Enabled(ent.Level) {
		return ce
	}
	return ce.AddCore(ent, c)
}
