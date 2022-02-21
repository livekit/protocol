package livekit

//----------------------------------------------------------------

type TrackID string

func StringsAsTrackIDs(trackIDs []string) []TrackID {
	asTrackID := make([]TrackID, 0, len(trackIDs))
	for _, trackID := range trackIDs {
		asTrackID = append(asTrackID, TrackID(trackID))
	}

	return asTrackID
}

//----------------------------------------------------------------

type ParticipantID string
type ParticipantIdentity string
type ParticipantName string

type RoomID string

//----------------------------------------------------------------

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

//----------------------------------------------------------------

type ConnectionID string
type NodeID string
type ParticipantKey string
