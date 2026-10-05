package sip

import (
	"strconv"
	"time"

	"github.com/livekit/protocol/livekit"
)

func InviteTimeAttributes(t time.Time) map[string]string {
	return map[string]string{
		livekit.AttrSIPInviteTime: strconv.FormatInt(t.UnixMilli(), 10),
	}
}

// ParseInviteTime returns the zero time if the attribute is missing or malformed.
func ParseInviteTime(attrs map[string]string) time.Time {
	ms, err := strconv.ParseInt(attrs[livekit.AttrSIPInviteTime], 10, 64)
	if err != nil {
		return time.Time{}
	}
	return time.UnixMilli(ms)
}
