package livekit

// ToProto implements DataPacket in Go SDK.
func (p *SipDTMF) ToProto() *DataPacket {
	return &DataPacket{
		Value: &DataPacket_SipDtmf{
			SipDtmf: p,
		},
	}
}

// AsInbound converts legacy SIPTrunkInfo to SIPInboundTrunkInfo.
func (p *SIPTrunkInfo) AsInbound() *SIPInboundTrunkInfo {
	if p == nil || p.Kind == SIPTrunkInfo_TRUNK_OUTBOUND {
		return nil
	}
	var nums []string
	if p.OutboundNumber != "" {
		nums = []string{p.OutboundNumber}
	}
	return &SIPInboundTrunkInfo{
		SipTrunkId:       p.SipTrunkId,
		Name:             p.Name,
		Metadata:         p.Metadata,
		Numbers:          nums,
		AllowedAddresses: p.InboundAddresses,
		AllowedNumbers:   p.InboundNumbers,
		AuthUsername:     p.InboundUsername,
		AuthPassword:     p.InboundPassword,
	}
}

// AsTrunkInfo converts SIPInboundTrunkInfo to legacy SIPTrunkInfo.
func (p *SIPInboundTrunkInfo) AsTrunkInfo() *SIPTrunkInfo {
	if p == nil {
		return nil
	}
	var num string
	if len(p.Numbers) != 0 {
		num = p.Numbers[0]
	}
	return &SIPTrunkInfo{
		SipTrunkId:       p.SipTrunkId,
		Kind:             SIPTrunkInfo_TRUNK_INBOUND,
		Name:             p.Name,
		Metadata:         p.Metadata,
		OutboundNumber:   num,
		InboundAddresses: p.AllowedAddresses,
		InboundNumbers:   p.AllowedNumbers,
		InboundUsername:  p.AuthUsername,
		InboundPassword:  p.AuthPassword,
	}
}

// AsOutbound converts legacy SIPTrunkInfo to SIPOutboundTrunkInfo.
func (p *SIPTrunkInfo) AsOutbound() *SIPOutboundTrunkInfo {
	if p == nil || p.Kind == SIPTrunkInfo_TRUNK_INBOUND {
		return nil
	}
	var nums []string
	if p.OutboundNumber != "" {
		nums = []string{p.OutboundNumber}
	}
	return &SIPOutboundTrunkInfo{
		SipTrunkId:   p.SipTrunkId,
		Name:         p.Name,
		Metadata:     p.Metadata,
		Address:      p.OutboundAddress,
		Transport:    p.Transport,
		Numbers:      nums,
		AuthUsername: p.OutboundUsername,
		AuthPassword: p.OutboundPassword,
	}
}

// AsTrunkInfo converts SIPOutboundTrunkInfo to legacy SIPTrunkInfo.
func (p *SIPOutboundTrunkInfo) AsTrunkInfo() *SIPTrunkInfo {
	if p == nil {
		return nil
	}
	var num string
	if len(p.Numbers) != 0 {
		num = p.Numbers[0]
	}
	return &SIPTrunkInfo{
		SipTrunkId:       p.SipTrunkId,
		Kind:             SIPTrunkInfo_TRUNK_OUTBOUND,
		Name:             p.Name,
		Metadata:         p.Metadata,
		OutboundAddress:  p.Address,
		Transport:        p.Transport,
		OutboundNumber:   num,
		OutboundUsername: p.AuthUsername,
		OutboundPassword: p.AuthPassword,
	}
}
