-- Tables inside a thread's database. The agents framework creates and writes
-- them (lk-agents-storage, statements/sql.rs). The platform reads items and
-- tasks (agent-db's threads service: AgentThreads.GetThreadHistory and
-- ListThreadTasks).

CREATE TABLE schema_version (
  version    INTEGER NOT NULL,
  applied_at INTEGER NOT NULL
);

-- One chat item per row, in each agent's order. Used by: the framework, as
-- the model's history, and GetThreadHistory.
CREATE TABLE items (
  id         TEXT PRIMARY KEY,
  lane       TEXT NOT NULL,   -- the agent
  seq        INTEGER NOT NULL,
  message_id TEXT,            -- the A2A messageId it came from, for idempotency
  payload    JSONB NOT NULL   -- livekit.agent.ChatContext.ChatItem as protobuf JSON
);
CREATE INDEX items_lane_seq ON items (lane, seq);
CREATE UNIQUE INDEX items_lane_message ON items (lane, message_id) WHERE message_id IS NOT NULL;

-- One A2A task per row: the task of an incoming message, or a background
-- AgentTask (name set). Used by: the framework's A2A methods (message/send,
-- message/stream, tasks/get, tasks/list, tasks/cancel), which rebuild the A2A
-- Task as id = task_id, contextId = the thread id, status = (status,
-- status_message, updated_at), plus the fields of data. And ListThreadTasks,
-- for the rows with a name.
CREATE TABLE tasks (
  task_id        TEXT PRIMARY KEY,
  lane           TEXT NOT NULL,   -- the agent
  seq            INTEGER NOT NULL,
  name           TEXT,            -- the AgentTask class
  message_id     TEXT,            -- the A2A messageId that created it, for idempotency
  attempts       INTEGER NOT NULL DEFAULT 0,  -- pickups, a second one means the worker was lost
  status         TEXT NOT NULL,   -- A2A TaskState
  status_message JSONB,           -- A2A Message
  updated_at     INTEGER NOT NULL,
  data           JSONB NOT NULL,  -- A2A Task fields: history, artifacts, metadata
  state          JSONB            -- AgentTask arguments and pending question, to resume it
);
CREATE INDEX tasks_lane_status ON tasks (lane, status);
CREATE UNIQUE INDEX tasks_lane_message ON tasks (lane, message_id) WHERE message_id IS NOT NULL;

-- One row per agent: its lease (one worker runs an agent at a time), its
-- fencing version, state and agent stack. Used by: the framework, when a run
-- starts, commits and ends.
CREATE TABLE lanes (
  lane         TEXT PRIMARY KEY,
  version      INTEGER NOT NULL,
  state        JSONB,
  agent_stack  JSONB,
  lease_holder TEXT,
  lease_token  TEXT,
  lease_until  INTEGER NOT NULL DEFAULT 0,
  waiter       TEXT,
  drained_seq  INTEGER NOT NULL DEFAULT 0
);

-- The run journal: every AgentSession event, plus run_started and run_ended.
-- Used by: the framework, to tell whether a run ended, and for replay.
CREATE TABLE events (
  seq     INTEGER PRIMARY KEY,
  ts_ms   INTEGER NOT NULL,
  lane    TEXT NOT NULL,
  run_id  TEXT NOT NULL,
  author  TEXT NOT NULL,
  type    TEXT NOT NULL,
  payload JSONB NOT NULL
);
CREATE INDEX events_lane_seq ON events (lane, seq);

-- Messages that arrived while the agent was running. Used by: the framework,
-- which drains them in the next run.
CREATE TABLE inbox (
  seq         INTEGER PRIMARY KEY,
  lane        TEXT NOT NULL,
  task_id     TEXT NOT NULL,
  message_id  TEXT NOT NULL,
  message     JSONB NOT NULL,
  received_at INTEGER NOT NULL
);
CREATE INDEX inbox_lane_seq ON inbox (lane, seq);
CREATE UNIQUE INDEX inbox_lane_message ON inbox (lane, message_id);
