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
	"context"
	"errors"
	"time"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/middleware"
)

const retries = 3

type EgressClient interface {
	EgressInternalClient
	EgressHandlerClient
}

type egressClient struct {
	EgressInternalClient
	EgressHandlerClient
}

func NewEgressClient(params ClientParams) (EgressClient, error) {
	if params.Bus == nil {
		return nil, nil
	}
	opts := clientOptions(params)
	timeout := params.Timeout
	if timeout < 10*time.Second {
		timeout = 10 * time.Second
	}
	internalOpts := append(opts, middleware.WithRPCRetries(middleware.RetryOptions{
		MaxAttempts: params.MaxAttempts,
		// use longer retry timeout
		Timeout: timeout,
		IsRecoverable: func(err error) bool {
			var e psrpc.Error
			if !errors.As(err, &e) {
				return true
			}
			return e.Code() == psrpc.DeadlineExceeded ||
				e.Code() == psrpc.ResourceExhausted ||
				e.Code() == psrpc.Unavailable
		},
	}))
	internalClient, err := NewEgressInternalClient(params.Bus, internalOpts...)
	if err != nil {
		return nil, err
	}

	handlerClient, err := NewEgressHandlerClient(params.Bus, opts...)
	if err != nil {
		return nil, err
	}

	return &egressClient{
		EgressInternalClient: internalClient,
		EgressHandlerClient:  handlerClient,
	}, nil
}

func (c *egressClient) StartEgress(ctx context.Context, topic string, req *StartEgressRequest, opts ...psrpc.RequestOption) (*livekit.EgressInfo, error) {
	o := append([]psrpc.RequestOption{
		psrpc.WithSelectionOpts(psrpc.SelectionOpts{
			MaximumAffinity:     1,
			AffinityTimeout:     time.Second,
			ShortCircuitTimeout: time.Millisecond * 500,
		}),
	}, opts...)
	return c.EgressInternalClient.StartEgress(ctx, topic, req, o...)
}
