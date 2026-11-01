#!/usr/bin/env bash
# otel/generate.sh against a stub weaver: a failing target leaves gen/ byte-identical, and a
# successful run replaces gen/ whole, so outputs that no target writes any more disappear.
set -euo pipefail
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT
cp "$(dirname "$0")/../generate.sh" "$tmp/generate.sh"
cat > "$tmp/weaver" <<'STUB'
#!/usr/bin/env bash
# weaver registry generate ... <target> <out-dir>
target=${@: -2:1} out=${@: -1}
[[ $target == "${FAIL_TARGET:-}" ]] && exit 3
mkdir -p "$out" && echo "new $target" > "$out/semconv.$target"
STUB
chmod +x "$tmp/weaver"
mkdir -p "$tmp/gen/rust" "$tmp/gen/obsolete-dir"
echo old > "$tmp/gen/rust/semconv.rust"
echo old > "$tmp/gen/obsolete.txt"
echo old > "$tmp/gen/obsolete-dir/file"
cp -Rp "$tmp/gen" "$tmp/before"

if WEAVER=$tmp/weaver FAIL_TARGET=python "$tmp/generate.sh" 2>/dev/null; then
  echo "FAIL  generate.sh succeeded although the python target failed" >&2; exit 1
fi
diff -r "$tmp/before" "$tmp/gen" || { echo "FAIL  a failing target changed gen/" >&2; exit 1; }

WEAVER=$tmp/weaver "$tmp/generate.sh"
got=$(cd "$tmp/gen" && find . -mindepth 1 | sort | tr '\n' ' ')
want="./docs ./docs/semconv.markdown ./python ./python/semconv.python ./rust ./rust/semconv.rust ./ts ./ts/semconv.typescript "
[[ $got == "$want" ]] || { echo "FAIL  gen/ after a successful run: $got" >&2; exit 1; }
[[ $(cat "$tmp/gen/rust/semconv.rust") == "new rust" ]] || { echo "FAIL  gen/ kept an old output" >&2; exit 1; }
echo "PASS  generate: a failing target leaves gen/ unchanged, a successful run removes obsolete outputs"
