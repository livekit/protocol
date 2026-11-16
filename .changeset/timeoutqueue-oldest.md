---
"github.com/livekit/protocol": patch
---

Add `TimeoutQueue.Oldest`, when the least recently reset item was reset, so a caller can sleep until it times out instead of polling
