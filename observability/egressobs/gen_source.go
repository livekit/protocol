// Code generated; DO NOT EDIT.
package egressobs

type EgressRequestType string

const (
	EgressRequestTypeUndefined      EgressRequestType = ""
	EgressRequestTypeRoomComposite  EgressRequestType = "room_composite"
	EgressRequestTypeTrackComposite EgressRequestType = "track_composite"
	EgressRequestTypeTrack          EgressRequestType = "track"
	EgressRequestTypeParticipant    EgressRequestType = "participant"
	EgressRequestTypeWeb            EgressRequestType = "web"
	EgressRequestTypeEgress         EgressRequestType = "egress"
	EgressRequestTypeReplay         EgressRequestType = "replay"
)

type SessionSourceType string

const (
	SessionSourceTypeUndefined SessionSourceType = ""
	SessionSourceTypeSdk       SessionSourceType = "sdk"
	SessionSourceTypeWeb       SessionSourceType = "web"
	SessionSourceTypeTemplate  SessionSourceType = "template"
	SessionSourceTypeMedia     SessionSourceType = "media"
)

type Rollup string

const (
	RollupUndefined Rollup = ""
	RollupProject   Rollup = "project"
	// Deprecated: removed from schema; retained for compatibility with stored data.
	RollupEgressIndex    Rollup = "egress_index"
	RollupEndTimeIndex   Rollup = "end_time_index"
	RollupStartTimeIndex Rollup = "start_time_index"
	RollupRoomNameIndex  Rollup = "room_name_index"
	RollupEgressEgress   Rollup = "egress_egress"
	RollupSessionIndex   Rollup = "session_index"
)
