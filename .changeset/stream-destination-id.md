---
"@livekit/protocol": patch
---

`StreamDestination` (`id`, `url`) on `StreamOutput.destinations` and `UpdateStreamRequest.add_destinations`, echoed as `StreamInfo.id`, so a client can tell its streams apart; `RedactStreamKeys` redacts destination URLs
