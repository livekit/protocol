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
# Use the go.mod toolchain's gofmt (`$(go env GOROOT)/bin/gofmt`), not whatever bare gofmt is on
# PATH, so formatting matches the Go version CI builds and tests the generated package with.
"$(go env GOROOT)/bin/gofmt" -w "$tmp/gen/go"
# Go has no namespaces: every attribute key, enum value, event, span, checkpoint and metric becomes
# a top-level constant, where Rust and Python keep them apart in modules and classes. Two that
# generate the same identifier are a redeclaration; `go build` rejects it, but catch it here with
# the offending name, so a model addition cannot slip a silent flat-namespace collision past review.
dupes=$(grep -hoE '^const [A-Za-z0-9_]+' "$tmp/gen/go/semconv.go" | awk '{print $2}' | sort | uniq -d)
if [[ -n $dupes ]]; then
  echo "generate.sh: generated Go has duplicate top-level constants (flat-namespace collision):" >&2
  echo "$dupes" >&2
  exit 1
fi
rm -rf gen && mv "$tmp/gen" gen
