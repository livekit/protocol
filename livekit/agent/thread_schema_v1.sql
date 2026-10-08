-- Tables inside a thread's database, created by the agents framework. The
-- platform reads items and tasks. items.payload and tasks.history hold
-- livekit.agent.ChatContext.ChatItem as protobuf JSON.

CREATE TABLE schema_version (
  version    INTEGER NOT NULL,
  applied_at INTEGER NOT NULL
);

CREATE TABLE items (
  lane       TEXT NOT NULL,
  id         TEXT PRIMARY KEY,
  seq        INTEGER NOT NULL,
  type       TEXT NOT NULL,
  message_id TEXT,
  payload    JSONB NOT NULL
);
CREATE INDEX items_lane_seq ON items (lane, seq);
CREATE UNIQUE INDEX items_lane_message ON items (lane, message_id) WHERE message_id IS NOT NULL;

CREATE TABLE tasks (
  task_id        TEXT PRIMARY KEY,
  lane           TEXT NOT NULL,
  seq            INTEGER NOT NULL,
  status         TEXT NOT NULL,
  message_id     TEXT,
  task           JSONB NOT NULL,
  agent          TEXT NOT NULL,
  due_at         INTEGER,
  deadline       INTEGER,
  ask            JSONB,
  status_message JSONB,
  history        JSONB NOT NULL,
  artifacts      JSONB,
  parent         TEXT,
  push_configs   JSONB,
  attempts       INTEGER NOT NULL DEFAULT 0,
  updated_at     INTEGER NOT NULL
);
CREATE INDEX tasks_lane_status_due ON tasks (lane, status, due_at);
CREATE UNIQUE INDEX tasks_lane_message ON tasks (lane, message_id) WHERE message_id IS NOT NULL;

-- One row per agent.
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

-- The tables below are private to the framework.

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
