// @ts-check
const { protoInt64 } = require("@bufbuild/protobuf");
const agentDispatch = require("./gen/cjs/livekit_agent_dispatch_pb.cjs");
const agent = require("./gen/cjs/livekit_agent_pb.cjs");
const egress = require("./gen/cjs/livekit_egress_pb.cjs");
const ingress = require("./gen/cjs/livekit_ingress_pb.cjs");
const metrics = require("./gen/cjs/livekit_metrics_pb.cjs");
const models = require("./gen/cjs/livekit_models_pb.cjs");
const room = require("./gen/cjs/livekit_room_pb.cjs");
const rtc = require("./gen/cjs/livekit_rtc_pb.cjs");
const sip = require("./gen/cjs/livekit_sip_pb.cjs");
const webhook = require("./gen/cjs/livekit_webhook_pb.cjs");
const version = require("./gen/cjs/version.cjs");

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
