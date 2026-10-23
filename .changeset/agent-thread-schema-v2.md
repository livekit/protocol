---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Agent thread schema version 2, for A2A v1: `tasks.status` holds the v1 task state names (`TASK_STATE_WORKING` and so on, `TASK_STATE_AUTH_REQUIRED` included), messages and tasks are v1 protobuf JSON, `tasks.updated_at` is in milliseconds, and `tasks.internal` is `tasks.ask`, the question a waiting task asked. `thread_schema_v1.sql` is `thread_schema.sql`, and `ThreadSchemaV1` is `ThreadSchema`.
