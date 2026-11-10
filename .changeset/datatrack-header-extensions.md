---
"github.com/livekit/protocol": patch
---

datatrack: size marshalled extensions from the extension list so a stale `ExtensionsSize` cannot corrupt the marshalled header.
datatrack: reset extensions on each `Unmarshal` and reject packets that repeat an extension ID.
