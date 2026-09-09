---
"github.com/livekit/protocol": minor
"@livekit/protocol": patch
---

Replace logger.WithTap with logger.WithTee, which duplicates every log entry to a caller-supplied zaputil.Tee. The tee's core is built from the level each derived logger resolves, so the copy follows component levels and WithMinLevel floors rather than carrying a level of its own.
