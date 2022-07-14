package egress

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils"
)

const (
	newEgressChannel      = "EG_START"
	updateChannel         = "EG_RESULTS"
	requestChannelPrefix  = "REQ_"
	responseChannelPrefix = "RES_"

	RequestExpiration = time.Second * 2
	requestTimeout    = time.Second * 3
	lockDuration      = time.Second * 3
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

type RedisRPC struct {
	nodeID livekit.NodeID
	bus    *utils.RedisMessageBus
}

func NewRedisRPCClient(nodeID livekit.NodeID, rc *redis.Client) RPCClient {
	if rc == nil {
		return nil
	}

	bus := utils.NewRedisMessageBus(rc)
	return &RedisRPC{
		nodeID: nodeID,
		bus:    bus.(*utils.RedisMessageBus),
	}
}

func (r *RedisRPC) GetUpdateChannel(ctx context.Context) (utils.PubSub, error) {
	return r.bus.SubscribeQueue(context.Background(), updateChannel)
}

func (r *RedisRPC) SendRequest(ctx context.Context, request proto.Message) (*livekit.EgressInfo, error) {
	requestID := utils.NewGuid(utils.RPCPrefix)
	var channel string

	switch req := request.(type) {
	case *livekit.StartEgressRequest:
		if req.EgressId == "" {
			req.EgressId = utils.NewGuid(utils.EgressPrefix)
		}
		req.RequestId = requestID
		req.SentAt = time.Now().UnixNano()
		req.SenderId = string(r.nodeID)
		channel = newEgressChannel

	case *livekit.EgressRequest:
		req.RequestId = requestID
		req.SenderId = string(r.nodeID)
		channel = requestChannel(req.EgressId)

	default:
		return nil, errors.New("invalid request type")
	}

	sub, err := r.bus.Subscribe(ctx, responseChannel(requestID))
	if err != nil {
		return nil, err
	}
	defer func() {
		err := sub.Close()
		if err != nil {
			logger.Errorw("failed to unsubscribe from response channel", err)
		}
	}()

	err = r.bus.Publish(ctx, channel, request)
	if err != nil {
		return nil, err
	}

	select {
	case raw := <-sub.Channel():
		res := &livekit.EgressResponse{}
		err := proto.Unmarshal(sub.Payload(raw), res)
		if err != nil {
			return nil, err
		} else if res.Error != "" {
			return nil, errors.New(res.Error)
		} else {
			return res.Info, nil
		}

	case <-time.After(requestTimeout):
		return nil, ErrNoResponse
	}
}

func NewRedisRPCServer(rc *redis.Client) RPCServer {
	bus := utils.NewRedisMessageBus(rc)
	return &RedisRPC{
		bus: bus.(*utils.RedisMessageBus),
	}
}

func (r *RedisRPC) GetRequestChannel(ctx context.Context) (utils.PubSub, error) {
	return r.bus.Subscribe(ctx, newEgressChannel)
}

func (r *RedisRPC) ClaimRequest(ctx context.Context, req *livekit.StartEgressRequest) (bool, error) {
	claimed, err := r.bus.Lock(ctx, requestChannel(req.EgressId), lockDuration)
	if err != nil || !claimed {
		return false, err
	}
	return true, nil
}

func (r *RedisRPC) EgressSubscription(ctx context.Context, egressID string) (utils.PubSub, error) {
	return r.bus.Subscribe(ctx, requestChannel(egressID))
}

func (r *RedisRPC) SendResponse(ctx context.Context, request proto.Message, info *livekit.EgressInfo, err error) error {
	res := &livekit.EgressResponse{
		Info: info,
	}

	switch req := request.(type) {
	case *livekit.StartEgressRequest:
		res.RequestId = req.RequestId
	case *livekit.EgressRequest:
		res.RequestId = req.RequestId
	}

	if err != nil {
		res.Error = err.Error()
	}

	return r.bus.Publish(ctx, responseChannel(res.RequestId), res)
}

func (r *RedisRPC) SendUpdate(ctx context.Context, info *livekit.EgressInfo) error {
	return r.bus.Publish(ctx, updateChannel, info)
}

func requestChannel(egressID string) string {
	return requestChannelPrefix + egressID
}

func responseChannel(nodeID string) string {
	return responseChannelPrefix + nodeID
}
