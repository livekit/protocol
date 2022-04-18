package livekit

import strings "strings"

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

type CodecMime string

func CodecMimeFromClientCodec(codec string, codecType TrackType) CodecMime {
	if codecType == TrackType_VIDEO && !strings.HasPrefix(codec, "video/") {
		codec = "video/" + codec
	} else if codecType == TrackType_AUDIO && !strings.HasPrefix(codec, "audio/") {
		codec = "audio/" + codec
	}
	return CodecMime(codec)
}

func (c CodecMime) ToClientCodec() string {
	codec := strings.TrimPrefix(string(c), "video/")
	codec = strings.TrimPrefix(codec, "audio/")
	return codec
}
