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
	"maps"
	"slices"
	sync "sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/atomic"

	"github.com/livekit/protocol/logger"
	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/middleware"
)

const (
	livekitNamespace = "livekit"
)

type psrpcMetrics struct {
	requestTime        prometheus.ObserverVec
	streamSendTime     prometheus.ObserverVec
	streamReceiveTotal *prometheus.CounterVec
	streamCurrent      *prometheus.GaugeVec
	errorTotal         *prometheus.CounterVec
	bytesTotal         *prometheus.CounterVec
	requestsReceived   *prometheus.CounterVec
	requestsExpired    *prometheus.CounterVec
	claimTotal         *prometheus.CounterVec
	claimWaitTime      prometheus.ObserverVec
}

var (
	metricsBase struct {
		mu          sync.RWMutex
		initialized bool
		curryLabels prometheus.Labels
		psrpcMetrics
	}
	metrics atomic.Pointer[psrpcMetrics]
)

type psrpcMetricsOptions struct {
	curryLabels prometheus.Labels
}

type PSRPCMetricsOption func(*psrpcMetricsOptions)

func WithCurryLabels(labels prometheus.Labels) PSRPCMetricsOption {
	return func(o *psrpcMetricsOptions) {
		maps.Copy(o.curryLabels, labels)
	}
}

func InitPSRPCStats(constLabels prometheus.Labels, opts ...PSRPCMetricsOption) {
	metricsBase.mu.Lock()
	if metricsBase.initialized {
		metricsBase.mu.Unlock()
		return
	}
	metricsBase.initialized = true

	o := psrpcMetricsOptions{
		curryLabels: prometheus.Labels{},
	}
	for _, opt := range opts {
		opt(&o)
	}

	metricsBase.curryLabels = o.curryLabels
	curryLabelNames := slices.Collect(maps.Keys(o.curryLabels))
	slices.Sort(curryLabelNames)

	labels := slices.Concat(curryLabelNames, []string{"role", "kind", "service", "method"})
	streamLabels := slices.Concat(curryLabelNames, []string{"role", "service", "method"})
	errorLabels := slices.Concat(labels, []string{"error_code"})
	bytesLabels := slices.Concat(labels, []string{"direction"})
	// Lifecycle metrics are server-side only, so they carry no role label.
	lifecycleLabels := slices.Concat(curryLabelNames, []string{"service", "method"})
	claimLabels := slices.Concat(lifecycleLabels, []string{"outcome"})

	metricsBase.requestTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "request_time_ms",
		ConstLabels: constLabels,
		Buckets:     []float64{10, 50, 100, 300, 500, 1000, 1500, 2000, 5000, 10000},
	}, labels)
	metricsBase.streamSendTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "stream_send_time_ms",
		ConstLabels: constLabels,
		Buckets:     []float64{10, 50, 100, 300, 500, 1000, 1500, 2000, 5000, 10000},
	}, streamLabels)
	metricsBase.streamReceiveTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "stream_receive_total",
		ConstLabels: constLabels,
	}, streamLabels)
	metricsBase.streamCurrent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "stream_count",
		ConstLabels: constLabels,
	}, streamLabels)
	metricsBase.errorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "error_total",
		ConstLabels: constLabels,
	}, errorLabels)
	metricsBase.bytesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "bytes_total",
		ConstLabels: constLabels,
	}, bytesLabels)

	metricsBase.requestsReceived = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "requests_received_total",
		ConstLabels: constLabels,
	}, lifecycleLabels)
	metricsBase.requestsExpired = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "requests_expired_total",
		ConstLabels: constLabels,
	}, lifecycleLabels)
	metricsBase.claimTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "claim_total",
		ConstLabels: constLabels,
	}, claimLabels)
	metricsBase.claimWaitTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "claim_wait_time_ms",
		ConstLabels: constLabels,
		// A granted claim settles in single-digit ms; a timed-out one runs to
		// the caller's selection timeout, 1s by default.
		Buckets: []float64{1, 5, 10, 25, 50, 100, 250, 500, 1000, 3000},
	}, claimLabels)

	metricsBase.mu.Unlock()

	prometheus.MustRegister(metricsBase.requestTime)
	prometheus.MustRegister(metricsBase.streamSendTime)
	prometheus.MustRegister(metricsBase.streamReceiveTotal)
	prometheus.MustRegister(metricsBase.streamCurrent)
	prometheus.MustRegister(metricsBase.errorTotal)
	prometheus.MustRegister(metricsBase.bytesTotal)
	prometheus.MustRegister(metricsBase.requestsReceived)
	prometheus.MustRegister(metricsBase.requestsExpired)
	prometheus.MustRegister(metricsBase.claimTotal)
	prometheus.MustRegister(metricsBase.claimWaitTime)

	CurryMetricLabels(o.curryLabels)
}

func CurryMetricLabels(labels prometheus.Labels) {
	metricsBase.mu.Lock()
	defer metricsBase.mu.Unlock()
	if !metricsBase.initialized {
		return
	}

	for k := range metricsBase.curryLabels {
		if v, ok := labels[k]; ok {
			metricsBase.curryLabels[k] = v
		}
	}

	metrics.Store(&psrpcMetrics{
		requestTime:        metricsBase.requestTime.MustCurryWith(metricsBase.curryLabels),
		streamSendTime:     metricsBase.streamSendTime.MustCurryWith(metricsBase.curryLabels),
		streamReceiveTotal: metricsBase.streamReceiveTotal.MustCurryWith(metricsBase.curryLabels),
		streamCurrent:      metricsBase.streamCurrent.MustCurryWith(metricsBase.curryLabels),
		errorTotal:         metricsBase.errorTotal.MustCurryWith(metricsBase.curryLabels),
		bytesTotal:         metricsBase.bytesTotal.MustCurryWith(metricsBase.curryLabels),
		requestsReceived:   metricsBase.requestsReceived.MustCurryWith(metricsBase.curryLabels),
		requestsExpired:    metricsBase.requestsExpired.MustCurryWith(metricsBase.curryLabels),
		claimTotal:         metricsBase.claimTotal.MustCurryWith(metricsBase.curryLabels),
		claimWaitTime:      metricsBase.claimWaitTime.MustCurryWith(metricsBase.curryLabels),
	})
}

func errorCodeLabel(err error) string {
	if code, ok := psrpc.GetErrorCode(err); ok && code != psrpc.OK {
		return string(code)
	}
	return string(psrpc.Unknown)
}

var (
	_ middleware.MetricsObserver = PSRPCMetricsObserver{}
	_ psrpc.RequestObserver      = PSRPCMetricsObserver{}
)

type PSRPCMetricsObserver struct{}

func (o PSRPCMetricsObserver) OnUnaryRequest(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, err error, rxBytes, txBytes int) {
	m := metrics.Load()
	m.bytesTotal.WithLabelValues(role.String(), "rpc", info.Service, info.Method, "rx").Add(float64(rxBytes))
	m.bytesTotal.WithLabelValues(role.String(), "rpc", info.Service, info.Method, "tx").Add(float64(txBytes))

	if err != nil {
		m.errorTotal.WithLabelValues(role.String(), "rpc", info.Service, info.Method, errorCodeLabel(err)).Inc()
	} else {
		m.requestTime.WithLabelValues(role.String(), "rpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnMultiRequest(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, responseCount, errorCount, rxBytes, txBytes int) {
	m := metrics.Load()
	m.bytesTotal.WithLabelValues(role.String(), "multirpc", info.Service, info.Method, "rx").Add(float64(rxBytes))
	m.bytesTotal.WithLabelValues(role.String(), "multirpc", info.Service, info.Method, "tx").Add(float64(txBytes))

	if responseCount == 0 {
		// psrpc's MetricsObserver doesn't surface an error for multi requests
		m.errorTotal.WithLabelValues(role.String(), "multirpc", info.Service, info.Method, string(psrpc.Unknown)).Inc()
	} else {
		m.requestTime.WithLabelValues(role.String(), "multirpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnStreamSend(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, err error, bytes int) {
	m := metrics.Load()
	m.bytesTotal.WithLabelValues(role.String(), "stream", info.Service, info.Method, "tx").Add(float64(bytes))

	if err != nil {
		m.errorTotal.WithLabelValues(role.String(), "stream", info.Service, info.Method, errorCodeLabel(err)).Inc()
	} else {
		m.streamSendTime.WithLabelValues(role.String(), info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnStreamRecv(role middleware.MetricRole, info psrpc.RPCInfo, err error, bytes int) {
	m := metrics.Load()
	m.bytesTotal.WithLabelValues(role.String(), "stream", info.Service, info.Method, "rx").Add(float64(bytes))

	if err != nil {
		m.errorTotal.WithLabelValues(role.String(), "stream", info.Service, info.Method, errorCodeLabel(err)).Inc()
	} else {
		m.streamReceiveTotal.WithLabelValues(role.String(), info.Service, info.Method).Inc()
	}
}

func (o PSRPCMetricsObserver) OnStreamOpen(role middleware.MetricRole, info psrpc.RPCInfo) {
	m := metrics.Load()
	m.streamCurrent.WithLabelValues(role.String(), info.Service, info.Method).Inc()
}

func (o PSRPCMetricsObserver) OnStreamClose(role middleware.MetricRole, info psrpc.RPCInfo) {
	m := metrics.Load()
	m.streamCurrent.WithLabelValues(role.String(), info.Service, info.Method).Dec()
}

var _ middleware.MetricsObserver = UnimplementedMetricsObserver{}

type UnimplementedMetricsObserver struct{}

func (o UnimplementedMetricsObserver) OnUnaryRequest(role middleware.MetricRole, rpcInfo psrpc.RPCInfo, duration time.Duration, err error, rxBytes, txBytes int) {
}
func (o UnimplementedMetricsObserver) OnMultiRequest(role middleware.MetricRole, rpcInfo psrpc.RPCInfo, duration time.Duration, responseCount, errorCount, reqBytes, txBytes int) {
}
func (o UnimplementedMetricsObserver) OnStreamSend(role middleware.MetricRole, rpcInfo psrpc.RPCInfo, duration time.Duration, err error, bytes int) {
}
func (o UnimplementedMetricsObserver) OnStreamRecv(role middleware.MetricRole, rpcInfo psrpc.RPCInfo, err error, bytes int) {
}
func (o UnimplementedMetricsObserver) OnStreamOpen(role middleware.MetricRole, rpcInfo psrpc.RPCInfo) {
}
func (o UnimplementedMetricsObserver) OnStreamClose(role middleware.MetricRole, rpcInfo psrpc.RPCInfo) {
}

// OnRequestReceived, OnRequestExpired and OnClaim report server-side lifecycle
// events that the interceptor chain cannot see, because in each case the
// handler is never invoked. Installed by psrpc.WithServerObserver, which is
// separate from middleware.WithServerMetrics.

func (o PSRPCMetricsObserver) OnRequestReceived(info psrpc.RPCInfo) {
	metrics.Load().requestsReceived.WithLabelValues(info.Service, info.Method).Inc()
}

func (o PSRPCMetricsObserver) OnRequestExpired(info psrpc.RPCInfo, lateBy time.Duration) {
	metrics.Load().requestsExpired.WithLabelValues(info.Service, info.Method).Inc()
	logger.Warnw("psrpc request dropped: expired before dispatch", nil,
		"service", info.Service, "method", info.Method, "lateBy", lateBy)
}

func (o PSRPCMetricsObserver) OnClaim(info psrpc.RPCInfo, outcome psrpc.ClaimOutcome, wait time.Duration) {
	m := metrics.Load()
	m.claimTotal.WithLabelValues(info.Service, info.Method, outcome.String()).Inc()
	m.claimWaitTime.WithLabelValues(info.Service, info.Method, outcome.String()).Observe(float64(wait.Milliseconds()))

	if outcome == psrpc.ClaimTimedOut {
		// The caller stopped waiting for a bid before ours was accepted. It has
		// already returned ErrNoResponse upstream, so without this line the
		// request leaves no record on either side.
		logger.Warnw("psrpc claim timed out before the caller granted it", nil,
			"service", info.Service, "method", info.Method, "waited", wait)
	}
}
