---
"@livekit/protocol": patch
---

Add `thread_id` to `CreateAgentDispatchRequest`, `RoomAgentDispatch`, `AgentDispatch` and `Job`, the agent thread a dispatched job's call reads and writes

Add `guid.IsValidWithPrefix`, a strict check that an id is a prefix followed by exactly `guid.Size` characters of `guid.IDAlphabet`
