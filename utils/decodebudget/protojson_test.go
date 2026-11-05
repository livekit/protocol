package decodebudget

import (
	"bytes"
	"encoding/json"
	"runtime"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

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

// quadraticJSONStringCost is jsonStringCost before rebuildScan, kept as the
// reference for its costs.
func quadraticJSONStringCost(raw []byte, decodedLen int) int {
	indexRebuild := func(in []byte) int {
		i := bytes.IndexByte(in, '\\')
		if i < 0 {
			i = len(in)
		}
		if j := bytes.Index(in[:i], []byte(replacementChar)); j >= 0 {
			return j
		}
		return i
	}
	in := raw[bytes.IndexByte(raw, '"')+1 : len(raw)-1]
	i := indexRebuild(in)
	if i == len(in) {
		return allocSize(decodedLen)
	}
	cost, l, c := 0, i, i
	grow := func(add int) {
		mem, newCap := appendGrow(l, c, add, 1)
		cost, l, c = cost+mem, l+add, newCap
	}
	for in = in[i:]; len(in) > 0; {
		if in[0] == '\\' {
			n, size := jsonEscapeLen(in)
			grow(n)
			in = in[size:]
			continue
		}
		_, n := utf8.DecodeRune(in)
		n += indexRebuild(in[n:])
		grow(n)
		in = in[n:]
	}
	return cost + allocSize(decodedLen)
}

func TestJSONStringCostMatchesReference(t *testing.T) {
	fffd := string(utf8.RuneError)
	for _, s := range []string{
		``, `plain`, `a\nb`, `é😀`, fffd, fffd + fffd + `x`, `x` + fffd + `\t` + fffd + `y`,
		strings.Repeat(`ab\"`+fffd, 100), strings.Repeat(fffd, 1000) + `\\`, "\xff\xfe" + fffd,
		strings.Repeat("x", 40) + `\n` + strings.Repeat(fffd+"yz", 50),
	} {
		raw := []byte(`, "` + s + `"`)
		var decoded string
		require.NoError(t, json.Unmarshal(raw[2:], &decoded), s)
		require.Equal(t, quadraticJSONStringCost(raw, len(decoded)), jsonStringCost(raw, len(decoded)), s)
	}
}

// A string of U+FFFD used to rescan its tail after each one.
func TestJSONStringCostLinear(t *testing.T) {
	elapsed := func(n int) time.Duration {
		raw := []byte(`"` + strings.Repeat(string(utf8.RuneError), n) + `"`)
		start := time.Now()
		jsonStringCost(raw, len(raw)-2)
		return time.Since(start)
	}
	elapsed(1 << 12)
	small, large := elapsed(1<<15), elapsed(1<<18)
	t.Logf("32k=%v 256k=%v", small, large)
	require.Less(t, large, 32*small+10*time.Millisecond, "8x the input must not take 64x as long")
}
