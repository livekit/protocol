---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

`configutil.Observable` gains `ObserveCurrent`, which replays the current config to late subscribers ordered against reloads; `NewAtomic*` use it. Implementations outside this package must add the method
