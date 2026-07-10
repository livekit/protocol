---
"github.com/livekit/protocol": minor
---

agent_simulations: add SimulationRun.JobEvent and GetSimulationRunEvents — a durable, append-only
per-job event stream (conversation turns, uttered-vs-heard transcripts with WER + word alignment,
playout timing, lifecycle phases) readable with a run-global cursor while a run is live or after it
ends. Heard events carry error rates computed by the simulator; clients render, never recompute.
