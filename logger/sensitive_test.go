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

package logger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

// recorded is a value the sink accepted, plus the sensitivity it was tagged
// with. A redacted field arrives as an ordinary untagged string.
type recorded struct {
	v    any
	sens logger.Sensitivity
}

// recorder is a sink that records values under dot-joined key paths, mirroring
// how the columnar sink in backend-common flattens nesting. max is the highest
// sensitivity it will accept; anything above that is refused, so the marshaller
// falls back to redaction.
type recorder struct {
	t    *testing.T
	max  logger.Sensitivity
	vals map[string]recorded
}

func newRecorder(t *testing.T, max logger.Sensitivity) (*recorder, *recEnc) {
	r := &recorder{t: t, max: max, vals: map[string]recorded{}}
	return r, &recEnc{ObjectEncoder: zapcore.NewMapObjectEncoder(), r: r}
}

func (r *recorder) fields(t *testing.T, m zapcore.ObjectMarshaler, e *recEnc) map[string]recorded {
	t.Helper()
	require.NoError(t, m.MarshalLogObject(e))
	return r.vals
}

func (r *recorder) put(path string, v any, s logger.Sensitivity) {
	r.vals[path] = recorded{v, s}
}

func join(prefix, key string) string {
	if prefix == "" {
		return key
	}
	return prefix + "." + key
}

// recEnc embeds a MapObjectEncoder to satisfy the ObjectEncoder methods proto.go
// never calls, and overrides the eight it does.
type recEnc struct {
	zapcore.ObjectEncoder
	r      *recorder
	prefix string
	sens   logger.Sensitivity
}

var _ logger.SensitiveObjectEncoder = (*recEnc)(nil)

func (e *recEnc) SensitiveObjectEncoder(s logger.Sensitivity) zapcore.ObjectEncoder {
	if s > e.r.max {
		return nil
	}
	if s <= e.sens {
		return e
	}
	return &recEnc{ObjectEncoder: e.ObjectEncoder, r: e.r, prefix: e.prefix, sens: s}
}

func (e *recEnc) add(key string, v any) { e.r.put(join(e.prefix, key), v, e.sens) }

func (e *recEnc) AddString(key, v string)          { e.add(key, v) }
func (e *recEnc) AddInt(key string, v int)         { e.add(key, v) }
func (e *recEnc) AddInt64(key string, v int64)     { e.add(key, v) }
func (e *recEnc) AddUint64(key string, v uint64)   { e.add(key, v) }
func (e *recEnc) AddFloat64(key string, v float64) { e.add(key, v) }
func (e *recEnc) AddBool(key string, v bool)       { e.add(key, v) }

func (e *recEnc) AddObject(key string, m zapcore.ObjectMarshaler) error {
	return m.MarshalLogObject(&recEnc{
		ObjectEncoder: e.ObjectEncoder,
		r:             e.r,
		prefix:        join(e.prefix, key),
		sens:          e.sens,
	})
}

func (e *recEnc) AddArray(key string, m zapcore.ArrayMarshaler) error {
	return m.MarshalLogArray(&recArr{r: e.r, prefix: join(e.prefix, key), sens: e.sens})
}

// recArr records array elements under an indexed path.
type recArr struct {
	zapcore.ArrayEncoder
	r      *recorder
	prefix string
	sens   logger.Sensitivity
	n      int
}

func (e *recArr) path() string {
	p := fmt.Sprintf("%s.%d", e.prefix, e.n)
	e.n++
	return p
}

func (e *recArr) append(v any) { e.r.put(e.path(), v, e.sens) }

func (e *recArr) AppendString(v string)   { e.append(v) }
func (e *recArr) AppendBool(v bool)       { e.append(v) }
func (e *recArr) AppendInt64(v int64)     { e.append(v) }
func (e *recArr) AppendUint64(v uint64)   { e.append(v) }
func (e *recArr) AppendFloat64(v float64) { e.append(v) }

func (e *recArr) AppendObject(m zapcore.ObjectMarshaler) error {
	return m.MarshalLogObject(&recEnc{
		ObjectEncoder: zapcore.NewMapObjectEncoder(),
		r:             e.r,
		prefix:        e.path(),
		sens:          e.sens,
	})
}

// The remaining ArrayEncoder methods are unreachable from proto.go. Reaching one
// means the marshaller grew a path this double does not model.
func (e *recArr) unsupported(name string) {
	e.r.t.Fatalf("recArr.%s: unmodelled encoder method", name)
}

func (e *recArr) AppendArray(zapcore.ArrayMarshaler) error { e.unsupported("AppendArray"); return nil }
func (e *recArr) AppendReflected(any) error                { e.unsupported("AppendReflected"); return nil }
func (e *recArr) AppendByteString([]byte)                  { e.unsupported("AppendByteString") }
func (e *recArr) AppendComplex128(complex128)              { e.unsupported("AppendComplex128") }
func (e *recArr) AppendComplex64(complex64)                { e.unsupported("AppendComplex64") }
func (e *recArr) AppendDuration(d time.Duration)           { e.unsupported("AppendDuration") }
func (e *recArr) AppendTime(t time.Time)                   { e.unsupported("AppendTime") }
func (e *recArr) AppendFloat32(float32)                    { e.unsupported("AppendFloat32") }
func (e *recArr) AppendInt(int)                            { e.unsupported("AppendInt") }
func (e *recArr) AppendInt32(int32)                        { e.unsupported("AppendInt32") }
func (e *recArr) AppendInt16(int16)                        { e.unsupported("AppendInt16") }
func (e *recArr) AppendInt8(int8)                          { e.unsupported("AppendInt8") }
func (e *recArr) AppendUint(uint)                          { e.unsupported("AppendUint") }
func (e *recArr) AppendUint32(uint32)                      { e.unsupported("AppendUint32") }
func (e *recArr) AppendUint16(uint16)                      { e.unsupported("AppendUint16") }
func (e *recArr) AppendUint8(uint8)                        { e.unsupported("AppendUint8") }
func (e *recArr) AppendUintptr(uintptr)                    { e.unsupported("AppendUintptr") }

func requireRecorded(t *testing.T, got map[string]recorded, path string, v any, s logger.Sensitivity) {
	t.Helper()
	r, ok := got[path]
	require.True(t, ok, "missing %q in %v", path, got)
	require.Equal(t, v, r.v, "value of %q", path)
	require.Equal(t, s, r.sens, "sensitivity of %q", path)
}

// A sink that can tag stores the real value rather than the redaction, even for
// a field whose redact_format would otherwise produce a size summary.
func TestProtoStoresPIIWhenEncoderCanTag(t *testing.T) {
	r, e := newRecorder(t, logger.SensitivityPII)
	got := r.fields(t, logger.Proto(&livekit.ParticipantInfo{
		Sid: "PA_x", Identity: "alice", Name: "Alice", Metadata: `{"a":1}`,
	}), e)

	requireRecorded(t, got, "sid", "PA_x", logger.SensitivityNone)
	requireRecorded(t, got, "identity", "alice", logger.SensitivityNone)
	requireRecorded(t, got, "name", "Alice", logger.SensitivityPII)
	requireRecorded(t, got, "metadata", `{"a":1}`, logger.SensitivityPII)
}

// A sink that cannot tag is unaffected: Proto still redacts.
func TestProtoRedactsWhenEncoderCannotTag(t *testing.T) {
	r, e := newRecorder(t, logger.SensitivityNone)
	got := r.fields(t, logger.Proto(&livekit.ParticipantInfo{
		Sid: "PA_x", Name: "Alice",
	}), e)

	requireRecorded(t, got, "sid", "PA_x", logger.SensitivityNone)
	requireRecorded(t, got, "name", "<redacted>", logger.SensitivityNone)
}

// SECRET is never stored, whatever the sink offers.
func TestProtoNeverStoresSecret(t *testing.T) {
	r, e := newRecorder(t, logger.SensitivityPII)
	got := r.fields(t, logger.Proto(&livekit.S3Upload{
		AccessKey: "AK", Secret: "SK", SessionToken: "ST",
		AssumeRoleArn: "arn:aws:iam::1:role/r", AssumeRoleExternalId: "EID",
		Region: "us-east-1", Bucket: "b",
	}), e)

	requireRecorded(t, got, "region", "us-east-1", logger.SensitivityNone)
	requireRecorded(t, got, "bucket", "b", logger.SensitivityNone)
	requireRecorded(t, got, "assumeRoleArn", "arn:aws:iam::1:role/r", logger.SensitivityPII)
	for _, k := range []string{"accessKey", "secret", "sessionToken", "assumeRoleExternalID"} {
		requireRecorded(t, got, k, "<redacted>", logger.SensitivityNone)
	}
}

// A sensitive map has no per-entry annotation, so its entries inherit the
// field's level through the encoder.
func TestSensitivityInheritsThroughNesting(t *testing.T) {
	r, e := newRecorder(t, logger.SensitivityPII)
	got := r.fields(t, logger.Proto(&livekit.ParticipantInfo{
		Sid: "PA_x", Attributes: map[string]string{"a": "1", "b": "2"},
	}), e)

	requireRecorded(t, got, "sid", "PA_x", logger.SensitivityNone)
	requireRecorded(t, got, "attributes.a", "1", logger.SensitivityPII)
	requireRecorded(t, got, "attributes.b", "2", logger.SensitivityPII)
}

// Mixed levels inside a repeated message: the PII field is stored tagged while
// the SECRET sibling is still redacted.
func TestSecretUnderPIIStillRedacted(t *testing.T) {
	r, e := newRecorder(t, logger.SensitivityPII)
	got := r.fields(t, logger.Proto(&livekit.JoinResponse{
		IceServers: []*livekit.ICEServer{{
			Urls: []string{"turn:x"}, Username: "u", Credential: "c",
		}},
	}), e)

	requireRecorded(t, got, "iceServers.0.username", "u", logger.SensitivityPII)
	requireRecorded(t, got, "iceServers.0.credential", "<redacted>", logger.SensitivityNone)
}

// Summary mode tags the scalars it keeps and leaves the non-disclosing counts
// untagged.
func TestSummaryTagsScalars(t *testing.T) {
	r, e := newRecorder(t, logger.SensitivityPII)
	got := r.fields(t, logger.ProtoWithLimit(&livekit.ParticipantInfo{
		Sid: "PA_x", Name: "Alice", Attributes: map[string]string{"a": "1"},
	}, 1), e)

	requireRecorded(t, got, "truncatedProto", "livekit.ParticipantInfo", logger.SensitivityNone)
	requireRecorded(t, got, "sid", "PA_x", logger.SensitivityNone)
	requireRecorded(t, got, "name", "Alice", logger.SensitivityPII)
	requireRecorded(t, got, "attributesCount", 1, logger.SensitivityNone)
	require.NotContains(t, got, "attributes.a")
}
