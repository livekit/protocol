// Code generated; DO NOT EDIT.

package agentsobs

import (
	"time"
)

const Version_P0N4MSO = true

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
	WithCloudAgent(id string) CloudAgentReporter
	WithDeferredCloudAgent() (CloudAgentReporter, KeyResolver)
}

type ProjectTx interface{}

type CloudAgentReporter interface {
	RegisterFunc(func(ts time.Time, tx CloudAgentTx) bool)
	Tx(func(tx CloudAgentTx))
	TxAt(time.Time, func(tx CloudAgentTx))
	WithAgent(name string) AgentReporter
	WithDeferredAgent() (AgentReporter, KeyResolver)
}

type CloudAgentTx interface{}

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
	ReportLoad(v float32)
	ReportStatus(v WorkerStatus)
	ReportStartTime(v time.Time)
	ReportJobsCurrent(v uint32)
	ReportLive(v uint8)
	ReportCPU(v int64)
	ReportCPULimit(v int64)
	ReportMem(v int64)
	ReportMemLimit(v int64)
	ReportRegion(v string)
	ReportVersion(v string)
	ReportSdkVersion(v string)
	ReportState(v WorkerState)
}

type WorkerTx interface {
	ReportLoad(v float32)
	ReportStatus(v WorkerStatus)
	ReportStartTime(v time.Time)
	ReportJobsCurrent(v uint32)
	ReportLive(v uint8)
	ReportCPU(v int64)
	ReportCPULimit(v int64)
	ReportMem(v int64)
	ReportMemLimit(v int64)
	ReportRegion(v string)
	ReportVersion(v string)
	ReportSdkVersion(v string)
	ReportState(v WorkerState)
}

type JobReporter interface {
	RegisterFunc(func(ts time.Time, tx JobTx) bool)
	Tx(func(tx JobTx))
	TxAt(time.Time, func(tx JobTx))
	ReportRoomSessionID(v string)
	ReportKind(v JobKind)
	ReportWorkerKind(v WorkerKind)
	ReportStatus(v JobStatus)
	ReportDuration(v uint32)
	ReportDurationMinutes(v uint8)
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
	ReportJoinLatency(v uint32)
}

type JobTx interface {
	ReportRoomSessionID(v string)
	ReportKind(v JobKind)
	ReportWorkerKind(v WorkerKind)
	ReportStatus(v JobStatus)
	ReportDuration(v uint32)
	ReportDurationMinutes(v uint8)
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
	ReportJoinLatency(v uint32)
}
