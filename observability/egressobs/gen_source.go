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
)

type EgressStatus string

const (
	EgressStatusUndefined    EgressStatus = ""
	EgressStatusStarting     EgressStatus = "starting"
	EgressStatusActive       EgressStatus = "active"
	EgressStatusEnding       EgressStatus = "ending"
	EgressStatusComplete     EgressStatus = "complete"
	EgressStatusFailed       EgressStatus = "failed"
	EgressStatusAborted      EgressStatus = "aborted"
	EgressStatusLimitReached EgressStatus = "limit_reached"
)

type SessionSourceType string

const (
	SessionSourceTypeUndefined SessionSourceType = ""
	SessionSourceTypeSdk       SessionSourceType = "sdk"
	SessionSourceTypeWeb       SessionSourceType = "web"
)

type SessionStatus string

const (
	SessionStatusUndefined    SessionStatus = ""
	SessionStatusStarting     SessionStatus = "starting"
	SessionStatusActive       SessionStatus = "active"
	SessionStatusEnding       SessionStatus = "ending"
	SessionStatusComplete     SessionStatus = "complete"
	SessionStatusFailed       SessionStatus = "failed"
	SessionStatusAborted      SessionStatus = "aborted"
	SessionStatusLimitReached SessionStatus = "limit_reached"
)

type Rollup string

const (
	RollupUndefined      Rollup = ""
	RollupProject        Rollup = "project"
	RollupEgressIndex    Rollup = "egress_index"
	RollupEndTimeIndex   Rollup = "end_time_index"
	RollupStartTimeIndex Rollup = "start_time_index"
	RollupRoomNameIndex  Rollup = "room_name_index"
	RollupEgressEgress   Rollup = "egress_egress"
	RollupSessionIndex   Rollup = "session_index"
)
