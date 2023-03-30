package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeSeries(t *testing.T) {
	t.Run("ordering", func(t *testing.T) {
		ts := NewTimeSeries[uint32](TimeSeriesParams{
			UpdateOp: TimeSeriesUpdateOpMax,
			Window:   time.Minute,
		})

		now := time.Now()
		expectedSamples := []TimeSeriesSample[uint32]{
			{
				Value: 10,
				At:    now,
			},
			{
				Value: 20,
				At:    now.Add(time.Second),
			},
			{
				Value: 30,
				At:    now.Add(2 * time.Second),
			},
			{
				Value: 40,
				At:    now.Add(3 * time.Second),
			},
		}

		testCases := []struct {
			name    string
			samples []TimeSeriesSample[uint32]
		}{
			{
				name: "regular",
				samples: []TimeSeriesSample[uint32]{
					{10, now},
					{20, now.Add(time.Second)},
					{30, now.Add(2 * time.Second)},
					{40, now.Add(3 * time.Second)},
				},
			},
			{
				name: "reverse",
				samples: []TimeSeriesSample[uint32]{
					{40, now.Add(3 * time.Second)},
					{30, now.Add(2 * time.Second)},
					{20, now.Add(time.Second)},
					{10, now},
				},
			},
			{
				name: "jumbled 1",
				samples: []TimeSeriesSample[uint32]{
					{20, now.Add(time.Second)},
					{40, now.Add(3 * time.Second)},
					{30, now.Add(2 * time.Second)},
					{10, now},
				},
			},
			{
				name: "jumbled 2",
				samples: []TimeSeriesSample[uint32]{
					{10, now},
					{40, now.Add(3 * time.Second)},
					{30, now.Add(2 * time.Second)},
					{20, now.Add(time.Second)},
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				ts.ClearSamples()

				for _, tss := range tc.samples {
					ts.AddSampleAt(tss.Value, tss.At)
				}

				require.Equal(t, expectedSamples, ts.GetSamples())
			})
		}
	})

	t.Run("sum", func(t *testing.T) {
		ts := NewTimeSeries[uint32](TimeSeriesParams{
			UpdateOp: TimeSeriesUpdateOpMax,
			Window:   2 * time.Minute,
		})

		ts.UpdateSample(10)
		ts.UpdateSample(20)
		ts.CommitActiveSampleAt(time.Now())
		require.Equal(t, uint32(20), ts.Sum())

		ts.AddSample(30)
		require.Equal(t, uint32(50), ts.Sum())
	})

	t.Run("min", func(t *testing.T) {
		ts := NewTimeSeries[uint32](TimeSeriesParams{
			UpdateOp: TimeSeriesUpdateOpMax,
			Window:   2 * time.Minute,
		})

		ts.UpdateSample(10)
		ts.UpdateSample(20)
		ts.CommitActiveSampleAt(time.Now())
		require.Equal(t, uint32(20), ts.Min())

		ts.AddSample(30)
		require.Equal(t, uint32(20), ts.Min())
	})

	t.Run("max", func(t *testing.T) {
		ts := NewTimeSeries[uint32](TimeSeriesParams{
			UpdateOp: TimeSeriesUpdateOpAdd,
			Window:   2 * time.Minute,
		})

		ts.UpdateSample(10)
		ts.UpdateSample(20)
		ts.CommitActiveSampleAt(time.Now())
		require.Equal(t, uint32(30), ts.Max())

		ts.AddSample(20)
		require.Equal(t, uint32(30), ts.Max())
	})
}
