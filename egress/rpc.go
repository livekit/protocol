package egress

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
	SendRequest(ctx context.Context, req proto.Message) (*livekit.EgressInfo, error)
}

// RPCServer is used by Egress
type RPCServer interface {
	// GetRequestChannel returns a subscription for egress requests
	GetRequestChannel(ctx context.Context) (utils.PubSub, error)
	// ClaimRequest is used to take ownership of a request
	ClaimRequest(ctx context.Context, request *livekit.StartEgressRequest) (bool, error)
	// EgressSubscription subscribes to requests for a specific egress ID
	EgressSubscription(ctx context.Context, egressID string) (utils.PubSub, error)
	// SendResponse returns an RPC response
	SendResponse(ctx context.Context, request proto.Message, info *livekit.EgressInfo, err error) error
	// SendUpdate sends an egress info update
	SendUpdate(ctx context.Context, info *livekit.EgressInfo) error
}
