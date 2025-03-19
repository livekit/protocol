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

package auth

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestGrants(t *testing.T) {
	t.Parallel()

	t.Run("clone default grant", func(t *testing.T) {
		grants := &ClaimGrants{}
		clone := grants.Clone()
		require.NotSame(t, grants, clone)
		require.Same(t, grants.Video, clone.Video)
		require.True(t, reflect.DeepEqual(grants, clone))
		require.True(t, reflect.DeepEqual(grants.Video, clone.Video))
	})

	t.Run("clone nil video", func(t *testing.T) {
		grants := &ClaimGrants{
			Identity: "identity",
			Name:     "name",
			Sha256:   "sha256",
			Metadata: "metadata",
		}
		clone := grants.Clone()
		require.NotSame(t, grants, clone)
		require.Same(t, grants.Video, clone.Video)
		require.True(t, reflect.DeepEqual(grants, clone))
		require.True(t, reflect.DeepEqual(grants.Video, clone.Video))
	})

	t.Run("clone with video", func(t *testing.T) {
		tr := true
		fa := false
		video := &VideoGrant{
			RoomCreate:          true,
			RoomList:            false,
			RoomRecord:          true,
			RoomAdmin:           false,
			RoomJoin:            true,
			Room:                "room",
			CanPublish:          &tr,
			CanSubscribe:        &fa,
			CanPublishData:      nil,
			Hidden:              true,
			Recorder:            false,
			CanSubscribeMetrics: &tr,
		}
		grants := &ClaimGrants{
			Identity: "identity",
			Name:     "name",
			Kind:     "kind",
			Video:    video,
			Sha256:   "sha256",
			Metadata: "metadata",
		}
		clone := grants.Clone()
		require.NotSame(t, grants, clone)
		require.NotSame(t, grants.Video, clone.Video)
		require.NotSame(t, grants.Video.CanPublish, clone.Video.CanPublish)
		require.NotSame(t, grants.Video.CanSubscribe, clone.Video.CanSubscribe)
		require.Same(t, grants.Video.CanPublishData, clone.Video.CanPublishData)
		require.True(t, reflect.DeepEqual(grants, clone))
		require.True(t, reflect.DeepEqual(grants.Video, clone.Video))
	})
}

func TestParticipantKind(t *testing.T) {
	const kindMin, kindMax = livekit.ParticipantInfo_STANDARD, livekit.ParticipantInfo_CLOUD_AGENT
	for k := kindMin; k <= kindMax; k++ {
		k := k
		t.Run(k.String(), func(t *testing.T) {
			require.Equal(t, k, kindToProto(kindFromProto(k)))
		})
	}
	const kindNext = kindMax + 1
	if _, err := strconv.Atoi(kindNext.String()); err != nil {
		t.Errorf("Please update kindMax to match protobuf. Missing value: %s", kindNext)
	}
}
