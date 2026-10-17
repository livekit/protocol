package after_resolution

# The LiveKit rules on the resolved registry, in both of weaver's representations: default mode
# hands this package `input.groups`, --v2 hands it `input.registry` and `input.refinements`. It
# sees what refs, `extends`, refinements and definition/2 files resolve to. The rules are in
# lib.rego.

import rego.v1

# --v2 skips the per-file pass, so the attribute rules run here on the resolved definitions.
deny contains f if {
	some attr in object.get(input, ["registry", "attributes"], [])
	startswith(attr.key, "lk.")
	not attr.deprecated
	some f in data.livekit.lib.attr_findings(object.union(attr, {"id": attr.key}))
}

# Every resolved metric, either representation: name, unit, dimension keys, exception, origin.
metrics contains {"name": g.metric_name, "unit": unit, "keys": keys, "exception": exception, "where": where} if {
	some g in object.get(input, "groups", [])
	g.type == "metric"
	unit := object.get(g, "unit", "")
	keys := [a.name | some a in g.attributes]
	exception := data.livekit.lib.lk(g, "cardinality_exception")
	where := object.get(g, ["lineage", "provenance", "path"], g.id)
}

metrics contains {"name": m.name, "unit": unit, "keys": keys, "exception": exception, "where": where} if {
	some section in ["registry", "refinements"]
	some m in object.get(input, [section, "metrics"], [])
	unit := object.get(m, "unit", "")
	keys := [a.key | some a in object.get(m, "attributes", [])]
	exception := data.livekit.lib.lk(m, "cardinality_exception")
	where := object.get(m, ["provenance", "path"], m.name)
}

deny contains data.livekit.lib.finding("livekit_metric_unit", msg) if {
	some m in metrics
	m.unit == "ms"
	msg := sprintf("%s (%s): metric unit must be s, not ms", [m.name, m.where])
}

deny contains f if {
	some m in metrics
	some f in data.livekit.lib.cardinality(m.name, m.where, m.keys, m.exception)
}

# A ref inherits the referenced key's classification. Overriding it would let one signal re-label
# a key, e.g. as non-PII, while the collector keeps redacting by key name.
override_msg(key, signal, where) := sprintf("%s in %s (%s): a ref may not override the annotations of a LiveKit key", [key, signal, where])

# Default mode: the lineage records which fields a ref overrode.
deny contains data.livekit.lib.finding("livekit_ref_override", msg) if {
	some group in object.get(input, "groups", [])
	some key, lineage in object.get(group, ["lineage", "attributes"], {})
	startswith(key, "lk.")
	"annotations" in object.get(lineage, "locally_overridden_fields", [])
	msg := override_msg(key, group.id, object.get(group, ["lineage", "provenance", "path"], group.id))
}

# --v2: no lineage, so compare each signal's copy of a key with the key's definition.
definitions := {a.key: data.livekit.lib.lk_all(a) | some a in object.get(input, ["registry", "attributes"], [])}

v2_signals contains s if {
	some kind in ["attribute_groups", "metrics", "spans", "events", "entities"]
	some s in object.get(input, ["registry", kind], [])
}

v2_signals contains s if {
	some kind in ["metrics", "spans", "events", "entities"]
	some s in object.get(input, ["refinements", kind], [])
}

deny contains data.livekit.lib.finding("livekit_ref_override", msg) if {
	some s in v2_signals
	some a in object.get(s, "attributes", [])
	startswith(a.key, "lk.")
	data.livekit.lib.lk_all(a) != object.get(definitions, a.key, null)
	signal := object.get(s, "id", object.get(s, "name", object.get(s, "type", "?")))
	msg := override_msg(a.key, signal, object.get(s, ["provenance", "path"], signal))
}
