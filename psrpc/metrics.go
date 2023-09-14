package middleware

import (
	"time"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/middleware"
)

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
