---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Agent thread schema version 2: `tasks.data` holds the task's A2A v1 fields as protobuf JSON, its status message included (`tasks.status_message` is gone), `tasks.updated_at` is in milliseconds, and `tasks.internal` is `tasks.ask`, the question a waiting task asked. `thread_schema_v1.sql` is `thread_schema.sql`, and `ThreadSchemaV1` is `ThreadSchema`.
