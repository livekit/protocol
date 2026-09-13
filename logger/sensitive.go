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
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/livekit/logger"
)

// Sensitivity aliases the schema enum so implementors of the interfaces below
// need only this package.
type Sensitivity = logger.Sensitivity

const (
	SensitivityNone   = logger.Sensitivity_SENSITIVITY_UNSPECIFIED
	SensitivityPII    = logger.Sensitivity_SENSITIVITY_PII
	SensitivitySecret = logger.Sensitivity_SENSITIVITY_SECRET
)

// SensitiveObjectEncoder is a zapcore.ObjectEncoder that can record a value's
// sensitivity instead of having the value destroyed. An encoder that cannot
// record sensitivity does not implement it, and marshalers redact instead.
type SensitiveObjectEncoder interface {
	zapcore.ObjectEncoder

	// SensitiveObjectEncoder returns an encoder that records every value
	// written to it, and to encoders derived from it, at sensitivity s. It
	// returns nil if this encoder must not record values at s, in which case
	// the caller must redact.
	//
	// Implementations must fail closed by returning nil for any level they do
	// not explicitly recognize. The returned encoder may share buffers with the
	// receiver and is only valid until the next write on the receiver.
	SensitiveObjectEncoder(s Sensitivity) zapcore.ObjectEncoder
}

// SensitiveArrayEncoder is the zapcore.ArrayEncoder counterpart to
// SensitiveObjectEncoder.
type SensitiveArrayEncoder interface {
	zapcore.ArrayEncoder

	SensitiveArrayEncoder(s Sensitivity) zapcore.ArrayEncoder
}

// ObjectEncoderFor returns an encoder that records values at sensitivity s, or
// nil if e cannot record them. Hand-written marshalers that carry sensitivity
// annotations should use this rather than asserting directly.
func ObjectEncoderFor(e zapcore.ObjectEncoder, s Sensitivity) zapcore.ObjectEncoder {
	se, ok := e.(SensitiveObjectEncoder)
	if !ok {
		return nil
	}
	return se.SensitiveObjectEncoder(s)
}

// ArrayEncoderFor is the ArrayEncoder counterpart to ObjectEncoderFor.
func ArrayEncoderFor(e zapcore.ArrayEncoder, s Sensitivity) zapcore.ArrayEncoder {
	se, ok := e.(SensitiveArrayEncoder)
	if !ok {
		return nil
	}
	return se.SensitiveArrayEncoder(s)
}
