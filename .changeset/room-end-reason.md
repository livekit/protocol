---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

add `RoomEndReason` enum and report it on room-ended telemetry: `AnalyticsEvent.room_end_reason` and `WebhookEvent.room_end_reason` now say why a room ended (`ROOM_END_API_DELETE`, `ROOM_END_IDLE_TIMEOUT`, `ROOM_END_SERVER_SHUTDOWN`, `ROOM_END_SUPERSEDED`, `ROOM_END_OPEN_FAILED`), so room-end volume can be broken down by cause the way participant disconnects already are
