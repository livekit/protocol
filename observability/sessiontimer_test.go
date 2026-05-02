package observability

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSessionTimer(t *testing.T) {
	t.Run("advance 100ms", func(t *testing.T) {
		ts := time.Now()
		st := NewSessionTimer(ts)

		millis, secs, mins := st.Advance(ts.Add(100 * time.Millisecond))
		require.EqualValues(t, 100, millis)
		require.EqualValues(t, 1, secs)
		require.EqualValues(t, 1, mins)

		millis, secs, mins = st.Advance(ts.Add(200 * time.Millisecond))
		require.EqualValues(t, 100, millis)
		require.EqualValues(t, 0, secs)
		require.EqualValues(t, 0, mins)
	})

	t.Run("advance 2.5m", func(t *testing.T) {
		ts := time.Now()
		st := NewSessionTimer(ts)

		millis, secs, mins := st.Advance(ts.Add(150 * time.Second))
		require.EqualValues(t, 150000, millis)
		require.EqualValues(t, 150, secs)
		require.EqualValues(t, 3, mins)
	})

	t.Run("WithMinSeconds floors first advance and consumes watermark", func(t *testing.T) {
		ts := time.Now()
		st := NewSessionTimer(ts, WithMinSeconds(60))

		_, secs, _ := st.Advance(ts.Add(time.Second))
		require.EqualValues(t, 60, secs)

		_, secs, _ = st.Advance(ts.Add(30 * time.Second))
		require.EqualValues(t, 0, secs)

		_, secs, _ = st.Advance(ts.Add(62 * time.Second))
		require.EqualValues(t, 2, secs)
	})

	t.Run("WithMinSeconds does not lower larger natural value", func(t *testing.T) {
		ts := time.Now()
		st := NewSessionTimer(ts, WithMinSeconds(60))

		_, secs, _ := st.Advance(ts.Add(120 * time.Second))
		require.EqualValues(t, 120, secs)
	})

	t.Run("WithMinSeconds does not trigger on no-op advance", func(t *testing.T) {
		ts := time.Now()
		st := NewSessionTimer(ts, WithMinSeconds(60))

		_, secs, _ := st.Advance(ts)
		require.EqualValues(t, 0, secs)

		_, secs, _ = st.Advance(ts.Add(time.Second))
		require.EqualValues(t, 60, secs)
	})

	t.Run("WithMinMinutes floors first advance", func(t *testing.T) {
		ts := time.Now()
		st := NewSessionTimer(ts, WithMinMinutes(5))

		_, _, mins := st.Advance(ts.Add(time.Second))
		require.EqualValues(t, 5, mins)

		_, _, mins = st.Advance(ts.Add(2 * time.Minute))
		require.EqualValues(t, 0, mins)

		_, _, mins = st.Advance(ts.Add(7 * time.Minute))
		require.EqualValues(t, 2, mins)
	})

	t.Run("WithMinSeconds and WithMinMinutes apply independently", func(t *testing.T) {
		ts := time.Now()
		st := NewSessionTimer(ts, WithMinSeconds(45), WithMinMinutes(3))

		_, secs, mins := st.Advance(ts.Add(time.Second))
		require.EqualValues(t, 45, secs)
		require.EqualValues(t, 3, mins)
	})
}
