package utils

import "time"

const (
	ReservationChannel = "RESERVE_RECORDER"
	RecorderTimeout    = time.Second * 3
	ReservationTimeout = time.Second * 2
)

func ReservationResponseChannel(id string) string {
	return "RESPONSE_" + id
}

func StartRecordingChannel(id string) string {
	return "START_RECORDING_" + id
}

func EndRecordingChannel(id string) string {
	return "END_RECORDING_" + id
}
