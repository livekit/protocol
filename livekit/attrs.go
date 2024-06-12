package livekit

// Names of participant attributes for SIP.
const (
	// AttrSIPPrefix is shared for all SIP attributes.
	AttrSIPPrefix = "lk.sip."
	// AttrSIPCallID attribute contains LiveKit SIP call ID.
	AttrSIPCallID = AttrSIPPrefix + "callID"
	// AttrSIPTrunkID attribute contains LiveKit SIP Trunk ID used for the call.
	AttrSIPTrunkID = AttrSIPPrefix + "trunkID"
	// AttrSIPDispatchRuleID attribute contains LiveKit SIP DispatchRule ID used for the inbound call.
	AttrSIPDispatchRuleID = AttrSIPPrefix + "ruleID"
	// AttrSIPFromNumber attribute contains number from which the call was made.
	// This attribute will be omitted if HidePhoneNumber is set.
	AttrSIPFromNumber = AttrSIPPrefix + "fromNumber"
	// AttrSIPToNumber attribute contains number to which the call was made.
	// This attribute will be omitted if HidePhoneNumber is set.
	AttrSIPToNumber = AttrSIPPrefix + "toNumber"
)
