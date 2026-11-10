-- Tables inside a thread's database, created and written by lk-agents-storage.
-- A partition shares history, session state and write ownership across channels
-- and endpoints. The public default key is "default". The endpoint selects
-- the handler independently of the partition.
--
-- Every column says who reads it:
--   [platform]   agent-db's threads service: AgentThreads.GetThreadItems and
--                GetThreadTasks. Changing it breaks the platform API.
--   [a2a]        the framework's A2A methods (message/send, message/stream,
--                tasks/get, tasks/list, tasks/cancel), rebuilt into A2A objects.
--   [framework]  the framework's own run loop only, free to change with it.

CREATE TABLE schema_version (
  version INTEGER NOT NULL  -- [framework] compared to its own version on open
);

-- One model-context item per row: a message, tool call, tool result or handoff.
-- The payload is typed ChatContext.ChatItem protobuf JSON, not an A2A Message.
-- Keep query keys outside it, do not split its variant-specific fields into SQL columns.
CREATE TABLE items (
  pos           INTEGER PRIMARY KEY AUTOINCREMENT,  -- [platform] insertion order, never reused: the page token's sort key
  id            TEXT NOT NULL UNIQUE,               -- [framework] the item's id, its upsert key (ON CONFLICT DO UPDATE: a REPLACE would move pos)
  partition_key TEXT NOT NULL DEFAULT 'default' CHECK (length(partition_key) > 0),                      -- [platform, framework]
  seq           INTEGER NOT NULL,                   -- [platform, framework] model-context position in the partition
  message_id    TEXT DEFAULT NULL,                  -- [platform, framework] optional source A2A messageId for retry deduplication
  payload       JSONB NOT NULL                      -- [platform, framework] livekit.agent.ChatContext.ChatItem as protobuf JSON
);
-- Model-context order and stable insertion pagination are separate access paths.
CREATE INDEX items_partition_seq ON items (partition_key, seq);
CREATE INDEX items_partition_pos ON items (partition_key, pos);
CREATE UNIQUE INDEX items_partition_message ON items (partition_key, message_id) WHERE message_id IS NOT NULL;

-- One A2A task per row. The A2A Task is id = task_id, contextId = the thread
-- id, status = (status, data.status_message, updated_at), plus the fields of data.
CREATE TABLE tasks (
  pos           INTEGER PRIMARY KEY AUTOINCREMENT,  -- [platform] insertion order, never reused: the page token's sort key
  task_id       TEXT NOT NULL UNIQUE,               -- [platform, a2a, framework] its upsert key (ON CONFLICT DO UPDATE: a REPLACE would move pos)
  partition_key TEXT NOT NULL DEFAULT 'default' CHECK (length(partition_key) > 0),                      -- [framework] history and ownership partition
  -- [framework] stable destination (agent/deployment/endpoint), NULL for non-request tasks
  endpoint      TEXT CHECK (endpoint IS NULL OR length(endpoint) > 0),
  name          TEXT,                               -- [platform, framework] optional application task label
  -- [framework] Optional initiating A2A messageId for retry deduplication.
  -- NULL for tasks initiated without an incoming message.
  message_id    TEXT DEFAULT NULL,
  status        TEXT NOT NULL,                      -- [platform, a2a, framework] A2A task state
  updated_at    INTEGER NOT NULL,                   -- [platform, a2a, framework] Unix milliseconds
  data          JSONB NOT NULL,                     -- [a2a] task fields: history, artifacts, metadata, status_message
  internal      JSONB                               -- [framework] transport metadata and pending question
);
CREATE INDEX tasks_partition_pos ON tasks (partition_key, pos);
CREATE INDEX tasks_partition_status ON tasks (partition_key, status, pos);
-- A2A task listings use update time, independently of FIFO acceptance order.
CREATE INDEX tasks_updated ON tasks (updated_at DESC, task_id DESC);
-- Cross-endpoint FIFO uses durable acceptance position, never endpoint order.
CREATE UNIQUE INDEX tasks_partition_message ON tasks (partition_key, message_id) WHERE message_id IS NOT NULL;

-- One row per partition: mutable turns serialize under its lease.
-- Session routing is stored in the typed state snapshot.
CREATE TABLE partitions (
  partition_key TEXT PRIMARY KEY NOT NULL DEFAULT 'default' CHECK (length(partition_key) > 0),             -- [framework]
  version       INTEGER NOT NULL,             -- [framework] fencing version, bumped by every commit
  state         JSONB,                        -- [framework] typed state snapshot, no implied frame identity
  lease_holder  TEXT,                         -- [framework]
  lease_token   TEXT,                         -- [framework]
  lease_until   INTEGER NOT NULL DEFAULT 0,   -- [framework]
  waiter        TEXT,                         -- [framework] a worker waiting to interrupt the run
  drained_id    INTEGER NOT NULL DEFAULT 0    -- [framework] the last inbox row a run consumed
);

-- The run journal, append-only: every AgentSession event, plus run_started and
-- run_ended. Incremental writes can carry a producer-chosen event_key that is
-- reused on retry. Managed run commits retain their existing run-id guards.
CREATE TABLE events (
  id            INTEGER PRIMARY KEY,  -- [framework] insertion order
  ts_ms         INTEGER NOT NULL,     -- [framework]
  partition_key TEXT NOT NULL DEFAULT 'default' CHECK (length(partition_key) > 0),        -- [framework]
  run_id        TEXT NOT NULL,        -- [framework]
  -- [framework] optional idempotency identity for incremental recording
  event_key     TEXT CHECK (event_key IS NULL OR length(event_key) > 0),
  type          TEXT NOT NULL,        -- [framework]
  payload       JSONB NOT NULL        -- [framework]
);
CREATE INDEX events_partition_id ON events (partition_key, id);
-- Only terminal run receipts participate in commit/retry checks.
CREATE INDEX events_partition_run ON events (partition_key, run_id) WHERE type = 'run_ended';
CREATE UNIQUE INDEX events_partition_key ON events (partition_key, event_key) WHERE event_key IS NOT NULL;

-- Messages that arrived while the partition was running. Dispatch joins tasks
-- by (partition_key, task_id) for the destination, routing is stored only once. Batch
-- only consecutive compatible requests to one endpoint in acceptance order.
-- Ids are compared to partitions.drained_id, so the newest row is never deleted
-- (SQLite would hand its id out again).
CREATE TABLE inbox (
  id            INTEGER PRIMARY KEY,  -- [framework] arrival order
  partition_key TEXT NOT NULL DEFAULT 'default' CHECK (length(partition_key) > 0),        -- [framework]
  task_id       TEXT NOT NULL,        -- [framework]
  message_id    TEXT NOT NULL,        -- [framework] the A2A messageId, for idempotency
  message       JSONB NOT NULL,       -- [framework] the A2A message
  received_at   INTEGER NOT NULL      -- [framework] Unix seconds
);
CREATE INDEX inbox_partition_id ON inbox (partition_key, id);
CREATE INDEX inbox_partition_task ON inbox (partition_key, task_id);
CREATE UNIQUE INDEX inbox_partition_message ON inbox (partition_key, message_id);
