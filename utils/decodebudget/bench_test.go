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

//go:build !race

package decodebudget

import (
	"testing"

	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/utils/protojson"
)

func BenchmarkEstimateTraces(b *testing.B) {
	data, _ := proto.Marshal(calibSpans(2048))
	desc := (&tracev1.TracesData{}).ProtoReflect().Descriptor()
	budget := Proto{MaxFields: 1 << 30, MaxDepth: 100}
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	for b.Loop() {
		_, _ = budget.Estimate(data, desc)
	}
}

func BenchmarkProtoStructSize(b *testing.B) {
	desc := (&tracev1.Span{}).ProtoReflect().Descriptor()
	for b.Loop() {
		_ = protoStructSize(desc)
	}
}

func benchmarkProtoJSON(b *testing.B, msg proto.Message) {
	data, err := protojson.Marshal(msg)
	if err != nil {
		b.Fatal(err)
	}
	desc := msg.ProtoReflect().Descriptor()
	budget := Proto{MaxFields: 1 << 30, MaxDepth: 100}
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	for b.Loop() {
		_, _ = budget.EstimateProtoJSON(data, desc)
	}
}

func BenchmarkEstimateProtoJSONTraces(b *testing.B) {
	benchmarkProtoJSON(b, calibSpans(2048))
}

func BenchmarkEstimateProtoJSONLogs(b *testing.B) {
	benchmarkProtoJSON(b, calibLogs(512))
}
