package decodebudget

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendAllocBytes(t *testing.T) {
	require.Equal(t, 0, appendAllocBytes(0, 8))
	// 8, 16, 32 and 64 byte arrays for caps 1, 2, 4 and 8.
	require.Equal(t, 8+16+32+64, appendAllocBytes(5, 8))
}
