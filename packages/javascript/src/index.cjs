// @ts-check
const { protoInt64 } = require("@bufbuild/protobuf");
const agentDispatch = require("./gen/cjs/livekit_agent_dispatch_pb.js");
const agent = require("./gen/cjs/livekit_agent_pb.js");
const egress = require("./gen/cjs/livekit_egress_pb.js");
const ingress = require("./gen/cjs/livekit_ingress_pb.js");
const metrics = require("./gen/cjs/livekit_metrics_pb.js");
const models = require("./gen/cjs/livekit_models_pb.js");
const room = require("./gen/cjs/livekit_room_pb.js");
const rtc = require("./gen/cjs/livekit_rtc_pb.js");
const sip = require("./gen/cjs/livekit_sip_pb.js");
const webhook = require("./gen/cjs/livekit_webhook_pb.js");
const version = require("./gen/cjs/version.js");

module.exports = {
  protoInt64,
  ...agentDispatch,
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
}
