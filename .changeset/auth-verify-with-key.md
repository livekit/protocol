---
"github.com/livekit/protocol": patch
---

auth: `VerifyWithKey` verifies a token with the algorithm its key picks, HS256 for a version 1 key and EdDSA for a version 2 key, and `NewAccessTokenWithPrivateKey` signs a token with EdDSA. `ParseAPIToken` accepts EdDSA tokens; `Verify` stays HS256 only. A version 2 token names its API key in the `kid` header and has the `typ` header `livekit-v2+jwt`, and `VerifyWithKey` refuses one that does not. Version 2 keys are not supported by any server yet, so do not use them
