---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

AMD v2 in agent session events (`livekit_agent_session.proto`): `AmdCategory` adds `AMD_MACHINE_SCREENING` and `AMD_WAIT`, `AmdPrediction` adds the turn ID, the stage, whether it corrects the stage, and the inference time, and a new `AmdCompleted` event reports the final result and the last prediction once per AMD run.
