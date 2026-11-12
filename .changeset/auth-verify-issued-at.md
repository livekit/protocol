---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Validate the `iat` claim when verifying access tokens (`auth.APIKeyTokenVerifier.Verify`). Previously, a correctly-signed token that omitted `nbf` and carried a far-future `iat` would verify successfully immediately, regardless of how far in the future it claimed to have been issued.
