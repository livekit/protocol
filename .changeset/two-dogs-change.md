---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

WHIP protocol change

Deprecate the bypass_transcoding property in all ingress APIs and introduce the optional enable_transcoding property. This property will default to false for WHIP and to true for all other ingress types.
