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

package utils

import (
	"runtime"
	"sync"

	"go.uber.org/atomic"
)

// parallelState holds all shared state for one parallel ParallelExec call in a
// single heap object. Keeping it in one struct (rather than separate local
// variables captured by the goroutines) means the workers can share it through
// one pointer, so launching them allocates only this struct plus a single
// reused worker funcval instead of one object per goroutine.
type parallelState[T any] struct {
	vals  []T
	fn    func(T)
	step  uint64
	end   uint64
	start atomic.Uint64
	wg    sync.WaitGroup
}

func (s *parallelState[T]) run() {
	defer s.wg.Done()
	for {
		n := s.start.Add(s.step)
		if n >= s.end+s.step {
			return
		}

		for i := n - s.step; i < n && i < s.end; i++ {
			s.fn(s.vals[i])
		}
	}
}

// ParallelExec will executes the given function with each element of vals, if len(vals) >= parallelThreshold,
// will execute them in parallel, with the given step size. So fn must be thread-safe.
func ParallelExec[T any](vals []T, parallelThreshold, step uint64, fn func(T)) {
	if uint64(len(vals)) < parallelThreshold {
		for _, v := range vals {
			fn(v)
		}
		return
	}

	// parallel - enables much more efficient multi-core utilization
	numCPU := runtime.NumCPU()
	if numCPU > len(vals) {
		numCPU = len(vals)
	}

	st := &parallelState[T]{
		vals: vals,
		fn:   fn,
		step: step,
		end:  uint64(len(vals)),
	}
	st.wg.Add(numCPU)

	// Hoist the worker funcval out of the spawn loop so every `go worker()`
	// reuses it. A `go` statement on a method value or generic function would
	// otherwise allocate a funcval (carrying the type dictionary) per goroutine.
	worker := st.run
	for p := 0; p < numCPU; p++ {
		go worker()
	}
	st.wg.Wait()
}
