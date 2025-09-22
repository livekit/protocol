package utils

import (
	"fmt"
	"time"

	"github.com/livekit/psrpc"
)

type DataChannelRpcErrorCode uint32

const (
	DataChannelRpcApplicationError DataChannelRpcErrorCode = 1500 + iota
	DataChannelRpcConnectionTimeout
	DataChannelRpcResponseTimeout
	DataChannelRpcRecipientDisconnected
	DataChannelRpcResponsePayloadTooLarge
	DataChannelRpcSendFailed
)

const (
	DataChannelRpcUnsupportedMethod DataChannelRpcErrorCode = 1400 + iota
	DataChannelRpcRecipientNotFound
	DataChannelRpcRequestPayloadTooLarge
	DataChannelRpcUnsupportedServer
	DataChannelRpcUnsupportedVersion
)

const (
	DataChannelRpcMaxRoundTripLatency    = 2000 * time.Millisecond
	DataChannelRpcMaxMessageBytes        = 256
	DataChannelRpcMaxDataBytes           = 15360 // 15KiB
	DataChannelRpcMaxPayloadBytes        = 15360 // 15KiB
	DataChannelRpcDefaultResponseTimeout = 10000 * time.Millisecond
)

var dataChannelRpcErrorMessages = map[DataChannelRpcErrorCode]string{
	DataChannelRpcApplicationError:        "Application error in method handler",
	DataChannelRpcConnectionTimeout:       "Connection timeout",
	DataChannelRpcResponseTimeout:         "Response timeout",
	DataChannelRpcRecipientDisconnected:   "Recipient disconnected",
	DataChannelRpcResponsePayloadTooLarge: "Response payload too large",
	DataChannelRpcSendFailed:              "Failed to send",

	DataChannelRpcUnsupportedMethod:      "Method not supported at destination",
	DataChannelRpcRecipientNotFound:      "Recipient not found",
	DataChannelRpcRequestPayloadTooLarge: "Request payload too large",
	DataChannelRpcUnsupportedServer:      "RPC not supported by server",
	DataChannelRpcUnsupportedVersion:     "Unsupported RPC version",
}

type DataChannelRpcError struct {
	Code    DataChannelRpcErrorCode
	Message string
	Data    string
}

func (e *DataChannelRpcError) Error() string {
	return fmt.Sprintf("RpcError %d: %s", e.Code, e.Message)
}

func (e *DataChannelRpcError) PsrpcError() psrpc.Error {
	switch e.Code {
	case DataChannelRpcApplicationError:
		return psrpc.NewErrorf(psrpc.Internal, e.Message, "data", e.Data)
	case DataChannelRpcConnectionTimeout:
		return psrpc.NewErrorf(psrpc.Canceled, e.Message, "data", e.Data)
	case DataChannelRpcResponseTimeout:
		return psrpc.NewErrorf(psrpc.Canceled, e.Message, "data", e.Data)
	case DataChannelRpcRecipientDisconnected:
		return psrpc.NewErrorf(psrpc.Unavailable, e.Message, "data", e.Data)
	case DataChannelRpcResponsePayloadTooLarge:
		return psrpc.NewErrorf(psrpc.MalformedResponse, e.Message, "data", e.Data)
	case DataChannelRpcSendFailed:
		return psrpc.NewErrorf(psrpc.Internal, e.Message, "data", e.Data)
	case DataChannelRpcUnsupportedMethod:
		return psrpc.NewErrorf(psrpc.InvalidArgument, e.Message, "data", e.Data)
	case DataChannelRpcRecipientNotFound:
		return psrpc.NewErrorf(psrpc.NotFound, e.Message, "data", e.Data)
	case DataChannelRpcRequestPayloadTooLarge:
		return psrpc.NewErrorf(psrpc.MalformedRequest, e.Message, "data", e.Data)
	case DataChannelRpcUnsupportedServer:
		return psrpc.NewErrorf(psrpc.Unimplemented, e.Message, "data", e.Data)
	case DataChannelRpcUnsupportedVersion:
		return psrpc.NewErrorf(psrpc.Unimplemented, e.Message, "data", e.Data)
	default:
		return psrpc.NewErrorf(psrpc.Internal, e.Message, "data", e.Data)
	}
}

func DataChannelRpcErrorFromBuiltInCodes(code DataChannelRpcErrorCode, data string) *DataChannelRpcError {
	return &DataChannelRpcError{Code: code, Message: dataChannelRpcErrorMessages[code], Data: data}
}

type DataChannelRpcPendingAckHandler struct {
	Resolve             func()
	ParticipantIdentity string
}

type DataChannelRpcPendingResponseHandler struct {
	Resolve             func(payload string, err *DataChannelRpcError)
	ParticipantIdentity string
}
