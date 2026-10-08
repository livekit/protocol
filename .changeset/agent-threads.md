---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

Agent threads, the durable conversation an agent has with a person across channels: `AgentThread` messages (`agent/livekit_agent_thread.proto`), `auth.AgentThreadGrant` under the `agentThread` claim (optionally limited to some of the agent's endpoints), and the `AT_` (thread) and `AST_` (agent stream) id prefixes in `utils/guid`.

The agent-db and thread services are internal gRPC in cloud-protocol, so `AgentDBService` and its Twirp code are removed; the public API is public-api-server. A database lives in its project's data region: `AgentDB.CreateRequest.region` is reserved, and AgentDB times are `google.protobuf.Timestamp` / `Duration`.
