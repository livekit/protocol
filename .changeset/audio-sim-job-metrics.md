---
"github.com/livekit/protocol": minor
---

agent_simulations: add JobMetrics/RunMetrics to SimulationRun — per-job quality metrics grouped by
pipeline stage (stt/llm/tts/conversation, plus a per-turn timeline and headline
accuracy/experience scores) and run-level aggregates. Unset optional fields mean "not measured";
