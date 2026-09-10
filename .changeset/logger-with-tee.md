---
"github.com/livekit/protocol": minor
"@livekit/protocol": patch
---

Replace logger.WithTap with logger.WithTee, which duplicates every log entry to a caller-supplied zaputil.Tee. The tee's core is built from the level each derived logger resolves, so the copy follows component levels rather than carrying a level of its own.

Replace ZapLogger.WithMinLevel with WithComponentLeveler, which attaches a zaputil.ComponentLeveler to a branch of the logger tree. A leveler owns the per-component level and write-enabler cache for one configuration source and can only widen its parent, so a caller can resolve levels per (tenant, component) without rebuilding loggers on a config change. ZapLogger.Leveler exposes the leveler a branch resolves through, for use as the parent of a derived one.
