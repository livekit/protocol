package utils

import (
	"time"

	"github.com/benbjohnson/clock"
)

type Clock interface {
	Now() time.Time
	Sleep(time.Duration)
}

type SystemClock struct{}

var _ Clock = &SystemClock{}

func (clock *SystemClock) Now() time.Time {
	return time.Now()
}

func (clock *SystemClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

type SimulatedClock struct{
  clock.Mock
}

var _ Clock = &SimulatedClock{}
