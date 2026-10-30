package before_resolution

# The LiveKit rules on each definition file, in weaver's default mode (with --v2 weaver skips
# before_resolution policies; resolved.rego covers that mode). The rules are in lib.rego.

import rego.v1

# Attribute definitions in this file, lk.* or not. Refs and deprecated keys are exempt, but
# resolved.rego still applies the flat-identity rule to deprecated keys.
lk_attrs contains attr if {
	some group in input.groups
	some attr in group.attributes
	is_string(attr.id) # a definition, not a ref
	not attr.deprecated
}

# definition/2 files list attributes at the top level, keyed by `key`.
lk_attrs contains object.union(attr, {"id": attr.key}) if {
	some attr in input.attributes
	not attr.deprecated
}

deny contains f if {
	some attr in lk_attrs
	some f in data.livekit.lib.attr_findings(attr)
}

# A definition/2 metric refinement is absent from the default mode's resolved registry, so its
# dimensions are checked here (resolved.rego sees them with --v2).
deny contains f if {
	some r in object.get(input, "metric_refinements", [])
	keys := [a.ref | some a in object.get(r, "attributes", []); a.ref]
	exception := data.livekit.lib.lk(r, "cardinality_exception")
	some f in data.livekit.lib.cardinality(r.id, "metric refinement", keys, exception)
}
