---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

Agent threads, the durable conversation an agent has with a person across channels: `AgentThread` messages (`agent/livekit_agent_thread.proto`), `thread_id` on `AgentHttp.StreamPreamble` and `Job`, `UpdateThreadStatus` on the worker control stream, `run_started` and `run_ended` on `AgentSessionEvent`, the `AT_` and `AST_` id prefixes in `utils/guid`, and `auth.AgentThreadGrant` under the `agentThread` claim. On the agent-db wire, `Columns`, `ColumnBatch` and `ExecResult` carry the 0-based `statement` index of the `Batch` statement they answer.

`AgentHttp.StreamPreamble.thread_grant` carries the verified `agentThread` claim to the worker as `AgentHttp.AgentThreadGrant` (agent name, subject, thread ids, action flags); the token itself never reaches the worker. `auth.AgentThreadGrant.ToProto` converts the claim.

The agent-db and thread services move to internal gRPC in cloud-protocol, so `AgentDBService` and its Twirp code are removed; the public API is public-api-server. A database lives in its project's data region: `AgentDB.CreateRequest.region` is reserved.
