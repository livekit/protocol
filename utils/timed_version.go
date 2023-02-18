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

type nocmp [0]func()

const tickBits uint64 = 16
const tickMask uint64 = 0xffff
const timeMask = ^tickMask
const timeGranularity = 1000

var epoch = time.Now().UnixNano() - nanotime1()

type TimedVersionGenerator interface {
	New() *TimedVersion
	Next() TimedVersion
}

func timedVersionComponents(t TimedVersion) (ts int64, ticks int32) {
	return int64(uint64(t.v)>>tickBits) + epoch/timeGranularity, int32(uint64(t.v) & tickMask)
}

func timedVersionFromComponents(ts int64, ticks int32) TimedVersion {
	return TimedVersion{v: uint64(ts-epoch/timeGranularity)<<tickBits | uint64(ticks)}
}

type timedVersionGenerator struct {
	_ nocmp // prevent comparison

	v uint64
}

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
		prev := atomic.LoadUint64(&g.v)
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
		if atomic.CompareAndSwapUint64(&g.v, prev, next) {
			break
		}
	}
	return TimedVersion{v: next}
}

type TimedVersion struct {
	_ nocmp // prevent comparison

	v uint64
}

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
	ov := atomic.LoadUint64(&other.v)
	for {
		prev := atomic.LoadUint64(&t.v)
		if ov < prev {
			return
		}
		if atomic.CompareAndSwapUint64(&t.v, prev, ov) {
			break
		}
	}
}

func (t *TimedVersion) Store(other *TimedVersion) {
	ov := atomic.LoadUint64(&other.v)
	atomic.StoreUint64(&t.v, ov)
}

func (t *TimedVersion) Load() TimedVersion {
	return TimedVersion{v: atomic.LoadUint64(&t.v)}
}

func (t *TimedVersion) After(other *TimedVersion) bool {
	ov := atomic.LoadUint64(&other.v)
	return atomic.LoadUint64(&t.v) > ov
}

func (t *TimedVersion) Compare(other *TimedVersion) int {
	ov := atomic.LoadUint64(&other.v)
	v := atomic.LoadUint64(&t.v)
	if v < ov {
		return -1
	}
	if v == ov {
		return 0
	}
	return 1
}

func (t *TimedVersion) ToProto() *livekit.TimedVersion {
	ts, ticks := timedVersionComponents(TimedVersion{v: atomic.LoadUint64(&t.v)})
	return &livekit.TimedVersion{
		UnixMicro: ts,
		Ticks:     ticks,
	}
}

func (t *TimedVersion) String() string {
	ts, ticks := timedVersionComponents(TimedVersion{v: atomic.LoadUint64(&t.v)})
	return fmt.Sprintf("%d.%d", ts, ticks)
}
