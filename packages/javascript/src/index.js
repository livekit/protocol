// @ts-check
const { protoInt64 } = require("@bufbuild/protobuf");
const agentDispatch = require("./gen/esm/livekit_agent_dispatch_pb.js");
const agent = require("./gen/esm/livekit_agent_pb.js");
const egress = require("./gen/esm/livekit_egress_pb.js");
const ingress = require("./gen/esm/livekit_ingress_pb.js");
const metrics = require("./gen/esm/livekit_metrics_pb.js");
const models = require("./gen/esm/livekit_models_pb.js");
const room = require("./gen/esm/livekit_room_pb.js");
const rtc = require("./gen/esm/livekit_rtc_pb.js");
const sip = require("./gen/esm/livekit_sip_pb.js");
const webhook = require("./gen/esm/livekit_webhook_pb.js");
const version = require("./gen/esm/version.js");

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
  ...version,
}
