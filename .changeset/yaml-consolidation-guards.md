---
"@livekit/protocol": patch
---

Guard the yaml package used by configutil.Observer. Observer decodes caller-owned
config trees, so custom unmarshallers in those trees must not be bound to one
yaml package. Adds tests asserting Observer honors func-based UnmarshalYAML at the
top level, nested, and behind a pointer.
