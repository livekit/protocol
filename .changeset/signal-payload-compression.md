---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

signalling: allow signal messages on the WebSocket to be compressed, the way `WrappedJoinRequest` already compresses the join payload in the connect URL. Adds a `compressed` arm to the `SignalRequest` and `SignalResponse` oneofs, carrying `CompressedSignalRequest` / `CompressedSignalResponse` plus the shared `SignalCompression.Type` enum. Because the compressed form is an arm of the oneof rather than an envelope around it, no negotiation handshake is needed — the receiver always parses a `SignalRequest`/`SignalResponse` and the oneof tag says whether the payload is compressed, so even the first message can be compressed. Senders use the arm only when the peer advertised `ClientInfo.CAP_COMPRESSION_DEFLATE_RAW`, so old and new peers interoperate unchanged.
