---
"github.com/livekit/protocol": minor
"@livekit/protocol": minor
---

`TransferSIPParticipant` now returns `TransferSIPParticipantResponse` instead of `google.protobuf.Empty`, carrying the transfer id, status, a new `SIPTransferReason` and the SIP status when the outcome came from a SIP response. `SIPTransferInfo` gains the same `reason`. This changes the generated method signature: a transfer that reports a non-successful status without an error means the transfer did not complete.
