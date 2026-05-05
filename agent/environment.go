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
		case c == '_':
			return fmt.Errorf("deployment contains reserved character %q", c)
		case c <= ' ' || c == 0x7f:
			return fmt.Errorf("deployment contains whitespace or control byte at position %d", i)
		}
	}
	return nil
}
