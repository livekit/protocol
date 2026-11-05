package decodebudget

import (
	"bytes"
	"encoding/json"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/livekit/protocol/livekit"
)

func TestProtoJSONStructure(t *testing.T) {
	desc := (&livekit.ParticipantInfo{}).ProtoReflect().Descriptor()
	budget := Proto{MaxFields: 4, MaxDepth: 2}
	require.NoError(t, budget.ValidateProtoJSON([]byte(`{"identity":"a","tracks":[{"sid":"TR"}]}`), desc))
	require.ErrorIs(t, budget.ValidateProtoJSON([]byte(`{"tracks":[{"sid":"a"},{"sid":"b"},{"sid":"c"}]}`), desc), ErrMessageComplexity)
	require.ErrorIs(t, budget.ValidateProtoJSON([]byte(`{"tracks":[{"layers":[{}]}]}`), desc), ErrMessageComplexity)
	require.ErrorIs(t, (Proto{MaxFields: 4, MaxDepth: 2}).ValidateProtoJSON([]byte(`{"unknown":[[[[1]]]]}`), desc), ErrMessageComplexity)
	require.Error(t, budget.ValidateProtoJSON([]byte(`{} {}`), desc))
	require.Error(t, budget.ValidateProtoJSON([]byte(`{"identity":`), desc))
}

func TestProtoJSONExtensions(t *testing.T) {
	file := protoTestFile(t)
	desc := file.Messages().ByName("Container")
	budget := Proto{MaxFields: 6, MaxDepth: 4}
	require.NoError(t, budget.ValidateProtoJSON([]byte(`{"[decode_budget_test.items]":[{"children":[{}]}]}`), desc))
	require.ErrorIs(t, budget.ValidateProtoJSON([]byte(`{"[decode_budget_test.items]":[{"children":[{"children":[{"children":[{}]}]}]}]}`), desc), ErrMessageComplexity)
	require.NoError(t, budget.ValidateProtoJSON([]byte(`{"[decode_budget_test.missing]":[{}]}`), desc))
	require.NoError(t, budget.ValidateProtoJSON([]byte(`{"[":1,"]":1,"[]":1}`), desc))
	require.NoError(t, budget.ValidateProtoJSON([]byte(`{"[decode_budget_test.items]":[{}]}`), file.Messages().ByName("Node")))

	// The livekit package links buf.validate, so protojson decodes its
	// extensions in full; they must cost what the same content does as a root.
	rules, err := protoregistry.GlobalTypes.FindMessageByName("buf.validate.FieldRules")
	require.NoError(t, err)
	content := `{"bytes":{"in":[` + strings.TrimSuffix(strings.Repeat(`"",`, 1<<12), ",") + `]}}`
	unlimited := Proto{MaxFields: 1 << 30, MaxDepth: 100}
	asRoot, err := unlimited.EstimateProtoJSON([]byte(content), rules.Descriptor())
	require.NoError(t, err)
	options := (&descriptorpb.FieldOptions{}).ProtoReflect().Descriptor()
	asExtension, err := unlimited.EstimateProtoJSON([]byte(`{"[buf.validate.field]":`+content+`}`), options)
	require.NoError(t, err)
	require.Greater(t, asExtension, asRoot)
	unregistered, err := unlimited.EstimateProtoJSON([]byte(`{"[buf.validate.nosuch]":`+content+`}`), options)
	require.NoError(t, err)
	require.Less(t, unregistered, asRoot)
}

// Each Any looks ahead for "@type" before anything is charged, so with
// "@type" last every level scans its whole subtree; the scan must not copy it.
func TestProtoJSONNestedAnyLookahead(t *testing.T) {
	const anyURL = "type.googleapis.com/google.protobuf.Any"
	doc := `{"value":"` + strings.Repeat("a", 1<<20) + `","@type":"type.googleapis.com/google.protobuf.StringValue"}`
	for range 90 {
		doc = `{"value":` + doc + `,"@type":"` + anyURL + `"}`
	}
	desc := (&anypb.Any{}).ProtoReflect().Descriptor()
	budget := Proto{MaxFields: 1 << 20, MaxDepth: 100, MaxDecodedBytes: 1 << 20}

	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	err := budget.ValidateProtoJSON([]byte(doc), desc)
	runtime.ReadMemStats(&after)
	require.ErrorIs(t, err, ErrMessageComplexity)
	require.Less(t, after.TotalAlloc-before.TotalAlloc, uint64(16<<20))
}

func TestProtoJSONAnyTypeURL(t *testing.T) {
	const url = "type.googleapis.com/google.protobuf.StringValue"
	for _, tc := range []struct{ doc, want string }{
		{`{"@type":"` + url + `","value":"x"}`, url},
		{`{"value":"x","@type":"` + url + `"}`, url},
		{` { "skip" : [ {"@type":"no"}, "]}\"", -1.5e5, true, null, {} ] , "o":{"a":[1]}, "\u0040type" : "type.googleapis.com/google.protobuf.String\u0056alue" } `, url},
		{`{}`, ""},
		{`{"value":"x"}`, ""},
		{`{"@type":5}`, ""},
		{`{"value":[1,"@type"],"other":{"@type":"` + url + `"}}`, ""},
		{`{"value":`, ""},
		{`{"value":"x" "@type":"` + url + `"}`, ""},
		{`{"value":}`, ""},
		{`{"value":"unterminated`, ""},
		{`{"@type":"bad\escape"}`, ""},
	} {
		data := []byte(tc.doc)
		dec := json.NewDecoder(bytes.NewReader(data))
		_, err := dec.Token()
		require.NoError(t, err)
		w := &jsonWalk{dec: dec, data: data}
		require.Equal(t, tc.want, w.anyTypeURL(), tc.doc)
	}
}
