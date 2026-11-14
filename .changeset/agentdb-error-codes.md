---
"@livekit/protocol": patch
"github.com/livekit/protocol": patch
---

AgentDB.ErrorCode: every agent-db error code, with whether it may have applied and whether it may be retried, as enum value options.
AgentDB.Wire: Hello carries client_id, client_version and max_frame; HelloOk lists capabilities; every request carries deadline_ms and a write its seq; a Batch may carry if_tip.
