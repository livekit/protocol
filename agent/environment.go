package agent

import "fmt"

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
