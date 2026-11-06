---
"github.com/livekit/protocol": patch
---

`sdp.CodecsFromMediaDescription` builds the codec map once and returns each payload type once, even when the m= line repeats it.
`SDPFragment.PatchICECredentialAndCandidatesIntoSDP` adds the fragment's candidates only to the media section with the fragment's mid.
