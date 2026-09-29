---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

signalling: allow signal messages on the WebSocket to be compressed, the way `WrappedJoinRequest` already compresses the join payload in the connect URL. Adds a `compressed` arm to the `SignalRequest` and `SignalResponse` oneofs, carrying `CompressedSignalRequest` / `CompressedSignalResponse` plus the shared `SignalCompression.Type` enum. Because the compressed form is an arm of the oneof rather than an envelope around it, no negotiation handshake is needed — the receiver always parses a `SignalRequest`/`SignalResponse` and the oneof tag says whether the payload is compressed, so even the first message can be compressed. Each direction is gated separately so old and new peers interoperate unchanged: the server compresses only for a client that advertised `ClientInfo.CAP_COMPRESSION_DEFLATE_RAW`, and the client compresses only for a server that set the new `JoinResponse.accepts_compressed_signal` / `ReconnectResponse.accepts_compressed_signal`.
