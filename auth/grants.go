package auth

import (
	"strings"

	"github.com/livekit/protocol/livekit"
)

type VideoGrant struct {
	// actions on rooms
	RoomCreate bool `json:"roomCreate,omitempty"`
	RoomList   bool `json:"roomList,omitempty"`
	RoomRecord bool `json:"roomRecord,omitempty"`

	// actions on a particular room
	RoomAdmin bool   `json:"roomAdmin,omitempty"`
	RoomJoin  bool   `json:"roomJoin,omitempty"`
	Room      string `json:"room,omitempty"`

	// permissions within a room, if none of the permissions are set explicitly
	// it will be granted with all publish and subscribe permissions
	CanPublish     *bool `json:"canPublish,omitempty"`
	CanSubscribe   *bool `json:"canSubscribe,omitempty"`
	CanPublishData *bool `json:"canPublishData,omitempty"`
	// TrackSource types that a participant may publish.
	// When set, it supercedes CanPublish. Only sources explicitly set here can be published
	CanPublishSources []string `json:"canPublishSources,omitempty"` // keys keep track of each source
	// by default, a participant is not allowed to update its own metadata
	CanUpdateOwnMetadata *bool `json:"canUpdateOwnMetadata,omitempty"`

	// actions on ingresses
	IngressAdmin bool `json:"ingressAdmin,omitempty"` // applies to all ingress

	// participant is not visible to other participants
	Hidden bool `json:"hidden,omitempty"`
	// indicates to the room that current participant is a recorder
	Recorder bool `json:"recorder,omitempty"`
}

type ClaimGrants struct {
	Identity string      `json:"-"`
	Name     string      `json:"name,omitempty"`
	Video    *VideoGrant `json:"video,omitempty"`
	// for verifying integrity of the message body
	Sha256   string `json:"sha256,omitempty"`
	Metadata string `json:"metadata,omitempty"`
}

func (c *ClaimGrants) Clone() *ClaimGrants {
	if c == nil {
		return nil
	}

	clone := *c
	clone.Video = c.Video.Clone()

	return &clone
}

func (v *VideoGrant) SetCanPublish(val bool) {
	v.CanPublish = &val
}

func (v *VideoGrant) SetCanPublishData(val bool) {
	v.CanPublishData = &val
}

func (v *VideoGrant) SetCanSubscribe(val bool) {
	v.CanSubscribe = &val
}

func (v *VideoGrant) SetCanPublishSources(sources []livekit.TrackSource) {
	v.CanPublishSources = make([]string, 0, len(sources))
	for _, s := range sources {
		v.CanPublishSources = append(v.CanPublishSources, sourceToString(s))
	}
}

func (v *VideoGrant) SetCanUpdateOwnMetadata(val bool) {
	v.CanUpdateOwnMetadata = &val
}

func (v *VideoGrant) GetCanPublish() bool {
	if v.CanPublish == nil {
		return true
	}
	return *v.CanPublish
}

func (v *VideoGrant) GetCanPublishSource(source livekit.TrackSource) bool {
	// don't differentiate between nil and unset, since that distinction doesn't survive serialization
	if len(v.CanPublishSources) == 0 {
		return v.GetCanPublish()
	}
	sourceStr := sourceToString(source)
	for _, s := range v.CanPublishSources {
		if s == sourceStr {
			return true
		}
	}
	return false
}

func (v *VideoGrant) GetCanPublishData() bool {
	if v.CanPublishData == nil {
		return v.GetCanPublish()
	}
	return *v.CanPublishData
}

func (v *VideoGrant) GetCanSubscribe() bool {
	if v.CanSubscribe == nil {
		return true
	}
	return *v.CanSubscribe
}

func (v *VideoGrant) GetCanUpdateOwnMetadata() bool {
	if v.CanUpdateOwnMetadata == nil {
		return false
	}
	return *v.CanUpdateOwnMetadata
}

func (v *VideoGrant) ToPermission() *livekit.ParticipantPermission {
	pp := &livekit.ParticipantPermission{
		CanPublish:     v.GetCanPublish(),
		CanPublishData: v.GetCanPublishData(),
		CanSubscribe:   v.GetCanSubscribe(),
		Hidden:         v.Hidden,
		Recorder:       v.Recorder,
	}

	for _, s := range v.CanPublishSources {
		pp.CanPublishSources = append(pp.CanPublishSources, sourceToProto(s))
	}

	return pp
}

func (v *VideoGrant) Clone() *VideoGrant {
	if v == nil {
		return nil
	}

	clone := *v

	if v.CanPublish != nil {
		canPublish := *v.CanPublish
		clone.CanPublish = &canPublish
	}

	if v.CanSubscribe != nil {
		canSubscribe := *v.CanSubscribe
		clone.CanSubscribe = &canSubscribe
	}

	if v.CanPublishData != nil {
		canPublishData := *v.CanPublishData
		clone.CanPublishData = &canPublishData
	}

	if v.CanPublishSources != nil {
		clone.CanPublishSources = make([]string, len(v.CanPublishSources))
		copy(clone.CanPublishSources, v.CanPublishSources)
	}

	if v.CanUpdateOwnMetadata != nil {
		canUpdateOwnMetadata := *v.CanUpdateOwnMetadata
		clone.CanUpdateOwnMetadata = &canUpdateOwnMetadata
	}

	return &clone
}

func sourceToString(source livekit.TrackSource) string {
	return strings.ToLower(source.String())
}

func sourceToProto(sourceStr string) livekit.TrackSource {
	switch sourceStr {
	case "camera":
		return livekit.TrackSource_CAMERA
	case "microphone":
		return livekit.TrackSource_MICROPHONE
	case "screen_share":
		return livekit.TrackSource_SCREEN_SHARE
	case "screen_share_audio":
		return livekit.TrackSource_SCREEN_SHARE_AUDIO
	default:
		return livekit.TrackSource_UNKNOWN
	}
}
