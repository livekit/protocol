---
"@livekit/protocol": patch
---

Scope agent database access: `agentThread.sql` opens a covered thread's database, whose id is the thread's; `agentDatabase` grants standalone databases, or every database with `admin`. Remove `agent.databaseAdmin` and `AgentThread.Thread.database_id`
