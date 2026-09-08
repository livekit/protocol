---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `deployment` to `SimulationRun.Create.Request`, `SimulationRun`, and `SimulationDispatch`. Empty/unset remains production, matching `CreateAgentDispatchRequest.deployment`, so Cloud simulations can pin a non-production worker for a shared agent name.
