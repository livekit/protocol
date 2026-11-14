---
"github.com/livekit/protocol": patch
---

A `TimeoutQueue` iteration no longer revisits items reset during it, so resetting every item while iterating ends even when the clock has not moved
