package ingress

import "errors"

var (
	ErrIngressOutOfDate = errors.New("trying to ovewrite an ingress with an older version")
)
