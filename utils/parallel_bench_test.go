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

package utils

import (
	"fmt"
	"sync"
	"testing"

	"go.uber.org/atomic"
)

// TestParallelExecConcurrent exercises the shared work-stealing state under many
// concurrent callers: every element must be visited exactly once per call.
func TestParallelExecConcurrent(t *testing.T) {
	const n = 137
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for iter := 0; iter < 300; iter++ {
				vals := make([]int, n)
				counts := make([]atomic.Uint32, n)
				for i := range vals {
					vals[i] = i
				}
				ParallelExec(vals, 1, 2, func(i int) { counts[i].Add(1) })
				for i := 0; i < n; i++ {
					if got := counts[i].Load(); got != 1 {
						t.Errorf("index %d visited %d times", i, got)
						return
					}
				}
			}
		}()
	}
	wg.Wait()
}

var benchSink atomic.Uint64

// perItemWork approximates a small per-element cost (~0.2us) so the benchmark
// reflects real fan-out rather than pure loop overhead.
func perItemWork() {
	var acc uint64
	for i := 0; i < 250; i++ {
		acc = acc*1099511628211 + uint64(i)
	}
	benchSink.Add(acc)
}

func benchmarkParallelExec(b *testing.B, fanout int, concurrent bool) {
	vals := make([]int, fanout)
	fn := func(i int) { perItemWork() }

	b.ResetTimer()
	b.ReportAllocs()
	if concurrent {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ParallelExec(vals, 1, 2, fn)
			}
		})
		return
	}
	for n := 0; n < b.N; n++ {
		ParallelExec(vals, 1, 2, fn)
	}
}

func BenchmarkParallelExec(b *testing.B) {
	for _, fanout := range []int{50, 200} {
		b.Run(fmt.Sprintf("single/f%d", fanout), func(b *testing.B) {
			benchmarkParallelExec(b, fanout, false)
		})
		b.Run(fmt.Sprintf("concurrent/f%d", fanout), func(b *testing.B) {
			benchmarkParallelExec(b, fanout, true)
		})
	}
}
