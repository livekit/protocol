// The race detector adds allocations of its own, so measured decode sizes only
// match production builds without it.

//go:build !race

package decodebudget

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	resourcev1 "go.opentelemetry.io/proto/otlp/resource/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/livekit/agent"
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

func calibEnumBytesLogs(n int) *logsv1.LogsData {
	sl := &logsv1.ScopeLogs{}
	for i := range n {
		sl.LogRecords = append(sl.LogRecords, &logsv1.LogRecord{
			TimeUnixNano: uint64(1763059861000000000 + i), SeverityNumber: logsv1.SeverityNumber_SEVERITY_NUMBER_INFO,
			TraceId: []byte("0123456789abcdef"), SpanId: []byte("01234567"),
			Body: &commonv1.AnyValue{Value: &commonv1.AnyValue_BytesValue{BytesValue: []byte("some bytes value")}},
		})
	}
	return &logsv1.LogsData{ResourceLogs: []*logsv1.ResourceLogs{{ScopeLogs: []*logsv1.ScopeLogs{sl}}}}
}

func calibOptionalPoints(n int) *metricsv1.MetricsData {
	h := &metricsv1.Histogram{}
	for i := range n {
		v := float64(i)
		h.DataPoints = append(h.DataPoints, &metricsv1.HistogramDataPoint{Sum: &v, Min: &v, Max: &v})
	}
	return &metricsv1.MetricsData{ResourceMetrics: []*metricsv1.ResourceMetrics{{ScopeMetrics: []*metricsv1.ScopeMetrics{{
		Metrics: []*metricsv1.Metric{{Name: "h", Data: &metricsv1.Metric_Histogram{Histogram: h}}},
	}}}}}
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

// calibOneEntryMaps is n agent dispatches with one attribute each, where the
// map header and first group outweigh the entry.
func calibOneEntryMaps(n int) *livekit.RoomConfiguration {
	cfg := &livekit.RoomConfiguration{}
	for range n {
		cfg.Agents = append(cfg.Agents, &livekit.RoomAgentDispatch{Attributes: map[string]string{"role": "assistant"}})
	}
	return cfg
}

// calibLargeMap is one map of n entries, past several rounds of table splits.
func calibLargeMap(n int) *livekit.ParticipantInfo {
	p := &livekit.ParticipantInfo{Attributes: map[string]string{}}
	for i := range n {
		p.Attributes[fmt.Sprintf("attr-%d", i)] = "value"
	}
	return p
}

// valuelessMapEntries is n sub_handles entries with a key and no value; the
// decoder allocates a value message for each anyway.
func valuelessMapEntries(n int) []byte {
	handles := num(&livekit.DataTrackSubscriberHandles{}, "sub_handles")
	var b []byte
	for i := range n {
		key := protowire.AppendVarint(protowire.AppendTag(nil, 1, protowire.VarintType), uint64(i))
		b = protowire.AppendBytes(protowire.AppendTag(b, handles, protowire.BytesType), key)
	}
	return b
}

func emptyChildren(field protowire.Number, n int) []byte {
	one := protowire.AppendBytes(protowire.AppendTag(nil, field, protowire.BytesType), nil)
	return []byte(strings.Repeat(string(one), n))
}

func field(num protowire.Number, payload []byte) []byte {
	return protowire.AppendBytes(protowire.AppendTag(nil, num, protowire.BytesType), payload)
}

// inLogRecord wraps raw LogRecord fields as LogsData.
func inLogRecord(rec []byte) []byte {
	return field(1, field(2, field(2, rec)))
}

// inHistogram wraps raw HistogramDataPoint messages as MetricsData.
func inHistogram(points []byte) []byte {
	return field(1, field(2, field(2, field(9, points))))
}

func packedRuns(k int) []byte {
	one := field(6, protowire.AppendFixed64(nil, 1)) // bucket_counts
	return inHistogram(field(1, []byte(strings.Repeat(string(one), k))))
}

func optionalPoints(n int) []byte {
	var p []byte
	for _, num := range []protowire.Number{5, 11, 12} { // sum, min, max
		p = protowire.AppendFixed64(protowire.AppendTag(p, num, protowire.Fixed64Type), 1)
	}
	return inHistogram([]byte(strings.Repeat(string(field(1, p)), n)))
}

func mismatchedLength(n int) []byte {
	// time_unix_nano is fixed64 on the wire.
	return inLogRecord(field(1, make([]byte, n)))
}

func mismatchedVarints(n int) []byte {
	// severity_text is a string.
	one := protowire.AppendVarint(protowire.AppendTag(nil, 3, protowire.VarintType), 1)
	return inLogRecord([]byte(strings.Repeat(string(one), n)))
}

// mismatchedOneof sends AnyValue.bytes_value as a varint n times; the decoder
// allocates and drops a oneof wrapper for each.
func mismatchedOneof(n int) []byte {
	one := protowire.AppendVarint(protowire.AppendTag(nil, num(&commonv1.AnyValue{}, "bytes_value"), protowire.VarintType), 0)
	return inLogRecord(field(num(&logsv1.LogRecord{}, "body"), []byte(strings.Repeat(string(one), n))))
}

// mismatchedOneofSwitching is n × rpc_response{compressed_payload as varint ×256}
// rpc_ack{}; switching DataPacket.value makes every RpcResponse a fresh message.
func mismatchedOneofSwitching(n int) []byte {
	payload := protowire.AppendVarint(protowire.AppendTag(nil, num(&livekit.RpcResponse{}, "compressed_payload"), protowire.VarintType), 0)
	packet := &livekit.DataPacket{}
	unit := field(num(packet, "rpc_response"), []byte(strings.Repeat(string(payload), 256)))
	unit = append(unit, field(num(packet, "rpc_ack"), nil)...)
	return []byte(strings.Repeat(string(unit), n))
}

// mismatchedMaps sends ParticipantInfo.attributes as a varint in each of n
// participants; the decoder makes the map before it checks the wire type.
func mismatchedMaps(n int) []byte {
	attrs := protowire.AppendVarint(protowire.AppendTag(nil, num(&livekit.ParticipantInfo{}, "attributes"), protowire.VarintType), 0)
	one := field(num(&livekit.ListParticipantsResponse{}, "participants"), attrs)
	return []byte(strings.Repeat(string(one), n))
}

func unknownVarints(n int) []byte {
	one := protowire.AppendVarint(protowire.AppendTag(nil, 1000, protowire.VarintType), 1)
	return inLogRecord([]byte(strings.Repeat(string(one), n)))
}

func num(m proto.Message, name protoreflect.Name) protowire.Number {
	return m.ProtoReflect().Descriptor().Fields().ByName(name).Number()
}

// mergedBuckets repeats a singular Buckets field n times with one packed count
// each; the decoder merges them, regrowing one slice every time.
func mergedBuckets(n int) []byte {
	point := &metricsv1.ExponentialHistogramDataPoint{}
	buckets := &metricsv1.ExponentialHistogramDataPoint_Buckets{}
	one := field(num(point, "positive"), field(num(buckets, "bucket_counts"), protowire.AppendVarint(nil, 1)))
	dp := field(num(&metricsv1.ExponentialHistogram{}, "data_points"), []byte(strings.Repeat(string(one), n)))
	metric := field(num(&metricsv1.Metric{}, "exponential_histogram"), dp)
	return field(1, field(2, field(2, metric)))
}

// mergedResources repeats a singular Resource with one attribute n times.
func mergedResources(n int) []byte {
	attr := must(proto.Marshal(&resourcev1.Resource{Attributes: []*commonv1.KeyValue{{Key: "k", Value: str("v")}}}))
	one := field(num(&logsv1.ResourceLogs{}, "resource"), attr)
	return field(1, []byte(strings.Repeat(string(one), n)))
}

// alternatingOneof switches an AnyValue between two message members n times.
func alternatingOneof(n int) []byte {
	v := &commonv1.AnyValue{}
	kv := field(num(v, "kvlist_value"), must(proto.Marshal(&commonv1.KeyValueList{Values: []*commonv1.KeyValue{{Key: "k"}}})))
	arr := field(num(v, "array_value"), must(proto.Marshal(&commonv1.ArrayValue{Values: []*commonv1.AnyValue{{}}})))
	body := []byte(strings.Repeat(string(kv)+string(arr), n/2))
	return inLogRecord(field(num(&logsv1.LogRecord{}, "body"), body))
}

type calibrationCase struct {
	name   string
	data   []byte
	newMsg func() proto.Message
}

// scalarExtensions registers n int64 extensions of FieldOptions, numbered from
// scalarExtensionBase, and returns their numbers.
const scalarExtensionBase = 60000

var scalarExtensions = sync.OnceValue(func() []protowire.Number {
	const n = 64
	fd := &descriptorpb.FileDescriptorProto{
		Name:       proto.String("decodebudget_scalar_ext_test.proto"),
		Package:    proto.String("decodebudget.test"),
		Dependency: []string{"google/protobuf/descriptor.proto"},
	}
	for i := range n {
		fd.Extension = append(fd.Extension, &descriptorpb.FieldDescriptorProto{
			Name:     proto.String(fmt.Sprintf("scalar_%d", i)),
			Number:   proto.Int32(int32(scalarExtensionBase + i)),
			Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Type:     descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
			Extendee: proto.String(".google.protobuf.FieldOptions"),
		})
	}
	fd.Extension = append(fd.Extension, &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("repeated_scalar"),
		Number:   proto.Int32(scalarExtensionBase + n),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
		Extendee: proto.String(".google.protobuf.FieldOptions"),
	})
	file := must(protodesc.NewFile(fd, protoregistry.GlobalFiles))
	if err := protoregistry.GlobalTypes.RegisterExtension(dynamicpb.NewExtensionType(file.Extensions().Get(n))); err != nil {
		panic(err)
	}
	nums := make([]protowire.Number, n)
	for i := range n {
		xd := file.Extensions().Get(i)
		if err := protoregistry.GlobalTypes.RegisterExtension(dynamicpb.NewExtensionType(xd)); err != nil {
			panic(err)
		}
		nums[i] = xd.Number()
	}
	return nums
})

// scalarExtensionOptions is n occurrences of FieldOptions' int64 extensions,
// cycling through the first distinct of them.
func scalarExtensionOptions(n, distinct int) []byte {
	nums := scalarExtensions()
	var b []byte
	for i := range n {
		b = protowire.AppendVarint(protowire.AppendTag(b, nums[i%distinct], protowire.VarintType), 1<<40)
	}
	return b
}

// repeatedExtensionOptions is n elements of a repeated extension of
// FieldOptions, packed into runs of run elements, or unpacked when run is 0.
func repeatedExtensionOptions(number protowire.Number, n, run int, value uint64) []byte {
	var b []byte
	if run == 0 {
		for range n {
			b = protowire.AppendVarint(protowire.AppendTag(b, number, protowire.VarintType), value)
		}
		return b
	}
	for ; n > 0; n -= run {
		var values []byte
		for range min(n, run) {
			values = protowire.AppendVarint(values, value)
		}
		b = protowire.AppendBytes(protowire.AppendTag(b, number, protowire.BytesType), values)
	}
	return b
}

// repeatedField is n zero-valued occurrences of number with wire type typ.
func repeatedField(number protowire.Number, typ protowire.Type, n int) []byte {
	scalarExtensions()
	one := protowire.AppendTag(nil, number, typ)
	switch typ {
	case protowire.VarintType:
		one = protowire.AppendVarint(one, 0)
	case protowire.Fixed32Type:
		one = protowire.AppendFixed32(one, 0)
	case protowire.BytesType:
		one = protowire.AppendBytes(one, nil)
	}
	return []byte(strings.Repeat(string(one), n))
}

// unknownOptions is DescriptorProto{field: n×{options: {99999: 1}}}: the
// unknown field creates each FieldOptions' extension map.
func unknownOptions(n int) []byte {
	options := protowire.AppendVarint(protowire.AppendTag(nil, 99999, protowire.VarintType), 1)
	field := protowire.AppendBytes(protowire.AppendTag(nil, 8, protowire.BytesType), options)
	return []byte(strings.Repeat(string(protowire.AppendBytes(protowire.AppendTag(nil, 2, protowire.BytesType), field)), n))
}

func repeatedStrings(number protowire.Number, n int, value string) []byte {
	var b []byte
	for range n {
		b = protowire.AppendString(protowire.AppendTag(b, number, protowire.BytesType), value)
	}
	return b
}

func protoCalibrationCases() []calibrationCase {
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
	// DescriptorProto{field: n×{options: {[buf.validate.field]: {}}}}: every
	// FieldOptions gets an extension map.
	extensionOptions := func(n int) []byte {
		options := protowire.AppendBytes(protowire.AppendTag(nil, 1159, protowire.BytesType), nil)
		field := protowire.AppendBytes(protowire.AppendTag(nil, 8, protowire.BytesType), options)
		var b []byte
		for range n {
			b = protowire.AppendBytes(protowire.AppendTag(b, 2, protowire.BytesType), field)
		}
		return b
	}
	fieldOptions := func() proto.Message { return &descriptorpb.FieldOptions{} }
	fieldBehavior := annotations.E_FieldBehavior.TypeDescriptor().Number()
	participants := calibParticipants()
	return []calibrationCase{
		{"traces 512", must(proto.Marshal(calibSpans(512))), func() proto.Message { return &tracev1.TracesData{} }},
		{"logs 512", must(proto.Marshal(calibLogs(512))), func() proto.Message { return &logsv1.LogsData{} }},
		{"histograms 512", must(proto.Marshal(calibHistograms(512))), func() proto.Message { return &metricsv1.MetricsData{} }},
		{"empty spans 1M", emptySpans(1 << 20), func() proto.Message { return &tracev1.TracesData{} }},
		{"empty attributes 1M", emptyAttrs(1 << 20), func() proto.Message { return &logsv1.LogsData{} }},
		{"participants maps", must(proto.Marshal(participants)), func() proto.Message { return &livekit.ListParticipantsResponse{} }},
		{"one-entry maps 64k", must(proto.Marshal(calibOneEntryMaps(1 << 16))), func() proto.Message { return &livekit.RoomConfiguration{} }},
		{"map 128k", must(proto.Marshal(calibLargeMap(1 << 17))), func() proto.Message { return &livekit.ParticipantInfo{} }},
		{"valueless map entries 64k", valuelessMapEntries(1 << 16), func() proto.Message { return &livekit.DataTrackSubscriberHandles{} }},
		{"packed runs 3000", packedRuns(3000), func() proto.Message { return &metricsv1.MetricsData{} }},
		{"optional scalars 64k", optionalPoints(1 << 16), func() proto.Message { return &metricsv1.MetricsData{} }},
		{"wire type mismatch length", mismatchedLength(1 << 20), func() proto.Message { return &logsv1.LogsData{} }},
		{"wire type mismatch varints", mismatchedVarints(1 << 18), func() proto.Message { return &logsv1.LogsData{} }},
		{"wire type mismatch oneof", mismatchedOneof(1 << 18), func() proto.Message { return &logsv1.LogsData{} }},
		{"wire type mismatch oneof switching", mismatchedOneofSwitching(1600), func() proto.Message { return &livekit.DataPacket{} }},
		{"wire type mismatch maps", mismatchedMaps(1 << 16), func() proto.Message { return &livekit.ListParticipantsResponse{} }},
		{"wire type mismatch then map entry", mismatchedThenEntry(1 << 14), func() proto.Message { return &livekit.ListParticipantsResponse{} }},
		{"unknown varints", unknownVarints(1 << 18), func() proto.Message { return &logsv1.LogsData{} }},
		{"merged buckets 3000", mergedBuckets(3000), func() proto.Message { return &metricsv1.MetricsData{} }},
		{"merged resources 64k", mergedResources(1 << 16), func() proto.Message { return &logsv1.LogsData{} }},
		{"alternating oneof 64k", alternatingOneof(1 << 16), func() proto.Message { return &logsv1.LogsData{} }},
		{"extension options 16k", extensionOptions(1 << 14), func() proto.Message { return &descriptorpb.DescriptorProto{} }},
		{"dynamic scalar extension 16k", scalarExtensionOptions(1<<14, 1), fieldOptions},
		{"dynamic scalar extensions 64", scalarExtensionOptions(64, 64), fieldOptions},
		{"dynamic repeated extension 16k", repeatedExtensionOptions(scalarExtensionBase+64, 1<<14, 0, 1<<40), fieldOptions},
		{"repeated extension 16k", repeatedExtensionOptions(fieldBehavior, 1<<14, 0, 2), fieldOptions},
		{"packed extension 32k", repeatedExtensionOptions(fieldBehavior, 1<<15, 1<<14, 2), fieldOptions},
		{"extension as varint 64k", repeatedField(1159, protowire.VarintType, 1<<16), fieldOptions},
		{"repeated extension as fixed32 64k", repeatedField(fieldBehavior, protowire.Fixed32Type, 1<<16), fieldOptions},
		{"dynamic repeated extension as fixed32 64k", repeatedField(scalarExtensionBase+64, protowire.Fixed32Type, 1<<16), fieldOptions},
		{"empty packed extension runs 64k", repeatedField(scalarExtensionBase+64, protowire.BytesType, 1<<16), fieldOptions},
		{"unknown fields in extendable options 64k", unknownOptions(1 << 16), func() proto.Message { return &descriptorpb.DescriptorProto{} }},
		{"string extension 16k", repeatedStrings(annotations.E_DefaultHost.TypeDescriptor().Number(), 1<<14, strings.Repeat("h", 40)),
			func() proto.Message { return &descriptorpb.ServiceOptions{} }},
	}
}

func TestProtoEstimateCalibration(t *testing.T) {
	budget := Proto{MaxFields: 1 << 30, MaxDepth: 100}
	for _, tc := range protoCalibrationCases() {
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

func protoJSONCalibrationCases() []calibrationCase {
	emptySpans := `{"resourceSpans":[{"scopeSpans":[{"spans":[` + strings.TrimSuffix(strings.Repeat("{},", 1<<18), ",") + `]}]}]}`
	emptyValues := `{"resourceLogs":[{"scopeLogs":[{"logRecords":[{"body":{"arrayValue":{"values":[` + strings.TrimSuffix(strings.Repeat("{},", 1<<18), ",") + `]}}}]}]}]}`
	// The livekit package links buf.validate, whose extensions protojson
	// resolves from the global registry.
	fieldRules := func(rules, elem string) []byte {
		return fmt.Appendf(nil, `{"[buf.validate.field]":`+rules+`}`, strings.TrimSuffix(strings.Repeat(elem+",", 1<<15), ","))
	}
	fieldOptions := func() proto.Message { return &descriptorpb.FieldOptions{} }
	longNumber := "0." + strings.Repeat("0", 1<<18) + "1"
	chat := func(msg string) []byte { return []byte(`{"chatMessage":{"message":"` + msg + `"}}`) }
	escapedKeys := strings.TrimSuffix(strings.Repeat(`{"ident\u0069ty":"a","is\u0050ublisher":true,"disconnect\u0052eason":"CLIENT_INITIATED"},`, 8192), ",")
	newDataPacket := func() proto.Message { return &livekit.DataPacket{} }
	return []calibrationCase{
		{"traces 512", must(protojson.Marshal(calibSpans(512))), func() proto.Message { return &tracev1.TracesData{} }},
		{"logs 512", must(protojson.Marshal(calibLogs(512))), func() proto.Message { return &logsv1.LogsData{} }},
		{"histograms 512", must(protojson.Marshal(calibHistograms(512))), func() proto.Message { return &metricsv1.MetricsData{} }},
		{"empty spans", []byte(emptySpans), func() proto.Message { return &tracev1.TracesData{} }},
		{"empty values", []byte(emptyValues), func() proto.Message { return &logsv1.LogsData{} }},
		{"participants maps", must(protojson.Marshal(calibParticipants())), func() proto.Message { return &livekit.ListParticipantsResponse{} }},
		{"one-entry maps 16k", must(protojson.Marshal(calibOneEntryMaps(1 << 14))), func() proto.Message { return &livekit.RoomConfiguration{} }},
		{"map 64k", must(protojson.Marshal(calibLargeMap(1 << 16))), func() proto.Message { return &livekit.ParticipantInfo{} }},
		{"enums bytes int64 strings", must(protojson.Marshal(calibEnumBytesLogs(20000))), func() proto.Message { return &logsv1.LogsData{} }},
		{"proto field names", must(protojson.MarshalOptions{UseProtoNames: true}.Marshal(calibLogs(512))), func() proto.Message { return &logsv1.LogsData{} }},
		{"optional scalars", must(protojson.Marshal(calibOptionalPoints(16384))), func() proto.Message { return &metricsv1.MetricsData{} }},
		{"extension bytes", fieldRules(`{"bytes":{"in":[%s]}}`, `""`), fieldOptions},
		{"extension messages", fieldRules(`{"cel":[%s]}`, `{}`), fieldOptions},
		{"long number", []byte(`{"speaker":{"speakers":[{"level":` + longNumber + `}]}}`), newDataPacket},
		{"long number string", []byte(`{"speaker":{"speakers":[{"level":"` + longNumber + `"}]}}`), newDataPacket},
		{"escaped string", chat(strings.Repeat(strings.Repeat("n", 49)+`\n`, 4000)), newDataPacket},
		{"dense escapes", chat(strings.Repeat(`\"\\\u00e9\ud83d\ude00`, 20000)), newDataPacket},
		{"single late escape", chat(strings.Repeat("n", 200000) + `\n`), newDataPacket},
		{"replacement chars", chat(strings.Repeat(strings.Repeat("n", 47)+string(utf8.RuneError), 4000)), newDataPacket},
		{"escaped keys", []byte(`{"participants":[` + escapedKeys + `]}`), func() proto.Message { return &livekit.ListParticipantsResponse{} }},
	}
}

func TestProtoEstimateProtoJSONCalibration(t *testing.T) {
	budget := Proto{MaxFields: 1 << 30, MaxDepth: 100}
	for _, tc := range protoJSONCalibrationCases() {
		t.Run(tc.name, func(t *testing.T) {
			est, err := budget.EstimateProtoJSON(tc.data, tc.newMsg().ProtoReflect().Descriptor())
			require.NoError(t, err)
			got := measureJSONDecode(t, tc.data, tc.newMsg)
			t.Logf("json=%d measured=%d estimate=%d ratio=%.2f", len(tc.data), got, est, float64(est)/float64(got))
			require.GreaterOrEqual(t, uint64(est), got, "the estimate must not undercount")
			require.LessOrEqual(t, float64(est), 1.15*float64(got), "the estimate should stay close enough to admit real payloads")
		})
	}
}

func TestProtoEstimateProtoJSONWellKnownTypes(t *testing.T) {
	list := func(element string, n int) string { return strings.TrimSuffix(strings.Repeat(element+",", n), ",") }
	debug := func(elements string) []byte { return []byte(`{"payload":{"a":[` + elements + `]}}`) }
	anys := func(elements string) []byte { return []byte(`{"callContext":[` + elements + `]}`) }
	nested := `{"@type":"type.googleapis.com/google.protobuf.ListValue","value":[` + list(`0`, 1<<16) + `]}`
	for range 28 {
		nested = `{"@type":"type.googleapis.com/google.protobuf.Any","value":` + nested + `}`
	}
	newDebug := func() proto.Message { return &agent.DebugMessage{} }
	newCall := func() proto.Message { return &livekit.SIPCallInfo{} }
	cases := []calibrationCase{
		{"struct numbers", debug(list(`0`, 1<<16)), newDebug},
		{"struct nulls", debug(list(`null`, 1<<16)), newDebug},
		{"struct strings", debug(list(`"aaaaaaaaaaaaaaaaaaaa"`, 1<<16)), newDebug},
		{"struct objects", debug(list(`{"a":1}`, 1<<16)), newDebug},
		{"struct lists", debug(list(`[]`, 1<<16)), newDebug},
		{"root struct", []byte(`{"a":[` + list(`{}`, 1<<16) + `]}`), func() proto.Message { return &structpb.Struct{} }},
		{"any message", anys(`{"participants":[` + list(`{}`, 1<<16) + `],"@type":"type.googleapis.com/livekit.ListParticipantsResponse"}`), newCall},
		{"any decoy type", anys(`{"identity":"@type","@type":"type.googleapis.com/livekit.ParticipantInfo","tracks":[` + list(`{}`, 1<<16) + `]}`), newCall},
		{"any descriptor", anys(`{"@type":"type.googleapis.com/google.protobuf.FileDescriptorProto","messageType":[` + list(`{}`, 1<<16) + `]}`), newCall},
		{"any struct", anys(`{"@type":"type.googleapis.com/google.protobuf.Struct","value":{"a":[` + list(`0`, 1<<16) + `]}}`), newCall},
		{"any field mask", anys(`{"@type":"type.googleapis.com/google.protobuf.FieldMask","value":"` + list(`aAAAAAAAAA`, 1<<14) + `"}`), newCall},
		{"empty anys", anys(list(`{}`, 1<<16)), newCall},
		{"nested any", anys(nested), newCall},
	}
	budget := Proto{MaxFields: 1 << 30, MaxDepth: 100}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			est, err := budget.EstimateProtoJSON(tc.data, tc.newMsg().ProtoReflect().Descriptor())
			require.NoError(t, err)
			got := measureJSONDecode(t, tc.data, tc.newMsg)
			t.Logf("json=%d measured=%d estimate=%d ratio=%.2f", len(tc.data), got, est, float64(est)/float64(got))
			require.GreaterOrEqual(t, uint64(est), got, "the estimate must not undercount")
			// Any is charged for its worst-case marshalled size.
			require.LessOrEqual(t, float64(est), 1.5*float64(got), "the estimate should stay close enough to admit real payloads")
		})
	}
}

func measureUnmarshal(b Proto, data []byte, newMsg func() proto.Message) (uint64, error) {
	var best uint64
	var err error
	for i := 0; i < 3; i++ {
		var before, after runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&before)
		m := newMsg()
		err = b.Unmarshal(data, m)
		runtime.ReadMemStats(&after)
		runtime.KeepAlive(m)
		if d := after.TotalAlloc - before.TotalAlloc; best == 0 || d < best {
			best = d
		}
	}
	return best, err
}

func TestUnmarshalDynamicWithinBudget(t *testing.T) {
	data := packedSources(240_000)
	desc := (&livekit.ParticipantPermission{}).ProtoReflect().Descriptor()
	budget := Proto{MaxFields: 1 << 20, MaxDepth: 16, MaxDecodedBytes: 1 << 20}
	for _, tc := range []calibrationCase{
		{"generated", data, func() proto.Message { return &livekit.ParticipantPermission{} }},
		{"dynamicpb", data, func() proto.Message { return dynamicpb.NewMessage(desc) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := measureUnmarshal(budget, tc.data, tc.newMsg)
			t.Logf("allocated=%d limit=%d err=%v", got, budget.MaxDecodedBytes, err)
			if err == nil {
				require.LessOrEqual(t, got, uint64(budget.MaxDecodedBytes), "an admitted decode must stay within the limit")
			} else {
				require.ErrorIs(t, err, ErrMessageComplexity)
			}
		})
	}
}

func TestProtoEstimateDynamicCalibration(t *testing.T) {
	budget := Proto{MaxFields: 1 << 30, MaxDepth: 100, dynamic: true}
	large := protowire.AppendTag(nil, 6, protowire.BytesType) // HistogramDataPoint.bucket_counts
	large = protowire.AppendBytes(large, []byte(strings.Repeat(string(protowire.AppendFixed64(nil, 1<<40)), 1000)))
	cases := append(protoCalibrationCases(),
		calibrationCase{"packed enums 240k", packedSources(240_000), func() proto.Message { return &livekit.ParticipantPermission{} }},
		calibrationCase{"boxed fixed64s", large, func() proto.Message { return &metricsv1.HistogramDataPoint{} }},
	)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			desc := tc.newMsg().ProtoReflect().Descriptor()
			est, err := budget.Estimate(tc.data, desc)
			require.NoError(t, err)
			got := measureDecode(t, tc.data, func() proto.Message { return dynamicpb.NewMessage(desc) })
			t.Logf("wire=%d measured=%d estimate=%d ratio=%.2f", len(tc.data), got, est, float64(est)/float64(got))
			require.GreaterOrEqual(t, uint64(est), got, "the estimate must not undercount")
			// Boxing values under 256 is free, which the model does not track.
			require.LessOrEqual(t, float64(est), 1.25*float64(got), "the estimate should stay close enough to admit real payloads")
		})
	}
}

// A message built at runtime has no generated type, so Estimate prices the
// dynamicpb decode it can only take.
func TestProtoEstimateRuntimeMessage(t *testing.T) {
	file := must(protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    proto.String("decodebudget_runtime_test.proto"),
		Package: proto.String("decodebudget.runtime"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: proto.String("Strings"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:   proto.String("values"),
				Number: proto.Int32(1),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			}},
		}},
	}, protoregistry.GlobalFiles))
	desc := file.Messages().Get(0)
	data := repeatedStrings(1, 1<<16, "x")
	est, err := Proto{MaxFields: 1 << 30, MaxDepth: 100}.Estimate(data, desc)
	require.NoError(t, err)
	got := measureDecode(t, data, func() proto.Message { return dynamicpb.NewMessage(desc) })
	t.Logf("wire=%d measured=%d estimate=%d ratio=%.2f", len(data), got, est, float64(est)/float64(got))
	require.GreaterOrEqual(t, uint64(est), got, "the estimate must not undercount")
	require.LessOrEqual(t, float64(est), 1.25*float64(got), "the estimate should stay close enough to admit real payloads")
}

// A mismatched occurrence charges the map header but must not count as an entry.
func mismatchedThenEntry(n int) []byte {
	attrs := num(&livekit.ParticipantInfo{}, "attributes")
	entry := protowire.AppendString(protowire.AppendTag(nil, 1, protowire.BytesType), "k")
	entry = protowire.AppendString(protowire.AppendTag(entry, 2, protowire.BytesType), "v")
	info := protowire.AppendVarint(protowire.AppendTag(nil, attrs, protowire.VarintType), 0)
	info = protowire.AppendBytes(protowire.AppendTag(info, attrs, protowire.BytesType), entry)
	one := field(num(&livekit.ListParticipantsResponse{}, "participants"), info)
	return []byte(strings.Repeat(string(one), n))
}
