#!/usr/bin/env python3
"""Summarize a weaver live-check JSON report: findings by level and id, with the keys involved.

usage: otel/live-check/summarize.py live_check.json
"""
import collections, json, sys

report = json.load(open(sys.argv[1]))
by_type, keys = collections.Counter(), collections.defaultdict(collections.Counter)


def walk(x):
    if isinstance(x, dict):
        for adv in x.get("live_check_result", {}).get("all_advice", []) if isinstance(x.get("live_check_result"), dict) else []:
            t = f'{adv["level"]:<11} {adv["id"]}'
            by_type[t] += 1
            ctx = adv.get("context") or {}
            keys[t][ctx.get("attribute_key") or ctx.get("event_name") or ctx.get("span_name") or adv.get("signal_name") or "-"] += 1
        for v in x.values(): walk(v)
    elif isinstance(x, list):
        for v in x: walk(v)


walk(report["samples"])
for t, n in by_type.most_common():
    top = ", ".join(f"{k}×{c}" for k, c in keys[t].most_common(8))
    print(f"{n:4} {t}: {top}")
st = report.get("statistics", {})
print("totals:", json.dumps({k: st.get(k) for k in ("total_entities_by_type", "advice_level_counts")}))
