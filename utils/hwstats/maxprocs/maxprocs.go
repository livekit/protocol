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

// Package maxprocs caps GOMAXPROCS to the Kubernetes CPU request so the Go
// scheduler doesn't create more OS threads than the pod is guaranteed.
//
// # Why not CPU limits / automaxprocs
//
// Real-time media workloads (egress, ingress) intentionally omit CPU limits
// to avoid CFS throttling — even brief throttle events cause audio glitches
// and video frame drops. Without a limit the cgroup quota is "max", so
// libraries like uber/automaxprocs fall back to runtime.NumCPU (the full
// node). This package uses the CPU request instead, which reflects what the
// Kubernetes scheduler actually guarantees to the pod.
//
// When a CPU limit IS set, the cgroup quota gives a tighter bound and Go
// (or automaxprocs) already derives GOMAXPROCS from it. This package still
// works in that case — it caps GOMAXPROCS down to the request but never
// increases it, so the stricter of (limit, request) wins.
//
// # Usage
//
// The CPU request is read from the LIVEKIT_CPU_REQUEST environment variable,
// which should be populated e.g via the Kubernetes Downward API:
//
//	env:
//	  - name: LIVEKIT_CPU_REQUEST
//	    valueFrom:
//	      resourceFieldRef:
//	        containerName: <name>
//	        resource: requests.cpu
//
// Without this variable the package is a no-op. A blank import is
// sufficient — init runs at process startup:
//
//	import _ "github.com/livekit/protocol/utils/hwstats/maxprocs"
package maxprocs

import (
	"math"
	"runtime"
	"sync"

	"github.com/livekit/protocol/utils/hwstats"
)

var once sync.Once

func init() { Install() }

// Install caps GOMAXPROCS to ceil(LIVEKIT_CPU_REQUEST) when that environment
// variable is set and valid. It never increases GOMAXPROCS beyond its current
// value. It is a no-op when the variable is absent. Runs at most once per
// process.
func Install() {
	once.Do(func() {
		req := hwstats.EffectiveCPURequest()
		if req <= 0 {
			return
		}
		cap := int(math.Ceil(req))
		if runtime.GOMAXPROCS(0) > cap {
			runtime.GOMAXPROCS(cap)
		}
	})
}
