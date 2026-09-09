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

package zaputil

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func encoderWithValues(enc zapcore.Encoder, kvs ...any) zapcore.Encoder {
	clone := enc.Clone()
	for i := 1; i < len(kvs); i += 2 {
		if key, ok := kvs[i-1].(string); ok {
			zap.Any(key, kvs[i]).AddTo(clone)
		}
	}
	return clone
}

// Field selection must match encoderWithValues, which cannot share this loop without allocating.
func valueFields(kvs ...any) []zapcore.Field {
	fields := make([]zapcore.Field, 0, len(kvs)/2)
	for i := 1; i < len(kvs); i += 2 {
		if key, ok := kvs[i-1].(string); ok {
			fields = append(fields, zap.Any(key, kvs[i]))
		}
	}
	return fields
}

type Encoder[T any] interface {
	WithValues(kvs ...any) T
	Core(out *WriteEnabler) zapcore.Core
}

type DevelopmentEncoder struct {
	console zapcore.Encoder
}

func NewDevelopmentEncoder() DevelopmentEncoder {
	return DevelopmentEncoder{
		console: zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	}
}

func (e DevelopmentEncoder) WithValues(kvs ...any) DevelopmentEncoder {
	e.console = encoderWithValues(e.console, kvs...)
	return e
}

func (e DevelopmentEncoder) Core(out *WriteEnabler) zapcore.Core {
	return NewEncoderCore(e.console, out)
}

type ProductionEncoder struct {
	json zapcore.Encoder
}

func NewProductionEncoder() ProductionEncoder {
	return ProductionEncoder{
		json: zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
	}
}

func (e ProductionEncoder) WithValues(kvs ...any) ProductionEncoder {
	e.json = encoderWithValues(e.json, kvs...)
	return e
}

func (e ProductionEncoder) Core(out *WriteEnabler) zapcore.Core {
	return NewEncoderCore(e.json, out)
}
