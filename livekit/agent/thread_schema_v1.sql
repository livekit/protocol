-- Tables inside a thread's database. The agents framework creates and writes
-- them (lk-agents-storage, statements/sql.rs). A lane is one registered
-- agent's part of the thread, named after the agent, across every channel. Two
-- agents in one thread, through a delegation, are two lanes.
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

-- One chat item per row.
CREATE TABLE items (
  id         TEXT PRIMARY KEY,  -- [framework] the item's id, its upsert key
  lane       TEXT NOT NULL,     -- [platform, framework]
  seq        INTEGER NOT NULL,  -- [platform, framework] position in the lane
  message_id TEXT,              -- [platform, framework] the A2A messageId it came from, for idempotency
  payload    JSONB NOT NULL     -- [platform, framework] livekit.agent.ChatContext.ChatItem as protobuf JSON
);
CREATE INDEX items_lane_seq ON items (lane, seq);
CREATE UNIQUE INDEX items_lane_message ON items (lane, message_id) WHERE message_id IS NOT NULL;

-- One A2A task per row: the task of an incoming message, or a background
-- AgentTask (name set). The A2A Task is id = task_id, contextId = the thread
-- id, status = (status, status_message, updated_at), plus the fields of data.
CREATE TABLE tasks (
  task_id        TEXT PRIMARY KEY,  -- [platform, a2a, framework]
  lane           TEXT NOT NULL,     -- [framework]
  seq            INTEGER NOT NULL,  -- [framework] submission order in the lane
  name           TEXT,              -- [platform, framework] the AgentTask class
  message_id     TEXT,              -- [framework] the A2A messageId that created it, for idempotency
  status         TEXT NOT NULL,     -- [platform, a2a, framework] A2A TaskState, working when a new run finds it means its worker was lost
  status_message JSONB,             -- [a2a] A2A Message
  updated_at     INTEGER NOT NULL,  -- [platform, a2a, framework]
  data           JSONB NOT NULL,    -- [a2a] A2A Task fields: history, artifacts, metadata
  internal       JSONB              -- [framework] AgentTask arguments and pending question, to resume it
);
CREATE INDEX tasks_lane_seq ON tasks (lane, seq);
CREATE INDEX tasks_lane_status ON tasks (lane, status, seq);
CREATE UNIQUE INDEX tasks_lane_message ON tasks (lane, message_id) WHERE message_id IS NOT NULL;

-- One row per lane: one worker runs a lane at a time.
CREATE TABLE lanes (
  lane         TEXT PRIMARY KEY,             -- [framework]
  version      INTEGER NOT NULL,             -- [framework] fencing version, bumped by every commit
  state        JSONB,                        -- [framework] the agent's typed state
  agent_stack  JSONB,                        -- [framework] who is speaking
  lease_holder TEXT,                         -- [framework]
  lease_token  TEXT,                         -- [framework]
  lease_until  INTEGER NOT NULL DEFAULT 0,   -- [framework]
  waiter       TEXT,                         -- [framework] a worker waiting to interrupt the run
  drained_id   INTEGER NOT NULL DEFAULT 0    -- [framework] the last inbox row a run consumed
);

-- The run journal, append-only: every AgentSession event, plus run_started and
-- run_ended.
CREATE TABLE events (
  id      INTEGER PRIMARY KEY,  -- [framework] insertion order
  ts_ms   INTEGER NOT NULL,     -- [framework]
  lane    TEXT NOT NULL,        -- [framework]
  run_id  TEXT NOT NULL,        -- [framework]
  type    TEXT NOT NULL,        -- [framework]
  payload JSONB NOT NULL        -- [framework]
);
CREATE INDEX events_lane_id ON events (lane, id);
CREATE INDEX events_lane_run ON events (lane, run_id);

-- Messages that arrived while the lane was running, drained by the next run.
-- Ids are compared to lanes.drained_id, so the newest row is never deleted
-- (SQLite would hand its id out again).
CREATE TABLE inbox (
  id          INTEGER PRIMARY KEY,  -- [framework] arrival order
  lane        TEXT NOT NULL,        -- [framework]
  task_id     TEXT NOT NULL,        -- [framework]
  message_id  TEXT NOT NULL,        -- [framework] the A2A messageId, for idempotency
  message     JSONB NOT NULL,       -- [framework] the A2A message
  received_at INTEGER NOT NULL      -- [framework]
);
CREATE INDEX inbox_lane_id ON inbox (lane, id);
CREATE INDEX inbox_lane_task ON inbox (lane, task_id);
CREATE UNIQUE INDEX inbox_lane_message ON inbox (lane, message_id);
