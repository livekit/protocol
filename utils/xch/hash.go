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

package xch

import "math/bits"

// SEE: https://github.com/ClickHouse/ClickHouse/blob/master/src/Common/HashTable/Hash.h

func intHash32(x uint64) uint32 {
	x = (^x) + (x << 18)
	x = x ^ (bits.RotateLeft64(x, -31))
	x = x * 21
	x = x ^ (bits.RotateLeft64(x, -11))
	x = x + (x << 6)
	x = x ^ (bits.RotateLeft64(x, -22))
	return uint32(x)
}

func intHash64(x uint64) uint64 {
	x ^= x >> 33
	x *= 0xff51afd7ed558ccd
	x ^= x >> 33
	x *= 0xc4ceb9fe1a85ec53
	x ^= x >> 33
	return x
}

func trivialHash(x uint64) uint64 {
	return x
}
