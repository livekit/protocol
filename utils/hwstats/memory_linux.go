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
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/livekit/protocol/logger"
)

const (
	memStatsPathV1 = "/sys/fs/memory/memory.usage_in_bytes"
	memStatsPathV2 = "/sys/fs/memory/memory.current"

	memUsagePathV1 = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
	memLimitPathV1 = "/sys/fs/cgroup/memory/memory.limit_in_bytes"

	memCurrentPathV2 = "sys/fs/cgroup/memory/memory.current"
	memMaxPathV2     = "sys/fs/cgroup/memory/memory.max"
)

type memInfoGetter interface {
	getMemory() (uint64, uint64, error)
}

type cgroupMemoryGetter struct {
	cg memInfoGetter
	os *osStatMemoryGetter
}

func newPlatformMemoryGetter() (platformMemoryGetter, error) {
	// backup getter when max is not available,
	// will get from /proc/meminfo which will be node level max though
	osStat, err := newOSStatMemoryMonitor()
	if err != nil {
		return nil, err
	}

	// probe for the cgroup version
	var cg memInfoGetter
	for k, v := range map[string]func() memInfoGetter{
		memStatsPathV1: newMemInfoGetterV1,
		memStatsPathV2: newMemInfoGetterV2,
	} {
		e, err := fileExists(k)
		if err != nil {
			return nil, err
		}
		if e {
			cg = v(osStat)
			break
		}
	}
	if cg == nil {
		logger.Infow("failed reading cgroup specific memory stats, falling back to system wide implementation")
		return osStat, nil
	}

	return &cgroupMemoryMonitor{
		cg:     cg,
		osStat: osStat,
	}, nil
}

func (c *cgroupMemoryGetter) getMemory() (uint64, uint64, error) {
	return p.cg.getMemory()
}

// -----------------------------------------

const (
	cSaneMemoryLimit = 32 * 1024 * 1024 * 1024 // 32 GB
)

type memInfoGetterV1 struct{}

func newMemInfoGetterV1() memInfoGetter {
	return &memInfoGetterV1{}
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

	if total > cSaneMemoryLimit {
		usage, total, err = cg.osStat.getMemory()
		if err != nil {
			return 0, 0, err
		}
	}

	return usage, total, nil
}

// --------------------------------------------

type memInfoGetterV2 struct {
	osStat *osStatsMemoryGetter
}

func newMemInfoGetterV2(osStat *osStatMemoryGetter) memInfoGetter {
	return &memInfoGetterV2{
		osStat: osStat,
	}
}

func (cg *memInfoGetterV2) getMemory() (uint64, uint64, error) {
	usage, err := readValueFromFile(memCurrentPathV2)
	if err != nil {
		return 0, 0, err
	}

	total, err := readValueFromFile(memMaxPathV2)
	if err != nil {
		if maxVal, err := readStringFromFile(memMaxPathV2); err != nil {
			return 0, 0, err
		} else if maxVal == "max" {
			// memory limit not set
			usage, total, err = cg.osStat.getMemory()
			if err != nil {
				return 0, 0, err
			}
		}
	}

	return usage, total, nil
}

// ---------------------------------

func fileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	default:
		return false, err
	}
}

func readValueFromFile(file string) (uint64, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return 0, 0, err
	}

	// Skip the trailing EOL
	return strconv.ParseUint(string(b[:len(b)-1]), 10, 64)
}

func readStringFromFile(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return 0, 0, err
	}

	// Remove trailing new line if any
	s := strings.TrimSuffix(string(b), "\n")

	// Remove trailing space if any
	return strings.TrimSuffix(s, " ")
}
