-- Tables inside one thread's database, created and written by lk-agents-storage.
-- A thread is one durable conversation, identified by the A2A contextId.
--
-- A partition is a named history, session-state and write-ownership scope within
-- that thread. Text and voice endpoints with the same partition_key share that
-- history, state and inbox. Mutable turns serialize under the partition's lease.
-- Different keys keep those records and ownership separate, for example a fast
-- brain in "default" and a slow brain in "backend". The framework resolves an
-- omitted key to DefaultPartitionKey before writing. SQL requires an explicit key.
-- The endpoint selects which handler runs, independently of the partition.
-- Application context fields retain their explicitly selected storage scopes.
--
-- Column readers:
--   [platform]  agent-db's AgentThreads.GetThreadItems and GetThreadTasks
--   [a2a]       framework A2A methods, which reconstruct wire objects
--   [framework] framework run loop, leases and persistence

CREATE TABLE schema_version (
  version INTEGER NOT NULL  -- [framework] checked when opening the database
);

-- One typed model-context item: message, tool call, tool result or handoff.
-- payload is ChatContext.ChatItem protobuf JSON. Query keys stay outside it.
-- Upsert by id with DO UPDATE, preserving pos. REPLACE would move the cursor.
CREATE TABLE items (
  pos           INTEGER PRIMARY KEY AUTOINCREMENT, -- [platform] stable insertion cursor
  id            TEXT NOT NULL UNIQUE,              -- [framework] item's upsert identity
  partition_key TEXT NOT NULL -- [platform, framework]
                CHECK(length(partition_key)>0),
  seq           INTEGER NOT NULL,                  -- [platform, framework] model-context order
  message_id    TEXT DEFAULT NULL,                 -- [platform, framework] optional messageId
  payload       JSONB NOT NULL                      -- [platform, framework] typed ChatItem
);
-- History order and insertion pagination are separate access paths.
CREATE INDEX items_partition_seq ON items (partition_key, seq);
CREATE INDEX items_partition_pos ON items (partition_key, pos);
CREATE UNIQUE INDEX items_partition_message
  ON items (partition_key, message_id) WHERE message_id IS NOT NULL;

-- One A2A task. Its contextId is the enclosing thread id.
-- status combines status, data.status_message and updated_at.
-- Optional message_id deduplicates the initiating message on retry. Tasks created
-- without an incoming message leave it NULL. Later messages use inbox/history.
-- endpoint stores the stable agent/deployment/endpoint destination when applicable.
CREATE TABLE tasks (
  pos           INTEGER PRIMARY KEY AUTOINCREMENT, -- [platform, framework] cursor and FIFO order
  task_id       TEXT NOT NULL UNIQUE,              -- [platform, a2a, framework] task identity
  partition_key TEXT NOT NULL -- [framework] owner
                CHECK(length(partition_key)>0),
  endpoint      TEXT CHECK(endpoint IS NULL OR length(endpoint)>0), -- [framework] optional route
  name          TEXT,                              -- [platform, framework] optional label
  message_id    TEXT DEFAULT NULL,                 -- [framework] optional initiating A2A messageId
  status        TEXT NOT NULL,                     -- [platform, a2a, framework] A2A task state
  updated_at    INTEGER NOT NULL,                  -- [platform, a2a, framework] Unix milliseconds
  data          JSONB NOT NULL,                    -- [a2a] task fields and status_message
  internal      JSONB                              -- [framework] transport state and question
);
CREATE INDEX tasks_partition_pos ON tasks (partition_key, pos);
CREATE INDEX tasks_partition_status ON tasks (partition_key, status, pos);
-- A2A task listings use update time, independently of acceptance order.
CREATE INDEX tasks_updated ON tasks (updated_at DESC, task_id DESC);
CREATE UNIQUE INDEX tasks_partition_message
  ON tasks (partition_key, message_id) WHERE message_id IS NOT NULL;

-- One row per partition. Local session routing lives in the typed state snapshot.
CREATE TABLE partitions (
  partition_key TEXT PRIMARY KEY NOT NULL -- [framework]
                CHECK(length(partition_key)>0),
  version       INTEGER NOT NULL,           -- [framework] fencing version, incremented per commit
  state         JSONB,                      -- [framework] typed session-state snapshot
  lease_holder  TEXT,                       -- [framework] owning worker
  lease_token   TEXT,                       -- [framework] identity of this lease acquisition
  lease_until   INTEGER NOT NULL DEFAULT 0, -- [framework] lease expiry, Unix seconds
  waiter        TEXT,                       -- [framework] worker requesting interruption
  drained_id    INTEGER NOT NULL DEFAULT 0  -- [framework] last consumed inbox id
);

-- Append-only journal: session events, run_started and run_ended.
-- Incremental writers reuse event_key on retry. Managed commits use run-id guards.
CREATE TABLE events (
  id            INTEGER PRIMARY KEY,       -- [framework] insertion order
  ts_ms         INTEGER NOT NULL,          -- [framework] Unix milliseconds
  partition_key TEXT NOT NULL -- [framework]
                CHECK(length(partition_key)>0),
  run_id        TEXT NOT NULL,             -- [framework] producing run
  event_key     TEXT CHECK(event_key IS NULL OR length(event_key)>0), -- [framework] retry id
  type          TEXT NOT NULL,             -- [framework] event kind
  payload       JSONB NOT NULL              -- [framework] event data
);
CREATE INDEX events_partition_id ON events (partition_key, id);
-- Only terminal run receipts participate in commit/retry checks.
CREATE INDEX events_partition_run ON events (partition_key, run_id) WHERE type = 'run_ended';
CREATE UNIQUE INDEX events_partition_key
  ON events (partition_key, event_key) WHERE event_key IS NOT NULL;

-- Messages accepted while a partition was running. Join tasks by
-- (partition_key, task_id) for their destination. Batch only consecutive compatible
-- requests for one endpoint. Keep the newest row when pruning, preventing id reuse
-- behind partitions.drained_id.
CREATE TABLE inbox (
  id            INTEGER PRIMARY KEY,       -- [framework] arrival order
  partition_key TEXT NOT NULL -- [framework]
                CHECK(length(partition_key)>0),
  task_id       TEXT NOT NULL,             -- [framework] destination task
  message_id    TEXT NOT NULL,             -- [framework] A2A messageId, deduplicated per partition
  message       JSONB NOT NULL,            -- [framework] A2A Message
  received_at   INTEGER NOT NULL            -- [framework] Unix seconds
);
CREATE INDEX inbox_partition_id ON inbox (partition_key, id);
CREATE INDEX inbox_partition_task ON inbox (partition_key, task_id);
CREATE UNIQUE INDEX inbox_partition_message ON inbox (partition_key, message_id);
