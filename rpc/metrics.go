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
	"sort"
	sync "sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/atomic"
	"golang.org/x/exp/maps"

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
	curryLabelNames := maps.Keys(o.curryLabels)
	sort.Strings(curryLabelNames)

	labels := append(curryLabelNames, "role", "kind", "service", "method")
	streamLabels := append(curryLabelNames, "role", "service", "method")

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
	}, labels)

	metricsBase.mu.Unlock()

	prometheus.MustRegister(metricsBase.requestTime)
	prometheus.MustRegister(metricsBase.streamSendTime)
	prometheus.MustRegister(metricsBase.streamReceiveTotal)
	prometheus.MustRegister(metricsBase.streamCurrent)
	prometheus.MustRegister(metricsBase.errorTotal)

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
	})
}

var _ middleware.MetricsObserver = PSRPCMetricsObserver{}

type PSRPCMetricsObserver struct{}

func (o PSRPCMetricsObserver) OnUnaryRequest(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, err error) {
	if err != nil {
		metrics.Load().errorTotal.WithLabelValues(role.String(), "rpc", info.Service, info.Method).Inc()
	} else if role == middleware.ClientRole {
		metrics.Load().requestTime.WithLabelValues(role.String(), "rpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	} else {
		metrics.Load().requestTime.WithLabelValues(role.String(), "rpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnMultiRequest(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, responseCount int, errorCount int) {
	if responseCount == 0 {
		metrics.Load().errorTotal.WithLabelValues(role.String(), "multirpc", info.Service, info.Method).Inc()
	} else if role == middleware.ClientRole {
		metrics.Load().requestTime.WithLabelValues(role.String(), "multirpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	} else {
		metrics.Load().requestTime.WithLabelValues(role.String(), "multirpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnStreamSend(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, err error) {
	if err != nil {
		metrics.Load().errorTotal.WithLabelValues(role.String(), "stream", info.Service, info.Method).Inc()
	} else {
		metrics.Load().streamSendTime.WithLabelValues(role.String(), info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnStreamRecv(role middleware.MetricRole, info psrpc.RPCInfo, err error) {
	if err != nil {
		metrics.Load().errorTotal.WithLabelValues(role.String(), "stream", info.Service, info.Method).Inc()
	} else {
		metrics.Load().streamReceiveTotal.WithLabelValues(role.String(), info.Service, info.Method).Inc()
	}
}

func (o PSRPCMetricsObserver) OnStreamOpen(role middleware.MetricRole, info psrpc.RPCInfo) {
	metrics.Load().streamCurrent.WithLabelValues(role.String(), info.Service, info.Method).Inc()
}

func (o PSRPCMetricsObserver) OnStreamClose(role middleware.MetricRole, info psrpc.RPCInfo) {
	metrics.Load().streamCurrent.WithLabelValues(role.String(), info.Service, info.Method).Dec()
}

type UnimplementedMetricsObserver struct{}

func (o UnimplementedMetricsObserver) OnUnaryRequest(role middleware.MetricRole, rpcInfo psrpc.RPCInfo, duration time.Duration, err error) {
}
func (o UnimplementedMetricsObserver) OnMultiRequest(role middleware.MetricRole, rpcInfo psrpc.RPCInfo, duration time.Duration, responseCount int, errorCount int) {
}
func (o UnimplementedMetricsObserver) OnStreamSend(role middleware.MetricRole, rpcInfo psrpc.RPCInfo, duration time.Duration, err error) {
}
func (o UnimplementedMetricsObserver) OnStreamRecv(role middleware.MetricRole, rpcInfo psrpc.RPCInfo, err error) {
}
func (o UnimplementedMetricsObserver) OnStreamOpen(role middleware.MetricRole, rpcInfo psrpc.RPCInfo) {
}
func (o UnimplementedMetricsObserver) OnStreamClose(role middleware.MetricRole, rpcInfo psrpc.RPCInfo) {
}
