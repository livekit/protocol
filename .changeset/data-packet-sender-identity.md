---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add a `sender` oneof (`sender_identity` / `sender_hidden`) to `DataPacket`, which the SFU sets to the real sender when an agent sends a packet on behalf of another participant
