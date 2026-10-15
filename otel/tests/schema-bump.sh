#!/usr/bin/env bash
# Fails when otel/model differs from <base> but its schema_url was not bumped.
# usage: otel/tests/schema-bump.sh <base-ref>   (CI passes origin/<base branch>)
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"
base=$(git merge-base "$1" HEAD)
manifest=otel/model/manifest.yaml

if git diff --quiet "$base" -- otel/model; then
  echo "otel/model unchanged since $base"; exit 0
fi
old=$(git show "$base:$manifest" 2>/dev/null | grep '^schema_url:' || true)
new=$(grep '^schema_url:' "$manifest")
if [[ -z $old ]]; then
  echo "new registry: $new"
elif [[ $old == "$new" ]]; then
  echo "otel/model changed since $base but $manifest still has $new: bump the version" >&2; exit 1
else
  echo "schema version bumped: ${old#schema_url: } -> ${new#schema_url: }"
fi
