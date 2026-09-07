---
"github.com/livekit/protocol": patch
---

hwstats: report the cgroup's own memory usage when the cgroup has no memory limit, instead of node-wide usage. For containers without a memory limit, memory used over memory total now describes the container rather than the node.
