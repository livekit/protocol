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
	StartChannel          = "EG_START"
	ResultsChannel        = "EG_RESULTS"
	requestChannelPrefix  = "REQ_"
	responseChannelPrefix = "RES_"
	LockDuration          = time.Second * 3
	requestTimeout        = time.Second * 3
)

func SendRequest(ctx context.Context, bus utils.MessageBus, req proto.Message) (*livekit.EgressInfo, error) {
	requestID := utils.NewGuid(utils.RPCPrefix)
	var channel string

	switch r := req.(type) {
	case *livekit.StartEgressRequest:
		r.EgressId = utils.NewGuid(utils.EgressPrefix)
		r.RequestId = requestID
		channel = StartChannel
	case *livekit.EgressRequest:
		r.RequestId = requestID
		channel = RequestChannel(r.EgressId)
	default:
		return nil, errors.New("invalid request type")
	}

	sub, err := bus.Subscribe(ctx, ResponseChannel(requestID))
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

func RequestChannel(egressID string) string {
	return requestChannelPrefix + egressID
}

func ResponseChannel(requestID string) string {
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
