// Code generated; DO NOT EDIT.
package telephonycallobs

type CallDirection string

const (
	CallDirectionUndefined CallDirection = ""
	CallDirectionUnknown   CallDirection = "unknown"
	CallDirectionInbound   CallDirection = "inbound"
	CallDirectionOutbound  CallDirection = "outbound"
)

type CallTransport string

const (
	CallTransportUndefined CallTransport = ""
	CallTransportAuto      CallTransport = "auto"
	CallTransportUDP       CallTransport = "udp"
	CallTransportTCP       CallTransport = "tcp"
	CallTransportTLS       CallTransport = "tls"
)

type CallStatus string

const (
	CallStatusUndefined         CallStatus = ""
	CallStatusActive            CallStatus = "active"
	CallStatusCallIncoming      CallStatus = "call_incoming"
	CallStatusParticipantJoined CallStatus = "participant_joined"
	CallStatusDisconnected      CallStatus = "disconnected"
	CallStatusError             CallStatus = "error"
)

type CallTransferStatus string

const (
	CallTransferStatusUndefined CallTransferStatus = ""
	CallTransferStatusOngoing   CallTransferStatus = "ongoing"
	CallTransferStatusSuccess   CallTransferStatus = "success"
	CallTransferStatusFailed    CallTransferStatus = "failed"
)

type CallEncryption string

const (
	CallEncryptionUndefined CallEncryption = ""
	CallEncryptionDisable   CallEncryption = "disable"
	CallEncryptionAllow     CallEncryption = "allow"
	CallEncryptionRequire   CallEncryption = "require"
)

type Rollup string

const (
	RollupUndefined      Rollup = ""
	RollupProject        Rollup = "project"
	RollupCallIndex      Rollup = "call_index"
	RollupStartTimeIndex Rollup = "start_time_index"
	RollupEndTimeIndex   Rollup = "end_time_index"
)
