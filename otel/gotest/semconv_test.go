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

// Package gotest verifies the Weaver-generated Go semantic conventions in
// github.com/livekit/protocol/otel/gen/go. It lives outside otel/gen so otel/generate.sh, which
// replaces otel/gen whole, never deletes it. It encodes the invariants a consumer (such as the
// collector) relies on, so a model or template change that breaks them fails `go test` rather than
// surfacing downstream.
package gotest

import (
	"testing"

	semconv "github.com/livekit/protocol/otel/gen/go"
)

// TestConstants spot-checks that representative constants carry the exact OTLP wire strings, so a
// template rename of a Go identifier cannot silently change the value a producer sends.
func TestConstants(t *testing.T) {
	cases := map[string]string{
		"AttrLkRoomID":            semconv.AttrLkRoomId,
		"AttrLkParticipantID":     semconv.AttrLkParticipantId,
		"AttrLkTrackID":           semconv.AttrLkTrackId,
		"AttrLkTrackSource":       semconv.AttrLkTrackSource,
		"AttrLkDisconnectReason":  semconv.AttrLkDisconnectReason,
		"TrackSourceCamera":       semconv.LkTrackSourceSourceCamera,
		"SpanLkConnect":           semconv.SpanLkConnect,
		"EventLkRoomDisconnected": semconv.EventLkRoomDisconnected,
	}
	want := map[string]string{
		"AttrLkRoomID":            "lk.room_id",
		"AttrLkParticipantID":     "lk.participant_id",
		"AttrLkTrackID":           "lk.track_id",
		"AttrLkTrackSource":       "lk.track_source",
		"AttrLkDisconnectReason":  "lk.disconnect_reason",
		"TrackSourceCamera":       "SOURCE_CAMERA",
		"SpanLkConnect":           "lk.connect",
		"EventLkRoomDisconnected": "lk.room.disconnected",
	}
	for name, got := range cases {
		if got != want[name] {
			t.Errorf("%s = %q, want %q", name, got, want[name])
		}
	}
}

// TestRenamedKeys checks the legacy-key alias map against the eight documented client renames, and
// its structural invariants: a legacy key is never empty, maps to a non-empty canonical key, and is
// never equal to its own canonical key (a rename changes the spelling).
func TestRenamedKeys(t *testing.T) {
	want := map[string]string{
		"lk.room.sid":                    semconv.AttrLkRoomId,
		"lk.participant.sid":             semconv.AttrLkParticipantId,
		"lk.track.sid":                   semconv.AttrLkTrackId,
		"lk.room.name":                   semconv.AttrLkPiiRoomName,
		"lk.participant.identity":        semconv.AttrLkPiiParticipantIdentity,
		"lk.participant.remote_identity": semconv.AttrLkPiiRemoteParticipantIdentity,
		"lk.track.source":                semconv.AttrLkTrackSource,
		"lk.disconnect.reason":           semconv.AttrLkDisconnectReason,
	}
	if len(semconv.RenamedKeys) != len(want) {
		t.Fatalf("RenamedKeys has %d entries, want %d: %v", len(semconv.RenamedKeys), len(want), semconv.RenamedKeys)
	}
	for legacy, canonical := range want {
		got, ok := semconv.RenamedKeys[legacy]
		if !ok {
			t.Errorf("RenamedKeys is missing legacy key %q", legacy)
			continue
		}
		if got != canonical {
			t.Errorf("RenamedKeys[%q] = %q, want %q", legacy, got, canonical)
		}
	}
	for legacy, canonical := range semconv.RenamedKeys {
		if legacy == "" || canonical == "" {
			t.Errorf("RenamedKeys has an empty mapping %q -> %q", legacy, canonical)
		}
		if legacy == canonical {
			t.Errorf("RenamedKeys[%q] maps to itself", legacy)
		}
	}
}

// TestRenamedValues checks the two enums whose values were also renamed, and that every key in
// RenamedValues is also in RenamedKeys (a value remapping implies a key remapping).
func TestRenamedValues(t *testing.T) {
	if got := semconv.RenamedValues["lk.track.source"]["camera"]; got != semconv.LkTrackSourceSourceCamera {
		t.Errorf("RenamedValues track.source camera = %q, want %q", got, semconv.LkTrackSourceSourceCamera)
	}
	if got := semconv.RenamedValues["lk.track.source"]["screen_share"]; got != semconv.LkTrackSourceSourceScreenshare {
		t.Errorf("RenamedValues track.source screen_share = %q, want %q", got, semconv.LkTrackSourceSourceScreenshare)
	}
	if got := semconv.RenamedValues["lk.disconnect.reason"]["client_initiated"]; got != semconv.LkDisconnectReasonClientInitiated {
		t.Errorf("RenamedValues disconnect.reason client_initiated = %q, want %q", got, semconv.LkDisconnectReasonClientInitiated)
	}
	for legacy := range semconv.RenamedValues {
		if _, ok := semconv.RenamedKeys[legacy]; !ok {
			t.Errorf("RenamedValues key %q is not in RenamedKeys", legacy)
		}
	}
}
