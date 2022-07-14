package egress

import "errors"

var (
	ErrNoResponse     = errors.New("no response from egress service")
	ErrEgressTimedOut = errors.New("egress timed out")
)
