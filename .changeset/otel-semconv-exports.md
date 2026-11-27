---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

Export the Weaver-generated OTEL telemetry semantic conventions. `@livekit/protocol/semconv` now exposes the TypeScript constants as a dependency-free subpath, and a new Go package `github.com/livekit/protocol/otel/gen/go` exposes the Go constants (attribute keys, enum values, events, spans, checkpoints, metrics) plus `RenamedKeys` and `RenamedValues`, the legacy-key and legacy-value alias maps derived from the registry's `renamed_from` annotations for normalizing older captures.
