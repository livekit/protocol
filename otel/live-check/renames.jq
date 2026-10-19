# Live-check advice preprocessor: the renames recorded in otel/model as `renamed_from`, keyed by
# the old spelling, for live-check/livekit.rego (`data.renames`).
{renames: ([.registry.registry.attributes[]
  | .annotations.livekit.renamed_from as $from
  | select($from)
  | {($from.key): {key: .key, values: ($from.values // {})}}] | add // {})}
