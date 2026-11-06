package connector

import (
	"strconv"
	"time"

	"github.com/livekit/protocol/livekit"
)

func RequestTimeAttributes(t time.Time) map[string]string {
	return timeAttributes(livekit.AttrConnectorRequestTime, t)
}

// ParseRequestTime returns the zero time if the attribute is missing or malformed.
func ParseRequestTime(attrs map[string]string) time.Time {
	return parseTimeAttribute(attrs, livekit.AttrConnectorRequestTime)
}

func AnswerTimeAttributes(t time.Time) map[string]string {
	return timeAttributes(livekit.AttrConnectorAnswerTime, t)
}

// ParseAnswerTime returns the zero time if the attribute is missing or malformed.
func ParseAnswerTime(attrs map[string]string) time.Time {
	return parseTimeAttribute(attrs, livekit.AttrConnectorAnswerTime)
}

func timeAttributes(key string, t time.Time) map[string]string {
	return map[string]string{
		key: strconv.FormatInt(t.UnixMilli(), 10),
	}
}

func parseTimeAttribute(attrs map[string]string, key string) time.Time {
	ms, err := strconv.ParseInt(attrs[key], 10, 64)
	if err != nil {
		return time.Time{}
	}
	return time.UnixMilli(ms)
}
