package egress

import (
	"context"
	"errors"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils"
)

const (
	startChannel          = "EG_START"
	resultsChannel        = "EG_RESULTS"
	requestChannelPrefix  = "REQ_"
	responseChannelPrefix = "RES_"
	requestTimeout        = time.Second * 3
)

func SubscribeToStartRequests(ctx context.Context, bus utils.MessageBus) (utils.PubSub, error) {
	return bus.Subscribe(ctx, startChannel)
}

func SubscribeToEgressRequests(ctx context.Context, bus utils.MessageBus, egressID string) (utils.PubSub, error) {
	return bus.Subscribe(ctx, requestChannel(egressID))
}

func SendRequest(ctx context.Context, bus utils.MessageBus, req proto.Message) (*livekit.EgressInfo, error) {
	requestID := utils.NewGuid(utils.RPCPrefix)
	var channel string

	switch r := req.(type) {
	case *livekit.StartEgressRequest:
		r.EgressId = utils.NewGuid(utils.EgressPrefix)
		r.RequestId = requestID
		channel = startChannel
	case *livekit.EgressRequest:
		r.RequestId = requestID
		channel = requestChannel(r.EgressId)
	default:
		return nil, errors.New("invalid request type")
	}

	sub, err := bus.Subscribe(ctx, responseChannel(requestID))
	if err != nil {
		return nil, err
	}
	defer func() {
		err := sub.Close()
		if err != nil {
			logger.Errorw("failed to unsubscribe from response channel", err)
		}
	}()

	err = bus.Publish(ctx, channel, req)
	if err != nil {
		return nil, err
	}

	select {
	case raw := <-sub.Channel():
		return unmarshalResponse(sub.Payload(raw))
	case <-time.After(requestTimeout):
		return nil, errors.New("no response from egress service")
	}
}

func SendResponse(ctx context.Context, bus utils.MessageBus, requestID string, res *livekit.EgressResponse) error {
	return bus.Publish(ctx, responseChannel(requestID), res)
}

func SubscribeToResults(ctx context.Context, bus utils.MessageBus) (utils.PubSub, error) {
	return bus.Subscribe(ctx, resultsChannel)
}

func requestChannel(egressID string) string {
	return requestChannelPrefix + egressID
}

func responseChannel(requestID string) string {
	return responseChannelPrefix + requestID
}

func unmarshalResponse(data []byte) (*livekit.EgressInfo, error) {
	res := &livekit.EgressResponse{}
	err := proto.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return res.Info, nil
}
