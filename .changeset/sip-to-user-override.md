---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `to_user_override` to `CreateSIPParticipantRequest` to override the user part of the `To` header, and warn when `To`/`Via` are set in the SIP headers map.
