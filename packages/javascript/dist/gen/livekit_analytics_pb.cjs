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
var livekit_analytics_pb_exports = {};
__export(livekit_analytics_pb_exports, {
  AnalyticsClientMeta: () => AnalyticsClientMeta,
  AnalyticsEvent: () => AnalyticsEvent,
  AnalyticsEventType: () => AnalyticsEventType,
  AnalyticsEvents: () => AnalyticsEvents,
  AnalyticsNodeRooms: () => AnalyticsNodeRooms,
  AnalyticsRoom: () => AnalyticsRoom,
  AnalyticsRoomParticipant: () => AnalyticsRoomParticipant,
  AnalyticsStat: () => AnalyticsStat,
  AnalyticsStats: () => AnalyticsStats,
  AnalyticsStream: () => AnalyticsStream,
  AnalyticsVideoLayer: () => AnalyticsVideoLayer,
  StreamType: () => StreamType
});
module.exports = __toCommonJS(livekit_analytics_pb_exports);
var import_protobuf = require("@bufbuild/protobuf");
var import_livekit_models_pb = require("./livekit_models_pb.cjs");
var import_livekit_egress_pb = require("./livekit_egress_pb.cjs");
var import_livekit_ingress_pb = require("./livekit_ingress_pb.cjs");
var import_livekit_sip_pb = require("./livekit_sip_pb.cjs");
var StreamType = /* @__PURE__ */ ((StreamType2) => {
  StreamType2[StreamType2["UPSTREAM"] = 0] = "UPSTREAM";
  StreamType2[StreamType2["DOWNSTREAM"] = 1] = "DOWNSTREAM";
  return StreamType2;
})(StreamType || {});
import_protobuf.proto3.util.setEnumType(StreamType, "livekit.StreamType", [
  { no: 0, name: "UPSTREAM" },
  { no: 1, name: "DOWNSTREAM" }
]);
var AnalyticsEventType = /* @__PURE__ */ ((AnalyticsEventType2) => {
  AnalyticsEventType2[AnalyticsEventType2["ROOM_CREATED"] = 0] = "ROOM_CREATED";
  AnalyticsEventType2[AnalyticsEventType2["ROOM_ENDED"] = 1] = "ROOM_ENDED";
  AnalyticsEventType2[AnalyticsEventType2["PARTICIPANT_JOINED"] = 2] = "PARTICIPANT_JOINED";
  AnalyticsEventType2[AnalyticsEventType2["PARTICIPANT_LEFT"] = 3] = "PARTICIPANT_LEFT";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_PUBLISHED"] = 4] = "TRACK_PUBLISHED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_PUBLISH_REQUESTED"] = 20] = "TRACK_PUBLISH_REQUESTED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_UNPUBLISHED"] = 5] = "TRACK_UNPUBLISHED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_SUBSCRIBED"] = 6] = "TRACK_SUBSCRIBED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_SUBSCRIBE_REQUESTED"] = 21] = "TRACK_SUBSCRIBE_REQUESTED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_SUBSCRIBE_FAILED"] = 25] = "TRACK_SUBSCRIBE_FAILED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_UNSUBSCRIBED"] = 7] = "TRACK_UNSUBSCRIBED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_PUBLISHED_UPDATE"] = 10] = "TRACK_PUBLISHED_UPDATE";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_MUTED"] = 23] = "TRACK_MUTED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_UNMUTED"] = 24] = "TRACK_UNMUTED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_PUBLISH_STATS"] = 26] = "TRACK_PUBLISH_STATS";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_SUBSCRIBE_STATS"] = 27] = "TRACK_SUBSCRIBE_STATS";
  AnalyticsEventType2[AnalyticsEventType2["PARTICIPANT_ACTIVE"] = 11] = "PARTICIPANT_ACTIVE";
  AnalyticsEventType2[AnalyticsEventType2["PARTICIPANT_RESUMED"] = 22] = "PARTICIPANT_RESUMED";
  AnalyticsEventType2[AnalyticsEventType2["EGRESS_STARTED"] = 12] = "EGRESS_STARTED";
  AnalyticsEventType2[AnalyticsEventType2["EGRESS_ENDED"] = 13] = "EGRESS_ENDED";
  AnalyticsEventType2[AnalyticsEventType2["EGRESS_UPDATED"] = 28] = "EGRESS_UPDATED";
  AnalyticsEventType2[AnalyticsEventType2["TRACK_MAX_SUBSCRIBED_VIDEO_QUALITY"] = 14] = "TRACK_MAX_SUBSCRIBED_VIDEO_QUALITY";
  AnalyticsEventType2[AnalyticsEventType2["RECONNECTED"] = 15] = "RECONNECTED";
  AnalyticsEventType2[AnalyticsEventType2["INGRESS_CREATED"] = 18] = "INGRESS_CREATED";
  AnalyticsEventType2[AnalyticsEventType2["INGRESS_DELETED"] = 19] = "INGRESS_DELETED";
  AnalyticsEventType2[AnalyticsEventType2["INGRESS_STARTED"] = 16] = "INGRESS_STARTED";
  AnalyticsEventType2[AnalyticsEventType2["INGRESS_ENDED"] = 17] = "INGRESS_ENDED";
  AnalyticsEventType2[AnalyticsEventType2["INGRESS_UPDATED"] = 29] = "INGRESS_UPDATED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_INBOUND_TRUNK_CREATED"] = 30] = "SIP_INBOUND_TRUNK_CREATED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_INBOUND_TRUNK_DELETED"] = 31] = "SIP_INBOUND_TRUNK_DELETED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_OUTBOUND_TRUNK_CREATED"] = 32] = "SIP_OUTBOUND_TRUNK_CREATED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_OUTBOUND_TRUNK_DELETED"] = 33] = "SIP_OUTBOUND_TRUNK_DELETED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_DISPATCH_RULE_CREATED"] = 34] = "SIP_DISPATCH_RULE_CREATED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_DISPATCH_RULE_DELETED"] = 35] = "SIP_DISPATCH_RULE_DELETED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_PARTICIPANT_CREATED"] = 36] = "SIP_PARTICIPANT_CREATED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_CALL_INCOMING"] = 37] = "SIP_CALL_INCOMING";
  AnalyticsEventType2[AnalyticsEventType2["SIP_CALL_STARTED"] = 38] = "SIP_CALL_STARTED";
  AnalyticsEventType2[AnalyticsEventType2["SIP_CALL_ENDED"] = 39] = "SIP_CALL_ENDED";
  return AnalyticsEventType2;
})(AnalyticsEventType || {});
import_protobuf.proto3.util.setEnumType(AnalyticsEventType, "livekit.AnalyticsEventType", [
  { no: 0, name: "ROOM_CREATED" },
  { no: 1, name: "ROOM_ENDED" },
  { no: 2, name: "PARTICIPANT_JOINED" },
  { no: 3, name: "PARTICIPANT_LEFT" },
  { no: 4, name: "TRACK_PUBLISHED" },
  { no: 20, name: "TRACK_PUBLISH_REQUESTED" },
  { no: 5, name: "TRACK_UNPUBLISHED" },
  { no: 6, name: "TRACK_SUBSCRIBED" },
  { no: 21, name: "TRACK_SUBSCRIBE_REQUESTED" },
  { no: 25, name: "TRACK_SUBSCRIBE_FAILED" },
  { no: 7, name: "TRACK_UNSUBSCRIBED" },
  { no: 10, name: "TRACK_PUBLISHED_UPDATE" },
  { no: 23, name: "TRACK_MUTED" },
  { no: 24, name: "TRACK_UNMUTED" },
  { no: 26, name: "TRACK_PUBLISH_STATS" },
  { no: 27, name: "TRACK_SUBSCRIBE_STATS" },
  { no: 11, name: "PARTICIPANT_ACTIVE" },
  { no: 22, name: "PARTICIPANT_RESUMED" },
  { no: 12, name: "EGRESS_STARTED" },
  { no: 13, name: "EGRESS_ENDED" },
  { no: 28, name: "EGRESS_UPDATED" },
  { no: 14, name: "TRACK_MAX_SUBSCRIBED_VIDEO_QUALITY" },
  { no: 15, name: "RECONNECTED" },
  { no: 18, name: "INGRESS_CREATED" },
  { no: 19, name: "INGRESS_DELETED" },
  { no: 16, name: "INGRESS_STARTED" },
  { no: 17, name: "INGRESS_ENDED" },
  { no: 29, name: "INGRESS_UPDATED" },
  { no: 30, name: "SIP_INBOUND_TRUNK_CREATED" },
  { no: 31, name: "SIP_INBOUND_TRUNK_DELETED" },
  { no: 32, name: "SIP_OUTBOUND_TRUNK_CREATED" },
  { no: 33, name: "SIP_OUTBOUND_TRUNK_DELETED" },
  { no: 34, name: "SIP_DISPATCH_RULE_CREATED" },
  { no: 35, name: "SIP_DISPATCH_RULE_DELETED" },
  { no: 36, name: "SIP_PARTICIPANT_CREATED" },
  { no: 37, name: "SIP_CALL_INCOMING" },
  { no: 38, name: "SIP_CALL_STARTED" },
  { no: 39, name: "SIP_CALL_ENDED" }
]);
const _AnalyticsVideoLayer = class _AnalyticsVideoLayer extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: int32 layer = 1;
     */
    this.layer = 0;
    /**
     * @generated from field: uint32 packets = 2;
     */
    this.packets = 0;
    /**
     * @generated from field: uint64 bytes = 3;
     */
    this.bytes = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint32 frames = 4;
     */
    this.frames = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsVideoLayer().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsVideoLayer().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsVideoLayer().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsVideoLayer, a, b);
  }
};
_AnalyticsVideoLayer.runtime = import_protobuf.proto3;
_AnalyticsVideoLayer.typeName = "livekit.AnalyticsVideoLayer";
_AnalyticsVideoLayer.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "layer",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 2,
    name: "packets",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "bytes",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 4,
    name: "frames",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let AnalyticsVideoLayer = _AnalyticsVideoLayer;
const _AnalyticsStream = class _AnalyticsStream extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: uint32 ssrc = 1;
     */
    this.ssrc = 0;
    /**
     * @generated from field: uint32 primary_packets = 2;
     */
    this.primaryPackets = 0;
    /**
     * @generated from field: uint64 primary_bytes = 3;
     */
    this.primaryBytes = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint32 retransmit_packets = 4;
     */
    this.retransmitPackets = 0;
    /**
     * @generated from field: uint64 retransmit_bytes = 5;
     */
    this.retransmitBytes = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint32 padding_packets = 6;
     */
    this.paddingPackets = 0;
    /**
     * @generated from field: uint64 padding_bytes = 7;
     */
    this.paddingBytes = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint32 packets_lost = 8;
     */
    this.packetsLost = 0;
    /**
     * @generated from field: uint32 frames = 9;
     */
    this.frames = 0;
    /**
     * @generated from field: uint32 rtt = 10;
     */
    this.rtt = 0;
    /**
     * @generated from field: uint32 jitter = 11;
     */
    this.jitter = 0;
    /**
     * @generated from field: uint32 nacks = 12;
     */
    this.nacks = 0;
    /**
     * @generated from field: uint32 plis = 13;
     */
    this.plis = 0;
    /**
     * @generated from field: uint32 firs = 14;
     */
    this.firs = 0;
    /**
     * @generated from field: repeated livekit.AnalyticsVideoLayer video_layers = 15;
     */
    this.videoLayers = [];
    /**
     * @generated from field: uint32 packets_out_of_order = 19;
     */
    this.packetsOutOfOrder = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsStream().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsStream().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsStream().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsStream, a, b);
  }
};
_AnalyticsStream.runtime = import_protobuf.proto3;
_AnalyticsStream.typeName = "livekit.AnalyticsStream";
_AnalyticsStream.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "ssrc",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 2,
    name: "primary_packets",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "primary_bytes",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 4,
    name: "retransmit_packets",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 5,
    name: "retransmit_bytes",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 6,
    name: "padding_packets",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 7,
    name: "padding_bytes",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 8,
    name: "packets_lost",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 9,
    name: "frames",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 10,
    name: "rtt",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 11,
    name: "jitter",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 12,
    name: "nacks",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 13,
    name: "plis",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 14,
    name: "firs",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 15, name: "video_layers", kind: "message", T: AnalyticsVideoLayer, repeated: true },
  { no: 17, name: "start_time", kind: "message", T: import_protobuf.Timestamp },
  { no: 18, name: "end_time", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 19,
    name: "packets_out_of_order",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let AnalyticsStream = _AnalyticsStream;
const _AnalyticsStat = class _AnalyticsStat extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * unique id for this stat
     *
     * @generated from field: string id = 14;
     */
    this.id = "";
    /**
     * @generated from field: string analytics_key = 1;
     */
    this.analyticsKey = "";
    /**
     * @generated from field: livekit.StreamType kind = 2;
     */
    this.kind = 0 /* UPSTREAM */;
    /**
     * @generated from field: string node = 4;
     */
    this.node = "";
    /**
     * @generated from field: string room_id = 5;
     */
    this.roomId = "";
    /**
     * @generated from field: string room_name = 6;
     */
    this.roomName = "";
    /**
     * @generated from field: string participant_id = 7;
     */
    this.participantId = "";
    /**
     * @generated from field: string track_id = 8;
     */
    this.trackId = "";
    /**
     * average score
     *
     * @generated from field: float score = 9;
     */
    this.score = 0;
    /**
     * @generated from field: repeated livekit.AnalyticsStream streams = 10;
     */
    this.streams = [];
    /**
     * @generated from field: string mime = 11;
     */
    this.mime = "";
    /**
     * @generated from field: float min_score = 12;
     */
    this.minScore = 0;
    /**
     * @generated from field: float median_score = 13;
     */
    this.medianScore = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsStat().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsStat().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsStat().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsStat, a, b);
  }
};
_AnalyticsStat.runtime = import_protobuf.proto3;
_AnalyticsStat.typeName = "livekit.AnalyticsStat";
_AnalyticsStat.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 14,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 1,
    name: "analytics_key",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "kind", kind: "enum", T: import_protobuf.proto3.getEnumType(StreamType) },
  { no: 3, name: "time_stamp", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 4,
    name: "node",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "room_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "participant_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "track_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "score",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  { no: 10, name: "streams", kind: "message", T: AnalyticsStream, repeated: true },
  {
    no: 11,
    name: "mime",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 12,
    name: "min_score",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 13,
    name: "median_score",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  }
]);
let AnalyticsStat = _AnalyticsStat;
const _AnalyticsStats = class _AnalyticsStats extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.AnalyticsStat stats = 1;
     */
    this.stats = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsStats().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsStats().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsStats().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsStats, a, b);
  }
};
_AnalyticsStats.runtime = import_protobuf.proto3;
_AnalyticsStats.typeName = "livekit.AnalyticsStats";
_AnalyticsStats.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "stats", kind: "message", T: AnalyticsStat, repeated: true }
]);
let AnalyticsStats = _AnalyticsStats;
const _AnalyticsClientMeta = class _AnalyticsClientMeta extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string region = 1;
     */
    this.region = "";
    /**
     * @generated from field: string node = 2;
     */
    this.node = "";
    /**
     * @generated from field: string client_addr = 3;
     */
    this.clientAddr = "";
    /**
     * @generated from field: uint32 client_connect_time = 4;
     */
    this.clientConnectTime = 0;
    /**
     * udp, tcp, turn
     *
     * @generated from field: string connection_type = 5;
     */
    this.connectionType = "";
    /**
     * @generated from field: livekit.ReconnectReason reconnect_reason = 6;
     */
    this.reconnectReason = import_livekit_models_pb.ReconnectReason.RR_UNKNOWN;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsClientMeta().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsClientMeta().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsClientMeta().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsClientMeta, a, b);
  }
};
_AnalyticsClientMeta.runtime = import_protobuf.proto3;
_AnalyticsClientMeta.typeName = "livekit.AnalyticsClientMeta";
_AnalyticsClientMeta.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "region",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "node",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "client_addr",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "client_connect_time",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 5,
    name: "connection_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "reconnect_reason", kind: "enum", T: import_protobuf.proto3.getEnumType(import_livekit_models_pb.ReconnectReason) },
  { no: 7, name: "geo_hash", kind: "scalar", T: 9, opt: true },
  { no: 8, name: "country", kind: "scalar", T: 9, opt: true },
  { no: 9, name: "isp_asn", kind: "scalar", T: 13, opt: true }
]);
let AnalyticsClientMeta = _AnalyticsClientMeta;
const _AnalyticsEvent = class _AnalyticsEvent extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * unique id for this event
     *
     * @generated from field: string id = 25;
     */
    this.id = "";
    /**
     * @generated from field: livekit.AnalyticsEventType type = 1;
     */
    this.type = 0 /* ROOM_CREATED */;
    /**
     * @generated from field: string room_id = 3;
     */
    this.roomId = "";
    /**
     * @generated from field: string participant_id = 5;
     */
    this.participantId = "";
    /**
     * @generated from field: string track_id = 7;
     */
    this.trackId = "";
    /**
     * @generated from field: string analytics_key = 10;
     */
    this.analyticsKey = "";
    /**
     * @generated from field: string egress_id = 13;
     */
    this.egressId = "";
    /**
     * @generated from field: string ingress_id = 19;
     */
    this.ingressId = "";
    /**
     * @generated from field: livekit.VideoQuality max_subscribed_video_quality = 14;
     */
    this.maxSubscribedVideoQuality = import_livekit_models_pb.VideoQuality.LOW;
    /**
     * @generated from field: string mime = 16;
     */
    this.mime = "";
    /**
     * @generated from field: string error = 20;
     */
    this.error = "";
    /**
     * @generated from field: int32 video_layer = 22;
     */
    this.videoLayer = 0;
    /**
     * @generated from field: string node_id = 24;
     */
    this.nodeId = "";
    /**
     * @generated from field: string sip_call_id = 26;
     */
    this.sipCallId = "";
    /**
     * @generated from field: string sip_trunk_id = 28;
     */
    this.sipTrunkId = "";
    /**
     * @generated from field: string sip_dispatch_rule_id = 31;
     */
    this.sipDispatchRuleId = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsEvent().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsEvent().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsEvent().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsEvent, a, b);
  }
};
_AnalyticsEvent.runtime = import_protobuf.proto3;
_AnalyticsEvent.typeName = "livekit.AnalyticsEvent";
_AnalyticsEvent.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 25,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 1, name: "type", kind: "enum", T: import_protobuf.proto3.getEnumType(AnalyticsEventType) },
  { no: 2, name: "timestamp", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 3,
    name: "room_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "room", kind: "message", T: import_livekit_models_pb.Room },
  {
    no: 5,
    name: "participant_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "participant", kind: "message", T: import_livekit_models_pb.ParticipantInfo },
  {
    no: 7,
    name: "track_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 8, name: "track", kind: "message", T: import_livekit_models_pb.TrackInfo },
  {
    no: 10,
    name: "analytics_key",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 11, name: "client_info", kind: "message", T: import_livekit_models_pb.ClientInfo },
  { no: 12, name: "client_meta", kind: "message", T: AnalyticsClientMeta },
  {
    no: 13,
    name: "egress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 19,
    name: "ingress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 14, name: "max_subscribed_video_quality", kind: "enum", T: import_protobuf.proto3.getEnumType(import_livekit_models_pb.VideoQuality) },
  { no: 15, name: "publisher", kind: "message", T: import_livekit_models_pb.ParticipantInfo },
  {
    no: 16,
    name: "mime",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 17, name: "egress", kind: "message", T: import_livekit_egress_pb.EgressInfo },
  { no: 18, name: "ingress", kind: "message", T: import_livekit_ingress_pb.IngressInfo },
  {
    no: 20,
    name: "error",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 21, name: "rtp_stats", kind: "message", T: import_livekit_models_pb.RTPStats },
  {
    no: 22,
    name: "video_layer",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 24,
    name: "node_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 26,
    name: "sip_call_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 27, name: "sip_call", kind: "message", T: import_livekit_sip_pb.SIPCallInfo },
  {
    no: 28,
    name: "sip_trunk_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 29, name: "sip_inbound_trunk", kind: "message", T: import_livekit_sip_pb.SIPInboundTrunkInfo },
  { no: 30, name: "sip_outbound_trunk", kind: "message", T: import_livekit_sip_pb.SIPOutboundTrunkInfo },
  {
    no: 31,
    name: "sip_dispatch_rule_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 32, name: "sip_dispatch_rule", kind: "message", T: import_livekit_sip_pb.SIPDispatchRuleInfo }
]);
let AnalyticsEvent = _AnalyticsEvent;
const _AnalyticsEvents = class _AnalyticsEvents extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.AnalyticsEvent events = 1;
     */
    this.events = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsEvents().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsEvents().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsEvents().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsEvents, a, b);
  }
};
_AnalyticsEvents.runtime = import_protobuf.proto3;
_AnalyticsEvents.typeName = "livekit.AnalyticsEvents";
_AnalyticsEvents.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "events", kind: "message", T: AnalyticsEvent, repeated: true }
]);
let AnalyticsEvents = _AnalyticsEvents;
const _AnalyticsRoomParticipant = class _AnalyticsRoomParticipant extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string identity = 2;
     */
    this.identity = "";
    /**
     * @generated from field: string name = 3;
     */
    this.name = "";
    /**
     * @generated from field: livekit.ParticipantInfo.State state = 4;
     */
    this.state = import_livekit_models_pb.ParticipantInfo_State.JOINING;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsRoomParticipant().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsRoomParticipant().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsRoomParticipant().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsRoomParticipant, a, b);
  }
};
_AnalyticsRoomParticipant.runtime = import_protobuf.proto3;
_AnalyticsRoomParticipant.typeName = "livekit.AnalyticsRoomParticipant";
_AnalyticsRoomParticipant.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "state", kind: "enum", T: import_protobuf.proto3.getEnumType(import_livekit_models_pb.ParticipantInfo_State) },
  { no: 5, name: "joined_at", kind: "message", T: import_protobuf.Timestamp }
]);
let AnalyticsRoomParticipant = _AnalyticsRoomParticipant;
const _AnalyticsRoom = class _AnalyticsRoom extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * @generated from field: string project_id = 5;
     */
    this.projectId = "";
    /**
     * @generated from field: repeated livekit.AnalyticsRoomParticipant participants = 4;
     */
    this.participants = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsRoom().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsRoom().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsRoom().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsRoom, a, b);
  }
};
_AnalyticsRoom.runtime = import_protobuf.proto3;
_AnalyticsRoom.typeName = "livekit.AnalyticsRoom";
_AnalyticsRoom.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "project_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "created_at", kind: "message", T: import_protobuf.Timestamp },
  { no: 4, name: "participants", kind: "message", T: AnalyticsRoomParticipant, repeated: true }
]);
let AnalyticsRoom = _AnalyticsRoom;
const _AnalyticsNodeRooms = class _AnalyticsNodeRooms extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string node_id = 1;
     */
    this.nodeId = "";
    /**
     * @generated from field: uint64 sequence_number = 2;
     */
    this.sequenceNumber = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: repeated livekit.AnalyticsRoom rooms = 4;
     */
    this.rooms = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AnalyticsNodeRooms().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AnalyticsNodeRooms().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AnalyticsNodeRooms().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AnalyticsNodeRooms, a, b);
  }
};
_AnalyticsNodeRooms.runtime = import_protobuf.proto3;
_AnalyticsNodeRooms.typeName = "livekit.AnalyticsNodeRooms";
_AnalyticsNodeRooms.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "node_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "sequence_number",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  { no: 3, name: "timestamp", kind: "message", T: import_protobuf.Timestamp },
  { no: 4, name: "rooms", kind: "message", T: AnalyticsRoom, repeated: true }
]);
let AnalyticsNodeRooms = _AnalyticsNodeRooms;
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  AnalyticsClientMeta,
  AnalyticsEvent,
  AnalyticsEventType,
  AnalyticsEvents,
  AnalyticsNodeRooms,
  AnalyticsRoom,
  AnalyticsRoomParticipant,
  AnalyticsStat,
  AnalyticsStats,
  AnalyticsStream,
  AnalyticsVideoLayer,
  StreamType
});
//# sourceMappingURL=livekit_analytics_pb.cjs.map