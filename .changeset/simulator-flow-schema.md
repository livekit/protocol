---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Simulator flows in agent simulation (`livekit_agent_simulation.proto`): `ScenarioGroup.nodes` holds named `SimulatorNode`s shared by the group's scenarios, and `Scenario.start` names the node a call starts at, so the simulator plays the far end of an outbound call (voicemail, IVR menus, screeners, holds, and people) instead of one persona. A scenario sets `start` or `instructions`, not both. `JobMetrics.EndReason` adds `END_REASON_SIMULATOR_HUNG_UP` for a flow that hangs up.
