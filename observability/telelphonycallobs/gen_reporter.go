// Code generated; DO NOT EDIT.

package telelphonycallobs

import (
	"time"
)

const Version_J8004V8 = true

type KeyResolver interface {
	Resolve(string)
	Reset()
}

type Reporter interface {
	WithProject(id string) ProjectReporter
	WithDeferredProject() (ProjectReporter, KeyResolver)
}

type ProjectTx interface{}

type ProjectReporter interface {
	RegisterFunc(func(ts time.Time, tx ProjectTx) bool)
	Tx(func(tx ProjectTx))
	TxAt(time.Time, func(tx ProjectTx))
	WithCall(id string) CallReporter
	WithDeferredCall() (CallReporter, KeyResolver)
	ProjectTx
}

type CallTx interface {
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
	ReportDuration(v uint64)
	ReportDurationMinutes(v uint16)
	ReportTrunkID(v string)
	ReportDispatchID(v string)
	ReportCalleeNumber(v string)
	ReportCalleeHost(v string)
	ReportCallerNumber(v string)
	ReportCallerHost(v string)
	ReportDirection(v CallDirection)
	ReportTransport(v CallTransport)
	ReportProviderCallID(v string)
	ReportProviderName(v string)
	ReportSIPCallID(v string)
	ReportRoomID(v string)
	ReportRoomName(v string)
	ReportParticipantIdentity(v string)
	ReportError(v string)
	ReportStatus(v CallStatus)
	ReportResponseCode(v uint16)
	ReportDisconnectReason(v string)
	ReportTransferID(v string)
	ReportTransferTo(v string)
	ReportTransferDuration(v uint32)
	ReportTransferStatus(v CallTransferStatus)
	ReportTransferStatusCode(v uint16)
	ReportTransferError(v string)
	ReportCodec(v string)
	ReportRegion(v string)
	ReportPcapLink(v string)
	ReportAttributes(v string)
	ReportFeatures(v uint16)
	ReportEncryption(v CallEncryption)
}

type CallReporter interface {
	RegisterFunc(func(ts time.Time, tx CallTx) bool)
	Tx(func(tx CallTx))
	TxAt(time.Time, func(tx CallTx))
	CallTx
}
