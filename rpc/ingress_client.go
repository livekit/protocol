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
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/psrpc"
)

type IngressClient interface {
	IngressInternalClient
	IngressHandlerClient
}

type ingressClient struct {
	IngressInternalClient
	IngressHandlerClient
}

func NewIngressClient(nodeID livekit.NodeID, bus psrpc.MessageBus) (IngressClient, error) {
	if bus == nil {
		return nil, nil
	}

	clientID := string(nodeID)
	internalClient, err := NewIngressInternalClient(clientID, bus)
	if err != nil {
		return nil, err
	}
	handlerClient, err := NewIngressHandlerClient(clientID, bus)
	if err != nil {
		return nil, err
	}
	return &ingressClient{
		IngressInternalClient: internalClient,
		IngressHandlerClient:  handlerClient,
	}, nil
}
