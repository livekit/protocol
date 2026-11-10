#!/usr/bin/env bash
# Generates Rust, TypeScript and Python from otel/model plus tests/codegen/escape.yaml, compiles
# them, and checks the tricky enum value comes back byte for byte from rustc, tsc + node and python3.
set -euo pipefail
cd "$(dirname "$0")/.."
W=${WEAVER:-weaver}
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT
want=$'a "quoted" back\\slash, \'single\', </tag> & */ tab\tnew\nline'

cp -R model "$tmp/model" && cp tests/codegen/escape.yaml "$tmp/model/"
for target in rust typescript python; do
  "$W" registry generate --v2 -r "$tmp/model" -t templates "$target" "$tmp/$target" > "$tmp/gen.log" 2>&1 ||
    { cat "$tmp/gen.log"; echo "FAIL  weaver could not generate $target" >&2; exit 1; }
done

rustc --edition 2021 --crate-type lib -D warnings -D missing_docs "$tmp/rust/semconv.rs" --out-dir "$tmp"
printf '#[allow(dead_code)]\n#[path = "%s"]\nmod semconv;\nfn main() { print!("{}", semconv::lk_test_escape::TRICKY); }\n' \
  "$tmp/rust/semconv.rs" > "$tmp/main.rs"
rustc --edition 2021 -D warnings "$tmp/main.rs" -o "$tmp/main"
[[ $("$tmp/main"; printf .) == "$want." ]] || { echo "FAIL  Rust literal does not round-trip" >&2; exit 1; }

npx --yes -p typescript@5.9.3 tsc --strict --target es2020 --module commonjs --outDir "$tmp/js" "$tmp/typescript/semconv.ts"
got=$(node -e 'process.stdout.write(require(process.argv[1]).LkTestEscape.TRICKY)' "$tmp/js/semconv.js"; printf .)
[[ $got == "$want." ]] || { echo "FAIL  TypeScript literal does not round-trip" >&2; exit 1; }
got=$(python3 -I -c 'import sys, importlib.util; s = importlib.util.spec_from_file_location("semconv", sys.argv[1]); m = importlib.util.module_from_spec(s); s.loader.exec_module(m); sys.stdout.write(m.LkTestEscape.TRICKY)' "$tmp/python/semconv.py"; printf .)
[[ $got == "$want." ]] || { echo "FAIL  Python literal does not round-trip" >&2; exit 1; }
echo "PASS  codegen: escaped literals round-trip through rustc, tsc and python3"
