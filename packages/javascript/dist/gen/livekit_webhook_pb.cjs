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
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);
var livekit_webhook_pb_exports = {};
__export(livekit_webhook_pb_exports, {
  WebhookEvent: () => WebhookEvent
});
module.exports = __toCommonJS(livekit_webhook_pb_exports);
var import_protobuf = require("@bufbuild/protobuf");
var import_livekit_models_pb = require("./livekit_models_pb.cjs");
var import_livekit_egress_pb = require("./livekit_egress_pb.cjs");
var import_livekit_ingress_pb = require("./livekit_ingress_pb.cjs");
const _WebhookEvent = class _WebhookEvent extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * one of room_started, room_finished, participant_joined, participant_left,
     * track_published, track_unpublished, egress_started, egress_updated, egress_ended,
     * ingress_started, ingress_ended
     *
     * @generated from field: string event = 1;
     */
    this.event = "";
    /**
     * unique event uuid
     *
     * @generated from field: string id = 6;
     */
    this.id = "";
    /**
     * timestamp in seconds
     *
     * @generated from field: int64 created_at = 7;
     */
    this.createdAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int32 num_dropped = 11;
     */
    this.numDropped = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _WebhookEvent().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _WebhookEvent().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _WebhookEvent().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_WebhookEvent, a, b);
  }
};
_WebhookEvent.runtime = import_protobuf.proto3;
_WebhookEvent.typeName = "livekit.WebhookEvent";
_WebhookEvent.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "event",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "room", kind: "message", T: import_livekit_models_pb.Room },
  { no: 3, name: "participant", kind: "message", T: import_livekit_models_pb.ParticipantInfo },
  { no: 9, name: "egress_info", kind: "message", T: import_livekit_egress_pb.EgressInfo },
  { no: 10, name: "ingress_info", kind: "message", T: import_livekit_ingress_pb.IngressInfo },
  { no: 8, name: "track", kind: "message", T: import_livekit_models_pb.TrackInfo },
  {
    no: 6,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "created_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 11,
    name: "num_dropped",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  }
]);
let WebhookEvent = _WebhookEvent;
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  WebhookEvent
});
//# sourceMappingURL=livekit_webhook_pb.cjs.map