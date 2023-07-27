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

package livekit

// ----------------------------------------------------------------

type TrackID string

func StringsAsTrackIDs(trackIDs []string) []TrackID {
	asTrackID := make([]TrackID, 0, len(trackIDs))
	for _, trackID := range trackIDs {
		asTrackID = append(asTrackID, TrackID(trackID))
	}

	return asTrackID
}

// ----------------------------------------------------------------

type ParticipantID string

func ParticipantIDsAsStrings(ids []ParticipantID) []string {
	strs := make([]string, 0, len(ids))
	for _, id := range ids {
		strs = append(strs, string(id))
	}
	return strs
}

// ----------------------------------------------------------------

type ParticipantIdentity string
type ParticipantName string

type RoomID string

// ----------------------------------------------------------------

type RoomName string

func RoomNamesAsStrings(roomNames []RoomName) []string {
	asString := make([]string, 0, len(roomNames))
	for _, roomName := range roomNames {
		asString = append(asString, string(roomName))
	}

	return asString
}

func StringsAsRoomNames(roomNames []string) []RoomName {
	asRoomName := make([]RoomName, 0, len(roomNames))
	for _, roomName := range roomNames {
		asRoomName = append(asRoomName, RoomName(roomName))
	}

	return asRoomName
}

// ----------------------------------------------------------------

type ConnectionID string

// ----------------------------------------------------------------

type NodeID string

func NodeIDsAsStrings(ids []NodeID) []string {
	strs := make([]string, 0, len(ids))
	for _, id := range ids {
		strs = append(strs, string(id))
	}
	return strs
}

// ----------------------------------------------------------------
type ParticipantKey string

// ----------------------------------------------------------------

type Guid interface {
	TrackID | ParticipantID | RoomID
}

type GuidBlock [9]byte
