package ingress

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils"
)

// RPCClient is used by LiveKit Server
type RPCClient interface {
	// GetUpdateChannel returns a subscription for egress info updates
	GetUpdateChannel(ctx context.Context) (utils.PubSub, error)
	// SendRequest sends a request to all available instances
	SendRequest(ctx context.Context, req proto.Message) (*livekit.IngressInfo, error)
}

// RPCServer is used by Ingress
type RPCServer interface {
	// GetRequestChannel returns a subscription for ingress requests
	GetRequestChannel(ctx context.Context) (utils.PubSub, error)
	// ClaimRequest is used to take ownership of a request
	ClaimRequest(ctx context.Context, request *livekit.StartIngressRequest) (bool, error)
	// EgressSubscription subscribes to requests for a specific egress ID
	IngressSubscription(ctx context.Context, ingressID string) (utils.PubSub, error)
	// SendResponse returns an RPC response
	SendResponse(ctx context.Context, request proto.Message, info *livekit.IngressInfo, err error) error
	// SendUpdate sends an egress info update
	SendUpdate(ctx context.Context, info *livekit.IngressInfo) error
}
