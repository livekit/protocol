---
"github.com/livekit/protocol": patch
---

Integrate feedback on the agents protocol:

- Rename JobDescription to AgentDispatch
- Remove participant_identity entry in the dispatch
- Deprecate namespace
- Add an agent_name field to specify what agent workers a job should be dispatched to

Also allow setting a room configuration in the token.
