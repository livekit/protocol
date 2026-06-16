---
"github.com/livekit/protocol": minor
---

agent simulation: add SimulationMode to the Create request and run info, so a
run can request audio mode. Unspecified stays TEXT, so existing callers are
unaffected.
