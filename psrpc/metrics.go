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

package psrpc

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
