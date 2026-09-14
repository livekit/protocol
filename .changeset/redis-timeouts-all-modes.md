---
"github.com/livekit/protocol": patch
---

Honour `dial_timeout`, `read_timeout` and `write_timeout` for single-address and cluster Redis, not just Sentinel. They were previously accepted by `RedisConfig` but only ever read in the Sentinel branch, so setting them on any other deployment silently did nothing.
