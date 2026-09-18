---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `utils/leaftest`, which asserts that a package imports nothing beyond the standard library and an explicit allow list. `leaftest.Assert(t)` names the package from the working directory, so a guarded package's test carries no import path to keep in sync. Test-only imports are not counted.
