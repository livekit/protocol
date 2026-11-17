---
"github.com/livekit/protocol": patch
---

Tag `whatsapp_biz_opaque_callback_data` on `DialWhatsAppCallRequest` and `AcceptWhatsAppCallRequest` as PII, so `logger.Proto` redacts it
