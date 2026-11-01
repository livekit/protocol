#!/usr/bin/env python3
"""OTLP/JSON lines (otelcol file exporter) -> weaver live-check samples JSON array.

usage: otel/live-check/otlp2samples.py session.jsonl > samples.json
Lets `weaver registry live-check --input-source <file> --input-format json` check a
capture without a gRPC replay. Stdlib only.
"""
import json, sys

KIND = {0: "internal", 1: "internal", 2: "server", 3: "client", 4: "producer", 5: "consumer"}
STATUS = {0: "unset", 1: "ok", 2: "error"}


def value(v):
    if "stringValue" in v: return v["stringValue"]
    if "intValue" in v: return int(v["intValue"])  # OTLP/JSON encodes int64 as a string
    if "doubleValue" in v: return float(v["doubleValue"])
    if "boolValue" in v: return v["boolValue"]
    if "arrayValue" in v: return [value(x) for x in v["arrayValue"].get("values", [])]
    raise ValueError(f"unsupported AnyValue {v}")  # ponytail: kvlist/bytes unused by the core today


def attrs(kvs):
    return [{"name": kv["key"], "value": value(kv["value"])} for kv in kvs or []]


def samples(row):
    for rs in row.get("resourceSpans", []) + row.get("resourceLogs", []):
        yield {"resource": {"attributes": attrs(rs.get("resource", {}).get("attributes"))}}
        for ss in rs.get("scopeSpans", []):
            for s in ss.get("spans", []):
                yield {"span": {
                    "name": s["name"],
                    "kind": KIND[s.get("kind", 0)],
                    "status": {"code": STATUS[s.get("status", {}).get("code", 0)], "message": s.get("status", {}).get("message", "")},
                    "attributes": attrs(s.get("attributes")),
                    "span_events": [{"name": e["name"], "attributes": attrs(e.get("attributes"))} for e in s.get("events", [])],
                    "span_links": [],
                }}
        for sl in rs.get("scopeLogs", []):
            for r in sl.get("logRecords", []):
                yield {"log": {
                    "event_name": r.get("eventName", ""),  # plain logs have none; weaver wants a string
                    "severity_number": r.get("severityNumber"),
                    "severity_text": r.get("severityText"),
                    "body": r.get("body", {}).get("stringValue"),
                    "attributes": attrs(r.get("attributes")),
                }}


if __name__ == "__main__":
    out = [s for line in open(sys.argv[1]) if line.strip() for s in samples(json.loads(line))]
    assert out, "no OTLP records found"
    json.dump(out, sys.stdout, indent=1)
