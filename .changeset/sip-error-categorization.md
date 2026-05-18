---
"github.com/livekit/protocol": patch
---

Wrap SIP validation errors with typed psrpc codes (`InvalidArgument` for unsupported dispatch rule and inbound trunk auth, `FailedPrecondition` for outbound trunk missing numbers) so Twirp surfaces them as 4xx instead of 500. Also fix `SIPStatus.GRPCStatus()` fallback ladder which compared the wrong variable.
