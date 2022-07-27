package ingress

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
	newIngressChannel     = "IN_START"
	updateChannel         = "IN_RESULTS"
	entityChannel         = "IN_ENTITY"
	requestChannelPrefix  = "REQ_"
	responseChannelPrefix = "RES_"

	RequestExpiration = time.Second * 2
	requestTimeout    = time.Second * 3
	lockDuration      = time.Second * 3
)

// RPCClient is used by LiveKit Server
type RPCClient interface {
	// GetUpdateChannel returns a subscription for ingress info updates
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
	// IngressSubscription subscribes to requests for a specific ingress ID
	IngressSubscription(ctx context.Context, ingressID string) (utils.PubSub, error)
	// SendResponse returns an RPC response
	SendResponse(ctx context.Context, request proto.Message, info *livekit.IngressInfo, err error) error
	// SendUpdate sends an ingress info update
	SendUpdate(ctx context.Context, info *livekit.IngressInfo) error
	// GetEntityChannel returns a subscription for entity requests
	GetEntityChannel(ctx context.Context) (utils.PubSub, error)
}

type RPC interface {
	RPCClient
	RPCServer
}

type RedisRPC struct {
	nodeID livekit.NodeID
	bus    *utils.RedisMessageBus
}

func NewRedisRPC(nodeID livekit.NodeID, rc *redis.Client) RPC {
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

func (r *RedisRPC) SendRequest(ctx context.Context, request proto.Message) (*livekit.IngressInfo, error) {
	requestID := utils.NewGuid(utils.RPCPrefix)
	var channel string

	switch req := request.(type) {
	case *livekit.StartIngressRequest:
		req.IngressId = utils.NewGuid(utils.IngressPrefix)
		req.RequestId = requestID
		req.SentAt = time.Now().UnixNano()
		req.SenderId = string(r.nodeID)
		channel = newIngressChannel

	case *livekit.IngressRequest:
		req.RequestId = requestID
		req.SenderId = string(r.nodeID)
		channel = requestChannel(req.IngressId)

	case *livekit.GetIngressInfoRequest:
		req.RequestId = requestID
		req.SenderId = string(r.nodeID)
		req.SentAt = time.Now().UnixNano()
		channel = entityChannel

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
		res := &livekit.IngressResponse{}
		err := proto.Unmarshal(sub.Payload(raw), res)
		if err != nil {
			return nil, err
		} else if res.Error != "" {
			return nil, errors.New(res.Error)
		} else {
			return res.Info, nil
		}

	case <-time.After(requestTimeout):
		return nil, errors.New("no response from ingress service")
	}
}

func (r *RedisRPC) GetRequestChannel(ctx context.Context) (utils.PubSub, error) {
	return r.bus.Subscribe(ctx, newIngressChannel)
}

func (r *RedisRPC) ClaimRequest(ctx context.Context, req *livekit.StartIngressRequest) (bool, error) {
	claimed, err := r.bus.Lock(ctx, requestChannel(req.IngressId), lockDuration)
	if err != nil || !claimed {
		return false, err
	}
	return true, nil
}

func (r *RedisRPC) IngressSubscription(ctx context.Context, ingressID string) (utils.PubSub, error) {
	return r.bus.Subscribe(ctx, requestChannel(ingressID))
}

func (r *RedisRPC) SendResponse(ctx context.Context, request proto.Message, info *livekit.IngressInfo, err error) error {
	res := &livekit.IngressResponse{
		Info: info,
	}

	switch req := request.(type) {
	case *livekit.StartIngressRequest:
		res.RequestId = req.RequestId
	case *livekit.IngressRequest:
		res.RequestId = req.RequestId
	case *livekit.GetIngressInfoRequest:
		res.RequestId = req.RequestId
	}

	if err != nil {
		res.Error = err.Error()
	}

	return r.bus.Publish(ctx, responseChannel(res.RequestId), res)
}

func (r *RedisRPC) SendUpdate(ctx context.Context, info *livekit.IngressInfo) error {
	return r.bus.Publish(ctx, updateChannel, info)
}

func (r *RedisRPC) GetEntityChannel(ctx context.Context) (utils.PubSub, error) {
	return r.bus.Subscribe(ctx, entityChannel)
}

func requestChannel(ingressID string) string {
	return requestChannelPrefix + ingressID
}

func responseChannel(nodeID string) string {
	return responseChannelPrefix + nodeID
}
