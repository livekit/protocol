package ingress

import "errors"

var (
	ErrIngressOutOfDate        = errors.New("trying to ovewrite an ingress with an older version")
	ErrIngressTimedOut         = errors.New("ingress timed out")
	ErrNoResponse              = errors.New("no response from ingress service")
	ErrInvalidOutputDimensions = NewInvalidVideoParamsError("invalid output media dimensions")
)

type InvalidVideoParamsError string

func NewInvalidVideoParamsError(s string) InvalidVideoParamsError {
	return InvalidVideoParamsError(s)
}

func (s InvalidVideoParamsError) Error() string {
	return "invalid video parameters: " + string(s)
}
