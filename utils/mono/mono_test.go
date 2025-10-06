package mono

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMonoZero(t *testing.T) {
	ts := time.Time{}
	ts2 := FromTime(ts)
	require.True(t, ts.IsZero())
	require.True(t, ts2.IsZero())
	require.True(t, ts.Equal(ts2))
	require.Equal(t, ts.String(), ts2.String())
}
