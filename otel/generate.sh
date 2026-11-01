#!/usr/bin/env bash
# Regenerates otel/gen from otel/model (Rust, TypeScript, Python, markdown docs). Generation runs
# the LiveKit policies (--v2 mode). gen/ is replaced whole, and only when every target passes, so
# an output that no target writes any more disappears. CI fails when the committed outputs differ
# from a fresh run. otel/tests/generate.sh tests the replacement.
set -euo pipefail
cd "$(dirname "$0")"
W=${WEAVER:-weaver}
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT
for target in rust:rust typescript:ts python:python markdown:docs; do
  "$W" registry generate --v2 -r model -p policies -t templates "${target%%:*}" "$tmp/gen/${target#*:}"
done
rm -rf gen && mv "$tmp/gen" gen
