---
"github.com/livekit/protocol": patch
"@livekit/protocol": patch
---

`AgentThread.GetItemsRequest.sort_order` (`AgentThread.SortOrder`, as the public API's `SortOrder`): `DESC` lists a thread's items from the latest back, so a chat view can load its last page first and page up through older items. Thread listings page like the other lists: `page_token`, `limit` and `next_page_token` (strings, as `AgentDB.ListRequest`) replace `before`/`after`.
