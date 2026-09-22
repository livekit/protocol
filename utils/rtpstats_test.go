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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/livekit/protocol/livekit"
)

func TestAggregateRTPStatsFEC(t *testing.T) {
	start := time.Now()
	end := start.Add(time.Second)
	stats := AggregateRTPStats([]*livekit.RTPStats{
		{
			StartTime:           timestamppb.New(start),
			EndTime:             timestamppb.New(end),
			FecPacketsReceived:  3,
			FecBytesReceived:    900,
			FecPacketsDiscarded: 1,
			PacketsRecovered:    2,
		},
		{
			StartTime:           timestamppb.New(start),
			EndTime:             timestamppb.New(end),
			FecPacketsReceived:  4,
			FecBytesReceived:    1200,
			FecPacketsDiscarded: 2,
			PacketsRecovered:    3,
		},
	}, 8)

	require.EqualValues(t, 7, stats.FecPacketsReceived)
	require.EqualValues(t, 2100, stats.FecBytesReceived)
	require.EqualValues(t, 3, stats.FecPacketsDiscarded)
	require.EqualValues(t, 5, stats.PacketsRecovered)
}
