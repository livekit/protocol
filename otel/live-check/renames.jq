# Live-check advice preprocessor, for live-check/livekit.rego: the renames recorded in otel/model as
# `renamed_from`, keyed by the old spelling (`data.renames`), and the span keys Agents sends for a
# shared key under another name (`bindings.agents_otlp`, its first entry), keyed by that name
# (`data.bindings`), and the keys with a `pii_heuristic_exception` (`data.pii_exceptions`).
[.registry.registry.attributes[]] as $attrs
| {renames: ([$attrs[]
    | .annotations.livekit.renamed_from as $from
    | select($from)
    | {($from.key): {key: .key, values: ($from.values // {})}}] | add // {}),
  # policies/lib.rego honors the annotation only on its pii_heuristic_allowlist, so once
  # otel/check.sh passes, this exempts the same keys the registry check does.
  pii_exceptions: ([$attrs[]
    | select(.annotations.livekit.pii_heuristic_exception | type == "string" and length > 0)
    | {(.key): true}] | add // {}),
  bindings: ([$attrs[]
    | .key as $key
    | .annotations.livekit.bindings.agents_otlp
    | select(.)
    | (if type == "array" then .[0] else . end) as $name
    | select($name != $key)
    | {($name): $key}] | add // {})}
