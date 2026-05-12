package observability

import (
	"time"

	"github.com/livekit/protocol/utils/options"
)

type SessionTimer struct {
	lastMilli int64
	lastSec   int64
	lastMin   int64
	minSecs   int64
	minMins   int64
}

type SessionTimerOption func(*SessionTimer)

// WithMinSeconds ensures the first Advance that produces a non-zero secs
// return reports at least n seconds. Subsequent advances behave normally.
func WithMinSeconds(n int64) SessionTimerOption {
	return func(h *SessionTimer) { h.minSecs = n }
}

// WithMinMinutes ensures the first Advance that produces a non-zero mins
// return reports at least n minutes. Subsequent advances behave normally.
func WithMinMinutes(n int64) SessionTimerOption {
	return func(h *SessionTimer) { h.minMins = n }
}

func NewSessionTimer(startTime time.Time, opts ...SessionTimerOption) *SessionTimer {
	ts := startTime.UnixMilli()
	return options.Apply(&SessionTimer{lastMilli: ts, lastSec: ts, lastMin: ts}, opts)
}

func (h *SessionTimer) Advance(now time.Time) (millis, secs, mins int64) {
	ts := now.UnixMilli()
	if ts > h.lastMilli {
		millis = ts - h.lastMilli
		h.lastMilli = ts
	}
	if ts > h.lastSec {
		secs = max((ts-h.lastSec+999)/1000, h.minSecs)
		h.minSecs = 0
		h.lastSec += secs * 1000
	}
	if ts > h.lastMin {
		mins = max((ts-h.lastMin+59999)/60000, h.minMins)
		h.minMins = 0
		h.lastMin += mins * 60000
	}
	return
}
