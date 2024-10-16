package mono

import "time"

var (
	epoch     = time.Now()
	epochNano = epoch.UnixNano()
)

func FromTime(t time.Time) time.Time {
	return epoch.Add(t.Sub(epoch))
}

func Now() time.Time {
	return FromTime(time.Now())
}

func UnixNano() int64 {
	return epochNano + int64(time.Since(epoch))
}

func UnixMicro() int64 {
	return (epochNano + int64(time.Since(epoch))) / 1000
}
