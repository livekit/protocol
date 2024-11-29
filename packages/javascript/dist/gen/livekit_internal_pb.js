import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { ClientInfo, PlayoutDelay, ReconnectReason } from "./livekit_models_pb.js";
import { CreateRoomRequest } from "./livekit_room_pb.js";
import { AutoParticipantEgress, AutoTrackEgress } from "./livekit_egress_pb.js";
import { RoomAgentDispatch } from "./livekit_agent_dispatch_pb.js";
var NodeType = /* @__PURE__ */ ((NodeType2) => {
  NodeType2[NodeType2["SERVER"] = 0] = "SERVER";
  NodeType2[NodeType2["CONTROLLER"] = 1] = "CONTROLLER";
  NodeType2[NodeType2["MEDIA"] = 2] = "MEDIA";
  NodeType2[NodeType2["TURN"] = 4] = "TURN";
  NodeType2[NodeType2["SWEEPER"] = 5] = "SWEEPER";
  NodeType2[NodeType2["DIRECTOR"] = 6] = "DIRECTOR";
  return NodeType2;
})(NodeType || {});
proto3.util.setEnumType(NodeType, "livekit.NodeType", [
  { no: 0, name: "SERVER" },
  { no: 1, name: "CONTROLLER" },
  { no: 2, name: "MEDIA" },
  { no: 4, name: "TURN" },
  { no: 5, name: "SWEEPER" },
  { no: 6, name: "DIRECTOR" }
]);
var NodeState = /* @__PURE__ */ ((NodeState2) => {
  NodeState2[NodeState2["STARTING_UP"] = 0] = "STARTING_UP";
  NodeState2[NodeState2["SERVING"] = 1] = "SERVING";
  NodeState2[NodeState2["SHUTTING_DOWN"] = 2] = "SHUTTING_DOWN";
  return NodeState2;
})(NodeState || {});
proto3.util.setEnumType(NodeState, "livekit.NodeState", [
  { no: 0, name: "STARTING_UP" },
  { no: 1, name: "SERVING" },
  { no: 2, name: "SHUTTING_DOWN" }
]);
var ICECandidateType = /* @__PURE__ */ ((ICECandidateType2) => {
  ICECandidateType2[ICECandidateType2["ICT_NONE"] = 0] = "ICT_NONE";
  ICECandidateType2[ICECandidateType2["ICT_TCP"] = 1] = "ICT_TCP";
  ICECandidateType2[ICECandidateType2["ICT_TLS"] = 2] = "ICT_TLS";
  return ICECandidateType2;
})(ICECandidateType || {});
proto3.util.setEnumType(ICECandidateType, "livekit.ICECandidateType", [
  { no: 0, name: "ICT_NONE" },
  { no: 1, name: "ICT_TCP" },
  { no: 2, name: "ICT_TLS" }
]);
const _Node = class _Node extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string ip = 2;
     */
    this.ip = "";
    /**
     * @generated from field: uint32 num_cpus = 3;
     */
    this.numCpus = 0;
    /**
     * @generated from field: livekit.NodeType type = 5;
     */
    this.type = 0 /* SERVER */;
    /**
     * @generated from field: livekit.NodeState state = 6;
     */
    this.state = 0 /* STARTING_UP */;
    /**
     * @generated from field: string region = 7;
     */
    this.region = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Node().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Node().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Node().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_Node, a, b);
  }
};
_Node.runtime = proto3;
_Node.typeName = "livekit.Node";
_Node.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "ip",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "num_cpus",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 4, name: "stats", kind: "message", T: NodeStats },
  { no: 5, name: "type", kind: "enum", T: proto3.getEnumType(NodeType) },
  { no: 6, name: "state", kind: "enum", T: proto3.getEnumType(NodeState) },
  {
    no: 7,
    name: "region",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let Node = _Node;
const _NodeStats = class _NodeStats extends Message {
  constructor(data) {
    super();
    /**
     * when server was started
     *
     * @generated from field: int64 started_at = 1;
     */
    this.startedAt = protoInt64.zero;
    /**
     * when server last reported its status
     *
     * @generated from field: int64 updated_at = 2;
     */
    this.updatedAt = protoInt64.zero;
    /**
     * room
     *
     * @generated from field: int32 num_rooms = 3;
     */
    this.numRooms = 0;
    /**
     * @generated from field: int32 num_clients = 4;
     */
    this.numClients = 0;
    /**
     * @generated from field: int32 num_tracks_in = 5;
     */
    this.numTracksIn = 0;
    /**
     * @generated from field: int32 num_tracks_out = 6;
     */
    this.numTracksOut = 0;
    /**
     * @generated from field: int32 num_track_publish_attempts = 36;
     */
    this.numTrackPublishAttempts = 0;
    /**
     * @generated from field: float track_publish_attempts_per_sec = 37;
     */
    this.trackPublishAttemptsPerSec = 0;
    /**
     * @generated from field: int32 num_track_publish_success = 38;
     */
    this.numTrackPublishSuccess = 0;
    /**
     * @generated from field: float track_publish_success_per_sec = 39;
     */
    this.trackPublishSuccessPerSec = 0;
    /**
     * @generated from field: int32 num_track_subscribe_attempts = 40;
     */
    this.numTrackSubscribeAttempts = 0;
    /**
     * @generated from field: float track_subscribe_attempts_per_sec = 41;
     */
    this.trackSubscribeAttemptsPerSec = 0;
    /**
     * @generated from field: int32 num_track_subscribe_success = 42;
     */
    this.numTrackSubscribeSuccess = 0;
    /**
     * @generated from field: float track_subscribe_success_per_sec = 43;
     */
    this.trackSubscribeSuccessPerSec = 0;
    /**
     * packet
     *
     * @generated from field: uint64 bytes_in = 7;
     */
    this.bytesIn = protoInt64.zero;
    /**
     * @generated from field: uint64 bytes_out = 8;
     */
    this.bytesOut = protoInt64.zero;
    /**
     * @generated from field: uint64 packets_in = 9;
     */
    this.packetsIn = protoInt64.zero;
    /**
     * @generated from field: uint64 packets_out = 10;
     */
    this.packetsOut = protoInt64.zero;
    /**
     * @generated from field: uint64 nack_total = 11;
     */
    this.nackTotal = protoInt64.zero;
    /**
     * @generated from field: float bytes_in_per_sec = 12;
     */
    this.bytesInPerSec = 0;
    /**
     * @generated from field: float bytes_out_per_sec = 13;
     */
    this.bytesOutPerSec = 0;
    /**
     * @generated from field: float packets_in_per_sec = 14;
     */
    this.packetsInPerSec = 0;
    /**
     * @generated from field: float packets_out_per_sec = 15;
     */
    this.packetsOutPerSec = 0;
    /**
     * @generated from field: float nack_per_sec = 16;
     */
    this.nackPerSec = 0;
    /**
     * system
     *
     * @generated from field: uint32 num_cpus = 17;
     */
    this.numCpus = 0;
    /**
     * @generated from field: float load_avg_last1min = 18;
     */
    this.loadAvgLast1min = 0;
    /**
     * @generated from field: float load_avg_last5min = 19;
     */
    this.loadAvgLast5min = 0;
    /**
     * @generated from field: float load_avg_last15min = 20;
     */
    this.loadAvgLast15min = 0;
    /**
     * @generated from field: float cpu_load = 21;
     */
    this.cpuLoad = 0;
    /**
     * deprecated
     *
     * @generated from field: float memory_load = 33;
     */
    this.memoryLoad = 0;
    /**
     * @generated from field: uint64 memory_total = 34;
     */
    this.memoryTotal = protoInt64.zero;
    /**
     * @generated from field: uint64 memory_used = 35;
     */
    this.memoryUsed = protoInt64.zero;
    /**
     * @generated from field: uint32 sys_packets_out = 28;
     */
    this.sysPacketsOut = 0;
    /**
     * @generated from field: uint32 sys_packets_dropped = 29;
     */
    this.sysPacketsDropped = 0;
    /**
     * @generated from field: float sys_packets_out_per_sec = 30;
     */
    this.sysPacketsOutPerSec = 0;
    /**
     * @generated from field: float sys_packets_dropped_per_sec = 31;
     */
    this.sysPacketsDroppedPerSec = 0;
    /**
     * @generated from field: float sys_packets_dropped_pct_per_sec = 32;
     */
    this.sysPacketsDroppedPctPerSec = 0;
    /**
     * retransmissions
     *
     * @generated from field: uint64 retransmit_bytes_out = 22;
     */
    this.retransmitBytesOut = protoInt64.zero;
    /**
     * @generated from field: uint64 retransmit_packets_out = 23;
     */
    this.retransmitPacketsOut = protoInt64.zero;
    /**
     * @generated from field: float retransmit_bytes_out_per_sec = 24;
     */
    this.retransmitBytesOutPerSec = 0;
    /**
     * @generated from field: float retransmit_packets_out_per_sec = 25;
     */
    this.retransmitPacketsOutPerSec = 0;
    /**
     * participant joins
     *
     * @generated from field: uint64 participant_signal_connected = 26;
     */
    this.participantSignalConnected = protoInt64.zero;
    /**
     * @generated from field: float participant_signal_connected_per_sec = 27;
     */
    this.participantSignalConnectedPerSec = 0;
    /**
     * @generated from field: uint64 participant_rtc_connected = 44;
     */
    this.participantRtcConnected = protoInt64.zero;
    /**
     * @generated from field: float participant_rtc_connected_per_sec = 45;
     */
    this.participantRtcConnectedPerSec = 0;
    /**
     * @generated from field: uint64 participant_rtc_init = 46;
     */
    this.participantRtcInit = protoInt64.zero;
    /**
     * @generated from field: float participant_rtc_init_per_sec = 47;
     */
    this.participantRtcInitPerSec = 0;
    /**
     * forward metrics
     *
     * @generated from field: uint32 forward_latency = 48;
     */
    this.forwardLatency = 0;
    /**
     * @generated from field: uint32 forward_jitter = 49;
     */
    this.forwardJitter = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _NodeStats().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _NodeStats().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _NodeStats().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_NodeStats, a, b);
  }
};
_NodeStats.runtime = proto3;
_NodeStats.typeName = "livekit.NodeStats";
_NodeStats.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "started_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 2,
    name: "updated_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 3,
    name: "num_rooms",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 4,
    name: "num_clients",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 5,
    name: "num_tracks_in",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 6,
    name: "num_tracks_out",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 36,
    name: "num_track_publish_attempts",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 37,
    name: "track_publish_attempts_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 38,
    name: "num_track_publish_success",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 39,
    name: "track_publish_success_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 40,
    name: "num_track_subscribe_attempts",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 41,
    name: "track_subscribe_attempts_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 42,
    name: "num_track_subscribe_success",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 43,
    name: "track_subscribe_success_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 7,
    name: "bytes_in",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 8,
    name: "bytes_out",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 9,
    name: "packets_in",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 10,
    name: "packets_out",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 11,
    name: "nack_total",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 12,
    name: "bytes_in_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 13,
    name: "bytes_out_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 14,
    name: "packets_in_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 15,
    name: "packets_out_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 16,
    name: "nack_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 17,
    name: "num_cpus",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 18,
    name: "load_avg_last1min",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 19,
    name: "load_avg_last5min",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 20,
    name: "load_avg_last15min",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 21,
    name: "cpu_load",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 33,
    name: "memory_load",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 34,
    name: "memory_total",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 35,
    name: "memory_used",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 28,
    name: "sys_packets_out",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 29,
    name: "sys_packets_dropped",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 30,
    name: "sys_packets_out_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 31,
    name: "sys_packets_dropped_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 32,
    name: "sys_packets_dropped_pct_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 22,
    name: "retransmit_bytes_out",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 23,
    name: "retransmit_packets_out",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 24,
    name: "retransmit_bytes_out_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 25,
    name: "retransmit_packets_out_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 26,
    name: "participant_signal_connected",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 27,
    name: "participant_signal_connected_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 44,
    name: "participant_rtc_connected",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 45,
    name: "participant_rtc_connected_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 46,
    name: "participant_rtc_init",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 47,
    name: "participant_rtc_init_per_sec",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 48,
    name: "forward_latency",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 49,
    name: "forward_jitter",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let NodeStats = _NodeStats;
const _StartSession = class _StartSession extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string room_name = 1;
     */
    this.roomName = "";
    /**
     * @generated from field: string identity = 2;
     */
    this.identity = "";
    /**
     * @generated from field: string connection_id = 3;
     */
    this.connectionId = "";
    /**
     * if a client is reconnecting (i.e. resume instead of restart)
     *
     * @generated from field: bool reconnect = 4;
     */
    this.reconnect = false;
    /**
     * @generated from field: bool auto_subscribe = 9;
     */
    this.autoSubscribe = false;
    /**
     * @generated from field: bool hidden = 10;
     */
    this.hidden = false;
    /**
     * @generated from field: bool recorder = 12;
     */
    this.recorder = false;
    /**
     * @generated from field: string name = 13;
     */
    this.name = "";
    /**
     * A user's ClaimGrants serialized in JSON
     *
     * @generated from field: string grants_json = 14;
     */
    this.grantsJson = "";
    /**
     * @generated from field: bool adaptive_stream = 15;
     */
    this.adaptiveStream = false;
    /**
     * if reconnect, client will set current sid
     *
     * @generated from field: string participant_id = 16;
     */
    this.participantId = "";
    /**
     * @generated from field: livekit.ReconnectReason reconnect_reason = 17;
     */
    this.reconnectReason = ReconnectReason.RR_UNKNOWN;
    /**
     * @generated from field: bool disable_ice_lite = 19;
     */
    this.disableIceLite = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _StartSession().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _StartSession().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _StartSession().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_StartSession, a, b);
  }
};
_StartSession.runtime = proto3;
_StartSession.typeName = "livekit.StartSession";
_StartSession.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_name",
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
    name: "connection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "reconnect",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 9,
    name: "auto_subscribe",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 10,
    name: "hidden",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 11, name: "client", kind: "message", T: ClientInfo },
  {
    no: 12,
    name: "recorder",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 13,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 14,
    name: "grants_json",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 15,
    name: "adaptive_stream",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 16,
    name: "participant_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 17, name: "reconnect_reason", kind: "enum", T: proto3.getEnumType(ReconnectReason) },
  { no: 18, name: "subscriber_allow_pause", kind: "scalar", T: 8, opt: true },
  {
    no: 19,
    name: "disable_ice_lite",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 20, name: "create_room", kind: "message", T: CreateRoomRequest }
]);
let StartSession = _StartSession;
const _RoomInternal = class _RoomInternal extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.RoomAgentDispatch agent_dispatches = 5;
     */
    this.agentDispatches = [];
    /**
     * @generated from field: bool sync_streams = 4;
     */
    this.syncStreams = false;
    /**
     * @generated from field: bool replay_enabled = 6;
     */
    this.replayEnabled = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RoomInternal().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RoomInternal().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RoomInternal().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_RoomInternal, a, b);
  }
};
_RoomInternal.runtime = proto3;
_RoomInternal.typeName = "livekit.RoomInternal";
_RoomInternal.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "track_egress", kind: "message", T: AutoTrackEgress },
  { no: 2, name: "participant_egress", kind: "message", T: AutoParticipantEgress },
  { no: 3, name: "playout_delay", kind: "message", T: PlayoutDelay },
  { no: 5, name: "agent_dispatches", kind: "message", T: RoomAgentDispatch, repeated: true },
  {
    no: 4,
    name: "sync_streams",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 6,
    name: "replay_enabled",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let RoomInternal = _RoomInternal;
const _ICEConfig = class _ICEConfig extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.ICECandidateType preference_subscriber = 1;
     */
    this.preferenceSubscriber = 0 /* ICT_NONE */;
    /**
     * @generated from field: livekit.ICECandidateType preference_publisher = 2;
     */
    this.preferencePublisher = 0 /* ICT_NONE */;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ICEConfig().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ICEConfig().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ICEConfig().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ICEConfig, a, b);
  }
};
_ICEConfig.runtime = proto3;
_ICEConfig.typeName = "livekit.ICEConfig";
_ICEConfig.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "preference_subscriber", kind: "enum", T: proto3.getEnumType(ICECandidateType) },
  { no: 2, name: "preference_publisher", kind: "enum", T: proto3.getEnumType(ICECandidateType) }
]);
let ICEConfig = _ICEConfig;
export {
  ICECandidateType,
  ICEConfig,
  Node,
  NodeState,
  NodeStats,
  NodeType,
  RoomInternal,
  StartSession
};
//# sourceMappingURL=livekit_internal_pb.js.map