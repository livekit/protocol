---
"@livekit/protocol": patch
---

Add `CreateJobTokenRequest` and `CreateJobTokenResponse` to the agent worker protocol, for a job to ask the server for an access token

Add `agent.ParseJobTokenClaims` and `agent.BuildJobToken`, which read the claims a job asks for strictly and sign a 10-minute job token; neither accepts `sha256` or `video.agent`
