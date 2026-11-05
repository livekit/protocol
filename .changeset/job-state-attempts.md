---
"@livekit/protocol": patch
---

Add `attempts`, `first_started_at` and `agent_failed_attempts_ns` to `JobState`, so every node that syncs a job sees how many attempts it took, when the first one started, and how long the attempts the agent's own code failed took
