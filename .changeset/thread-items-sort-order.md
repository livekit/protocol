---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

`AgentThread.GetItemsRequest.sort_order` (`AgentThread.SortOrder`): `DESC` lists a thread's items from the latest back, so a chat view can load its last page first and page up through older items. Thread listings page like the other lists: `page_token`, `limit` and `next_page_token` (strings: `livekit.agent` cannot use `TokenPagination`) replace `before`/`after`, and `GetTasks` pages too. In `thread_schema_v1.sql`, `items` gains `pos INTEGER PRIMARY KEY AUTOINCREMENT` (insertion order, never reused, so a page token never skips an item) and an index on `(lane, pos)`; `id` stays the upsert key as `TEXT NOT NULL UNIQUE`.
