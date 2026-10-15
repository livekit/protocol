#!/usr/bin/env bash
# The smoke sample (a real client capture, anonymized, canonical names) must pass with no
# findings, and so must a sparse RTC window; the same smoke sample with the client's current
# identity/PII keys must be flagged.
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
  tests/live-sample.json > "$tmp/legacy.json"
legacy=$(./live-check.sh "$tmp/legacy.json" "$tmp/legacy" --fail-on none)
echo "$legacy"
for want in "livekit_flat_identity: lk.room.sid" "livekit_pii_unmarked: lk.room.name"; do
  [[ $legacy == *"$want"* ]] || { echo "FAIL  legacy keys not flagged: $want" >&2; exit 1; }
done
echo "PASS  live-check: smoke and sparse samples clean, legacy keys flagged"
