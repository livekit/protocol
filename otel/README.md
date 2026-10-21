# LiveKit telemetry semantic conventions

The machine-readable registry of the attribute keys, events and spans LiveKit producers send over
OTLP, in [OpenTelemetry Weaver](https://github.com/open-telemetry/weaver) format. It depends on
[OTel semantic conventions v1.44.0](https://github.com/open-telemetry/semantic-conventions/tree/v1.44.0/model):
upstream keys (`service.*`, `os.*`, `device.*`, `network.connection.type`, `error.type`,
`session.id`, `code.*`) are referenced with `ref:`, never redefined.

Status: prototype, `development` stability. It registers the shared identity, PII and SDK keys,
what the client cores (rust-sdks `livekit-telemetry` and the JS port) emit today, and what
LiveKit Agents 2.0 (Python and the Rust core) emits today, all under their current names. The
OTel GenAI conventions are a second dependency
([semantic-conventions-genai](https://github.com/open-telemetry/semantic-conventions-genai),
pinned by commit: it has no tags), so Agents' `gen_ai.*` keys are referenced, not redefined.

The registry is the spec of these conventions: names, types, units, requirement levels, meaning,
event severity and body, span kind, checkpoints and outcome rules. Transport, cadence, upload
policy, storage and the app API are not conventions and live with each SDK.

| Path | What |
|---|---|
| `model/manifest.yaml` | The pinned OTel dependencies, and a `schema_url` that points at this registry (see [Versioning](#versioning)) |
| `model/shared/registry.yaml` | Shared keys: identity (`lk.room_sid`, `lk.participant_id`, `lk.track_sid`, `lk.job_id`), PII (`lk.pii.*`), SDK identity (`lk.sdk.*`), and the concepts clients and Agents both send (`lk.track_source`, `lk.disconnect_reason`, `lk.agent_name`) |
| `model/client/registry.yaml` | Client-owned attributes (`lk.device.*`, `lk.rtc.*`, `lk.track.*`, `lk.log.*`, `lk.telemetry.*`, …) |
| `model/client/signals.yaml` | Client log events and spans |
| `model/client/groups.yaml` | Attribute sets on every client resource / record (used by live-check) |
| `model/agents/registry.yaml` | Agents-owned attributes (`lk.speech_id`, `lk.eou.*`, `lk.rpc.*`, `lk.blocking.*`, …), the `lk.agents.*` metric dimensions, the `gen_ai.*`/LangFuse keys no upstream registry defines, and the session-report keys |
| `model/agents/signals.yaml` | Agents spans (50, unprefixed as shipped), the 20 `lk.agents.*` metrics, and the GenAI span events and metrics that are deprecated upstream or cannot be imported |
| `model/agents/groups.yaml` | GenAI metrics imported from the GenAI registry, and the values Agents sends for upstream keys that the upstream definitions do not list |
| `policies/` | LiveKit rules (`lib.rego`), applied per file (`livekit.rego`, default mode) and to the resolved registry (`resolved.rego`, default and `--v2`); any finding fails `check.sh` and `generate.sh` |
| `advice/` | Non-failing advice, printed by `check.sh` |
| `templates/registry/{rust,typescript,python,markdown}/` | Codegen templates |
| `gen/` | Generated outputs, committed: `rust/semconv.rs`, `ts/semconv.ts`, `python/semconv.py`, `docs/semconv.md` |
| `live-check/` | Live-check config, advice rules and the OTLP/JSON converter |
| `tests/` | Policy negatives, live-check samples, codegen escaping |

The generated docs, [`gen/docs/semconv.md`](gen/docs/semconv.md), are the human-readable spec:
naming rules, the Resource, the attributes on every record, each event and span with its
requirement table, enum values and bindings.

## Ownership

Every LiveKit key carries `annotations.livekit.owner`:

- **`shared`** (`model/shared/`): concepts several producers send. A change needs Client, Agents and
  Observability review.
- **`client`** (`model/client/`): the Client team approves alone, unless the change touches or
  duplicates a shared concept.
- **`agents`** (`model/agents/`): the Agents team approves alone, unless the change touches or
  duplicates a shared concept. Registered as emitted.

Where a client and Agents named the same concept differently, the Agents spelling and value set
is the shared key, and the client is the side that changes: `lk.track_source` and
`lk.disconnect_reason` take the enum names Agents sends as values (`SOURCE_MICROPHONE`,
`CLIENT_INITIATED`), and `lk.agent_name` is the key a client uses for the agent. The track
source names are the Agents RTC FFI `TrackSource` names, not the signal protocol's (`CAMERA`,
`SCREEN_SHARE`), so clients map them explicitly.

All eight client renames (the six identity/PII keys, `lk.track.source` and
`lk.disconnect.reason`) are recorded the same way: `renamed_from` on the shared key names the old
spelling and, for an enum, maps each old value to a member. The old spellings are never
registered, deprecated or not. A policy rejects defining one, and live-check reads the map to
name the key and value to send.
- `cloud` is a valid owner for when that team registers its keys, in its own directory.

The Notion conventions put client-owned attributes under `lk.client.*`. The client core does not do
that yet, so that rule is advice: `check.sh` lists every client key it would move, and never
fails on it.

How a concept is named in other producers or stores (Agents' `room_id`, the collector's `room_id`
column, roomobs) is recorded in `annotations.livekit.bindings`. Weaver carries it into the docs;
nothing validates it.

Agents' session report sends plain log records whose keys predate this registry. They are
registered as emitted, owner `agents`: `session.options`, `session.report_timestamp` and
`session.tags` sit in the upstream `session.*` namespace and carry
`annotations.livekit.upstream_namespace: session`, and `agent_name`, `sdk_version`, `usage`,
`chat.item`, `evaluation`, `tag` and `outcome` have no namespace. A map-valued key carries an OTLP
map, and its fields (such as `reason` in `outcome`) are payload, not attribute keys.

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
   | `unit_exception` | a reason | allows a `*_ms` key with `unit: ms` (the RTC stats keys: `lk.rtc.window_ms` as measured, the others converted from the W3C stats' seconds) |
   | `bindings` | map of producer/store → name | informational |
   | `cardinality_exception` | `true`, on a metric | allows a high-cardinality ID as a metric dimension |
   | `renamed_from` | `{key, values}`, on a shared key | the spelling it replaces, and for an enum a map of old value → member; the old key stays unregistered |
   | `pii_heuristic_exception` | a reason | documents a shipped key the PII heuristic would flag; it only takes effect for a key listed in `pii_heuristic_allowlist` in `policies/lib.rego` (today `lk.rpc.caller_identity`, `lk.rpc.destination_identity`), so a new key cannot opt out by annotating itself. The collector still does not redact them |

   Events and spans carry their own `annotations.livekit`, rendered in the docs; live-check
   enforces `severity`, `checkpoints` and `checkpoint_pattern`:

   | Annotation | On | Meaning |
   |---|---|---|
   | `severity` | event | OTLP severity texts the event is sent with, e.g. `[INFO]` |
   | `body` | event | What the log body carries |
   | `reserved`, `limits` | `custom` event | Keys and prefixes apps may not use; size caps on app events and attributes |
   | `checkpoints` | span | Span events in order, each `{name, requirement: required\|best_effort, brief}` |
   | `checkpoint_pattern` | span | A dynamic checkpoint name: `{name, pattern, brief}` |
   | `limits` | span | Checkpoints and attributes per span; the rest is counted in OTLP `dropped_*_count` |

   Give every key, event and span a `brief` and a `note`, and every non-enum key `examples`; a
   `ref` to an OTel key gets a short `note` saying how LiveKit uses it.

   Naming: shared identity keys are flat (`lk.room_sid`, never `lk.room.sid`). Customer-supplied
   values go under `lk.pii.*`: the collector deletes keys with a `pii` segment when redaction is
   on, and keeps everything else.
4. Reference it from the events or spans that carry it, with a `requirement_level`.
5. Run `otel/check.sh`, then `go tool mage otel` (it runs `otel/generate.sh`), and commit `gen/`
   with the change.

## Versioning

The registry is unversioned until a producer needs to emit `schema_url`. Weaver requires one, so
`model/manifest.yaml` sets `https://github.com/livekit/protocol/tree/main/otel`, which points at
this registry. Weaver reads its last path segment, `otel`, as the version. No producer emits it,
and the link resolves once this directory is on main. When a producer needs a version, the plan
is to stamp the protocol release version into the manifest in the release step, as OTel
semantic-conventions does.

## Running the checks locally

CI (`.github/workflows/otel.yaml`) runs the same scripts on pull requests and pushes to `main`
that touch `otel/**` or the workflow file itself, and on manual dispatch.

`go tool mage otel` regenerates `otel/gen`, next to `go tool mage proto` for the protobufs. Only a
change under `otel/` needs Weaver: `mage proto` and the JS generation do not. As for the
protobufs, the Generate workflow (`.github/workflows/generate.yaml`) regenerates `otel/gen` on
every push to a branch other than `main` and commits any change. That commit is pushed with the
workflow token, so it does not run the checks by itself (GitHub can ask a maintainer to approve
that run), the same as for the regenerated protobuf code. After a regeneration commit, push
again to run the checks.

Install weaver 0.28.0, the version CI pins, from its
[release](https://github.com/open-telemetry/weaver/releases/tag/v0.28.0) (`weaver-installer.sh`),
or put the binary anywhere and point `WEAVER` at it. Every run clones OTel semconv v1.44.0 and
the GenAI registry. It prints about 90 warnings: almost all about upstream's own files, and one
for each LiveKit file in the not-yet-stable `definition/2` format (`client/groups.yaml`,
`agents/groups.yaml`).

```sh
otel/check.sh                          # registry + LiveKit policies, default and --v2, then advice
otel/tests/negatives.sh                # each negative fixture fails with its named finding, both modes
go tool mage otel && git status otel/gen  # regenerate (runs otel/generate.sh); no diff expected
otel/tests/live-check.sh               # smoke and sparse samples clean, legacy keys flagged
otel/tests/generate.sh                 # generate.sh replaces gen/ whole, and keeps it when a target fails
otel/tests/codegen.sh                  # a value with quotes, backslashes, newlines round-trips through rustc, tsc and python3

rustc --edition 2021 --crate-type lib -D warnings -D missing_docs otel/gen/rust/semconv.rs --out-dir /tmp
npx -p typescript@5.9.3 tsc --strict --noEmit --target es2020 otel/gen/ts/semconv.ts
python3 -I -m py_compile otel/gen/python/semconv.py
```

The generated Rust and Python also carry an `upstream` section: every OTel key a LiveKit signal
references (`error.type`, `gen_ai.request.model`, `rpc.method`, …), defined by a dependency and
never here, so a producer names them from one place. TypeScript does not yet. Three mechanisms
add to it:

- Imported metrics. `model/agents/groups.yaml` imports `gen_ai.invoke_agent.duration` and
  `gen_ai.execute_tool.duration` from the GenAI registry, and the upstream section lists them with
  their units. Weaver 0.28.0 refuses to import `gen_ai.client.operation.duration`, because semconv
  v1.44.0 keeps a deprecated copy excluded from dependency resolution. `model/agents/signals.yaml`
  declares that metric as emitted instead.
- Values beyond upstream. A ref to an upstream key can list `annotations.livekit.shipped_values`:
  values a producer sends that the upstream definition does not list, for example
  `gen_ai.provider.name = mistral_ai`. Each value gets a constant next to the upstream values
  (`GEN_AI_PROVIDER_NAME_MISTRAL_AI`). Two rules keep those constants unique.
  A shipped `id` must be lowercase snake_case (`^[a-z][a-z0-9_]*$`), so its uppercase form is
  exactly the constant name, and it must differ from every upstream member id read the same way
  (`aws.bedrock` as `aws_bedrock`).
  At most one shipped_values list exists per upstream key in the whole registry, because codegen
  merges every ref's list into one set of constants.
  `livekit_shipped_value` enforces both in both modes. The CI compile of the generated Rust is
  the backstop for any collision the policy misses: a repeated constant does not compile (E0428).
- Rust `log` keys. In a `log` key-value list, wrap a generated constant in parentheses:
  `log::warn!(target: "livekit.agents", (semconv::LK_PII_ERROR) = error; "...")`. A bare
  identifier also compiles, but `log` sends the identifier's name (`LK_PII_ERROR`) as the key.

| Finding | Fails | Rule |
|---|---|---|
| `livekit_owner`, `livekit_sensitivity` | yes | annotations present and valid |
| `livekit_secret` | yes | `secret` is never emitted |
| `livekit_pii_marker` | yes | `sensitivity: pii` iff a `pii` key segment |
| `livekit_pii_heuristic` | yes | identity, email, phone, room/participant name outside `lk.pii.*`, unless the key is on the policy's `pii_heuristic_allowlist` and carries a `pii_heuristic_exception` reason |
| `livekit_flat_identity` | yes | no `lk.room.sid`, `lk.participant.sid`, `lk.track.sid`, `lk.job.id` |
| `livekit_shared_rename` | yes | no old spelling named in a `renamed_from` is defined, deprecated or not (resolved, both modes) |
| `livekit_rename_value` | yes | every value a `renamed_from` maps to is a member of the shared key's enum |
| `livekit_seconds`, `livekit_unit` | yes | seconds, unless `unit_exception` |
| `livekit_metric_unit` | yes | metric units are not `ms` (resolved, so `definition/2` metrics and refinements count) |
| `livekit_cardinality` | yes | high-cardinality IDs are not metric dimensions (resolved, so `extends` and refinements count) |
| `livekit_ref_override` | yes | a `ref` to an `lk.*` key may not override its `annotations` (owner, sensitivity, unit) |
| `livekit_shipped_value` | yes | a `shipped_values` list is on a non-`lk.*` key of type string, string[] or an enum whose values are all strings; every entry has a lowercase snake_case id, a value and a brief; no value repeats an upstream value; no id repeats another entry or an upstream member id (lowercased, non-alphanumerics as `_`); and each key has at most one distinct list in the whole registry. Refs and definitions, resolved, both modes |
| `livekit_client_namespace` | no (advice) | client-owned keys under `lk.client.*` |

Weaver itself rejects duplicate keys, unknown fields and unresolved references. Every rule runs in
weaver's default mode and with `--v2`: weaver skips per-file (`before_resolution`) policies with
`--v2`, so `resolved.rego` applies the same rules (from `lib.rego`) to the resolved registry there.

## Live-check

`live-check.sh` checks captured telemetry against the registry:

```sh
otel/live-check.sh session.jsonl [out-dir] [--fail-on none]
```

It takes OTLP/JSON lines from an otelcol `file` exporter (the SDK e2e setup), or a weaver samples
JSON file. `live-check/weaver.toml` matches spans by name and checks the resource and every
record against `model/client/groups.yaml`. It drops two kinds of finding: `not_stable` (every key
is `development`) and absent `conditionally_required` attributes (their conditions are prose).
`live-check/livekit.rego` adds `livekit_shared_rename` for an old spelling. `live-check/renames.jq`
builds its map from the model's `renamed_from`, so each finding names the key to send and, for an
enum, the value. It also adds `livekit_pii_unmarked` for any other PII-looking key, and
`livekit_unknown_span` for an `lk.*` span no matcher places (weaver
alone reports nothing for an unmatched span that has no attributes). From the event and span
annotations it adds `livekit_severity` (an event at a severity it is not declared with),
`livekit_checkpoint` (a span event the span does not list) and `livekit_span_outcome` (span status
disagreeing with `lk.outcome`, or `error.type` without `lk.outcome = error`).

On the reference session (Swift 2.17.0 with the core at rust-sdks `a9ec8d52`: 6 spans, 22 logs,
349 attributes), every finding is one of the eight expected renames to shared keys (136
violations, 62 improvements):

| Emitted today | Registered key | Findings |
|---|---|---|
| `lk.room.sid` | `lk.room_sid` | `missing_attribute`, `unexpected_attribute`, `livekit_shared_rename` ×10 each |
| `lk.participant.sid` | `lk.participant_id` | same, ×10 each |
| `lk.track.sid` | `lk.track_sid` | same, ×12 each, plus `required_attribute_not_present: lk.track_sid` ×10 |
| `lk.room.name` | `lk.pii.room_name` | `missing_attribute`, `unexpected_attribute`, `livekit_shared_rename` ×10 each |
| `lk.participant.identity` | `lk.pii.participant_identity` | same, ×10 each |
| `lk.participant.remote_identity` | `lk.pii.remote_participant_identity` | same, ×2 each, plus `recommended_attribute_not_present` ×2 |
| `lk.track.source` = `camera`, `microphone` | `lk.track_source` = `SOURCE_CAMERA`, `SOURCE_MICROPHONE` | `missing_attribute`, `unexpected_attribute`, `livekit_shared_rename`, `required_attribute_not_present: lk.track_source` ×4 each |
| `lk.disconnect.reason` = `client_initiated` | `lk.disconnect_reason` = `CLIENT_INITIATED` | same, ×2 each |

With those eight keys renamed (and the two values mapped) in the same capture, live-check reports
nothing and exits 0. `tests/live-sample.json` is part of that capture, anonymized, and
`tests/live-agents.json` holds the shared keys as Agents sends them; both must pass.

Known limits:

- Span names are checked only in the `lk.*` namespace. A span outside it (an app's own
  `Span::detached` name) passes without a check; with no attributes it gets no finding at all.
  A registered span also needs its matcher in `live-check/weaver.toml`, or it is reported as
  `livekit_unknown_span`.
- Checkpoint order is not checked, only names.
- App-defined attributes (custom events, `Scope::set_attribute`) are reported as unknown keys.
- `lk.sdk.*` is not reported when missing: resource attributes are only compared, not required.
- Units are annotations; live-check checks names, types and enum values, never magnitudes.

### In SDK integration tests

Every SDK's telemetry e2e test already writes its session to a local collector with a file
exporter. That capture is live-check input as is, so each SDK can validate what it actually
sends against this registry, in CI, with no new harness:

```
otel/live-check.sh <sdk-capture>.jsonl                 # names, types, enums, severities, checkpoints, outcomes
otel/live-check.sh <sdk-capture>.jsonl . --fail-on none  # report only
```

Rollout: run report-only until the eight renames to shared keys land in a core, then fail the
test on violations. One registry version checks all producers the same way, so a platform
cannot drift from the shared keys without its own CI noticing. It checks the vocabulary, not
transport, cadence or storage.
