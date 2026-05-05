// Code generated; DO NOT EDIT.

package agentsv3obs

import (
	"time"
)

const Version_MAT8U20 = true

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
	WithCloudAgent(id string) CloudAgentReporter
	WithDeferredCloudAgent() (CloudAgentReporter, KeyResolver)
	projectReporter
}

type cloudAgentReporter interface {
}

type CloudAgentTx interface {
	Project() ProjectTx
	cloudAgentReporter
}

type CloudAgentReporter interface {
	RegisterFunc(func(ts time.Time, tx CloudAgentTx) bool)
	Tx(func(tx CloudAgentTx))
	TxAt(time.Time, func(tx CloudAgentTx))
	WithAgent(name string) AgentReporter
	WithDeferredAgent() (AgentReporter, KeyResolver)
	cloudAgentReporter
}

type agentReporter interface {
}

type AgentTx interface {
	CloudAgent() CloudAgentTx
	agentReporter
}

type AgentReporter interface {
	RegisterFunc(func(ts time.Time, tx AgentTx) bool)
	Tx(func(tx AgentTx))
	TxAt(time.Time, func(tx AgentTx))
	WithDeployment(name string) DeploymentReporter
	WithDeferredDeployment() (DeploymentReporter, KeyResolver)
	agentReporter
}

type deploymentReporter interface {
}

type DeploymentTx interface {
	Agent() AgentTx
	deploymentReporter
}

type DeploymentReporter interface {
	RegisterFunc(func(ts time.Time, tx DeploymentTx) bool)
	Tx(func(tx DeploymentTx))
	TxAt(time.Time, func(tx DeploymentTx))
	WithWorker(id string) WorkerReporter
	WithDeferredWorker() (WorkerReporter, KeyResolver)
	deploymentReporter
}

type workerReporter interface {
	ReportLoad(v float32)
	ReportStatus(v WorkerStatus)
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
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
	Deployment() DeploymentTx
	workerReporter
}

type WorkerReporter interface {
	RegisterFunc(func(ts time.Time, tx WorkerTx) bool)
	Tx(func(tx WorkerTx))
	TxAt(time.Time, func(tx WorkerTx))
	WithJob(id string) JobReporter
	WithDeferredJob() (JobReporter, KeyResolver)
	workerReporter
}

type jobReporter interface {
	ReportRoomSessionID(v string)
	ReportKind(v JobKind)
	ReportWorkerKind(v WorkerKind)
	ReportStatus(v JobStatus)
	ReportDuration(v uint32)
	ReportDurationSeconds(v uint32)
	ReportDurationMinutes(v uint8)
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
	ReportJoinLatency(v uint32)
}

type JobTx interface {
	Worker() WorkerTx
	jobReporter
}

type JobReporter interface {
	RegisterFunc(func(ts time.Time, tx JobTx) bool)
	Tx(func(tx JobTx))
	TxAt(time.Time, func(tx JobTx))
	jobReporter
}
