---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

AMD v2 in agent session events (`livekit_agent_session.proto`): `AmdCategory` adds `AMD_MACHINE_SCREENING` and `AMD_WAIT`, `AmdPrediction` adds the turn, stage, and previous categories, and a new `AmdCompleted` event reports the final result once per AMD run.
