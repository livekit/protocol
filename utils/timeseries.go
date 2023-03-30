package utils

import (
	"container/list"
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
	samples        *list.List
	activeSample   T
	isActiveSample bool
}

func NewTimeSeries[T number](params TimeSeriesParams) *TimeSeries[T] {
	t := &TimeSeries[T]{
		params:  params,
		samples: list.New(),
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
}

func (t *TimeSeries[T]) GetSamples() []TimeSeriesSample[T] {
	t.lock.RLock()
	defer t.lock.RUnlock()

	samples := make([]TimeSeriesSample[T], 0, t.samples.Len())
	for e := t.samples.Front(); e != nil; e = e.Next() {
		samples = append(samples, e.Value.(TimeSeriesSample[T]))
	}
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
	for e := t.samples.Front(); e != nil; e = e.Next() {
		s := e.Value.(TimeSeriesSample[T])
		sum += s.Value
	}

	return sum
}

func (t *TimeSeries[T]) Min() T {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()

	min := T(0)
	for e := t.samples.Front(); e != nil; e = e.Next() {
		s := e.Value.(TimeSeriesSample[T])
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
	for e := t.samples.Front(); e != nil; e = e.Next() {
		s := e.Value.(TimeSeriesSample[T])
		if max < s.Value {
			max = s.Value
		}
	}

	return max
}

func (t *TimeSeries[T]) initSamples() {
	t.samples = t.samples.Init()
}

func (t *TimeSeries[T]) addSampleAt(val T, at time.Time) {
	// insert in time order
	var e *list.Element
	for e = t.samples.Back(); e != nil; e = e.Prev() {
		s := e.Value.(TimeSeriesSample[T])
		if at.After(s.At) {
			break
		}
	}

	sample := TimeSeriesSample[T]{
		Value: val,
		At:    at,
	}
	switch {
	case e != nil: // in the middle
		t.samples.InsertAfter(sample, e)

	case t.samples.Front() != nil: // at the end
		t.samples.PushFront(sample)

	default: // in the front
		t.samples.PushBack(sample)
	}

	t.prune()
}

func (t *TimeSeries[T]) prune() {
	thresh := time.Now().Add(-t.params.Window)

	toRemove := make([]*list.Element, 0, t.samples.Len())
	for e := t.samples.Front(); e != nil; e = e.Next() {
		s := e.Value.(TimeSeriesSample[T])
		if s.At.After(thresh) {
			break
		}

		toRemove = append(toRemove, e)
	}

	for _, e := range toRemove {
		t.samples.Remove(e)
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
