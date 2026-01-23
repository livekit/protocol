package livekit

func (p *IngressInfo) ID() string {
	if p == nil {
		return ""
	}
	return p.IngressId
}

func (p *ListIngressRequest) Filter(info *IngressInfo) bool {
	if info == nil {
		return true // for FilterSlice to work correctly with missing IDs
	}
	// Filter by room_name if specified
	if p.RoomName != "" && info.RoomName != p.RoomName {
		return false
	}
	// Filter by ingress_id if specified
	if p.IngressId != "" && info.IngressId != p.IngressId {
		return false
	}
	return true
}
