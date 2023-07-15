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

func NewEgressClient(nodeID livekit.NodeID, bus psrpc.MessageBus) (EgressClient, error) {
	if bus == nil {
		return nil, nil
	}

	clientID := string(nodeID)
	internalClient, err := NewEgressInternalClient(clientID, bus, middleware.WithRPCRetries(middleware.RetryOptions{
		MaxAttempts: retries,
		Timeout:     psrpc.DefaultClientTimeout,
		IsRecoverable: func(err error) bool {
			var e psrpc.Error
			if !errors.As(err, &e) {
				return true
			}
			return e.Code() == psrpc.DeadlineExceeded || e.Code() == psrpc.ResourceExhausted
		},
	}))
	if err != nil {
		return nil, err
	}

	handlerClient, err := NewEgressHandlerClient(clientID, bus)
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
