// @ts-check
const { protoInt64 } = require("@bufbuild/protobuf");
const agentDispatch = require("./gen/livekit_agent_dispatch_pb.cjs");
const agent = require("./gen/livekit_agent_pb.cjs");
const egress = require("./gen/livekit_egress_pb.cjs");
const ingress = require("./gen/livekit_ingress_pb.cjs");
const metrics = require("./gen/livekit_metrics_pb.cjs");
const models = require("./gen/livekit_models_pb.cjs");
const room = require("./gen/livekit_room_pb.cjs");
const rtc = require("./gen/livekit_rtc_pb.cjs");
const sip = require("./gen/livekit_sip_pb.cjs");
const webhook = require("./gen/livekit_webhook_pb.cjs");
const version = require("./gen/version.cjs");

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
