package observability

import "github.com/livekit/protocol/observability/roomobs"

const Project = "livekit"

type Reporter interface {
	Room() roomobs.Reporter
	Close()
}

func NewReporter() Reporter {
	return reporter{}
}

type reporter struct{}

func (reporter) Room() roomobs.Reporter {
	return roomobs.NewNoopReporter()
}

func (reporter) Close() {}
