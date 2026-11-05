package decodebudget

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/dynamicpb"

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

// packedSources is a ParticipantPermission whose can_publish_sources holds n
// CAMERA values in one packed run.
func packedSources(n int) []byte {
	return protowire.AppendBytes(protowire.AppendTag(nil, 9, protowire.BytesType), []byte(strings.Repeat("\x01", n)))
}

func TestProtoUnmarshalDynamicDestination(t *testing.T) {
	data := packedSources(240_000)
	desc := (&livekit.ParticipantPermission{}).ProtoReflect().Descriptor()
	budget := Proto{MaxFields: 1 << 20, MaxDepth: 16, MaxDecodedBytes: 1 << 20}

	require.NoError(t, budget.Unmarshal(data, &livekit.ParticipantPermission{}))
	dynamic := dynamicpb.NewMessage(desc)
	require.ErrorIs(t, budget.Unmarshal(data, dynamic), ErrMessageComplexity)
	require.True(t, proto.Equal(dynamicpb.NewMessage(desc), dynamic))
	wrapped := struct{ *dynamicpb.Message }{dynamicpb.NewMessage(desc)}
	require.ErrorIs(t, budget.Unmarshal(data, wrapped), ErrMessageComplexity)

	p := DefaultProfile()
	p.DynamicMessages = true
	estimate, err := (Proto{MaxFields: 1 << 20, MaxDepth: 16, Profile: &p}).Estimate(data, desc)
	require.NoError(t, err)
	budget.MaxDecodedBytes = estimate
	require.NoError(t, budget.Unmarshal(data, dynamic))
	require.Equal(t, 240_000, dynamic.Get(desc.Fields().ByNumber(9)).List().Len())
}
