---
"@livekit/protocol": patch
---

agent_session: rename `CustomEvent` → `DebugMessage`; drop `type` field, keep only `payload`

Renames the agent-session event added in #1588 before any consumer ships it. The
message is repositioned as an internal debug/trace channel surfaced only to the
debugger/recorder, not to user code, so the `type` discriminator was unnecessary
— callers just emit a free-form JSON `payload`.

Wire-level: `AgentSessionEvent.custom_event` (field 21) → `AgentSessionEvent.debug_message`
(same field number 21, same type slot, no schema-compat concerns since nothing
has been built against it yet downstream).
