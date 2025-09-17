package utils

import (
	"fmt"
	"time"

	"github.com/livekit/psrpc"
)

type RpcErrorCode uint32

const (
	RpcApplicationError RpcErrorCode = 1500 + iota
	RpcConnectionTimeout
	RpcResponseTimeout
	RpcRecipientDisconnected
	RpcResponsePayloadTooLarge
	RpcSendFailed
)

const (
	RpcUnsupportedMethod RpcErrorCode = 1400 + iota
	RpcRecipientNotFound
	RpcRequestPayloadTooLarge
	RpcUnsupportedServer
	RpcUnsupportedVersion
)

const (
	RpcMaxRoundTripLatency    = 2000 * time.Millisecond
	RpcMaxMessageBytes        = 256
	RpcMaxDataBytes           = 15360 // 15KiB
	RpcMaxPayloadBytes        = 15360 // 15KiB
	RpcDefaultResponseTimeout = 10000 * time.Millisecond
)

var rpcErrorMessages = map[RpcErrorCode]string{
	RpcApplicationError:        "Application error in method handler",
	RpcConnectionTimeout:       "Connection timeout",
	RpcResponseTimeout:         "Response timeout",
	RpcRecipientDisconnected:   "Recipient disconnected",
	RpcResponsePayloadTooLarge: "Response payload too large",
	RpcSendFailed:              "Failed to send",

	RpcUnsupportedMethod:      "Method not supported at destination",
	RpcRecipientNotFound:      "Recipient not found",
	RpcRequestPayloadTooLarge: "Request payload too large",
	RpcUnsupportedServer:      "RPC not supported by server",
	RpcUnsupportedVersion:     "Unsupported RPC version",
}

type RpcError struct {
	Code    RpcErrorCode
	Message string
	Data    string
}

func (e *RpcError) Error() string {
	return fmt.Sprintf("RpcError %d: %s", e.Code, e.Message)
}

func (e *RpcError) PsrpcError() psrpc.Error {
	switch e.Code {
	case RpcApplicationError:
		return psrpc.NewErrorf(psrpc.Internal, e.Message, "data", e.Data)
	case RpcConnectionTimeout:
		return psrpc.NewErrorf(psrpc.DeadlineExceeded, e.Message, "data", e.Data)
	case RpcResponseTimeout:
		return psrpc.NewErrorf(psrpc.DeadlineExceeded, e.Message, "data", e.Data)
	case RpcRecipientDisconnected:
		return psrpc.NewErrorf(psrpc.Unavailable, e.Message, "data", e.Data)
	case RpcResponsePayloadTooLarge:
		return psrpc.NewErrorf(psrpc.MalformedResponse, e.Message, "data", e.Data)
	case RpcSendFailed:
		return psrpc.NewErrorf(psrpc.Internal, e.Message, "data", e.Data)
	case RpcUnsupportedMethod:
		return psrpc.NewErrorf(psrpc.InvalidArgument, e.Message, "data", e.Data)
	case RpcRecipientNotFound:
		return psrpc.NewErrorf(psrpc.NotFound, e.Message, "data", e.Data)
	case RpcRequestPayloadTooLarge:
		return psrpc.NewErrorf(psrpc.MalformedRequest, e.Message, "data", e.Data)
	case RpcUnsupportedServer:
		return psrpc.NewErrorf(psrpc.Unimplemented, e.Message, "data", e.Data)
	case RpcUnsupportedVersion:
		return psrpc.NewErrorf(psrpc.Unimplemented, e.Message, "data", e.Data)
	default:
		return psrpc.NewErrorf(psrpc.Internal, e.Message, "data", e.Data)
	}
}

func RpcErrorFromBuiltInCodes(code RpcErrorCode, data *string) *RpcError {
	return &RpcError{Code: code, Message: rpcErrorMessages[code], Data: *data}
}

type RpcPendingAckHandler struct {
	Resolve             func()
	ParticipantIdentity string
}

type RpcPendingResponseHandler struct {
	Resolve             func(payload string, err *RpcError)
	ParticipantIdentity string
}
