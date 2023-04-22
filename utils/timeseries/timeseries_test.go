package timeseries

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
		require.Equal(t, float64(20.0), ts.Sum())

		ts.AddSample(30)
		require.Equal(t, float64(50.0), ts.Sum())
	})

	t.Run("min", func(t *testing.T) {
		ts := NewTimeSeries[uint32](TimeSeriesParams{
			UpdateOp: TimeSeriesUpdateOpLatest,
			Window:   2 * time.Minute,
		})

		ts.UpdateSample(10)
		ts.UpdateSample(20)
		ts.UpdateSample(15)
		ts.CommitActiveSampleAt(time.Now())
		require.Equal(t, uint32(15), ts.Min())

		ts.AddSample(30)
		require.Equal(t, uint32(15), ts.Min())
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

	t.Run("current_run", func(t *testing.T) {
		ts := NewTimeSeries[uint32](TimeSeriesParams{
			UpdateOp: TimeSeriesUpdateOpMax,
			Window:   time.Minute,
		})

		testCases := []struct {
			name           string
			values         []uint32
			timeStep       time.Duration
			compareOp      TimeSeriesCompareOp
			threshold      uint32
			expectedResult time.Duration
		}{
			{
				name: "eq_run",
				values: []uint32{
					10,
					20,
					30,
					40,
					40,
					40,
				},
				timeStep:       time.Second,
				compareOp:      TimeSeriesCompareOpEQ,
				threshold:      40,
				expectedResult: 2 * time.Second,
			},
			{
				name: "eq_no_run",
				values: []uint32{
					10,
					20,
					30,
					40,
					40,
					40,
				},
				timeStep:       time.Second,
				compareOp:      TimeSeriesCompareOpEQ,
				threshold:      50,
				expectedResult: 0,
			},
			{
				name: "ne",
				values: []uint32{
					10,
					20,
					30,
					40,
					40,
					40,
				},
				timeStep:       time.Second,
				compareOp:      TimeSeriesCompareOpNE,
				threshold:      50,
				expectedResult: 5 * time.Second,
			},
			{
				name: "gt",
				values: []uint32{
					10,
					20,
					30,
					40,
					40,
					40,
				},
				timeStep:       time.Second,
				compareOp:      TimeSeriesCompareOpGT,
				threshold:      20,
				expectedResult: 3 * time.Second,
			},
			{
				name: "gte",
				values: []uint32{
					10,
					20,
					30,
					40,
					40,
					40,
				},
				timeStep:       time.Second,
				compareOp:      TimeSeriesCompareOpGTE,
				threshold:      20,
				expectedResult: 4 * time.Second,
			},
			{
				name: "lt",
				values: []uint32{
					50,
					20,
					30,
					40,
					40,
					40,
				},
				timeStep:       time.Second,
				compareOp:      TimeSeriesCompareOpLT,
				threshold:      50,
				expectedResult: 4 * time.Second,
			},
			{
				name: "lte",
				values: []uint32{
					10,
					20,
					30,
					40,
					40,
					40,
				},
				timeStep:       time.Second,
				compareOp:      TimeSeriesCompareOpLTE,
				threshold:      40,
				expectedResult: 5 * time.Second,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				ts.ClearSamples()

				now := time.Now()
				for idx, value := range tc.values {
					ts.AddSampleAt(value, now.Add(time.Duration(idx)*tc.timeStep))
				}

				require.Equal(t, tc.expectedResult, ts.CurrentRun(tc.threshold, tc.compareOp))
			})
		}
	})

	t.Run("online", func(t *testing.T) {
		ts := NewTimeSeries[uint32](TimeSeriesParams{
			UpdateOp: TimeSeriesUpdateOpMax,
			Window:   time.Minute,
		})

		now := time.Now()
		for val := uint32(1); val <= 10; val++ {
			ts.AddSampleAt(val, now.Add(time.Duration(val)*time.Second))
		}

		require.Equal(t, float64(5.5), ts.OnlineAverage())
		onlineVariance := ts.OnlineVariance()
		require.Condition(t, func() bool { return onlineVariance > 9.16 && onlineVariance < 9.17 }, "online variance out of range")
		onlineStdDev := ts.OnlineStdDev()
		require.Condition(t, func() bool { return onlineStdDev > 3.02 && onlineStdDev < 3.03 }, "online std dev out of range")
	})

	t.Run("collapse", func(t *testing.T) {
		ts := NewTimeSeries[uint32](TimeSeriesParams{
			UpdateOp:         TimeSeriesUpdateOpMax,
			Window:           time.Minute,
			CollapseDuration: 2 * time.Second,
		})

		// add same value spaced apart by half the collapse duration, should add only five to the list
		now := time.Now()
		for i := 0; i < 10; i++ {
			ts.AddSampleAt(42, now.Add(time.Duration(i)*time.Second))
		}
		samples := ts.GetSamples()
		require.Equal(t, 5, len(samples))
		require.Equal(t, uint32(42), samples[0].Value) // spot check
		require.Equal(t, uint32(42), samples[3].Value) // spot check

		// add a sample of different value within the collapse window, it should get added
		ts.AddSampleAt(43, now.Add(time.Duration(9)*time.Second)) // same time offset as last sample to keep within collapse window
		samples = ts.GetSamples()
		require.Equal(t, 6, len(samples))
		require.Equal(t, uint32(42), samples[0].Value) // spot check
		require.Equal(t, uint32(42), samples[3].Value) // spot check
		require.Equal(t, uint32(43), samples[5].Value)

		// add a sample with same value as initial burst within the collapse window, it should get added
		ts.AddSampleAt(42, now.Add(time.Duration(10)*time.Second))
		samples = ts.GetSamples()
		require.Equal(t, 7, len(samples))
		require.Equal(t, uint32(42), samples[0].Value) // spot check
		require.Equal(t, uint32(42), samples[3].Value) // spot check
		require.Equal(t, uint32(43), samples[5].Value)
		require.Equal(t, uint32(42), samples[6].Value)
	})

	t.Run("kendall's tau", func(t *testing.T) {
		ts := NewTimeSeries[int64](TimeSeriesParams{
			UpdateOp: TimeSeriesUpdateOpMax,
			Window:   time.Minute,
		})

		// increasing values
		now := time.Now()
		for val := int64(1); val <= 10; val++ {
			ts.AddSampleAt(val, now.Add(time.Duration(val)*time.Second))
		}

		// asking to use more samples than available should return 0.0
		require.Equal(t, float64(0.0), ts.KendallsTau(11))

		// ever increasing should return 1.0
		require.Equal(t, float64(1.0), ts.KendallsTau(8))

		ts.ClearSamples()

		// decreasing values
		now = time.Now()
		for val := int64(1); val <= 10; val++ {
			ts.AddSampleAt(11-val, now.Add(time.Duration(val)*time.Second))
		}

		// ever decreasing should return -1.0
		require.Equal(t, float64(-1.0), ts.KendallsTau(8))

		ts.ClearSamples()

		// overall increasing
		now = time.Now()
		for val := int64(1); val <= 10; val++ {
			if val&0x1 == 0 {
				ts.AddSampleAt(2*val, now.Add(time.Duration(val)*time.Second))
			} else {
				ts.AddSampleAt(val, now.Add(time.Duration(val)*time.Second))
			}
		}

		// increasing envelope should trend positive
		require.Less(t, float64(0.0), ts.KendallsTau(8))

		// overall decreasing
		now = time.Now()
		for val := int64(1); val <= 10; val++ {
			if val&0x1 == 0 {
				ts.AddSampleAt(2*(11-val), now.Add(time.Duration(val)*time.Second))
			} else {
				ts.AddSampleAt(11-val, now.Add(time.Duration(val)*time.Second))
			}
		}

		// decreasing envelope should trend negative
		require.Greater(t, float64(0.0), ts.KendallsTau(8))
	})

}
