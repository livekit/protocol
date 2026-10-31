// The race detector adds allocations of its own, so measured decode sizes only
// match production builds without it.

//go:build !race

package decodebudget

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	resourcev1 "go.opentelemetry.io/proto/otlp/resource/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils/protojson"
)

func measureDecode(t *testing.T, data []byte, newMsg func() proto.Message) uint64 {
	var best uint64
	for i := 0; i < 3; i++ {
		var before, after runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&before)
		m := newMsg()
		require.NoError(t, proto.Unmarshal(data, m))
		runtime.ReadMemStats(&after)
		runtime.KeepAlive(m)
		if d := after.TotalAlloc - before.TotalAlloc; best == 0 || d < best {
			best = d
		}
	}
	return best
}

func str(s string) *commonv1.AnyValue {
	return &commonv1.AnyValue{Value: &commonv1.AnyValue_StringValue{StringValue: s}}
}

func calibSpans(n int) *tracev1.TracesData {
	ss := &tracev1.ScopeSpans{Scope: &commonv1.InstrumentationScope{Name: "livekit-agents"}}
	for i := range n {
		attrs := []*commonv1.KeyValue{}
		for j := range 20 {
			attrs = append(attrs, &commonv1.KeyValue{Key: fmt.Sprintf("lk.attr_%d", j), Value: str(strings.Repeat("v", 10+j))})
		}
		attrs = append(attrs, &commonv1.KeyValue{Key: "n", Value: &commonv1.AnyValue{Value: &commonv1.AnyValue_IntValue{IntValue: int64(i)}}})
		ss.Spans = append(ss.Spans, &tracev1.Span{
			TraceId: make([]byte, 16), SpanId: make([]byte, 8), Name: "llm_node",
			StartTimeUnixNano: uint64(i), EndTimeUnixNano: uint64(i + 1), Attributes: attrs,
			Events: []*tracev1.Span_Event{{Name: "e", Attributes: attrs[:3]}},
		})
	}
	return &tracev1.TracesData{ResourceSpans: []*tracev1.ResourceSpans{{
		Resource:   &resourcev1.Resource{Attributes: []*commonv1.KeyValue{{Key: "service.name", Value: str("agent")}}},
		ScopeSpans: []*tracev1.ScopeSpans{ss},
	}}}
}

func calibLogs(n int) *logsv1.LogsData {
	sl := &logsv1.ScopeLogs{}
	for i := range n {
		nested := &commonv1.AnyValue{Value: &commonv1.AnyValue_KvlistValue{KvlistValue: &commonv1.KeyValueList{Values: []*commonv1.KeyValue{
			{Key: "role", Value: str("assistant")},
			{Key: "content", Value: &commonv1.AnyValue{Value: &commonv1.AnyValue_ArrayValue{ArrayValue: &commonv1.ArrayValue{Values: []*commonv1.AnyValue{str(strings.Repeat("hello ", 40))}}}}},
			{Key: "ttft", Value: &commonv1.AnyValue{Value: &commonv1.AnyValue_DoubleValue{DoubleValue: 0.4}}},
		}}}}
		sl.LogRecords = append(sl.LogRecords, &logsv1.LogRecord{
			TimeUnixNano: uint64(i), SeverityText: "INFO", Body: str("chat item"),
			Attributes: []*commonv1.KeyValue{{Key: "job_id", Value: str("AJ_FuWSAg3gYutS")}, {Key: "chat.item", Value: nested}},
			TraceId:    make([]byte, 16), SpanId: make([]byte, 8),
		})
	}
	return &logsv1.LogsData{ResourceLogs: []*logsv1.ResourceLogs{{ScopeLogs: []*logsv1.ScopeLogs{sl}}}}
}

func calibHistograms(n int) *metricsv1.MetricsData {
	sm := &metricsv1.ScopeMetrics{}
	for i := range n {
		sm.Metrics = append(sm.Metrics, &metricsv1.Metric{Name: "lat", Data: &metricsv1.Metric_Histogram{Histogram: &metricsv1.Histogram{
			DataPoints: []*metricsv1.HistogramDataPoint{{
				TimeUnixNano: uint64(i), Count: 10, BucketCounts: make([]uint64, 32), ExplicitBounds: make([]float64, 31),
			}},
		}}})
	}
	return &metricsv1.MetricsData{ResourceMetrics: []*metricsv1.ResourceMetrics{{ScopeMetrics: []*metricsv1.ScopeMetrics{sm}}}}
}

func calibParticipants() *livekit.ListParticipantsResponse {
	participants := &livekit.ListParticipantsResponse{}
	for i := range 200 {
		p := &livekit.ParticipantInfo{Sid: fmt.Sprintf("PA_%d", i), Identity: "user", Attributes: map[string]string{}}
		for j := range 10 {
			p.Attributes[fmt.Sprintf("k%d", j)] = "value"
		}
		p.Tracks = []*livekit.TrackInfo{{Sid: "TR_x", Name: "mic", Layers: []*livekit.VideoLayer{{}, {}}}}
		participants.Participants = append(participants.Participants, p)
	}
	return participants
}

func emptyChildren(field protowire.Number, n int) []byte {
	one := protowire.AppendBytes(protowire.AppendTag(nil, field, protowire.BytesType), nil)
	return []byte(strings.Repeat(string(one), n))
}

func TestProtoEstimateCalibration(t *testing.T) {
	// ScopeSpans{spans: n×{}} inside ResourceSpans inside TracesData.
	emptySpans := func(n int) []byte {
		scope := emptyChildren(2, n)
		rs := protowire.AppendBytes(protowire.AppendTag(nil, 2, protowire.BytesType), scope)
		return protowire.AppendBytes(protowire.AppendTag(nil, 1, protowire.BytesType), rs)
	}
	emptyAttrs := func(n int) []byte {
		// LogsData > ResourceLogs > ScopeLogs > LogRecord{attributes: n×{}}
		rec := emptyChildren(6, n)
		sl := protowire.AppendBytes(protowire.AppendTag(nil, 2, protowire.BytesType), rec)
		rl := protowire.AppendBytes(protowire.AppendTag(nil, 2, protowire.BytesType), sl)
		return protowire.AppendBytes(protowire.AppendTag(nil, 1, protowire.BytesType), rl)
	}
	participants := calibParticipants()
	cases := []struct {
		name   string
		data   []byte
		newMsg func() proto.Message
	}{
		{"traces 512", must(proto.Marshal(calibSpans(512))), func() proto.Message { return &tracev1.TracesData{} }},
		{"logs 512", must(proto.Marshal(calibLogs(512))), func() proto.Message { return &logsv1.LogsData{} }},
		{"histograms 512", must(proto.Marshal(calibHistograms(512))), func() proto.Message { return &metricsv1.MetricsData{} }},
		{"empty spans 1M", emptySpans(1 << 20), func() proto.Message { return &tracev1.TracesData{} }},
		{"empty attributes 1M", emptyAttrs(1 << 20), func() proto.Message { return &logsv1.LogsData{} }},
		{"participants maps", must(proto.Marshal(participants)), func() proto.Message { return &livekit.ListParticipantsResponse{} }},
	}
	budget := Proto{MaxFields: 1 << 30, MaxDepth: 100}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			est, err := budget.Estimate(tc.data, tc.newMsg().ProtoReflect().Descriptor())
			require.NoError(t, err)
			got := measureDecode(t, tc.data, tc.newMsg)
			t.Logf("wire=%d measured=%d estimate=%d ratio=%.2f", len(tc.data), got, est, float64(est)/float64(got))
			require.GreaterOrEqual(t, uint64(est), got, "the estimate must not undercount")
			require.LessOrEqual(t, float64(est), 1.15*float64(got), "the estimate should stay close enough to admit real payloads")
		})
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func measureJSONDecode(t *testing.T, data []byte, newMsg func() proto.Message) uint64 {
	var best uint64
	for i := 0; i < 3; i++ {
		var before, after runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&before)
		m := newMsg()
		require.NoError(t, protojson.Unmarshal(data, m))
		runtime.ReadMemStats(&after)
		runtime.KeepAlive(m)
		if d := after.TotalAlloc - before.TotalAlloc; best == 0 || d < best {
			best = d
		}
	}
	return best
}

func TestProtoEstimateJSONCalibration(t *testing.T) {
	emptySpans := `{"resourceSpans":[{"scopeSpans":[{"spans":[` + strings.TrimSuffix(strings.Repeat("{},", 1<<18), ",") + `]}]}]}`
	emptyValues := `{"resourceLogs":[{"scopeLogs":[{"logRecords":[{"body":{"arrayValue":{"values":[` + strings.TrimSuffix(strings.Repeat("{},", 1<<18), ",") + `]}}}]}]}]}`
	cases := []struct {
		name   string
		data   []byte
		newMsg func() proto.Message
	}{
		{"traces 512", must(protojson.Marshal(calibSpans(512))), func() proto.Message { return &tracev1.TracesData{} }},
		{"logs 512", must(protojson.Marshal(calibLogs(512))), func() proto.Message { return &logsv1.LogsData{} }},
		{"histograms 512", must(protojson.Marshal(calibHistograms(512))), func() proto.Message { return &metricsv1.MetricsData{} }},
		{"empty spans", []byte(emptySpans), func() proto.Message { return &tracev1.TracesData{} }},
		{"empty values", []byte(emptyValues), func() proto.Message { return &logsv1.LogsData{} }},
		{"participants maps", must(protojson.Marshal(calibParticipants())), func() proto.Message { return &livekit.ListParticipantsResponse{} }},
	}
	budget := Proto{MaxFields: 1 << 30, MaxDepth: 100}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			est, err := budget.EstimateJSON(tc.data, tc.newMsg().ProtoReflect().Descriptor())
			require.NoError(t, err)
			got := measureJSONDecode(t, tc.data, tc.newMsg)
			t.Logf("json=%d measured=%d estimate=%d ratio=%.2f", len(tc.data), got, est, float64(est)/float64(got))
			require.GreaterOrEqual(t, uint64(est), got, "the estimate must not undercount")
			require.LessOrEqual(t, float64(est), 1.15*float64(got), "the estimate should stay close enough to admit real payloads")
		})
	}
}
