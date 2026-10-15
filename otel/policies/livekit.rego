package before_resolution

# LiveKit rules per definition file (violations), run by otel/check.sh. Each finding id is
# documented in otel/README.md; otel/tests/negatives has a failing fixture for each.

import rego.v1

# Mirrors the collector's redaction matcher, agents-private
# pii-redaction/redact/oteltraces.go:16. Keep the two in sync (and otel/live-check/livekit.rego).
pii_segment := `(^|\.)pii(\.|$)`

# ponytail: keyword heuristic, misses PII with innocent names; reviewers still own it.
pii_looking := `(identity|email|phone|ip_address|user_agent|(room|participant)[._]name)`

canonical_identity := {
	"lk.room.sid": "lk.room_sid",
	"lk.participant.sid": "lk.participant_id",
	"lk.participant_sid": "lk.participant_id",
	"lk.track.sid": "lk.track_sid",
	"lk.job.id": "lk.job_id",
}

owners := {"shared", "client", "agents", "cloud"}

# Mirrors protocol protobufs/logger/options.proto Sensitivity (UNSPECIFIED/PII/SECRET).
sensitivities := {"none", "pii", "secret"}

# LiveKit attribute definitions in this file (refs and deprecated keys are exempt).
lk_attrs contains attr if {
	some group in input.groups
	some attr in group.attributes
	startswith(attr.id, "lk.")
	not attr.deprecated
}

# definition/2 files list attributes at the top level, keyed by `key`.
lk_attrs contains object.union(attr, {"id": attr.key}) if {
	some attr in input.attributes
	startswith(attr.key, "lk.")
	not attr.deprecated
}

lk(obj, field) := object.get(obj, ["annotations", "livekit", field], null)

# A documented reason to keep milliseconds (e.g. RTC stats in webrtc-stats units).
ms_exception(attr) if {
	lk(attr, "unit") == "ms"
	is_string(lk(attr, "unit_exception"))
	lk(attr, "unit_exception") != ""
}

finding(id, msg) := {"id": id, "level": "violation", "message": msg}

deny contains finding("livekit_owner", sprintf("%s: annotations.livekit.owner must be one of %v", [attr.id, owners])) if {
	some attr in lk_attrs
	not owners[lk(attr, "owner")]
}

deny contains finding("livekit_sensitivity", sprintf("%s: annotations.livekit.sensitivity must be one of %v (protocol logger.Sensitivity)", [attr.id, sensitivities])) if {
	some attr in lk_attrs
	not sensitivities[lk(attr, "sensitivity")]
}

deny contains finding("livekit_secret", sprintf("%s: SECRET values are never emitted as telemetry", [attr.id])) if {
	some attr in lk_attrs
	lk(attr, "sensitivity") == "secret"
}

deny contains finding("livekit_pii_marker", sprintf("%s: sensitivity=%v, but the collector redacts by key, so a `pii` key segment must be present iff sensitivity=pii", [attr.id, lk(attr, "sensitivity")])) if {
	some attr in lk_attrs
	sensitivities[lk(attr, "sensitivity")]
	(lk(attr, "sensitivity") == "pii") != regex.match(pii_segment, attr.id)
}

deny contains finding("livekit_pii_heuristic", sprintf("%s looks like PII; name it lk.pii.* so the collector redacts it", [attr.id])) if {
	some attr in lk_attrs
	regex.match(pii_looking, attr.id)
	not regex.match(pii_segment, attr.id)
}

deny contains finding("livekit_flat_identity", sprintf("%s: shared identity keys are flat, use %s", [attr.id, canonical_identity[attr.id]])) if {
	some attr in lk_attrs
	canonical_identity[attr.id]
}

deny contains finding("livekit_seconds", sprintf("%s: LiveKit durations are seconds; a *_ms key needs annotations.livekit.unit: ms plus a unit_exception reason", [attr.id])) if {
	some attr in lk_attrs
	regex.match(`_ms($|\.)`, attr.id)
	not ms_exception(attr)
}

# A duration word ending the key, optionally followed by `_ms` and one aggregation segment:
# lk.client.rtt, lk.rtc.freezes_duration_ms, lk.client.rtt.avg, lk.rtc.rtt_ms.max.
duration_like := `(^|[._])(duration|latency|delay|rtt|jitter)(_ms)?([._](min|max|avg|mean|median|sum|total|last|p[0-9]+))?$`

deny contains finding("livekit_unit", sprintf("%s: duration-like attribute must declare annotations.livekit.unit: s (or unit: ms with a unit_exception)", [attr.id])) if {
	some attr in lk_attrs
	regex.match(duration_like, attr.id)
	lk(attr, "unit") != "s"
	not ms_exception(attr)
}

# Metric rules (units, cardinality) and ref overrides live in otel/policies/resolved.rego.
