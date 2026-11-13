---
"@livekit/protocol": patch
---

Scope agent database access: `agentThread.sql` opens a covered thread's database, whose id is the thread's; `agentDatabase` reads, writes or deletes the standalone databases it lists, and creates or lists the project's. Remove `agent.databaseAdmin` and `AgentThread.Thread.database_id`
