#!/usr/bin/env bash
# Regenerates otel/gen from otel/model (Rust, TypeScript, markdown docs). CI fails when the
# committed outputs differ from a fresh run.
set -euo pipefail
cd "$(dirname "$0")"
W=${WEAVER:-weaver}
for target in rust:rust typescript:ts markdown:docs; do
  "$W" registry generate --v2 -r model -p policies -t templates "${target%%:*}" "gen/${target#*:}"
done
