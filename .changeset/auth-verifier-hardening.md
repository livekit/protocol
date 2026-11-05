---
"github.com/livekit/protocol": patch
---

`auth.ParseAPIToken` reads only `alg`, `iss`, `sub` and `jti`, and `APIKeyTokenVerifier.Verify` checks the signature before it decodes the header and grants, so an unsigned token costs no more than its size to reject.
`Verify` rejects an empty key of any type with `ErrKeysMissing`, and rejects a token whose `iat` is more than one minute in the future.
