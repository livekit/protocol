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

func TestMono(t *testing.T) {
	t.Cleanup(resetClock)

	ts1 := time.Now()
	ts2 := ts1.Add(time.Second)

	ts1m := FromTime(ts1)
	// emulate a clock reset, +1h jump
	// TODO: use synctest when we switch to Go 1.25
	jumpClock(time.Hour)
	ts2m := FromTime(ts2)

	require.Equal(t, ts2.Sub(ts1), ts2m.Sub(ts1m))
}

func TestNoGoMonotonicPayload(t *testing.T) {
	t.Cleanup(resetClock)

	now := Now()
	fromTime := FromTime(time.Now())
	fromUnix := Unix(123, 456)
	fromParse, err := Parse(time.RFC3339Nano, "2026-03-05T12:34:56.789123456Z")
	require.NoError(t, err)

	require.Equal(t, now, now.Round(0))
	require.Equal(t, fromTime, fromTime.Round(0))
	require.Equal(t, fromUnix, fromUnix.Round(0))
	require.Equal(t, fromParse, fromParse.Round(0))
}

func TestSerializationRoundTripComparableAcrossClockJump(t *testing.T) {
	t.Cleanup(resetClock)

	t1 := Now()
	t2 := t1.Add(10 * time.Second)

	t1Encoded := t1.UnixNano()
	t2Encoded := t2.UnixNano()

	// emulate a clock reset, +1h jump
	jumpClock(time.Hour)

	t1Decoded := Unix(0, t1Encoded)
	t2Decoded := Unix(0, t2Encoded)

	require.Equal(t, t2.Sub(t1), t2Decoded.Sub(t1Decoded))
}

func TestNowProgressesAcrossClockJump(t *testing.T) {
	t.Cleanup(resetClock)

	t1 := Now()
	time.Sleep(time.Millisecond)

	// emulate a clock reset, +1h jump
	jumpClock(time.Hour)
	t2 := Now()

	require.Greater(t, t2.Sub(t1), time.Duration(0))
}

func BenchmarkTime(b *testing.B) {
	b.Run("Now()", func(b *testing.B) {
		for b.Loop() {
			_ = Now()
		}
	})
	b.Run("time.Now()", func(b *testing.B) {
		for b.Loop() {
			_ = time.Now()
		}
	})
}
