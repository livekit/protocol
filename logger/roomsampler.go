package logger

import (
	"fmt"
	"math"
	"reflect"

	"github.com/zeebo/xxh3"
	"go.uber.org/zap/zapcore"
)

const RoomKey = "room"

func NewRoomSampler(core zapcore.Core, rate float64) zapcore.Core {
	return &roomSampler{
		Core:      core,
		threshold: math.MaxUint64 / 10000 * uint64(math.Max(math.Min(rate, 1), 0)*10000),
	}
}

type roomSampler struct {
	zapcore.Core

	threshold uint64
	filter    bool
}

var _ zapcore.Core = (*roomSampler)(nil)

func (s *roomSampler) With(fields []zapcore.Field) zapcore.Core {
	write, ok := s.checkSampleField(fields)

	return &roomSampler{
		Core:      s.Core.With(fields),
		threshold: s.threshold,
		filter:    s.filter || (!write && ok),
	}
}

func (s *roomSampler) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if s.Enabled(ent.Level) {
		return ce.AddCore(ent, s)
	}
	return ce
}

func (s *roomSampler) checkSampleField(fields []zapcore.Field) (write, ok bool) {
	v, ok := s.findSampleField(fields)
	if !ok {
		return false, false
	}
	return xxh3.HashString(v) <= s.threshold, true
}

func (s *roomSampler) findSampleField(fields []zapcore.Field) (string, bool) {
	for _, f := range fields {
		if f.Key == RoomKey {
			switch f.Type {
			case zapcore.StringType:
				return f.String, true
			case zapcore.ReflectType:
				rv := reflect.ValueOf(f.Interface)
				if rv.Kind() == reflect.String {
					return rv.String(), true
				}
			case zapcore.StringerType:
				if str, ok := f.Interface.(fmt.Stringer); ok {
					return str.String(), true
				}
			}
			return "", false
		}
	}
	return "", false
}

func (s *roomSampler) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if s.filter {
		return nil
	}
	if write, ok := s.checkSampleField(fields); !write && ok {
		return nil
	}
	return s.Core.Write(entry, fields)
}
