package ingress

import (
	"fmt"

	"github.com/livekit/psrpc"
)

var (
	ErrIngressOutOfDate        = psrpc.NewErrorf(psrpc.FailedPrecondition, "trying to ovewrite an ingress with an older version")
	ErrIngressTimedOut         = psrpc.NewErrorf(psrpc.DeadlineExceeded, "ingress timed out")
	ErrNoResponse              = psrpc.NewErrorf(psrpc.Unavailable, "no response from ingress service")
	ErrInvalidOutputDimensions = NewInvalidVideoParamsError("invalid output media dimensions")
)

func NewInvalidVideoParamsError(s string) error {
	return psrpc.NewError(psrpc.InvalidArgument, fmt.Errorf("invalid video parameters: %s", s))
}

func NewInvalidAudioParamsError(s string) error {
	return psrpc.NewError(psrpc.InvalidArgument, fmt.Errorf("invalid audio parameters: %s", s))
}
