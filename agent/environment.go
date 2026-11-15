package agent

import (
	"fmt"

	"github.com/livekit/protocol/utils/guid"
)

const MaxDeploymentLength = 64

func ValidateDeployment(deployment string) error {
	if deployment == "" {
		return nil
	}
	if len(deployment) > MaxDeploymentLength {
		return fmt.Errorf("deployment exceeds %d bytes", MaxDeploymentLength)
	}
	for i := 0; i < len(deployment); i++ {
		c := deployment[i]
		switch {
		case c == '_':
			return fmt.Errorf("deployment contains reserved character %q", c)
		case c <= ' ' || c == 0x7f:
			return fmt.Errorf("deployment contains whitespace or control byte at position %d", i)
		}
	}
	return nil
}

// ValidateThreadID accepts an empty thread id (the call keeps no durable thread) or an agent
// thread id as guid.New(guid.AgentThreadPrefix) mints it.
func ValidateThreadID(threadID string) error {
	if threadID != "" && !guid.IsValidWithPrefix(guid.AgentThreadPrefix, threadID) {
		return fmt.Errorf("thread id must be an %s agent thread id", guid.AgentThreadPrefix)
	}
	return nil
}
