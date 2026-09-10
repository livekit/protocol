---
"github.com/livekit/protocol": minor
---

Agent simulation: add `dispatch_agent_name` to `SimulationRun.Create.Request`. `agent_name` is now the agent under test as named in the project's `livekit.toml`, stable across runs so list/counts filters and the dashboard can group by it; `dispatch_agent_name`, when set, is the explicit-dispatch name the simulator sends jobs to (the CLI's throwaway local-worker name). Unset dispatches to `agent_name`, so existing callers are unchanged.
