package rpc

import (
	"github.com/livekit/psrpc"
)

func NewAgentClient(bus psrpc.MessageBus) (AgentInternalClient, error) {
	return NewAgentInternalClient(bus)
}
