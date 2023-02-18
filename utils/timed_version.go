package utils

import (
	"fmt"
	"sync"
	"time"
	_ "unsafe" // required for linkname

	"go.uber.org/atomic"

	"github.com/livekit/protocol/livekit"
)

//go:linkname nanotime1 runtime.nanotime1
func nanotime1() int64

//go:linkname usleep runtime.usleep
func usleep(usec uint32)

const tickBits uint64 = 12
const tickMask uint64 = (1 << tickBits) - 1

var epoch = uint64(time.Date(2000, 0, 0, 0, 0, 0, 0, time.UTC).UnixMicro())

type TimedVersionGenerator interface {
	New() *TimedVersion
	Next() TimedVersion
}

func timedVersionComponents(v uint64) (ts int64, ticks int32) {
	return int64(v>>tickBits) + int64(epoch), int32(v & tickMask)
}

func timedVersionFromComponents(ts int64, ticks int32) TimedVersion {
	return TimedVersion{v: *atomic.NewUint64((uint64(ts-int64(epoch)) << tickBits) | uint64(ticks))}
}

type timedVersionGenerator struct {
	mu    sync.Mutex
	ts    uint64
	ticks uint64
}

func NewDefaultTimedVersionGenerator() TimedVersionGenerator {
	return new(timedVersionGenerator)
}

func (g *timedVersionGenerator) New() *TimedVersion {
	v := g.Next()
	return &v
}

func (g *timedVersionGenerator) Next() TimedVersion {
	now := uint64(time.Now().UnixMicro()) - epoch

	g.mu.Lock()
	defer g.mu.Unlock()

	for {
		if now < g.ts {
			now = g.ts
		}
		if g.ts == now {
			// if incrementing the ticks would overflow the version sleep for a
			// microsecond then try again.
			if g.ticks == tickMask {
				usleep(1)
				now = uint64(time.Now().UnixMicro()) - epoch
				continue
			}
			g.ticks++
		} else {
			g.ts = now
			g.ticks = 0
		}
		return TimedVersion{v: *atomic.NewUint64(now<<tickBits | g.ticks)}
	}
}

type TimedVersion struct {
	v atomic.Uint64
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
	ov := other.v.Load()
	for {
		prev := t.v.Load()
		if ov <= prev {
			return
		}
		if t.v.CompareAndSwap(prev, ov) {
			break
		}
	}
}

func (t *TimedVersion) Store(other *TimedVersion) {
	t.v.Store(other.v.Load())
}

func (t *TimedVersion) Load() TimedVersion {
	return TimedVersion{v: *atomic.NewUint64(t.v.Load())}
}

func (t *TimedVersion) After(other *TimedVersion) bool {
	return t.v.Load() > other.v.Load()
}

func (t *TimedVersion) Compare(other *TimedVersion) int {
	ov := other.v.Load()
	v := t.v.Load()
	if v < ov {
		return -1
	}
	if v == ov {
		return 0
	}
	return 1
}

func (t *TimedVersion) ToProto() *livekit.TimedVersion {
	ts, ticks := timedVersionComponents(t.v.Load())
	return &livekit.TimedVersion{
		UnixMicro: ts,
		Ticks:     ticks,
	}
}

func (t *TimedVersion) String() string {
	ts, ticks := timedVersionComponents(t.v.Load())
	return fmt.Sprintf("%d.%d", ts, ticks)
}
