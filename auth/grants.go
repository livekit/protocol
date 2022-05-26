package auth

import (
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

	// permissions within a room, if none of the permissions are set
	// it's interpreted as both are permissible
	CanPublish     *bool `json:"canPublish,omitempty"`
	CanSubscribe   *bool `json:"canSubscribe,omitempty"`
	CanPublishData *bool `json:"canPublishData,omitempty"`

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

func (v *VideoGrant) GetCanPublish() bool {
	if v.CanPublish == nil {
		return true
	}
	return *v.CanPublish
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

func (v *VideoGrant) ToPermission() *livekit.ParticipantPermission {
	return &livekit.ParticipantPermission{
		CanPublish:     v.GetCanPublish(),
		CanPublishData: v.GetCanPublishData(),
		CanSubscribe:   v.GetCanSubscribe(),
		Hidden:         v.Hidden,
		Recorder:       v.Recorder,
	}
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

	return &clone
}
