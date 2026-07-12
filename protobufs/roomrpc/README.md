# LiveKit participant RPCs registry (roomrpc)

These protos describe a list of well-known [LiveKit participant RPCs](https://docs.livekit.io/transport/data/rpc/)
and corresponding request/response messages.

## Directory structure

Each participant type or API should have its own directory under `roomrpc` with the name `<type>rpc`
to avoid Go package name collisions and isolate types for each API.

## Versioning

Proto files should follow the name convention: `roomrpc_<type>_v<vers>.proto`. Version is a single sequential number, not a semver.
For example: `roomrpc_sip_v1.proto`.

Corresponding `V<vers>` should be added as a suffix to the service definition and the request/response types
(method name should omit the version, as the versioned service name is part of the method path).

```protobuf
service MyServiceV1 {
  rpc DoSomething(DoSomethingV1Request) returns (DoSomethingV1Response);
}
message DoSomethingV1Request {}
message DoSomethingV1Response {}
```