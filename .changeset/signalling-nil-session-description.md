---
"github.com/livekit/protocol": patch
---

`signalling.FromProtoSessionDescription` returns an empty description for a nil input, the same zero value that `ToProtoSessionDescription` maps to nil.
