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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/livekit/protocol/logger"
)

const (
	memUsagePathV1          = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
	memLimitPathV1          = "/sys/fs/cgroup/memory/memory.limit_in_bytes"
	memStatPathV1           = "/sys/fs/cgroup/memory/memory.stat"
	totalInactiveFilePathV1 = "total_inactive_file"

	memCurrentPathV2   = "/sys/fs/cgroup/memory.current"
	memMaxPathV2       = "/sys/fs/cgroup/memory.max"
	memStatPathV2      = "/sys/fs/cgroup/memory.stat"
	inactiveFilePathV2 = "inactive_file"
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
		memUsagePathV1:   newMemInfoGetterV1,
		memCurrentPathV2: newMemInfoGetterV2,
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
	totalInactiveFile, err := readValueOfKeyFromFile(memStatPathV1, totalInactiveFileV1)
	if err == nil && usage > totalInactiveFile {
		usage -= totalInactiveFile
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
	usage, err := readValueFromFile(memCurrentPathV2)
	if err != nil {
		return 0, 0, err
	}
	inactiveFile, err := readValueOfKeyFromFile(memStatPathV2, inactiveFileV2)
	if err == nil && usage > inactiveFile {
		usage -= inactiveFile
	}

	total, err := readValueFromFile(memMaxPathV2)
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

	// Skip the trailing EOL
	return strconv.ParseUint(string(b[:len(b)-1]), 10, 64)
}

func readValueOfKeyFromFile(file string, key string) (uint64, error) {
	fd, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var memory Stats
	memStats := map[string]*uint64{
		"MemTotal":     &memory.Total,
		"MemFree":      &memory.Free,
		"MemAvailable": &memory.Available,
		"Buffers":      &memory.Buffers,
		"Cached":       &memory.Cached,
		"Active":       &memory.Active,
		"Inactive":     &memory.Inactive,
		"SwapCached":   &memory.SwapCached,
		"SwapTotal":    &memory.SwapTotal,
		"SwapFree":     &memory.SwapFree,
		"Mapped":       &memory.Mapped,
		"Shmem":        &memory.Shmem,
		"Slab":         &memory.Slab,
		"PageTables":   &memory.PageTables,
		"Committed_AS": &memory.Committed,
		"VmallocUsed":  &memory.VmallocUsed,
	}
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.IndexRune(line, ':')
		if i < 0 {
			continue
		}

		field := line[:i]
		if field != key {
			continue
		}

		val := strings.TrimSpace(line[i+1:])
		if v, err := strconv.ParseUint(val, 10, 64); err == nil {
			return v, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scan error for %s, key: %s: %s", file, key, err)
	}

	return 0, nil
}
