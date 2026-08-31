---
"github.com/livekit/protocol": patch
---

`utils.Promise` gains `Then(func(T, error))`, registering a callback to run when the promise resolves. Callbacks run on the goroutine that calls `Resolve`, in registration order; registering on an already-resolved promise runs the callback inline on the caller. `NewPromise` no longer allocates the done channel up front — `Done()` already created it lazily.
