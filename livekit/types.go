package livekit

//----------------------------------------------------------------

type TrackID string

func StringsAsTrackIDs(trackIDs []string) []TrackID {
	var asTrackID []TrackID
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
	var asString []string
	for _, roomName := range roomNames {
		asString = append(asString, string(roomName))
	}

	return asString
}

func StringsAsRoomNames(roomNames []string) []RoomName {
	var asRoomName []RoomName
	for _, roomName := range roomNames {
		asRoomName = append(asRoomName, RoomName(roomName))
	}

	return asRoomName
}

//----------------------------------------------------------------
