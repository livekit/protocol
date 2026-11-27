#!/usr/bin/env bash
# Regenerates otel/gen from otel/model (Rust, TypeScript, Python, markdown docs). Generation runs
# the LiveKit policies (--v2 mode). gen/ is replaced whole, and only when every target passes, so
# an output that no target writes any more disappears. CI fails when the committed outputs differ
# from a fresh run. otel/tests/generate.sh tests the replacement.
set -euo pipefail
cd "$(dirname "$0")"
W=${WEAVER:-weaver}
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT
for target in go:go rust:rust typescript:ts python:python markdown:docs; do
  "$W" registry generate --v2 -r model -p policies -t templates "${target%%:*}" "$tmp/gen/${target#*:}"
done
# The Go output is gofmt-canonical, so the committed file matches `gofmt`/`go fmt` and stays
# diff-stable. weaver emits unaligned const and map literals that gofmt would otherwise rewrite;
# the other targets need no such pass. gofmt fails loudly on a template that emits invalid Go.
gofmt -w "$tmp/gen/go"
rm -rf gen && mv "$tmp/gen" gen
