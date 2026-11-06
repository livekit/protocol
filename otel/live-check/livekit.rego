package live_check_advice

# Same contract as otel/policies/livekit.rego, applied to emitted telemetry. A finding on an old
# spelling names the registered key, and value, a producer must send instead.

import rego.v1

pii_segment := `(^|\.)pii(\.|$)`

pii_looking := `(identity|email|phone|ip_address|user_agent|(room|participant)[._]name)`

# The renames recorded in otel/model (`renamed_from` on each shared key), keyed by the old
# spelling: `data.renames[old] = {key, values}`. Built by live-check/renames.jq.
renames := object.get(data, "renames", {})

# The span keys Agents sends for a shared key under another name (`bindings.agents_otlp` in
# otel/model), keyed by that name: `data.bindings[name] = key`. Also built by renames.jq.
bindings := object.get(data, "bindings", {})

lk_advice(type, ctx, msg) := lk_finding("violation", type, ctx, msg)

lk_finding(level, type, ctx, msg) := {
	"type": "advice",
	"advice_type": type,
	"advice_level": level,
	"advice_context": ctx,
	"message": msg,
}

# A key the registry check exempts from its PII heuristic: the same rule, allowlist and annotation,
# from policies/lib.rego (loaded here through the live-check/lib.rego symlink), applied to the
# key's registered definition. An unregistered key has none and is never exempt.
pii_exempt(key) if {
	attr := object.get(input, "registry_attribute", null)
	is_object(attr)
	data.livekit.lib.heuristic_exception(object.union(attr, {"id": key}))
}

# A PII-looking key that is not an old spelling (those get livekit_shared_rename) or exempt.
deny contains lk_advice("livekit_pii_unmarked", {"attribute_key": key}, sprintf("'%s' looks like PII but has no `pii` key segment, so the collector will not redact it", [key])) if {
	key := input.sample.attribute.name
	startswith(key, "lk.")
	regex.match(pii_looking, key)
	not regex.match(pii_segment, key)
	not renames[key]
	not pii_exempt(key)
}

# An old spelling: the model names the key to send and, for an enum, the value.
deny contains lk_advice("livekit_shared_rename", {"attribute_key": key, "expected": to.key, "expected_value": value}, sprintf("'%s' is renamed: send '%s'%s", [key, to.key, hint])) if {
	key := input.sample.attribute.name
	to := renames[key]
	value := object.get(to.values, input.sample.attribute.value, null)
	hint := rename_value_hint(to, value)
}

# An active Agents span binding names the registered key as an improvement. Historical bindings
# are not in data.bindings, so they still fail as unregistered attributes.
deny contains lk_finding("improvement", "livekit_agents_binding", {"attribute_key": key, "expected": to}, sprintf("'%s' is how Agents spans send '%s' (bindings.agents_otlp)", [key, to])) if {
	key := input.sample.attribute.name
	to := bindings[key]
}

rename_value_hint(to, value) := "" if count(to.values) == 0

rename_value_hint(to, value) := sprintf(" = '%s'", [value]) if {
	count(to.values) > 0
	value != null
}

rename_value_hint(to, value) := sprintf(" (no value mapped for '%v')", [input.sample.attribute.value]) if {
	count(to.values) > 0
	value == null
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

# A span event must be one of the span's checkpoints, or match a pattern the span declares.
matches_declared_pattern(name) if {
	pattern := object.get(input, ["registry_group", "annotations", "livekit", "checkpoint_pattern", "pattern"], null)
	is_string(pattern)
	regex.match(pattern, name)
}

deny contains lk_advice("livekit_checkpoint", {"span_name": span.name, "checkpoint": name}, sprintf("span '%s' has checkpoint '%s', which otel/model does not list for it", [span.name, name])) if {
	span := input.sample.span
	declared("checkpoints") != null
	some event in span.span_events
	name := event.name
	not name in {c.name | some c in declared("checkpoints")}
	not matches_declared_pattern(name)
}

# OTLP severity numbers per severity text, all six ranges of the logs data model.
severity_range := {"TRACE": [1, 4], "DEBUG": [5, 8], "INFO": [9, 12], "WARN": [13, 16], "ERROR": [17, 20], "FATAL": [21, 24]}

# The matched event's annotations.livekit, {} when the sample matched no event.
event_livekit := object.get(input, ["registry_group", "annotations", "livekit"], {})

# `severity` is declared, also as an explicit null: absent means no constraint.
severity_declared if {
	is_object(event_livekit)
	some field, _ in event_livekit
	field == "severity"
}

# The same declaration rule as livekit_event_severity in policies/lib.rego, so live-check does not
# depend on otel/check.sh having run on the model.
severity_valid if {
	severity_declared
	data.livekit.lib.severity_texts_only(event_livekit.severity)
}

# A declaration no record can be compared against: reported once per record, and no record is
# compared against it.
deny contains lk_advice("livekit_event_severity", {"event_name": log.event_name}, sprintf("event '%s': otel/model declares severity %v, which is not a list of %v; records are not compared against it", [log.event_name, event_livekit.severity, data.livekit.lib.severity_texts])) if {
	log := input.sample.log
	severity_declared
	not severity_valid
}

# `severity: []` declares an event that is never a log record (the GenAI span events): a log record
# under its name is one finding that says so, not a severity mismatch.
deny contains lk_advice("livekit_severity", {"event_name": log.event_name}, sprintf("event '%s' is declared as not a log record (severity: []), but was sent as one", [log.event_name])) if {
	log := input.sample.log
	severity_valid
	event_livekit.severity == []
}

# A valid declaration with one or more severities.
declared_levels if {
	severity_valid
	count(event_livekit.severity) > 0
}

number_declared(n) if {
	some text in declared("severity")
	severity_range[text][0] <= n
	n <= severity_range[text][1]
}

# The number decides; a text, when present, must agree with the declaration too.
deny contains lk_advice("livekit_severity", {"event_name": log.event_name, "severity_number": log.severity_number}, sprintf("event '%s' is sent at severity number %v; otel/model allows %v", [log.event_name, log.severity_number, declared("severity")])) if {
	log := input.sample.log
	declared_levels
	not number_declared(object.get(log, "severity_number", 0))
}

deny contains lk_advice("livekit_severity", {"event_name": log.event_name, "severity_text": log.severity_text}, sprintf("event '%s' is sent at %v; otel/model allows %v", [log.event_name, log.severity_text, declared("severity")])) if {
	log := input.sample.log
	declared_levels
	text := object.get(log, "severity_text", null) # optional in OTLP; omitted arrives as null
	is_string(text)
	text != ""
	not text in declared("severity")
}

# The span status lk.outcome implies: Error for `error`, Unset for `ok` and `cancelled`.
status_for := {"ok": "unset", "cancelled": "unset", "error": "error"}

deny contains lk_advice("livekit_span_outcome", {"span_name": span.name, "outcome": outcome, "status": span.status.code}, sprintf("span '%s': lk.outcome=%s requires status %s, got %s", [span.name, outcome, status_for[outcome], span.status.code])) if {
	span := input.sample.span
	startswith(span.name, "lk.")
	some outcome in span_attr("lk.outcome")
	span.status.code != status_for[outcome]
}

deny contains lk_advice("livekit_span_outcome", {"span_name": span.name, "outcome": outcome}, sprintf("span '%s': error.type without lk.outcome=error (lk.outcome=%s)", [span.name, outcome])) if {
	span := input.sample.span
	startswith(span.name, "lk.")
	count(span_attr("error.type")) > 0
	some outcome in span_attr("lk.outcome")
	outcome != "error"
}
