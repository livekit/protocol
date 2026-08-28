---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

signalling: add `WrappedSignalRequest` / `WrappedSignalResponse` envelopes and the `SignalCompression` enum, so signal messages on the WebSocket can be compressed the way `WrappedJoinRequest` already compresses the join payload in the connect URL. Negotiated via the existing `ClientInfo.CAP_COMPRESSION_DEFLATE_RAW` capability and acknowledged by the new `JoinResponse.signal_compression` / `ReconnectResponse.signal_compression` fields; unset means the existing uncompressed wire format is used, so old and new peers interoperate unchanged.
