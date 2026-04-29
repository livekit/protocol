// Code generated; DO NOT EDIT.

package corecallobs

import (
	"time"
)

const Version_857KEC0 = true

type KeyResolver interface {
	Resolve(string)
	Reset()
}

type Reporter interface {
	WithProject(id string) ProjectReporter
	WithDeferredProject() (ProjectReporter, KeyResolver)
}

type projectReporter interface {
}

type ProjectTx interface {
	projectReporter
}

type ProjectReporter interface {
	RegisterFunc(func(ts time.Time, tx ProjectTx) bool)
	Tx(func(tx ProjectTx))
	TxAt(time.Time, func(tx ProjectTx))
	WithCall(id string) CallReporter
	WithDeferredCall() (CallReporter, KeyResolver)
	projectReporter
}

type callReporter interface {
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
	ReportDuration(v uint64)
	ReportDurationMinutes(v uint16)
	ReportDirection(v CallDirection)
	ReportCallType(v CallCallType)
	ReportFrom(v string)
	ReportTo(v string)
	ReportRegion(v string)
	ReportRoomID(v string)
	ReportRoomName(v string)
	ReportError(v string)
	ReportStatus(v CallStatus)
}

type CallTx interface {
	Project() ProjectTx
	callReporter
}

type CallReporter interface {
	RegisterFunc(func(ts time.Time, tx CallTx) bool)
	Tx(func(tx CallTx))
	TxAt(time.Time, func(tx CallTx))
	callReporter
}
