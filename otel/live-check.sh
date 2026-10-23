#!/usr/bin/env bash
# Checks captured client telemetry against otel/model.
# usage: otel/live-check.sh <capture> [out-dir] [extra `weaver registry live-check` args]
#   capture: OTLP/JSON lines from an otelcol file exporter (*.jsonl), or weaver samples (*.json)
#   out-dir: where live_check.json lands (default: a new temporary directory)
# Exits like weaver: 1 when there is a violation (pass `--fail-on none` to only report). When weaver
# writes no report, prints the weaver log and exits with weaver's status, or 1 if that was 0.
set -euo pipefail
here=$(cd "$(dirname "$0")" && pwd)
W=${WEAVER:-weaver}
capture=$1 out=${2:-$(mktemp -d)}
mkdir -p "$out"
# Only the previous report goes: out-dir can be the caller's and hold other files.
rm -f "$out/live_check.json"

samples=$capture
if [[ $capture == *.jsonl ]]; then
  samples=$out/samples.json
  python3 -I "$here/live-check/otlp2samples.py" "$capture" > "$samples"
fi

status=0
"$W" registry live-check --config "$here/live-check/weaver.toml" -r "$here/model" --v2 \
  --advice-policies "$here/live-check" --advice-preprocessor "$here/live-check/renames.jq" \
  --input-source "$samples" --input-format json \
  --format json -o "$out" "${@:3}" > "$out/weaver.log" 2>&1 || status=$?
if [[ ! -f $out/live_check.json ]]; then
  cat "$out/weaver.log" >&2
  echo "live-check: weaver wrote no report (exit $status)" >&2
  exit "$(( status ? status : 1 ))"
fi
python3 -I "$here/live-check/summarize.py" "$out/live_check.json"
echo "report: $out/live_check.json"
exit "$status"
