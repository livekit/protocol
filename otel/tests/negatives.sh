#!/usr/bin/env bash
# Each otel/tests/negatives/<case>.yaml is added to a copy of otel/model; `weaver registry
# check` must then fail and name the finding on the fixture's `# expect:` line.
set -uo pipefail
cd "$(dirname "$0")/.." || exit
W=${WEAVER:-weaver}
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT
failed=0
for fixture in tests/negatives/*.yaml; do
  expect=$(sed -n '1s/^# expect: //p' "$fixture")
  model=$tmp/$(basename "$fixture" .yaml)
  cp -R model "$model" && cp "$fixture" "$model/negative.yaml"
  out=$("$W" registry check -r "$model" -p policies 2>&1 | sed $'s/\x1b\\[[0-9;]*m//g')
  status=${PIPESTATUS[0]}
  if [[ $status != 0 && -n $expect && $out == *"$expect"* ]]; then
    echo "PASS  $fixture fails with: $expect"
  else
    echo "FAIL  $fixture (exit $status), expected: $expect"; echo "$out" | tail -20; failed=1
  fi
done
exit "$failed"
