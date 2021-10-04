package recording

import "time"

const (
	ReservationChannel = "RESERVE_RECORDER"
	ResultChannel      = "RECORDING_RESULT"
	RecorderTimeout    = time.Second * 3
	ReservationTimeout = time.Second * 2
)

func RequestChannel(id string) string {
	return "RECORDING_REQUEST_" + id
}

func ResponseChannel(id string) string {
	return "RECORDING_RESPONSE_" + id
}
