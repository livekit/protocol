var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);
var src_exports = {};
__export(src_exports, {
  protoInt64: () => import_protobuf.protoInt64
});
module.exports = __toCommonJS(src_exports);
var import_protobuf = require("@bufbuild/protobuf");
__reExport(src_exports, require("./gen/livekit_agent_dispatch_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_agent_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_egress_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_ingress_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_metrics_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_models_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_room_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_rtc_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_sip_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/livekit_webhook_pb.cjs"), module.exports);
__reExport(src_exports, require("./gen/version.cjs"), module.exports);
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  protoInt64,
  ...require("./gen/livekit_agent_dispatch_pb.cjs"),
  ...require("./gen/livekit_agent_pb.cjs"),
  ...require("./gen/livekit_egress_pb.cjs"),
  ...require("./gen/livekit_ingress_pb.cjs"),
  ...require("./gen/livekit_metrics_pb.cjs"),
  ...require("./gen/livekit_models_pb.cjs"),
  ...require("./gen/livekit_room_pb.cjs"),
  ...require("./gen/livekit_rtc_pb.cjs"),
  ...require("./gen/livekit_sip_pb.cjs"),
  ...require("./gen/livekit_webhook_pb.cjs"),
  ...require("./gen/version.cjs")
});
//# sourceMappingURL=index.cjs.map