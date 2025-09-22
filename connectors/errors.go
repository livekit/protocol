package connectors

import "github.com/livekit/psrpc"

var (
	ErrMissingWhatsAppPhoneNumberId = psrpc.NewErrorf(psrpc.InvalidArgument, "whatsapp phone number id is required")
	ErrMissingWhatsAppToNumber      = psrpc.NewErrorf(psrpc.InvalidArgument, "whatsapp to_phone_number is required")
	ErrMissingWhatsAppCallId        = psrpc.NewErrorf(psrpc.InvalidArgument, "whatsapp call id is required")
	ErrMissingWhatsAppApiKey        = psrpc.NewErrorf(psrpc.InvalidArgument, "whatsapp api key is required")
	ErrIncorrectSDPType             = psrpc.NewErrorf(psrpc.InvalidArgument, "incorrect sdp type")
)
