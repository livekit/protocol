#!/usr/bin/env bash
# The smoke sample (a real client capture, anonymized, canonical names) must pass with no
# findings, and so must a sparse RTC window and a JS network event without its flags. The same
# smoke sample with the client's current identity/PII keys, an unregistered lk.* span, unlisted
# checkpoints (one empty), outcomes that disagree with the span status (Error for ok, Ok for ok)
# and events at the wrong severity (by text, and by number) must be flagged.
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

# A browser without the Network Information API: the network event without its two flags.
js=$(./live-check.sh tests/live-js-network.json "$tmp/js")
echo "$js"
[[ $js == *'"advice_level_counts": {}'* ]] || { echo "FAIL  JS network sample has findings" >&2; exit 1; }

sed -e 's/"lk\.room_sid"/"lk.room.sid"/g' -e 's/"lk\.pii\.room_name"/"lk.room.name"/g' \
  -e 's/"name": "lk\.publish"/"name": "lk.bogus"/' \
  -e '/"name": "lk\.connect"/s/"name": "ws_open"/"name": "ws_opened"/' \
  -e '/"name": "lk\.connect"/s/"name": "lk\.outcome", "value": "ok"/"name": "lk.outcome", "value": "error"/' \
  -e '/"lk\.device\.network\.changed"/s/"severity_text": "INFO"/"severity_text": "WARN"/' \
  -e '/"name": "lk\.subscribe"/s/"span_events": \[/"span_events": [{"name": "", "attributes": []}, /' \
  -e '/"name": "lk\.subscribe"/s/"status": {"code": "unset"/"status": {"code": "ok"/' \
  -e '/"event_name": "lk\.rtc\.stats\.sample"/s/"severity_number": 9/"severity_number": 17/' \
  tests/live-sample.json > "$tmp/legacy.json"
legacy=$(./live-check.sh "$tmp/legacy.json" "$tmp/legacy" --fail-on none)
echo "$legacy"
for want in "livekit_flat_identity lk.room.sid" "livekit_pii_unmarked lk.room.name" "livekit_unknown_span lk.bogus" \
  "livekit_checkpoint lk.connect" "livekit_span_outcome lk.connect" "livekit_severity lk.device.network.changed" \
  "livekit_checkpoint lk.subscribe" "livekit_span_outcome lk.subscribe" "livekit_severity lk.rtc.stats.sample"; do
  id=${want% *} key=${want#* }
  grep -q -E -- "${id}: .*${key//./\\.}" <<<"$legacy" || { echo "FAIL  not flagged: $id on $key" >&2; exit 1; }
done
echo "PASS  live-check: smoke, sparse and JS samples clean, every planted violation flagged"
