package before_resolution

# Advice, never a failure. `weaver registry check` exits 1 on any finding whatever its level, so
# otel/check.sh runs this directory in a second pass after otel/policies and only prints it.

import rego.v1

# Attribute definitions, lk.* or not, from v1 `groups:` files and definition/2 files alike.
lk_attrs contains attr if {
	some group in input.groups
	some attr in group.attributes
	is_string(attr.id) # a definition, not a ref
	not attr.deprecated
}

lk_attrs contains object.union(attr, {"id": attr.key}) if {
	some attr in input.attributes
	not attr.deprecated
}

# The namespace convention puts client-owned attributes under lk.client.*. The registry keeps the
# client's names, so this rule is advice: it lists the client-owned keys outside lk.client.*.
deny contains {"id": "livekit_client_namespace", "level": "improvement", "message": sprintf("%s: client-owned key outside lk.client.*", [attr.id])} if {
	some attr in lk_attrs
	object.get(attr, ["annotations", "livekit", "owner"], null) == "client"
	not startswith(attr.id, "lk.client.")
}
