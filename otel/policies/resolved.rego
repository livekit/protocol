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

# Renames recorded in the model: each shared key's `renamed_from` names the old spelling and, for
# an enum, maps the old values to its members. Read from every attribute, either representation.
attributes contains {"key": a.name, "attr": a} if {
	some g in object.get(input, "groups", [])
	some a in g.attributes
}

attributes contains {"key": a.key, "attr": a} if {
	some a in object.get(input, ["registry", "attributes"], [])
}

renamed_to[old] := a.key if {
	some a in attributes
	old := data.livekit.lib.lk(a.attr, "renamed_from").key
}

# An old spelling stays unregistered, deprecated or not: the rename lives on the shared key.
deny contains data.livekit.lib.finding("livekit_shared_rename", msg) if {
	some a in attributes
	new := renamed_to[a.key]
	msg := sprintf("%s: renamed to %s (renamed_from on %s); the old spelling may not be defined, deprecated or not", [a.key, new, new])
}

# A mapped value must be a member of the shared key's enum.
deny contains data.livekit.lib.finding("livekit_rename_value", msg) if {
	some a in attributes
	some old_value, new_value in object.get(a.attr, ["annotations", "livekit", "renamed_from", "values"], {})
	members := {m.value | some m in object.get(a.attr, ["type", "members"], [])}
	not new_value in members
	msg := sprintf("%s: renamed_from maps %s to %s, which is not a member", [a.key, old_value, new_value])
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
