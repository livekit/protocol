package sip

import (
	"strconv"
	"time"

	"github.com/livekit/protocol/livekit"
)

func InviteTimeAttributes(t time.Time) map[string]string {
	return timeAttributes(livekit.AttrSIPInviteTime, t)
}

// ParseInviteTime returns the zero time if the attribute is missing or malformed.
func ParseInviteTime(attrs map[string]string) time.Time {
	return parseTimeAttribute(attrs, livekit.AttrSIPInviteTime)
}

func PinEnteredTimeAttributes(t time.Time) map[string]string {
	return timeAttributes(livekit.AttrSIPPinEnteredTime, t)
}

// ParsePinEnteredTime returns the zero time if the attribute is missing or malformed.
func ParsePinEnteredTime(attrs map[string]string) time.Time {
	return parseTimeAttribute(attrs, livekit.AttrSIPPinEnteredTime)
}

func AnswerTimeAttributes(t time.Time) map[string]string {
	return timeAttributes(livekit.AttrSIPAnswerTime, t)
}

// ParseAnswerTime returns the zero time if the attribute is missing or malformed.
func ParseAnswerTime(attrs map[string]string) time.Time {
	return parseTimeAttribute(attrs, livekit.AttrSIPAnswerTime)
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
