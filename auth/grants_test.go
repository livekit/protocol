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
	"encoding/json"
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
		require.Same(t, grants.Agent, clone.Agent)
		require.Same(t, grants.Inference, clone.Inference)
		require.Same(t, grants.SIP, clone.SIP)
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

		// require SIP
		require.Same(t, grants.SIP, clone.SIP)
		require.True(t, reflect.DeepEqual(grants.SIP, clone.SIP))
		// require Agent
		require.Same(t, grants.Agent, clone.Agent)
		require.True(t, reflect.DeepEqual(grants.Agent, clone.Agent))
		// require Inference
		require.Same(t, grants.Inference, clone.Inference)
		require.True(t, reflect.DeepEqual(grants.Inference, clone.Inference))
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

	t.Run("clone with SIP", func(t *testing.T) {
		sip := &SIPGrant{
			Admin: true,
		}
		grants := &ClaimGrants{
			Identity: "identity",
			Name:     "name",
			Kind:     "kind",
			SIP:      sip,
			Sha256:   "sha256",
			Metadata: "metadata",
		}
		clone := grants.Clone()
		require.NotSame(t, grants, clone)
		require.NotSame(t, grants.SIP, clone.SIP)
		require.Equal(t, grants.SIP.Admin, clone.SIP.Admin)
		require.True(t, reflect.DeepEqual(grants, clone))
		require.True(t, reflect.DeepEqual(grants.SIP, clone.SIP))
	})

	t.Run("clone with Agent", func(t *testing.T) {
		agent := &AgentGrant{
			Admin: true,
		}
		grants := &ClaimGrants{
			Identity: "identity",
			Name:     "name",
			Kind:     "kind",
			Agent:    agent,
			Sha256:   "sha256",
			Metadata: "metadata",
		}
		clone := grants.Clone()
		require.NotSame(t, grants, clone)
		require.NotSame(t, grants.Agent, clone.Agent)
		require.Equal(t, grants.Agent.Admin, clone.Agent.Admin)
		require.True(t, reflect.DeepEqual(grants, clone))
		require.True(t, reflect.DeepEqual(grants.Agent, clone.Agent))
	})

	t.Run("clone with Inference", func(t *testing.T) {
		inference := &InferenceGrant{
			Perform: true,
		}
		grants := &ClaimGrants{
			Identity:  "identity",
			Name:      "name",
			Kind:      "kind",
			Inference: inference,
			Sha256:    "sha256",
			Metadata:  "metadata",
		}
		clone := grants.Clone()
		require.NotSame(t, grants, clone)
		require.NotSame(t, grants.Inference, clone.Inference)
		require.Equal(t, grants.Inference.Perform, clone.Inference.Perform)
		require.True(t, reflect.DeepEqual(grants, clone))
		require.True(t, reflect.DeepEqual(grants.Inference, clone.Inference))
	})
}

func TestClaimGrantsVideoRTCCompat(t *testing.T) {
	t.Parallel()

	t.Run("unmarshal with video field", func(t *testing.T) {
		jsonData := `{"identity":"user1","video":{"roomJoin":true,"room":"test-room","canPublish":true}}`

		var grants ClaimGrants
		err := json.Unmarshal([]byte(jsonData), &grants)
		require.NoError(t, err)
		require.Equal(t, "user1", grants.Identity)
		require.NotNil(t, grants.Video)
		require.True(t, grants.Video.RoomJoin)
		require.Equal(t, "test-room", grants.Video.Room)
		require.True(t, grants.Video.GetCanPublish())
	})

	t.Run("unmarshal with rtc field", func(t *testing.T) {
		jsonData := `{"identity":"user2","rtc":{"roomJoin":true,"room":"rtc-room","canSubscribe":false}}`

		var grants ClaimGrants
		err := json.Unmarshal([]byte(jsonData), &grants)
		require.NoError(t, err)
		require.Equal(t, "user2", grants.Identity)
		require.NotNil(t, grants.Video)
		require.True(t, grants.Video.RoomJoin)
		require.Equal(t, "rtc-room", grants.Video.Room)
		require.False(t, grants.Video.GetCanSubscribe())
	})

	t.Run("unmarshal with both video and rtc fields prefers rtc", func(t *testing.T) {
		jsonData := `{"identity":"user3","video":{"room":"video-room"},"rtc":{"room":"rtc-room"}}`

		var grants ClaimGrants
		err := json.Unmarshal([]byte(jsonData), &grants)
		require.NoError(t, err)
		require.NotNil(t, grants.Video)
		require.Equal(t, "rtc-room", grants.Video.Room, "rtc field should take precedence over video")
	})

	t.Run("marshal outputs video field", func(t *testing.T) {
		grants := &ClaimGrants{
			Identity: "user4",
			Video: &VideoGrant{
				RoomJoin: true,
				Room:     "my-room",
			},
		}

		data, err := json.Marshal(grants)
		require.NoError(t, err)

		// Verify the output contains "video" not "rtc"
		var rawMap map[string]interface{}
		err = json.Unmarshal(data, &rawMap)
		require.NoError(t, err)
		require.Contains(t, rawMap, "video", "marshaled JSON should contain 'video' field")
		require.NotContains(t, rawMap, "rtc", "marshaled JSON should not contain 'rtc' field")
	})

	t.Run("roundtrip with video field preserves data", func(t *testing.T) {
		original := &ClaimGrants{
			Identity: "user5",
			Name:     "Test User",
			Video: &VideoGrant{
				RoomJoin:   true,
				Room:       "test-room",
				RoomCreate: true,
			},
			SIP: &SIPGrant{Admin: true},
		}

		data, err := json.Marshal(original)
		require.NoError(t, err)

		var decoded ClaimGrants
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		require.Equal(t, original.Identity, decoded.Identity)
		require.Equal(t, original.Name, decoded.Name)
		require.Equal(t, original.Video.RoomJoin, decoded.Video.RoomJoin)
		require.Equal(t, original.Video.Room, decoded.Video.Room)
		require.Equal(t, original.Video.RoomCreate, decoded.Video.RoomCreate)
		require.Equal(t, original.SIP.Admin, decoded.SIP.Admin)
	})

	t.Run("unmarshal with all grant types via rtc", func(t *testing.T) {
		jsonData := `{
			"identity": "agent1",
			"kind": "agent",
			"rtc": {"roomJoin": true, "room": "agent-room", "agent": true},
			"sip": {"admin": true},
			"agent": {"admin": true},
			"inference": {"perform": true}
		}`

		var grants ClaimGrants
		err := json.Unmarshal([]byte(jsonData), &grants)
		require.NoError(t, err)
		require.Equal(t, "agent1", grants.Identity)
		require.Equal(t, "agent", grants.Kind)
		require.NotNil(t, grants.Video)
		require.True(t, grants.Video.RoomJoin)
		require.Equal(t, "agent-room", grants.Video.Room)
		require.True(t, grants.Video.Agent)
		require.NotNil(t, grants.SIP)
		require.True(t, grants.SIP.Admin)
		require.NotNil(t, grants.Agent)
		require.True(t, grants.Agent.Admin)
		require.NotNil(t, grants.Inference)
		require.True(t, grants.Inference.Perform)
	})
}

func TestParticipantKind(t *testing.T) {
	const kindMin, kindMax = livekit.ParticipantInfo_STANDARD, livekit.ParticipantInfo_AGENT
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

func TestParticipantKindDetail(t *testing.T) {
	const detailMin, detailMax = livekit.ParticipantInfo_CLOUD_AGENT, livekit.ParticipantInfo_CONNECTOR_TWILIO
	var details []livekit.ParticipantInfo_KindDetail
	for k := detailMin; k <= detailMax; k++ {
		details = append(details, k)
	}

	require.EqualValues(t, details, kindDetailsToProto(kindDetailsFromProto(details)))
}

func TestRoomConfiguration_CheckCredentials(t *testing.T) {
	t.Parallel()

	t.Run("nil egress returns nil", func(t *testing.T) {
		config := &RoomConfiguration{}
		require.NoError(t, config.CheckCredentials())
	})

	t.Run("empty egress returns nil", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{},
		}
		require.NoError(t, config.CheckCredentials())
	})

	t.Run("participant file output with S3 secret fails", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Participant: &livekit.AutoParticipantEgress{
					FileOutputs: []*livekit.EncodedFileOutput{
						{
							Output: &livekit.EncodedFileOutput_S3{
								S3: &livekit.S3Upload{
									AccessKey: "access",
									Secret:    "secret", // This should trigger error
									Bucket:    "bucket",
								},
							},
						},
					},
				},
			},
		}
		require.ErrorIs(t, config.CheckCredentials(), ErrSensitiveCredentials)
	})

	t.Run("participant file output with S3 but no secret passes", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Participant: &livekit.AutoParticipantEgress{
					FileOutputs: []*livekit.EncodedFileOutput{
						{
							Output: &livekit.EncodedFileOutput_S3{
								S3: &livekit.S3Upload{
									AccessKey: "access",
									Secret:    "", // No secret
									Bucket:    "bucket",
									Region:    "us-west-2",
								},
							},
						},
					},
				},
			},
		}
		require.NoError(t, config.CheckCredentials())
	})

	t.Run("participant segment output with GCP credentials fails", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Participant: &livekit.AutoParticipantEgress{
					SegmentOutputs: []*livekit.SegmentedFileOutput{
						{
							Output: &livekit.SegmentedFileOutput_Gcp{
								Gcp: &livekit.GCPUpload{
									Credentials: "credentials", // This should trigger error
									Bucket:      "bucket",
								},
							},
						},
					},
				},
			},
		}
		require.ErrorIs(t, config.CheckCredentials(), ErrSensitiveCredentials)
	})

	t.Run("participant segment output with GCP but no credentials passes", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Participant: &livekit.AutoParticipantEgress{
					SegmentOutputs: []*livekit.SegmentedFileOutput{
						{
							Output: &livekit.SegmentedFileOutput_Gcp{
								Gcp: &livekit.GCPUpload{
									Credentials: "", // No credentials
									Bucket:      "bucket",
								},
							},
						},
					},
				},
			},
		}
		require.NoError(t, config.CheckCredentials())
	})

	t.Run("room file output with Azure account key fails", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Room: &livekit.RoomCompositeEgressRequest{
					FileOutputs: []*livekit.EncodedFileOutput{
						{
							Output: &livekit.EncodedFileOutput_Azure{
								Azure: &livekit.AzureBlobUpload{
									AccountName:   "account",
									AccountKey:    "key", // This should trigger error
									ContainerName: "container",
								},
							},
						},
					},
				},
			},
		}
		require.ErrorIs(t, config.CheckCredentials(), ErrSensitiveCredentials)
	})

	t.Run("room segment output with AliOSS secret fails", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Room: &livekit.RoomCompositeEgressRequest{
					SegmentOutputs: []*livekit.SegmentedFileOutput{
						{
							Output: &livekit.SegmentedFileOutput_AliOSS{
								AliOSS: &livekit.AliOSSUpload{
									AccessKey: "access",
									Secret:    "secret", // This should trigger error
									Bucket:    "bucket",
								},
							},
						},
					},
				},
			},
		}
		require.ErrorIs(t, config.CheckCredentials(), ErrSensitiveCredentials)
	})

	t.Run("room image output with valid config passes", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Room: &livekit.RoomCompositeEgressRequest{
					ImageOutputs: []*livekit.ImageOutput{
						{
							CaptureInterval: 5,
							Width:           1920,
							Height:          1080,
							Output: &livekit.ImageOutput_S3{
								S3: &livekit.S3Upload{
									AccessKey: "access",
									Secret:    "", // No secret
									Bucket:    "bucket",
								},
							},
						},
					},
				},
			},
		}
		require.NoError(t, config.CheckCredentials())
	})

	t.Run("room stream outputs always fail", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Room: &livekit.RoomCompositeEgressRequest{
					StreamOutputs: []*livekit.StreamOutput{
						{
							Protocol: livekit.StreamProtocol_RTMP,
							Urls:     []string{"rtmp://example.com/live"},
						},
					},
				},
			},
		}
		require.ErrorIs(t, config.CheckCredentials(), ErrSensitiveCredentials)
	})

	t.Run("tracks output with S3 secret fails", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Tracks: &livekit.AutoTrackEgress{
					Filepath: "output.mp4",
					Output: &livekit.AutoTrackEgress_S3{
						S3: &livekit.S3Upload{
							AccessKey: "access",
							Secret:    "secret", // This should trigger error
							Bucket:    "bucket",
						},
					},
				},
			},
		}
		require.ErrorIs(t, config.CheckCredentials(), ErrSensitiveCredentials)
	})

	t.Run("tracks output without credentials passes", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Tracks: &livekit.AutoTrackEgress{
					Filepath: "output.mp4",
					Output: &livekit.AutoTrackEgress_Gcp{
						Gcp: &livekit.GCPUpload{
							Credentials: "", // No credentials
							Bucket:      "bucket",
						},
					},
				},
			},
		}
		require.NoError(t, config.CheckCredentials())
	})

	t.Run("multiple outputs with mixed credentials", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Participant: &livekit.AutoParticipantEgress{
					FileOutputs: []*livekit.EncodedFileOutput{
						{
							Output: &livekit.EncodedFileOutput_S3{
								S3: &livekit.S3Upload{
									AccessKey: "access",
									Secret:    "", // No secret - OK
									Bucket:    "bucket1",
								},
							},
						},
						{
							Output: &livekit.EncodedFileOutput_Gcp{
								Gcp: &livekit.GCPUpload{
									Credentials: "credentials", // Has credentials - should fail
									Bucket:      "bucket2",
								},
							},
						},
					},
				},
			},
		}
		require.ErrorIs(t, config.CheckCredentials(), ErrSensitiveCredentials)
	})

	t.Run("all cloud providers without credentials pass", func(t *testing.T) {
		config := &RoomConfiguration{
			Egress: &livekit.RoomEgress{
				Room: &livekit.RoomCompositeEgressRequest{
					FileOutputs: []*livekit.EncodedFileOutput{
						{
							Output: &livekit.EncodedFileOutput_S3{
								S3: &livekit.S3Upload{
									AccessKey: "access",
									Secret:    "", // No secret
									Bucket:    "s3bucket",
								},
							},
						},
						{
							Output: &livekit.EncodedFileOutput_Gcp{
								Gcp: &livekit.GCPUpload{
									Credentials: "", // No credentials
									Bucket:      "gcpbucket",
								},
							},
						},
						{
							Output: &livekit.EncodedFileOutput_Azure{
								Azure: &livekit.AzureBlobUpload{
									AccountName:   "account",
									AccountKey:    "", // No key
									ContainerName: "container",
								},
							},
						},
						{
							Output: &livekit.EncodedFileOutput_AliOSS{
								AliOSS: &livekit.AliOSSUpload{
									AccessKey: "access",
									Secret:    "", // No secret
									Bucket:    "alibucket",
								},
							},
						},
					},
				},
			},
		}
		require.NoError(t, config.CheckCredentials())
	})
}
