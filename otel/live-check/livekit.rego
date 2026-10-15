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
	"lk.participant.remote_identity": "lk.pii.remote_participant_identity",
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

# What otel/model/client/signals.yaml declares in annotations.livekit, checked on matched samples.
declared(field) := object.get(input, ["registry_group", "annotations", "livekit", field], null)

span_attr(key) := [a.value | some a in input.sample.span.attributes; a.name == key]

# A span event must be one of the span's checkpoints or match its checkpoint pattern.
deny contains lk_advice("livekit_checkpoint", {"span_name": span.name, "checkpoint": name}, sprintf("span '%s' has checkpoint '%s', which otel/model does not list for it", [span.name, name])) if {
	span := input.sample.span
	declared("checkpoints") != null
	some event in span.span_events
	name := event.name
	not name in {c.name | some c in declared("checkpoints")}
	not regex.match(object.get(input, ["registry_group", "annotations", "livekit", "checkpoint_pattern", "pattern"], "^$"), name)
}

deny contains lk_advice("livekit_severity", {"event_name": log.event_name, "severity_text": log.severity_text}, sprintf("event '%s' is sent at %v; otel/model allows %v", [log.event_name, log.severity_text, declared("severity")])) if {
	log := input.sample.log
	declared("severity") != null
	not log.severity_text in declared("severity")
}

# Status follows lk.outcome (Error only for `error`), and error.type comes only with `error`.
deny contains lk_advice("livekit_span_outcome", {"span_name": span.name, "outcome": outcome, "status": span.status.code}, sprintf("span '%s': lk.outcome=%s but status=%s", [span.name, outcome, span.status.code])) if {
	span := input.sample.span
	startswith(span.name, "lk.")
	some outcome in span_attr("lk.outcome")
	(outcome == "error") != (span.status.code == "error")
}

deny contains lk_advice("livekit_span_outcome", {"span_name": span.name, "outcome": outcome}, sprintf("span '%s': error.type without lk.outcome=error (lk.outcome=%s)", [span.name, outcome])) if {
	span := input.sample.span
	startswith(span.name, "lk.")
	count(span_attr("error.type")) > 0
	some outcome in span_attr("lk.outcome")
	outcome != "error"
}
