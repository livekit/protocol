package livekit

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestAgentDBErrorCodes(t *testing.T) {
	t.Run("a code reads back from its wire name", func(t *testing.T) {
		for n := range AgentDB_ErrorCode_name {
			code := AgentDB_ErrorCode(n)
			require.Equal(t, code, ParseAgentDBErrorCode(code.Wire()))
		}
	})

	t.Run("an unknown code may have applied and is never retried", func(t *testing.T) {
		for _, code := range []AgentDB_ErrorCode{ParseAgentDBErrorCode("no_such_code"), AgentDB_ErrorCode(999)} {
			require.Equal(t, AgentDB_UNKNOWN, ParseAgentDBErrorCode(code.Wire()))
			require.Equal(t, AgentDB_APPLIED_MAYBE, code.Applied())
			require.Equal(t, AgentDB_RETRY_NEVER, code.Retry())
		}
	})

	t.Run("each code carries its options", func(t *testing.T) {
		require.Equal(t, AgentDB_APPLIED_MAYBE, AgentDB_OUTCOME_UNKNOWN.Applied())
		require.Equal(t, AgentDB_RETRY_HERE, AgentDB_OUTCOME_UNKNOWN.Retry())
		require.Equal(t, AgentDB_APPLIED_NO, AgentDB_IDLE_EVICTED.Applied())
		require.Equal(t, AgentDB_RETRY_ELSEWHERE, AgentDB_IDLE_EVICTED.Retry())
		require.Equal(t, AgentDB_APPLIED_MAYBE, AgentDB_INTERNAL.Applied())
		require.Equal(t, AgentDB_RETRY_NEVER, AgentDB_INTERNAL.Retry())
	})

	t.Run("a gRPC status code is the code of the same name", func(t *testing.T) {
		for c := codes.Canceled; c <= codes.Unauthenticated; c++ {
			want := AgentDB_UNKNOWN
			if c != codes.Unknown {
				want = ParseAgentDBErrorCode(map[codes.Code]string{
					codes.Canceled: "canceled", codes.InvalidArgument: "invalid_argument",
					codes.DeadlineExceeded: "deadline_exceeded", codes.NotFound: "not_found",
					codes.AlreadyExists: "already_exists", codes.PermissionDenied: "permission_denied",
					codes.ResourceExhausted: "resource_exhausted", codes.FailedPrecondition: "failed_precondition",
					codes.Aborted: "aborted", codes.OutOfRange: "out_of_range", codes.Unimplemented: "unimplemented",
					codes.Internal: "internal", codes.Unavailable: "unavailable", codes.DataLoss: "data_loss",
					codes.Unauthenticated: "unauthenticated",
				}[c])
				require.NotEqual(t, AgentDB_UNKNOWN, want, c)
			}
			require.Equal(t, want, AgentDBErrorCodeOf(c), c)
		}
		require.Equal(t, AgentDB_UNKNOWN, AgentDBErrorCodeOf(codes.OK))
	})
}
