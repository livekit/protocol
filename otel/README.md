# LiveKit telemetry semantic conventions

The machine-readable registry of the attribute keys, events and spans LiveKit producers send over
OTLP, in [OpenTelemetry Weaver](https://github.com/open-telemetry/weaver) format. It depends on
[OTel semantic conventions v1.44.0](https://github.com/open-telemetry/semantic-conventions/tree/v1.44.0/model):
upstream keys (`service.*`, `os.*`, `device.*`, `network.connection.type`, `error.type`,
`session.id`, `code.*`) are referenced with `ref:`, never redefined.

Status: prototype, `development` stability. It registers the shared identity, PII and SDK keys,
and what the client core (rust-sdks `livekit-telemetry`) emits today, under its current names.
Agents-owned keys are not registered yet.

| Path | What |
|---|---|
| `model/manifest.yaml` | Registry version (`schema_url`) and the pinned OTel dependency |
| `model/shared/registry.yaml` | Shared keys: identity (`lk.room_sid`, `lk.participant_id`, `lk.track_sid`, `lk.job_id`), PII (`lk.pii.*`), SDK identity (`lk.sdk.*`) |
| `model/client/registry.yaml` | Client-owned attributes (`lk.device.*`, `lk.rtc.*`, `lk.track.*`, `lk.log.*`, `lk.telemetry.*`, …) |
| `model/client/signals.yaml` | Client log events and spans |
| `model/client/groups.yaml` | Attribute sets on every client resource / record (used by live-check) |
| `policies/` | LiveKit rules; any finding fails `check.sh` |
| `advice/` | Non-failing advice, printed by `check.sh` |
| `templates/registry/{rust,typescript,markdown}/` | Codegen templates |
| `gen/` | Generated outputs, committed: `rust/semconv.rs`, `ts/semconv.ts`, `docs/semconv.md` |
| `live-check/` | Live-check config, advice rules and the OTLP/JSON converter |
| `tests/` | Policy negatives, live-check smoke sample, schema-version rule |

The generated docs, [`gen/docs/semconv.md`](gen/docs/semconv.md), list every key with its type,
owner, sensitivity and unit.

## Ownership

Every LiveKit key carries `annotations.livekit.owner`:

- **`shared`** (`model/shared/`): concepts several producers send. A change needs Client, Agents and
  Observability review.
- **`client`** (`model/client/`): the Client team approves alone, unless the change touches or
  duplicates a shared concept.
- `agents` and `cloud` are valid owners for when those teams register their keys, in their own
  directories.

The Notion conventions put client-owned attributes under `lk.client.*`. The client core does not do
that yet, so that rule is advice: `check.sh` lists every client key it would move, and never
fails on it.

How a concept is named in other producers or stores (Agents' `room_id`, the collector's `room_id`
column, roomobs) is recorded in `annotations.livekit.bindings`. Weaver carries it into the docs;
nothing validates it.

## Adding a key

1. Pick the file by owner. Define keys in the v1 `groups:` files. `client/groups.yaml` is
   `definition/2` only because live-check needs public attribute groups; it holds references.
2. Reuse an upstream OTel key when one fits, with `ref:`.
3. Give the key a `brief`, `stability: development`, `examples` and
   `annotations.livekit`:

   | Annotation | Values | Rule |
   |---|---|---|
   | `owner` | `shared`, `client`, `agents`, `cloud` | required |
   | `sensitivity` | `none`, `pii`, `secret` (protocol `logger.Sensitivity`) | required; `pii` iff the key has a `pii` segment; `secret` is never emitted |
   | `unit` | UCUM, e.g. `s`, `By`, `{packet}` | durations are seconds |
   | `unit_exception` | a reason | allows a `*_ms` key with `unit: ms` (RTC stats keep webrtc-stats milliseconds) |
   | `bindings` | map of producer/store → name | informational |
   | `cardinality_exception` | `true`, on a metric | allows a high-cardinality ID as a metric dimension |

   Naming: shared identity keys are flat (`lk.room_sid`, never `lk.room.sid`). Customer-supplied
   values go under `lk.pii.*`: the collector deletes keys with a `pii` segment when redaction is
   on, and keeps everything else.
4. Reference it from the events or spans that carry it, with a `requirement_level`.
5. Run `otel/check.sh`, then `otel/generate.sh`, and commit `gen/` with the change.
6. Bump the schema version (below).

## Bumping the schema version

`schema_url` in `model/manifest.yaml` is the telemetry schema version. Producers set it as the
OTLP scope `schema_url` (generated as `SCHEMA_URL`), which the collector already stores as a
column. It is versioned independently of the protocol package.

Every PR that changes `model/` bumps it; CI fails otherwise (`tests/schema-bump.sh`). While the
registry is `0.x`, bump the minor version for added, renamed or removed keys, events or spans, and
the patch version for documentation-only changes. To rename a key that has shipped, keep the old
key with `deprecated: {reason: renamed, renamed_to: <new key>}` instead of deleting it;
`weaver registry diff` and live-check both report the rename.

## Running the checks locally

Install weaver 0.28.0, the version CI pins, from its
[release](https://github.com/open-telemetry/weaver/releases/tag/v0.28.0) (`weaver-installer.sh`),
or put the binary anywhere and point `WEAVER` at it. Every run clones OTel semconv v1.44.0 and
prints 34 warnings: 33 about upstream's own files, and one because `client/groups.yaml` uses the
not-yet-stable `definition/2` format.

```sh
otel/check.sh                          # registry + LiveKit policies, then advice
otel/tests/negatives.sh                # each negative fixture fails with its named finding
otel/generate.sh && git status otel/gen   # regenerate; no diff expected
otel/tests/live-check.sh               # smoke sample clean, legacy keys flagged
otel/tests/schema-bump.sh origin/main  # schema_url bumped if model/ changed

rustc --edition 2021 --crate-type lib -D warnings -D missing_docs otel/gen/rust/semconv.rs --out-dir /tmp
npx -p typescript@5.9.3 tsc --strict --noEmit --target es2020 otel/gen/ts/semconv.ts
```

| Finding | Fails | Rule |
|---|---|---|
| `livekit_owner`, `livekit_sensitivity` | yes | annotations present and valid |
| `livekit_secret` | yes | `secret` is never emitted |
| `livekit_pii_marker` | yes | `sensitivity: pii` iff a `pii` key segment |
| `livekit_pii_heuristic` | yes | identity, email, phone, room/participant name outside `lk.pii.*` |
| `livekit_flat_identity` | yes | no `lk.room.sid`, `lk.participant.sid`, `lk.track.sid`, `lk.job.id` |
| `livekit_seconds`, `livekit_unit`, `livekit_metric_unit` | yes | seconds, unless `unit_exception` |
| `livekit_cardinality` | yes | high-cardinality IDs are not metric dimensions |
| `livekit_client_namespace` | no (advice) | client-owned keys under `lk.client.*` |

Weaver itself rejects duplicate keys, unknown fields and unresolved references.

## Live-check

`live-check.sh` checks captured telemetry against the registry:

```sh
otel/live-check.sh session.jsonl [out-dir] [--fail-on none]
```

It takes OTLP/JSON lines from an otelcol `file` exporter (the SDK e2e setup), or a weaver samples
JSON file. `live-check/weaver.toml` matches spans by name and checks the resource and every
record against `model/client/groups.yaml`. It drops two kinds of finding: `not_stable` (every key
is `development`) and absent `conditionally_required` attributes (their conditions are prose).
`live-check/livekit.rego` adds `livekit_flat_identity` and `livekit_pii_unmarked`, each naming
the key to send instead.

On the reference session (Swift 2.17.0 with the core at rust-sdks `a9ec8d52`: 6 spans, 22 logs,
349 attributes), every finding is one of the expected shared identity / PII renames:

| Emitted today | Registered key | Findings |
|---|---|---|
| `lk.room.sid` | `lk.room_sid` | `missing_attribute`, `unexpected_attribute`, `livekit_flat_identity` ×10 each |
| `lk.participant.sid` | `lk.participant_id` | same, ×10 each |
| `lk.track.sid` | `lk.track_sid` | same, ×12 each, plus `required_attribute_not_present: lk.track_sid` ×10 |
| `lk.room.name` | `lk.pii.room_name` | `missing_attribute`, `unexpected_attribute`, `livekit_pii_unmarked` ×10 each |
| `lk.participant.identity` | `lk.pii.participant_identity` | same, ×10 each |
| `lk.participant.remote_identity` | `lk.pii.remote_participant_identity` (name TBD) | same, ×2 each, plus `recommended_attribute_not_present` ×2 |

With those six keys renamed in the same capture, live-check reports nothing and exits 0.
`tests/live-sample.json` is part of that capture, anonymized.

Known limits:

- Span event names (connect checkpoints, `attempt <n> <mode>`) are not checked; weaver compares
  span events only by attributes.
- App-defined attributes (custom events, `Scope::set_attribute`) are reported as unknown keys.
- `lk.sdk.*` is not reported when missing: resource attributes are only compared, not required.
- Units are annotations; live-check checks names, types and enum values, never magnitudes.
