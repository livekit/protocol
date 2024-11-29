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
var livekit_models_pb_exports = {};
__export(livekit_models_pb_exports, {
  ActiveSpeakerUpdate: () => ActiveSpeakerUpdate,
  AudioCodec: () => AudioCodec,
  AudioTrackFeature: () => AudioTrackFeature,
  ChatMessage: () => ChatMessage,
  ClientConfigSetting: () => ClientConfigSetting,
  ClientConfiguration: () => ClientConfiguration,
  ClientInfo: () => ClientInfo,
  ClientInfo_SDK: () => ClientInfo_SDK,
  Codec: () => Codec,
  ConnectionQuality: () => ConnectionQuality,
  DataPacket: () => DataPacket,
  DataPacket_Kind: () => DataPacket_Kind,
  DataStream: () => DataStream,
  DataStream_Chunk: () => DataStream_Chunk,
  DataStream_FileHeader: () => DataStream_FileHeader,
  DataStream_Header: () => DataStream_Header,
  DataStream_OperationType: () => DataStream_OperationType,
  DataStream_TextHeader: () => DataStream_TextHeader,
  DisabledCodecs: () => DisabledCodecs,
  DisconnectReason: () => DisconnectReason,
  Encryption: () => Encryption,
  Encryption_Type: () => Encryption_Type,
  ImageCodec: () => ImageCodec,
  ParticipantInfo: () => ParticipantInfo,
  ParticipantInfo_Kind: () => ParticipantInfo_Kind,
  ParticipantInfo_State: () => ParticipantInfo_State,
  ParticipantPermission: () => ParticipantPermission,
  ParticipantTracks: () => ParticipantTracks,
  PlayoutDelay: () => PlayoutDelay,
  RTCPSenderReportState: () => RTCPSenderReportState,
  RTPDrift: () => RTPDrift,
  RTPForwarderState: () => RTPForwarderState,
  RTPMungerState: () => RTPMungerState,
  RTPStats: () => RTPStats,
  ReconnectReason: () => ReconnectReason,
  Room: () => Room,
  RpcAck: () => RpcAck,
  RpcError: () => RpcError,
  RpcRequest: () => RpcRequest,
  RpcResponse: () => RpcResponse,
  ServerInfo: () => ServerInfo,
  ServerInfo_Edition: () => ServerInfo_Edition,
  SimulcastCodecInfo: () => SimulcastCodecInfo,
  SipDTMF: () => SipDTMF,
  SpeakerInfo: () => SpeakerInfo,
  SubscriptionError: () => SubscriptionError,
  TimedVersion: () => TimedVersion,
  TrackInfo: () => TrackInfo,
  TrackSource: () => TrackSource,
  TrackType: () => TrackType,
  Transcription: () => Transcription,
  TranscriptionSegment: () => TranscriptionSegment,
  UserPacket: () => UserPacket,
  VP8MungerState: () => VP8MungerState,
  VideoCodec: () => VideoCodec,
  VideoConfiguration: () => VideoConfiguration,
  VideoLayer: () => VideoLayer,
  VideoQuality: () => VideoQuality
});
module.exports = __toCommonJS(livekit_models_pb_exports);
var import_protobuf = require("@bufbuild/protobuf");
var import_livekit_metrics_pb = require("./livekit_metrics_pb.cjs");
var AudioCodec = /* @__PURE__ */ ((AudioCodec2) => {
  AudioCodec2[AudioCodec2["DEFAULT_AC"] = 0] = "DEFAULT_AC";
  AudioCodec2[AudioCodec2["OPUS"] = 1] = "OPUS";
  AudioCodec2[AudioCodec2["AAC"] = 2] = "AAC";
  return AudioCodec2;
})(AudioCodec || {});
import_protobuf.proto3.util.setEnumType(AudioCodec, "livekit.AudioCodec", [
  { no: 0, name: "DEFAULT_AC" },
  { no: 1, name: "OPUS" },
  { no: 2, name: "AAC" }
]);
var VideoCodec = /* @__PURE__ */ ((VideoCodec2) => {
  VideoCodec2[VideoCodec2["DEFAULT_VC"] = 0] = "DEFAULT_VC";
  VideoCodec2[VideoCodec2["H264_BASELINE"] = 1] = "H264_BASELINE";
  VideoCodec2[VideoCodec2["H264_MAIN"] = 2] = "H264_MAIN";
  VideoCodec2[VideoCodec2["H264_HIGH"] = 3] = "H264_HIGH";
  VideoCodec2[VideoCodec2["VP8"] = 4] = "VP8";
  return VideoCodec2;
})(VideoCodec || {});
import_protobuf.proto3.util.setEnumType(VideoCodec, "livekit.VideoCodec", [
  { no: 0, name: "DEFAULT_VC" },
  { no: 1, name: "H264_BASELINE" },
  { no: 2, name: "H264_MAIN" },
  { no: 3, name: "H264_HIGH" },
  { no: 4, name: "VP8" }
]);
var ImageCodec = /* @__PURE__ */ ((ImageCodec2) => {
  ImageCodec2[ImageCodec2["IC_DEFAULT"] = 0] = "IC_DEFAULT";
  ImageCodec2[ImageCodec2["IC_JPEG"] = 1] = "IC_JPEG";
  return ImageCodec2;
})(ImageCodec || {});
import_protobuf.proto3.util.setEnumType(ImageCodec, "livekit.ImageCodec", [
  { no: 0, name: "IC_DEFAULT" },
  { no: 1, name: "IC_JPEG" }
]);
var TrackType = /* @__PURE__ */ ((TrackType2) => {
  TrackType2[TrackType2["AUDIO"] = 0] = "AUDIO";
  TrackType2[TrackType2["VIDEO"] = 1] = "VIDEO";
  TrackType2[TrackType2["DATA"] = 2] = "DATA";
  return TrackType2;
})(TrackType || {});
import_protobuf.proto3.util.setEnumType(TrackType, "livekit.TrackType", [
  { no: 0, name: "AUDIO" },
  { no: 1, name: "VIDEO" },
  { no: 2, name: "DATA" }
]);
var TrackSource = /* @__PURE__ */ ((TrackSource2) => {
  TrackSource2[TrackSource2["UNKNOWN"] = 0] = "UNKNOWN";
  TrackSource2[TrackSource2["CAMERA"] = 1] = "CAMERA";
  TrackSource2[TrackSource2["MICROPHONE"] = 2] = "MICROPHONE";
  TrackSource2[TrackSource2["SCREEN_SHARE"] = 3] = "SCREEN_SHARE";
  TrackSource2[TrackSource2["SCREEN_SHARE_AUDIO"] = 4] = "SCREEN_SHARE_AUDIO";
  return TrackSource2;
})(TrackSource || {});
import_protobuf.proto3.util.setEnumType(TrackSource, "livekit.TrackSource", [
  { no: 0, name: "UNKNOWN" },
  { no: 1, name: "CAMERA" },
  { no: 2, name: "MICROPHONE" },
  { no: 3, name: "SCREEN_SHARE" },
  { no: 4, name: "SCREEN_SHARE_AUDIO" }
]);
var VideoQuality = /* @__PURE__ */ ((VideoQuality2) => {
  VideoQuality2[VideoQuality2["LOW"] = 0] = "LOW";
  VideoQuality2[VideoQuality2["MEDIUM"] = 1] = "MEDIUM";
  VideoQuality2[VideoQuality2["HIGH"] = 2] = "HIGH";
  VideoQuality2[VideoQuality2["OFF"] = 3] = "OFF";
  return VideoQuality2;
})(VideoQuality || {});
import_protobuf.proto3.util.setEnumType(VideoQuality, "livekit.VideoQuality", [
  { no: 0, name: "LOW" },
  { no: 1, name: "MEDIUM" },
  { no: 2, name: "HIGH" },
  { no: 3, name: "OFF" }
]);
var ConnectionQuality = /* @__PURE__ */ ((ConnectionQuality2) => {
  ConnectionQuality2[ConnectionQuality2["POOR"] = 0] = "POOR";
  ConnectionQuality2[ConnectionQuality2["GOOD"] = 1] = "GOOD";
  ConnectionQuality2[ConnectionQuality2["EXCELLENT"] = 2] = "EXCELLENT";
  ConnectionQuality2[ConnectionQuality2["LOST"] = 3] = "LOST";
  return ConnectionQuality2;
})(ConnectionQuality || {});
import_protobuf.proto3.util.setEnumType(ConnectionQuality, "livekit.ConnectionQuality", [
  { no: 0, name: "POOR" },
  { no: 1, name: "GOOD" },
  { no: 2, name: "EXCELLENT" },
  { no: 3, name: "LOST" }
]);
var ClientConfigSetting = /* @__PURE__ */ ((ClientConfigSetting2) => {
  ClientConfigSetting2[ClientConfigSetting2["UNSET"] = 0] = "UNSET";
  ClientConfigSetting2[ClientConfigSetting2["DISABLED"] = 1] = "DISABLED";
  ClientConfigSetting2[ClientConfigSetting2["ENABLED"] = 2] = "ENABLED";
  return ClientConfigSetting2;
})(ClientConfigSetting || {});
import_protobuf.proto3.util.setEnumType(ClientConfigSetting, "livekit.ClientConfigSetting", [
  { no: 0, name: "UNSET" },
  { no: 1, name: "DISABLED" },
  { no: 2, name: "ENABLED" }
]);
var DisconnectReason = /* @__PURE__ */ ((DisconnectReason2) => {
  DisconnectReason2[DisconnectReason2["UNKNOWN_REASON"] = 0] = "UNKNOWN_REASON";
  DisconnectReason2[DisconnectReason2["CLIENT_INITIATED"] = 1] = "CLIENT_INITIATED";
  DisconnectReason2[DisconnectReason2["DUPLICATE_IDENTITY"] = 2] = "DUPLICATE_IDENTITY";
  DisconnectReason2[DisconnectReason2["SERVER_SHUTDOWN"] = 3] = "SERVER_SHUTDOWN";
  DisconnectReason2[DisconnectReason2["PARTICIPANT_REMOVED"] = 4] = "PARTICIPANT_REMOVED";
  DisconnectReason2[DisconnectReason2["ROOM_DELETED"] = 5] = "ROOM_DELETED";
  DisconnectReason2[DisconnectReason2["STATE_MISMATCH"] = 6] = "STATE_MISMATCH";
  DisconnectReason2[DisconnectReason2["JOIN_FAILURE"] = 7] = "JOIN_FAILURE";
  DisconnectReason2[DisconnectReason2["MIGRATION"] = 8] = "MIGRATION";
  DisconnectReason2[DisconnectReason2["SIGNAL_CLOSE"] = 9] = "SIGNAL_CLOSE";
  DisconnectReason2[DisconnectReason2["ROOM_CLOSED"] = 10] = "ROOM_CLOSED";
  DisconnectReason2[DisconnectReason2["USER_UNAVAILABLE"] = 11] = "USER_UNAVAILABLE";
  DisconnectReason2[DisconnectReason2["USER_REJECTED"] = 12] = "USER_REJECTED";
  DisconnectReason2[DisconnectReason2["SIP_TRUNK_FAILURE"] = 13] = "SIP_TRUNK_FAILURE";
  return DisconnectReason2;
})(DisconnectReason || {});
import_protobuf.proto3.util.setEnumType(DisconnectReason, "livekit.DisconnectReason", [
  { no: 0, name: "UNKNOWN_REASON" },
  { no: 1, name: "CLIENT_INITIATED" },
  { no: 2, name: "DUPLICATE_IDENTITY" },
  { no: 3, name: "SERVER_SHUTDOWN" },
  { no: 4, name: "PARTICIPANT_REMOVED" },
  { no: 5, name: "ROOM_DELETED" },
  { no: 6, name: "STATE_MISMATCH" },
  { no: 7, name: "JOIN_FAILURE" },
  { no: 8, name: "MIGRATION" },
  { no: 9, name: "SIGNAL_CLOSE" },
  { no: 10, name: "ROOM_CLOSED" },
  { no: 11, name: "USER_UNAVAILABLE" },
  { no: 12, name: "USER_REJECTED" },
  { no: 13, name: "SIP_TRUNK_FAILURE" }
]);
var ReconnectReason = /* @__PURE__ */ ((ReconnectReason2) => {
  ReconnectReason2[ReconnectReason2["RR_UNKNOWN"] = 0] = "RR_UNKNOWN";
  ReconnectReason2[ReconnectReason2["RR_SIGNAL_DISCONNECTED"] = 1] = "RR_SIGNAL_DISCONNECTED";
  ReconnectReason2[ReconnectReason2["RR_PUBLISHER_FAILED"] = 2] = "RR_PUBLISHER_FAILED";
  ReconnectReason2[ReconnectReason2["RR_SUBSCRIBER_FAILED"] = 3] = "RR_SUBSCRIBER_FAILED";
  ReconnectReason2[ReconnectReason2["RR_SWITCH_CANDIDATE"] = 4] = "RR_SWITCH_CANDIDATE";
  return ReconnectReason2;
})(ReconnectReason || {});
import_protobuf.proto3.util.setEnumType(ReconnectReason, "livekit.ReconnectReason", [
  { no: 0, name: "RR_UNKNOWN" },
  { no: 1, name: "RR_SIGNAL_DISCONNECTED" },
  { no: 2, name: "RR_PUBLISHER_FAILED" },
  { no: 3, name: "RR_SUBSCRIBER_FAILED" },
  { no: 4, name: "RR_SWITCH_CANDIDATE" }
]);
var SubscriptionError = /* @__PURE__ */ ((SubscriptionError2) => {
  SubscriptionError2[SubscriptionError2["SE_UNKNOWN"] = 0] = "SE_UNKNOWN";
  SubscriptionError2[SubscriptionError2["SE_CODEC_UNSUPPORTED"] = 1] = "SE_CODEC_UNSUPPORTED";
  SubscriptionError2[SubscriptionError2["SE_TRACK_NOTFOUND"] = 2] = "SE_TRACK_NOTFOUND";
  return SubscriptionError2;
})(SubscriptionError || {});
import_protobuf.proto3.util.setEnumType(SubscriptionError, "livekit.SubscriptionError", [
  { no: 0, name: "SE_UNKNOWN" },
  { no: 1, name: "SE_CODEC_UNSUPPORTED" },
  { no: 2, name: "SE_TRACK_NOTFOUND" }
]);
var AudioTrackFeature = /* @__PURE__ */ ((AudioTrackFeature2) => {
  AudioTrackFeature2[AudioTrackFeature2["TF_STEREO"] = 0] = "TF_STEREO";
  AudioTrackFeature2[AudioTrackFeature2["TF_NO_DTX"] = 1] = "TF_NO_DTX";
  AudioTrackFeature2[AudioTrackFeature2["TF_AUTO_GAIN_CONTROL"] = 2] = "TF_AUTO_GAIN_CONTROL";
  AudioTrackFeature2[AudioTrackFeature2["TF_ECHO_CANCELLATION"] = 3] = "TF_ECHO_CANCELLATION";
  AudioTrackFeature2[AudioTrackFeature2["TF_NOISE_SUPPRESSION"] = 4] = "TF_NOISE_SUPPRESSION";
  AudioTrackFeature2[AudioTrackFeature2["TF_ENHANCED_NOISE_CANCELLATION"] = 5] = "TF_ENHANCED_NOISE_CANCELLATION";
  return AudioTrackFeature2;
})(AudioTrackFeature || {});
import_protobuf.proto3.util.setEnumType(AudioTrackFeature, "livekit.AudioTrackFeature", [
  { no: 0, name: "TF_STEREO" },
  { no: 1, name: "TF_NO_DTX" },
  { no: 2, name: "TF_AUTO_GAIN_CONTROL" },
  { no: 3, name: "TF_ECHO_CANCELLATION" },
  { no: 4, name: "TF_NOISE_SUPPRESSION" },
  { no: 5, name: "TF_ENHANCED_NOISE_CANCELLATION" }
]);
const _Room = class _Room extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sid = 1;
     */
    this.sid = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * @generated from field: uint32 empty_timeout = 3;
     */
    this.emptyTimeout = 0;
    /**
     * @generated from field: uint32 departure_timeout = 14;
     */
    this.departureTimeout = 0;
    /**
     * @generated from field: uint32 max_participants = 4;
     */
    this.maxParticipants = 0;
    /**
     * @generated from field: int64 creation_time = 5;
     */
    this.creationTime = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: string turn_password = 6;
     */
    this.turnPassword = "";
    /**
     * @generated from field: repeated livekit.Codec enabled_codecs = 7;
     */
    this.enabledCodecs = [];
    /**
     * @generated from field: string metadata = 8;
     */
    this.metadata = "";
    /**
     * @generated from field: uint32 num_participants = 9;
     */
    this.numParticipants = 0;
    /**
     * @generated from field: uint32 num_publishers = 11;
     */
    this.numPublishers = 0;
    /**
     * @generated from field: bool active_recording = 10;
     */
    this.activeRecording = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Room().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Room().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Room().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Room, a, b);
  }
};
_Room.runtime = import_protobuf.proto3;
_Room.typeName = "livekit.Room";
_Room.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sid",
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
    no: 3,
    name: "empty_timeout",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 14,
    name: "departure_timeout",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "max_participants",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 5,
    name: "creation_time",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 6,
    name: "turn_password",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 7, name: "enabled_codecs", kind: "message", T: Codec, repeated: true },
  {
    no: 8,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "num_participants",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 11,
    name: "num_publishers",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 10,
    name: "active_recording",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 13, name: "version", kind: "message", T: TimedVersion }
]);
let Room = _Room;
const _Codec = class _Codec extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string mime = 1;
     */
    this.mime = "";
    /**
     * @generated from field: string fmtp_line = 2;
     */
    this.fmtpLine = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Codec().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Codec().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Codec().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Codec, a, b);
  }
};
_Codec.runtime = import_protobuf.proto3;
_Codec.typeName = "livekit.Codec";
_Codec.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "mime",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "fmtp_line",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let Codec = _Codec;
const _PlayoutDelay = class _PlayoutDelay extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool enabled = 1;
     */
    this.enabled = false;
    /**
     * @generated from field: uint32 min = 2;
     */
    this.min = 0;
    /**
     * @generated from field: uint32 max = 3;
     */
    this.max = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _PlayoutDelay().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _PlayoutDelay().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _PlayoutDelay().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_PlayoutDelay, a, b);
  }
};
_PlayoutDelay.runtime = import_protobuf.proto3;
_PlayoutDelay.typeName = "livekit.PlayoutDelay";
_PlayoutDelay.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "enabled",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 2,
    name: "min",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "max",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let PlayoutDelay = _PlayoutDelay;
const _ParticipantPermission = class _ParticipantPermission extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * allow participant to subscribe to other tracks in the room
     *
     * @generated from field: bool can_subscribe = 1;
     */
    this.canSubscribe = false;
    /**
     * allow participant to publish new tracks to room
     *
     * @generated from field: bool can_publish = 2;
     */
    this.canPublish = false;
    /**
     * allow participant to publish data
     *
     * @generated from field: bool can_publish_data = 3;
     */
    this.canPublishData = false;
    /**
     * sources that are allowed to be published
     *
     * @generated from field: repeated livekit.TrackSource can_publish_sources = 9;
     */
    this.canPublishSources = [];
    /**
     * indicates that it's hidden to others
     *
     * @generated from field: bool hidden = 7;
     */
    this.hidden = false;
    /**
     * indicates it's a recorder instance
     * deprecated: use ParticipantInfo.kind instead
     *
     * @generated from field: bool recorder = 8 [deprecated = true];
     * @deprecated
     */
    this.recorder = false;
    /**
     * indicates that participant can update own metadata and attributes
     *
     * @generated from field: bool can_update_metadata = 10;
     */
    this.canUpdateMetadata = false;
    /**
     * indicates that participant is an agent
     * deprecated: use ParticipantInfo.kind instead
     *
     * @generated from field: bool agent = 11 [deprecated = true];
     * @deprecated
     */
    this.agent = false;
    /**
     * if a participant can subscribe to metrics
     *
     * @generated from field: bool can_subscribe_metrics = 12;
     */
    this.canSubscribeMetrics = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ParticipantPermission().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ParticipantPermission().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ParticipantPermission().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ParticipantPermission, a, b);
  }
};
_ParticipantPermission.runtime = import_protobuf.proto3;
_ParticipantPermission.typeName = "livekit.ParticipantPermission";
_ParticipantPermission.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "can_subscribe",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 2,
    name: "can_publish",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 3,
    name: "can_publish_data",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 9, name: "can_publish_sources", kind: "enum", T: import_protobuf.proto3.getEnumType(TrackSource), repeated: true },
  {
    no: 7,
    name: "hidden",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 8,
    name: "recorder",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 10,
    name: "can_update_metadata",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 11,
    name: "agent",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 12,
    name: "can_subscribe_metrics",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let ParticipantPermission = _ParticipantPermission;
const _ParticipantInfo = class _ParticipantInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sid = 1;
     */
    this.sid = "";
    /**
     * @generated from field: string identity = 2;
     */
    this.identity = "";
    /**
     * @generated from field: livekit.ParticipantInfo.State state = 3;
     */
    this.state = 0 /* JOINING */;
    /**
     * @generated from field: repeated livekit.TrackInfo tracks = 4;
     */
    this.tracks = [];
    /**
     * @generated from field: string metadata = 5;
     */
    this.metadata = "";
    /**
     * timestamp when participant joined room, in seconds
     *
     * @generated from field: int64 joined_at = 6;
     */
    this.joinedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: string name = 9;
     */
    this.name = "";
    /**
     * @generated from field: uint32 version = 10;
     */
    this.version = 0;
    /**
     * @generated from field: string region = 12;
     */
    this.region = "";
    /**
     * indicates the participant has an active publisher connection
     * and can publish to the server
     *
     * @generated from field: bool is_publisher = 13;
     */
    this.isPublisher = false;
    /**
     * @generated from field: livekit.ParticipantInfo.Kind kind = 14;
     */
    this.kind = 0 /* STANDARD */;
    /**
     * @generated from field: map<string, string> attributes = 15;
     */
    this.attributes = {};
    /**
     * @generated from field: livekit.DisconnectReason disconnect_reason = 16;
     */
    this.disconnectReason = 0 /* UNKNOWN_REASON */;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ParticipantInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ParticipantInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ParticipantInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ParticipantInfo, a, b);
  }
};
_ParticipantInfo.runtime = import_protobuf.proto3;
_ParticipantInfo.typeName = "livekit.ParticipantInfo";
_ParticipantInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sid",
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
  { no: 3, name: "state", kind: "enum", T: import_protobuf.proto3.getEnumType(ParticipantInfo_State) },
  { no: 4, name: "tracks", kind: "message", T: TrackInfo, repeated: true },
  {
    no: 5,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "joined_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 9,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 10,
    name: "version",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 11, name: "permission", kind: "message", T: ParticipantPermission },
  {
    no: 12,
    name: "region",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 13,
    name: "is_publisher",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 14, name: "kind", kind: "enum", T: import_protobuf.proto3.getEnumType(ParticipantInfo_Kind) },
  { no: 15, name: "attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  { no: 16, name: "disconnect_reason", kind: "enum", T: import_protobuf.proto3.getEnumType(DisconnectReason) }
]);
let ParticipantInfo = _ParticipantInfo;
var ParticipantInfo_State = /* @__PURE__ */ ((ParticipantInfo_State2) => {
  ParticipantInfo_State2[ParticipantInfo_State2["JOINING"] = 0] = "JOINING";
  ParticipantInfo_State2[ParticipantInfo_State2["JOINED"] = 1] = "JOINED";
  ParticipantInfo_State2[ParticipantInfo_State2["ACTIVE"] = 2] = "ACTIVE";
  ParticipantInfo_State2[ParticipantInfo_State2["DISCONNECTED"] = 3] = "DISCONNECTED";
  return ParticipantInfo_State2;
})(ParticipantInfo_State || {});
import_protobuf.proto3.util.setEnumType(ParticipantInfo_State, "livekit.ParticipantInfo.State", [
  { no: 0, name: "JOINING" },
  { no: 1, name: "JOINED" },
  { no: 2, name: "ACTIVE" },
  { no: 3, name: "DISCONNECTED" }
]);
var ParticipantInfo_Kind = /* @__PURE__ */ ((ParticipantInfo_Kind2) => {
  ParticipantInfo_Kind2[ParticipantInfo_Kind2["STANDARD"] = 0] = "STANDARD";
  ParticipantInfo_Kind2[ParticipantInfo_Kind2["INGRESS"] = 1] = "INGRESS";
  ParticipantInfo_Kind2[ParticipantInfo_Kind2["EGRESS"] = 2] = "EGRESS";
  ParticipantInfo_Kind2[ParticipantInfo_Kind2["SIP"] = 3] = "SIP";
  ParticipantInfo_Kind2[ParticipantInfo_Kind2["AGENT"] = 4] = "AGENT";
  return ParticipantInfo_Kind2;
})(ParticipantInfo_Kind || {});
import_protobuf.proto3.util.setEnumType(ParticipantInfo_Kind, "livekit.ParticipantInfo.Kind", [
  { no: 0, name: "STANDARD" },
  { no: 1, name: "INGRESS" },
  { no: 2, name: "EGRESS" },
  { no: 3, name: "SIP" },
  { no: 4, name: "AGENT" }
]);
const _Encryption = class _Encryption extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Encryption().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Encryption().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Encryption().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Encryption, a, b);
  }
};
_Encryption.runtime = import_protobuf.proto3;
_Encryption.typeName = "livekit.Encryption";
_Encryption.fields = import_protobuf.proto3.util.newFieldList(() => []);
let Encryption = _Encryption;
var Encryption_Type = /* @__PURE__ */ ((Encryption_Type2) => {
  Encryption_Type2[Encryption_Type2["NONE"] = 0] = "NONE";
  Encryption_Type2[Encryption_Type2["GCM"] = 1] = "GCM";
  Encryption_Type2[Encryption_Type2["CUSTOM"] = 2] = "CUSTOM";
  return Encryption_Type2;
})(Encryption_Type || {});
import_protobuf.proto3.util.setEnumType(Encryption_Type, "livekit.Encryption.Type", [
  { no: 0, name: "NONE" },
  { no: 1, name: "GCM" },
  { no: 2, name: "CUSTOM" }
]);
const _SimulcastCodecInfo = class _SimulcastCodecInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string mime_type = 1;
     */
    this.mimeType = "";
    /**
     * @generated from field: string mid = 2;
     */
    this.mid = "";
    /**
     * @generated from field: string cid = 3;
     */
    this.cid = "";
    /**
     * @generated from field: repeated livekit.VideoLayer layers = 4;
     */
    this.layers = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SimulcastCodecInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SimulcastCodecInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SimulcastCodecInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_SimulcastCodecInfo, a, b);
  }
};
_SimulcastCodecInfo.runtime = import_protobuf.proto3;
_SimulcastCodecInfo.typeName = "livekit.SimulcastCodecInfo";
_SimulcastCodecInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "mime_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "mid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "cid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "layers", kind: "message", T: VideoLayer, repeated: true }
]);
let SimulcastCodecInfo = _SimulcastCodecInfo;
const _TrackInfo = class _TrackInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sid = 1;
     */
    this.sid = "";
    /**
     * @generated from field: livekit.TrackType type = 2;
     */
    this.type = 0 /* AUDIO */;
    /**
     * @generated from field: string name = 3;
     */
    this.name = "";
    /**
     * @generated from field: bool muted = 4;
     */
    this.muted = false;
    /**
     * original width of video (unset for audio)
     * clients may receive a lower resolution version with simulcast
     *
     * @generated from field: uint32 width = 5;
     */
    this.width = 0;
    /**
     * original height of video (unset for audio)
     *
     * @generated from field: uint32 height = 6;
     */
    this.height = 0;
    /**
     * true if track is simulcasted
     *
     * @generated from field: bool simulcast = 7;
     */
    this.simulcast = false;
    /**
     * true if DTX (Discontinuous Transmission) is disabled for audio
     *
     * @generated from field: bool disable_dtx = 8;
     */
    this.disableDtx = false;
    /**
     * source of media
     *
     * @generated from field: livekit.TrackSource source = 9;
     */
    this.source = 0 /* UNKNOWN */;
    /**
     * @generated from field: repeated livekit.VideoLayer layers = 10;
     */
    this.layers = [];
    /**
     * mime type of codec
     *
     * @generated from field: string mime_type = 11;
     */
    this.mimeType = "";
    /**
     * @generated from field: string mid = 12;
     */
    this.mid = "";
    /**
     * @generated from field: repeated livekit.SimulcastCodecInfo codecs = 13;
     */
    this.codecs = [];
    /**
     * @generated from field: bool stereo = 14;
     */
    this.stereo = false;
    /**
     * true if RED (Redundant Encoding) is disabled for audio
     *
     * @generated from field: bool disable_red = 15;
     */
    this.disableRed = false;
    /**
     * @generated from field: livekit.Encryption.Type encryption = 16;
     */
    this.encryption = 0 /* NONE */;
    /**
     * @generated from field: string stream = 17;
     */
    this.stream = "";
    /**
     * @generated from field: repeated livekit.AudioTrackFeature audio_features = 19;
     */
    this.audioFeatures = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TrackInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TrackInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TrackInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_TrackInfo, a, b);
  }
};
_TrackInfo.runtime = import_protobuf.proto3;
_TrackInfo.typeName = "livekit.TrackInfo";
_TrackInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "type", kind: "enum", T: import_protobuf.proto3.getEnumType(TrackType) },
  {
    no: 3,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "muted",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 5,
    name: "width",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 6,
    name: "height",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 7,
    name: "simulcast",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 8,
    name: "disable_dtx",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 9, name: "source", kind: "enum", T: import_protobuf.proto3.getEnumType(TrackSource) },
  { no: 10, name: "layers", kind: "message", T: VideoLayer, repeated: true },
  {
    no: 11,
    name: "mime_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 12,
    name: "mid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 13, name: "codecs", kind: "message", T: SimulcastCodecInfo, repeated: true },
  {
    no: 14,
    name: "stereo",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 15,
    name: "disable_red",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 16, name: "encryption", kind: "enum", T: import_protobuf.proto3.getEnumType(Encryption_Type) },
  {
    no: 17,
    name: "stream",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 18, name: "version", kind: "message", T: TimedVersion },
  { no: 19, name: "audio_features", kind: "enum", T: import_protobuf.proto3.getEnumType(AudioTrackFeature), repeated: true }
]);
let TrackInfo = _TrackInfo;
const _VideoLayer = class _VideoLayer extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * for tracks with a single layer, this should be HIGH
     *
     * @generated from field: livekit.VideoQuality quality = 1;
     */
    this.quality = 0 /* LOW */;
    /**
     * @generated from field: uint32 width = 2;
     */
    this.width = 0;
    /**
     * @generated from field: uint32 height = 3;
     */
    this.height = 0;
    /**
     * target bitrate in bit per second (bps), server will measure actual
     *
     * @generated from field: uint32 bitrate = 4;
     */
    this.bitrate = 0;
    /**
     * @generated from field: uint32 ssrc = 5;
     */
    this.ssrc = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _VideoLayer().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _VideoLayer().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _VideoLayer().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_VideoLayer, a, b);
  }
};
_VideoLayer.runtime = import_protobuf.proto3;
_VideoLayer.typeName = "livekit.VideoLayer";
_VideoLayer.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "quality", kind: "enum", T: import_protobuf.proto3.getEnumType(VideoQuality) },
  {
    no: 2,
    name: "width",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "height",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "bitrate",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 5,
    name: "ssrc",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let VideoLayer = _VideoLayer;
const _DataPacket = class _DataPacket extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.DataPacket.Kind kind = 1 [deprecated = true];
     * @deprecated
     */
    this.kind = 0 /* RELIABLE */;
    /**
     * participant identity of user that sent the message
     *
     * @generated from field: string participant_identity = 4;
     */
    this.participantIdentity = "";
    /**
     * identities of participants who will receive the message (sent to all by default)
     *
     * @generated from field: repeated string destination_identities = 5;
     */
    this.destinationIdentities = [];
    /**
     * @generated from oneof livekit.DataPacket.value
     */
    this.value = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DataPacket().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DataPacket().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DataPacket().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DataPacket, a, b);
  }
};
_DataPacket.runtime = import_protobuf.proto3;
_DataPacket.typeName = "livekit.DataPacket";
_DataPacket.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "kind", kind: "enum", T: import_protobuf.proto3.getEnumType(DataPacket_Kind) },
  {
    no: 4,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "destination_identities", kind: "scalar", T: 9, repeated: true },
  { no: 2, name: "user", kind: "message", T: UserPacket, oneof: "value" },
  { no: 3, name: "speaker", kind: "message", T: ActiveSpeakerUpdate, oneof: "value" },
  { no: 6, name: "sip_dtmf", kind: "message", T: SipDTMF, oneof: "value" },
  { no: 7, name: "transcription", kind: "message", T: Transcription, oneof: "value" },
  { no: 8, name: "metrics", kind: "message", T: import_livekit_metrics_pb.MetricsBatch, oneof: "value" },
  { no: 9, name: "chat_message", kind: "message", T: ChatMessage, oneof: "value" },
  { no: 10, name: "rpc_request", kind: "message", T: RpcRequest, oneof: "value" },
  { no: 11, name: "rpc_ack", kind: "message", T: RpcAck, oneof: "value" },
  { no: 12, name: "rpc_response", kind: "message", T: RpcResponse, oneof: "value" },
  { no: 13, name: "stream_header", kind: "message", T: DataStream_Header, oneof: "value" },
  { no: 14, name: "stream_chunk", kind: "message", T: DataStream_Chunk, oneof: "value" }
]);
let DataPacket = _DataPacket;
var DataPacket_Kind = /* @__PURE__ */ ((DataPacket_Kind2) => {
  DataPacket_Kind2[DataPacket_Kind2["RELIABLE"] = 0] = "RELIABLE";
  DataPacket_Kind2[DataPacket_Kind2["LOSSY"] = 1] = "LOSSY";
  return DataPacket_Kind2;
})(DataPacket_Kind || {});
import_protobuf.proto3.util.setEnumType(DataPacket_Kind, "livekit.DataPacket.Kind", [
  { no: 0, name: "RELIABLE" },
  { no: 1, name: "LOSSY" }
]);
const _ActiveSpeakerUpdate = class _ActiveSpeakerUpdate extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.SpeakerInfo speakers = 1;
     */
    this.speakers = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ActiveSpeakerUpdate().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ActiveSpeakerUpdate().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ActiveSpeakerUpdate().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ActiveSpeakerUpdate, a, b);
  }
};
_ActiveSpeakerUpdate.runtime = import_protobuf.proto3;
_ActiveSpeakerUpdate.typeName = "livekit.ActiveSpeakerUpdate";
_ActiveSpeakerUpdate.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "speakers", kind: "message", T: SpeakerInfo, repeated: true }
]);
let ActiveSpeakerUpdate = _ActiveSpeakerUpdate;
const _SpeakerInfo = class _SpeakerInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sid = 1;
     */
    this.sid = "";
    /**
     * audio level, 0-1.0, 1 is loudest
     *
     * @generated from field: float level = 2;
     */
    this.level = 0;
    /**
     * true if speaker is currently active
     *
     * @generated from field: bool active = 3;
     */
    this.active = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpeakerInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpeakerInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpeakerInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_SpeakerInfo, a, b);
  }
};
_SpeakerInfo.runtime = import_protobuf.proto3;
_SpeakerInfo.typeName = "livekit.SpeakerInfo";
_SpeakerInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "level",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 3,
    name: "active",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let SpeakerInfo = _SpeakerInfo;
const _UserPacket = class _UserPacket extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * participant ID of user that sent the message
     *
     * @generated from field: string participant_sid = 1 [deprecated = true];
     * @deprecated
     */
    this.participantSid = "";
    /**
     * @generated from field: string participant_identity = 5 [deprecated = true];
     * @deprecated
     */
    this.participantIdentity = "";
    /**
     * user defined payload
     *
     * @generated from field: bytes payload = 2;
     */
    this.payload = new Uint8Array(0);
    /**
     * the ID of the participants who will receive the message (sent to all by default)
     *
     * @generated from field: repeated string destination_sids = 3 [deprecated = true];
     * @deprecated
     */
    this.destinationSids = [];
    /**
     * identities of participants who will receive the message (sent to all by default)
     *
     * @generated from field: repeated string destination_identities = 6 [deprecated = true];
     * @deprecated
     */
    this.destinationIdentities = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UserPacket().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UserPacket().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UserPacket().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UserPacket, a, b);
  }
};
_UserPacket.runtime = import_protobuf.proto3;
_UserPacket.typeName = "livekit.UserPacket";
_UserPacket.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "participant_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "payload",
    kind: "scalar",
    T: 12
    /* ScalarType.BYTES */
  },
  { no: 3, name: "destination_sids", kind: "scalar", T: 9, repeated: true },
  { no: 6, name: "destination_identities", kind: "scalar", T: 9, repeated: true },
  { no: 4, name: "topic", kind: "scalar", T: 9, opt: true },
  { no: 8, name: "id", kind: "scalar", T: 9, opt: true },
  { no: 9, name: "start_time", kind: "scalar", T: 4, opt: true },
  { no: 10, name: "end_time", kind: "scalar", T: 4, opt: true }
]);
let UserPacket = _UserPacket;
const _SipDTMF = class _SipDTMF extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: uint32 code = 3;
     */
    this.code = 0;
    /**
     * @generated from field: string digit = 4;
     */
    this.digit = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SipDTMF().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SipDTMF().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SipDTMF().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_SipDTMF, a, b);
  }
};
_SipDTMF.runtime = import_protobuf.proto3;
_SipDTMF.typeName = "livekit.SipDTMF";
_SipDTMF.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 3,
    name: "code",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "digit",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let SipDTMF = _SipDTMF;
const _Transcription = class _Transcription extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * Participant that got its speech transcribed
     *
     * @generated from field: string transcribed_participant_identity = 2;
     */
    this.transcribedParticipantIdentity = "";
    /**
     * @generated from field: string track_id = 3;
     */
    this.trackId = "";
    /**
     * @generated from field: repeated livekit.TranscriptionSegment segments = 4;
     */
    this.segments = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Transcription().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Transcription().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Transcription().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Transcription, a, b);
  }
};
_Transcription.runtime = import_protobuf.proto3;
_Transcription.typeName = "livekit.Transcription";
_Transcription.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 2,
    name: "transcribed_participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "track_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "segments", kind: "message", T: TranscriptionSegment, repeated: true }
]);
let Transcription = _Transcription;
const _TranscriptionSegment = class _TranscriptionSegment extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string text = 2;
     */
    this.text = "";
    /**
     * @generated from field: uint64 start_time = 3;
     */
    this.startTime = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 end_time = 4;
     */
    this.endTime = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: bool final = 5;
     */
    this.final = false;
    /**
     * @generated from field: string language = 6;
     */
    this.language = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TranscriptionSegment().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TranscriptionSegment().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TranscriptionSegment().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_TranscriptionSegment, a, b);
  }
};
_TranscriptionSegment.runtime = import_protobuf.proto3;
_TranscriptionSegment.typeName = "livekit.TranscriptionSegment";
_TranscriptionSegment.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "text",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "start_time",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 4,
    name: "end_time",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 5,
    name: "final",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 6,
    name: "language",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let TranscriptionSegment = _TranscriptionSegment;
const _ChatMessage = class _ChatMessage extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * uuid
     *
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: int64 timestamp = 2;
     */
    this.timestamp = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: string message = 4;
     */
    this.message = "";
    /**
     * true to remove message
     *
     * @generated from field: bool deleted = 5;
     */
    this.deleted = false;
    /**
     * true if the chat message has been generated by an agent from a participant's audio transcription
     *
     * @generated from field: bool generated = 6;
     */
    this.generated = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ChatMessage().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ChatMessage().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ChatMessage().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ChatMessage, a, b);
  }
};
_ChatMessage.runtime = import_protobuf.proto3;
_ChatMessage.typeName = "livekit.ChatMessage";
_ChatMessage.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "timestamp",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  { no: 3, name: "edit_timestamp", kind: "scalar", T: 3, opt: true },
  {
    no: 4,
    name: "message",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "deleted",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 6,
    name: "generated",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let ChatMessage = _ChatMessage;
const _RpcRequest = class _RpcRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string method = 2;
     */
    this.method = "";
    /**
     * @generated from field: string payload = 3;
     */
    this.payload = "";
    /**
     * @generated from field: uint32 response_timeout_ms = 4;
     */
    this.responseTimeoutMs = 0;
    /**
     * @generated from field: uint32 version = 5;
     */
    this.version = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RpcRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RpcRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RpcRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RpcRequest, a, b);
  }
};
_RpcRequest.runtime = import_protobuf.proto3;
_RpcRequest.typeName = "livekit.RpcRequest";
_RpcRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "method",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "payload",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "response_timeout_ms",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 5,
    name: "version",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let RpcRequest = _RpcRequest;
const _RpcAck = class _RpcAck extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string request_id = 1;
     */
    this.requestId = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RpcAck().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RpcAck().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RpcAck().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RpcAck, a, b);
  }
};
_RpcAck.runtime = import_protobuf.proto3;
_RpcAck.typeName = "livekit.RpcAck";
_RpcAck.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "request_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let RpcAck = _RpcAck;
const _RpcResponse = class _RpcResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string request_id = 1;
     */
    this.requestId = "";
    /**
     * @generated from oneof livekit.RpcResponse.value
     */
    this.value = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RpcResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RpcResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RpcResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RpcResponse, a, b);
  }
};
_RpcResponse.runtime = import_protobuf.proto3;
_RpcResponse.typeName = "livekit.RpcResponse";
_RpcResponse.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "request_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "payload", kind: "scalar", T: 9, oneof: "value" },
  { no: 3, name: "error", kind: "message", T: RpcError, oneof: "value" }
]);
let RpcResponse = _RpcResponse;
const _RpcError = class _RpcError extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: uint32 code = 1;
     */
    this.code = 0;
    /**
     * @generated from field: string message = 2;
     */
    this.message = "";
    /**
     * @generated from field: string data = 3;
     */
    this.data = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RpcError().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RpcError().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RpcError().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RpcError, a, b);
  }
};
_RpcError.runtime = import_protobuf.proto3;
_RpcError.typeName = "livekit.RpcError";
_RpcError.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "code",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 2,
    name: "message",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "data",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let RpcError = _RpcError;
const _ParticipantTracks = class _ParticipantTracks extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * participant ID of participant to whom the tracks belong
     *
     * @generated from field: string participant_sid = 1;
     */
    this.participantSid = "";
    /**
     * @generated from field: repeated string track_sids = 2;
     */
    this.trackSids = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ParticipantTracks().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ParticipantTracks().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ParticipantTracks().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ParticipantTracks, a, b);
  }
};
_ParticipantTracks.runtime = import_protobuf.proto3;
_ParticipantTracks.typeName = "livekit.ParticipantTracks";
_ParticipantTracks.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "participant_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "track_sids", kind: "scalar", T: 9, repeated: true }
]);
let ParticipantTracks = _ParticipantTracks;
const _ServerInfo = class _ServerInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.ServerInfo.Edition edition = 1;
     */
    this.edition = 0 /* Standard */;
    /**
     * @generated from field: string version = 2;
     */
    this.version = "";
    /**
     * @generated from field: int32 protocol = 3;
     */
    this.protocol = 0;
    /**
     * @generated from field: string region = 4;
     */
    this.region = "";
    /**
     * @generated from field: string node_id = 5;
     */
    this.nodeId = "";
    /**
     * additional debugging information. sent only if server is in development mode
     *
     * @generated from field: string debug_info = 6;
     */
    this.debugInfo = "";
    /**
     * @generated from field: int32 agent_protocol = 7;
     */
    this.agentProtocol = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ServerInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ServerInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ServerInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ServerInfo, a, b);
  }
};
_ServerInfo.runtime = import_protobuf.proto3;
_ServerInfo.typeName = "livekit.ServerInfo";
_ServerInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "edition", kind: "enum", T: import_protobuf.proto3.getEnumType(ServerInfo_Edition) },
  {
    no: 2,
    name: "version",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "protocol",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 4,
    name: "region",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "node_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "debug_info",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "agent_protocol",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  }
]);
let ServerInfo = _ServerInfo;
var ServerInfo_Edition = /* @__PURE__ */ ((ServerInfo_Edition2) => {
  ServerInfo_Edition2[ServerInfo_Edition2["Standard"] = 0] = "Standard";
  ServerInfo_Edition2[ServerInfo_Edition2["Cloud"] = 1] = "Cloud";
  return ServerInfo_Edition2;
})(ServerInfo_Edition || {});
import_protobuf.proto3.util.setEnumType(ServerInfo_Edition, "livekit.ServerInfo.Edition", [
  { no: 0, name: "Standard" },
  { no: 1, name: "Cloud" }
]);
const _ClientInfo = class _ClientInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.ClientInfo.SDK sdk = 1;
     */
    this.sdk = 0 /* UNKNOWN */;
    /**
     * @generated from field: string version = 2;
     */
    this.version = "";
    /**
     * @generated from field: int32 protocol = 3;
     */
    this.protocol = 0;
    /**
     * @generated from field: string os = 4;
     */
    this.os = "";
    /**
     * @generated from field: string os_version = 5;
     */
    this.osVersion = "";
    /**
     * @generated from field: string device_model = 6;
     */
    this.deviceModel = "";
    /**
     * @generated from field: string browser = 7;
     */
    this.browser = "";
    /**
     * @generated from field: string browser_version = 8;
     */
    this.browserVersion = "";
    /**
     * @generated from field: string address = 9;
     */
    this.address = "";
    /**
     * wifi, wired, cellular, vpn, empty if not known
     *
     * @generated from field: string network = 10;
     */
    this.network = "";
    /**
     * comma separated list of additional LiveKit SDKs in use of this client, with versions
     * e.g. "components-js:1.2.3,track-processors-js:1.2.3"
     *
     * @generated from field: string other_sdks = 11;
     */
    this.otherSdks = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ClientInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ClientInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ClientInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ClientInfo, a, b);
  }
};
_ClientInfo.runtime = import_protobuf.proto3;
_ClientInfo.typeName = "livekit.ClientInfo";
_ClientInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "sdk", kind: "enum", T: import_protobuf.proto3.getEnumType(ClientInfo_SDK) },
  {
    no: 2,
    name: "version",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "protocol",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 4,
    name: "os",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "os_version",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "device_model",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "browser",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "browser_version",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "address",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 10,
    name: "network",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 11,
    name: "other_sdks",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let ClientInfo = _ClientInfo;
var ClientInfo_SDK = /* @__PURE__ */ ((ClientInfo_SDK2) => {
  ClientInfo_SDK2[ClientInfo_SDK2["UNKNOWN"] = 0] = "UNKNOWN";
  ClientInfo_SDK2[ClientInfo_SDK2["JS"] = 1] = "JS";
  ClientInfo_SDK2[ClientInfo_SDK2["SWIFT"] = 2] = "SWIFT";
  ClientInfo_SDK2[ClientInfo_SDK2["ANDROID"] = 3] = "ANDROID";
  ClientInfo_SDK2[ClientInfo_SDK2["FLUTTER"] = 4] = "FLUTTER";
  ClientInfo_SDK2[ClientInfo_SDK2["GO"] = 5] = "GO";
  ClientInfo_SDK2[ClientInfo_SDK2["UNITY"] = 6] = "UNITY";
  ClientInfo_SDK2[ClientInfo_SDK2["REACT_NATIVE"] = 7] = "REACT_NATIVE";
  ClientInfo_SDK2[ClientInfo_SDK2["RUST"] = 8] = "RUST";
  ClientInfo_SDK2[ClientInfo_SDK2["PYTHON"] = 9] = "PYTHON";
  ClientInfo_SDK2[ClientInfo_SDK2["CPP"] = 10] = "CPP";
  ClientInfo_SDK2[ClientInfo_SDK2["UNITY_WEB"] = 11] = "UNITY_WEB";
  ClientInfo_SDK2[ClientInfo_SDK2["NODE"] = 12] = "NODE";
  return ClientInfo_SDK2;
})(ClientInfo_SDK || {});
import_protobuf.proto3.util.setEnumType(ClientInfo_SDK, "livekit.ClientInfo.SDK", [
  { no: 0, name: "UNKNOWN" },
  { no: 1, name: "JS" },
  { no: 2, name: "SWIFT" },
  { no: 3, name: "ANDROID" },
  { no: 4, name: "FLUTTER" },
  { no: 5, name: "GO" },
  { no: 6, name: "UNITY" },
  { no: 7, name: "REACT_NATIVE" },
  { no: 8, name: "RUST" },
  { no: 9, name: "PYTHON" },
  { no: 10, name: "CPP" },
  { no: 11, name: "UNITY_WEB" },
  { no: 12, name: "NODE" }
]);
const _ClientConfiguration = class _ClientConfiguration extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.ClientConfigSetting resume_connection = 3;
     */
    this.resumeConnection = 0 /* UNSET */;
    /**
     * @generated from field: livekit.ClientConfigSetting force_relay = 5;
     */
    this.forceRelay = 0 /* UNSET */;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ClientConfiguration().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ClientConfiguration().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ClientConfiguration().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ClientConfiguration, a, b);
  }
};
_ClientConfiguration.runtime = import_protobuf.proto3;
_ClientConfiguration.typeName = "livekit.ClientConfiguration";
_ClientConfiguration.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "video", kind: "message", T: VideoConfiguration },
  { no: 2, name: "screen", kind: "message", T: VideoConfiguration },
  { no: 3, name: "resume_connection", kind: "enum", T: import_protobuf.proto3.getEnumType(ClientConfigSetting) },
  { no: 4, name: "disabled_codecs", kind: "message", T: DisabledCodecs },
  { no: 5, name: "force_relay", kind: "enum", T: import_protobuf.proto3.getEnumType(ClientConfigSetting) }
]);
let ClientConfiguration = _ClientConfiguration;
const _VideoConfiguration = class _VideoConfiguration extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.ClientConfigSetting hardware_encoder = 1;
     */
    this.hardwareEncoder = 0 /* UNSET */;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _VideoConfiguration().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _VideoConfiguration().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _VideoConfiguration().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_VideoConfiguration, a, b);
  }
};
_VideoConfiguration.runtime = import_protobuf.proto3;
_VideoConfiguration.typeName = "livekit.VideoConfiguration";
_VideoConfiguration.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "hardware_encoder", kind: "enum", T: import_protobuf.proto3.getEnumType(ClientConfigSetting) }
]);
let VideoConfiguration = _VideoConfiguration;
const _DisabledCodecs = class _DisabledCodecs extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * disabled for both publish and subscribe
     *
     * @generated from field: repeated livekit.Codec codecs = 1;
     */
    this.codecs = [];
    /**
     * only disable for publish
     *
     * @generated from field: repeated livekit.Codec publish = 2;
     */
    this.publish = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DisabledCodecs().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DisabledCodecs().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DisabledCodecs().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DisabledCodecs, a, b);
  }
};
_DisabledCodecs.runtime = import_protobuf.proto3;
_DisabledCodecs.typeName = "livekit.DisabledCodecs";
_DisabledCodecs.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "codecs", kind: "message", T: Codec, repeated: true },
  { no: 2, name: "publish", kind: "message", T: Codec, repeated: true }
]);
let DisabledCodecs = _DisabledCodecs;
const _RTPDrift = class _RTPDrift extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: double duration = 3;
     */
    this.duration = 0;
    /**
     * @generated from field: uint64 start_timestamp = 4;
     */
    this.startTimestamp = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 end_timestamp = 5;
     */
    this.endTimestamp = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 rtp_clock_ticks = 6;
     */
    this.rtpClockTicks = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 drift_samples = 7;
     */
    this.driftSamples = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: double drift_ms = 8;
     */
    this.driftMs = 0;
    /**
     * @generated from field: double clock_rate = 9;
     */
    this.clockRate = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RTPDrift().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RTPDrift().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RTPDrift().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RTPDrift, a, b);
  }
};
_RTPDrift.runtime = import_protobuf.proto3;
_RTPDrift.typeName = "livekit.RTPDrift";
_RTPDrift.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "start_time", kind: "message", T: import_protobuf.Timestamp },
  { no: 2, name: "end_time", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 3,
    name: "duration",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 4,
    name: "start_timestamp",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 5,
    name: "end_timestamp",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 6,
    name: "rtp_clock_ticks",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 7,
    name: "drift_samples",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 8,
    name: "drift_ms",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 9,
    name: "clock_rate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  }
]);
let RTPDrift = _RTPDrift;
const _RTPStats = class _RTPStats extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: double duration = 3;
     */
    this.duration = 0;
    /**
     * @generated from field: uint32 packets = 4;
     */
    this.packets = 0;
    /**
     * @generated from field: double packet_rate = 5;
     */
    this.packetRate = 0;
    /**
     * @generated from field: uint64 bytes = 6;
     */
    this.bytes = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 header_bytes = 39;
     */
    this.headerBytes = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: double bitrate = 7;
     */
    this.bitrate = 0;
    /**
     * @generated from field: uint32 packets_lost = 8;
     */
    this.packetsLost = 0;
    /**
     * @generated from field: double packet_loss_rate = 9;
     */
    this.packetLossRate = 0;
    /**
     * @generated from field: float packet_loss_percentage = 10;
     */
    this.packetLossPercentage = 0;
    /**
     * @generated from field: uint32 packets_duplicate = 11;
     */
    this.packetsDuplicate = 0;
    /**
     * @generated from field: double packet_duplicate_rate = 12;
     */
    this.packetDuplicateRate = 0;
    /**
     * @generated from field: uint64 bytes_duplicate = 13;
     */
    this.bytesDuplicate = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 header_bytes_duplicate = 40;
     */
    this.headerBytesDuplicate = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: double bitrate_duplicate = 14;
     */
    this.bitrateDuplicate = 0;
    /**
     * @generated from field: uint32 packets_padding = 15;
     */
    this.packetsPadding = 0;
    /**
     * @generated from field: double packet_padding_rate = 16;
     */
    this.packetPaddingRate = 0;
    /**
     * @generated from field: uint64 bytes_padding = 17;
     */
    this.bytesPadding = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 header_bytes_padding = 41;
     */
    this.headerBytesPadding = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: double bitrate_padding = 18;
     */
    this.bitratePadding = 0;
    /**
     * @generated from field: uint32 packets_out_of_order = 19;
     */
    this.packetsOutOfOrder = 0;
    /**
     * @generated from field: uint32 frames = 20;
     */
    this.frames = 0;
    /**
     * @generated from field: double frame_rate = 21;
     */
    this.frameRate = 0;
    /**
     * @generated from field: double jitter_current = 22;
     */
    this.jitterCurrent = 0;
    /**
     * @generated from field: double jitter_max = 23;
     */
    this.jitterMax = 0;
    /**
     * @generated from field: map<int32, uint32> gap_histogram = 24;
     */
    this.gapHistogram = {};
    /**
     * @generated from field: uint32 nacks = 25;
     */
    this.nacks = 0;
    /**
     * @generated from field: uint32 nack_acks = 37;
     */
    this.nackAcks = 0;
    /**
     * @generated from field: uint32 nack_misses = 26;
     */
    this.nackMisses = 0;
    /**
     * @generated from field: uint32 nack_repeated = 38;
     */
    this.nackRepeated = 0;
    /**
     * @generated from field: uint32 plis = 27;
     */
    this.plis = 0;
    /**
     * @generated from field: uint32 firs = 29;
     */
    this.firs = 0;
    /**
     * @generated from field: uint32 rtt_current = 31;
     */
    this.rttCurrent = 0;
    /**
     * @generated from field: uint32 rtt_max = 32;
     */
    this.rttMax = 0;
    /**
     * @generated from field: uint32 key_frames = 33;
     */
    this.keyFrames = 0;
    /**
     * @generated from field: uint32 layer_lock_plis = 35;
     */
    this.layerLockPlis = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RTPStats().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RTPStats().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RTPStats().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RTPStats, a, b);
  }
};
_RTPStats.runtime = import_protobuf.proto3;
_RTPStats.typeName = "livekit.RTPStats";
_RTPStats.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "start_time", kind: "message", T: import_protobuf.Timestamp },
  { no: 2, name: "end_time", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 3,
    name: "duration",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 4,
    name: "packets",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 5,
    name: "packet_rate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 6,
    name: "bytes",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 39,
    name: "header_bytes",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 7,
    name: "bitrate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
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
    name: "packet_loss_rate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 10,
    name: "packet_loss_percentage",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 11,
    name: "packets_duplicate",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 12,
    name: "packet_duplicate_rate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 13,
    name: "bytes_duplicate",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 40,
    name: "header_bytes_duplicate",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 14,
    name: "bitrate_duplicate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 15,
    name: "packets_padding",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 16,
    name: "packet_padding_rate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 17,
    name: "bytes_padding",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 41,
    name: "header_bytes_padding",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 18,
    name: "bitrate_padding",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 19,
    name: "packets_out_of_order",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 20,
    name: "frames",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 21,
    name: "frame_rate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 22,
    name: "jitter_current",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  {
    no: 23,
    name: "jitter_max",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  { no: 24, name: "gap_histogram", kind: "map", K: 5, V: {
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  } },
  {
    no: 25,
    name: "nacks",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 37,
    name: "nack_acks",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 26,
    name: "nack_misses",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 38,
    name: "nack_repeated",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 27,
    name: "plis",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 28, name: "last_pli", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 29,
    name: "firs",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 30, name: "last_fir", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 31,
    name: "rtt_current",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 32,
    name: "rtt_max",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 33,
    name: "key_frames",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 34, name: "last_key_frame", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 35,
    name: "layer_lock_plis",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 36, name: "last_layer_lock_pli", kind: "message", T: import_protobuf.Timestamp },
  { no: 44, name: "packet_drift", kind: "message", T: RTPDrift },
  { no: 45, name: "ntp_report_drift", kind: "message", T: RTPDrift },
  { no: 46, name: "rebased_report_drift", kind: "message", T: RTPDrift },
  { no: 47, name: "received_report_drift", kind: "message", T: RTPDrift }
]);
let RTPStats = _RTPStats;
const _RTCPSenderReportState = class _RTCPSenderReportState extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: uint32 rtp_timestamp = 1;
     */
    this.rtpTimestamp = 0;
    /**
     * @generated from field: uint64 rtp_timestamp_ext = 2;
     */
    this.rtpTimestampExt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 ntp_timestamp = 3;
     */
    this.ntpTimestamp = import_protobuf.protoInt64.zero;
    /**
     * time at which this happened
     *
     * @generated from field: int64 at = 4;
     */
    this.at = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 at_adjusted = 5;
     */
    this.atAdjusted = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint32 packets = 6;
     */
    this.packets = 0;
    /**
     * @generated from field: uint64 octets = 7;
     */
    this.octets = import_protobuf.protoInt64.zero;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RTCPSenderReportState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RTCPSenderReportState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RTCPSenderReportState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RTCPSenderReportState, a, b);
  }
};
_RTCPSenderReportState.runtime = import_protobuf.proto3;
_RTCPSenderReportState.typeName = "livekit.RTCPSenderReportState";
_RTCPSenderReportState.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "rtp_timestamp",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 2,
    name: "rtp_timestamp_ext",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 3,
    name: "ntp_timestamp",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 4,
    name: "at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 5,
    name: "at_adjusted",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 6,
    name: "packets",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 7,
    name: "octets",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  }
]);
let RTCPSenderReportState = _RTCPSenderReportState;
const _RTPForwarderState = class _RTPForwarderState extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool started = 1;
     */
    this.started = false;
    /**
     * @generated from field: int32 reference_layer_spatial = 2;
     */
    this.referenceLayerSpatial = 0;
    /**
     * @generated from field: int64 pre_start_time = 3;
     */
    this.preStartTime = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 ext_first_timestamp = 4;
     */
    this.extFirstTimestamp = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 dummy_start_timestamp_offset = 5;
     */
    this.dummyStartTimestampOffset = import_protobuf.protoInt64.zero;
    /**
     * @generated from oneof livekit.RTPForwarderState.codec_munger
     */
    this.codecMunger = { case: void 0 };
    /**
     * @generated from field: repeated livekit.RTCPSenderReportState sender_report_state = 8;
     */
    this.senderReportState = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RTPForwarderState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RTPForwarderState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RTPForwarderState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RTPForwarderState, a, b);
  }
};
_RTPForwarderState.runtime = import_protobuf.proto3;
_RTPForwarderState.typeName = "livekit.RTPForwarderState";
_RTPForwarderState.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "started",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 2,
    name: "reference_layer_spatial",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 3,
    name: "pre_start_time",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 4,
    name: "ext_first_timestamp",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 5,
    name: "dummy_start_timestamp_offset",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  { no: 6, name: "rtp_munger", kind: "message", T: RTPMungerState },
  { no: 7, name: "vp8_munger", kind: "message", T: VP8MungerState, oneof: "codec_munger" },
  { no: 8, name: "sender_report_state", kind: "message", T: RTCPSenderReportState, repeated: true }
]);
let RTPForwarderState = _RTPForwarderState;
const _RTPMungerState = class _RTPMungerState extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: uint64 ext_last_sequence_number = 1;
     */
    this.extLastSequenceNumber = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 ext_second_last_sequence_number = 2;
     */
    this.extSecondLastSequenceNumber = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 ext_last_timestamp = 3;
     */
    this.extLastTimestamp = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: uint64 ext_second_last_timestamp = 4;
     */
    this.extSecondLastTimestamp = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: bool last_marker = 5;
     */
    this.lastMarker = false;
    /**
     * @generated from field: bool second_last_marker = 6;
     */
    this.secondLastMarker = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RTPMungerState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RTPMungerState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RTPMungerState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RTPMungerState, a, b);
  }
};
_RTPMungerState.runtime = import_protobuf.proto3;
_RTPMungerState.typeName = "livekit.RTPMungerState";
_RTPMungerState.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "ext_last_sequence_number",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 2,
    name: "ext_second_last_sequence_number",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 3,
    name: "ext_last_timestamp",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 4,
    name: "ext_second_last_timestamp",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 5,
    name: "last_marker",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 6,
    name: "second_last_marker",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let RTPMungerState = _RTPMungerState;
const _VP8MungerState = class _VP8MungerState extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: int32 ext_last_picture_id = 1;
     */
    this.extLastPictureId = 0;
    /**
     * @generated from field: bool picture_id_used = 2;
     */
    this.pictureIdUsed = false;
    /**
     * @generated from field: uint32 last_tl0_pic_idx = 3;
     */
    this.lastTl0PicIdx = 0;
    /**
     * @generated from field: bool tl0_pic_idx_used = 4;
     */
    this.tl0PicIdxUsed = false;
    /**
     * @generated from field: bool tid_used = 5;
     */
    this.tidUsed = false;
    /**
     * @generated from field: uint32 last_key_idx = 6;
     */
    this.lastKeyIdx = 0;
    /**
     * @generated from field: bool key_idx_used = 7;
     */
    this.keyIdxUsed = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _VP8MungerState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _VP8MungerState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _VP8MungerState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_VP8MungerState, a, b);
  }
};
_VP8MungerState.runtime = import_protobuf.proto3;
_VP8MungerState.typeName = "livekit.VP8MungerState";
_VP8MungerState.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "ext_last_picture_id",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 2,
    name: "picture_id_used",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 3,
    name: "last_tl0_pic_idx",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "tl0_pic_idx_used",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 5,
    name: "tid_used",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 6,
    name: "last_key_idx",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 7,
    name: "key_idx_used",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let VP8MungerState = _VP8MungerState;
const _TimedVersion = class _TimedVersion extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: int64 unix_micro = 1;
     */
    this.unixMicro = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int32 ticks = 2;
     */
    this.ticks = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TimedVersion().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TimedVersion().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TimedVersion().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_TimedVersion, a, b);
  }
};
_TimedVersion.runtime = import_protobuf.proto3;
_TimedVersion.typeName = "livekit.TimedVersion";
_TimedVersion.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "unix_micro",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 2,
    name: "ticks",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  }
]);
let TimedVersion = _TimedVersion;
const _DataStream = class _DataStream extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DataStream().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DataStream().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DataStream().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DataStream, a, b);
  }
};
_DataStream.runtime = import_protobuf.proto3;
_DataStream.typeName = "livekit.DataStream";
_DataStream.fields = import_protobuf.proto3.util.newFieldList(() => []);
let DataStream = _DataStream;
var DataStream_OperationType = /* @__PURE__ */ ((DataStream_OperationType2) => {
  DataStream_OperationType2[DataStream_OperationType2["CREATE"] = 0] = "CREATE";
  DataStream_OperationType2[DataStream_OperationType2["UPDATE"] = 1] = "UPDATE";
  DataStream_OperationType2[DataStream_OperationType2["DELETE"] = 2] = "DELETE";
  DataStream_OperationType2[DataStream_OperationType2["REACTION"] = 3] = "REACTION";
  return DataStream_OperationType2;
})(DataStream_OperationType || {});
import_protobuf.proto3.util.setEnumType(DataStream_OperationType, "livekit.DataStream.OperationType", [
  { no: 0, name: "CREATE" },
  { no: 1, name: "UPDATE" },
  { no: 2, name: "DELETE" },
  { no: 3, name: "REACTION" }
]);
const _DataStream_TextHeader = class _DataStream_TextHeader extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.DataStream.OperationType operation_type = 1;
     */
    this.operationType = 0 /* CREATE */;
    /**
     * Optional: Version for updates/edits
     *
     * @generated from field: int32 version = 2;
     */
    this.version = 0;
    /**
     * Optional: Reply to specific message
     *
     * @generated from field: string reply_to_stream_id = 3;
     */
    this.replyToStreamId = "";
    /**
     * file attachments for text streams
     *
     * @generated from field: repeated string attached_stream_ids = 4;
     */
    this.attachedStreamIds = [];
    /**
     * true if the text has been generated by an agent from a participant's audio transcription
     *
     * @generated from field: bool generated = 5;
     */
    this.generated = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DataStream_TextHeader().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DataStream_TextHeader().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DataStream_TextHeader().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DataStream_TextHeader, a, b);
  }
};
_DataStream_TextHeader.runtime = import_protobuf.proto3;
_DataStream_TextHeader.typeName = "livekit.DataStream.TextHeader";
_DataStream_TextHeader.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "operation_type", kind: "enum", T: import_protobuf.proto3.getEnumType(DataStream_OperationType) },
  {
    no: 2,
    name: "version",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 3,
    name: "reply_to_stream_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "attached_stream_ids", kind: "scalar", T: 9, repeated: true },
  {
    no: 5,
    name: "generated",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let DataStream_TextHeader = _DataStream_TextHeader;
const _DataStream_FileHeader = class _DataStream_FileHeader extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * name of the file
     *
     * @generated from field: string file_name = 1;
     */
    this.fileName = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DataStream_FileHeader().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DataStream_FileHeader().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DataStream_FileHeader().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DataStream_FileHeader, a, b);
  }
};
_DataStream_FileHeader.runtime = import_protobuf.proto3;
_DataStream_FileHeader.typeName = "livekit.DataStream.FileHeader";
_DataStream_FileHeader.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "file_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let DataStream_FileHeader = _DataStream_FileHeader;
const _DataStream_Header = class _DataStream_Header extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * unique identifier for this data stream
     *
     * @generated from field: string stream_id = 1;
     */
    this.streamId = "";
    /**
     * using int64 for Unix timestamp
     *
     * @generated from field: int64 timestamp = 2;
     */
    this.timestamp = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: string topic = 3;
     */
    this.topic = "";
    /**
     * @generated from field: string mime_type = 4;
     */
    this.mimeType = "";
    /**
     * defaults to NONE
     *
     * @generated from field: livekit.Encryption.Type encryption_type = 7;
     */
    this.encryptionType = 0 /* NONE */;
    /**
     * user defined extensions map that can carry additional info
     *
     * @generated from field: map<string, string> extensions = 8;
     */
    this.extensions = {};
    /**
     * oneof to choose between specific header types
     *
     * @generated from oneof livekit.DataStream.Header.content_header
     */
    this.contentHeader = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DataStream_Header().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DataStream_Header().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DataStream_Header().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DataStream_Header, a, b);
  }
};
_DataStream_Header.runtime = import_protobuf.proto3;
_DataStream_Header.typeName = "livekit.DataStream.Header";
_DataStream_Header.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "stream_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "timestamp",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 3,
    name: "topic",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "mime_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "total_length", kind: "scalar", T: 4, opt: true },
  { no: 6, name: "total_chunks", kind: "scalar", T: 4, opt: true },
  { no: 7, name: "encryption_type", kind: "enum", T: import_protobuf.proto3.getEnumType(Encryption_Type) },
  { no: 8, name: "extensions", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  { no: 9, name: "text_header", kind: "message", T: DataStream_TextHeader, oneof: "content_header" },
  { no: 10, name: "file_header", kind: "message", T: DataStream_FileHeader, oneof: "content_header" }
]);
let DataStream_Header = _DataStream_Header;
const _DataStream_Chunk = class _DataStream_Chunk extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * unique identifier for this data stream to map it to the correct header
     *
     * @generated from field: string stream_id = 1;
     */
    this.streamId = "";
    /**
     * @generated from field: uint64 chunk_index = 2;
     */
    this.chunkIndex = import_protobuf.protoInt64.zero;
    /**
     * content as binary (bytes)
     *
     * @generated from field: bytes content = 3;
     */
    this.content = new Uint8Array(0);
    /**
     * true only if this is the last chunk of this stream - can also be sent with empty content
     *
     * @generated from field: bool complete = 4;
     */
    this.complete = false;
    /**
     * a version indicating that this chunk_index has been retroactively modified and the original one needs to be replaced
     *
     * @generated from field: int32 version = 5;
     */
    this.version = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DataStream_Chunk().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DataStream_Chunk().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DataStream_Chunk().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DataStream_Chunk, a, b);
  }
};
_DataStream_Chunk.runtime = import_protobuf.proto3;
_DataStream_Chunk.typeName = "livekit.DataStream.Chunk";
_DataStream_Chunk.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "stream_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "chunk_index",
    kind: "scalar",
    T: 4
    /* ScalarType.UINT64 */
  },
  {
    no: 3,
    name: "content",
    kind: "scalar",
    T: 12
    /* ScalarType.BYTES */
  },
  {
    no: 4,
    name: "complete",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 5,
    name: "version",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  { no: 6, name: "iv", kind: "scalar", T: 12, opt: true }
]);
let DataStream_Chunk = _DataStream_Chunk;
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  ActiveSpeakerUpdate,
  AudioCodec,
  AudioTrackFeature,
  ChatMessage,
  ClientConfigSetting,
  ClientConfiguration,
  ClientInfo,
  ClientInfo_SDK,
  Codec,
  ConnectionQuality,
  DataPacket,
  DataPacket_Kind,
  DataStream,
  DataStream_Chunk,
  DataStream_FileHeader,
  DataStream_Header,
  DataStream_OperationType,
  DataStream_TextHeader,
  DisabledCodecs,
  DisconnectReason,
  Encryption,
  Encryption_Type,
  ImageCodec,
  ParticipantInfo,
  ParticipantInfo_Kind,
  ParticipantInfo_State,
  ParticipantPermission,
  ParticipantTracks,
  PlayoutDelay,
  RTCPSenderReportState,
  RTPDrift,
  RTPForwarderState,
  RTPMungerState,
  RTPStats,
  ReconnectReason,
  Room,
  RpcAck,
  RpcError,
  RpcRequest,
  RpcResponse,
  ServerInfo,
  ServerInfo_Edition,
  SimulcastCodecInfo,
  SipDTMF,
  SpeakerInfo,
  SubscriptionError,
  TimedVersion,
  TrackInfo,
  TrackSource,
  TrackType,
  Transcription,
  TranscriptionSegment,
  UserPacket,
  VP8MungerState,
  VideoCodec,
  VideoConfiguration,
  VideoLayer,
  VideoQuality
});
//# sourceMappingURL=livekit_models_pb.cjs.map