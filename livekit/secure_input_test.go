package livekit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/utils/xtwirp"
	"github.com/livekit/psrpc"
)

func TestSecureInputErrorFrom(t *testing.T) {
	want := &SecureInputError{Reason: SecureInputErrorReason_SECURE_INPUT_ERROR_NO_TELEPHONE_EVENTS}
	twerr := xtwirp.ToError(psrpc.NewError(psrpc.FailedPrecondition, errors.New("refused"), want))
	require.Equal(t, "SECURE_INPUT_ERROR_NO_TELEPHONE_EVENTS", twerr.Meta("secure_input_reason"))
	st, ok := xtwirp.StatusFromError(twerr)
	require.True(t, ok)
	require.True(t, proto.Equal(want, SecureInputErrorFrom(st.Err())))
}
