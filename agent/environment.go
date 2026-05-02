package agent

import "fmt"

const MaxEnvironmentLength = 64

func ValidateEnvironment(env string) error {
	if env == "" {
		return nil
	}
	if len(env) > MaxEnvironmentLength {
		return fmt.Errorf("environment exceeds %d bytes", MaxEnvironmentLength)
	}
	for i := 0; i < len(env); i++ {
		c := env[i]
		switch {
		case c == '_':
			return fmt.Errorf("environment contains reserved character %q", c)
		case c <= ' ' || c == 0x7f:
			return fmt.Errorf("environment contains whitespace or control byte at position %d", i)
		}
	}
	return nil
}
