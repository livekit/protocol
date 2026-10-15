#!/usr/bin/env bash
# Exercises otel/tests/schema-bump.sh in a throwaway repo: a changed model passes only with a
# strictly higher schema_url version.
set -uo pipefail
bump=$(cd "$(dirname "$0")" && pwd)/schema-bump.sh
tmp=$(mktemp -d) && trap 'rm -rf "$tmp"' EXIT
cd "$tmp" && git init -q && mkdir -p otel/model
url() { printf 'file_format: definition_manifest/2.0\nschema_url: %s\n' "$1" > otel/model/manifest.yaml; }
url https://livekit.io/schemas/telemetry/0.1.0 && echo 'groups: []' > otel/model/a.yaml
git add -A && git -c user.name=t -c user.email=t@t commit -qm base

failed=0
case_() { # case_ <name> <want exit> <manifest schema_url line value> <model change? 0|1>
  git checkout -q -- . && url "$3"
  [[ $4 == 1 ]] && echo '# changed' >> otel/model/a.yaml
  "$bump" HEAD > out.txt 2>&1
  local got=$?
  if [[ $got == "$2" ]]; then echo "PASS  $1 (exit $got)"; else echo "FAIL  $1 (exit $got, want $2)"; cat out.txt; failed=1; fi
}
case_ "unchanged" 0 https://livekit.io/schemas/telemetry/0.1.0 0
case_ "changed, not bumped" 1 https://livekit.io/schemas/telemetry/0.1.0 1
case_ "changed, only a comment on the URL" 1 "https://livekit.io/schemas/telemetry/0.1.0 # bumped?" 1
case_ "changed, downgraded" 1 https://livekit.io/schemas/telemetry/0.0.1 1
case_ "changed, same version quoted" 1 '"https://livekit.io/schemas/telemetry/0.1.0"' 1
case_ "changed, version not semver" 1 https://livekit.io/schemas/telemetry/next 1
case_ "changed, patch bump" 0 https://livekit.io/schemas/telemetry/0.1.1 1
case_ "changed, minor bump with a comment" 0 "https://livekit.io/schemas/telemetry/0.2.0 # renames" 1
case_ "changed, major bump" 0 https://livekit.io/schemas/telemetry/1.0.0 1
exit "$failed"
