// Copyright 2025 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connector

import "github.com/livekit/psrpc"

var (
	ErrMissingWhatsAppPhoneNumberId = psrpc.NewErrorf(psrpc.InvalidArgument, "whatsapp phone number id is required")
	ErrMissingWhatsAppToNumber      = psrpc.NewErrorf(psrpc.InvalidArgument, "whatsapp to_phone_number is required")
	ErrMissingWhatsAppCallId        = psrpc.NewErrorf(psrpc.InvalidArgument, "whatsapp call id is required")
	ErrMissingWhatsAppApiKey        = psrpc.NewErrorf(psrpc.InvalidArgument, "whatsapp api key is required")
	ErrIncorrectSDPType             = psrpc.NewErrorf(psrpc.InvalidArgument, "incorrect sdp type")
	ErrConnectorNotFound            = psrpc.NewErrorf(psrpc.NotFound, "connector not found")
	ErrWhatsAppCallNotFound         = psrpc.NewErrorf(psrpc.NotFound, "whatsapp call not found")
)
