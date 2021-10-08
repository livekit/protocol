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

func ReserveRecorder(bus utils.MessageBus) (string, error) {
	recordingId := utils.NewGuid(utils.RecordingPrefix)
	req := &livekit.RecordingReservation{
		Id:          recordingId,
		SubmittedAt: time.Now().UnixNano() / 1e6,
	}

	b, err := proto.Marshal(req)
	if err != nil {
		return "", err
	}

	sub, _ := bus.Subscribe(context.Background(), ResponseChannel(recordingId))
	defer sub.Close()

	err = bus.Publish(context.Background(), ReservationChannel, string(b))
	if err != nil {
		return "", err
	}

	select {
	case <-sub.Channel():
		return recordingId, nil
	case <-time.After(RequestTimeout):
		return "", errors.New("no recorders available")
	}
}

func RPC(ctx context.Context, bus utils.MessageBus, recordingId string, req *livekit.RecordingRequest) error {
	sub, err := bus.Subscribe(ctx, ResponseChannel(recordingId))
	if err != nil {
		return err
	}
	defer sub.Close()

	if req.RequestId == "" {
		req.RequestId = utils.NewGuid(utils.RPCPrefix)
	}

	b, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	err = bus.Publish(ctx, RequestChannel(recordingId), b)
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
			if resp.Error != "" {
				return errors.New(resp.Error)
			}
			return nil
		}
	}
}
