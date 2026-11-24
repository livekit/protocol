#!/usr/bin/env bash
# Each otel/tests/negatives/<case>.yaml is added to a copy of otel/model; `weaver registry
# check` must then fail and name the finding on the fixture's `# expect:` line, in the default
# mode and with --v2 (the mode generate.sh runs the policies in).
set -uo pipefail
cd "$(dirname "$0")/.." || exit
W=${WEAVER:-weaver}
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT
failed=0
for fixture in tests/negatives/*.yaml; do
  expect=$(sed -n '1s/^# expect: //p' "$fixture")
  model=$tmp/$(basename "$fixture" .yaml)
  cp -R model "$model" && cp "$fixture" "$model/negative.yaml"
  for mode in default --v2; do
    flags=() && [[ $mode == --v2 ]] && flags=(--v2)
    out=$("$W" registry check ${flags[@]+"${flags[@]}"} -r "$model" -p policies 2>&1 | sed $'s/\x1b\\[[0-9;]*m//g')
    status=${PIPESTATUS[0]}
    if [[ $status != 0 && -n $expect && $out == *"$expect"* ]]; then
      echo "PASS  $fixture ($mode) fails with: $expect"
    else
      echo "FAIL  $fixture ($mode, exit $status), expected: $expect"; echo "$out" | tail -20; failed=1
    fi
  done
done
exit "$failed"
