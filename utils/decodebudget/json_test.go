package decodebudget

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {
	budget := JSON{MaxTokens: 20, MaxDepth: 3}
	for _, data := range []string{`null`, `true`, `false`, `1e100000`, `"value"`, `{}`, `[]`, `{"name":"camera","layers":[{}]}`} {
		t.Run(data, func(t *testing.T) { require.NoError(t, budget.Validate([]byte(data))) })
	}
	for _, data := range []string{"", " \n\t", `{} {}`, `true false`, `null 0`, `{"a":`, `[1,]`, `]`} {
		t.Run("invalid "+data, func(t *testing.T) { require.Error(t, budget.Validate([]byte(data))) })
	}
	require.NoError(t, (JSON{MaxTokens: 3, MaxDepth: 1}).Validate([]byte(`[0]`)))
	require.ErrorIs(t, (JSON{MaxTokens: 2, MaxDepth: 1}).Validate([]byte(`[0]`)), ErrMessageComplexity)
	require.NoError(t, (JSON{MaxTokens: 4, MaxDepth: 1}).Validate([]byte(`{"a":0}`)))
	require.ErrorIs(t, (JSON{MaxTokens: 3, MaxDepth: 1}).Validate([]byte(`{"a":0}`)), ErrMessageComplexity)
	require.NoError(t, budget.Validate([]byte(`[[[]]]`)))
	require.ErrorIs(t, budget.Validate([]byte(`[[[[]]]]`)), ErrMessageComplexity)
	require.ErrorIs(t, budget.Validate([]byte("["+strings.Repeat("{},", 20)+"{}]")), ErrMessageComplexity)
	for _, invalid := range []JSON{{}, {MaxTokens: 1}, {MaxDepth: 1}, {MaxTokens: -1, MaxDepth: 1}} {
		require.Error(t, invalid.Validate([]byte(`null`)))
	}
}
