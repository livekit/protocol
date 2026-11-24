---
"@livekit/protocol": patch
---

v2 auto-egress config on `RoomEgress` (`room_egress`, `participant_egress`, `track_egress`) and `RoomInternal`, where the participant and track configs carry encoding, outputs, storage and webhooks (plus `prefer_screen_share` and `capture_data` for participants), and `egress.AutoRoomRequest`/`AutoParticipantRequest`/`AutoTrackRequest`/`ValidateRoomEgress` to build egress requests from v1 or v2 config
