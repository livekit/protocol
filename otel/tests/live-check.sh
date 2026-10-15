#!/usr/bin/env bash
# The smoke sample (a real client capture, anonymized, canonical names) must pass with no
# findings, and so must a sparse RTC window. The same smoke sample with the client's current
# identity/PII keys, an unregistered lk.* span, an unlisted checkpoint, an outcome that disagrees
# with the span status and an event at the wrong severity must be flagged.
set -euo pipefail
cd "$(dirname "$0")/.."
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT

good=$(./live-check.sh tests/live-sample.json "$tmp/good")
echo "$good"
[[ $good == *'"advice_level_counts": {}'* ]] || { echo "FAIL  smoke sample has findings" >&2; exit 1; }

# The core omits every counter and gauge a reading lacks (RtcStatsSample's Option fields).
sparse=$(./live-check.sh tests/live-sparse.json "$tmp/sparse")
echo "$sparse"
[[ $sparse == *'"advice_level_counts": {}'* ]] || { echo "FAIL  sparse RTC sample has findings" >&2; exit 1; }

sed -e 's/"lk\.room_sid"/"lk.room.sid"/g' -e 's/"lk\.pii\.room_name"/"lk.room.name"/g' \
  -e 's/"name": "lk\.publish"/"name": "lk.bogus"/' \
  -e '/"name": "lk\.connect"/s/"name": "ws_open"/"name": "ws_opened"/' \
  -e '/"name": "lk\.connect"/s/"name": "lk\.outcome", "value": "ok"/"name": "lk.outcome", "value": "error"/' \
  -e '/"lk\.device\.network\.changed"/s/"severity_text": "INFO"/"severity_text": "WARN"/' \
  tests/live-sample.json > "$tmp/legacy.json"
legacy=$(./live-check.sh "$tmp/legacy.json" "$tmp/legacy" --fail-on none)
echo "$legacy"
for want in "livekit_flat_identity: lk.room.sid" "livekit_pii_unmarked: lk.room.name" "livekit_unknown_span: lk.bogus" \
  "livekit_checkpoint: lk.connect" "livekit_span_outcome: lk.connect" "livekit_severity: lk.device.network.changed"; do
  [[ $legacy == *"$want"* ]] || { echo "FAIL  legacy keys not flagged: $want" >&2; exit 1; }
done
echo "PASS  live-check: smoke and sparse samples clean, every planted violation flagged"
