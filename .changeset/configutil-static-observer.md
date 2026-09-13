---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `configutil.NewStaticObserver`, which builds an `Observer` over an already-built config with no file to watch, for stubbing observable config in tests. `EmitConfigUpdate` now also stores the config it emits, so `Load` reflects an update pushed by hand.
