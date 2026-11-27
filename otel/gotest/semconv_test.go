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

// TestRenamedKeys checks the legacy-key alias map against the eight documented OTLP client renames
// as a subset (the map now also carries the pre-#7973 Agents span names and the Go logger /
// collector spellings as aliases), and its structural invariants: a legacy key is never empty, maps
// to a non-empty canonical key, and is never equal to its own canonical key (a rename changes the
// spelling).
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

// TestKeyAliases checks the canonical-to-aliases map and its cross-map invariants: every canonical
// key is a registered attribute, every alias maps back to that canonical in RenamedKeys, an alias
// is never a canonical key itself, and the two maps agree (RenamedKeys has one entry per alias
// across all KeyAliases lists).
func TestKeyAliases(t *testing.T) {
	registered := map[string]bool{
		semconv.AttrLkRoomId:                       true,
		semconv.AttrLkParticipantId:                true,
		semconv.AttrLkTrackId:                      true,
		semconv.AttrLkPiiRoomName:                  true,
		semconv.AttrLkPiiParticipantIdentity:       true,
		semconv.AttrLkPiiRemoteParticipantIdentity: true,
		semconv.AttrLkTrackSource:                  true,
		semconv.AttrLkDisconnectReason:             true,
	}
	aliasCount := 0
	for canonical, aliases := range semconv.KeyAliases {
		if !registered[canonical] {
			t.Errorf("KeyAliases has canonical key %q that is not a registered attribute", canonical)
		}
		for _, alias := range aliases {
			aliasCount++
			if alias == canonical {
				t.Errorf("KeyAliases[%q] lists the canonical key as its own alias", canonical)
			}
			if got := semconv.RenamedKeys[alias]; got != canonical {
				t.Errorf("KeyAliases[%q] has alias %q, but RenamedKeys[%q] = %q", canonical, alias, alias, got)
			}
			if _, ok := semconv.KeyAliases[alias]; ok {
				t.Errorf("alias %q is also a canonical key in KeyAliases", alias)
			}
		}
	}
	if aliasCount != len(semconv.RenamedKeys) {
		t.Errorf("KeyAliases lists %d aliases, RenamedKeys has %d entries; they must match", aliasCount, len(semconv.RenamedKeys))
	}
}

// TestNormalizeKey checks that a legacy spelling folds onto its canonical key, an unknown key is
// returned unchanged, and a canonical key is already canonical.
func TestNormalizeKey(t *testing.T) {
	cases := map[string]string{
		"roomID":             semconv.AttrLkRoomId,
		"room_id":            semconv.AttrLkRoomId,
		"lk.room.sid":        semconv.AttrLkRoomId,
		"participantID":      semconv.AttrLkParticipantId,
		semconv.AttrLkRoomId: semconv.AttrLkRoomId,
		"app.custom.key":     "app.custom.key",
	}
	for key, want := range cases {
		if got := semconv.NormalizeKey(key); got != want {
			t.Errorf("NormalizeKey(%q) = %q, want %q", key, got, want)
		}
	}
}

// TestAliasesFor checks that the canonical key comes first, its aliases follow, and a key with no
// recorded legacy spelling yields just itself.
func TestAliasesFor(t *testing.T) {
	got := semconv.AliasesFor(semconv.AttrLkRoomId)
	if len(got) == 0 || got[0] != semconv.AttrLkRoomId {
		t.Fatalf("AliasesFor(%q) = %v, want canonical first", semconv.AttrLkRoomId, got)
	}
	found := false
	for _, a := range got[1:] {
		if a == "roomID" {
			found = true
		}
		if a == semconv.AttrLkRoomId {
			t.Errorf("AliasesFor(%q) repeats the canonical key", semconv.AttrLkRoomId)
		}
	}
	if !found {
		t.Errorf("AliasesFor(%q) = %v, missing alias roomID", semconv.AttrLkRoomId, got)
	}
	if got := semconv.AliasesFor("app.custom.key"); len(got) != 1 || got[0] != "app.custom.key" {
		t.Errorf("AliasesFor(%q) = %v, want just itself", "app.custom.key", got)
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

// TestNormalizeValue checks that a legacy value under a legacy key folds onto its canonical value,
// and an unmapped key or value is returned unchanged.
func TestNormalizeValue(t *testing.T) {
	if got := semconv.NormalizeValue("lk.track.source", "camera"); got != semconv.LkTrackSourceSourceCamera {
		t.Errorf("NormalizeValue(track.source, camera) = %q, want %q", got, semconv.LkTrackSourceSourceCamera)
	}
	if got := semconv.NormalizeValue("lk.disconnect.reason", "client_initiated"); got != semconv.LkDisconnectReasonClientInitiated {
		t.Errorf("NormalizeValue(disconnect.reason, client_initiated) = %q, want %q", got, semconv.LkDisconnectReasonClientInitiated)
	}
	if got := semconv.NormalizeValue("lk.track.source", "SOURCE_CAMERA"); got != "SOURCE_CAMERA" {
		t.Errorf("NormalizeValue of an already-canonical value = %q, want unchanged", got)
	}
	if got := semconv.NormalizeValue("lk.room.sid", "RM_abc"); got != "RM_abc" {
		t.Errorf("NormalizeValue of a key with no value map = %q, want unchanged", got)
	}
}
