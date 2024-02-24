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
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/middleware"
)

const (
	livekitNamespace = "livekit"
)

var (
	psrpcRequestTime        *prometheus.HistogramVec
	psrpcStreamSendTime     *prometheus.HistogramVec
	psrpcStreamReceiveTotal *prometheus.CounterVec
	psrpcStreamCurrent      *prometheus.GaugeVec
	psrpcErrorTotal         *prometheus.CounterVec
)

func InitPSRPCStats(constLabels prometheus.Labels) {
	labels := []string{"role", "kind", "service", "method"}
	streamLabels := []string{"role", "service", "method"}

	psrpcRequestTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "request_time_ms",
		ConstLabels: constLabels,
		Buckets:     []float64{10, 50, 100, 300, 500, 1000, 1500, 2000, 5000, 10000},
	}, labels)
	psrpcStreamSendTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "stream_send_time_ms",
		ConstLabels: constLabels,
		Buckets:     []float64{10, 50, 100, 300, 500, 1000, 1500, 2000, 5000, 10000},
	}, streamLabels)
	psrpcStreamReceiveTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "stream_receive_total",
		ConstLabels: constLabels,
	}, streamLabels)
	psrpcStreamCurrent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "stream_count",
		ConstLabels: constLabels,
	}, streamLabels)
	psrpcErrorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   livekitNamespace,
		Subsystem:   "psrpc",
		Name:        "error_total",
		ConstLabels: constLabels,
	}, labels)

	prometheus.MustRegister(psrpcRequestTime)
	prometheus.MustRegister(psrpcStreamSendTime)
	prometheus.MustRegister(psrpcStreamReceiveTotal)
	prometheus.MustRegister(psrpcStreamCurrent)
	prometheus.MustRegister(psrpcErrorTotal)
}

var _ middleware.MetricsObserver = PSRPCMetricsObserver{}

type PSRPCMetricsObserver struct{}

func (o PSRPCMetricsObserver) OnUnaryRequest(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, err error) {
	if err != nil {
		psrpcErrorTotal.WithLabelValues(role.String(), "rpc", info.Service, info.Method).Inc()
	} else if role == middleware.ClientRole {
		psrpcRequestTime.WithLabelValues(role.String(), "rpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	} else {
		psrpcRequestTime.WithLabelValues(role.String(), "rpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnMultiRequest(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, responseCount int, errorCount int) {
	if responseCount == 0 {
		psrpcErrorTotal.WithLabelValues(role.String(), "multirpc", info.Service, info.Method).Inc()
	} else if role == middleware.ClientRole {
		psrpcRequestTime.WithLabelValues(role.String(), "multirpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	} else {
		psrpcRequestTime.WithLabelValues(role.String(), "multirpc", info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnStreamSend(role middleware.MetricRole, info psrpc.RPCInfo, duration time.Duration, err error) {
	if err != nil {
		psrpcErrorTotal.WithLabelValues(role.String(), "stream", info.Service, info.Method).Inc()
	} else {
		psrpcStreamSendTime.WithLabelValues(role.String(), info.Service, info.Method).Observe(float64(duration.Milliseconds()))
	}
}

func (o PSRPCMetricsObserver) OnStreamRecv(role middleware.MetricRole, info psrpc.RPCInfo, err error) {
	if err != nil {
		psrpcErrorTotal.WithLabelValues(role.String(), "stream", info.Service, info.Method).Inc()
	} else {
		psrpcStreamReceiveTotal.WithLabelValues(role.String(), info.Service, info.Method).Inc()
	}
}

func (o PSRPCMetricsObserver) OnStreamOpen(role middleware.MetricRole, info psrpc.RPCInfo) {
	psrpcStreamCurrent.WithLabelValues(role.String(), info.Service, info.Method).Inc()
}

func (o PSRPCMetricsObserver) OnStreamClose(role middleware.MetricRole, info psrpc.RPCInfo) {
	psrpcStreamCurrent.WithLabelValues(role.String(), info.Service, info.Method).Dec()
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
