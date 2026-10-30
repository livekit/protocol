# Live-check advice preprocessor, for live-check/livekit.rego: the renames recorded in otel/model as
# `renamed_from`, keyed by the old spelling (`data.renames`), and the span keys Agents sends for a
# shared key under another name (`bindings.agents_otlp`, its first entry), keyed by that name
# (`data.bindings`).
[.registry.registry.attributes[]] as $attrs
| {renames: ([$attrs[]
    | .annotations.livekit.renamed_from as $from
    | select($from)
    | {($from.key): {key: .key, values: ($from.values // {})}}] | add // {}),
  bindings: ([$attrs[]
    | .key as $key
    | .annotations.livekit.bindings.agents_otlp
    | select(.)
    | (if type == "array" then .[0] else . end) as $name
    | select($name != $key)
    | {($name): $key}] | add // {})}
