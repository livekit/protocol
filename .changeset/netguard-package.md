---
"github.com/livekit/protocol": patch
---

Add `utils/netguard`, a check that refuses internal and reserved addresses, for clients that connect to destinations a caller supplies: `IsBlockedIP`, `CheckHost`, `BlockInternal`, `NewDialControl`, `ParseCIDR`, and `ReportOnly` and `ReportOnlyContext`, which report what a check would refuse and let the connection go ahead. Existing packages are unchanged.
