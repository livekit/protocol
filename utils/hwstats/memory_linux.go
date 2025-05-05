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

//go:build linux

package hwstats

import (
	"os"
	"strconv"

	"github.com/livekit/protocol/logger"
)

const (
	memStatsPathV1 = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
	memStatsPathV2 = "/sys/fs/memory/memory.current"

	memUsagePathV1 = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
	memLimitPathV1 = "/sys/fs/cgroup/memory/memory.limit_in_bytes"

	memCurrentPathV2 = "/sys/fs/cgroup/memory.current"
	memMaxPathV2     = "/sys/fs/cgroup/memory.max"
)

type memInfoGetter interface {
	getMemory() (uint64, uint64, error)
}

type cgroupMemoryGetter struct {
	cg     memInfoGetter
	osStat *osStatMemoryGetter
}

func newPlatformMemoryGetter() (platformMemoryGetter, error) {
	// backup getter when max is not available,
	// will get from /proc/meminfo which will be node level max though
	osStat, err := newOSStatMemoryGetter()
	if err != nil {
		return nil, err
	}

	// probe for the cgroup version
	var cg memInfoGetter
	for k, v := range map[string]func(osStat *osStatMemoryGetter) memInfoGetter{
		memStatsPathV1: newMemInfoGetterV1,
		memStatsPathV2: newMemInfoGetterV2,
	} {
		e, err := fileExists(k)
		if err != nil {
			return nil, err
		}
		if e {
			logger.Infow("using getter", "key", k) // REMOVE
			cg = v(osStat)
			break
		}
	}
	if cg == nil {
		logger.Infow("failed reading cgroup specific memory stats, falling back to system wide implementation")
		return osStat, nil
	}

	return &cgroupMemoryGetter{
		cg:     cg,
		osStat: osStat,
	}, nil
}

func (c *cgroupMemoryGetter) getMemory() (uint64, uint64, error) {
	return c.cg.getMemory()
}

// -----------------------------------------

type memInfoGetterV1 struct {
	osStat *osStatMemoryGetter
}

func newMemInfoGetterV1(osStat *osStatMemoryGetter) memInfoGetter {
	return &memInfoGetterV1{
		osStat: osStat,
	}
}

func (cg *memInfoGetterV1) getMemory() (uint64, uint64, error) {
	usage, err := readValueFromFile(memUsagePathV1)
	if err != nil {
		return 0, 0, err
	}

	total, err := readValueFromFile(memLimitPathV1)
	if err != nil {
		return 0, 0, err
	}

	// fallback if limit from cgroup is more than physical available memory
	// when limit is not set explicitly, it could be very high
	usage1, total1, err := cg.osStat.getMemory()
	if err != nil {
		return 0, 0, err
	}
	if total > total1 {
		return usage1, total1, nil
	}

	return usage, total, nil
}

// --------------------------------------------

type memInfoGetterV2 struct {
	osStat *osStatMemoryGetter
}

func newMemInfoGetterV2(osStat *osStatMemoryGetter) memInfoGetter {
	return &memInfoGetterV2{
		osStat: osStat,
	}
}

func (cg *memInfoGetterV2) getMemory() (uint64, uint64, error) {
	usage, err1 := readValueFromFile(memCurrentPathV2)
	if err1 != nil {
		return 0, 0, err1
	}

	total, err := readValueFromFile(memMaxPathV2)
	logger.Infow("memory v2 usage", "usage", usage, "err1", err1, "total", total, "err", err) // REMOVE
	if err != nil {
		// when memory limit not set, it has the string "max"
		usage, total, err = cg.osStat.getMemory()
		if err != nil {
			return 0, 0, err
		}
	}

	return usage, total, nil
}

// ---------------------------------

func readValueFromFile(file string) (uint64, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}

	logger.Infow("read file", "file", file, "value", string(b), "truncated", string(b[:len(b)-1])) // REMOVE
	// Skip the trailing EOL
	return strconv.ParseUint(string(b[:len(b)-1]), 10, 64)
}
