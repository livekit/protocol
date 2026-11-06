package decodebudget

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils/protojson"
)

func TestProtoDecodedBytes(t *testing.T) {
	input := &livekit.ParticipantInfo{Identity: strings.Repeat("x", 1000), Attributes: map[string]string{"a": "one"}}
	data, err := proto.Marshal(input)
	require.NoError(t, err)
	desc := input.ProtoReflect().Descriptor()

	unlimited := Proto{MaxFields: 100, MaxDepth: 4}
	estimate, err := unlimited.Estimate(data, desc)
	require.NoError(t, err)
	require.Greater(t, estimate, 1000)

	require.NoError(t, (Proto{MaxFields: 100, MaxDepth: 4, MaxDecodedBytes: estimate}).Validate(data, desc))
	require.ErrorIs(t, (Proto{MaxFields: 100, MaxDepth: 4, MaxDecodedBytes: estimate - 1}).Validate(data, desc), ErrMessageComplexity)
	require.Error(t, (Proto{MaxFields: 100, MaxDepth: 4, MaxDecodedBytes: -1}).Validate(data, desc))

	jsonData, err := protojson.Marshal(input)
	require.NoError(t, err)
	jsonEstimate, err := unlimited.EstimateProtoJSON(jsonData, desc)
	require.NoError(t, err)
	require.ErrorIs(t, (Proto{MaxFields: 100, MaxDepth: 4, MaxDecodedBytes: jsonEstimate - 1}).ValidateProtoJSON(jsonData, desc), ErrMessageComplexity)
}
