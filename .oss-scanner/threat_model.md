# LiveKit protocol threat model

`github.com/livekit/protocol` is a Go library shared by the LiveKit server, the
Go server SDK, and LiveKit's media services. It holds the protobuf definitions
and generated code, plus the Go helpers that interpret untrusted input on their
behalf: access-token and webhook verification, SDP and RTP helpers, data-track
packet parsing, and decode budgets for protobuf and JSON.

## Scope and trust boundaries

Audit the hand-written Go on `main`. Prioritise:

- `auth/`: access-token minting, parsing, and verification, and grant
  semantics (`ClaimGrants`, `VideoGrant`, `SIPGrant`, agent and inference
  grants).
- `webhook/`: `Receive` / `ReceiveWebhookEvent` (signature and body-hash
  verification) and the outbound notifier.
- `utils/jwtutil/`, `utils/decodebudget/`: JWT helpers and resource limits for
  decoding untrusted protobuf and JSON.
- `datatrack/`, `sdp/`, `utils/rtputil/`, `codecs/`: parsers for packets and
  SDP that arrive from clients.
- `sip/`: SIP token building, phone-number normalisation, and trunk and
  dispatch-rule matching.
- `utils/` helpers that redact secrets or handle IDs.

Out of scope: generated code (`*.pb.go`, `*.twirp.go`, `*.psrpc.go`), the
`.proto` files themselves, the JS packages under `packages/`, and build tooling
under `magefile.go` and `tools/`. Report a bug in a generated file only when
the generator template or the `.proto` definition causes it.

Untrusted input: access tokens and webhook requests from the network, SDP,
RTP/RTCP, and data-track packets from clients, and any protobuf or JSON a
client can send that these packages decode.

Trusted: API secrets, the caller's configuration, and values the embedding
application passes in after it has authenticated them.

LiveKit server (`livekit/livekit`) is enrolled separately. Report bugs in how
the server uses these packages there; report bugs in the packages here.

## Security properties

- **Token verification:** a token must verify only with its issuer's secret,
  only with HS256, and only inside its validity window. Parsing an unverified
  token must never be treated as verification. Grants must deserialize exactly
  as signed; a token must not gain a grant, room, or identity it was not issued.
- **Webhook verification:** an event must be accepted only when its token
  verifies and the body's SHA-256 matches the signed hash.
- **Parsing:** SDP, RTP extensions, data-track packets, and decode-budget
  inputs must not panic, loop, or allocate without bound on malformed input.
- **Decode budgets:** the budget must bound the memory and CPU of decoding the
  message it is applied to.
- **Secrets:** API secrets must not appear in logs, errors, or redacted output.

## Severity guidance

- **Critical:** forging a token or webhook that verifies without the secret.
- **High:** accepting an expired, wrong-algorithm, or wrong-issuer token;
  grant escalation through parsing; a panic reachable from untrusted network
  input that crashes the embedding process.
- **Medium:** unbounded resource use from one untrusted input; a decode budget
  that undercounts by a large factor; secret leakage into logs.
- **Low / not a finding:** misuse that needs the API secret or caller
  configuration; issues only in generated code; outbound webhook delivery
  behaviour (retries, ordering) without a security impact.

## Testing

The checkout is at `/src`. Go dependencies are cached. `GOPROXY=off` and
`GOSUMDB=off` prevent downloads. No external services are needed.

```sh
cd /src
go test -count=1 ./...
go test -race -count=1 ./auth/... ./webhook/... ./datatrack/... ./sdp/...
go test -run '^$' -fuzz FuzzPacketUnmarshal -fuzztime 60s ./datatrack
go test -run '^$' -fuzz FuzzSDPFragmentUnmarshal -fuzztime 60s ./sdp
```

## Reports

State the attacker's required access (none, any valid token, or a specific
grant), the entry point, and the demonstrated impact. Provide a minimal
reproducer, preferably a Go test in the affected package.
