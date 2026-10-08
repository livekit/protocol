---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

Agent threads on the agent runtime: `Job.thread_id`, `UpdateThreadStatus` on the worker control stream, and `AgentHttp.StreamPreamble.thread_id` / `thread_grant` (`AgentHttp.AgentThreadGrant`, with `endpoints`). `auth.AgentThreadGrant.ToProto` converts the claim for the preamble.
