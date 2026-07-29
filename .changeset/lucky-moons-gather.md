---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

agent simulation: add `RunMetrics.issues_count`, the number of issues in the run's `SimulationRunSummary`, so a run listing can show an issue count without shipping `summary_zstd`. Presence-tracked: absent means the run has not been summarized, 0 means the summary found no issues.
