const { protoInt64 } = require("@bufbuild/protobuf");
const dispatch = require("./gen/livekit_agent_dispatch_pb");
const agent = require("./gen/livekit_agent_pb");
const egress = require("./gen/livekit_egress_pb");
const ingress = require("./gen/livekit_ingress_pb");
const metrics = require("./gen/livekit_metrics_pb");
const models = require("./gen/livekit_models_pb");
const room = require("./gen/livekit_room_pb");
const rtc = require("./gen/livekit_rtc_pb");
const sip = require("./gen/livekit_sip_pb");
const webhook = require("./gen/livekit_webhook_pb");
const { version } = require("./gen/version");

module.exports = {
  protoInt64,
  ...dispatch,
  ...agent,
  ...egress,
  ...ingress,
  ...metrics,
  ...models,
  ...room,
  ...rtc,
  ...sip,
  ...webhook,
  version,
};
