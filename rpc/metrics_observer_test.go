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
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/middleware"
)

// TestRequestObserverMetrics asserts the server-side lifecycle events register
// and emit. These are the only signals available for a request whose handler is
// never invoked, so a silent regression here would be invisible in production.
func TestRequestObserverMetrics(t *testing.T) {
	InitPSRPCStats(prometheus.Labels{})
	o := PSRPCMetricsObserver{}
	info := psrpc.RPCInfo{Service: "LifecycleSvc", Method: "TestMethod"}

	o.OnRequestReceived(info)
	o.OnRequestExpired(info, 20*time.Millisecond)
	o.OnClaim(info, psrpc.ClaimGranted, 3*time.Millisecond)
	o.OnClaim(info, psrpc.ClaimTimedOut, 1005*time.Millisecond)

	got := gatherPSRPCSeries(t, "LifecycleSvc")
	require.Equal(t, 1.0, got["livekit_psrpc_requests_received_total"])
	require.Equal(t, 1.0, got["livekit_psrpc_requests_expired_total"])
	require.Equal(t, 1.0, got["livekit_psrpc_claim_total|granted"])
	require.Equal(t, 1.0, got["livekit_psrpc_claim_total|timed_out"])
	require.Equal(t, 1.0, got["livekit_psrpc_claim_wait_time_ms|granted"])
	require.Equal(t, 1.0, got["livekit_psrpc_claim_wait_time_ms|timed_out"])
}

// TestMetricsObserverMetrics covers the interceptor-driven series. Each method
// routes to a different metric depending on whether the call errored, so the
// error and success paths are asserted separately.
func TestMetricsObserverMetrics(t *testing.T) {
	InitPSRPCStats(prometheus.Labels{})
	o := PSRPCMetricsObserver{}
	info := psrpc.RPCInfo{Service: "ObserverSvc", Method: "TestMethod"}
	boom := errors.New("boom")

	o.OnUnaryRequest(middleware.ClientRole, info, 5*time.Millisecond, nil, 10, 20)
	o.OnUnaryRequest(middleware.ClientRole, info, 5*time.Millisecond, boom, 1, 2)
	o.OnMultiRequest(middleware.ServerRole, info, 7*time.Millisecond, 2, 0, 30, 40)
	o.OnMultiRequest(middleware.ServerRole, info, 7*time.Millisecond, 0, 1, 0, 0)
	o.OnStreamSend(middleware.ClientRole, info, 3*time.Millisecond, nil, 50)
	o.OnStreamRecv(middleware.ClientRole, info, nil, 60)
	o.OnStreamOpen(middleware.ServerRole, info)
	o.OnStreamOpen(middleware.ServerRole, info)
	o.OnStreamClose(middleware.ServerRole, info)

	got := gatherPSRPCSeries(t, "ObserverSvc")

	require.Equal(t, 1.0, got["livekit_psrpc_request_time_ms|client|rpc"])
	require.Equal(t, 1.0, got["livekit_psrpc_error_total|client|rpc"])
	require.Equal(t, 1.0, got["livekit_psrpc_request_time_ms|server|multirpc"])
	require.Equal(t, 1.0, got["livekit_psrpc_error_total|server|multirpc"])
	require.Equal(t, 1.0, got["livekit_psrpc_stream_send_time_ms|client"])
	require.Equal(t, 1.0, got["livekit_psrpc_stream_receive_total|client"])

	// stream_count is a gauge: two opens and one close leave one stream live.
	require.Equal(t, 1.0, got["livekit_psrpc_stream_count|server"])

	require.Equal(t, 11.0, got["livekit_psrpc_bytes_total|client|rpc|rx"])
	require.Equal(t, 22.0, got["livekit_psrpc_bytes_total|client|rpc|tx"])
	require.Equal(t, 30.0, got["livekit_psrpc_bytes_total|server|multirpc|rx"])
	require.Equal(t, 40.0, got["livekit_psrpc_bytes_total|server|multirpc|tx"])
	require.Equal(t, 60.0, got["livekit_psrpc_bytes_total|client|stream|rx"])
	require.Equal(t, 50.0, got["livekit_psrpc_bytes_total|client|stream|tx"])
}

// discriminatingLabels are appended to each key in the order listed, so a key
// reads livekit_psrpc_bytes_total|client|rpc|rx.
var discriminatingLabels = []string{"role", "kind", "direction", "outcome"}

// gatherPSRPCSeries returns livekit_psrpc_* values for one service: counter and
// gauge values, and sample counts for histograms. Filtering on service keeps
// tests in this package independent — the registry is global and accumulates
// across them, so a shared key would make assertions order-dependent.
func gatherPSRPCSeries(t *testing.T, service string) map[string]float64 {
	t.Helper()
	mfs, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	out := map[string]float64{}
	for _, mf := range mfs {
		if !strings.HasPrefix(mf.GetName(), "livekit_psrpc_") {
			continue
		}
		for _, m := range mf.GetMetric() {
			labels := map[string]string{}
			for _, l := range m.GetLabel() {
				labels[l.GetName()] = l.GetValue()
			}
			if labels["service"] != service {
				continue
			}

			key := mf.GetName()
			for _, name := range discriminatingLabels {
				if v, ok := labels[name]; ok {
					key += "|" + v
				}
			}

			if c := m.GetCounter(); c != nil {
				out[key] += c.GetValue()
			}
			if g := m.GetGauge(); g != nil {
				out[key] += g.GetValue()
			}
			if h := m.GetHistogram(); h != nil {
				out[key] += float64(h.GetSampleCount())
			}
		}
	}
	return out
}
