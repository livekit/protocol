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

type Encoder interface {
	WithValues(kvs ...any) Encoder
	Core(console, json *WriteEnabler) zapcore.Core
}

type DevelopmentEncoder struct {
	console zapcore.Encoder
	json    zapcore.Encoder
}

func NewDevelopmentEncoder() Encoder {
	return &DevelopmentEncoder{
		console: zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		json:    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
	}
}

func (e *DevelopmentEncoder) WithValues(kvs ...any) Encoder {
	clone := *e
	clone.console = encoderWithValues(clone.console, kvs...)
	clone.json = encoderWithValues(clone.json, kvs...)
	return &clone
}

func (e *DevelopmentEncoder) Core(console, json *WriteEnabler) zapcore.Core {
	return zapcore.NewTee(NewEncoderCore(e.console, console), NewEncoderCore(e.json, json))
}

type ProductionEncoder struct {
	json zapcore.Encoder
}

func NewProductionEncoder() Encoder {
	return &ProductionEncoder{
		json: zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
	}
}

func (e *ProductionEncoder) WithValues(kvs ...any) Encoder {
	clone := *e
	clone.json = encoderWithValues(clone.json, kvs...)
	return &clone
}

func (e *ProductionEncoder) Core(console, json *WriteEnabler) zapcore.Core {
	return NewEncoderCore(e.json, console, json)
}
