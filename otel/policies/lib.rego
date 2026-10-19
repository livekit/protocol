package livekit.lib

# The LiveKit rules, shared by livekit.rego (each definition file, default mode) and resolved.rego
# (the resolved registry, the only pass weaver runs with --v2). Call them by full path:
# data.livekit.lib.<rule>. Each finding id is documented in otel/README.md, and
# otel/tests/negatives has a fixture for each, run in both modes.

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

high_cardinality := {"lk.room_sid", "lk.participant_id", "lk.track_sid", "lk.job_id", "lk.client.connect_id"}

owners := {"shared", "client", "agents", "cloud"}

# Mirrors protocol protobufs/logger/options.proto Sensitivity (UNSPECIFIED/PII/SECRET).
sensitivities := {"none", "pii", "secret"}

# A duration word ending the key, optionally followed by `_ms` and one aggregation segment:
# lk.client.rtt, lk.rtc.freezes_duration_ms, lk.client.rtt.avg, lk.rtc.rtt_ms.max.
duration_like := `(^|[._])(duration|latency|delay|rtt|jitter)(_ms)?([._](min|max|avg|mean|median|sum|total|last|p[0-9]+))?$`

finding(id, msg) := {"id": id, "level": "violation", "message": msg}

lk(obj, field) := object.get(obj, ["annotations", "livekit", field], null)

# A documented reason to keep milliseconds (the RTC stats keys).
ms_exception(attr) if {
	lk(attr, "unit") == "ms"
	is_string(lk(attr, "unit_exception"))
	lk(attr, "unit_exception") != ""
}

# Every violation of one LiveKit attribute definition, its key in `id`.
attr_findings(attr) := ((((((owner(attr) | sensitivity(attr)) | secret(attr)) | pii_marker(attr)) | pii_heuristic(attr)) | flat_identity(attr)) | seconds(attr)) | unit(attr)

owner(attr) := {finding("livekit_owner", sprintf("%s: annotations.livekit.owner must be one of %v", [attr.id, owners])) |
	not owners[lk(attr, "owner")]
}

sensitivity(attr) := {finding("livekit_sensitivity", sprintf("%s: annotations.livekit.sensitivity must be one of %v (protocol logger.Sensitivity)", [attr.id, sensitivities])) |
	not sensitivities[lk(attr, "sensitivity")]
}

secret(attr) := {finding("livekit_secret", sprintf("%s: SECRET values are never emitted as telemetry", [attr.id])) |
	lk(attr, "sensitivity") == "secret"
}

pii_marker(attr) := {finding("livekit_pii_marker", sprintf("%s: sensitivity=%v, but the collector redacts by key, so a `pii` key segment must be present iff sensitivity=pii", [attr.id, lk(attr, "sensitivity")])) |
	sensitivities[lk(attr, "sensitivity")]
	(lk(attr, "sensitivity") == "pii") != regex.match(pii_segment, attr.id)
}

pii_heuristic(attr) := {finding("livekit_pii_heuristic", sprintf("%s looks like PII; name it lk.pii.* so the collector redacts it", [attr.id])) |
	regex.match(pii_looking, attr.id)
	not regex.match(pii_segment, attr.id)
	not heuristic_exception(attr)
}

# The shipped keys the heuristic would flag, registered as emitted. Only a key listed HERE may
# carry `annotations.livekit.pii_heuristic_exception` (the annotation documents it at the key;
# this map is the gate), so a new identity/email/phone key cannot opt out of the rule by
# annotating itself. The collector still does not redact these: this documents, it does not
# excuse. Adding a key means editing this file, in review.
pii_heuristic_allowlist := {
	"lk.rpc.caller_identity": "Agents 2.0 ships the RPC caller identity unmarked; renaming it to lk.pii.rpc.caller_identity is a producer change",
	"lk.rpc.destination_identity": "Agents 2.0 ships the RPC destination identity unmarked; same as lk.rpc.caller_identity",
}

heuristic_exception(attr) if {
	pii_heuristic_allowlist[attr.id]
	is_string(lk(attr, "pii_heuristic_exception"))
	lk(attr, "pii_heuristic_exception") != ""
}

flat_identity(attr) := {finding("livekit_flat_identity", sprintf("%s: shared identity keys are flat, use %s", [attr.id, canonical_identity[attr.id]])) |
	canonical_identity[attr.id]
}

seconds(attr) := {finding("livekit_seconds", sprintf("%s: LiveKit durations are seconds; a *_ms key needs annotations.livekit.unit: ms plus a unit_exception reason", [attr.id])) |
	regex.match(`_ms($|\.)`, attr.id)
	not ms_exception(attr)
}

unit(attr) := {finding("livekit_unit", sprintf("%s: duration-like attribute must declare annotations.livekit.unit: s (or unit: ms with a unit_exception)", [attr.id])) |
	regex.match(duration_like, attr.id)
	lk(attr, "unit") != "s"
	not ms_exception(attr)
}

# A metric dimension list against the cardinality rule; `exception` is the metric's own flag.
cardinality(name, where, keys, exception) := {finding("livekit_cardinality", sprintf("%s (%s): %s is high-cardinality, not a default metric dimension", [name, where, key])) |
	exception != true # Agents' job-scoped exception
	some key in keys
	high_cardinality[key]
}

lk_all(obj) := object.get(obj, ["annotations", "livekit"], null)
