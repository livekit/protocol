---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

signalling: add `WrappedSignalRequest` / `WrappedSignalResponse` envelopes, the `SignalCompression.Type` enum, and a `SignalCompressionAck` message on the `SignalResponse` oneof, so signal messages on the WebSocket can be compressed the way `WrappedJoinRequest` already compresses the join payload in the connect URL. Negotiated via the existing `ClientInfo.CAP_COMPRESSION_DEFLATE_RAW` capability: a server honouring it sends `SignalCompressionAck` as the first message and wraps everything after, including the `JoinResponse`. Servers and clients that do not exchange that message keep using the existing uncompressed wire format, so old and new peers interoperate unchanged.
