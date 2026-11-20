---
"github.com/livekit/protocol": minor
---

logger: add an opt-in log/slog backend (`WithSlogBackend`, byte-identical output) and the slog migration API: `ToSlog`/`FromSlog`, `*slog.Logger` helpers (`WithName`, `WithComponent`, `WithDeferredValues`, `WithItemSampler`, `WithoutSampler`), `Err`, `Float32`, `LogDepth`, `WithTeeHandler`, `WithZapBackend`, `ComponentLevels`, and `ObjectEncoder`/`ObjectMarshaler` types to implement instead of zapcore's; deprecate `ZapLogger.ToZap`
