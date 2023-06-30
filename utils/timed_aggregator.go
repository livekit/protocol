package utils

import (
	"errors"
	"sync"
	"time"
)

// ------------------------------------------------

var (
	errAnachronousSample = errors.New("anachronous sample")
)

// ------------------------------------------------

type number interface {
	int64 | float64
}

type TimedAggregatorParams struct {
	CapNegativeValues bool
}

type TimedAggregator[T number] struct {
	params TimedAggregatorParams

	lock              sync.RWMutex
	lastSample        T
	lastSampleAt      time.Time
	aggregate         T
	aggregateDuration time.Duration
}

func NewTimedAggregator[T number](params TimedAggregatorParams) *TimedAggregator[T] {
	return &TimedAggregator[T]{
		params: params,
	}
}

func (t *TimedAggregator[T]) AddSampleAt(val T, at time.Time) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if val < 0 && t.params.CapNegativeValues {
		val = 0
	}

	return t.addSampleAtLocked(val, at)
}

func (t *TimedAggregator[T]) AddSample(val T) error {
	return t.AddSampleAt(val, time.Now())
}

func (t *TimedAggregator[T]) addSampleAtLocked(val T, at time.Time) error {
	var sinceLast time.Duration
	if !t.lastSampleAt.IsZero() {
		if t.lastSampleAt.After(at) {
			return errAnachronousSample
		}

		sinceLast = at.Sub(t.lastSampleAt)
	}
	lastVal := t.lastSample

	t.lastSample = val
	t.lastSampleAt = at

	t.aggregate += T(sinceLast.Seconds() * float64(lastVal))
	t.aggregateDuration += sinceLast
	return nil
}

func (t *TimedAggregator[T]) GetAggregate() (T, time.Duration) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.aggregate, t.aggregateDuration
}

func (t *TimedAggregator[T]) GetAggregateAt(at time.Time) (T, time.Duration) {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.getAggregateAtLocked(at)
}

func (t *TimedAggregator[T]) GetAggregateAndRestartAt(at time.Time) (T, time.Duration) {
	t.lock.Lock()
	defer t.lock.Unlock()

	aggregate, aggregateDuration := t.getAggregateAtLocked(at)
	t.restartAtLocked(at)
	return aggregate, aggregateDuration
}

func (t *TimedAggregator[T]) getAggregateAtLocked(at time.Time) (T, time.Duration) {
	if !t.lastSampleAt.IsZero() {
		// re-add last sample at given time
		t.addSampleAtLocked(t.lastSample, at)
	}

	return t.aggregate, t.aggregateDuration
}

func (t *TimedAggregator[T]) GetAverage() float64 {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.getAverageLocked()
}

func (t *TimedAggregator[T]) GetAverageAt(at time.Time) float64 {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.getAverageAtLocked(at)
}

func (t *TimedAggregator[T]) GetAverageAndRestartAt(at time.Time) float64 {
	t.lock.Lock()
	defer t.lock.Unlock()

	average := t.getAverageAtLocked(at)
	t.restartAtLocked(at)
	return average
}

func (t *TimedAggregator[T]) getAverageLocked() float64 {
	seconds := t.aggregateDuration.Seconds()
	if seconds == 0.0 {
		return 0.0
	}

	return float64(t.aggregate) / seconds
}

func (t *TimedAggregator[T]) getAverageAtLocked(at time.Time) float64 {
	if !t.lastSampleAt.IsZero() {
		// re-add last sample at given time
		t.addSampleAtLocked(t.lastSample, at)
	}

	return t.getAverageLocked()
}

func (t *TimedAggregator[T]) Reset() {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.lastSample = 0
	t.lastSampleAt = time.Time{}
	t.aggregate = 0
	t.aggregateDuration = 0
}

func (t *TimedAggregator[T]) RestartAt(at time.Time) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.restartAtLocked(at)
}

func (t *TimedAggregator[T]) Restart() {
	t.RestartAt(time.Now())
}

func (t *TimedAggregator[T]) restartAtLocked(at time.Time) {
	t.lastSampleAt = at
	t.aggregate = 0
	t.aggregateDuration = 0
}
