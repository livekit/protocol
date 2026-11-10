---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Update the existing v1 agent thread schema in place: partition-keyed history and ownership,
optional task message links, endpoint destinations and journal retry identities. Use task
insertion positions for FIFO acceptance, store task status messages in data, and index task
update listings and terminal run receipts. Keep the v1 schema filename, exports and version.

Define DefaultPartitionKey for framework default resolution and require every SQL write to supply an explicit partition key.

Standardize stored timestamps and lease arithmetic on Unix milliseconds, using explicit _ms field names.

Remove the cross-worker interruption marker from partition lease rows. New requests wait for ownership without signaling the active turn to stop.
