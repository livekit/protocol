// Copyright 2026 LiveKit, Inc.
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

package decodebudget

import (
	"slices"
)

// The cost model overestimates on purpose: a budget that admits a payload must
// not let its decode allocate more than the estimate. This file models the Go
// allocator; generated.go and dynamic.go model what each decoder asks it for.

// Go runtime size classes; an allocation is rounded up to the next one.
var goSizeClasses = []int{
	8, 16, 24, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256,
	288, 320, 352, 384, 416, 448, 480, 512, 576, 640, 704, 768, 896, 1024, 1152, 1280,
	1408, 1536, 1792, 2048, 2304, 2688, 3072, 3200, 3456, 4096, 4864, 5376, 6144, 6528,
	6784, 6912, 8192, 9472, 9728, 10240, 10880, 12288, 13568, 14336, 16384, 18432,
	19072, 20480, 21760, 24576, 27264, 28672, 32768,
}

// sizeClassByWord maps ceil(n/8) to n's size class, since every class is a
// multiple of 8, to spare allocSize a search on the hot path.
var sizeClassByWord = func() []uint16 {
	maxSmall := goSizeClasses[len(goSizeClasses)-1]
	table := make([]uint16, maxSmall/8+1)
	for i := range table {
		c, _ := slices.BinarySearch(goSizeClasses, i*8)
		table[i] = uint16(goSizeClasses[c])
	}
	return table
}()

func allocSize(n int) int {
	if n <= 0 {
		return 0
	}
	if i := (n + 7) >> 3; i < len(sizeClassByWord) {
		return int(sizeClassByWord[i])
	}
	const page = 8192
	return (n + page - 1) / page * page
}

// pointerAlloc is allocSize for an object holding pointers, which carries an
// 8-byte malloc header between 512 bytes and 32 KiB, and the usable size.
func pointerAlloc(n int) (size, usable int) {
	const minHeader, maxSmall, header = 512, 32768, 8
	if n > minHeader && n <= maxSmall-header {
		size = allocSize(n + header)
		return size, size - header
	}
	size = allocSize(n)
	return size, size
}

// tmpCopySize is the allocation for a temporary copy of n bytes; one that does
// not escape and fits runtime.tmpBuf stays on the stack.
func tmpCopySize(n int) int {
	if n <= 32 {
		return 0
	}
	return allocSize(n)
}

// appendGrow is the allocation for appending add elements of elem bytes to a
// slice of length l and capacity c, and the resulting capacity, following
// runtime.growslice.
func appendGrow(l, c, add, elem int) (mem, newCap int) {
	newCap = nextSliceCap(l+add, c)
	if newCap == c {
		return 0, c
	}
	mem = allocSize(newCap * elem)
	return mem, mem / elem
}

// nextSliceCap is the capacity growslice requests, before size-class rounding,
// to hold newLen elements in a slice of capacity c.
func nextSliceCap(newLen, c int) int {
	if newLen <= c {
		return c
	}
	newCap := c + c
	switch {
	case newLen > newCap:
		newCap = newLen
	case c >= 256:
		newCap = c
		for newCap < newLen {
			newCap += (newCap + 3*256) >> 2
		}
	}
	return newCap
}

// appendAllocBytes is the total size of the backing arrays append allocates
// while growing a nil slice to n elements one at a time.
func appendAllocBytes(n, elem int) int {
	total, c := 0, 0
	for l := 0; l < n; l++ {
		mem, newCap := appendGrow(l, c, 1, elem)
		total += mem
		c = newCap
	}
	return total
}

// Go's map layout (internal/runtime/maps): a header, then one group of 8
// slots, then a table that doubles at 7/8 load up to 1024 slots, after which
// each full table splits into two new ones.
const (
	mapHeaderSize    = 48
	mapTableSize     = 32
	mapGroupSlots    = 8
	mapMaxTableSlots = 1024
)

// mapGrowth is the allocation for inserting a new key into a map holding n
// entries, excluding the header, for slots of slot bytes.
func mapGrowth(n, slot int, pointers bool) int {
	group := mapGroupSlots + mapGroupSlots*slot
	alloc := func(n int) int {
		if !pointers {
			return allocSize(n)
		}
		size, _ := pointerAlloc(n)
		return size
	}
	table := func(slots int) int {
		return allocSize(mapTableSize) + alloc(slots/mapGroupSlots*group)
	}
	switch {
	case n == 0:
		return alloc(group)
	case n < mapGroupSlots:
		return 0
	case n == mapGroupSlots:
		// The first table comes with a one-entry directory.
		return table(2*mapGroupSlots) + allocSize(8)
	case n < mapMaxTableSlots*7/8:
		for c := 2 * mapGroupSlots; c < mapMaxTableSlots; c *= 2 {
			if n == c*7/8 {
				return table(2 * c)
			}
		}
		return 0
	}
	return mapSplitCost(n, table(mapMaxTableSlots))
}

// mapSplitCost charges the round of splits that takes a map from t to 2t
// full-size tables, plus the directory that doubles for it. Tables fill
// unevenly, so a round with several tables is spread over the inserts from
// 3/4 to 7/8 average load.
func mapSplitCost(n, table int) int {
	t := 1
	for n >= 2*t*mapMaxTableSlots*3/4 {
		t *= 2
	}
	full := t * mapMaxTableSlots * 7 / 8
	begin := full
	if t > 1 {
		begin = t * mapMaxTableSlots * 3 / 4
	}
	if n < begin || n > full {
		return 0
	}
	round := 2*t*table + allocSize(2*t*8)
	span := full - begin + 1
	return round*(n-begin+1)/span - round*(n-begin)/span
}
