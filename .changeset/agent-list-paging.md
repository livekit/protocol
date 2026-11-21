---
"@livekit/protocol": patch
"github.com/livekit/protocol": patch
---

AgentThread.ListRequest and AgentDB.ListRequest: unspecified order, stable across pages; pages of 50 by default, 100 at most, possibly short while a token is set; list rows leave activity, expiry and tip unset.
