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

package signalling

import (
	"testing"

	"github.com/pion/webrtc/v4"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
)

func TestSessionDescriptionRoundTrip(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		pb := ToProtoSessionDescription(webrtc.SessionDescription{}, 0, nil)
		require.Nil(t, pb)

		sd, id, midToTrackID := FromProtoSessionDescription(pb)
		require.Equal(t, webrtc.SessionDescription{}, sd)
		require.Zero(t, id)
		require.Nil(t, midToTrackID)
	})

	t.Run("offer", func(t *testing.T) {
		in := webrtc.SessionDescription{Type: webrtc.SDPTypeOffer, SDP: "v=0\r\n"}
		mids := map[string]string{"0": "TR_1"}

		sd, id, midToTrackID := FromProtoSessionDescription(ToProtoSessionDescription(in, 7, mids))
		require.Equal(t, in, sd)
		require.Equal(t, uint32(7), id)
		require.Equal(t, mids, midToTrackID)
	})
}

func TestFromProtoSessionDescriptionEmptySyncState(t *testing.T) {
	// field 12 (sync_state) with length 0: a SyncState with no answer and no offer
	req := &livekit.SignalRequest{}
	require.NoError(t, proto.Unmarshal([]byte{0x62, 0x00}, req))
	state := req.GetSyncState()
	require.NotNil(t, state)

	for _, pb := range []*livekit.SessionDescription{state.GetAnswer(), state.GetOffer()} {
		require.Nil(t, pb)
		sd, id, midToTrackID := FromProtoSessionDescription(pb)
		require.Equal(t, webrtc.SessionDescription{}, sd)
		require.Zero(t, id)
		require.Nil(t, midToTrackID)
	}
}
