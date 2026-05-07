---
"@livekit/protocol": patch
"github.com/livekit/protocol": patch
---

Fix SIP trunk-level MediaEncryption being silently dropped on outbound and inbound calls. The early `req.Upgrade()` / `rule.Upgrade()` calls pinned `Media.Encryption` to the (legacy) request/rule field before the trunk's MediaEncryption was merged, causing INVITEs to omit SRTP when only the trunk had it configured.
