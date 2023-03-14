package livekit

import (
	context "context"
	fmt "fmt"
)

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

type ParticipantTopic string
type RoomTopic string

func FormatParticipantTopic(roomName RoomName, identity ParticipantIdentity) ParticipantTopic {
	return ParticipantTopic(fmt.Sprintf("%s_%s", roomName, identity))
}

func FormatRoomTopic(roomName RoomName) RoomTopic {
	return RoomTopic(roomName)
}

type TopicFormatter interface {
	ParticipantTopic(ctx context.Context, roomName RoomName, identity ParticipantIdentity) ParticipantTopic
	RoomTopic(ctx context.Context, roomName RoomName) RoomTopic
}

type topicFormatter struct{}

func NewTopicFormatter() TopicFormatter {
	return topicFormatter{}
}

func (f topicFormatter) ParticipantTopic(ctx context.Context, roomName RoomName, identity ParticipantIdentity) ParticipantTopic {
	return FormatParticipantTopic(roomName, identity)
}

func (f topicFormatter) RoomTopic(ctx context.Context, roomName RoomName) RoomTopic {
	return FormatRoomTopic(roomName)
}
