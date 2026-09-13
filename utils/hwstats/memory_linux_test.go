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
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCgroupCharged      = uint64(3 << 30)
	testCgroupInactiveFile = uint64(1 << 30)
	testCgroupWorkingSet   = testCgroupCharged - testCgroupInactiveFile
)

// writeCgroupFiles points the given path variables at freshly written files for the duration of the test.
func writeCgroupFiles(t *testing.T, usagePath, limitPath, statPath *string, limit, inactiveFileKey string) {
	dir := t.TempDir()
	files := map[*string]string{
		usagePath: strconv.FormatUint(testCgroupCharged, 10) + "\n",
		limitPath: limit + "\n",
		statPath:  "anon 1024\n" + inactiveFileKey + " " + strconv.FormatUint(testCgroupInactiveFile, 10) + "\n",
	}
	for ptr, content := range files {
		path := filepath.Join(dir, filepath.Base(*ptr))
		require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
		orig := *ptr
		*ptr = path
		t.Cleanup(func() { *ptr = orig })
	}
}

func nodeMemory(t *testing.T) uint64 {
	osStat, err := newOSStatMemoryGetter()
	require.NoError(t, err)
	_, total, err := osStat.getMemory()
	require.NoError(t, err)
	return total
}

func TestMemInfoGetterV2(t *testing.T) {
	nodeTotal := nodeMemory(t)
	osStat, err := newOSStatMemoryGetter()
	require.NoError(t, err)
	cg := newMemInfoGetterV2(osStat)

	t.Run("limit set", func(t *testing.T) {
		// below physical memory so the cgroup limit is reported as is
		limit := nodeTotal / 2
		writeCgroupFiles(t, &memCurrentPathV2, &memMaxPathV2, &memStatPathV2, strconv.FormatUint(limit, 10), inactiveFileV2)
		usage, total, err := cg.getMemory()
		require.NoError(t, err)
		require.Equal(t, testCgroupWorkingSet, usage)
		require.Equal(t, limit, total)
	})

	t.Run("unlimited keeps cgroup usage", func(t *testing.T) {
		writeCgroupFiles(t, &memCurrentPathV2, &memMaxPathV2, &memStatPathV2, "max", inactiveFileV2)
		usage, total, err := cg.getMemory()
		require.NoError(t, err)
		require.Equal(t, testCgroupWorkingSet, usage)
		require.Equal(t, nodeTotal, total)
	})
}

func TestMemInfoGetterV1(t *testing.T) {
	nodeTotal := nodeMemory(t)
	osStat, err := newOSStatMemoryGetter()
	require.NoError(t, err)
	cg := newMemInfoGetterV1(osStat)

	t.Run("limit set", func(t *testing.T) {
		// below physical memory so the cgroup limit is reported as is
		limit := nodeTotal / 2
		writeCgroupFiles(t, &memUsagePathV1, &memLimitPathV1, &memStatPathV1, strconv.FormatUint(limit, 10), totalInactiveFileV1)
		usage, total, err := cg.getMemory()
		require.NoError(t, err)
		require.Equal(t, testCgroupWorkingSet, usage)
		require.Equal(t, limit, total)
	})

	t.Run("limit above physical keeps cgroup usage", func(t *testing.T) {
		// cgroup v1 reports an unlimited cgroup as a limit far above physical memory
		writeCgroupFiles(t, &memUsagePathV1, &memLimitPathV1, &memStatPathV1, "9223372036854771712", totalInactiveFileV1)
		usage, total, err := cg.getMemory()
		require.NoError(t, err)
		require.Equal(t, testCgroupWorkingSet, usage)
		require.Equal(t, nodeTotal, total)
	})
}
