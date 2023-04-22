package timeseries

import (
	"container/list"
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	minLinearFitSamples = 5
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

type TimeSeriesCompareOp int

const (
	TimeSeriesCompareOpEQ TimeSeriesCompareOp = iota
	TimeSeriesCompareOpNE
	TimeSeriesCompareOpGT
	TimeSeriesCompareOpGTE
	TimeSeriesCompareOpLT
	TimeSeriesCompareOpLTE
)

func (t TimeSeriesCompareOp) String() string {
	switch t {
	case TimeSeriesCompareOpEQ:
		return "EQ"
	case TimeSeriesCompareOpNE:
		return "NE"
	case TimeSeriesCompareOpGT:
		return "GT"
	case TimeSeriesCompareOpGTE:
		return "GTE"
	case TimeSeriesCompareOpLT:
		return "LT"
	case TimeSeriesCompareOpLTE:
		return "LTE"
	default:
		return fmt.Sprintf("%d", int(t))
	}
}

// ------------------------------------------------

type number interface {
	uint32 | uint64 | int | int32 | int64 | float32 | float64
}

type TimeSeriesSample[T number] struct {
	Value T
	At    time.Time
}

type TimeSeriesParams struct {
	UpdateOp         TimeSeriesUpdateOp
	Window           time.Duration
	CollapseDuration time.Duration
}

type TimeSeries[T number] struct {
	params TimeSeriesParams

	lock           sync.RWMutex
	samples        *list.List
	activeSample   T
	isActiveSample bool

	welfordCount int
	welfordM     float64
	welfordS     float64
	welfordStart time.Time
	welfordLast  time.Time
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

func (t *TimeSeries[T]) Sum() float64 {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()

	sum := float64(0.0)
	for e := t.samples.Front(); e != nil; e = e.Next() {
		s := e.Value.(TimeSeriesSample[T])
		sum += float64(s.Value)
	}

	return sum
}

func (t *TimeSeries[T]) Min() T {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()
	return t.minLocked()
}

func (t *TimeSeries[T]) minLocked() T {
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

	return t.maxLocked()
}

func (t *TimeSeries[T]) maxLocked() T {
	max := T(0)
	for e := t.samples.Front(); e != nil; e = e.Next() {
		s := e.Value.(TimeSeriesSample[T])
		if max < s.Value {
			max = s.Value
		}
	}

	return max
}

func (t *TimeSeries[T]) CurrentRun(threshold T, op TimeSeriesCompareOp) time.Duration {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()

	start := time.Time{}
	end := time.Time{}

	for e := t.samples.Back(); e != nil; e = e.Prev() {
		cond := false
		s := e.Value.(TimeSeriesSample[T])
		switch op {
		case TimeSeriesCompareOpEQ:
			cond = s.Value == threshold
		case TimeSeriesCompareOpNE:
			cond = s.Value != threshold
		case TimeSeriesCompareOpGT:
			cond = s.Value > threshold
		case TimeSeriesCompareOpGTE:
			cond = s.Value >= threshold
		case TimeSeriesCompareOpLT:
			cond = s.Value < threshold
		case TimeSeriesCompareOpLTE:
			cond = s.Value <= threshold
		}
		if !cond {
			break
		}
		if end.IsZero() {
			end = s.At
		}
		start = s.At
	}

	if end.IsZero() || start.IsZero() {
		return 0
	}

	return end.Sub(start)
}

func (t *TimeSeries[T]) OnlineAverage() float64 {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.welfordM
}

func (t *TimeSeries[T]) OnlineVariance() float64 {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.onlineVarianceLocked()
}

func (t *TimeSeries[T]) onlineVarianceLocked() float64 {
	if t.welfordCount > 1 {
		return t.welfordS / float64(t.welfordCount-1)
	}

	return 0.0
}

func (t *TimeSeries[T]) OnlineStdDev() float64 {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.onlineStdDevLocked()
}

func (t *TimeSeries[T]) onlineStdDevLocked() float64 {
	return math.Sqrt(t.onlineVarianceLocked())
}

func (t *TimeSeries[T]) ZScore(val T) float64 {
	t.lock.RLock()
	defer t.lock.RUnlock()

	onlineStdDev := t.onlineStdDevLocked()
	if onlineStdDev != 0.0 {
		return (float64(val) - t.welfordM) / t.onlineStdDevLocked()
	}

	return 0.0
}

func (t *TimeSeries[T]) LinearFit() (float64, float64) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.prune()

	if t.samples.Len() < minLinearFitSamples {
		return 0.0, 0.0
	}

	min := t.minLocked()
	max := t.maxLocked()
	diff := float64(max - min)

	start := time.Time{}

	sx := float64(0.0)
	sxsq := float64(0.0)
	sy := float64(0.0)
	sysq := float64(0.0)
	sxy := float64(0.0)

	for e := t.samples.Front(); e != nil; e = e.Next() {
		s := e.Value.(TimeSeriesSample[T])
		if start.IsZero() {
			start = s.At
		}
		since := s.At.Sub(start).Seconds()
		normy := 0.0
		if diff != 0.0 {
			normy = float64(s.Value-min) / diff
		}
		sx += since
		sxsq += since * since
		sy += normy
		sysq += normy * normy
		sxy += since * normy
	}

	N := float64(t.samples.Len())
	sxwsq := sx * sx
	denom := N*sxsq - sxwsq
	slope := float64(0.0)
	if denom != 0.0 {
		slope = (N*sxy - sx*sy) / denom
	}
	intercept := (sy - slope*sx) / N
	return slope, intercept
}

func (t *TimeSeries[T]) KendallsTau(numSamplesToUse int) float64 {
	t.lock.Lock()
	t.prune()

	if t.samples.Len() < numSamplesToUse {
		t.lock.Unlock()
		return 0.0
	}

	values := make([]T, numSamplesToUse)
	idx := numSamplesToUse - 1
	for e := t.samples.Back(); e != nil; e = e.Prev() {
		if idx < 0 {
			break
		}

		s := e.Value.(TimeSeriesSample[T])
		values[idx] = s.Value
		idx--
	}
	t.lock.Unlock()

	concordantPairs := 0
	discordantPairs := 0
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				concordantPairs++
			} else if values[i] > values[j] {
				discordantPairs++
			}
		}
	}

	if (concordantPairs + discordantPairs) == 0 {
		return 0.0
	}

	return (float64(concordantPairs) - float64(discordantPairs)) / (float64(concordantPairs) + float64(discordantPairs))
}

func (t *TimeSeries[T]) initSamples() {
	t.samples = t.samples.Init()
}

func (t *TimeSeries[T]) addSampleAt(val T, at time.Time) {
	// insert in time order
	e := t.samples.Back()
	if e != nil {
		lastSample := e.Value.(TimeSeriesSample[T])
		if val == lastSample.Value && at.Sub(lastSample.At) < t.params.CollapseDuration {
			// repeated value within collapse duration
			t.prune()
			return
		}
	}
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

	case t.samples.Front() != nil: // in the front
		t.samples.PushFront(sample)

	default: // at the end
		t.samples.PushBack(sample)
	}

	t.updateWelfordStats(val, at)

	t.prune()
}

func (t *TimeSeries[T]) updateWelfordStats(val T, at time.Time) {
	t.welfordCount++
	mLast := t.welfordM
	t.welfordM += (float64(val) - t.welfordM) / float64(t.welfordCount)
	t.welfordS += (float64(val) - mLast) * (float64(val) - t.welfordM)

	if t.welfordStart.IsZero() {
		t.welfordStart = at
	}
	t.welfordLast = at
}

func (t *TimeSeries[T]) prune() {
	thresh := t.welfordLast.Add(-t.params.Window)
	//thresh := time.Now().Add(-t.params.Window)

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
// - run
// - z-score
