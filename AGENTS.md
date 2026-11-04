# AGENTS.md

Instructions for everyone who changes this repository, people and coding agents.

## Generate

```bash
go tool mage proto                                     # Go, Twirp, and psrpc code
pnpm --filter @livekit/protocol run generate:proto     # the JS package
```

On a push to a branch other than `main`, the Generate workflow regenerates the code and commits it as "generated protobuf". CI pins the versions of `protoc` and the plugins. A different local version changes the version stamps and the formatting of every generated file, so do not commit generated files that you made with other versions.

A change that consumers see, such as a change to a public proto, adds a changeset in `.changeset/` for `@livekit/protocol` and `github.com/livekit/protocol`.

## Proto comments are public documentation

The protos in `protobufs/` are public. Their comments reach several audiences:

- `protoc-gen-go` copies each comment into the Go doc comments that SDK users read.
- `protoc-gen-es` copies each comment into the JSDoc of `@livekit/protocol`.
- The LiveKit Cloud Public API imports some of these messages, and its OpenAPI spec publishes their comments as descriptions. Examples are the egress outputs, `RoomConfiguration`, the agent simulation messages, and the agent session and thread messages.

The protos in `protobufs/rpc/`, `protobufs/roomrpc/`, and `protobufs/infra/`, and `livekit_internal.proto`, are for LiveKit services only. These rules apply to every other proto. Before you add or change a comment on a public proto, read it as a customer.

### Keep a note internal

**A detached comment** has one blank line between the comment and the declaration. No generator copies it. Use it for a note that is only for the engineers who work on this repository:

```proto
// A note for the engineers who work on this repository.

// A scenario that a simulation run plays.
message Scenario {
```

Do not use an `(-- ... --)` marker for an internal note. The Public API spec removes it, but the Go and JS output keep it.

Some comments are published by accident:

- A trailing comment (`string x = 1; // note`) is published. Put a published comment on the line above the declaration.
- A comment that labels a group of fields is published as the comment of the first field. Put a blank line after it to detach it.
- A comment on the line after the last field of a message is published as the comment of that field. Put a blank line before it to detach it.
- A comment on a `reserved` line is not published.

### Write a published comment

- Write for the person who uses the API or the SDK. Address them as "you", not "we".
- Do not name the systems behind LiveKit, such as a LiveKit Cloud service, a database, a queue, or an internal RPC. Do not describe how or where LiveKit stores data, as in "stored as JSON" or "server-side". Do not describe the API by what it is built on, with words such as "mirrors" or "upstream". Do not write `TODO` or `FIXME`.
- Use the proto field names, in code font: `room_name`, not `roomName` or `RoomName`. Each language converts the name in its own way, and the proto name is the same for all of them. Keep the real spelling of OpenTelemetry attributes and agent hook names.
- Use code font for field names, message names, enum values, and literal values: `room_name`, `RoomConfiguration`, `H264_720P_30`, `true`.
- Use ASD-STE100 Simplified Technical English: short sentences, active voice, present tense, and one instruction in each sentence. Write "returns", not "will return". Write "use", not "utilize".
- Describe the API as it is now. Do not tell what changed or how to migrate, and do not write a time-bound word such as "today", "currently", "yet", or "no longer". Do not refer to an API version, as in "not filterable in v1". For a deprecated field, write "Deprecated." and name the field to use instead.
- Write to inform the reader, not to sell. Do not use marketing language.
- Use plain CommonMark. Do not use headings, tables, or HTML.
- Do not use em dashes. Use a colon, a comma, parentheses, or a new sentence.
- Write "for example" or "such as", not "e.g." or "i.e.".
- Do not use a possessive for a thing: write "the name of the room", not "the room's name". A person and an agent can take a possessive.
- When two nouns make a known term, write the term instead of an "of" phrase: "the track SID", "the egress status", "the trace ID". Keep "of" when the term is not known or is ambiguous, or when the term has more than three nouns: "the participants of the room".

### Patterns for each element

Tools show the first sentence by itself, so make it complete and specific. Do not start with the symbol name: write "A room session", not "Session is a room session".

| Element | First sentence | Example |
| --- | --- | --- |
| Service | A verb that tells what the caller can do. | "Manages the rooms of a project." |
| RPC | A third-person verb: Lists, Returns, Creates, Updates, Deletes. | "Lists the active rooms." |
| Message | A noun phrase. | "The encoding settings of an egress." |
| Field | A noun phrase that starts with "The" or "A". | "The name of the room." |
| Request filter | A noun phrase, then "Returns only ...". | "The name of a room. Returns only the egresses of this room." |
| Boolean field | What true does, or "True when ...". | "When true, the egress does not upload the manifest file." |
| Enum value | What the value means. | "The run is waiting for its source upload." |

After the first sentence, give the facts that the reader needs. Leave out a fact that does not apply.

- **Unset or empty:** what happens when the field is not set. Write "When empty, ..." or "Unset means ...".
- **Default:** "Defaults to 1920." Write a full sentence, not "(default 1920)".
- **Limits:** the range and whether it is inclusive, or the maximum length. Tell what happens to a value outside the range: the API rejects it or clamps it.
- **Units:** use the unit in the field name (`ttl_seconds`) or state it.
- **Format:** an ID prefix (`RM_...`), a timestamp, or an example value.
- **Order:** the sort order of a list.
- **Errors:** the conditions that cause an error, and the error code, in code font.
- **Side effects and idempotency:** for an RPC that changes data, tell what else changes and what a retry does.
