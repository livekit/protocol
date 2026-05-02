// Code generated; DO NOT EDIT.

package egressobs

import (
	"time"
)

const Version_HAKUONO = true

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
	WithEgress(id string) EgressReporter
	WithDeferredEgress() (EgressReporter, KeyResolver)
	projectReporter
}

type egressReporter interface {
	ReportRequestType(v EgressRequestType)
	ReportRoomName(v string)
	ReportRequest(v string)
	ReportAudioOnly(v bool)
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
	ReportUpdateTime(v time.Time)
	ReportStatus(v string)
	ReportDetails(v string)
	ReportError(v string)
	ReportErrorCode(v int32)
	ReportResult(v string)
	ReportManifestLocation(v string)
}

type EgressTx interface {
	Project() ProjectTx
	egressReporter
}

type EgressReporter interface {
	RegisterFunc(func(ts time.Time, tx EgressTx) bool)
	Tx(func(tx EgressTx))
	TxAt(time.Time, func(tx EgressTx))
	WithSession(id string) SessionReporter
	WithDeferredSession() (SessionReporter, KeyResolver)
	egressReporter
}

type sessionReporter interface {
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
	ReportUpdateTime(v time.Time)
	ReportDuration(v uint64)
	ReportDurationSeconds(v uint64)
	ReportRetryCount(v uint32)
	ReportSourceType(v SessionSourceType)
	ReportRegion(v string)
	ReportRoomID(v string)
	ReportStatus(v string)
	ReportDetails(v string)
	ReportError(v string)
	ReportErrorCode(v int32)
	ReportManifestLocation(v string)
	ReportBackupStorageUsed(v bool)
	ReportResult(v string)
}

type SessionTx interface {
	Egress() EgressTx
	sessionReporter
}

type SessionReporter interface {
	RegisterFunc(func(ts time.Time, tx SessionTx) bool)
	Tx(func(tx SessionTx))
	TxAt(time.Time, func(tx SessionTx))
	sessionReporter
}
