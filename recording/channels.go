package recording

import (
	"context"
	"errors"
	"time"

	"google.golang.org/protobuf/proto"

	livekit "github.com/livekit/protocol/proto"
	"github.com/livekit/protocol/utils"
)

const (
	ReservationChannel = "RESERVE_RECORDER"
	ResultChannel      = "RECORDING_RESULT"
	RequestTimeout     = time.Second * 3
	ReservationTimeout = time.Second * 2
)

func RequestChannel(id string) string {
	return "RECORDING_REQUEST_" + id
}

func ResponseChannel(id string) string {
	return "RECORDING_RESPONSE_" + id
}

func RecordingRPC(ctx context.Context, mb utils.MessageBus, recordingId string, req *livekit.RecordingRequest) error {
	sub, err := mb.Subscribe(ctx, ResponseChannel(recordingId))
	if err != nil {
		return err
	}

	b, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	err = mb.Publish(ctx, RequestChannel(recordingId), b)
	if err != nil {
		return err
	}

	for {
		select {
		case <-time.After(RequestTimeout):
			return errors.New("request timeout")
		case msg := <-sub.Channel():
			b = sub.Payload(msg)
			resp := &livekit.RecordingResponse{}
			err = proto.Unmarshal(b, resp)
			if err != nil {
				return err
			}
			if resp.RequestId != req.RequestId {
				continue
			}
			if !resp.Success {
				return errors.New(resp.Error)
			}
			return nil
		}
	}
}
