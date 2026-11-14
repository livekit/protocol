---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Remove `FunctionCall.extra`: livekit-agents 1.7.0 to 1.8.8 fill it from their own `FunctionCall.extra` and fail to encode the session history.
