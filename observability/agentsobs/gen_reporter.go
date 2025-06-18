// Code generated; DO NOT EDIT.

package agentsobs

import (
	"time"
)

const Version_1MM59JG = true

type KeyResolver interface {
	Resolve(string)
	Reset()
}

type Reporter interface {
	WithProject(id string) ProjectReporter
	WithDeferredProject() (ProjectReporter, KeyResolver)
}

type ProjectReporter interface {
	RegisterFunc(func(ts time.Time, tx ProjectTx) bool)
	Tx(func(tx ProjectTx))
	TxAt(time.Time, func(tx ProjectTx))
	WithAgent(name string) AgentReporter
	WithDeferredAgent() (AgentReporter, KeyResolver)
}

type ProjectTx interface{}

type AgentReporter interface {
	RegisterFunc(func(ts time.Time, tx AgentTx) bool)
	Tx(func(tx AgentTx))
	TxAt(time.Time, func(tx AgentTx))
	WithWorker(id string) WorkerReporter
	WithDeferredWorker() (WorkerReporter, KeyResolver)
}

type AgentTx interface{}

type WorkerReporter interface {
	RegisterFunc(func(ts time.Time, tx WorkerTx) bool)
	Tx(func(tx WorkerTx))
	TxAt(time.Time, func(tx WorkerTx))
	WithJob(id string) JobReporter
	WithDeferredJob() (JobReporter, KeyResolver)
	ReportCPU(v int64)
	ReportCPULimit(v int64)
	ReportMem(v int64)
	ReportMemLimit(v int64)
	ReportLoad(v float32)
	ReportStatus(v WorkerStatus)
	ReportRegion(v string)
	ReportVersion(v string)
	ReportSdkVersion(v string)
	ReportState(v WorkerState)
	ReportStartedAt(v time.Time)
	ReportJobsCurrent(v uint16)
	ReportKind(v AgentKind)
}

type WorkerTx interface {
	ReportCPU(v int64)
	ReportCPULimit(v int64)
	ReportMem(v int64)
	ReportMemLimit(v int64)
	ReportLoad(v float32)
	ReportStatus(v WorkerStatus)
	ReportRegion(v string)
	ReportVersion(v string)
	ReportSdkVersion(v string)
	ReportState(v WorkerState)
	ReportStartedAt(v time.Time)
	ReportJobsCurrent(v uint16)
	ReportKind(v AgentKind)
}

type JobReporter interface {
	RegisterFunc(func(ts time.Time, tx JobTx) bool)
	Tx(func(tx JobTx))
	TxAt(time.Time, func(tx JobTx))
	ReportRoomSessionID(v string)
	ReportKind(v JobKind)
	ReportStatus(v JobStatus)
	ReportDuration(v uint32)
	ReportDispatchedAt(v time.Time)
	ReportJoinedAt(v time.Time)
	ReportCompletedAt(v time.Time)
}

type JobTx interface {
	ReportRoomSessionID(v string)
	ReportKind(v JobKind)
	ReportStatus(v JobStatus)
	ReportDuration(v uint32)
	ReportDispatchedAt(v time.Time)
	ReportJoinedAt(v time.Time)
	ReportCompletedAt(v time.Time)
}
