package livekit

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/status"

	"github.com/livekit/protocol/utils/xtwirp"
)

var (
	_ xtwirp.ErrorMeta = (*SecureInputError)(nil)
	_ error            = (*SecureInputError)(nil)
)

func SecureInputErrorFrom(err error) *SecureInputError {
	if e, ok := errors.AsType[*SecureInputError](err); ok {
		return e
	}
	st, ok := status.FromError(err)
	if !ok {
		return nil
	}
	for _, d := range st.Details() {
		if e, ok := d.(*SecureInputError); ok {
			return e
		}
	}
	return nil
}

func (p *SecureInputError) Error() string {
	return fmt.Sprintf("secure input refused: %s", p.Reason)
}

func (p *SecureInputError) TwirpErrorMeta() map[string]string {
	return map[string]string{"secure_input_reason": p.Reason.String()}
}
