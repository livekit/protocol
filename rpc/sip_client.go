// Copyright 2023 LiveKit, Inc.
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

package rpc

import (
	"time"

	"github.com/livekit/psrpc"
)

// sipClaimTimeout is how long a SIP request waits for a SIP instance to claim it. psrpc's 1s
// default is too short for a claim published from a SIP pod whose Redis connections went cold
// while it sat idle. The window only bounds the wait for an unclaimed request; nothing is sent
// again, so a longer one cannot place a call twice.
const sipClaimTimeout = 3 * time.Second

type SIPClient interface {
	SIPInternalClient
}

type sipClient struct {
	SIPInternalClient
}

func NewSIPClient(bus psrpc.MessageBus) (SIPClient, error) {
	return NewSIPClientWithParams(ClientParams{Bus: bus})
}

func NewSIPClientWithParams(params ClientParams) (SIPClient, error) {
	if params.Bus == nil {
		return nil, nil
	}
	// First, so a caller's own selection timeout still wins.
	opts := append([]psrpc.ClientOption{psrpc.WithClientSelectTimeout(sipClaimTimeout)}, params.Options()...)

	internalClient, err := NewSIPInternalClient(params.Bus, opts...)
	if err != nil {
		return nil, err
	}

	return &sipClient{
		SIPInternalClient: internalClient,
	}, nil
}
