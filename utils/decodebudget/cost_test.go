package decodebudget

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendAllocBytes(t *testing.T) {
	require.Equal(t, 0, appendAllocBytes(0, 8))
	// 8, 16, 32 and 64 byte arrays for caps 1, 2, 4 and 8.
	require.Equal(t, 8+16+32+64, appendAllocBytes(5, 8))
}

func TestAllocSizeMatchesSizeClasses(t *testing.T) {
	for n := 1; n <= 40000; n++ {
		want := (n + 8191) / 8192 * 8192
		if i, _ := slices.BinarySearch(goSizeClasses, n); i < len(goSizeClasses) {
			want = goSizeClasses[i]
		}
		if got := allocSize(n); got != want {
			t.Fatalf("allocSize(%d) = %d, want %d", n, got, want)
		}
	}
}
