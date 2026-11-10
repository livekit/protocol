---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Update the existing v1 agent thread schema in place: partition-keyed history and ownership,
optional task message links, endpoint destinations and journal retry identities. Use task
insertion positions for FIFO acceptance, store task status messages in data, and index task
update listings and terminal run receipts. Keep the v1 schema filename, exports and version.
