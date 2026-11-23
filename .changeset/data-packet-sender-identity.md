---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `data_stream_actual_sender_overwritten`, `data_stream_sending_on_behalf` and `data_stream_actual_sender_identity` to `DataPacket`, which the SFU sets so receivers know the real sender of a packet an agent sent on behalf of another participant
