---
"@livekit/protocol": patch
---

This adds a transfer_id as an easier to implement way to deduplicate transfer requests. It also adds a transfer_status to indicate if the event is for the start of end of a transfer, and if it's been successful.
