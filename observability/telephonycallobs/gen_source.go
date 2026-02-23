// Code generated; DO NOT EDIT.
package telephonycallobs

type CallTrunkType string

const (
	CallTrunkTypeUndefined CallTrunkType = ""
	CallTrunkTypeInternal  CallTrunkType = "internal"
	CallTrunkTypeExternal  CallTrunkType = "external"
)

type CallNumberType string

const (
	CallNumberTypeUndefined CallNumberType = ""
	CallNumberTypeTollFree  CallNumberType = "toll_free"
	CallNumberTypeRegular   CallNumberType = "regular"
)

type CallTransport string

const (
	CallTransportUndefined CallTransport = ""
	CallTransportUDP       CallTransport = "udp"
	CallTransportTCP       CallTransport = "tcp"
	CallTransportTLS       CallTransport = "tls"
)

type CallTransferStatus string

const (
	CallTransferStatusUndefined CallTransferStatus = ""
	CallTransferStatusOngoing   CallTransferStatus = "ongoing"
	CallTransferStatusSuccess   CallTransferStatus = "success"
	CallTransferStatusFailed    CallTransferStatus = "failed"
)

type CallMediaEncryptionSettings string

const (
	CallMediaEncryptionSettingsUndefined CallMediaEncryptionSettings = ""
	CallMediaEncryptionSettingsDisable   CallMediaEncryptionSettings = "disable"
	CallMediaEncryptionSettingsAllow     CallMediaEncryptionSettings = "allow"
	CallMediaEncryptionSettingsRequire   CallMediaEncryptionSettings = "require"
)

type Rollup string

const (
	RollupUndefined      Rollup = ""
	RollupProject        Rollup = "project"
	RollupCallIndex      Rollup = "call_index"
	RollupStartTimeIndex Rollup = "start_time_index"
	RollupEndTimeIndex   Rollup = "end_time_index"
)
