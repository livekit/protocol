---
"github.com/livekit/protocol": patch
---

`VideoGrant.UpdateFromPermission` and `MatchesPermission` no longer read `ParticipantPermission.agent`, so a permission update does not change the agent grant.
`RoomConfiguration.CheckCredentials` checks every egress output type, including image outputs, the deprecated single room outputs and upload proxy passwords.
