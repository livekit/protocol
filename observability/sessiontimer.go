package observability

import "time"

type SessionTimer struct {
	lastMilli int64
	lastMin   int64
}

func NewSessionTimer(startTime time.Time) *SessionTimer {
	ts := startTime.UnixMilli()
	return &SessionTimer{ts, ts}
}

func (h *SessionTimer) Advance(now time.Time) (millis, mins int64) {
	ts := now.UnixMilli()
	if ts > h.lastMilli {
		millis = ts - h.lastMilli
		h.lastMilli = ts
	}
	if ts > h.lastMin {
		mins = 1
		h.lastMin += 60000
	}
	return
}
