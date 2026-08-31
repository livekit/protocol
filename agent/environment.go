package agent

import "fmt"

const MaxDeploymentLength = 64

const MaxAgentNameLength = 64

// ValidateAgentName enforces that an agent name is usable as a single URL path
// segment, so a worker can be addressed at /agents/{agent_name}/{deployment}/...
// Unlike a deployment, an agent name commonly contains underscores, so they are
// permitted here.
func ValidateAgentName(agentName string) error {
	if len(agentName) > MaxAgentNameLength {
		return fmt.Errorf("agent name exceeds %d bytes", MaxAgentNameLength)
	}
	for i := 0; i < len(agentName); i++ {
		c := agentName[i]
		switch {
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9', c == '-', c == '.', c == '_':
		default:
			return fmt.Errorf("agent name contains invalid character %q at position %d", c, i)
		}
	}
	return nil
}

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
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9', c == '-', c == '.':
		case c == '_':
			return fmt.Errorf("deployment contains reserved character %q", c)
		default:
			// deployments appear as a single URL path segment (/agents/{deployment}/...),
			// so only alphanumerics, '-' and '.' are accepted
			return fmt.Errorf("deployment contains invalid character %q at position %d", c, i)
		}
	}
	return nil
}
