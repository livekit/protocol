#!/usr/bin/env bash
# The smoke sample (a real client capture, anonymized, canonical names) must pass with no
# findings, and so must a sparse RTC window, a JS network event without its flags and a record
# without severity text; the shared keys as Agents sends them report only its two span bindings,
# as improvements. The same smoke sample with the client's current identity/PII keys and
# spellings of shared keys, an unregistered lk.* span, unlisted checkpoints (one empty), outcomes
# that disagree with the span status (Error for ok, Ok for ok) and events at the wrong severity
# (by text, and by number) must be flagged. Last, a stub weaver checks the wrapper: it keeps
# weaver's status, shows weaver's log when there is no report, and never summarizes a report left
# over from an earlier run.
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

# What Agents sends for the shared keys, as bare attributes (Agents spans have no matchers). Its
# span names for the room and track keys are bindings: improvements that name the key, nothing else.
agents=$(./live-check.sh tests/live-agents.json "$tmp/agents" --search-all-attributes)
echo "$agents"
[[ $agents == *'"advice_level_counts": {"improvement": 2}'* ]] || { echo "FAIL  Agents shared-key sample has findings beyond its two bindings" >&2; exit 1; }
for want in "'lk.room_sid' is how Agents spans send 'lk.room_id'" "'lk.track_sid' is how Agents spans send 'lk.track_id'"; do
  grep -q -F -- "$want" "$tmp/agents/live_check.json" || { echo "FAIL  no finding says: $want" >&2; exit 1; }
done

# severity_text is optional in OTLP: a record with only a valid severity number passes, whether
# it comes as weaver samples or as OTLP/JSON lines.
for sample in tests/live-no-severity-text.json tests/live-no-severity-text.jsonl; do
  out=$(./live-check.sh "$sample" "$tmp/$(basename "$sample")")
  echo "$out"
  [[ $out == *'"advice_level_counts": {}'* ]] || { echo "FAIL  $sample has findings" >&2; exit 1; }
done

sed -e 's/"lk\.room_id"/"lk.room.sid"/g' -e 's/"lk\.pii\.room_name"/"lk.room.name"/g' \
  -e 's/"name": "lk\.publish"/"name": "lk.bogus"/' \
  -e '/"name": "lk\.connect"/s/"name": "ws_open"/"name": "ws_opened"/' \
  -e '/"name": "lk\.connect"/s/"name": "lk\.outcome", "value": "ok"/"name": "lk.outcome", "value": "error"/' \
  -e '/"lk\.device\.network\.changed"/s/"severity_text": "INFO"/"severity_text": "WARN"/' \
  -e '/"name": "lk\.subscribe"/s/"span_events": \[/"span_events": [{"name": "", "attributes": []}, /' \
  -e '/"name": "lk\.subscribe"/s/"status": {"code": "unset"/"status": {"code": "ok"/' \
  -e '/"event_name": "lk\.rtc\.stats\.sample"/s/"severity_number": 9/"severity_number": 17/' \
  -e 's/"name": "lk\.track_source", "value": "SOURCE_CAMERA"/"name": "lk.track.source", "value": "camera"/g' \
  -e 's/"name": "lk\.disconnect_reason", "value": "CLIENT_INITIATED"/"name": "lk.disconnect.reason", "value": "client_initiated"/' \
  -e '/"name": "lk\.connect"/s/"name": "lk\.pii\.participant_identity"/"name": "lk.caller_identity"/' \
  tests/live-sample.json > "$tmp/legacy.json"
legacy=$(./live-check.sh "$tmp/legacy.json" "$tmp/legacy" --fail-on none)
echo "$legacy"
for want in "livekit_shared_rename lk.room.sid" "livekit_shared_rename lk.room.name" "livekit_unknown_span lk.bogus" \
  "livekit_checkpoint lk.connect" "livekit_span_outcome lk.connect" "livekit_severity lk.device.network.changed" \
  "livekit_checkpoint lk.subscribe" "livekit_span_outcome lk.subscribe" "livekit_severity lk.rtc.stats.sample" \
  "livekit_shared_rename lk.track.source" "livekit_shared_rename lk.disconnect.reason" \
  "livekit_pii_unmarked lk.caller_identity"; do
  id=${want% *} key=${want#* }
  grep -q -E -- "${id}: .*${key//./\\.}" <<<"$legacy" || { echo "FAIL  not flagged: $id on $key" >&2; exit 1; }
done
# The key and value to send come from renamed_from in the model.
for want in "send 'lk.room_id'" "send 'lk.track_source' = 'SOURCE_CAMERA'" "send 'lk.disconnect_reason' = 'CLIENT_INITIATED'"; do
  grep -q -F -- "$want" "$tmp/legacy/live_check.json" || { echo "FAIL  no finding says: $want" >&2; exit 1; }
done

cat > "$tmp/weaver" <<'STUB'
#!/usr/bin/env bash
# weaver registry live-check ... -o <out-dir>: copies $STUB_REPORT there when set, exits $STUB_STATUS.
while [[ $# -gt 0 ]]; do [[ $1 == -o ]] && dir=$2; shift; done
echo "stub weaver: diagnostic"
[[ -n ${STUB_REPORT:-} ]] && cp "$STUB_REPORT" "$dir/live_check.json"
exit "${STUB_STATUS:-0}"
STUB
chmod +x "$tmp/weaver"
# stub <weaver status> <report to write, or ""> <out-dir>: sets $res (output) and $st (exit).
stub() { st=0; res=$(STUB_STATUS=$1 STUB_REPORT=$2 WEAVER=$tmp/weaver ./live-check.sh tests/live-sample.json "$3" 2>&1) || st=$?; }
fail() { echo "FAIL  $1 (exit $st): $res" >&2; exit 1; }
report=$tmp/legacy/live_check.json # a valid report with violations, from the run above

stub 7 "" "$tmp/stub7"
[[ $st == 7 && $res == *"stub weaver: diagnostic"* && $res != *Traceback* && $res != *"report:"* ]] ||
  fail "no report, weaver exit 7: must show the weaver log and exit 7"
stub 0 "" "$tmp/stub0"
[[ $st == 1 && $res == *"stub weaver: diagnostic"* && $res != *"report:"* ]] ||
  fail "no report, weaver exit 0: must show the weaver log and exit 1"
stub 1 "$report" "$tmp/stub1"
[[ $st == 1 && $res == *"totals:"* && $res == *"report: $tmp/stub1/live_check.json"* ]] ||
  fail "valid report, weaver exit 1: must summarize the report and exit 1"
mkdir -p "$tmp/reused" && cp "$report" "$tmp/reused/live_check.json" && echo keep > "$tmp/reused/unrelated.txt"
stub 7 "" "$tmp/reused"
[[ $st == 7 && $res != *"totals:"* && $res != *"report:"* && -f $tmp/reused/unrelated.txt && ! -e $tmp/reused/live_check.json ]] ||
  fail "stale report in a reused directory: must not be summarized, and an unrelated file must stay"
echo "PASS  live-check: smoke, sparse, JS and no-severity-text samples clean, Agents sample only its two bindings, every planted violation flagged, wrapper status and stale-report cases"
