---
"github.com/livekit/protocol": patch
---

`utils/list`: mutations involving a node the list does not own — removing or moving an unlinked node, inserting an already-linked node, or inserting relative to an unlinked mark — are now no-ops, as in `container/list`. Previously an unlinked node's zeroed hook made `Remove`/`MoveToFront`/`MoveToBack` clear the list's head and tail, silently stranding every remaining element, and re-inserting a linked node forked the list.
