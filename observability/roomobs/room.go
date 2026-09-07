package roomobs

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils"
)

const tagDelimiter = "\x1e"

type Tag string

func ToTag(key, value string) Tag {
	return Tag(key + tagDelimiter + value)
}

func (t Tag) KeyValue() (string, string) {
	key, value, ok := strings.Cut(string(t), tagDelimiter)
	if !ok {
		return string(t), ""
	}
	return key, value
}

type Tags []Tag

func (t Tags) Strings() []string {
	return utils.CastStringSlice[string](t)
}

func ToTags(m map[string]string) Tags {
	t := make(Tags, 0, len(m))
	for k, v := range m {
		t = append(t, ToTag(k, v))
	}
	return t
}

func PackTrackLayer(x, y uint32) uint32 {
	return uint32(x<<16 | y)
}

func UnpackTrackLayer(layer uint32) (x, y int) {
	return int(layer >> 16), int(layer & 0xffff)
}

func PackCountryCode(isoAlpha2 string) uint16 {
	if len(isoAlpha2) != 2 {
		return PackCountryCode("??")
	}
	return uint16(isoAlpha2[0])<<8 | uint16(isoAlpha2[1])
}

func UnpackCountryCode(code uint16) (isoAlpha2 string) {
	b := [2]byte{byte(code >> 8), byte(code)}
	return unsafe.String(&b[0], 2)
}

func ToClientOS(os string) ClientOS {
	switch strings.ToLower(os) {
	case "":
		return ClientOSUndefined
	case "ios":
		return ClientOSIos
	case "android":
		return ClientOSAndroid
	case "windows":
		return ClientOSWindows
	case "mac", "mac os x", "darwin", "macos":
		return ClientOSMac
	case "linux", "chrome os":
		return ClientOSLinux
	default:
		return ClientOSUndefined
	}
}

func FormatBrowser(clientInfo *livekit.ClientInfo) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", clientInfo.GetBrowser(), clientInfo.GetBrowserVersion()))
}

func FormatSDKVersion(clientInfo *livekit.ClientInfo) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", clientInfo.GetSdk(), clientInfo.GetVersion()))
}

func TrackKindFromProto(p livekit.StreamType) TrackKind {
	switch p {
	case livekit.StreamType_UPSTREAM:
		return TrackKindSub
	case livekit.StreamType_DOWNSTREAM:
		return TrackKindPub
	default:
		return TrackKindUndefined
	}
}

func TrackTypeFromProto(p livekit.TrackType) TrackType {
	switch p {
	case livekit.TrackType_AUDIO:
		return TrackTypeAudio
	case livekit.TrackType_VIDEO:
		return TrackTypeVideo
	case livekit.TrackType_DATA:
		return TrackTypeData
	default:
		return TrackTypeUndefined
	}
}

func TrackSourceFromProto(p livekit.TrackSource) TrackSource {
	switch p {
	case livekit.TrackSource_UNKNOWN:
		return TrackSourceUndefined
	case livekit.TrackSource_CAMERA:
		return TrackSourceCamera
	case livekit.TrackSource_MICROPHONE:
		return TrackSourceMicrophone
	case livekit.TrackSource_SCREEN_SHARE:
		return TrackSourceScreenShare
	case livekit.TrackSource_SCREEN_SHARE_AUDIO:
		return TrackSourceScreenShareAudio
	default:
		return TrackSourceUndefined
	}
}

type RoomFeature uint16

func (f RoomFeature) HasIngress() bool    { return f&IngressRoomFeature != 0 }
func (f RoomFeature) HasEgress() bool     { return f&EgressRoomFeature != 0 }
func (f RoomFeature) HasSIP() bool        { return f&SIPRoomFeature != 0 }
func (f RoomFeature) HasAgent() bool      { return f&AgentRoomFeature != 0 }
func (f RoomFeature) HasConnector() bool  { return f&ConnectorRoomFeature != 0 }
func (f RoomFeature) HasSimulation() bool { return f&SimulationRoomFeature != 0 }

const (
	IngressRoomFeature RoomFeature = 1 << iota
	EgressRoomFeature
	SIPRoomFeature
	AgentRoomFeature
	ConnectorRoomFeature
	SimulationRoomFeature
)

// RoomFeatureFromParticipantKind derives the room-session features implied by a
// participant's kind and any kind details. Features are additive: a single
// participant can contribute multiple bits (e.g. an AGENT participant flagged
// with the SIMULATION kind detail yields both AgentRoomFeature and
// SimulationRoomFeature).
func RoomFeatureFromParticipantKind(k livekit.ParticipantInfo_Kind, details ...livekit.ParticipantInfo_KindDetail) RoomFeature {
	var f RoomFeature
	switch k {
	case livekit.ParticipantInfo_INGRESS:
		f = IngressRoomFeature
	case livekit.ParticipantInfo_EGRESS:
		f = EgressRoomFeature
	case livekit.ParticipantInfo_SIP:
		f = SIPRoomFeature
	case livekit.ParticipantInfo_AGENT:
		f = AgentRoomFeature
	case livekit.ParticipantInfo_CONNECTOR:
		f = ConnectorRoomFeature
	}
	for _, d := range details {
		switch d {
		case livekit.ParticipantInfo_SIMULATION:
			f |= SimulationRoomFeature
		}
	}
	return f
}

func ParticipantKindCode(k livekit.ParticipantInfo_Kind) int32 {
	return int32(k)
}

func ParticipantKindDetailsCodes(d []livekit.ParticipantInfo_KindDetail) []int32 {
	return *(*[]int32)(unsafe.Pointer(&d))
}
