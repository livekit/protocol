#!/usr/bin/env bash
# Regenerates otel/gen from otel/model (Rust, TypeScript, markdown docs). Generation runs the
# LiveKit policies (--v2 mode); gen/ is only replaced when every target passes. CI fails when the
# committed outputs differ from a fresh run.
set -euo pipefail
cd "$(dirname "$0")"
W=${WEAVER:-weaver}
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT
for target in rust:rust typescript:ts python:python markdown:docs; do
  "$W" registry generate --v2 -r model -p policies -t templates "${target%%:*}" "$tmp/${target#*:}"
done
cp -R "$tmp"/. gen/
