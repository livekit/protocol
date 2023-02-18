package utils

import (
	"fmt"
	"sync/atomic"
	"time"
	_ "unsafe"

	"github.com/livekit/protocol/livekit"
)

//go:linkname nanotime1 runtime.nanotime1
func nanotime1() int64

//go:linkname usleep runtime.usleep
func usleep(usec uint32)

const tickBits uint64 = 16
const timeMask uint64 = 0xffffffffffff0000
const tickMask = ^timeMask
const timeGranularity = 1000

var epoch = time.Now().UnixNano() - nanotime1()

type TimedVersionGenerator interface {
	New() *TimedVersion
	Next() TimedVersion
}

func timedVersionComponents(v TimedVersion) (ts int64, ticks int32) {
	return int64(uint64(v)>>tickBits) + epoch/timeGranularity, int32(uint64(v) & tickMask)
}

func timedVersionFromComponents(ts int64, ticks int32) TimedVersion {
	return TimedVersion(uint64(ts-epoch/timeGranularity)<<tickBits | uint64(ticks))
}

type timedVersionGenerator uint64

func NewDefaultTimedVersionGenerator() TimedVersionGenerator {
	return new(timedVersionGenerator)
}

func (g *timedVersionGenerator) New() *TimedVersion {
	v := g.Next()
	return &v
}

func (g *timedVersionGenerator) Next() TimedVersion {
	var next uint64
	for {
		prev := atomic.LoadUint64((*uint64)(g))
		next = uint64(nanotime1()) / timeGranularity << tickBits

		// if the version timestamp and next timestamp match increment the ticks
		if prev&timeMask == next {
			// if incrementing the ticks would overflow the version sleep for a
			// microsecond then try again
			if prev&tickMask == tickMask {
				usleep(1)
				continue
			}
			next = prev + 1
		}
		if atomic.CompareAndSwapUint64((*uint64)(g), prev, next) {
			break
		}
	}
	return TimedVersion(next)
}

type TimedVersion uint64

func NewTimedVersionFromProto(proto *livekit.TimedVersion) *TimedVersion {
	v := timedVersionFromComponents(proto.UnixMicro, proto.Ticks)
	return &v
}

func NewTimedVersionFromTime(t time.Time) *TimedVersion {
	v := timedVersionFromComponents(t.UnixMicro(), 0)
	return &v
}

func TimedVersionFromProto(proto *livekit.TimedVersion) TimedVersion {
	return timedVersionFromComponents(proto.UnixMicro, proto.Ticks)
}

func TimedVersionFromTime(t time.Time) TimedVersion {
	return timedVersionFromComponents(t.UnixMicro(), 0)
}

func (t *TimedVersion) Update(other *TimedVersion) {
	ov := atomic.LoadUint64((*uint64)(other))
	for {
		prev := atomic.LoadUint64((*uint64)(t))
		if ov < prev {
			return
		}
		if atomic.CompareAndSwapUint64((*uint64)(t), prev, ov) {
			break
		}
	}
}

func (t *TimedVersion) Set(other *TimedVersion) {
	ov := atomic.LoadUint64((*uint64)(other))
	atomic.StoreUint64((*uint64)(t), ov)
}

func (t *TimedVersion) After(other *TimedVersion) bool {
	ov := atomic.LoadUint64((*uint64)(other))
	return atomic.LoadUint64((*uint64)(t)) > ov
}

func (t *TimedVersion) Compare(other *TimedVersion) int {
	ov := atomic.LoadUint64((*uint64)(other))
	v := atomic.LoadUint64((*uint64)(t))
	if v < ov {
		return -1
	}
	if v == ov {
		return 0
	}
	return 1
}

func (t *TimedVersion) ToProto() *livekit.TimedVersion {
	ts, ticks := timedVersionComponents(TimedVersion(atomic.LoadUint64((*uint64)(t))))
	return &livekit.TimedVersion{
		UnixMicro: ts,
		Ticks:     ticks,
	}
}

func (t *TimedVersion) String() string {
	ts, ticks := timedVersionComponents(TimedVersion(atomic.LoadUint64((*uint64)(t))))
	return fmt.Sprintf("%d.%d", ts, ticks)
}
