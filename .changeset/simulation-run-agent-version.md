---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `agent_version` to `SimulationRun` and `SimulationRun.Job` — the deployed cloud agent version (the deploy tag shown by `lk agent versions`) that served a run, as self-reported by the agent participant's `lk.agent.version` attribute. The run-level field is set only when every job in the run reports the same version.
