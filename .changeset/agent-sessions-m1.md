---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Agent sessions: add the `AgentSessionRegistry` Twirp service (`agent/livekit_agent_session_registry.proto`), `session_id` on `AgentHttp.StreamPreamble` and `Job`, `UpdateSessionStatus` on the worker control stream, `run_started` and `run_ended` on `AgentSessionEvent`, the `AT_` and `AST_` id prefixes in `utils/guid`, and `auth.AgentSessionGrant` under the `agentSession` claim. On the agent-db wire, `Columns`, `ColumnBatch` and `ExecResult` carry the 0-based `statement` index of the `Batch` statement they answer.

`AgentSessionRegistry.RedirectSession` sets `redirect_to` on a fresh session after its items were copied into the found one; project credential only, `failed_precondition` when the row already redirects elsewhere.

`AgentHttp.StreamPreamble.session_grant` carries the verified `agentSession` claim to the worker as `AgentHttp.AgentSessionGrant` (agent name, subject, session ids, action flags); the token itself still never reaches the worker. `auth.AgentSessionGrant.ToProto` converts the claim.
