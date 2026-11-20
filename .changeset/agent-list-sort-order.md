---
"@livekit/protocol": patch
"github.com/livekit/protocol": patch
---

AgentThread.ListRequest and AgentDB.ListRequest: sort_order, newest first by default; pages of 50 by default, 100 at most; list rows leave activity, expiry and tip unset.
