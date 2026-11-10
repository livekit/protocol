---
"@livekit/protocol": patch
---

Add `END_REASON_SIMULATOR_ERROR` to `SimulationRun.JobMetrics.EndReason`, for a simulated call that ended because the caller's own models failed rather than anything the agent under test did
