package decodebudget

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestProtoJSONStructure(t *testing.T) {
	desc := (&livekit.ParticipantInfo{}).ProtoReflect().Descriptor()
	budget := Proto{MaxFields: 4, MaxDepth: 2}
	require.NoError(t, budget.ValidateJSON([]byte(`{"identity":"a","tracks":[{"sid":"TR"}]}`), desc))
	require.ErrorIs(t, budget.ValidateJSON([]byte(`{"tracks":[{"sid":"a"},{"sid":"b"},{"sid":"c"}]}`), desc), ErrMessageComplexity)
	require.ErrorIs(t, budget.ValidateJSON([]byte(`{"tracks":[{"layers":[{}]}]}`), desc), ErrMessageComplexity)
	require.ErrorIs(t, (Proto{MaxFields: 4, MaxDepth: 2}).ValidateJSON([]byte(`{"unknown":[[[[1]]]]}`), desc), ErrMessageComplexity)
	require.Error(t, budget.ValidateJSON([]byte(`{} {}`), desc))
	require.Error(t, budget.ValidateJSON([]byte(`{"identity":`), desc))
}
