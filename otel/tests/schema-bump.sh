#!/usr/bin/env bash
# Fails when otel/model differs from <base> but the schema_url version did not strictly increase.
# usage: otel/tests/schema-bump.sh <base-ref>   (CI passes origin/<base branch>)
# otel/tests/schema-bump-test.sh exercises it.
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"
base=$(git merge-base "$1" HEAD)

if git diff --quiet "$base" -- otel/model; then
  echo "otel/model unchanged since $base"; exit 0
fi
python3 -I - "$base" otel/model/manifest.yaml <<'PY'
import re, subprocess, sys

base, path = sys.argv[1:]

def version(text, where):
    """schema_url as YAML reads it (quotes, trailing comment dropped) and its MAJOR.MINOR.PATCH."""
    m = re.search(r"""^schema_url:[ \t]*(["']?)([^"'#\s]+)\1[ \t]*(#.*)?$""", text, re.M)
    v = m and re.fullmatch(r".*/(\d+)\.(\d+)\.(\d+)", m.group(2))
    if not v:
        sys.exit(f"{where}: schema_url must end in a MAJOR.MINOR.PATCH version")
    return m.group(2), tuple(int(n) for n in v.groups())

new_url, new = version(open(path).read(), path)
old = subprocess.run(["git", "show", f"{base}:{path}"], capture_output=True, text=True)
if old.returncode:
    print(f"new registry: {new_url}")
    sys.exit()
old_url, old_v = version(old.stdout, f"{base}:{path}")
if new <= old_v:
    sys.exit(f"otel/model changed since {base[:12]}, but {new_url} is not a higher version than {old_url}: bump it")
print(f"schema version bumped: {old_url} -> {new_url}")
PY
