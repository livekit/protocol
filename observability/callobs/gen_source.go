// Code generated; DO NOT EDIT.
package callobs

type CallDirection string

const (
	CallDirectionUndefined CallDirection = ""
	CallDirectionUnknown   CallDirection = "unknown"
	CallDirectionInbound   CallDirection = "inbound"
	CallDirectionOutbound  CallDirection = "outbound"
)

type CallCallType string

const (
	CallCallTypeUndefined CallCallType = ""
	CallCallTypeSIP       CallCallType = "sip"
	CallCallTypeTwilio    CallCallType = "twilio"
	CallCallTypeWhatsapp  CallCallType = "whatsapp"
)

type CallStatus string

const (
	CallStatusUndefined         CallStatus = ""
	CallStatusActive            CallStatus = "active"
	CallStatusCallIncoming      CallStatus = "call_incoming"
	CallStatusParticipantJoined CallStatus = "participant_joined"
	CallStatusDisconnected      CallStatus = "disconnected"
	CallStatusError             CallStatus = "error"
	CallStatusPending           CallStatus = "pending"
	CallStatusSuccess           CallStatus = "success"
	CallStatusFailed            CallStatus = "failed"
)

type Rollup string

const (
	RollupUndefined      Rollup = ""
	RollupProject        Rollup = "project"
	RollupCallIndex      Rollup = "call_index"
	RollupStartTimeIndex Rollup = "start_time_index"
	RollupEndTimeIndex   Rollup = "end_time_index"
)
