---
"github.com/livekit/protocol": patch
---

webhook: export `RedactSendError` for a host whose webhook owners are tenants. It gives the same text for a destination the dial guard refuses and one that does not resolve, and in any other error replaces each address in the private and reserved ranges that `BlockInternal` refuses. `SendError` itself is unchanged. `NewDialControl` is exported so a caller can add blocked and allowed ranges. The dial guard refuses any zoned address that is not explicitly allowed. An IPv4-mapped range in `blocked_cidrs` or `allowed_cidrs` matches as its IPv4 range, and one shorter than 96 bits is rejected. A blocked range also catches the IPv4 address that a NAT64, v4-compatible or 6to4 address embeds
