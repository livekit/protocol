package decodebudget

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"

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
