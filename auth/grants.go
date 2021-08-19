package auth

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

	// used for recording
	Hidden bool `json:"hidden,omitempty"`
}

type ClaimGrants struct {
	Identity string      `json:"-"`
	Video    *VideoGrant `json:"video,omitempty"`
	// for verifying integrity of the message body
	Sha256   string `json:"sha256,omitempty"`
	Metadata string `json:"metadata,omitempty"`
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
