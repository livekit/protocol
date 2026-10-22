package xtwirp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/twitchtv/twirp"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/livekit/protocol/utils/xtwirp"
	"github.com/livekit/psrpc"
)

func TestStatus(t *testing.T) {
	st := status.New(codes.FailedPrecondition, "test")
	st, err := st.WithDetails(&errdetails.ErrorInfo{Reason: "reason"})
	require.NoError(t, err)

	e := twirp.NewError(twirp.InvalidArgument, "twirp")
	e = xtwirp.WithDetailsFromStatus(e, st)

	got, ok := xtwirp.StatusFromError(e)
	require.True(t, ok)
	require.Equal(t, st, got)
	require.Equal(t, st.Details(), got.Details())
}

// A non-twirp error is first wrapped as twirp.Unknown, then re-coded once the
// gRPC status is known. The re-coding must not fold the "twirp error unknown:"
// prefix into the message.
func TestToErrorKeepsMessage(t *testing.T) {
	cases := []struct {
		name string
		err  error
		msg  string
	}{
		{
			name: "psrpc",
			err:  psrpc.NewErrorFromResponse(string(psrpc.NotFound), "object cannot be found"),
			msg:  "object cannot be found",
		},
		{
			name: "grpc status",
			err:  status.Error(codes.NotFound, "object cannot be found"),
			msg:  "rpc error: code = NotFound desc = object cannot be found",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			e := xtwirp.ToError(c.err)
			require.Equal(t, twirp.NotFound, e.Code())
			require.Equal(t, c.msg, e.Msg())
		})
	}
}
