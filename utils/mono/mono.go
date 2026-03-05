// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// These APIs encode monotonic time into time.Time wall-clock fields. Returned
// values intentionally do not carry Go's internal monotonic payload and are
// meant to be compared with other mono timestamps.
//
// More details: https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md
package mono

import "time"

var (
	epoch     = time.Now()
	epochNano = epoch.UnixNano()
)

// resetClock resets the reference timestamp.
// Used in tests only.
func resetClock() {
	epoch = time.Now()
	epochNano = epoch.UnixNano()
}

// jumpClock adjusts reference timestamp by a given duration emulating a clock
// reset/jump. Used in tests only.
func jumpClock(dt time.Duration) {
	epoch = epoch.Add(-dt) // we pretend time.Now() jumps, not the reference
	epochNano = epoch.UnixNano()
}

// FromTime creates a Time from the monotonic part of t. Note that the monotonic
// part of t could have been erased when using functions like Truncate, Round,
// In, UTC, etc... Be careful when using this
func FromTime(t time.Time) time.Time {
	if t.IsZero() {
		return time.Time{}
	}
	return time.Unix(0, epochNano+int64(t.Sub(epoch)))
}

// Now creates a monotonic time without reading the system wall clock
func Now() time.Time {
	return time.Unix(0, epochNano+int64(time.Since(epoch)))
}

// Unix is an analog of time.Unix that produces monotonic time.
func Unix(sec, nsec int64) time.Time {
	return FromTime(time.Unix(sec, nsec))
}

// Parse is an analog of time.Parse that produces monotonic time.
func Parse(layout, value string) (time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, err
	}
	return FromTime(t), nil
}

// UnixNano returns the number of nanoseconds elapsed, based on the application
// start time.
func UnixNano() int64 {
	return epochNano + int64(time.Since(epoch))
}

// UnixMicro returns the number of microseconds elapsed, based on the
// application start time.
func UnixMicro() int64 {
	return UnixNano() / 1000
}
