// Code generated; DO NOT EDIT.

package telephonycallobs

import (
	"time"
)

const Version_T93RR0G = true

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
	ReportTrunkID(v string)
	ReportTrunkType(v CallTrunkType)
	ReportDispatchID(v string)
	ReportToNumber(v string)
	ReportToHost(v string)
	ReportFromNumber(v string)
	ReportFromHost(v string)
	ReportNumberType(v CallNumberType)
	ReportCountryCode(v string)
	ReportTransport(v CallTransport)
	ReportProviderCallID(v string)
	ReportProviderName(v string)
	ReportSIPCallID(v string)
	ReportParticipantIdentity(v string)
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
	ReportMediaEncryptionSettings(v CallMediaEncryptionSettings)
	ReportMediaEncryption(v string)
}

type CallReporter interface {
	RegisterFunc(func(ts time.Time, tx CallTx) bool)
	Tx(func(tx CallTx))
	TxAt(time.Time, func(tx CallTx))
	CallTx
}
