// Code generated; DO NOT EDIT.

package agentsobs

import (
	"time"
)

var (
	_ Reporter           = (*noopReporter)(nil)
	_ ProjectReporter    = (*noopProjectReporter)(nil)
	_ ProjectTx          = (*noopProjectTx)(nil)
	_ CloudAgentReporter = (*noopCloudAgentReporter)(nil)
	_ CloudAgentTx       = (*noopCloudAgentTx)(nil)
	_ AgentReporter      = (*noopAgentReporter)(nil)
	_ AgentTx            = (*noopAgentTx)(nil)
	_ WorkerReporter     = (*noopWorkerReporter)(nil)
	_ WorkerTx           = (*noopWorkerTx)(nil)
	_ JobReporter        = (*noopJobReporter)(nil)
	_ JobTx              = (*noopJobTx)(nil)
)

type noopKeyResolver struct{}

func (noopKeyResolver) Resolve(string) {}
func (noopKeyResolver) Reset()         {}

type noopReporter struct{}

func NewNoopReporter() Reporter {
	return &noopReporter{}
}

func (r *noopReporter) WithProject(id string) ProjectReporter {
	return &noopProjectReporter{}
}

func (r *noopReporter) WithDeferredProject() (ProjectReporter, KeyResolver) {
	return &noopProjectReporter{}, noopKeyResolver{}
}

type noopProjectReporter struct{}

func NewNoopProjectReporter() ProjectReporter {
	return &noopProjectReporter{}
}

func (r *noopProjectReporter) RegisterFunc(f func(ts time.Time, tx ProjectTx) bool) {}
func (r *noopProjectReporter) Tx(f func(ProjectTx))                                 {}
func (r *noopProjectReporter) TxAt(ts time.Time, f func(ProjectTx))                 {}
func (r *noopProjectReporter) WithCloudAgent(id string) CloudAgentReporter {
	return &noopCloudAgentReporter{}
}
func (r *noopProjectReporter) WithDeferredCloudAgent() (CloudAgentReporter, KeyResolver) {
	return &noopCloudAgentReporter{}, noopKeyResolver{}
}

type noopProjectTx struct{}

type noopCloudAgentReporter struct{}

func NewNoopCloudAgentReporter() CloudAgentReporter {
	return &noopCloudAgentReporter{}
}

func (r *noopCloudAgentReporter) RegisterFunc(f func(ts time.Time, tx CloudAgentTx) bool) {}
func (r *noopCloudAgentReporter) Tx(f func(CloudAgentTx))                                 {}
func (r *noopCloudAgentReporter) TxAt(ts time.Time, f func(CloudAgentTx))                 {}
func (r *noopCloudAgentReporter) WithAgent(name string) AgentReporter {
	return &noopAgentReporter{}
}
func (r *noopCloudAgentReporter) WithDeferredAgent() (AgentReporter, KeyResolver) {
	return &noopAgentReporter{}, noopKeyResolver{}
}

type noopCloudAgentTx struct{}

func (t *noopCloudAgentTx) Project() ProjectTx {
	return &noopProjectTx{}
}

type noopAgentReporter struct{}

func NewNoopAgentReporter() AgentReporter {
	return &noopAgentReporter{}
}

func (r *noopAgentReporter) RegisterFunc(f func(ts time.Time, tx AgentTx) bool) {}
func (r *noopAgentReporter) Tx(f func(AgentTx))                                 {}
func (r *noopAgentReporter) TxAt(ts time.Time, f func(AgentTx))                 {}
func (r *noopAgentReporter) WithWorker(id string) WorkerReporter {
	return &noopWorkerReporter{}
}
func (r *noopAgentReporter) WithDeferredWorker() (WorkerReporter, KeyResolver) {
	return &noopWorkerReporter{}, noopKeyResolver{}
}

type noopAgentTx struct{}

func (t *noopAgentTx) CloudAgent() CloudAgentTx {
	return &noopCloudAgentTx{}
}

type noopWorkerReporter struct{}

func NewNoopWorkerReporter() WorkerReporter {
	return &noopWorkerReporter{}
}

func (r *noopWorkerReporter) RegisterFunc(f func(ts time.Time, tx WorkerTx) bool) {}
func (r *noopWorkerReporter) Tx(f func(WorkerTx))                                 {}
func (r *noopWorkerReporter) TxAt(ts time.Time, f func(WorkerTx))                 {}
func (r *noopWorkerReporter) ReportLoad(v float32)                                {}
func (r *noopWorkerReporter) ReportStatus(v WorkerStatus)                         {}
func (r *noopWorkerReporter) ReportStartTime(v time.Time)                         {}
func (r *noopWorkerReporter) ReportEndTime(v time.Time)                           {}
func (r *noopWorkerReporter) ReportJobsCurrent(v uint32)                          {}
func (r *noopWorkerReporter) ReportLive(v uint8)                                  {}
func (r *noopWorkerReporter) ReportCPU(v int64)                                   {}
func (r *noopWorkerReporter) ReportCPULimit(v int64)                              {}
func (r *noopWorkerReporter) ReportMem(v int64)                                   {}
func (r *noopWorkerReporter) ReportMemLimit(v int64)                              {}
func (r *noopWorkerReporter) ReportRegion(v string)                               {}
func (r *noopWorkerReporter) ReportVersion(v string)                              {}
func (r *noopWorkerReporter) ReportSdkVersion(v string)                           {}
func (r *noopWorkerReporter) ReportState(v WorkerState)                           {}
func (r *noopWorkerReporter) WithJob(id string) JobReporter {
	return &noopJobReporter{}
}
func (r *noopWorkerReporter) WithDeferredJob() (JobReporter, KeyResolver) {
	return &noopJobReporter{}, noopKeyResolver{}
}

type noopWorkerTx struct{}

func (t *noopWorkerTx) Agent() AgentTx {
	return &noopAgentTx{}
}

func (t *noopWorkerTx) ReportLoad(v float32)        {}
func (t *noopWorkerTx) ReportStatus(v WorkerStatus) {}
func (t *noopWorkerTx) ReportStartTime(v time.Time) {}
func (t *noopWorkerTx) ReportEndTime(v time.Time)   {}
func (t *noopWorkerTx) ReportJobsCurrent(v uint32)  {}
func (t *noopWorkerTx) ReportLive(v uint8)          {}
func (t *noopWorkerTx) ReportCPU(v int64)           {}
func (t *noopWorkerTx) ReportCPULimit(v int64)      {}
func (t *noopWorkerTx) ReportMem(v int64)           {}
func (t *noopWorkerTx) ReportMemLimit(v int64)      {}
func (t *noopWorkerTx) ReportRegion(v string)       {}
func (t *noopWorkerTx) ReportVersion(v string)      {}
func (t *noopWorkerTx) ReportSdkVersion(v string)   {}
func (t *noopWorkerTx) ReportState(v WorkerState)   {}

type noopJobReporter struct{}

func NewNoopJobReporter() JobReporter {
	return &noopJobReporter{}
}

func (r *noopJobReporter) RegisterFunc(f func(ts time.Time, tx JobTx) bool) {}
func (r *noopJobReporter) Tx(f func(JobTx))                                 {}
func (r *noopJobReporter) TxAt(ts time.Time, f func(JobTx))                 {}
func (r *noopJobReporter) ReportRoomSessionID(v string)                     {}
func (r *noopJobReporter) ReportKind(v JobKind)                             {}
func (r *noopJobReporter) ReportWorkerKind(v WorkerKind)                    {}
func (r *noopJobReporter) ReportStatus(v JobStatus)                         {}
func (r *noopJobReporter) ReportDuration(v uint32)                          {}
func (r *noopJobReporter) ReportDurationSeconds(v uint32)                   {}
func (r *noopJobReporter) ReportDurationMinutes(v uint8)                    {}
func (r *noopJobReporter) ReportStartTime(v time.Time)                      {}
func (r *noopJobReporter) ReportEndTime(v time.Time)                        {}
func (r *noopJobReporter) ReportJoinLatency(v uint32)                       {}

type noopJobTx struct{}

func (t *noopJobTx) Worker() WorkerTx {
	return &noopWorkerTx{}
}

func (t *noopJobTx) ReportRoomSessionID(v string)   {}
func (t *noopJobTx) ReportKind(v JobKind)           {}
func (t *noopJobTx) ReportWorkerKind(v WorkerKind)  {}
func (t *noopJobTx) ReportStatus(v JobStatus)       {}
func (t *noopJobTx) ReportDuration(v uint32)        {}
func (t *noopJobTx) ReportDurationSeconds(v uint32) {}
func (t *noopJobTx) ReportDurationMinutes(v uint8)  {}
func (t *noopJobTx) ReportStartTime(v time.Time)    {}
func (t *noopJobTx) ReportEndTime(v time.Time)      {}
func (t *noopJobTx) ReportJoinLatency(v uint32)     {}
