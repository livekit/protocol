---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

Agent threads, the durable conversation an agent has with a person across channels (`agent/livekit_agent_thread.proto`): `AgentThread` (the thread, its items as `ChatContext.ChatItem`) and `AgentTask` (background work an agent runs in it); the conversation tables in `livekit/agent/thread_schema_v1.sql`; `auth.AgentThreadGrant` under the `agentThread` claim (list and create by scope, read, write and delete by thread id); the `AT_` (thread) and `AST_` (agent stream) id prefixes in `utils/guid`.

The agent-db and thread services are internal gRPC in cloud-protocol, so `AgentDBService` and its Twirp code are removed; the public API is public-api-server. A database lives in its project's data region: `AgentDB.CreateRequest.region` is removed, and AgentDB times are `google.protobuf.Timestamp` / `Duration`.
