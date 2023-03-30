package utils

import (
	"fmt"
	"sync"
	"time"
)

// ------------------------------------------------

type TimeSeriesUpdateOp int

const (
	TimeSeriesUpdateOpAdd TimeSeriesUpdateOp = iota
	TimeSeriesUpdateOpMax
)

func (t TimeSeriesUpdateOp) String() string {
	switch t {
	case TimeSeriesUpdateOpAdd:
		return "ADD"
	case TimeSeriesUpdateOpMax:
		return "MAX"
	default:
		return fmt.Sprintf("%d", int(t))
	}
}

// ------------------------------------------------

type number interface {
	uint32 | float64
}

type TimeSeriesSample[T number] struct {
	Value T
	At    time.Time
}

type TimeSeriesParams struct {
	UpdateOp TimeSeriesUpdateOp
	Window   time.Duration
}

type TimeSeries[T number] struct {
	params TimeSeriesParams

	lock           sync.RWMutex
	samples        []TimeSeriesSample[T]
	activeSample   T
	isActiveSample bool
}

func NewTimeSeries[T number](params TimeSeriesParams) *TimeSeries[T] {
	t := &TimeSeries[T]{
		params: params,
	}

	t.initSamples()
	return t
}

func (t *TimeSeries[T]) UpdateSample(val T) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if !t.isActiveSample {
		t.isActiveSample = true
		t.activeSample = val
		return
	}

	switch t.params.UpdateOp {
	case TimeSeriesUpdateOpAdd:
		t.activeSample += val
	case TimeSeriesUpdateOpMax:
		if val > t.activeSample {
			t.activeSample = val
		}
	}
}

func (t *TimeSeries[T]) CommitActiveSample() {
	t.CommitActiveSampleAt(time.Now())
}

func (t *TimeSeries[T]) CommitActiveSampleAt(at time.Time) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()

	if !t.isActiveSample {
		return
	}

	t.addSampleAt(t.activeSample, at)
	t.isActiveSample = false
}

func (t *TimeSeries[T]) AddSample(val T) {
	t.AddSampleAt(val, time.Now())
}

func (t *TimeSeries[T]) AddSampleAt(val T, at time.Time) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.addSampleAt(val, at)
	t.prune()
}

func (t *TimeSeries[T]) GetSamples() []TimeSeriesSample[T] {
	t.lock.RLock()
	defer t.lock.RUnlock()

	samples := make([]TimeSeriesSample[T], len(t.samples))
	copy(samples, t.samples)
	return samples
}

func (t *TimeSeries[T]) ClearSamples() {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.initSamples()
}

func (t *TimeSeries[T]) Sum() T {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()

	sum := T(0)
	for _, s := range t.samples {
		sum += s.Value
	}

	return sum
}

func (t *TimeSeries[T]) Min() T {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()

	min := T(0)
	for _, s := range t.samples {
		if min == T(0) || min > s.Value {
			min = s.Value
		}
	}

	return min
}

func (t *TimeSeries[T]) Max() T {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()

	max := T(0)
	for _, s := range t.samples {
		if max < s.Value {
			max = s.Value
		}
	}

	return max
}

func (t *TimeSeries[T]) initSamples() {
	// set initial capacity assuming 1 second granularity
	numSamples := (t.params.Window + time.Second - 1) / time.Second
	t.samples = make([]TimeSeriesSample[T], 0, numSamples)
}

func (t *TimeSeries[T]) addSampleAt(val T, at time.Time) {
	// insert in time order
	switch {
	case len(t.samples) == 0 || at.After(t.samples[len(t.samples)-1].At): // empty or at end
		t.samples = append(t.samples, TimeSeriesSample[T]{
			Value: val,
			At:    at,
		})

	case at.Before(t.samples[0].At): // at start
		t.samples = append([]TimeSeriesSample[T]{
			{
				Value: val,
				At:    at,
			},
		}, t.samples...)

	default: // in between
		for idx := len(t.samples) - 1; idx >= 0; idx-- {
			if at.After(t.samples[idx].At) {
				t.samples = append(t.samples[:idx+1], t.samples[idx:]...)
				t.samples[idx+1] = TimeSeriesSample[T]{
					Value: val,
					At:    at,
				}
				break
			}
		}
	}

	t.prune()
}

func (t *TimeSeries[T]) prune() {
	thresh := time.Now().Add(-t.params.Window)
	for idx, s := range t.samples {
		if s.At.After(thresh) {
			t.samples = t.samples[idx:]
			break
		}
	}
}

// TODO - a bunch of stats
// - sum
// - moving average
// - EWMA
// - min
// - max
// - average
// - median
// - variance
// - stddev
// - trend
