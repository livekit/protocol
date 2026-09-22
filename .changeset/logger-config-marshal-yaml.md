---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

`logger.Config` implements `yaml.Marshaler`, so marshaling a `*Config` snapshots it under the same lock `Update` takes instead of reading fields alongside a concurrent update.
