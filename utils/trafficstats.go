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
	"time"

	"github.com/livekit/protocol/livekit"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RTPStatsDiffToTrafficStats(before, after *livekit.RTPStats) *livekit.TrafficStats {
	if after == nil {
		return nil
	}

	startTime := after.StartTime
	if before != nil {
		startTime = before.EndTime
	}

	if before == nil {
		return &livekit.TrafficStats{
			StartTime: startTime,
			EndTime:   after.EndTime,
			Packets:   after.Packets,
			Bytes:     after.Bytes + after.BytesDuplicate + after.BytesPadding,
		}
	}

	return &livekit.TrafficStats{
		StartTime: startTime,
		EndTime:   after.EndTime,
		Packets:   after.Packets - before.Packets,
		Bytes:     (after.Bytes + after.BytesDuplicate + after.BytesPadding) - (before.Bytes + before.BytesDuplicate + before.BytesPadding),
	}
}

func AggregateTrafficStats(statsList []*livekit.TrafficStats) *livekit.TrafficStats {
	if len(statsList) == 0 {
		return nil
	}

	startTime := time.Time{}
	endTime := time.Time{}

	packets := uint32(0)
	bytes := uint64(0)

	for _, stats := range statsList {
		if startTime.IsZero() || startTime.After(stats.StartTime.AsTime()) {
			startTime = stats.StartTime.AsTime()
		}

		if endTime.IsZero() || endTime.Before(stats.EndTime.AsTime()) {
			endTime = stats.EndTime.AsTime()
		}

		packets += stats.Packets
		bytes += stats.Bytes
	}

	if endTime.IsZero() {
		endTime = time.Now()
	}
	return &livekit.TrafficStats{
		StartTime: timestamppb.New(startTime),
		EndTime:   timestamppb.New(endTime),
		Packets:   packets,
		Bytes:     bytes,
	}
}
