import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { AudioTrackFeature, ClientConfiguration, Codec, ConnectionQuality, DisconnectReason, Encryption_Type, ParticipantInfo, ParticipantTracks, Room, ServerInfo, SpeakerInfo, SubscriptionError, TrackInfo, TrackSource, TrackType, VideoLayer, VideoQuality } from "./livekit_models_pb.js";
var SignalTarget = /* @__PURE__ */ ((SignalTarget2) => {
  SignalTarget2[SignalTarget2["PUBLISHER"] = 0] = "PUBLISHER";
  SignalTarget2[SignalTarget2["SUBSCRIBER"] = 1] = "SUBSCRIBER";
  return SignalTarget2;
})(SignalTarget || {});
proto3.util.setEnumType(SignalTarget, "livekit.SignalTarget", [
  { no: 0, name: "PUBLISHER" },
  { no: 1, name: "SUBSCRIBER" }
]);
var StreamState = /* @__PURE__ */ ((StreamState2) => {
  StreamState2[StreamState2["ACTIVE"] = 0] = "ACTIVE";
  StreamState2[StreamState2["PAUSED"] = 1] = "PAUSED";
  return StreamState2;
})(StreamState || {});
proto3.util.setEnumType(StreamState, "livekit.StreamState", [
  { no: 0, name: "ACTIVE" },
  { no: 1, name: "PAUSED" }
]);
var CandidateProtocol = /* @__PURE__ */ ((CandidateProtocol2) => {
  CandidateProtocol2[CandidateProtocol2["UDP"] = 0] = "UDP";
  CandidateProtocol2[CandidateProtocol2["TCP"] = 1] = "TCP";
  CandidateProtocol2[CandidateProtocol2["TLS"] = 2] = "TLS";
  return CandidateProtocol2;
})(CandidateProtocol || {});
proto3.util.setEnumType(CandidateProtocol, "livekit.CandidateProtocol", [
  { no: 0, name: "UDP" },
  { no: 1, name: "TCP" },
  { no: 2, name: "TLS" }
]);
const _SignalRequest = class _SignalRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from oneof livekit.SignalRequest.message
     */
    this.message = { case: void 0 };
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SignalRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SignalRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SignalRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SignalRequest, a, b);
  }
};
_SignalRequest.runtime = proto3;
_SignalRequest.typeName = "livekit.SignalRequest";
_SignalRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "offer", kind: "message", T: SessionDescription, oneof: "message" },
  { no: 2, name: "answer", kind: "message", T: SessionDescription, oneof: "message" },
  { no: 3, name: "trickle", kind: "message", T: TrickleRequest, oneof: "message" },
  { no: 4, name: "add_track", kind: "message", T: AddTrackRequest, oneof: "message" },
  { no: 5, name: "mute", kind: "message", T: MuteTrackRequest, oneof: "message" },
  { no: 6, name: "subscription", kind: "message", T: UpdateSubscription, oneof: "message" },
  { no: 7, name: "track_setting", kind: "message", T: UpdateTrackSettings, oneof: "message" },
  { no: 8, name: "leave", kind: "message", T: LeaveRequest, oneof: "message" },
  { no: 10, name: "update_layers", kind: "message", T: UpdateVideoLayers, oneof: "message" },
  { no: 11, name: "subscription_permission", kind: "message", T: SubscriptionPermission, oneof: "message" },
  { no: 12, name: "sync_state", kind: "message", T: SyncState, oneof: "message" },
  { no: 13, name: "simulate", kind: "message", T: SimulateScenario, oneof: "message" },
  { no: 14, name: "ping", kind: "scalar", T: 3, oneof: "message" },
  { no: 15, name: "update_metadata", kind: "message", T: UpdateParticipantMetadata, oneof: "message" },
  { no: 16, name: "ping_req", kind: "message", T: Ping, oneof: "message" },
  { no: 17, name: "update_audio_track", kind: "message", T: UpdateLocalAudioTrack, oneof: "message" },
  { no: 18, name: "update_video_track", kind: "message", T: UpdateLocalVideoTrack, oneof: "message" }
]);
let SignalRequest = _SignalRequest;
const _SignalResponse = class _SignalResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from oneof livekit.SignalResponse.message
     */
    this.message = { case: void 0 };
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SignalResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SignalResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SignalResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SignalResponse, a, b);
  }
};
_SignalResponse.runtime = proto3;
_SignalResponse.typeName = "livekit.SignalResponse";
_SignalResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "join", kind: "message", T: JoinResponse, oneof: "message" },
  { no: 2, name: "answer", kind: "message", T: SessionDescription, oneof: "message" },
  { no: 3, name: "offer", kind: "message", T: SessionDescription, oneof: "message" },
  { no: 4, name: "trickle", kind: "message", T: TrickleRequest, oneof: "message" },
  { no: 5, name: "update", kind: "message", T: ParticipantUpdate, oneof: "message" },
  { no: 6, name: "track_published", kind: "message", T: TrackPublishedResponse, oneof: "message" },
  { no: 8, name: "leave", kind: "message", T: LeaveRequest, oneof: "message" },
  { no: 9, name: "mute", kind: "message", T: MuteTrackRequest, oneof: "message" },
  { no: 10, name: "speakers_changed", kind: "message", T: SpeakersChanged, oneof: "message" },
  { no: 11, name: "room_update", kind: "message", T: RoomUpdate, oneof: "message" },
  { no: 12, name: "connection_quality", kind: "message", T: ConnectionQualityUpdate, oneof: "message" },
  { no: 13, name: "stream_state_update", kind: "message", T: StreamStateUpdate, oneof: "message" },
  { no: 14, name: "subscribed_quality_update", kind: "message", T: SubscribedQualityUpdate, oneof: "message" },
  { no: 15, name: "subscription_permission_update", kind: "message", T: SubscriptionPermissionUpdate, oneof: "message" },
  { no: 16, name: "refresh_token", kind: "scalar", T: 9, oneof: "message" },
  { no: 17, name: "track_unpublished", kind: "message", T: TrackUnpublishedResponse, oneof: "message" },
  { no: 18, name: "pong", kind: "scalar", T: 3, oneof: "message" },
  { no: 19, name: "reconnect", kind: "message", T: ReconnectResponse, oneof: "message" },
  { no: 20, name: "pong_resp", kind: "message", T: Pong, oneof: "message" },
  { no: 21, name: "subscription_response", kind: "message", T: SubscriptionResponse, oneof: "message" },
  { no: 22, name: "request_response", kind: "message", T: RequestResponse, oneof: "message" },
  { no: 23, name: "track_subscribed", kind: "message", T: TrackSubscribed, oneof: "message" }
]);
let SignalResponse = _SignalResponse;
const _SimulcastCodec = class _SimulcastCodec extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string codec = 1;
     */
    this.codec = "";
    /**
     * @generated from field: string cid = 2;
     */
    this.cid = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SimulcastCodec().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SimulcastCodec().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SimulcastCodec().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SimulcastCodec, a, b);
  }
};
_SimulcastCodec.runtime = proto3;
_SimulcastCodec.typeName = "livekit.SimulcastCodec";
_SimulcastCodec.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "codec",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "cid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let SimulcastCodec = _SimulcastCodec;
const _AddTrackRequest = class _AddTrackRequest extends Message {
  constructor(data) {
    super();
    /**
     * client ID of track, to match it when RTC track is received
     *
     * @generated from field: string cid = 1;
     */
    this.cid = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * @generated from field: livekit.TrackType type = 3;
     */
    this.type = TrackType.AUDIO;
    /**
     * to be deprecated in favor of layers
     *
     * @generated from field: uint32 width = 4;
     */
    this.width = 0;
    /**
     * @generated from field: uint32 height = 5;
     */
    this.height = 0;
    /**
     * true to add track and initialize to muted
     *
     * @generated from field: bool muted = 6;
     */
    this.muted = false;
    /**
     * true if DTX (Discontinuous Transmission) is disabled for audio
     *
     * @generated from field: bool disable_dtx = 7;
     */
    this.disableDtx = false;
    /**
     * @generated from field: livekit.TrackSource source = 8;
     */
    this.source = TrackSource.UNKNOWN;
    /**
     * @generated from field: repeated livekit.VideoLayer layers = 9;
     */
    this.layers = [];
    /**
     * @generated from field: repeated livekit.SimulcastCodec simulcast_codecs = 10;
     */
    this.simulcastCodecs = [];
    /**
     * server ID of track, publish new codec to exist track
     *
     * @generated from field: string sid = 11;
     */
    this.sid = "";
    /**
     * @generated from field: bool stereo = 12;
     */
    this.stereo = false;
    /**
     * true if RED (Redundant Encoding) is disabled for audio
     *
     * @generated from field: bool disable_red = 13;
     */
    this.disableRed = false;
    /**
     * @generated from field: livekit.Encryption.Type encryption = 14;
     */
    this.encryption = Encryption_Type.NONE;
    /**
     * which stream the track belongs to, used to group tracks together.
     * if not specified, server will infer it from track source to bundle camera/microphone, screenshare/audio together
     *
     * @generated from field: string stream = 15;
     */
    this.stream = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AddTrackRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AddTrackRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AddTrackRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_AddTrackRequest, a, b);
  }
};
_AddTrackRequest.runtime = proto3;
_AddTrackRequest.typeName = "livekit.AddTrackRequest";
_AddTrackRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "cid",
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
  { no: 3, name: "type", kind: "enum", T: proto3.getEnumType(TrackType) },
  {
    no: 4,
    name: "width",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 5,
    name: "height",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 6,
    name: "muted",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 7,
    name: "disable_dtx",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 8, name: "source", kind: "enum", T: proto3.getEnumType(TrackSource) },
  { no: 9, name: "layers", kind: "message", T: VideoLayer, repeated: true },
  { no: 10, name: "simulcast_codecs", kind: "message", T: SimulcastCodec, repeated: true },
  {
    no: 11,
    name: "sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 12,
    name: "stereo",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 13,
    name: "disable_red",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 14, name: "encryption", kind: "enum", T: proto3.getEnumType(Encryption_Type) },
  {
    no: 15,
    name: "stream",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let AddTrackRequest = _AddTrackRequest;
const _TrickleRequest = class _TrickleRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string candidateInit = 1;
     */
    this.candidateInit = "";
    /**
     * @generated from field: livekit.SignalTarget target = 2;
     */
    this.target = 0 /* PUBLISHER */;
    /**
     * @generated from field: bool final = 3;
     */
    this.final = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TrickleRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TrickleRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TrickleRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_TrickleRequest, a, b);
  }
};
_TrickleRequest.runtime = proto3;
_TrickleRequest.typeName = "livekit.TrickleRequest";
_TrickleRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "candidateInit",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "target", kind: "enum", T: proto3.getEnumType(SignalTarget) },
  {
    no: 3,
    name: "final",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let TrickleRequest = _TrickleRequest;
const _MuteTrackRequest = class _MuteTrackRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sid = 1;
     */
    this.sid = "";
    /**
     * @generated from field: bool muted = 2;
     */
    this.muted = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _MuteTrackRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _MuteTrackRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _MuteTrackRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_MuteTrackRequest, a, b);
  }
};
_MuteTrackRequest.runtime = proto3;
_MuteTrackRequest.typeName = "livekit.MuteTrackRequest";
_MuteTrackRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "muted",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let MuteTrackRequest = _MuteTrackRequest;
const _JoinResponse = class _JoinResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.ParticipantInfo other_participants = 3;
     */
    this.otherParticipants = [];
    /**
     * deprecated. use server_info.version instead.
     *
     * @generated from field: string server_version = 4;
     */
    this.serverVersion = "";
    /**
     * @generated from field: repeated livekit.ICEServer ice_servers = 5;
     */
    this.iceServers = [];
    /**
     * use subscriber as the primary PeerConnection
     *
     * @generated from field: bool subscriber_primary = 6;
     */
    this.subscriberPrimary = false;
    /**
     * when the current server isn't available, return alternate url to retry connection
     * when this is set, the other fields will be largely empty
     *
     * @generated from field: string alternative_url = 7;
     */
    this.alternativeUrl = "";
    /**
     * deprecated. use server_info.region instead.
     *
     * @generated from field: string server_region = 9;
     */
    this.serverRegion = "";
    /**
     * @generated from field: int32 ping_timeout = 10;
     */
    this.pingTimeout = 0;
    /**
     * @generated from field: int32 ping_interval = 11;
     */
    this.pingInterval = 0;
    /**
     * Server-Injected-Frame byte trailer, used to identify unencrypted frames when e2ee is enabled
     *
     * @generated from field: bytes sif_trailer = 13;
     */
    this.sifTrailer = new Uint8Array(0);
    /**
     * @generated from field: repeated livekit.Codec enabled_publish_codecs = 14;
     */
    this.enabledPublishCodecs = [];
    /**
     * when set, client should attempt to establish publish peer connection when joining room to speed up publishing
     *
     * @generated from field: bool fast_publish = 15;
     */
    this.fastPublish = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _JoinResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _JoinResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _JoinResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_JoinResponse, a, b);
  }
};
_JoinResponse.runtime = proto3;
_JoinResponse.typeName = "livekit.JoinResponse";
_JoinResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "room", kind: "message", T: Room },
  { no: 2, name: "participant", kind: "message", T: ParticipantInfo },
  { no: 3, name: "other_participants", kind: "message", T: ParticipantInfo, repeated: true },
  {
    no: 4,
    name: "server_version",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "ice_servers", kind: "message", T: ICEServer, repeated: true },
  {
    no: 6,
    name: "subscriber_primary",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 7,
    name: "alternative_url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 8, name: "client_configuration", kind: "message", T: ClientConfiguration },
  {
    no: 9,
    name: "server_region",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 10,
    name: "ping_timeout",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 11,
    name: "ping_interval",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  { no: 12, name: "server_info", kind: "message", T: ServerInfo },
  {
    no: 13,
    name: "sif_trailer",
    kind: "scalar",
    T: 12
    /* ScalarType.BYTES */
  },
  { no: 14, name: "enabled_publish_codecs", kind: "message", T: Codec, repeated: true },
  {
    no: 15,
    name: "fast_publish",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let JoinResponse = _JoinResponse;
const _ReconnectResponse = class _ReconnectResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.ICEServer ice_servers = 1;
     */
    this.iceServers = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReconnectResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReconnectResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReconnectResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ReconnectResponse, a, b);
  }
};
_ReconnectResponse.runtime = proto3;
_ReconnectResponse.typeName = "livekit.ReconnectResponse";
_ReconnectResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "ice_servers", kind: "message", T: ICEServer, repeated: true },
  { no: 2, name: "client_configuration", kind: "message", T: ClientConfiguration }
]);
let ReconnectResponse = _ReconnectResponse;
const _TrackPublishedResponse = class _TrackPublishedResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string cid = 1;
     */
    this.cid = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TrackPublishedResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TrackPublishedResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TrackPublishedResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_TrackPublishedResponse, a, b);
  }
};
_TrackPublishedResponse.runtime = proto3;
_TrackPublishedResponse.typeName = "livekit.TrackPublishedResponse";
_TrackPublishedResponse.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "cid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "track", kind: "message", T: TrackInfo }
]);
let TrackPublishedResponse = _TrackPublishedResponse;
const _TrackUnpublishedResponse = class _TrackUnpublishedResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string track_sid = 1;
     */
    this.trackSid = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TrackUnpublishedResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TrackUnpublishedResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TrackUnpublishedResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_TrackUnpublishedResponse, a, b);
  }
};
_TrackUnpublishedResponse.runtime = proto3;
_TrackUnpublishedResponse.typeName = "livekit.TrackUnpublishedResponse";
_TrackUnpublishedResponse.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let TrackUnpublishedResponse = _TrackUnpublishedResponse;
const _SessionDescription = class _SessionDescription extends Message {
  constructor(data) {
    super();
    /**
     * "answer" | "offer" | "pranswer" | "rollback"
     *
     * @generated from field: string type = 1;
     */
    this.type = "";
    /**
     * @generated from field: string sdp = 2;
     */
    this.sdp = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SessionDescription().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SessionDescription().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SessionDescription().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SessionDescription, a, b);
  }
};
_SessionDescription.runtime = proto3;
_SessionDescription.typeName = "livekit.SessionDescription";
_SessionDescription.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "sdp",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let SessionDescription = _SessionDescription;
const _ParticipantUpdate = class _ParticipantUpdate extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.ParticipantInfo participants = 1;
     */
    this.participants = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ParticipantUpdate().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ParticipantUpdate().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ParticipantUpdate().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ParticipantUpdate, a, b);
  }
};
_ParticipantUpdate.runtime = proto3;
_ParticipantUpdate.typeName = "livekit.ParticipantUpdate";
_ParticipantUpdate.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "participants", kind: "message", T: ParticipantInfo, repeated: true }
]);
let ParticipantUpdate = _ParticipantUpdate;
const _UpdateSubscription = class _UpdateSubscription extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated string track_sids = 1;
     */
    this.trackSids = [];
    /**
     * @generated from field: bool subscribe = 2;
     */
    this.subscribe = false;
    /**
     * @generated from field: repeated livekit.ParticipantTracks participant_tracks = 3;
     */
    this.participantTracks = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateSubscription().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateSubscription().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateSubscription().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateSubscription, a, b);
  }
};
_UpdateSubscription.runtime = proto3;
_UpdateSubscription.typeName = "livekit.UpdateSubscription";
_UpdateSubscription.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "track_sids", kind: "scalar", T: 9, repeated: true },
  {
    no: 2,
    name: "subscribe",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 3, name: "participant_tracks", kind: "message", T: ParticipantTracks, repeated: true }
]);
let UpdateSubscription = _UpdateSubscription;
const _UpdateTrackSettings = class _UpdateTrackSettings extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated string track_sids = 1;
     */
    this.trackSids = [];
    /**
     * when true, the track is placed in a paused state, with no new data returned
     *
     * @generated from field: bool disabled = 3;
     */
    this.disabled = false;
    /**
     * deprecated in favor of width & height
     *
     * @generated from field: livekit.VideoQuality quality = 4;
     */
    this.quality = VideoQuality.LOW;
    /**
     * for video, width to receive
     *
     * @generated from field: uint32 width = 5;
     */
    this.width = 0;
    /**
     * for video, height to receive
     *
     * @generated from field: uint32 height = 6;
     */
    this.height = 0;
    /**
     * @generated from field: uint32 fps = 7;
     */
    this.fps = 0;
    /**
     * subscription priority. 1 being the highest (0 is unset)
     * when unset, server sill assign priority based on the order of subscription
     * server will use priority in the following ways:
     * 1. when subscribed tracks exceed per-participant subscription limit, server will
     *    pause the lowest priority tracks
     * 2. when the network is congested, server will assign available bandwidth to
     *    higher priority tracks first. lowest priority tracks can be paused
     *
     * @generated from field: uint32 priority = 8;
     */
    this.priority = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateTrackSettings().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateTrackSettings().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateTrackSettings().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateTrackSettings, a, b);
  }
};
_UpdateTrackSettings.runtime = proto3;
_UpdateTrackSettings.typeName = "livekit.UpdateTrackSettings";
_UpdateTrackSettings.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "track_sids", kind: "scalar", T: 9, repeated: true },
  {
    no: 3,
    name: "disabled",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 4, name: "quality", kind: "enum", T: proto3.getEnumType(VideoQuality) },
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
    name: "fps",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 8,
    name: "priority",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let UpdateTrackSettings = _UpdateTrackSettings;
const _UpdateLocalAudioTrack = class _UpdateLocalAudioTrack extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string track_sid = 1;
     */
    this.trackSid = "";
    /**
     * @generated from field: repeated livekit.AudioTrackFeature features = 2;
     */
    this.features = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateLocalAudioTrack().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateLocalAudioTrack().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateLocalAudioTrack().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateLocalAudioTrack, a, b);
  }
};
_UpdateLocalAudioTrack.runtime = proto3;
_UpdateLocalAudioTrack.typeName = "livekit.UpdateLocalAudioTrack";
_UpdateLocalAudioTrack.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "features", kind: "enum", T: proto3.getEnumType(AudioTrackFeature), repeated: true }
]);
let UpdateLocalAudioTrack = _UpdateLocalAudioTrack;
const _UpdateLocalVideoTrack = class _UpdateLocalVideoTrack extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string track_sid = 1;
     */
    this.trackSid = "";
    /**
     * @generated from field: uint32 width = 2;
     */
    this.width = 0;
    /**
     * @generated from field: uint32 height = 3;
     */
    this.height = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateLocalVideoTrack().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateLocalVideoTrack().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateLocalVideoTrack().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateLocalVideoTrack, a, b);
  }
};
_UpdateLocalVideoTrack.runtime = proto3;
_UpdateLocalVideoTrack.typeName = "livekit.UpdateLocalVideoTrack";
_UpdateLocalVideoTrack.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
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
  }
]);
let UpdateLocalVideoTrack = _UpdateLocalVideoTrack;
const _LeaveRequest = class _LeaveRequest extends Message {
  constructor(data) {
    super();
    /**
     * sent when server initiates the disconnect due to server-restart
     * indicates clients should attempt full-reconnect sequence
     * NOTE: `can_reconnect` obsoleted by `action` starting in protocol version 13
     *
     * @generated from field: bool can_reconnect = 1;
     */
    this.canReconnect = false;
    /**
     * @generated from field: livekit.DisconnectReason reason = 2;
     */
    this.reason = DisconnectReason.UNKNOWN_REASON;
    /**
     * @generated from field: livekit.LeaveRequest.Action action = 3;
     */
    this.action = 0 /* DISCONNECT */;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _LeaveRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _LeaveRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _LeaveRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_LeaveRequest, a, b);
  }
};
_LeaveRequest.runtime = proto3;
_LeaveRequest.typeName = "livekit.LeaveRequest";
_LeaveRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "can_reconnect",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 2, name: "reason", kind: "enum", T: proto3.getEnumType(DisconnectReason) },
  { no: 3, name: "action", kind: "enum", T: proto3.getEnumType(LeaveRequest_Action) },
  { no: 4, name: "regions", kind: "message", T: RegionSettings }
]);
let LeaveRequest = _LeaveRequest;
var LeaveRequest_Action = /* @__PURE__ */ ((LeaveRequest_Action2) => {
  LeaveRequest_Action2[LeaveRequest_Action2["DISCONNECT"] = 0] = "DISCONNECT";
  LeaveRequest_Action2[LeaveRequest_Action2["RESUME"] = 1] = "RESUME";
  LeaveRequest_Action2[LeaveRequest_Action2["RECONNECT"] = 2] = "RECONNECT";
  return LeaveRequest_Action2;
})(LeaveRequest_Action || {});
proto3.util.setEnumType(LeaveRequest_Action, "livekit.LeaveRequest.Action", [
  { no: 0, name: "DISCONNECT" },
  { no: 1, name: "RESUME" },
  { no: 2, name: "RECONNECT" }
]);
const _UpdateVideoLayers = class _UpdateVideoLayers extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string track_sid = 1;
     */
    this.trackSid = "";
    /**
     * @generated from field: repeated livekit.VideoLayer layers = 2;
     */
    this.layers = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateVideoLayers().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateVideoLayers().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateVideoLayers().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateVideoLayers, a, b);
  }
};
_UpdateVideoLayers.runtime = proto3;
_UpdateVideoLayers.typeName = "livekit.UpdateVideoLayers";
_UpdateVideoLayers.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "layers", kind: "message", T: VideoLayer, repeated: true }
]);
let UpdateVideoLayers = _UpdateVideoLayers;
const _UpdateParticipantMetadata = class _UpdateParticipantMetadata extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string metadata = 1;
     */
    this.metadata = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * attributes to update. it only updates attributes that have been set
     * to delete attributes, set the value to an empty string
     *
     * @generated from field: map<string, string> attributes = 3;
     */
    this.attributes = {};
    /**
     * @generated from field: uint32 request_id = 4;
     */
    this.requestId = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateParticipantMetadata().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateParticipantMetadata().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateParticipantMetadata().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateParticipantMetadata, a, b);
  }
};
_UpdateParticipantMetadata.runtime = proto3;
_UpdateParticipantMetadata.typeName = "livekit.UpdateParticipantMetadata";
_UpdateParticipantMetadata.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "metadata",
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
  { no: 3, name: "attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  {
    no: 4,
    name: "request_id",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let UpdateParticipantMetadata = _UpdateParticipantMetadata;
const _ICEServer = class _ICEServer extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated string urls = 1;
     */
    this.urls = [];
    /**
     * @generated from field: string username = 2;
     */
    this.username = "";
    /**
     * @generated from field: string credential = 3;
     */
    this.credential = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ICEServer().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ICEServer().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ICEServer().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ICEServer, a, b);
  }
};
_ICEServer.runtime = proto3;
_ICEServer.typeName = "livekit.ICEServer";
_ICEServer.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "urls", kind: "scalar", T: 9, repeated: true },
  {
    no: 2,
    name: "username",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "credential",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let ICEServer = _ICEServer;
const _SpeakersChanged = class _SpeakersChanged extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.SpeakerInfo speakers = 1;
     */
    this.speakers = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpeakersChanged().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpeakersChanged().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpeakersChanged().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SpeakersChanged, a, b);
  }
};
_SpeakersChanged.runtime = proto3;
_SpeakersChanged.typeName = "livekit.SpeakersChanged";
_SpeakersChanged.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "speakers", kind: "message", T: SpeakerInfo, repeated: true }
]);
let SpeakersChanged = _SpeakersChanged;
const _RoomUpdate = class _RoomUpdate extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RoomUpdate().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RoomUpdate().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RoomUpdate().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_RoomUpdate, a, b);
  }
};
_RoomUpdate.runtime = proto3;
_RoomUpdate.typeName = "livekit.RoomUpdate";
_RoomUpdate.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "room", kind: "message", T: Room }
]);
let RoomUpdate = _RoomUpdate;
const _ConnectionQualityInfo = class _ConnectionQualityInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string participant_sid = 1;
     */
    this.participantSid = "";
    /**
     * @generated from field: livekit.ConnectionQuality quality = 2;
     */
    this.quality = ConnectionQuality.POOR;
    /**
     * @generated from field: float score = 3;
     */
    this.score = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ConnectionQualityInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ConnectionQualityInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ConnectionQualityInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ConnectionQualityInfo, a, b);
  }
};
_ConnectionQualityInfo.runtime = proto3;
_ConnectionQualityInfo.typeName = "livekit.ConnectionQualityInfo";
_ConnectionQualityInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "participant_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "quality", kind: "enum", T: proto3.getEnumType(ConnectionQuality) },
  {
    no: 3,
    name: "score",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  }
]);
let ConnectionQualityInfo = _ConnectionQualityInfo;
const _ConnectionQualityUpdate = class _ConnectionQualityUpdate extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.ConnectionQualityInfo updates = 1;
     */
    this.updates = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ConnectionQualityUpdate().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ConnectionQualityUpdate().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ConnectionQualityUpdate().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ConnectionQualityUpdate, a, b);
  }
};
_ConnectionQualityUpdate.runtime = proto3;
_ConnectionQualityUpdate.typeName = "livekit.ConnectionQualityUpdate";
_ConnectionQualityUpdate.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "updates", kind: "message", T: ConnectionQualityInfo, repeated: true }
]);
let ConnectionQualityUpdate = _ConnectionQualityUpdate;
const _StreamStateInfo = class _StreamStateInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string participant_sid = 1;
     */
    this.participantSid = "";
    /**
     * @generated from field: string track_sid = 2;
     */
    this.trackSid = "";
    /**
     * @generated from field: livekit.StreamState state = 3;
     */
    this.state = 0 /* ACTIVE */;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _StreamStateInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _StreamStateInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _StreamStateInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_StreamStateInfo, a, b);
  }
};
_StreamStateInfo.runtime = proto3;
_StreamStateInfo.typeName = "livekit.StreamStateInfo";
_StreamStateInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "participant_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "state", kind: "enum", T: proto3.getEnumType(StreamState) }
]);
let StreamStateInfo = _StreamStateInfo;
const _StreamStateUpdate = class _StreamStateUpdate extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.StreamStateInfo stream_states = 1;
     */
    this.streamStates = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _StreamStateUpdate().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _StreamStateUpdate().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _StreamStateUpdate().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_StreamStateUpdate, a, b);
  }
};
_StreamStateUpdate.runtime = proto3;
_StreamStateUpdate.typeName = "livekit.StreamStateUpdate";
_StreamStateUpdate.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "stream_states", kind: "message", T: StreamStateInfo, repeated: true }
]);
let StreamStateUpdate = _StreamStateUpdate;
const _SubscribedQuality = class _SubscribedQuality extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.VideoQuality quality = 1;
     */
    this.quality = VideoQuality.LOW;
    /**
     * @generated from field: bool enabled = 2;
     */
    this.enabled = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SubscribedQuality().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SubscribedQuality().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SubscribedQuality().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SubscribedQuality, a, b);
  }
};
_SubscribedQuality.runtime = proto3;
_SubscribedQuality.typeName = "livekit.SubscribedQuality";
_SubscribedQuality.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "quality", kind: "enum", T: proto3.getEnumType(VideoQuality) },
  {
    no: 2,
    name: "enabled",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let SubscribedQuality = _SubscribedQuality;
const _SubscribedCodec = class _SubscribedCodec extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string codec = 1;
     */
    this.codec = "";
    /**
     * @generated from field: repeated livekit.SubscribedQuality qualities = 2;
     */
    this.qualities = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SubscribedCodec().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SubscribedCodec().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SubscribedCodec().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SubscribedCodec, a, b);
  }
};
_SubscribedCodec.runtime = proto3;
_SubscribedCodec.typeName = "livekit.SubscribedCodec";
_SubscribedCodec.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "codec",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "qualities", kind: "message", T: SubscribedQuality, repeated: true }
]);
let SubscribedCodec = _SubscribedCodec;
const _SubscribedQualityUpdate = class _SubscribedQualityUpdate extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string track_sid = 1;
     */
    this.trackSid = "";
    /**
     * @generated from field: repeated livekit.SubscribedQuality subscribed_qualities = 2;
     */
    this.subscribedQualities = [];
    /**
     * @generated from field: repeated livekit.SubscribedCodec subscribed_codecs = 3;
     */
    this.subscribedCodecs = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SubscribedQualityUpdate().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SubscribedQualityUpdate().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SubscribedQualityUpdate().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SubscribedQualityUpdate, a, b);
  }
};
_SubscribedQualityUpdate.runtime = proto3;
_SubscribedQualityUpdate.typeName = "livekit.SubscribedQualityUpdate";
_SubscribedQualityUpdate.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "subscribed_qualities", kind: "message", T: SubscribedQuality, repeated: true },
  { no: 3, name: "subscribed_codecs", kind: "message", T: SubscribedCodec, repeated: true }
]);
let SubscribedQualityUpdate = _SubscribedQualityUpdate;
const _TrackPermission = class _TrackPermission extends Message {
  constructor(data) {
    super();
    /**
     * permission could be granted either by participant sid or identity
     *
     * @generated from field: string participant_sid = 1;
     */
    this.participantSid = "";
    /**
     * @generated from field: bool all_tracks = 2;
     */
    this.allTracks = false;
    /**
     * @generated from field: repeated string track_sids = 3;
     */
    this.trackSids = [];
    /**
     * @generated from field: string participant_identity = 4;
     */
    this.participantIdentity = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TrackPermission().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TrackPermission().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TrackPermission().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_TrackPermission, a, b);
  }
};
_TrackPermission.runtime = proto3;
_TrackPermission.typeName = "livekit.TrackPermission";
_TrackPermission.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "participant_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "all_tracks",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 3, name: "track_sids", kind: "scalar", T: 9, repeated: true },
  {
    no: 4,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let TrackPermission = _TrackPermission;
const _SubscriptionPermission = class _SubscriptionPermission extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool all_participants = 1;
     */
    this.allParticipants = false;
    /**
     * @generated from field: repeated livekit.TrackPermission track_permissions = 2;
     */
    this.trackPermissions = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SubscriptionPermission().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SubscriptionPermission().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SubscriptionPermission().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SubscriptionPermission, a, b);
  }
};
_SubscriptionPermission.runtime = proto3;
_SubscriptionPermission.typeName = "livekit.SubscriptionPermission";
_SubscriptionPermission.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "all_participants",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 2, name: "track_permissions", kind: "message", T: TrackPermission, repeated: true }
]);
let SubscriptionPermission = _SubscriptionPermission;
const _SubscriptionPermissionUpdate = class _SubscriptionPermissionUpdate extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string participant_sid = 1;
     */
    this.participantSid = "";
    /**
     * @generated from field: string track_sid = 2;
     */
    this.trackSid = "";
    /**
     * @generated from field: bool allowed = 3;
     */
    this.allowed = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SubscriptionPermissionUpdate().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SubscriptionPermissionUpdate().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SubscriptionPermissionUpdate().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SubscriptionPermissionUpdate, a, b);
  }
};
_SubscriptionPermissionUpdate.runtime = proto3;
_SubscriptionPermissionUpdate.typeName = "livekit.SubscriptionPermissionUpdate";
_SubscriptionPermissionUpdate.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "participant_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "allowed",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let SubscriptionPermissionUpdate = _SubscriptionPermissionUpdate;
const _SyncState = class _SyncState extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.TrackPublishedResponse publish_tracks = 3;
     */
    this.publishTracks = [];
    /**
     * @generated from field: repeated livekit.DataChannelInfo data_channels = 4;
     */
    this.dataChannels = [];
    /**
     * @generated from field: repeated string track_sids_disabled = 6;
     */
    this.trackSidsDisabled = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SyncState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SyncState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SyncState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SyncState, a, b);
  }
};
_SyncState.runtime = proto3;
_SyncState.typeName = "livekit.SyncState";
_SyncState.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "answer", kind: "message", T: SessionDescription },
  { no: 2, name: "subscription", kind: "message", T: UpdateSubscription },
  { no: 3, name: "publish_tracks", kind: "message", T: TrackPublishedResponse, repeated: true },
  { no: 4, name: "data_channels", kind: "message", T: DataChannelInfo, repeated: true },
  { no: 5, name: "offer", kind: "message", T: SessionDescription },
  { no: 6, name: "track_sids_disabled", kind: "scalar", T: 9, repeated: true }
]);
let SyncState = _SyncState;
const _DataChannelInfo = class _DataChannelInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string label = 1;
     */
    this.label = "";
    /**
     * @generated from field: uint32 id = 2;
     */
    this.id = 0;
    /**
     * @generated from field: livekit.SignalTarget target = 3;
     */
    this.target = 0 /* PUBLISHER */;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DataChannelInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DataChannelInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DataChannelInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_DataChannelInfo, a, b);
  }
};
_DataChannelInfo.runtime = proto3;
_DataChannelInfo.typeName = "livekit.DataChannelInfo";
_DataChannelInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "label",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "id",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 3, name: "target", kind: "enum", T: proto3.getEnumType(SignalTarget) }
]);
let DataChannelInfo = _DataChannelInfo;
const _SimulateScenario = class _SimulateScenario extends Message {
  constructor(data) {
    super();
    /**
     * @generated from oneof livekit.SimulateScenario.scenario
     */
    this.scenario = { case: void 0 };
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SimulateScenario().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SimulateScenario().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SimulateScenario().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SimulateScenario, a, b);
  }
};
_SimulateScenario.runtime = proto3;
_SimulateScenario.typeName = "livekit.SimulateScenario";
_SimulateScenario.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "speaker_update", kind: "scalar", T: 5, oneof: "scenario" },
  { no: 2, name: "node_failure", kind: "scalar", T: 8, oneof: "scenario" },
  { no: 3, name: "migration", kind: "scalar", T: 8, oneof: "scenario" },
  { no: 4, name: "server_leave", kind: "scalar", T: 8, oneof: "scenario" },
  { no: 5, name: "switch_candidate_protocol", kind: "enum", T: proto3.getEnumType(CandidateProtocol), oneof: "scenario" },
  { no: 6, name: "subscriber_bandwidth", kind: "scalar", T: 3, oneof: "scenario" },
  { no: 7, name: "disconnect_signal_on_resume", kind: "scalar", T: 8, oneof: "scenario" },
  { no: 8, name: "disconnect_signal_on_resume_no_messages", kind: "scalar", T: 8, oneof: "scenario" },
  { no: 9, name: "leave_request_full_reconnect", kind: "scalar", T: 8, oneof: "scenario" }
]);
let SimulateScenario = _SimulateScenario;
const _Ping = class _Ping extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: int64 timestamp = 1;
     */
    this.timestamp = protoInt64.zero;
    /**
     * rtt in milliseconds calculated by client
     *
     * @generated from field: int64 rtt = 2;
     */
    this.rtt = protoInt64.zero;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Ping().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Ping().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Ping().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_Ping, a, b);
  }
};
_Ping.runtime = proto3;
_Ping.typeName = "livekit.Ping";
_Ping.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "timestamp",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 2,
    name: "rtt",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  }
]);
let Ping = _Ping;
const _Pong = class _Pong extends Message {
  constructor(data) {
    super();
    /**
     * timestamp field of last received ping request
     *
     * @generated from field: int64 last_ping_timestamp = 1;
     */
    this.lastPingTimestamp = protoInt64.zero;
    /**
     * @generated from field: int64 timestamp = 2;
     */
    this.timestamp = protoInt64.zero;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Pong().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Pong().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Pong().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_Pong, a, b);
  }
};
_Pong.runtime = proto3;
_Pong.typeName = "livekit.Pong";
_Pong.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "last_ping_timestamp",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 2,
    name: "timestamp",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  }
]);
let Pong = _Pong;
const _RegionSettings = class _RegionSettings extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.RegionInfo regions = 1;
     */
    this.regions = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RegionSettings().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RegionSettings().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RegionSettings().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_RegionSettings, a, b);
  }
};
_RegionSettings.runtime = proto3;
_RegionSettings.typeName = "livekit.RegionSettings";
_RegionSettings.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "regions", kind: "message", T: RegionInfo, repeated: true }
]);
let RegionSettings = _RegionSettings;
const _RegionInfo = class _RegionInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string region = 1;
     */
    this.region = "";
    /**
     * @generated from field: string url = 2;
     */
    this.url = "";
    /**
     * @generated from field: int64 distance = 3;
     */
    this.distance = protoInt64.zero;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RegionInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RegionInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RegionInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_RegionInfo, a, b);
  }
};
_RegionInfo.runtime = proto3;
_RegionInfo.typeName = "livekit.RegionInfo";
_RegionInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "region",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "distance",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  }
]);
let RegionInfo = _RegionInfo;
const _SubscriptionResponse = class _SubscriptionResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string track_sid = 1;
     */
    this.trackSid = "";
    /**
     * @generated from field: livekit.SubscriptionError err = 2;
     */
    this.err = SubscriptionError.SE_UNKNOWN;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SubscriptionResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SubscriptionResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SubscriptionResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SubscriptionResponse, a, b);
  }
};
_SubscriptionResponse.runtime = proto3;
_SubscriptionResponse.typeName = "livekit.SubscriptionResponse";
_SubscriptionResponse.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "err", kind: "enum", T: proto3.getEnumType(SubscriptionError) }
]);
let SubscriptionResponse = _SubscriptionResponse;
const _RequestResponse = class _RequestResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: uint32 request_id = 1;
     */
    this.requestId = 0;
    /**
     * @generated from field: livekit.RequestResponse.Reason reason = 2;
     */
    this.reason = 0 /* OK */;
    /**
     * @generated from field: string message = 3;
     */
    this.message = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RequestResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RequestResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RequestResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_RequestResponse, a, b);
  }
};
_RequestResponse.runtime = proto3;
_RequestResponse.typeName = "livekit.RequestResponse";
_RequestResponse.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "request_id",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 2, name: "reason", kind: "enum", T: proto3.getEnumType(RequestResponse_Reason) },
  {
    no: 3,
    name: "message",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let RequestResponse = _RequestResponse;
var RequestResponse_Reason = /* @__PURE__ */ ((RequestResponse_Reason2) => {
  RequestResponse_Reason2[RequestResponse_Reason2["OK"] = 0] = "OK";
  RequestResponse_Reason2[RequestResponse_Reason2["NOT_FOUND"] = 1] = "NOT_FOUND";
  RequestResponse_Reason2[RequestResponse_Reason2["NOT_ALLOWED"] = 2] = "NOT_ALLOWED";
  RequestResponse_Reason2[RequestResponse_Reason2["LIMIT_EXCEEDED"] = 3] = "LIMIT_EXCEEDED";
  return RequestResponse_Reason2;
})(RequestResponse_Reason || {});
proto3.util.setEnumType(RequestResponse_Reason, "livekit.RequestResponse.Reason", [
  { no: 0, name: "OK" },
  { no: 1, name: "NOT_FOUND" },
  { no: 2, name: "NOT_ALLOWED" },
  { no: 3, name: "LIMIT_EXCEEDED" }
]);
const _TrackSubscribed = class _TrackSubscribed extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string track_sid = 1;
     */
    this.trackSid = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TrackSubscribed().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TrackSubscribed().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TrackSubscribed().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_TrackSubscribed, a, b);
  }
};
_TrackSubscribed.runtime = proto3;
_TrackSubscribed.typeName = "livekit.TrackSubscribed";
_TrackSubscribed.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let TrackSubscribed = _TrackSubscribed;
export {
  AddTrackRequest,
  CandidateProtocol,
  ConnectionQualityInfo,
  ConnectionQualityUpdate,
  DataChannelInfo,
  ICEServer,
  JoinResponse,
  LeaveRequest,
  LeaveRequest_Action,
  MuteTrackRequest,
  ParticipantUpdate,
  Ping,
  Pong,
  ReconnectResponse,
  RegionInfo,
  RegionSettings,
  RequestResponse,
  RequestResponse_Reason,
  RoomUpdate,
  SessionDescription,
  SignalRequest,
  SignalResponse,
  SignalTarget,
  SimulateScenario,
  SimulcastCodec,
  SpeakersChanged,
  StreamState,
  StreamStateInfo,
  StreamStateUpdate,
  SubscribedCodec,
  SubscribedQuality,
  SubscribedQualityUpdate,
  SubscriptionPermission,
  SubscriptionPermissionUpdate,
  SubscriptionResponse,
  SyncState,
  TrackPermission,
  TrackPublishedResponse,
  TrackSubscribed,
  TrackUnpublishedResponse,
  TrickleRequest,
  UpdateLocalAudioTrack,
  UpdateLocalVideoTrack,
  UpdateParticipantMetadata,
  UpdateSubscription,
  UpdateTrackSettings,
  UpdateVideoLayers
};
//# sourceMappingURL=livekit_rtc_pb.js.map