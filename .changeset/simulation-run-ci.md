---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

Add `SimulationRun.CI` — the provider, commit, ref, pull request, run URL, and actor of the pipeline that started a run — to `SimulationRun.Create.Request` and to `SimulationRun`, so the dashboard can link a run back to the CI job that produced it.
