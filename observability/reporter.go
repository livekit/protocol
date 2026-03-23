package observability

import (
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/observability/agentsobs"
	"github.com/livekit/protocol/observability/core_callobs"
	"github.com/livekit/protocol/observability/egressobs"
	"github.com/livekit/protocol/observability/gatewayobs"
	"github.com/livekit/protocol/observability/ingressobs"
	"github.com/livekit/protocol/observability/roomobs"
	"github.com/livekit/protocol/observability/storageobs"
	"github.com/livekit/protocol/observability/telephony_callobs"
	"github.com/livekit/protocol/observability/telephonyobs"
)

const Project = "livekit"

type Reporter interface {
	Logger(name, projectID string) (logger.Logger, error)
	Room() roomobs.Reporter
	Agent() agentsobs.Reporter
	Gateway() gatewayobs.Reporter
	Telephony() telephonyobs.Reporter
	Egress() egressobs.Reporter
	Ingress() ingressobs.Reporter
	TelephonyCall() telephony_callobs.Reporter
	CoreCall() core_callobs.Reporter
	Storage() storageobs.Reporter
	Close()
}

func NewReporter() Reporter {
	return reporter{}
}

type reporter struct{}

func (reporter) Logger(name, projectID string) (logger.Logger, error) {
	return logger.GetDiscardLogger(), nil
}

func (reporter) Room() roomobs.Reporter {
	return roomobs.NewNoopReporter()
}

func (reporter) Agent() agentsobs.Reporter {
	return agentsobs.NewNoopReporter()
}

func (reporter) Gateway() gatewayobs.Reporter {
	return gatewayobs.NewNoopReporter()
}

func (reporter) Telephony() telephonyobs.Reporter {
	return telephonyobs.NewNoopReporter()
}

func (reporter) Egress() egressobs.Reporter {
	return egressobs.NewNoopReporter()
}

func (reporter) Ingress() ingressobs.Reporter {
	return ingressobs.NewNoopReporter()
}

func (reporter) TelephonyCall() telephony_callobs.Reporter {
	return telephony_callobs.NewNoopReporter()
}

func (reporter) CoreCall() core_callobs.Reporter {
	return core_callobs.NewNoopReporter()
}

func (reporter) Storage() storageobs.Reporter { return storageobs.NewNoopReporter() }

func (reporter) Close() {
}
