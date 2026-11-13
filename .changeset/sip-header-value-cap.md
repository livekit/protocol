---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Raise the SIP header value cap from 1 KB to 4 KB so carrier headers like Twilio's `X-Twilio-CallToken` (Immutable Call Forwarding token with SHAKEN/STIR PASSporTs, >1 KB) pass `CreateSIPParticipant` validation