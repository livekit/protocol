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

package logger_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

func marshalFields(t *testing.T, m zapcore.ObjectMarshaler) map[string]any {
	t.Helper()
	enc := zapcore.NewMapObjectEncoder()
	require.NoError(t, m.MarshalLogObject(enc))
	return enc.Fields
}

func TestProtoRedactsPIIAndSecret(t *testing.T) {
	// S3Upload has SECRET access_key/secret and PII assume_role_arn.
	msg := &livekit.S3Upload{
		AccessKey:            "AKIAEXAMPLE",
		Secret:               "supersecretvalue",
		SessionToken:         "tokenvalue",
		AssumeRoleArn:        "arn:aws:iam::123456789012:role/MyRole",
		AssumeRoleExternalId: "external-id-1",
		Region:               "us-east-1",
		Bucket:               "my-bucket",
	}
	fields := marshalFields(t, logger.Proto(msg))

	// Unannotated fields appear verbatim.
	require.Equal(t, "us-east-1", fields["region"])
	require.Equal(t, "my-bucket", fields["bucket"])

	// PII and SECRET are both redacted by Proto.
	require.Equal(t, "<redacted>", fields["accessKey"])
	require.Equal(t, "<redacted>", fields["secret"])
	require.Equal(t, "<redacted>", fields["sessionToken"])
	require.Equal(t, "<redacted>", fields["assumeRoleArn"])
	require.Equal(t, "<redacted>", fields["assumeRoleExternalID"])
}

func TestUnredactedProtoShowsPIIRedactsSecret(t *testing.T) {
	msg := &livekit.S3Upload{
		AccessKey:            "AKIAEXAMPLE",
		Secret:               "supersecretvalue",
		SessionToken:         "tokenvalue",
		AssumeRoleArn:        "arn:aws:iam::123456789012:role/MyRole",
		AssumeRoleExternalId: "external-id-1",
		Region:               "us-east-1",
	}
	fields := marshalFields(t, logger.UnredactedProto(msg))

	require.Equal(t, "us-east-1", fields["region"])

	// PII is exposed.
	require.Equal(t, "arn:aws:iam::123456789012:role/MyRole", fields["assumeRoleArn"])

	// SECRETs remain redacted.
	require.Equal(t, "<redacted>", fields["accessKey"])
	require.Equal(t, "<redacted>", fields["secret"])
	require.Equal(t, "<redacted>", fields["sessionToken"])
	require.Equal(t, "<redacted>", fields["assumeRoleExternalID"])
}

func TestProtoRedactFormatPreservedAtPIITier(t *testing.T) {
	// ParticipantInfo.metadata is PII with a size-showing redact_format.
	msg := &livekit.ParticipantInfo{
		Identity: "user-123",
		Name:     "Alice",
		Metadata: `{"plan":"pro"}`,
	}

	got := marshalFields(t, logger.Proto(msg))
	require.Equal(t, "user-123", got["identity"])
	require.Equal(t, "<redacted>", got["name"])
	require.Contains(t, got["metadata"].(string), "<redacted (")
	require.Contains(t, got["metadata"].(string), "bytes)>")

	gotUnredacted := marshalFields(t, logger.UnredactedProto(msg))
	require.Equal(t, "Alice", gotUnredacted["name"])
	require.Equal(t, `{"plan":"pro"}`, gotUnredacted["metadata"])
}

func TestProtoNestedListPIIRedacted(t *testing.T) {
	// TrackInfo.name is PII; it lives inside a repeated list on ParticipantInfo.
	msg := &livekit.ParticipantInfo{
		Identity: "id",
		Tracks: []*livekit.TrackInfo{
			{Sid: "TR1", Name: "Microphone"},
			{Sid: "TR2", Name: "Camera"},
		},
	}

	gotRedacted := marshalFields(t, logger.Proto(msg))
	tracks := gotRedacted["tracks"].([]any)
	require.Len(t, tracks, 2)
	for _, raw := range tracks {
		track := raw.(map[string]any)
		require.NotEmpty(t, track["sid"])
		require.Equal(t, "<redacted>", track["name"])
	}

	gotUnredacted := marshalFields(t, logger.UnredactedProto(msg))
	tracksU := gotUnredacted["tracks"].([]any)
	names := make([]string, 0, 2)
	for _, raw := range tracksU {
		names = append(names, raw.(map[string]any)["name"].(string))
	}
	require.ElementsMatch(t, []string{"Microphone", "Camera"}, names)
}

func TestProtoNestedSecretAlwaysRedacted(t *testing.T) {
	// ICEServer.credential is SECRET inside an ICEServer list.
	servers := []*livekit.ICEServer{
		{Urls: []string{"turn:turn.example.com"}, Username: "user", Credential: "shhh"},
	}
	wrap := &livekit.JoinResponse{IceServers: servers}

	for _, name := range []string{"Proto", "UnredactedProto"} {
		var m zapcore.ObjectMarshaler
		if name == "Proto" {
			m = logger.Proto(wrap)
		} else {
			m = logger.UnredactedProto(wrap)
		}
		fields := marshalFields(t, m)
		iceServers := fields["iceServers"].([]any)
		require.Len(t, iceServers, 1)
		entry := iceServers[0].(map[string]any)
		require.Equal(t, "<redacted>", entry["credential"], "credential must be redacted by %s", name)
	}
}

func TestProtoNilSafe(t *testing.T) {
	require.Nil(t, logger.Proto(nil))
	require.Nil(t, logger.UnredactedProto(nil))

	// Typed-nil pointer is not interface-nil; the returned marshaller is
	// non-nil but emits no fields.
	require.NotPanics(t, func() {
		marshalFields(t, logger.Proto((*livekit.S3Upload)(nil)))
	})
}

func TestProtoUnknownFieldPrefix(t *testing.T) {
	// Sanity: keys come from JSONName (proto JSON convention), not from the
	// snake_case proto field name. Use this when reading the assertions above.
	msg := &livekit.ParticipantInfo{Identity: "x"}
	fields := marshalFields(t, logger.Proto(msg))
	for k := range fields {
		require.False(t, strings.Contains(k, "_"), "key %q should be camelCase", k)
	}
}
