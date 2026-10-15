package before_resolution

# Advice, never a failure. `weaver registry check` exits 1 on any finding whatever its level, so
# otel/check.sh runs this directory in a second pass after otel/policies and only prints it.

import rego.v1

# LiveKit attribute definitions, from v1 `groups:` files and definition/2 files alike.
lk_attrs contains attr if {
	some group in input.groups
	some attr in group.attributes
	startswith(attr.id, "lk.")
	not attr.deprecated
}

lk_attrs contains object.union(attr, {"id": attr.key}) if {
	some attr in input.attributes
	startswith(attr.key, "lk.")
	not attr.deprecated
}

# The Notion conventions put client-owned attributes under lk.client.*. The client core ships
# its current names until that decision is made; this lists what the rule would move.
deny contains {"id": "livekit_client_namespace", "level": "improvement", "message": sprintf("%s: client-owned; the Notion convention would move it under lk.client.*", [attr.id])} if {
	some attr in lk_attrs
	object.get(attr, ["annotations", "livekit", "owner"], null) == "client"
	not startswith(attr.id, "lk.client.")
}
