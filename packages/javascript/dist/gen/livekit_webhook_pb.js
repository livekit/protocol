import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { ParticipantInfo, Room, TrackInfo } from "./livekit_models_pb.js";
import { EgressInfo } from "./livekit_egress_pb.js";
import { IngressInfo } from "./livekit_ingress_pb.js";
const _WebhookEvent = class _WebhookEvent extends Message {
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
    this.createdAt = protoInt64.zero;
    /**
     * @generated from field: int32 num_dropped = 11;
     */
    this.numDropped = 0;
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_WebhookEvent, a, b);
  }
};
_WebhookEvent.runtime = proto3;
_WebhookEvent.typeName = "livekit.WebhookEvent";
_WebhookEvent.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "event",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "room", kind: "message", T: Room },
  { no: 3, name: "participant", kind: "message", T: ParticipantInfo },
  { no: 9, name: "egress_info", kind: "message", T: EgressInfo },
  { no: 10, name: "ingress_info", kind: "message", T: IngressInfo },
  { no: 8, name: "track", kind: "message", T: TrackInfo },
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
export {
  WebhookEvent
};
//# sourceMappingURL=livekit_webhook_pb.js.map