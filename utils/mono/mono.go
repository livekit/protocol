package mono

import "time"

var (
	epoch     = time.Now()
	epochNano = epoch.UnixNano()
)

func Now() time.Time {
	return epoch.Add(time.Since(epoch))
}

func UnixNano() int64 {
	return epochNano + int64(time.Since(epoch))
}
