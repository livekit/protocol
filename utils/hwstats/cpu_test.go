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

package hwstats

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCPUIdle(t *testing.T) {
	st, err := NewCPUStats(nil)
	require.NoError(t, err)

	idle := st.GetCPUIdle()
	require.GreaterOrEqual(t, 0.0, idle)
}

func TestAggregateMemoryStats(t *testing.T) {
	const pageSize = 4096
	const selfPID = 1

	t.Run("self process only", func(t *testing.T) {
		entries := []procEntry{
			{pid: 1, ppid: 0, comm: "main", rss: 100},
		}
		memTotal, memory, groups := aggregateMemoryStats(entries, selfPID, pageSize)

		require.Equal(t, 100*pageSize, memTotal)
		require.Len(t, memory, 1)
		require.Equal(t, 100*pageSize, memory[1].Total)
		require.Len(t, memory[1].Procs, 1)
		require.Equal(t, ProcMemoryEntry{Name: "main", Memory: 100 * pageSize}, memory[1].Procs[1])
		require.Equal(t, 1, groups[1])
	})

	t.Run("direct children form separate groups", func(t *testing.T) {
		entries := []procEntry{
			{pid: 1, ppid: 0, comm: "main", rss: 100},
			{pid: 10, ppid: 1, comm: "child1", rss: 200},
			{pid: 20, ppid: 1, comm: "child2", rss: 300},
		}
		memTotal, memory, groups := aggregateMemoryStats(entries, selfPID, pageSize)

		require.Equal(t, 600*pageSize, memTotal)
		require.Len(t, memory, 3)

		require.Equal(t, 10, groups[10])
		require.Equal(t, 20, groups[20])

		require.Equal(t, 200*pageSize, memory[10].Total)
		require.Len(t, memory[10].Procs, 1)
		require.Equal(t, 300*pageSize, memory[20].Total)
		require.Len(t, memory[20].Procs, 1)
	})

	t.Run("nested processes grouped under first child", func(t *testing.T) {
		// self(1) -> child(10) -> grandchild(100) -> great-grandchild(1000)
		entries := []procEntry{
			{pid: 1, ppid: 0, comm: "main", rss: 50},
			{pid: 10, ppid: 1, comm: "child", rss: 100},
			{pid: 100, ppid: 10, comm: "grandchild", rss: 200},
			{pid: 1000, ppid: 100, comm: "great-grandchild", rss: 300},
		}
		memTotal, memory, groups := aggregateMemoryStats(entries, selfPID, pageSize)

		require.Equal(t, 650*pageSize, memTotal)

		// all descendants grouped under pid 10
		require.Equal(t, 10, groups[10])
		require.Equal(t, 10, groups[100])
		require.Equal(t, 10, groups[1000])

		// group 10 total = 100+200+300
		require.Equal(t, 600*pageSize, memory[10].Total)
		require.Len(t, memory[10].Procs, 3)

		// self is its own group
		require.Equal(t, 50*pageSize, memory[1].Total)
		require.Len(t, memory[1].Procs, 1)
	})

	t.Run("multiple groups with nesting", func(t *testing.T) {
		// self(1) -> child1(10) -> gc1(100), gc2(101)
		//         -> child2(20) -> gc3(200)
		entries := []procEntry{
			{pid: 1, ppid: 0, comm: "main", rss: 10},
			{pid: 10, ppid: 1, comm: "child1", rss: 20},
			{pid: 100, ppid: 10, comm: "gc1", rss: 30},
			{pid: 101, ppid: 10, comm: "gc2", rss: 40},
			{pid: 20, ppid: 1, comm: "child2", rss: 50},
			{pid: 200, ppid: 20, comm: "gc3", rss: 60},
		}
		memTotal, memory, groups := aggregateMemoryStats(entries, selfPID, pageSize)

		require.Equal(t, 210*pageSize, memTotal)

		// group assignments
		require.Equal(t, 10, groups[10])
		require.Equal(t, 10, groups[100])
		require.Equal(t, 10, groups[101])
		require.Equal(t, 20, groups[20])
		require.Equal(t, 20, groups[200])

		// group 10: child1(20) + gc1(30) + gc2(40) = 90
		require.Equal(t, 90*pageSize, memory[10].Total)
		require.Len(t, memory[10].Procs, 3)
		require.Equal(t, ProcMemoryEntry{Name: "child1", Memory: 20 * pageSize}, memory[10].Procs[10])
		require.Equal(t, ProcMemoryEntry{Name: "gc1", Memory: 30 * pageSize}, memory[10].Procs[100])
		require.Equal(t, ProcMemoryEntry{Name: "gc2", Memory: 40 * pageSize}, memory[10].Procs[101])

		// group 20: child2(50) + gc3(60) = 110
		require.Equal(t, 110*pageSize, memory[20].Total)
		require.Len(t, memory[20].Procs, 2)

		// self group
		require.Equal(t, 10*pageSize, memory[1].Total)
		require.Len(t, memory[1].Procs, 1)
	})

	t.Run("unrelated processes group under their top ancestor", func(t *testing.T) {
		// Processes not in self's tree (ppid chain leads to 0)
		entries := []procEntry{
			{pid: 1, ppid: 0, comm: "main", rss: 10},
			{pid: 50, ppid: 0, comm: "unrelated", rss: 100},
			{pid: 500, ppid: 50, comm: "unrelated-child", rss: 200},
		}
		memTotal, memory, groups := aggregateMemoryStats(entries, selfPID, pageSize)

		require.Equal(t, 310*pageSize, memTotal)

		// unrelated processes group under top ancestor (50)
		require.Equal(t, 50, groups[50])
		require.Equal(t, 50, groups[500])

		require.Equal(t, 300*pageSize, memory[50].Total)
		require.Len(t, memory[50].Procs, 2)
	})

	t.Run("empty entries", func(t *testing.T) {
		memTotal, memory, groups := aggregateMemoryStats(nil, selfPID, pageSize)

		require.Equal(t, 0, memTotal)
		require.Empty(t, memory)
		require.Empty(t, groups)
	})
}
