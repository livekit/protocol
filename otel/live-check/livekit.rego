package live_check_advice

# Same contract as otel/policies/livekit.rego, applied to emitted telemetry. Each finding names
# the registered key a producer should send instead.

import rego.v1

pii_segment := `(^|\.)pii(\.|$)`

pii_looking := `(identity|email|phone|ip_address|user_agent|(room|participant)[._]name)`

canonical_identity := {
	"lk.room.sid": "lk.room_sid",
	"lk.participant.sid": "lk.participant_id",
	"lk.participant_sid": "lk.participant_id",
	"lk.track.sid": "lk.track_sid",
	"lk.job.id": "lk.job_id",
}

canonical_pii := {
	"lk.room.name": "lk.pii.room_name",
	"lk.participant.identity": "lk.pii.participant_identity",
	"lk.participant.remote_identity": "lk.pii.remote_participant_identity", # name TBD
}

lk_advice(type, ctx, msg) := {
	"type": "advice",
	"advice_type": type,
	"advice_level": "violation",
	"advice_context": ctx,
	"message": msg,
}

deny contains lk_advice("livekit_pii_unmarked", {"attribute_key": key, "expected": object.get(canonical_pii, key, null)}, sprintf("'%s' looks like PII but has no `pii` key segment, so the collector will not redact it%s", [key, hint])) if {
	key := input.sample.attribute.name
	startswith(key, "lk.")
	regex.match(pii_looking, key)
	not regex.match(pii_segment, key)
	hint := object.get({k: sprintf("; send '%s'", [v]) | some k, v in canonical_pii}, key, "")
}

deny contains lk_advice("livekit_flat_identity", {"attribute_key": key, "expected": want}, sprintf("'%s' must be emitted as '%s'", [key, want])) if {
	key := input.sample.attribute.name
	want := canonical_identity[key]
}

# Spans have no name-based match of their own (otel/live-check/weaver.toml matches the registered
# ones), so weaver says nothing about an lk.* span it cannot place. Like an unknown event, flag it.
deny contains lk_advice("livekit_unknown_span", {"span_name": name}, sprintf("span '%s' is in the LiveKit namespace but not registered in otel/model", [name])) if {
	name := input.sample.span.name
	startswith(name, "lk.")
	object.get(input, "registry_group", null) == null # always present, null when unmatched
}
