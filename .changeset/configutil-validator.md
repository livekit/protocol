---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `configutil.Validator`, an optional builder hook run on every config load after `InitDefaults`. A config that fails validation, or whose `InitDefaults` returns an error, now fails `NewObserver`; on reload the failure is logged and the previous config stays in effect.
