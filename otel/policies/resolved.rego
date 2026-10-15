package after_resolution

# LiveKit rules on the resolved registry, run by otel/check.sh with otel/policies/livekit.rego.
# They see what each group ends up with after refs, `extends` and definition/2 files are
# resolved, which the per-file rules cannot.

import rego.v1

high_cardinality := {"lk.room_sid", "lk.participant_id", "lk.track_sid", "lk.job_id", "lk.client.connect_id"}

finding(id, msg) := {"id": id, "level": "violation", "message": msg}

where(group) := object.get(group, ["lineage", "provenance", "path"], group.id)

deny contains finding("livekit_metric_unit", sprintf("%s (%s): metric unit must be s, not ms", [group.metric_name, where(group)])) if {
	some group in input.groups
	group.type == "metric"
	group.unit == "ms"
}

deny contains finding("livekit_cardinality", sprintf("%s (%s): %s is high-cardinality, not a default metric dimension", [group.metric_name, where(group), attr.name])) if {
	some group in input.groups
	group.type == "metric"
	some attr in group.attributes
	high_cardinality[attr.name]
	object.get(group, ["annotations", "livekit", "cardinality_exception"], false) != true # Agents' job-scoped exception
}

# A ref inherits the referenced key's classification. Overriding it would let one signal
# re-label a key, e.g. as non-PII, while the collector keeps redacting by key name.
deny contains finding("livekit_ref_override", sprintf("%s in %s (%s): a ref may not override the annotations of a LiveKit key", [key, group.id, where(group)])) if {
	some group in input.groups
	some key, lineage in object.get(group, ["lineage", "attributes"], {})
	startswith(key, "lk.")
	"annotations" in object.get(lineage, "locally_overridden_fields", [])
}
