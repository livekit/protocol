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
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"
)

func ObjectSlice[T zapcore.ObjectMarshaler](s []T) zapcore.ArrayMarshaler {
	return objectSlice[T](s)
}

type objectSlice[T zapcore.ObjectMarshaler] []T

func (a objectSlice[T]) MarshalLogArray(e zapcore.ArrayEncoder) error {
	var err error
	for _, o := range a {
		err = multierr.Append(err, e.AppendObject(o))
	}
	return err
}
