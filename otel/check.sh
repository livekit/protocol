#!/usr/bin/env bash
# Checks otel/model against OTel semconv v1.44.0 and the LiveKit policies (fails on any
# violation), then prints the non-failing advice from otel/advice.
# usage: otel/check.sh [extra `weaver registry check` args, e.g. --diagnostic-format gh_workflow_command]
set -euo pipefail
cd "$(dirname "$0")"
W=${WEAVER:-weaver}

"$W" registry check -r model -p policies "$@"

# weaver exits 1 whenever a policy has findings, whatever their level, so the advice pass reads
# its JSON diagnostics instead and fails only on something that is not advice.
{ "$W" registry check -r model -p advice --diagnostic-format json --diagnostic-stdout 2>/dev/null || true; } |
  python3 -I -c '
import json, sys
advice, other = [], []
for d in json.load(sys.stdin):
    e = d.get("error", {})
    if e.get("type") == "policy_violation" and e["violation"]["level"] != "violation":
        advice.append(e["violation"])
    elif d.get("diagnostic", {}).get("severity") != "Warning":  # the strict pass printed warnings
        other.append(d.get("diagnostic", {}).get("message", d))
print("Advice (non-failing): %d finding(s)" % len(advice))
for a in sorted(advice, key=lambda a: a["message"]):
    print("  %s: %s" % (a["id"], a["message"]))
sys.exit("unexpected findings in the advice pass: %s" % other if other else 0)
'
