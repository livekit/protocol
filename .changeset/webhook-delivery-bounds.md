---
"github.com/livekit/protocol": patch
---

webhook: bound each delivery attempt by default (10 s client timeout when none is set) and cap a `Retry-After` wait at `RetryWaitMax`.
webhook: mask userinfo and query values of the delivery URL in logs and send errors, log `http.StatusText` instead of the endpoint's status text, and cap the response head at 64 KiB.
