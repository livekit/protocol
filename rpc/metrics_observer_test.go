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

package rpc

import (
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"github.com/livekit/psrpc"
)

// TestRequestObserverMetrics asserts the server-side lifecycle events register
// and emit. These are the only signals available for a request whose handler is
// never invoked, so a silent regression here would be invisible in production.
func TestRequestObserverMetrics(t *testing.T) {
	InitPSRPCStats(prometheus.Labels{})
	o := PSRPCMetricsObserver{}
	info := psrpc.RPCInfo{Service: "TestSvc", Method: "TestMethod"}

	o.OnRequestReceived(info)
	o.OnRequestExpired(info, 20*time.Millisecond)
	o.OnClaim(info, psrpc.ClaimGranted, 3*time.Millisecond)
	o.OnClaim(info, psrpc.ClaimTimedOut, 1005*time.Millisecond)

	got := gatherPSRPCCounts(t)
	require.Equal(t, 1.0, got["livekit_psrpc_requests_received_total"])
	require.Equal(t, 1.0, got["livekit_psrpc_requests_expired_total"])
	require.Equal(t, 1.0, got["livekit_psrpc_claim_total|granted"])
	require.Equal(t, 1.0, got["livekit_psrpc_claim_total|timed_out"])
	require.Equal(t, 1.0, got["livekit_psrpc_claim_wait_time_ms|timed_out"])
}

// gatherPSRPCCounts returns counter values and histogram sample counts for
// livekit_psrpc_* series, keyed by name and outcome label where present.
func gatherPSRPCCounts(t *testing.T) map[string]float64 {
	t.Helper()
	mfs, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	out := map[string]float64{}
	for _, mf := range mfs {
		if !strings.HasPrefix(mf.GetName(), "livekit_psrpc_") {
			continue
		}
		for _, m := range mf.GetMetric() {
			key := mf.GetName()
			for _, l := range m.GetLabel() {
				if l.GetName() == "outcome" {
					key += "|" + l.GetValue()
				}
			}
			if c := m.GetCounter(); c != nil {
				out[key] += c.GetValue()
			}
			if h := m.GetHistogram(); h != nil {
				out[key] += float64(h.GetSampleCount())
			}
		}
	}
	return out
}
