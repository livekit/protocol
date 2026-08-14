---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `to_user_override` to `CreateSIPParticipantRequest` to override the user part of the `To` header, and forbid `To`/`Via` in the SIP headers map.
