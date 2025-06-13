// Code generated; DO NOT EDIT.
package agentsobs

type WorkerStatus string

const (
	WorkerStatusUndefined WorkerStatus = ""
	WorkerStatusAvailable WorkerStatus = "available"
	WorkerStatusFull      WorkerStatus = "full"
)

type AgentKind string

const (
	AgentKindUndefined AgentKind = ""
	AgentKindCloud     AgentKind = "cloud"
	AgentKindSelfhost  AgentKind = "selfhost"
)

type JobKind string

const (
	JobKindUndefined   JobKind = ""
	JobKindRoom        JobKind = "room"
	JobKindPublisher   JobKind = "publisher"
	JobKindParticipant JobKind = "participant"
)

type JobStatus string

const (
	JobStatusUndefined JobStatus = ""
	JobStatusPending   JobStatus = "pending"
	JobStatusRunning   JobStatus = "running"
	JobStatusSuccess   JobStatus = "success"
	JobStatusFailed    JobStatus = "failed"
)

type Rollup string

const (
	RollupUndefined    Rollup = ""
	RollupWorker       Rollup = "worker"
	RollupWorkerSeries Rollup = "worker_series"
	RollupJob          Rollup = "job"
)
